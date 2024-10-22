package main

// 1588. Sum of All Odd Length Subarrays
// Given an array of positive integers arr, return the sum of all possible odd-length subarrays of arr.

// A subarray is a contiguous subsequence of the array.

// Example 1:
// Input: arr = [1,4,2,5,3]
// Output: 58
// Explanation: The odd-length subarrays of arr and their sums are:
// [1] = 1
// [4] = 4
// [2] = 2
// [5] = 5
// [3] = 3
// [1,4,2] = 7
// [4,2,5] = 11
// [2,5,3] = 10
// [1,4,2,5,3] = 15
// If we add all these together we get 1 + 4 + 2 + 5 + 3 + 7 + 11 + 10 + 15 = 58

// Example 2:
// Input: arr = [1,2]
// Output: 3
// Explanation: There are only 2 subarrays of odd length, [1] and [2]. Their sum is 3.

// Example 3:
// Input: arr = [10,11,12]
// Output: 66

// Constraints:
//     1 <= arr.length <= 100
//     1 <= arr[i] <= 1000

// Follow up:
//     Could you solve this problem in O(n) time complexity?

import "fmt"

// sliding window
func sumOddLengthSubarrays(arr []int) int {
    res := 0
    for i := 1; i <= len(arr); i += 2 {
        l, r := 0, i - 1
        for r < len(arr) {
            for j := l; j <= r; j++ {
                res += arr[j]
            }
            l++
            r++
        }
    }
    return res
}

func sumOddLengthSubarrays1(arr []int) int {
    res, n := 0, len(arr)
    odd, even := make([]int, n), make([]int, n) // 以 i 结尾的的奇数序列和, 以 i 结尾的偶数序列和
    for i, v := range arr {
        if i == 0 {
            even[i], odd[i] = 0, v
        } else {
            // 在 i-1 的偶数序列上，加上当前数，构成奇数序列
            // 长度为 i 的序列，有 i/2 个偶数序列，+1 表示当前数单独构成奇数序列
            odd[i] = (i / 2 + 1) * v + even[i-1]
            even[i] = (i + 1) / 2 * v + odd[i-1]
        }
        res += odd[i]
    }
    return res
}

func main() {
    // Example 1:
    // Input: arr = [1,4,2,5,3]
    // Output: 58
    // Explanation: The odd-length subarrays of arr and their sums are:
    // [1] = 1
    // [4] = 4
    // [2] = 2
    // [5] = 5
    // [3] = 3
    // [1,4,2] = 7
    // [4,2,5] = 11
    // [2,5,3] = 10
    // [1,4,2,5,3] = 15
    // If we add all these together we get 1 + 4 + 2 + 5 + 3 + 7 + 11 + 10 + 15 = 58
    fmt.Println(sumOddLengthSubarrays([]int{1,4,2,5,3})) // 58
    // Example 2:
    // Input: arr = [1,2]
    // Output: 3
    // Explanation: There are only 2 subarrays of odd length, [1] and [2]. Their sum is 3.
    fmt.Println(sumOddLengthSubarrays([]int{1,2})) // 3
    // Example 3:
    // Input: arr = [10,11,12]
    // Output: 66
    fmt.Println(sumOddLengthSubarrays([]int{10,11,12})) // 66

    fmt.Println(sumOddLengthSubarrays1([]int{1,4,2,5,3})) // 58
    fmt.Println(sumOddLengthSubarrays1([]int{1,2})) // 3
    fmt.Println(sumOddLengthSubarrays1([]int{10,11,12})) // 66
}