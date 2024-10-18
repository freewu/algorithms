package main

// 3185. Count Pairs That Form a Complete Day II
// Given an integer array hours representing times in hours, 
// return an integer denoting the number of pairs i, j where i < j and hours[i] + hours[j] forms a complete day.

// A complete day is defined as a time duration that is an exact multiple of 24 hours.

// For example, 1 day is 24 hours, 2 days is 48 hours, 3 days is 72 hours, and so on.

// Example 1:
// Input: hours = [12,12,30,24,24]
// Output: 2
// Explanation: The pairs of indices that form a complete day are (0, 1) and (3, 4).

// Example 2:
// Input: hours = [72,48,24,3]
// Output: 3
// Explanation: The pairs of indices that form a complete day are (0, 1), (0, 2), and (1, 2).

// Constraints:
//     1 <= hours.length <= 5 * 10^5
//     1 <= hours[i] <= 10^9

import "fmt"

func countCompleteDayPairs(hours []int) int64 {
    res, count := 0, make([]int, 24)
    for _, h := range hours {
        remainder := h % 24
        res += count[(24 - remainder) % 24]
        count[remainder]++
    }
    return int64(res)
}

func countCompleteDayPairs1(hours []int) int64 {
    mp := make([]int64, 24)
    for _, h := range hours {
        mp[h%24] += 1
    }
    res := mp[0] * (mp[0] - 1) / 2
    res += mp[12] * (mp[12] - 1) / 2
    for i := 1; i < 12; i++ {
        res += mp[i] * mp[24 - i]
    }
    return res
}

func main() {
    // Example 1:
    // Input: hours = [12,12,30,24,24]
    // Output: 2
    // Explanation: The pairs of indices that form a complete day are (0, 1) and (3, 4).
    fmt.Println(countCompleteDayPairs([]int{12,12,30,24,24})) // 2
    // Example 2:
    // Input: hours = [72,48,24,3]
    // Output: 3
    // Explanation: The pairs of indices that form a complete day are (0, 1), (0, 2), and (1, 2).
    fmt.Println(countCompleteDayPairs([]int{72,48,24,3})) // 3

    fmt.Println(countCompleteDayPairs1([]int{12,12,30,24,24})) // 2
    fmt.Println(countCompleteDayPairs1([]int{72,48,24,3})) // 3
}