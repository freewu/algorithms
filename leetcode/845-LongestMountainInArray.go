package main

// 845. Longest Mountain in Array
// You may recall that an array arr is a mountain array if and only if:
//     arr.length >= 3
//     There exists some index i (0-indexed) with 0 < i < arr.length - 1 such that:
//         arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
//         arr[i] > arr[i + 1] > ... > arr[arr.length - 1]

// Given an integer array arr, return the length of the longest subarray, which is a mountain. 
// Return 0 if there is no mountain subarray.

// Example 1:
// Input: arr = [2,1,4,7,3,2,5]
// Output: 5
// Explanation: The largest mountain is [1,4,7,3,2] which has length 5.

// Example 2:
// Input: arr = [2,2,2]
// Output: 0
// Explanation: There is no mountain.

// Constraints:
//     1 <= arr.length <= 10^4
//     0 <= arr[i] <= 10^4

// Follow up:
//     Can you solve it using only one pass?
//     Can you solve it in O(1) space?

import "fmt"

func longestMountain(arr []int) int {
    n, res := len(arr), 0
    if n < 3 { // arr.length >= 3
        return 0
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < n - 1; i++ {
        if arr[i] > arr[i - 1] && arr[i] > arr[i + 1] {
            left, right := i, i
            for left > 0 && arr[left - 1] < arr[left] {
                left--
            }
            for right < n - 1 && arr[right + 1] < arr[right] {
                right++
            }
            res = max(res, right - left + 1)
        }
    }
    return res
}

func longestMountain1(arr []int) int {
    n, res := len(arr), 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < n - 1; i++ {
        if arr[i-1] >= arr[i] || arr[i] <= arr[i+1] {
            continue
        }
        // process 统计最长山脉子数组   right - left + 1
        left, right := i, i
        for left > 0 && arr[left - 1] < arr[left] {
            left--
        }
        for right < n-1 && arr[right] > arr[right + 1] {
            right++
        }
        res = max(res, right - left + 1)
        i = right
    }
    return res
}

func main() {
    // Example 1:
    // Input: arr = [2,1,4,7,3,2,5]
    // Output: 5
    // Explanation: The largest mountain is [1,4,7,3,2] which has length 5.
    fmt.Println(longestMountain([]int{2,1,4,7,3,2,5})) // 5
    // Example 2:
    // Input: arr = [2,2,2]
    // Output: 0
    // Explanation: There is no mountain.
    fmt.Println(longestMountain([]int{2,2,2})) // 0

    fmt.Println(longestMountain1([]int{2,1,4,7,3,2,5})) // 5
    fmt.Println(longestMountain1([]int{2,2,2})) // 0
}
