package main

// 2488. Count Subarrays With Median K
// You are given an array nums of size n consisting of distinct integers from 1 to n and a positive integer k.

// Return the number of non-empty subarrays in nums that have a median equal to k.

// Note:
//     1. The median of an array is the middle element after sorting the array in ascending order. 
//        If the array is of even length, the median is the left middle element.
//             For example, the median of [2,3,1,4] is 2, and the median of [8,4,3,5,1] is 4.
//     2. A subarray is a contiguous part of an array.

// Example 1:
// Input: nums = [3,2,1,4,5], k = 4
// Output: 3
// Explanation: The subarrays that have a median equal to 4 are: [4], [4,5] and [1,4,5].

// Example 2:
// Input: nums = [2,3,1], k = 3
// Output: 1
// Explanation: [3] is the only subarray that has a median equal to 3.

// Constraints:
//     n == nums.length
//     1 <= n <= 10^5
//     1 <= nums[i], k <= n
//     The integers in nums are distinct.

import "fmt"

func countSubarrays(nums []int, k int) int {
    res, s, flag := 0, 0, false
    mp := map[int]int{0 : 1}
    for _, v := range nums {
        if v > k {
            s++
        } else if v < k {
            s--
        } else {
            flag = true
        }
        if !flag {
            mp[s]++
        } else {
            res += (mp[s] + mp[s - 1])
        }
    }
    return res
}

func countSubarrays1(nums []int, k int) int {
    pos, count, n := 0, 0, len(nums)
    mp := make(map[int]int)
    for i, v := range nums {
        if v == k {
            pos = i
        }
    }
    mp[0]++
    for i := pos - 1; i >= 0; i-- {
        if nums[i] < k {
            count--
        } else {
            count++
        }
        mp[count]++
    }
    count = 0
    res := mp[0] + mp[1]
    for i := pos + 1; i < n; i++ {
        if nums[i] < k {
            count--
        } else {
            count++
        }
        res += (mp[-count] + mp[-count + 1])
    }
    return res
}

func countSubarrays2(nums []int, k int) int {
    res, x, i, n := 1, 0, 0, len(nums)
    for nums[i] != k {
        i++
    }
    count := make([]int, n << 1 | 1)
    for j := i + 1; j < n; j++ {
        if nums[j] > k {
            x++
        } else {
            x--
        }
        if x >= 0 && x <= 1 {
            res++
        }
        count[x + n]++
    }
    x = 0
    for j := i - 1; j >= 0; j-- {
        if nums[j] > k {
            x++
        } else {
            x--
        }
        if x >= 0 && x <= 1 {
            res++
        }
        res += (count[-x + n] + count[-x + n + 1])
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,2,1,4,5], k = 4
    // Output: 3
    // Explanation: The subarrays that have a median equal to 4 are: [4], [4,5] and [1,4,5].
    fmt.Println(countSubarrays([]int{3,2,1,4,5}, 4)) // 3
    // Example 2:
    // Input: nums = [2,3,1], k = 3
    // Output: 1
    // Explanation: [3] is the only subarray that has a median equal to 3.
    fmt.Println(countSubarrays([]int{2,3,1}, 3)) // 1

    fmt.Println(countSubarrays([]int{1,2,3,4,5,6,7,8,9}, 3)) // 6
    fmt.Println(countSubarrays([]int{9,8,7,6,5,4,3,2,1}, 3)) // 6

    fmt.Println(countSubarrays1([]int{3,2,1,4,5}, 4)) // 3
    fmt.Println(countSubarrays1([]int{2,3,1}, 3)) // 1
    fmt.Println(countSubarrays1([]int{1,2,3,4,5,6,7,8,9}, 3)) // 6
    fmt.Println(countSubarrays1([]int{9,8,7,6,5,4,3,2,1}, 3)) // 6

    fmt.Println(countSubarrays2([]int{3,2,1,4,5}, 4)) // 3
    fmt.Println(countSubarrays2([]int{2,3,1}, 3)) // 1
    fmt.Println(countSubarrays2([]int{1,2,3,4,5,6,7,8,9}, 3)) // 6
    fmt.Println(countSubarrays2([]int{9,8,7,6,5,4,3,2,1}, 3)) // 6
}