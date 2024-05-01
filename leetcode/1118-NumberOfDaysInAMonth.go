package main

// 1118. Number of Days in a Month
// Given a year year and a month month, return the number of days of that month.

// Example 1:
// Input: year = 1992, month = 7
// Output: 31

// Example 2:
// Input: year = 2000, month = 2
// Output: 29

// Example 3:
// Input: year = 1900, month = 2
// Output: 28
 
// Constraints:
//     1583 <= year <= 2100
//     1 <= month <= 12

import "fmt"

func numberOfDays(year int, month int) int {
    isLeapYear := func(year int) bool { // 判断是否是闰年
        return year % 4 == 0 && year % 100 != 0 || year % 400 == 0
    }
    switch month {
        case 2:
            if isLeapYear(year) { // 处理一下闰年的问题即可
                return 29
            }
            return 28
        case 1,3,5,7,8,10,12:
            return 31
        case 4,6,9,11:
            return 30
    }
    return -1
}

func main() {
    // Example 1:
    // Input: year = 1992, month = 7
    // Output: 31
    fmt.Println("1992-7 => ", numberOfDays(1992,7)) // 31
    // Example 2:
    // Input: year = 2000, month = 2
    // Output: 29
    fmt.Println("2000-2 => ", numberOfDays(2000,2)) // 29
    // Example 3:
    // Input: year = 1900, month = 2
    // Output: 28
    fmt.Println("1900-2 => ", numberOfDays(1900,2)) // 28
}