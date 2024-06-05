package main

// 453. Minimum Moves to Equal Array Elements
// Given an integer array nums of size n, return the minimum number of moves required to make all array elements equal.
// In one move, you can increment n - 1 elements of the array by 1.

// Example 1:
// Input: nums = [1,2,3]
// Output: 3
// Explanation: Only three moves are needed (remember each move increments two elements):
// [1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]

// Example 2:
// Input: nums = [1,1,1]
// Output: 0
 
// Constraints:
//     n == nums.length
//     1 <= nums.length <= 10^5
//     -10^9 <= nums[i] <= 10^9
//     The answer is guaranteed to fit in a 32-bit integer.

import "fmt"

// 使得每个元素都相同，意思让所有元素的差异变为 0 。
// 每次移动的过程中，都有 n - 1 个元素 + 1，那么没有 + 1 的那个元素和其他 n - 1 个元素相对差异就缩小了。
// 所以这道题让所有元素都变为相等的最少步数，即等于让所有元素相对差异减少到最小的那个数
func minMoves(nums []int) int {
    sum, mn, l := 0, 1 << 32 -1, len(nums)
    for _, v := range nums {
        sum += v  // 累加数组的值
        if mn > v { // 得到数组中最小的值
            mn = v
        }
    }
    return sum - mn * l // 让所有元素相对差异减少到最小的那个数
}

func minMoves1(nums []int) int {
    res, mn := 0, nums[0]
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < len(nums); i++ {
        mn = min(mn, nums[i])
    }
    for i := 0; i < len(nums); i++ {
        res += nums[i] - mn
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3]
    // Output: 3
    // Explanation: Only three moves are needed (remember each move increments two elements):
    // [1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
    fmt.Println(minMoves([]int{1,2,3})) // 3
    // Example 2:
    // Input: nums = [1,1,1]
    // Output: 0
    fmt.Println(minMoves([]int{1,1,1})) // 0

    fmt.Println(minMoves([]int{1,2,3})) // 3
    fmt.Println(minMoves([]int{1,1,1})) // 0
}