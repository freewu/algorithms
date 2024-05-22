package main

// 1360. Number of Days Between Two Dates
// Write a program to count the number of days between two dates.
// The two dates are given as strings, their format is YYYY-MM-DD as shown in the examples.

// Example 1:
// Input: date1 = "2019-06-29", date2 = "2019-06-30"
// Output: 1

// Example 2:
// Input: date1 = "2020-01-15", date2 = "2019-12-31"
// Output: 15
 
// Constraints:
//     The given dates are valid dates between the years 1971 and 2100.

import "fmt"
import "time"
import "math"

import "strings"
import "strconv"

// use time lib
func daysBetweenDates(date1 string, date2 string) int {
    t1, _ := time.Parse(time.RFC3339, date1 + "T00:00:00.000Z")
    t2, _ := time.Parse(time.RFC3339, date2 + "T00:00:00.000Z")
    return int(math.Abs(t2.Sub(t1).Hours() / 24))
}

func daysBetweenDates1(date1 string, date2 string) int {
    isLeap := func (year int) (int, bool){ // 是否闰年
        // check if year is leapyear and add the leap year days in-between years
        checkLeap := year % 400==0 || (year % 4 == 0 && year % 100 != 0)
        return year / 4 - year / 100 + year / 400, checkLeap
    }
    getDates := func (date string) int{
        year, d1, months := 0, []int{}, []int{ 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
        str := strings.Split(date, "-")
        for _, i := range str {
            conv , _ := strconv.Atoi(i)
            d1 = append(d1, conv)
        }
        year += d1[0]*365
        for i := 0; i < d1[1]-1; i++{
            year += months[i] 
        }
        leapCount, leapYear := isLeap(d1[0])
        if (d1[1] <= 2 && leapYear) {// 闰年的处理
            leapCount--
        }    
        year += leapCount
        return year + d1[2]
    }
    d1, d2 := getDates(date1), getDates(date2)
    if d2 > d1 {
        return d2 - d1
    }
    return d1 - d2
}

func main() {
    // Example 1:
    // Input: date1 = "2019-06-29", date2 = "2019-06-30"
    // Output: 1
    fmt.Println(daysBetweenDates("2019-06-29", "2019-06-30")) // 1
    // Example 2:
    // Input: date1 = "2020-01-15", date2 = "2019-12-31"
    // Output: 15
    fmt.Println(daysBetweenDates("2020-01-15", "2019-12-31")) // 15

    fmt.Println(daysBetweenDates1("2019-06-29", "2019-06-30")) // 1
    fmt.Println(daysBetweenDates1("2020-01-15", "2019-12-31")) // 15
}