package main

// 949. Largest Time for Given Digits
// Given an array arr of 4 digits, find the latest 24-hour time that can be made using each digit exactly once.

// 24-hour times are formatted as "HH:MM", where HH is between 00 and 23, and MM is between 00 and 59. 
// The earliest 24-hour time is 00:00, and the latest is 23:59.

// Return the latest 24-hour time in "HH:MM" format. If no valid time can be made, return an empty string.

// Example 1:
// Input: arr = [1,2,3,4]
// Output: "23:41"
// Explanation: The valid 24-hour times are "12:34", "12:43", "13:24", "13:42", "14:23", "14:32", "21:34", "21:43", "23:14", and "23:41". Of these times, "23:41" is the latest.

// Example 2:
// Input: arr = [5,5,5,5]
// Output: ""
// Explanation: There are no valid 24-hour times as "55:55" is not valid.

// Constraints:
//     arr.length == 4
//     0 <= arr[i] <= 9

import "fmt"

func largestTimeFromDigits(arr []int) string {
    tranToTime := func (h1, h2, m1, m2 int) string {
        hour := h1 * 10 + h2
        if hour >= 24 { return "" }
        minute := m1 * 10 + m2
        if minute >= 60 { return "" }
        return fmt.Sprintf("%02d:%02d", hour, minute)
    }
    res := ""
    for x := 0; x < 4; x++ {
        for y := 0; y < 4; y++ {
            if y == x { continue  }
            for z := 0; z < 4; z++ {
                if z == y || z == x { continue }
                for m := 0; m < 4; m++ {
                    if m == z || m == y || m == x { continue }
                    if curTime := tranToTime(arr[x], arr[y], arr[z], arr[m]); curTime > res {
                        res = curTime
                    }
                }
            }
        }
    }
    return res
}

func largestTimeFromDigits1(arr []int) string {
    maxTime := -1
    checkTime := func(arr []int) {
        hour, minute := arr[0] * 10 + arr[1], arr[2] * 10 + arr[3]
        if hour < 24 && minute < 60 {
            curTime := hour * 60 + minute
            if curTime > maxTime {
                maxTime = curTime
            }
        }
    }
    var permute func(arr []int, start int)
    permute = func(arr []int, start int) {
        if start == len(arr) {
          checkTime(arr)
          return  
        }
        for i := start; i < len(arr); i++ {
            arr[start], arr[i] = arr[i], arr[start]
            permute(arr,start + 1)
            arr[start], arr[i] = arr[i], arr[start]
        }
    }
    permute(arr, 0)
    if maxTime == -1 { return "" }
    return fmt.Sprintf("%02d:%02d", maxTime / 60, maxTime % 60)
}

func main() {
    // Example 1:
    // Input: arr = [1,2,3,4]
    // Output: "23:41"
    // Explanation: The valid 24-hour times are "12:34", "12:43", "13:24", "13:42", "14:23", "14:32", "21:34", "21:43", "23:14", and "23:41". Of these times, "23:41" is the latest.
    fmt.Println(largestTimeFromDigits([]int{1,2,3,4})) // "23:41"
    // Example 2:
    // Input: arr = [5,5,5,5]
    // Output: ""
    // Explanation: There are no valid 24-hour times as "55:55" is not valid.
    fmt.Println(largestTimeFromDigits([]int{5,5,5,5})) // ""

    fmt.Println(largestTimeFromDigits([]int{1,2,3,0})) // "23:10"

    fmt.Println(largestTimeFromDigits1([]int{1,2,3,4})) // "23:41"
    fmt.Println(largestTimeFromDigits1([]int{5,5,5,5})) // ""
    fmt.Println(largestTimeFromDigits1([]int{1,2,3,0})) // "23:10"
}