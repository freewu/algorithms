package main

// 1154. Day of the Year
// Given a string date representing a Gregorian calendar date formatted as YYYY-MM-DD, return the day number of the year.

// Example 1:
// Input: date = "2019-01-09"
// Output: 9
// Explanation: Given date is the 9th day of the year in 2019.

// Example 2:
// Input: date = "2019-02-10"
// Output: 41

// Constraints:
//     date.length == 10
//     date[4] == date[7] == '-', and all other date[i]'s are digits
//     date represents a calendar date between Jan 1st, 1900 and Dec 31th, 2019.

import "fmt"
import "strings"
import "strconv"

func dayOfYear(date string) int {
    months := []int{31,28,31,30,31,30,31,31,30,31,30,31}
    isLeapYear := func(year int) bool { if (year % 4 == 0 && year % 100 != 0) || (year % 400 == 0) { return true; }; return false; }
    arr := strings.Split(date, "-") // 分解
    y, _ := strconv.Atoi(arr[0])
    m, _ := strconv.Atoi(arr[1])
    d, _ := strconv.Atoi(arr[2])
    res := d
    for i := 0; i < m - 1; i++ { // 累加月份天数
        res += months[i]
    }
    if isLeapYear(y) && m > 2 {  // 加闰年天数
        res++
    }
    return res
}

func main() {
    // Example 1:
    // Input: date = "2019-01-09"
    // Output: 9
    // Explanation: Given date is the 9th day of the year in 2019.
    fmt.Println(dayOfYear("2019-01-09")) // 9
    // Example 2:
    // Input: date = "2019-02-10"
    // Output: 41
    fmt.Println(dayOfYear("2019-02-10")) // 41

    fmt.Println(dayOfYear("2020-02-10")) // 42
    fmt.Println(dayOfYear("2012-01-02")) // 2
    fmt.Println(dayOfYear("2016-02-09")) // 40
}