package main

// 2799. Count Complete Subarrays in an Array
// You are given an array nums consisting of positive integers.

// We call a subarray of an array complete if the following condition is satisfied:
//     The number of distinct elements in the subarray is equal to the number of distinct elements in the whole array.

// Return the number of complete subarrays.

// A subarray is a contiguous non-empty part of an array.

// Example 1:
// Input: nums = [1,3,1,2,2]
// Output: 4
// Explanation: The complete subarrays are the following: [1,3,1,2], [1,3,1,2,2], [3,1,2] and [3,1,2,2].

// Example 2:
// Input: nums = [5,5,5,5]
// Output: 10
// Explanation: The array consists only of the integer 5, so any subarray is complete. The number of subarrays that we can choose is 10.

// Constraints:
//     1 <= nums.length <= 1000
//     1 <= nums[i] <= 2000

import "fmt"

func countCompleteSubarrays(nums []int) int {
    res, l, freq, set := 0, 0, make(map[int]int), make(map[int]bool)
    for _, v := range nums {
        set[v] = true
    }
    for r, v := range nums {
        freq[v]++
        for len(set) == len(freq) {
            freq[nums[l]]--
            if freq[nums[l]] == 0 {
                delete(freq, nums[l])
            }
            res += (len(nums) - r)
            l++
        }
    }
    return res
}

func countCompleteSubarrays1(nums []int) int {
    n, unique := len(nums), 0
    count, count2 := make([]int,2001), make([]int,2001)
    for _, v := range nums {
        count[v]++
        if count[v] == 1 {
            unique++
        }
    }
    res, i, j, unique2 := 0, 0, 0, 0
    for j < n {
        count2[nums[j]] ++
        if count2[nums[j]] == 1 {
            unique2++
        }
        for i <= j && unique2 == unique {
            res += (n - j)
            count2[nums[i]] --
            if count2[nums[i]] == 0 {
                unique2--
            }
            i++
        }
        j++
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,3,1,2,2]
    // Output: 4
    // Explanation: The complete subarrays are the following: [1,3,1,2], [1,3,1,2,2], [3,1,2] and [3,1,2,2].
    fmt.Println(countCompleteSubarrays([]int{1,3,1,2,2})) // 4
    // Example 2:
    // Input: nums = [5,5,5,5]
    // Output: 10
    // Explanation: The array consists only of the integer 5, so any subarray is complete. The number of subarrays that we can choose is 10.
    fmt.Println(countCompleteSubarrays([]int{5,5,5,5})) // 10

    fmt.Println(countCompleteSubarrays([]int{1,2,3,4,5,6,7,8,9})) // 1
    fmt.Println(countCompleteSubarrays([]int{9,8,7,6,5,4,3,2,1})) // 1

    fmt.Println(countCompleteSubarrays1([]int{1,3,1,2,2})) // 4
    fmt.Println(countCompleteSubarrays1([]int{5,5,5,5})) // 10
    fmt.Println(countCompleteSubarrays1([]int{1,2,3,4,5,6,7,8,9})) // 1
    fmt.Println(countCompleteSubarrays1([]int{9,8,7,6,5,4,3,2,1})) // 1
}