package main

// 945. Minimum Increment to Make Array Unique
// You are given an integer array nums. 
// In one move, you can pick an index i where 0 <= i < nums.length and increment nums[i] by 1.

// Return the minimum number of moves to make every value in nums unique.
// The test cases are generated so that the answer fits in a 32-bit integer.

// Example 1:
// Input: nums = [1,2,2]
// Output: 1
// Explanation: After 1 move, the array could be [1, 2, 3].

// Example 2:
// Input: nums = [3,2,1,2,1,7]
// Output: 6
// Explanation: After 6 moves, the array could be [3, 4, 1, 2, 5, 7].
// It can be shown with 5 or less moves that it is impossible for the array to have all unique values.
 
// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^5

import "fmt"

func minIncrementForUnique(nums []int) int {
    res, arr, mn, mx := 0, make([]int, 100001), 100001, -1
    for _, v := range nums { // 获取最大值 & 最小值
        if v > mx { mx = v }
        if v < mn { mn = v }
        arr[v]++
    }
    for i := mn; i <= mx; i++ {
        if i == mx  {
            return res + (((arr[i] - 1) * arr[i]) / 2)
        }
        if arr[i] > 0 {
            res += arr[i] - 1
            arr[i+1] += arr[i]-1
        }
    } 
    return -1
}

func main() {
    // Example 1:
    // Input: nums = [1,2,2]
    // Output: 1
    // Explanation: After 1 move, the array could be [1, 2, 3].
    fmt.Println(minIncrementForUnique([]int{1,2,2})) // 1
    // Example 2:
    // Input: nums = [3,2,1,2,1,7]
    // Output: 6
    // Explanation: After 6 moves, the array could be [3, 4, 1, 2, 5, 7].
    // It can be shown with 5 or less moves that it is impossible for the array to have all unique values.
    fmt.Println(minIncrementForUnique([]int{3,2,1,2,1,7})) // 6
}