package main

// 2104. Sum of Subarray Ranges
// You are given an integer array nums. 
// The range of a subarray of nums is the difference between the largest and smallest element in the subarray.

// Return the sum of all subarray ranges of nums.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [1,2,3]
// Output: 4
// Explanation: The 6 subarrays of nums are the following:
// [1], range = largest - smallest = 1 - 1 = 0 
// [2], range = 2 - 2 = 0
// [3], range = 3 - 3 = 0
// [1,2], range = 2 - 1 = 1
// [2,3], range = 3 - 2 = 1
// [1,2,3], range = 3 - 1 = 2
// So the sum of all ranges is 0 + 0 + 0 + 1 + 1 + 2 = 4.

// Example 2:
// Input: nums = [1,3,3]
// Output: 4
// Explanation: The 6 subarrays of nums are the following:
// [1], range = largest - smallest = 1 - 1 = 0
// [3], range = 3 - 3 = 0
// [3], range = 3 - 3 = 0
// [1,3], range = 3 - 1 = 2
// [3,3], range = 3 - 3 = 0
// [1,3,3], range = 3 - 1 = 2
// So the sum of all ranges is 0 + 0 + 0 + 2 + 0 + 2 = 4.

// Example 3:
// Input: nums = [4,-2,-3,4,1]
// Output: 59
// Explanation: The sum of all subarray ranges of nums is 59.

// Constraints:
//     1 <= nums.length <= 1000
//     -10^9 <= nums[i] <= 10^9

// Follow-up: Could you find a solution with O(n) time complexity?

import "fmt"

// monotic stack
func subArrayRanges(nums []int) int64 {
    res, n, stack := int64(0), len(nums), []int{}
    for right := 0; right <= n; right++ {
        for len(stack) > 0 && (right == n || nums[stack[len(stack)-1]] >= nums[right]) {
            mn := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            left := -1
            if len(stack) > 0 {
                left = stack[len(stack)-1]
            }
            res -= int64(nums[mn] * (mn - left) * (right - mn))
        }
        stack = append(stack, right)
    }
    stack = []int{}
    for right := 0; right <= n; right++ {
        for len(stack) > 0 && (right == n || nums[stack[len(stack)-1]] <= nums[right]) {
            mn := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            left := -1
            if len(stack) > 0 {
                left = stack[len(stack)-1]
            }
            res += int64(nums[mn] * (mn - left) * (right - mn))
        }
        stack = append(stack, right)
    }
    return res
}

func subArrayRanges1(nums []int) int64 {
    calcMaxSum := func(nums []int) int {
        nums = append(nums, 1 << 31)
        sum, s := 0, []int{ -1 }
        for r, v := range nums {
            for len(s) > 1 && nums[s[len(s)-1]]  <= v {
                i := s[len(s)-1]
                s = s[:len(s)-1]
                l := s[len(s)-1]
                
                sum += nums[i] * (i-l) * (r-i)
            }
            s = append(s, r)
        }
        return sum
    }
    calcMinSum := func(nums []int) int {
        nums = append(nums, -1 << 31)
        sum, s := 0, []int{ -1 }
        for r, v := range nums {
            for len(s) > 1 && nums[s[len(s)-1]]  >= v {
                i := s[len(s)-1]
                s = s[:len(s)-1]
                l := s[len(s)-1]
                sum += nums[i] * (i-l) * (r-i)
            }
            s = append(s, r)
        }
        return sum
    }
    return int64(calcMaxSum(nums) - calcMinSum(nums))
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3]
    // Output: 4
    // Explanation: The 6 subarrays of nums are the following:
    // [1], range = largest - smallest = 1 - 1 = 0 
    // [2], range = 2 - 2 = 0
    // [3], range = 3 - 3 = 0
    // [1,2], range = 2 - 1 = 1
    // [2,3], range = 3 - 2 = 1
    // [1,2,3], range = 3 - 1 = 2
    // So the sum of all ranges is 0 + 0 + 0 + 1 + 1 + 2 = 4.
    fmt.Println(subArrayRanges([]int{1,2,3})) // 4
    // Example 2:
    // Input: nums = [1,3,3]
    // Output: 4
    // Explanation: The 6 subarrays of nums are the following:
    // [1], range = largest - smallest = 1 - 1 = 0
    // [3], range = 3 - 3 = 0
    // [3], range = 3 - 3 = 0
    // [1,3], range = 3 - 1 = 2
    // [3,3], range = 3 - 3 = 0
    // [1,3,3], range = 3 - 1 = 2
    // So the sum of all ranges is 0 + 0 + 0 + 2 + 0 + 2 = 4.
    fmt.Println(subArrayRanges([]int{1,3,3})) // 4
    // Example 3:
    // Input: nums = [4,-2,-3,4,1]
    // Output: 59
    // Explanation: The sum of all subarray ranges of nums is 59.
    fmt.Println(subArrayRanges([]int{4,-2,-3,4,1})) // 59

    fmt.Println(subArrayRanges1([]int{1,2,3})) // 4
    fmt.Println(subArrayRanges1([]int{1,3,3})) // 4
    fmt.Println(subArrayRanges1([]int{4,-2,-3,4,1})) // 59
}