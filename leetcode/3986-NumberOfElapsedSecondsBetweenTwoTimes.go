package main

// 3986. Number of Elapsed Seconds Between Two Times
// You are given two valid times startTime and endTime, each represented as a string in the format "HH:MM:SS".

// Return the number of seconds that have elapsed from startTime to endTime, inclusive of both endpoints.

// Example 1:
// Input: startTime = "01:00:00", endTime = "01:00:25"
// Output: 25
// Explanation:
// endTime is 25 seconds ahead of startTime.

// Example 2:
// Input: startTime = "12:34:56", endTime = "13:00:00"
// Output: 1504
// Explanation:
// endTime is 25 minutes and 4 seconds ahead of startTime, which equals 1504 seconds.

// Constraints:
//     startTime.length == 8
//     endTime.length == 8
//     startTime and endTime are valid times in the format "HH:MM:SS"
//     00 <= HH <= 23
//     00 <= MM <= 59
//     00 <= SS <= 59
//     endTime is not earlier than startTime

import "fmt"

func secondsBetweenTimes(startTime string, endTime string) int {
    parse := func(t string) int {
        hour := int(t[0]-'0')*10 + int(t[1]-'0')
        minute := int(t[3]-'0')*10 + int(t[4]-'0')
        second := int(t[6]-'0')*10 + int(t[7]-'0')
        return hour * 3600 + minute * 60 + second
    }
    return parse(endTime) - parse(startTime)
}

func main() {
    // Example 1:
    // Input: startTime = "01:00:00", endTime = "01:00:25"
    // Output: 25
    // Explanation:
    // endTime is 25 seconds ahead of startTime.
    fmt.Println(secondsBetweenTimes("01:00:00", "01:00:25")) // 25
    // Example 2:
    // Input: startTime = "12:34:56", endTime = "13:00:00"
    // Output: 1504
    // Explanation:
    // endTime is 25 minutes and 4 seconds ahead of startTime, which equals 1504 seconds.
    fmt.Println(secondsBetweenTimes("12:34:56", "13:00:00")) // 1504

    fmt.Println(secondsBetweenTimes("00:00:00", "23:59:59")) // 86399
}