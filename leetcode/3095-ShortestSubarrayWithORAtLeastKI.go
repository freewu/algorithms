package main

// 3095. Shortest Subarray With OR at Least K I
// You are given an array nums of non-negative integers and an integer k.

// An array is called special if the bitwise OR of all of its elements is at least k.

// Return the length of the shortest special non-empty subarray of nums, or return -1 if no special subarray exists.

// Example 1:
// Input: nums = [1,2,3], k = 2
// Output: 1
// Explanation:
// The subarray [3] has OR value of 3. Hence, we return 1.
// Note that [2] is also a special subarray.

// Example 2:
// Input: nums = [2,1,8], k = 10
// Output: 3
// Explanation:
// The subarray [2,1,8] has OR value of 11. Hence, we return 3.

// Example 3:
// Input: nums = [1,2], k = 0
// Output: 1
// Explanation:
// The subarray [1] has OR value of 1. Hence, we return 1.

// Constraints:
//     1 <= nums.length <= 50
//     0 <= nums[i] <= 50
//     0 <= k < 64

import "fmt"

// brute force 
func minimumSubarrayLength(nums []int, k int) int {
    res, n := 1 << 31, len(nums)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {
            or := 0
            for k := i; k <= j; k++ {
                or = or | nums[k]
            }
            if or >= k {
                res = min(res, j - i + 1)
            }
        }
    }
    if res == 1 << 31 { return -1 }
    return res
}

func minimumSubarrayLength1(nums []int, k int) int {
    // 由于nums确保至少有一个元素，并且k为0，那么任一一个元素都可以自己组成长度为1的数组，使得满足或值>=0的条件。
    if k == 0 { return 1 }
    res := 1 << 31
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i, v := range nums {
        if v >= k { // 如果有元素直接满足条件，则自己组成结果数组
            return 1
        }
        for j := 0; j <= i; j++ {
            nums[j] = nums[i] | nums[j]
            if nums[j] >= k {
                // 这轮结果窗口是从nums[j]开始到nums[i]结束
                res = min(res, i - j + 1)
            }
        }
    }
    if res == 1 << 31 { return -1 }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3], k = 2
    // Output: 1
    // Explanation:
    // The subarray [3] has OR value of 3. Hence, we return 1.
    // Note that [2] is also a special subarray.
    fmt.Println(minimumSubarrayLength([]int{1,2,3}, 2)) // 1
    // Example 2:
    // Input: nums = [2,1,8], k = 10
    // Output: 3
    // Explanation:
    // The subarray [2,1,8] has OR value of 11. Hence, we return 3.
    fmt.Println(minimumSubarrayLength([]int{2,1,8}, 10)) // 3
    // Example 3:
    // Input: nums = [1,2], k = 0
    // Output: 1
    // Explanation:
    // The subarray [1] has OR value of 1. Hence, we return 1.
    fmt.Println(minimumSubarrayLength([]int{1, 2}, 0)) // 1

    fmt.Println(minimumSubarrayLength1([]int{1,2,3}, 2)) // 1
    fmt.Println(minimumSubarrayLength1([]int{2,1,8}, 10)) // 3
    fmt.Println(minimumSubarrayLength1([]int{1, 2}, 0)) // 1
}