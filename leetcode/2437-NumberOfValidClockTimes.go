package main

// 2437. Number of Valid Clock Times
// You are given a string of length 5 called time, representing the current time on a digital clock in the format "hh:mm". 
// The earliest possible time is "00:00" and the latest possible time is "23:59".

// In the string time, the digits represented by the ? symbol are unknown, and must be replaced with a digit from 0 to 9.

// Return an integer answer, the number of valid clock times that can be created by replacing every ? with a digit from 0 to 9.

// Example 1:
// Input: time = "?5:00"
// Output: 2
// Explanation: We can replace the ? with either a 0 or 1, producing "05:00" or "15:00". 
// Note that we cannot replace it with a 2, since the time "25:00" is invalid. In total, we have two choices.

// Example 2:
// Input: time = "0?:0?"
// Output: 100
// Explanation: Each ? can be replaced by any digit from 0 to 9, so we have 100 total choices.

// Example 3:
// Input: time = "??:??"
// Output: 1440
// Explanation: There are 24 possible choices for the hours, and 60 possible choices for the minutes. 
// In total, we have 24 * 60 = 1440 choices.

// Constraints:
//     time is a valid string of length 5 in the format "hh:mm".
//     "00" <= hh <= "23"
//     "00" <= mm <= "59"
//     Some of the digits might be replaced with '?' and need to be replaced with digits from 0 to 9.

import "fmt"

func countTime(time string) int {
    res := 0
    for hour := 0; hour < 24; hour++ {
        for minute := 0; minute < 60; minute++ {
            hourHigh, hourLow, minuteHigh, minuteLow := hour / 10, hour % 10, minute / 10, minute % 10
            if time[0] != '?' && time[0] != byte(hourHigh +'0')   { continue }
            if time[1] != '?' && time[1] != byte(hourLow +'0')    { continue }
            if time[3] != '?' && time[3] != byte(minuteHigh +'0') { continue }
            if time[4] != '?' && time[4] != byte(minuteLow +'0')  { continue }
            res++
        }
    }
    return res
}

func countTime1(time string) int {
    res, hour, minute := 1, time[0:2],time[3:]
    if hour[0] == '?' {
        if hour[1] == '?' {
            res *= 24
        } else if hour[1] <= '3' {
            res *= 3
        } else {
            res *= 2
        }
    } else if hour[1] == '?' {
        if hour[0] < '2' {
            res *= 10
        } else {
            res *= 4
        }
    }
    if minute[0] == '?' {
        if minute[1] == '?' {
            res *= 60
        } else {
            res *= 6
        }
    } else if minute[1] == '?' {
        res *= 10
    }
    return res
}

func main() {
    // Example 1:
    // Input: time = "?5:00"
    // Output: 2
    // Explanation: We can replace the ? with either a 0 or 1, producing "05:00" or "15:00". 
    // Note that we cannot replace it with a 2, since the time "25:00" is invalid. In total, we have two choices.
    fmt.Println(countTime("?5:00")) // 2
    // Example 2:
    // Input: time = "0?:0?"
    // Output: 100
    // Explanation: Each ? can be replaced by any digit from 0 to 9, so we have 100 total choices.
    fmt.Println(countTime("0?:0?")) // 100
    // Example 3:
    // Input: time = "??:??"
    // Output: 1440
    // Explanation: There are 24 possible choices for the hours, and 60 possible choices for the minutes. 
    // In total, we have 24 * 60 = 1440 choices.
    fmt.Println(countTime("??:??")) // 1440

    fmt.Println(countTime1("?5:00")) // 2
    fmt.Println(countTime1("0?:0?")) // 100
    fmt.Println(countTime1("??:??")) // 1440
}

