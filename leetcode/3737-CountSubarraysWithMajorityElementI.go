package main

// 3737. Count Subarrays With Majority Element I
// You are given an integer array nums and an integer target.

// Return the number of subarrays of nums in which target is the majority element.

// The majority element of a subarray is the element that appears strictly more than half of the times in that subarray.

// A subarray is a contiguous non-empty sequence of elements within an array.
 

// Example 1:
// Input: nums = [1,2,2,3], target = 2
// Output: 5
// Explanation:
// Valid subarrays with target = 2 as the majority element:
// nums[1..1] = [2]
// nums[2..2] = [2]
// nums[1..2] = [2,2]
// nums[0..2] = [1,2,2]
// nums[1..3] = [2,2,3]
// So there are 5 such subarrays.

// Example 2:
// Input: nums = [1,1,1,1], target = 1
// Output: 10
// Explanation:
// ​​​​​​​All 10 subarrays have 1 as the majority element.

// Example 3:
// Input: nums = [1,2,3], target = 4
// Output: 0
// Explanation:
// target = 4 does not appear in nums at all. Therefore, there cannot be any subarray where 4 is the majority element. Hence the answer is 0.

// Constraints:
//     1 <= nums.length <= 1000
//     1 <= nums[i] <= 10​​​​​​^​9
//     1 <= target <= 10^9

import "fmt"

// brute force
func countMajoritySubarrays(nums []int, target int) int {
    res, n := 0, len(nums)
    for i := 0; i < n; i++ {
        match := 0
        for j := i; j < n; j++ {
            if nums[j] == target {
                match++
            }
            if match > (j - i + 1) / 2 {
                res++
            }
        }
    }
    return res
}

func countMajoritySubarrays1(nums []int, target int) int {
    res, count, n := 0, 0, len(nums)
    for i, v := range nums {
        if v == target {
            nums[i] = 1
        } else {
            nums[i] = -1
        }
    }
    prefix := make([]int, n + 1)     // 前缀和
    prefix[0] = 0
    for i, v := range nums {
        count += v
        prefix[i+1] = count
    }
    for i := 1; i < n + 1; i++ {
        for j := 0; j < i; j++ {
            if prefix[i] - prefix[j] > 0 {
                res++
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,2,3], target = 2
    // Output: 5
    // Explanation:
    // Valid subarrays with target = 2 as the majority element:
    // nums[1..1] = [2]
    // nums[2..2] = [2]
    // nums[1..2] = [2,2]
    // nums[0..2] = [1,2,2]
    // nums[1..3] = [2,2,3]
    // So there are 5 such subarrays.
    fmt.Println(countMajoritySubarrays([]int{1,2,2,3}, 2)) // 5 
    // Example 2:
    // Input: nums = [1,1,1,1], target = 1
    // Output: 10
    // Explanation:
    // ​​​​​​​All 10 subarrays have 1 as the majority element.
    fmt.Println(countMajoritySubarrays([]int{1,1,1,1}, 1)) // 10
    // Example 3:
    // Input: nums = [1,2,3], target = 4
    // Output: 0
    // Explanation:
    // target = 4 does not appear in nums at all. Therefore, there cannot be any subarray where 4 is the majority element. Hence the answer is 0.
    fmt.Println(countMajoritySubarrays([]int{1,2,3}, 4)) // 0

    fmt.Println(countMajoritySubarrays([]int{1,2,3,4,5,6,7,8,9}, 2)) // 1
    fmt.Println(countMajoritySubarrays([]int{9,8,7,6,5,4,3,2,1}, 2)) // 1

    fmt.Println(countMajoritySubarrays1([]int{1,2,2,3}, 2)) // 5 
    fmt.Println(countMajoritySubarrays1([]int{1,1,1,1}, 1)) // 10
    fmt.Println(countMajoritySubarrays1([]int{1,2,3}, 4)) // 0
    fmt.Println(countMajoritySubarrays1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 1
    fmt.Println(countMajoritySubarrays1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 1
}