package main

// 2224. Minimum Number of Operations to Convert Time
// You are given two strings current and correct representing two 24-hour times.

// 24-hour times are formatted as "HH:MM", where HH is between 00 and 23, and MM is between 00 and 59. 
// The earliest 24-hour time is 00:00, and the latest is 23:59.

// In one operation you can increase the time current by 1, 5, 15, or 60 minutes. 
// You can perform this operation any number of times.

// Return the minimum number of operations needed to convert current to correct.

// Example 1:
// Input: current = "02:30", correct = "04:35"
// Output: 3
// Explanation:
// We can convert current to correct in 3 operations as follows:
// - Add 60 minutes to current. current becomes "03:30".
// - Add 60 minutes to current. current becomes "04:30".
// - Add 5 minutes to current. current becomes "04:35".
// It can be proven that it is not possible to convert current to correct in fewer than 3 operations.

// Example 2:
// Input: current = "11:00", correct = "11:01"
// Output: 1
// Explanation: We only have to add one minute to current, so the minimum number of operations needed is 1.

// Constraints:
//     current and correct are in the format "HH:MM"
//     current <= correct

import "fmt"
import "strconv"

func convertTime(current, correct string) (op int) {
    convertor := func(s string) int {
        hours, _ := strconv.Atoi(s[:2])
        minutes, _ := strconv.Atoi(s[3:5])
        return hours * 60 + minutes
    }
    res, diff :=  0, convertor(correct) - convertor(current)
    for _, v := range []int{ 60, 15, 5, 1 } {
        res += diff/v
        diff %= v
    }
    return res
}

func main() {
    // Example 1:
    // Input: current = "02:30", correct = "04:35"
    // Output: 3
    // Explanation:
    // We can convert current to correct in 3 operations as follows:
    // - Add 60 minutes to current. current becomes "03:30".
    // - Add 60 minutes to current. current becomes "04:30".
    // - Add 5 minutes to current. current becomes "04:35".
    // It can be proven that it is not possible to convert current to correct in fewer than 3 operations.
    fmt.Println(convertTime("02:30", "04:35")) // 3
    // Example 2:
    // Input: current = "11:00", correct = "11:01"
    // Output: 1
    // Explanation: We only have to add one minute to current, so the minimum number of operations needed is 1.
    fmt.Println(convertTime("11:00", "11:01")) // 3
}