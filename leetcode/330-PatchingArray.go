package main

// 330. Patching Array
// Given a sorted integer array nums and an integer n, 
// add/patch elements to the array such that any number in the range [1, n] inclusive can be formed by the sum of some elements in the array.
// Return the minimum number of patches required.

// Example 1:
// Input: nums = [1,3], n = 6
// Output: 1
// Explanation:
// Combinations of nums are [1], [3], [1,3], which form possible sums of: 1, 3, 4.
// Now if we add/patch 2 to nums, the combinations are: [1], [2], [3], [1,3], [2,3], [1,2,3].
// Possible sums are 1, 2, 3, 4, 5, 6, which now covers the range [1, 6].
// So we only need 1 patch.

// Example 2:
// Input: nums = [1,5,10], n = 20
// Output: 2
// Explanation: The two patches can be [2, 4].

// Example 3:
// Input: nums = [1,2,2], n = 5
// Output: 0
 
// Constraints:
//     1 <= nums.length <= 1000
//     1 <= nums[i] <= 10^4
//     nums is sorted in ascending order.
//     1 <= n <= 2^31 - 1

import "fmt"

func minPatches(nums []int, n int) int {
    res, covered, index := 0, 0, 0
    for index < len(nums) && covered < n {
        if nums[index] <= covered + 1{
            covered += nums[index]
            index++ 
        } else {
            covered += covered + 1
            res++
        }
    }
    for covered < n { // 补齐剩余的
        covered += covered + 1
        res++
    }
    return res
}

func minPatches1(nums []int, n int) int {
    res := 0
    for i, x := 0, 1; x <= n; {
        if i < len(nums) && nums[i] <= x {
            x += nums[i]
            i++
        } else {
            x *= 2
            res++
        }
    }
    return res
}

func minPatches2(nums []int, n int) int {
    res, missing, index := 0, 1, 0
    for missing <= n {
        if index < len(nums) && nums[index] <= missing {
            missing += nums[index]
            index++
        } else {
            missing += missing
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,3], n = 6
    // Output: 1
    // Explanation:
    // Combinations of nums are [1], [3], [1,3], which form possible sums of: 1, 3, 4.
    // Now if we add/patch 2 to nums, the combinations are: [1], [2], [3], [1,3], [2,3], [1,2,3].
    // Possible sums are 1, 2, 3, 4, 5, 6, which now covers the range [1, 6].
    // So we only need 1 patch.
    fmt.Println(minPatches([]int{1,3}, 6)) // 1
    // Example 2:
    // Input: nums = [1,5,10], n = 20
    // Output: 2
    // Explanation: The two patches can be [2, 4].
    fmt.Println(minPatches([]int{1,5,10}, 20)) // 2
    // Example 3:
    // Input: nums = [1,2,2], n = 5
    // Output: 0
    fmt.Println(minPatches([]int{1,2,2}, 5)) // 0

    fmt.Println(minPatches1([]int{1,3}, 6)) // 1
    fmt.Println(minPatches1([]int{1,5,10}, 20)) // 2
    fmt.Println(minPatches1([]int{1,2,2}, 5)) // 0

    fmt.Println(minPatches2([]int{1,3}, 6)) // 1
    fmt.Println(minPatches2([]int{1,5,10}, 20)) // 2
    fmt.Println(minPatches2([]int{1,2,2}, 5)) // 0
}