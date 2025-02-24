package main

// 3209. Number of Subarrays With AND Value of K
// Given an array of integers nums and an integer k, 
// return the number of subarrays of nums where the bitwise AND of the elements of the subarray equals k.

// Example 1:
// Input: nums = [1,1,1], k = 1
// Output: 6
// Explanation:
// All subarrays contain only 1's.

// Example 2:
// Input: nums = [1,1,2], k = 1
// Output: 3
// Explanation:
// Subarrays having an AND value of 1 are: [1,1,2], [1,1,2], [1,1,2].

// Example 3:
// Input: nums = [1,2,3], k = 2
// Output: 2
// Explanation:
// Subarrays having an AND value of 2 are: [1,2,3], [1,2,3].

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i], k <= 10^9

import "fmt"
import "sort"

func countSubarrays(nums []int, k int) int64 {
    res, count := 0, make(map[int]int)
    for _, n := range nums {
        mp := make(map[int]int)
        if n & k == k {
            mp[n] = 1
            for v, c := range count {
                val := v & n
                if val & k == k {
                    mp[val] += c
                }
            }
            res += mp[k]
        }
        count = mp
    }
    return int64(res)
}

func countSubarrays1(nums []int, k int) int64 {
    res := 0
    for i, v := range nums {
        for j := i - 1; j >= 0 && nums[j] & v != nums[j]; j-- {
            nums[j] &= v
        }
        arr := nums[:i + 1]
        res += (sort.SearchInts(arr, k + 1) - sort.SearchInts(arr, k))
    }
    return int64(res)
}

func countSubarrays2(nums []int, k int) int64 {
    res, count := 0, 0
    for i, v := range nums {
        if v == k {
            count++
        }
        for j := i - 1; j >= 0 && nums[j] & v != nums[j]; j-- {
            if nums[j] == k {
                count--
            }
            nums[j] &= v
            if nums[j] == k {
                count++
            }
        }
        res += count
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [1,1,1], k = 1
    // Output: 6
    // Explanation:
    // All subarrays contain only 1's.
    fmt.Println(countSubarrays([]int{1,1,1}, 1)) // 6
    // Example 2:
    // Input: nums = [1,1,2], k = 1
    // Output: 3
    // Explanation:
    // Subarrays having an AND value of 1 are: [1,1,2], [1,1,2], [1,1,2].
    fmt.Println(countSubarrays([]int{1,1,2}, 1)) // 3
    // Example 3:
    // Input: nums = [1,2,3], k = 2
    // Output: 2
    // Explanation:
    // Subarrays having an AND value of 2 are: [1,2,3], [1,2,3].
    fmt.Println(countSubarrays([]int{1,2,3}, 2)) // 2

    fmt.Println(countSubarrays([]int{1,2,3,4,5,6,7,8,9}, 2)) // 2
    fmt.Println(countSubarrays([]int{9,8,7,6,5,4,3,2,1}, 2)) // 2

    fmt.Println(countSubarrays1([]int{1,1,1}, 1)) // 6
    fmt.Println(countSubarrays1([]int{1,1,2}, 1)) // 3
    fmt.Println(countSubarrays1([]int{1,2,3}, 2)) // 2
    fmt.Println(countSubarrays1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 2
    fmt.Println(countSubarrays1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 2

    fmt.Println(countSubarrays2([]int{1,1,1}, 1)) // 6
    fmt.Println(countSubarrays2([]int{1,1,2}, 1)) // 3
    fmt.Println(countSubarrays2([]int{1,2,3}, 2)) // 2
    fmt.Println(countSubarrays2([]int{1,2,3,4,5,6,7,8,9}, 2)) // 2
    fmt.Println(countSubarrays2([]int{9,8,7,6,5,4,3,2,1}, 2)) // 2
}