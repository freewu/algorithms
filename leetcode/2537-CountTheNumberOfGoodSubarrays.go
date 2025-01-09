package main

// 2537. Count the Number of Good Subarrays
// Given an integer array nums and an integer k, return the number of good subarrays of nums.

// A subarray arr is good if there are at least k pairs of indices (i, j) such that i < j and arr[i] == arr[j].

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [1,1,1,1,1], k = 10
// Output: 1
// Explanation: The only good subarray is the array nums itself.

// Example 2:
// Input: nums = [3,1,4,3,2,2,4], k = 2
// Output: 4
// Explanation: There are 4 different good subarrays:
// - [3,1,4,3,2,2] that has 2 pairs.
// - [3,1,4,3,2,2,4] that has 3 pairs.
// - [1,4,3,2,2,4] that has 2 pairs.
// - [4,3,2,2,4] that has 2 pairs.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i], k <= 10^9

import "fmt"

func countGood(nums []int, k int) int64 {
    count := make(map[int]int)
    res, start, end := 0, 0, 0
    for end < len(nums) {
        k -= count[nums[end]]
        count[nums[end]]++
        for k <= 0 {
            count[nums[start]]--
            k += count[nums[start]]
            start++
        }
        res += start
        end++
    }
    return int64(res)
}

func countGood1(nums []int, k int) int64 {
    count := make(map[int]int)
    res, pairs, left := 0, 0, 0
    for _, v := range nums {
        pairs += count[v]
        count[v]++
        for pairs >= k {
            count[nums[left]]--
            pairs -= count[nums[left]]
            left++
        }
        res += left
    }
    return int64(res)
}

func countGood2(nums []int, k int) int64 {
    count := make(map[int]int)
    res, ctr, left, n := 0, 0, 0, len(nums)
    for i := 0; i < n; i++ {
        ctr += count[nums[i]]
        count[nums[i]]++
        if ctr >= k {
            for left <= i {
                count[nums[left]]--
                ctr -= count[nums[left]]
                left++
                if ctr < k { break }
            }   
        }
        res += left
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [1,1,1,1,1], k = 10
    // Output: 1
    // Explanation: The only good subarray is the array nums itself.
    fmt.Println(countGood([]int{1,1,1,1,1}, 10)) // 1
    // Example 2:
    // Input: nums = [3,1,4,3,2,2,4], k = 2
    // Output: 4
    // Explanation: There are 4 different good subarrays:
    // - [3,1,4,3,2,2] that has 2 pairs.
    // - [3,1,4,3,2,2,4] that has 3 pairs.
    // - [1,4,3,2,2,4] that has 2 pairs.
    // - [4,3,2,2,4] that has 2 pairs.
    fmt.Println(countGood([]int{3,1,4,3,2,2,4}, 2)) // 4 

    fmt.Println(countGood([]int{1,2,3,4,5,6,7,8,9}, 2)) // 0

    fmt.Println(countGood1([]int{1,1,1,1,1}, 10)) // 1
    fmt.Println(countGood1([]int{3,1,4,3,2,2,4}, 2)) // 4 
    fmt.Println(countGood1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 0

    fmt.Println(countGood2([]int{1,1,1,1,1}, 10)) // 1
    fmt.Println(countGood2([]int{3,1,4,3,2,2,4}, 2)) // 4 
    fmt.Println(countGood2([]int{1,2,3,4,5,6,7,8,9}, 2)) // 0
}