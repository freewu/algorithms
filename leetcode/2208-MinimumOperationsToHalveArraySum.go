package main

// 2208. Minimum Operations to Halve Array Sum
// You are given an array nums of positive integers. 
// In one operation, you can choose any number from nums and reduce it to exactly half the number. 
// (Note that you may choose this reduced number in future operations.)

// Return the minimum number of operations to reduce the sum of nums by at least half.

// Example 1:
// Input: nums = [5,19,8,1]
// Output: 3
// Explanation: The initial sum of nums is equal to 5 + 19 + 8 + 1 = 33.
// The following is one of the ways to reduce the sum by at least half:
// Pick the number 19 and reduce it to 9.5.
// Pick the number 9.5 and reduce it to 4.75.
// Pick the number 8 and reduce it to 4.
// The final array is [5, 4.75, 4, 1] with a total sum of 5 + 4.75 + 4 + 1 = 14.75. 
// The sum of nums has been reduced by 33 - 14.75 = 18.25, which is at least half of the initial sum, 18.25 >= 33/2 = 16.5.
// Overall, 3 operations were used so we return 3.
// It can be shown that we cannot reduce the sum by at least half in less than 3 operations.

// Example 2:
// Input: nums = [3,8,20]
// Output: 3
// Explanation: The initial sum of nums is equal to 3 + 8 + 20 = 31.
// The following is one of the ways to reduce the sum by at least half:
// Pick the number 20 and reduce it to 10.
// Pick the number 10 and reduce it to 5.
// Pick the number 3 and reduce it to 1.5.
// The final array is [1.5, 8, 5] with a total sum of 1.5 + 8 + 5 = 14.5. 
// The sum of nums has been reduced by 31 - 14.5 = 16.5, which is at least half of the initial sum, 16.5 >= 31/2 = 15.5.
// Overall, 3 operations were used so we return 3.
// It can be shown that we cannot reduce the sum by at least half in less than 3 operations.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^7

import "fmt"
import "sort"
import "container/heap"

type MaxHeap struct{ sort.Float64Slice }

func (h MaxHeap) Less(i, j int) bool { return h.Float64Slice[i] > h.Float64Slice[j] }
func (h *MaxHeap) Push(v any)        { h.Float64Slice = append(h.Float64Slice, v.(float64)) }
func (h *MaxHeap) Pop() any {
    a := h.Float64Slice
    v := a[len(a)-1]
    h.Float64Slice = a[:len(a)-1]
    return v
}

func halveArray(nums []int) int {
    res, sum := 0, float64(0)
    mxh := MaxHeap{}
    for _, v := range nums {
        sum += float64(v)
        heap.Push(&mxh, float64(v))
    }
    sum /= 2
    for sum > 0 {
        v := heap.Pop(&mxh).(float64)
        res++
        sum -= v / 2
        heap.Push(&mxh, v/2)
    }
    return res
}

func halveArray1(nums []int) int {
    res, n, sum := 0, len(nums), int64(0)
    memo := make([]int64, n)
    heapify := func(nums []int64, i int, n int) {  // 调整堆结构
        l := i * 2 + 1
        for l < n {
            best, r := l, i * 2+2 // 右孩子节点
            if r < n && nums[r] > nums[l] {
                best = r
            } // 找出最好的孩子节点
            if nums[best] <= nums[i] { break  }
            nums[i], nums[best] = nums[best], nums[i] // 和父节点对比 如果大的话交换位置
            i = best // 转到子节点 
            l = i * 2 + 1 // 继续看孩子的子节点
        }
    }
    for i := n - 1; i >=0 ; i-- {
        memo[i] = int64(nums[i]) << 20 // 放大2^20倍主要是为了减少浮点数带来的误差 \U0001f31f
        sum += int64(nums[i]) << 20
        heapify(memo, i, n) // 自底部向顶部建立大根堆 时间复杂度O(n)
    }
    sum /= 2 // 目标值
    for minus := int64(0); minus < sum; res++ { // minus表示gap值的和
        memo[0] /= 2
        minus += memo[0]
        heapify(memo, 0, n)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [5,19,8,1]
    // Output: 3
    // Explanation: The initial sum of nums is equal to 5 + 19 + 8 + 1 = 33.
    // The following is one of the ways to reduce the sum by at least half:
    // Pick the number 19 and reduce it to 9.5.
    // Pick the number 9.5 and reduce it to 4.75.
    // Pick the number 8 and reduce it to 4.
    // The final array is [5, 4.75, 4, 1] with a total sum of 5 + 4.75 + 4 + 1 = 14.75. 
    // The sum of nums has been reduced by 33 - 14.75 = 18.25, which is at least half of the initial sum, 18.25 >= 33/2 = 16.5.
    // Overall, 3 operations were used so we return 3.
    // It can be shown that we cannot reduce the sum by at least half in less than 3 operations.
    fmt.Println(halveArray([]int{5,19,8,1})) // 3
    // Example 2:
    // Input: nums = [3,8,20]
    // Output: 3
    // Explanation: The initial sum of nums is equal to 3 + 8 + 20 = 31.
    // The following is one of the ways to reduce the sum by at least half:
    // Pick the number 20 and reduce it to 10.
    // Pick the number 10 and reduce it to 5.
    // Pick the number 3 and reduce it to 1.5.
    // The final array is [1.5, 8, 5] with a total sum of 1.5 + 8 + 5 = 14.5. 
    // The sum of nums has been reduced by 31 - 14.5 = 16.5, which is at least half of the initial sum, 16.5 >= 31/2 = 15.5.
    // Overall, 3 operations were used so we return 3.
    // It can be shown that we cannot reduce the sum by at least half in less than 3 operations.
    fmt.Println(halveArray([]int{3,8,20})) // 3
    fmt.Println(halveArray([]int{6,58,10,84,35,8,22,64,1,78,86,71,77})) // 9

    fmt.Println(halveArray1([]int{5,19,8,1})) // 3
    fmt.Println(halveArray1([]int{3,8,20})) // 3
    fmt.Println(halveArray1([]int{6,58,10,84,35,8,22,64,1,78,86,71,77})) // 9
}