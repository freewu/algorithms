package main

// 3719. Longest Balanced Subarray I
// You are given an integer array nums.

// A subarray is called balanced if the number of distinct even numbers in the subarray is equal to the number of distinct odd numbers.

// Return the length of the longest balanced subarray.

// A subarray is a contiguous non-empty sequence of elements within an array.
 
// Example 1:
// Input: nums = [2,5,4,3]
// Output: 4
// Explanation:
// The longest balanced subarray is [2, 5, 4, 3].
// It has 2 distinct even numbers [2, 4] and 2 distinct odd numbers [5, 3]. Thus, the answer is 4.

// Example 2:
// Input: nums = [3,2,2,5,4]
// Output: 5
// Explanation:
// The longest balanced subarray is [3, 2, 2, 5, 4].
// It has 2 distinct even numbers [2, 4] and 2 distinct odd numbers [3, 5]. Thus, the answer is 5.

// Example 3:
// Input: nums = [1,2,3,2]
// Output: 3
// Explanation:
// The longest balanced subarray is [2, 3, 2].
// It has 1 distinct even number [2] and 1 distinct odd number [3]. Thus, the answer is 3.
 
// Constraints:
//     1 <= nums.length <= 1500
//     1 <= nums[i] <= 10^5

import "fmt"
import "math"

func longestBalanced(nums []int) int {
    res, n := 0,len(nums)
    m := int(math.Sqrt(float64(n+1)))/2 + 1
    sum := make([]int, n + 1)
    // === 分块模板开始 ===
    // 用分块维护 sum
    type Block struct {
        l, r int // [l,r) 左闭右开
        todo int
        pos  map[int]int
    }
    blocks := make([]Block, n / m + 1)
    calcPos := func(l, r int) map[int]int {
        pos := map[int]int{}
        for j := r - 1; j >= l; j-- {
            pos[sum[j]] = j
        }
        return pos
    }
    for i := 0; i <= n; i += m {
        r := min(i+m, n+1)
        pos := calcPos(i, r)
        blocks[i/m] = Block{i, r, 0, pos}
    }
    // sum[l:r] 增加 v
    rangeAdd := func(l, r, v int) {
        for i := range blocks {
            b := &blocks[i]
            if b.r <= l { continue }
            if b.l >= r { break }
            if l <= b.l && b.r <= r { // 完整块
                b.todo += v
            } else { // 部分块，直接重算
                for j := b.l; j < b.r; j++ {
                    sum[j] += b.todo
                    if l <= j && j < r {
                        sum[j] += v
                    }
                }
                b.pos = calcPos(b.l, b.r)
                b.todo = 0
            }
        }
    }
    // 返回 sum[:r] 中第一个 v 的下标
    // 如果没有 v，返回 n
    findFirst := func(r, v int) int {
        for i := range blocks {
            b := &blocks[i]
            if b.r <= r { // 完整块，直接查哈希表
                if j, ok := b.pos[v - b.todo]; ok {
                    return j
                }
            } else { // 部分块，暴力查找
                for j := b.l; j < r; j++ {
                    if sum[j] == v-b.todo {
                        return j
                    }
                }
                break
            }
        }
        return n
    }
    // === 分块模板结束 ===
    last := map[int]int{} // nums 的元素上一次出现的位置
    for i := 1; i <= n; i++ {
        x := nums[i-1]
        v := x % 2 * 2 - 1
        if j := last[x]; j == 0 { // 首次遇到 x
            rangeAdd(i, n+1, v) // sum[i:] 增加 v
        } else { // 再次遇到 x
            rangeAdd(j, i, -v) // 撤销之前对 sum[j:i] 的增加
        }
        last[x] = i
        s := sum[i] + blocks[i/m].todo // sum[i] 的实际值
        res = max(res, i - findFirst(i - res, s)) // 优化右边界
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,5,4,3]
    // Output: 4
    // Explanation:
    // The longest balanced subarray is [2, 5, 4, 3].
    // It has 2 distinct even numbers [2, 4] and 2 distinct odd numbers [5, 3]. Thus, the answer is 4.
    fmt.Println(longestBalanced([]int{2,5,4,3})) // 4
    // Example 2:
    // Input: nums = [3,2,2,5,4]
    // Output: 5
    // Explanation:
    // The longest balanced subarray is [3, 2, 2, 5, 4].
    // It has 2 distinct even numbers [2, 4] and 2 distinct odd numbers [3, 5]. Thus, the answer is 5.
    fmt.Println(longestBalanced([]int{3,2,2,5,4})) // 5
    // Example 3:
    // Input: nums = [1,2,3,2]
    // Output: 3
    // Explanation:
    // The longest balanced subarray is [2, 3, 2].
    // It has 1 distinct even number [2] and 1 distinct odd number [3]. Thus, the answer is 3.
    fmt.Println(longestBalanced([]int{1,2,3,2})) // 3

    fmt.Println(longestBalanced([]int{1,2,3,4,5,6,7,8,9})) // 8
    fmt.Println(longestBalanced([]int{9,8,7,6,5,4,3,2,1})) // 8
}