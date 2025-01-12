package main

// 2446. Determine if Two Events Have Conflict
// You are given two arrays of strings that represent two inclusive events that happened on the same day, event1 and event2, where:
//     event1 = [startTime1, endTime1] and
//     event2 = [startTime2, endTime2].

// Event times are valid 24 hours format in the form of HH:MM.

// A conflict happens when two events have some non-empty intersection (i.e., some moment is common to both events).

// Return true if there is a conflict between two events. Otherwise, return false.

// Example 1:
// Input: event1 = ["01:15","02:00"], event2 = ["02:00","03:00"]
// Output: true
// Explanation: The two events intersect at time 2:00.

// Example 2:
// Input: event1 = ["01:00","02:00"], event2 = ["01:20","03:00"]
// Output: true
// Explanation: The two events intersect starting from 01:20 to 02:00.

// Example 3:
// Input: event1 = ["10:00","11:00"], event2 = ["14:00","15:00"]
// Output: false
// Explanation: The two events do not intersect.

// Constraints:
//     event1.length == event2.length == 2
//     event1[i].length == event2[i].length == 5
//     startTime1 <= endTime1
//     startTime2 <= endTime2
//     All the event times follow the HH:MM format.

import "fmt"
import "strings"
import "strconv"

func haveConflict(event1 []string, event2 []string) bool {
    calc := func(event1 string) int {
        arr := strings.Split(event1, ":")
        h, _ := strconv.Atoi(arr[0])
        m, _ := strconv.Atoi(arr[1])
        return h * 60 + m
    }
    if (calc(event1[1]) < calc(event2[0]) && calc(event1[1]) < calc(event2[1])) || 
       (calc(event2[1]) < calc(event1[0]) && calc(event2[1]) < calc(event1[1])) {
        return false
    }
    return true
}

func haveConflict1(event1 []string, event2 []string) bool {
    return event2[0] <= event1[1] && event1[0] <= event2[1] 
}

func main() {
    // Example 1:
    // Input: event1 = ["01:15","02:00"], event2 = ["02:00","03:00"]
    // Output: true
    // Explanation: The two events intersect at time 2:00.
    fmt.Println(haveConflict([]string{"01:15","02:00"}, []string{"02:00","03:00"})) // true
    // Example 2:
    // Input: event1 = ["01:00","02:00"], event2 = ["01:20","03:00"]
    // Output: true
    // Explanation: The two events intersect starting from 01:20 to 02:00.
    fmt.Println(haveConflict([]string{"01:00","02:00"}, []string{"01:20","03:00"})) // true
    // Example 3:
    // Input: event1 = ["10:00","11:00"], event2 = ["14:00","15:00"]
    // Output: false
    // Explanation: The two events do not intersect.
    fmt.Println(haveConflict([]string{"10:00","11:00"}, []string{"14:00","15:00"})) // false

    fmt.Println(haveConflict1([]string{"01:15","02:00"}, []string{"02:00","03:00"})) // true
    fmt.Println(haveConflict1([]string{"01:00","02:00"}, []string{"01:20","03:00"})) // true
    fmt.Println(haveConflict1([]string{"10:00","11:00"}, []string{"14:00","15:00"})) // false
}