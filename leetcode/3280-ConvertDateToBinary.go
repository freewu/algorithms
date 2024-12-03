package main

// 3280. Convert Date to Binary
// You are given a string date representing a Gregorian calendar date in the yyyy-mm-dd format.

// date can be written in its binary representation obtained by converting year, month, and day 
// to their binary representations without any leading zeroes and writing them down in year-month-day format.

// Return the binary representation of date.

// Example 1:
// Input: date = "2080-02-29"
// Output: "100000100000-10-11101"
// Explanation:
// 100000100000, 10, and 11101 are the binary representations of 2080, 02, and 29 respectively.

// Example 2:
// Input: date = "1900-01-01"
// Output: "11101101100-1-1"
// Explanation:
// 11101101100, 1, and 1 are the binary representations of 1900, 1, and 1 respectively.

// Constraints:
//     date.length == 10
//     date[4] == date[7] == '-', and all other date[i]'s are digits.
//     The input is generated such that date represents a valid Gregorian calendar date between Jan 1st, 1900 and Dec 31st, 2100 (both inclusive).

import "fmt"
import "strings"
import "strconv"

func convertDateToBinary(date string) string {
    ymd := strings.Split(date, "-")
    y, _ := strconv.Atoi(ymd[0])
    m, _ := strconv.Atoi(ymd[1])
    d, _ := strconv.Atoi(ymd[2])
    return fmt.Sprintf("%b-%b-%b", y, m , d)
}

func convertDateToBinary1(date string) string {
    year := 1000 * int(date[0] - '0') + 100 * int(date[1] - '0') + 10 * int(date[2] - '0') + int(date[3] - '0')
    month := 10 * int(date[5] - '0') + int(date[6] - '0')
    day := 10 * int(date[8] - '0') + int(date[9] - '0')
    return fmt.Sprintf("%b-%b-%b", year, month, day)
}

func main() {
    // Example 1:
    // Input: date = "2080-02-29"
    // Output: "100000100000-10-11101"
    // Explanation:
    // 100000100000, 10, and 11101 are the binary representations of 2080, 02, and 29 respectively.
    fmt.Println(convertDateToBinary("2080-02-29")) // "100000100000-10-11101"
    // Example 2:
    // Input: date = "1900-01-01"
    // Output: "11101101100-1-1"
    // Explanation:
    // 11101101100, 1, and 1 are the binary representations of 1900, 1, and 1 respectively.
    fmt.Println(convertDateToBinary("1900-01-01")) // "11101101100-1-1"

    fmt.Println(convertDateToBinary("2014-12-04")) // "11111011110-1100-100"

    fmt.Println(convertDateToBinary1("2080-02-29")) // "100000100000-10-11101"
    fmt.Println(convertDateToBinary1("1900-01-01")) // "11101101100-1-1"
    fmt.Println(convertDateToBinary1("2014-12-04")) // "11111011110-1100-100"
}