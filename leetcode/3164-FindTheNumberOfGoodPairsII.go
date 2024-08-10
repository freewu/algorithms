package main

// 3164. Find the Number of Good Pairs II
// You are given 2 integer arrays nums1 and nums2 of lengths n and m respectively. 
// You are also given a positive integer k.

// A pair (i, j) is called good if nums1[i] is divisible by nums2[j] * k (0 <= i <= n - 1, 0 <= j <= m - 1).

// Return the total number of good pairs.

// Example 1:
// Input: nums1 = [1,3,4], nums2 = [1,3,4], k = 1
// Output: 5
// Explanation:
// The 5 good pairs are (0, 0), (1, 0), (1, 1), (2, 0), and (2, 2).

// Example 2:
// Input: nums1 = [1,2,4,12], nums2 = [2,4], k = 3
// Output: 2
// Explanation:
// The 2 good pairs are (3, 0) and (3, 1).

// Constraints:
//     1 <= n, m <= 10^5
//     1 <= nums1[i], nums2[j] <= 10^6
//     1 <= k <= 10^3

import "fmt"

func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {
    res, mp := 0, map[int]int{}
    for _, v := range nums1 {
        if v % k > 0 {
            continue;
        }
        v /= k
        factor := 1
        for factor * factor < v {
            if v % factor == 0 {
                mp[factor]++
                mp[v / factor]++
            }
            factor++
        }
        if factor * factor == v {
            mp[factor]++
        }
    }
    for _, v := range nums2 {
        res += mp[v]
    }
    return int64(res)
}

func numberOfPairs1(nums1 []int, nums2 []int, k int) int64 {
    res, mx, mp1, mp2 := 0, 0, map[int]int{}, map[int]int{}
    for _, v := range nums1 {
        if v % k == 0 {
            mp1[v / k]++
        }
    }
    if len(mp1) == 0 {
        return 0
    }
    for _, v := range nums2 {
        mp2[v]++
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for v := range mp1 {
        mx = max(mx, v)
    }
    for i, v := range mp2 {
        s := 0
        for factor := i; factor <= mx; factor += i {
            s += mp1[factor]
        }
        res += s * v 
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums1 = [1,3,4], nums2 = [1,3,4], k = 1
    // Output: 5
    // Explanation:
    // The 5 good pairs are (0, 0), (1, 0), (1, 1), (2, 0), and (2, 2).
    fmt.Println(numberOfPairs([]int{1,3,4}, []int{1,3,4}, 1)) // 5
    // Example 2:
    // Input: nums1 = [1,2,4,12], nums2 = [2,4], k = 3
    // Output: 2
    // Explanation:
    // The 2 good pairs are (3, 0) and (3, 1).
    fmt.Println(numberOfPairs([]int{1,2,4,12}, []int{2,4}, 3)) // 2

    fmt.Println(numberOfPairs1([]int{1,3,4}, []int{1,3,4}, 1)) // 5
    fmt.Println(numberOfPairs1([]int{1,2,4,12}, []int{2,4}, 3)) // 2
}