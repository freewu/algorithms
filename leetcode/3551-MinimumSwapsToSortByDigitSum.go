package main

// 3551. Minimum Swaps to Sort by Digit Sum
// You are given an array nums of distinct positive integers. 
// You need to sort the array in increasing order based on the sum of the digits of each number. 
// If two numbers have the same digit sum, the smaller number appears first in the sorted order.

// Return the minimum number of swaps required to rearrange nums into this sorted order.

// A swap is defined as exchanging the values at two distinct positions in the array.

// Example 1:
// Input: nums = [37,100]
// Output: 1
// Explanation:
// Compute the digit sum for each integer: [3 + 7 = 10, 1 + 0 + 0 = 1] → [10, 1]
// Sort the integers based on digit sum: [100, 37]. Swap 37 with 100 to obtain the sorted order.
// Thus, the minimum number of swaps required to rearrange nums is 1.

// Example 2:
// Input: nums = [22,14,33,7]
// Output: 0
// Explanation:
// Compute the digit sum for each integer: [2 + 2 = 4, 1 + 4 = 5, 3 + 3 = 6, 7 = 7] → [4, 5, 6, 7]
// Sort the integers based on digit sum: [22, 14, 33, 7]. The array is already sorted.
// Thus, the minimum number of swaps required to rearrange nums is 0.

// Example 3:
// Input: nums = [18,43,34,16]
// Output: 2
// Explanation:
// Compute the digit sum for each integer: [1 + 8 = 9, 4 + 3 = 7, 3 + 4 = 7, 1 + 6 = 7] → [9, 7, 7, 7]
// Sort the integers based on digit sum: [16, 34, 43, 18]. Swap 18 with 16, and swap 43 with 34 to obtain the sorted order.
// Thus, the minimum number of swaps required to rearrange nums is 2.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     nums consists of distinct positive integers.

import "fmt"
import "sort"

func minSwaps(nums []int) int {
    res, n := 0, len(nums)
    type Elem struct { val, sum, index int } // 原数值, 数字和, 原始索引
    arr := make([]Elem, n)
    digitSum := func(num int) int { // 计算数字各位之和
        sum := 0
        for num > 0 {
            sum += num % 10 // 取最后一位相加
            num /= 10       // 去掉最后一位
        }
        return sum
    }
    for i, v := range nums { // 预处理
        arr[i] = Elem{ v, digitSum(v), i }
    }
    // 排序：先按数字和升序，相同则按原数值升序
    sort.Slice(arr, func(i, j int) bool {
        if arr[i].sum != arr[j].sum {
            return arr[i].sum < arr[j].sum
        }
        return arr[i].val < arr[j].val
    })
    visited := make([]bool, n) // 标记是否已访问
    for i := 0; i < n; i++ {
        if visited[i] || arr[i].index == i { continue } // 如果已访问或已在正确位置，跳过
        // 开始检测环
        cycleSize, j := 0, i
        for !visited[j] {
            visited[j] = true  // 标记为已访问
            j = arr[j].index   // 移动到它应该在的位置
            cycleSize++        // 环大小+1
        }
        if cycleSize > 0 { // 每个大小为 k 的环需要 k-1 次交换
            res += (cycleSize - 1)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [37,100]
    // Output: 1
    // Explanation:
    // Compute the digit sum for each integer: [3 + 7 = 10, 1 + 0 + 0 = 1] → [10, 1]
    // Sort the integers based on digit sum: [100, 37]. Swap 37 with 100 to obtain the sorted order.
    // Thus, the minimum number of swaps required to rearrange nums is 1.
    fmt.Println(minSwaps([]int{37,100})) // 1
    // Example 2:
    // Input: nums = [22,14,33,7]
    // Output: 0
    // Explanation:
    // Compute the digit sum for each integer: [2 + 2 = 4, 1 + 4 = 5, 3 + 3 = 6, 7 = 7] → [4, 5, 6, 7]
    // Sort the integers based on digit sum: [22, 14, 33, 7]. The array is already sorted.
    // Thus, the minimum number of swaps required to rearrange nums is 0.
    fmt.Println(minSwaps([]int{22,14,33,7})) // 0
    // Example 3:
    // Input: nums = [18,43,34,16]
    // Output: 2
    // Explanation:
    // Compute the digit sum for each integer: [1 + 8 = 9, 4 + 3 = 7, 3 + 4 = 7, 1 + 6 = 7] → [9, 7, 7, 7]
    // Sort the integers based on digit sum: [16, 34, 43, 18]. Swap 18 with 16, and swap 43 with 34 to obtain the sorted order.
    // Thus, the minimum number of swaps required to rearrange nums is 2.
    fmt.Println(minSwaps([]int{18,43,34,16})) // 2

    fmt.Println(minSwaps([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minSwaps([]int{9,8,7,6,5,4,3,2,1})) // 4
}