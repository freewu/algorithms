package main

// 551. Student Attendance Record I
// You are given a string s representing an attendance record
// for a student where each character signifies whether the student was absent, late, or present on that day. 
// The record only contains the following three characters:
//     'A': Absent.
//     'L': Late.
//     'P': Present.

// The student is eligible for an attendance award if they meet both of the following criteria:
//     The student was absent ('A') for strictly fewer than 2 days total.
//     The student was never late ('L') for 3 or more consecutive days.

// Return true if the student is eligible for an attendance award, or false otherwise.

// Example 1:
// Input: s = "PPALLP"
// Output: true
// Explanation: The student has fewer than 2 absences and was never late 3 or more consecutive days.

// Example 2:
// Input: s = "PPALLL"
// Output: false
// Explanation: The student was late 3 consecutive days in the last 3 days, so is not eligible for the award.
 
// Constraints:
//     1 <= s.length <= 1000
//     s[i] is either 'A', 'L', or 'P'.

import "fmt"

func checkRecord(s string) bool {
    absent, late := 0, 0
    for i := 0; i < len(s); i++ {
        if s[i] == 'L' { // 迟到
            late++
            if late >= 3 { // The student was never late ('L') for 3 or more consecutive days.
                return false
            }
        } else {
            late = 0
        }
        if s[i] == 'A' { // 缺勤
            absent++ 
            if absent >= 2 { // The student was absent ('A') for strictly fewer than 2 days total.
                return false
            }
        }
        if s[i] == 'P' { // 正常出勤
            continue
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: s = "PPALLP"
    // Output: true
    // Explanation: The student has fewer than 2 absences and was never late 3 or more consecutive days.
    fmt.Println(checkRecord("PPALLP")) // true
    // Example 2:
    // Input: s = "PPALLL"
    // Output: false
    // Explanation: The student was late 3 consecutive days in the last 3 days, so is not eligible for the award.
    fmt.Println(checkRecord("PPALLL")) // false
    fmt.Println(checkRecord("AA")) // false
}