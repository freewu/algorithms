package main

// 1185. Day of the Week
// Given a date, return the corresponding day of the week for that date.
// The input is given as three integers representing the day, month and year respectively.
// Return the answer as one of the following values {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}.

// Example 1:
// Input: day = 31, month = 8, year = 2019
// Output: "Saturday"

// Example 2:
// Input: day = 18, month = 7, year = 1999
// Output: "Sunday"

// Example 3:
// Input: day = 15, month = 8, year = 1993
// Output: "Sunday"

// Constraints:
//     The given dates are valid dates between the years 1971 and 2100.

import "fmt"
import "time"

func dayOfTheWeek(day int, month int, year int) string {
    return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).Format("Monday")
}

// Zeller's algorithm
func dayOfTheWeek1(day int, month int, year int) string {
    // https://en.wikipedia.org/wiki/Determination_of_the_day_of_the_week#Zeller's_algorithm
    days := []string{ "Saturday", "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", }
    y := year
    if month < 3 {
        y--
    }
    y, c, d, m := y % 100, y / 100, day, month
    if m < 3 {
        m += 12
    }
    w := (13 * (m + 1)) / 5
    w += y/4
    w += c/4
    w += d
    w += y
    w -= 2*c
    w %= 7
    if w < 0 {
        w += 7
    }
    return days[w]
}

func main() {
    // Example 1:
    // Input: day = 31, month = 8, year = 2019
    // Output: "Saturday"
    fmt.Println(dayOfTheWeek(31, 8, 2019)) // "Saturday"
    // Example 2:
    // Input: day = 18, month = 7, year = 1999
    // Output: "Sunday"
    fmt.Println(dayOfTheWeek(18, 7, 1999)) // "Sunday"
    // Example 3:
    // Input: day = 15, month = 8, year = 1993
    // Output: "Sunday"
    fmt.Println(dayOfTheWeek(15, 8, 1993)) // "Sunday"

    fmt.Println(dayOfTheWeek1(31, 8, 2019)) // "Saturday"
    fmt.Println(dayOfTheWeek1(18, 7, 1999)) // "Sunday"
    fmt.Println(dayOfTheWeek1(15, 8, 1993)) // "Sunday"
}