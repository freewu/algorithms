package main

// 2111. Minimum Operations to Make the Array K-Increasing
// You are given a 0-indexed array arr consisting of n positive integers, and a positive integer k.

// The array arr is called K-increasing if arr[i-k] <= arr[i] holds for every index i, where k <= i <= n-1.
//     1. For example, arr = [4, 1, 5, 2, 6, 2] is K-increasing for k = 2 because:
//         arr[0] <= arr[2] (4 <= 5)
//         arr[1] <= arr[3] (1 <= 2)
//         arr[2] <= arr[4] (5 <= 6)
//         arr[3] <= arr[5] (2 <= 2)
//     2. However, the same arr is not K-increasing for k = 1 (because arr[0] > arr[1]) or k = 3 (because arr[0] > arr[3]).

// In one operation, you can choose an index i and change arr[i] into any positive integer.

// Return the minimum number of operations required to make the array K-increasing for the given k.

// Example 1:
// Input: arr = [5,4,3,2,1], k = 1
// Output: 4
// Explanation:
// For k = 1, the resultant array has to be non-decreasing.
// Some of the K-increasing arrays that can be formed are [5,6,7,8,9], [1,1,1,1,1], [2,2,3,4,4]. All of them require 4 operations.
// It is suboptimal to change the array to, for example, [6,7,8,9,10] because it would take 5 operations.
// It can be shown that we cannot make the array K-increasing in less than 4 operations.

// Example 2:
// Input: arr = [4,1,5,2,6,2], k = 2
// Output: 0
// Explanation:
// This is the same example as the one in the problem description.
// Here, for every index i where 2 <= i <= 5, arr[i-2] <= arr[i].
// Since the given array is already K-increasing, we do not need to perform any operations.

// Example 3:
// Input: arr = [4,1,5,2,6,2], k = 3
// Output: 2
// Explanation:
// Indices 3 and 5 are the only ones not satisfying arr[i-3] <= arr[i] for 3 <= i <= 5.
// One of the ways we can make the array K-increasing is by changing arr[3] to 4 and arr[5] to 5.
// The array will now be [4,1,5,4,6,5].
// Note that there can be other ways to make the array K-increasing, but none of them require less than 2 operations.

// Constraints:
//     1 <= arr.length <= 10^5
//     1 <= arr[i], k <= arr.length

import "fmt"

func kIncreasing(arr []int, k int) int {
    res := 0
    binarySearch := func(arr []int, left, right, target int) int {
        for left <= right {
            mid := left + (right - left)>>1
            if arr[mid] <= target {
                left = mid + 1
            } else if arr[mid] > target {
                right = mid - 1
            }
        }
        return left
    }
    maxNonDecreaseSubarray := func (arr []int) int {
        n, index := len(arr), 0
        dp := make([]int, n)
        dp[0] = arr[0]
        for i := 1; i < n; i++ {
            if arr[i] >= dp[index] {
                index++
                dp[index] = arr[i]
            } else {
                dp[binarySearch(dp, 0, index, arr[i])] = arr[i]
            }
        }
        return index + 1
    }
    for i := 0; i < k; i++ {
        sub := make([]int, 0)
        for j := i; j < len(arr); j += k {
            sub = append(sub, arr[j])
        }
        res += len(sub) - maxNonDecreaseSubarray(sub)
    }
    return res
}

func kIncreasing1(arr []int, k int) int {
    res, n := 0, len(arr)
    binarySearch := func(arr []int, target int) int {
        left, right := 0, len(arr) - 1
        for left < right {
            mid := (left + right) / 2
            if arr[mid] > target {
                right = mid
            } else {
                left = mid + 1
            }
        }
        return left
    }
    for i := 0; i < k; i++ {
        sub, left := make([]int,0), 1
        sub = append(sub,arr[i])
        for j := i + k; j < n; j += k {
            if arr[j] >= sub[len(sub)-1] {
                sub = append(sub,arr[j])
            } else {
                index := binarySearch(sub,arr[j])
                sub[index] = arr[j]
            }
            left++
        }
        res += (left - len(sub))
    }
    return res
}

// func kIncreasing1(arr []int, k int) int {
//     res, n := 0, len(arr)
//     binarySearch := func(arr []int, target int) int {
//         left, right := 0, len(arr) - 1
//         for left <= right {
//             mid := left + (right - left) >> 1
//             if arr[mid] > target {
//                 right = mid
//             } else {
//                 left = mid + 1
//             }
//         }
//         return left
//     }
//     for i := 0; i < k;i++ {
//         sub, left := make([]int,0), 1
//         sub = append(sub, arr[i])
//         for j := i + k; j < n; j += k {
//             if arr[j] >= sub[len(sub) - 1] {
//                 sub = append(sub, arr[j])
//             } else {
//                 index := binarySearch(sub,arr[j])
//                 sub[index] = arr[j]
//             }
//             left++
//         }
//         res += (left - len(sub))
//     }
//     return res
// }

func main() {
    // Example 1:
    // Input: arr = [5,4,3,2,1], k = 1
    // Output: 4
    // Explanation:
    // For k = 1, the resultant array has to be non-decreasing.
    // Some of the K-increasing arrays that can be formed are [5,6,7,8,9], [1,1,1,1,1], [2,2,3,4,4]. All of them require 4 operations.
    // It is suboptimal to change the array to, for example, [6,7,8,9,10] because it would take 5 operations.
    // It can be shown that we cannot make the array K-increasing in less than 4 operations.
    fmt.Println(kIncreasing([]int{5,4,3,2,1}, 1)) // 4
    // Example 2:
    // Input: arr = [4,1,5,2,6,2], k = 2
    // Output: 0
    // Explanation:
    // This is the same example as the one in the problem description.
    // Here, for every index i where 2 <= i <= 5, arr[i-2] <= arr[i].
    // Since the given array is already K-increasing, we do not need to perform any operations.
    fmt.Println(kIncreasing([]int{4,1,5,2,6,2}, 2)) // 0
    // Example 3:
    // Input: arr = [4,1,5,2,6,2], k = 3
    // Output: 2
    // Explanation:
    // Indices 3 and 5 are the only ones not satisfying arr[i-3] <= arr[i] for 3 <= i <= 5.
    // One of the ways we can make the array K-increasing is by changing arr[3] to 4 and arr[5] to 5.
    // The array will now be [4,1,5,4,6,5].
    // Note that there can be other ways to make the array K-increasing, but none of them require less than 2 operations.
    fmt.Println(kIncreasing([]int{4,1,5,2,6,2}, 3)) // 2

    fmt.Println(kIncreasing1([]int{5,4,3,2,1}, 1)) // 4
    fmt.Println(kIncreasing1([]int{4,1,5,2,6,2}, 2)) // 0
    fmt.Println(kIncreasing1([]int{4,1,5,2,6,2}, 3)) // 2
}