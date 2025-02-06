package main

// 3114. Latest Time You Can Obtain After Replacing Characters
// You are given a string s representing a 12-hour format time where some of the digits (possibly none) are replaced with a "?".

// 12-hour times are formatted as "HH:MM", where HH is between 00 and 11, and MM is between 00 and 59. 
// The earliest 12-hour time is 00:00, and the latest is 11:59.

// You have to replace all the "?" characters in s with digits such that the time we obtain by the resulting string is a valid 12-hour format time and is the latest possible.

// Return the resulting string.

// Example 1:
// Input: s = "1?:?4"
// Output: "11:54"
// Explanation: The latest 12-hour format time we can achieve by replacing "?" characters is "11:54".

// Example 2:
// Input: s = "0?:5?"
// Output: "09:59"
// Explanation: The latest 12-hour format time we can achieve by replacing "?" characters is "09:59".

// Constraints:
//     s.length == 5
//     s[2] is equal to the character ":".
//     All characters except s[2] are digits or "?" characters.
//     The input is generated such that there is at least one time between "00:00" and "11:59" that you can obtain after replacing the "?" characters.

import "fmt"

func findLatestTime(s string) string {
    res := []byte(s)
    for i := 0; i < len(res); i++ {
        if res[i] == '?' {
            switch i {
            case 0:
                if res[1] == '0' || res[1] == '1' || res[1] == '?' {
                    res[i] = '1'
                } else {
                    res[i] = '0'
                }
            case 1:
                if res[0] == '1' {
                    res[1] = '1'
                } else {
                    res[1] = '9'
                }
            case 3:
                res[i] = '5'
            case 4:
                res[i] = '9'
            }
        }
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: s = "1?:?4"
    // Output: "11:54"
    // Explanation: The latest 12-hour format time we can achieve by replacing "?" characters is "11:54".
    fmt.Println(findLatestTime("1?:?4")) // "11:54"
    // Example 2:
    // Input: s = "0?:5?"
    // Output: "09:59"
    // Explanation: The latest 12-hour format time we can achieve by replacing "?" characters is "09:59".
    fmt.Println(findLatestTime("0?:5?")) // "09:59"
}