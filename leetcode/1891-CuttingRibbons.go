package main

// 1891. Cutting Ribbons
// You are given an integer array ribbons, where ribbons[i] represents the length of the ith ribbon, and an integer k. 
// You may cut any of the ribbons into any number of segments of positive integer lengths, or perform no cuts at all.

//     For example, if you have a ribbon of length 4, you can:
//         Keep the ribbon of length 4,
//         Cut it into one ribbon of length 3 and one ribbon of length 1,
//         Cut it into two ribbons of length 2,
//         Cut it into one ribbon of length 2 and two ribbons of length 1, or
//         Cut it into four ribbons of length 1.

// Your goal is to obtain k ribbons of all the same positive integer length. 
// You are allowed to throw away any excess ribbon as a result of cutting.

// Return the maximum possible positive integer length that you can obtain k ribbons of, 
// or 0 if you cannot obtain k ribbons of the same length.

// Example 1:
// Input: ribbons = [9,7,5], k = 3
// Output: 5
// Explanation:
// - Cut the first ribbon to two ribbons, one of length 5 and one of length 4.
// - Cut the second ribbon to two ribbons, one of length 5 and one of length 2.
// - Keep the third ribbon as it is.
// Now you have 3 ribbons of length 5.

// Example 2:
// Input: ribbons = [7,5,9], k = 4
// Output: 4
// Explanation:
// - Cut the first ribbon to two ribbons, one of length 4 and one of length 3.
// - Cut the second ribbon to two ribbons, one of length 4 and one of length 1.
// - Cut the third ribbon to three ribbons, two of length 4 and one of length 1.
// Now you have 4 ribbons of length 4.

// Example 3:
// Input: ribbons = [5,7,9], k = 22
// Output: 0
// Explanation: You cannot obtain k ribbons of the same positive integer length.

// Constraints:
//     1 <= ribbons.length <= 10^5
//     1 <= ribbons[i] <= 10^5
//     1 <= k <= 10^9

import "fmt"

func maxLength(ribbons []int, k int) int {
    low, high := 1, 100000
    canMake := func(ribbons []int, i, k int) bool {
        n := 0
        for _, ribbon := range ribbons {
            n += ribbon / i
        }
        if n >= k { return true }
        return false
    }
    for low <= high {
        mid := low + (high - low) / 2
        if canMake(ribbons, mid, k) { // 遍历过程中判断当前长度是否可以形成k条及以上的绳子
            low = mid + 1 // 如果可以形成，则往更多的方向划分
        } else {
            high = mid - 1 // 如果不可以，则往更少的方向划分
        }
    }
    return low - 1
}

func maxLength1(ribbons []int, k int) int {
    sum := 0
    for _, v := range ribbons {
        sum += v
    }
    if sum < k  { return 0 } // 长度不足
    if sum == k { return 1 } // 刚好满足
    r, l, m := sum / k, 1, 0
    for l <= r {
        m = (l + r) / 2
        t := 0
        for _, v := range ribbons {
            t += v / m
        }
        if t >= k {
            l = m + 1
        } else {
            r = m - 1
        }
    }
    return r
}

func main() {
    // Example 1:
    // Input: ribbons = [9,7,5], k = 3
    // Output: 5
    // Explanation:
    // - Cut the first ribbon to two ribbons, one of length 5 and one of length 4.
    // - Cut the second ribbon to two ribbons, one of length 5 and one of length 2.
    // - Keep the third ribbon as it is.
    // Now you have 3 ribbons of length 5.
    fmt.Println(maxLength([]int{9,7,5}, 3)) // 5
    // Example 2:
    // Input: ribbons = [7,5,9], k = 4
    // Output: 4
    // Explanation:
    // - Cut the first ribbon to two ribbons, one of length 4 and one of length 3.
    // - Cut the second ribbon to two ribbons, one of length 4 and one of length 1.
    // - Cut the third ribbon to three ribbons, two of length 4 and one of length 1.
    // Now you have 4 ribbons of length 4.
    fmt.Println(maxLength([]int{7,5,9}, 4)) // 4
    // Example 3:
    // Input: ribbons = [5,7,9], k = 22
    // Output: 0
    // Explanation: You cannot obtain k ribbons of the same positive integer length.
    fmt.Println(maxLength([]int{5,7,9}, 22)) // 0

    fmt.Println(maxLength1([]int{9,7,5}, 3)) // 5
    fmt.Println(maxLength1([]int{7,5,9}, 4)) // 4
    fmt.Println(maxLength1([]int{5,7,9}, 22)) // 0
}