package main

// 2855. Minimum Right Shifts to Sort the Array
// You are given a 0-indexed array nums of length n containing distinct positive integers. 
// Return the minimum number of right shifts required to sort nums and -1 if this is not possible.

// A right shift is defined as shifting the element at index i to index (i + 1) % n, for all indices.

// Example 1:
// Input: nums = [3,4,5,1,2]
// Output: 2
// Explanation: 
// After the first right shift, nums = [2,3,4,5,1].
// After the second right shift, nums = [1,2,3,4,5].
// Now nums is sorted; therefore the answer is 2.

// Example 2:
// Input: nums = [1,3,5]
// Output: 0
// Explanation: nums is already sorted therefore, the answer is 0.

// Example 3:
// Input: nums = [2,1,4]
// Output: -1
// Explanation: It's impossible to sort the array using right shifts.

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 100
//     nums contains distinct integers.

import "fmt"

func minimumRightShifts(nums []int) int {
    res, n, count := 0, len(nums), 0; // How many breaks , if breaks are > 1 return -1;
    for i := 0; i < n - 1; i++ {
        if nums[i] > nums[i + 1] {
            if res == 0 {
                res = n - i - 1
            }
            count++
        }
    }
    if count > 0 {
        if count > 1 || nums[n - 1] > nums[0] {
            return -1
        } else {
            return res
        }
    }
    return res
}

func minimumRightShifts1(nums []int) int {
    n := len(nums)
    if n == 1 { return 0 }
    res, count := 0, 0
    for i := n - 1; i > 0; i-- {
        if nums[i] < nums[i - 1] {
            if count > 0 {
                return -1
            }
            res = n - i
            count++
        }
    }
    if n > 2 && count > 0 && nums[n - 1] > nums[0] {
        return -1
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,4,5,1,2]
    // Output: 2
    // Explanation: 
    // After the first right shift, nums = [2,3,4,5,1].
    // After the second right shift, nums = [1,2,3,4,5].
    // Now nums is sorted; therefore the answer is 2.
    fmt.Println(minimumRightShifts([]int{3,4,5,1,2})) // 2
    // Example 2:
    // Input: nums = [1,3,5]
    // Output: 0
    // Explanation: nums is already sorted therefore, the answer is 0.
    fmt.Println(minimumRightShifts([]int{1,3,5})) // 0
    // Example 3:
    // Input: nums = [2,1,4]
    // Output: -1
    // Explanation: It's impossible to sort the array using right shifts.
    fmt.Println(minimumRightShifts([]int{2,1,4})) // -1

    fmt.Println(minimumRightShifts([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minimumRightShifts([]int{9,8,7,6,5,4,3,2,1})) // -1

    fmt.Println(minimumRightShifts1([]int{3,4,5,1,2})) // 2
    fmt.Println(minimumRightShifts1([]int{1,3,5})) // 0
    fmt.Println(minimumRightShifts1([]int{2,1,4})) // -1
    fmt.Println(minimumRightShifts1([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minimumRightShifts1([]int{9,8,7,6,5,4,3,2,1})) // -1
}