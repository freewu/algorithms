package main 

// 1736. Latest Time by Replacing Hidden Digits
// You are given a string time in the form of hh:mm, 
// where some of the digits in the string are hidden (represented by ?).

// The valid times are those inclusively between 00:00 and 23:59.

// Return the latest valid time you can get from time by replacing the hidden digits.

// Example 1:
// Input: time = "2?:?0"
// Output: "23:50"
// Explanation: The latest hour beginning with the digit '2' is 23 and the latest minute ending with the digit '0' is 50.

// Example 2:
// Input: time = "0?:3?"
// Output: "09:39"

// Example 3:
// Input: time = "1?:22"
// Output: "19:22"

// Constraints:
//     time is in the format hh:mm.
//     It is guaranteed that you can produce a valid time from the given string.

import "fmt"
import "strings"

func maximumTime(time string) string {
    res := []byte(time)
    if res[0] == '?' {
        if res[1] == '?' || (res[1] - 48) < 4 {
            res[0] = '2'
        } else {
            res[0] = '1'
        }
    }
    if res[1] == '?' {
        if res[0] == '2' {
            res[1] = '3'
        } else {
            res[1] = '9'
        }
    }
    if res[3] == '?' { res[3] = '5' }
    if res[4] == '?' { res[4] = '9' }
    return string(res)
}

func maximumTime1(time string) string {
    res := make([]byte, 5)
    res[2] = ':'
    for i := 0; i < 5; i++ {
        switch i {
            case 0:
                if time[i] == '?' {
                    if strings.Contains("0123?", string(time[1])) {
                        res[0] = '2'
                    } else {
                        res[0] = '1'
                    }
                } else {
                    res[0] = time[i]
                }
            case 1:
                if time[i] == '?' {
                    if res[0] == '2' {
                        res[1] = '3'
                    } else {
                        res[1] = '9'
                    }
                } else {
                    res[1] = time[i]
                }
            case 3:
                if time[i] == '?' {
                    res[3] = '5'
                } else {
                    res[3] = time[i]
                }
            case 4:
                if time[i] == '?' {
                    res[4] = '9'
                } else {
                    res[4] = time[i]
                }
        }
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: time = "2?:?0"
    // Output: "23:50"
    // Explanation: The latest hour beginning with the digit '2' is 23 and the latest minute ending with the digit '0' is 50.
    fmt.Println(maximumTime("2?:?0")) // "23:50"
    // Example 2:
    // Input: time = "0?:3?"
    // Output: "09:39"
    fmt.Println(maximumTime("0?:3?")) // "09:39"
    // Example 3:
    // Input: time = "1?:22"
    // Output: "19:22"
    fmt.Println(maximumTime("1?:22")) // "19:22"
    fmt.Println(maximumTime("??:??")) // "23:59"

    fmt.Println(maximumTime1("2?:?0")) // "23:50"
    fmt.Println(maximumTime1("0?:3?")) // "09:39"
    fmt.Println(maximumTime1("1?:22")) // "19:22"
    fmt.Println(maximumTime1("??:??")) // "23:59"
}