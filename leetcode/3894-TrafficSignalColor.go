package main

// 3894. Traffic Signal Color
// You are given an integer timer representing the remaining time (in seconds) on a traffic signal.
// The signal follows these rules:
//     1. If timer == 0, the signal is "Green"
//     2. If timer == 30, the signal is "Orange"
//     3. If 30 < timer <= 90, the signal is "Red"

// Return the current state of the signal. 
// If none of the above conditions are met, return "Invalid".

// Example 1:
// Input: timer = 60
// Output: "Red"
// Explanation:
// Since timer = 60, and 30 < timer <= 90, the answer is "Red".

// Example 2:
// Input: timer = 5
// Output: "Invalid"
// Explanation:
// Since timer = 5, it does not satisfy any of the given conditions, the answer is "Invalid".
 
// Constraints:
//     0 <= timer <= 1000

import "fmt"

func trafficSignal(timer int) string {
    if timer == 0 {
        return "Green"
    }
    if timer == 30 {
        return "Orange"
    }
    if timer > 30 && timer <= 90 {
        return "Red"
    }
    return "Invalid"
}

func main() {
    // Example 1:
    // Input: timer = 60
    // Output: "Red"
    // Explanation:
    // Since timer = 60, and 30 < timer <= 90, the answer is "Red".
    fmt.Println(trafficSignal(60)) // "Red"
    // Example 2:
    // Input: timer = 5
    // Output: "Invalid"
    // Explanation:
    // Since timer = 5, it does not satisfy any of the given conditions, the answer is "Invalid".
    fmt.Println(trafficSignal(5)) // "Invalid"

    fmt.Println(trafficSignal(0)) // "Green"
    fmt.Println(trafficSignal(1)) // "Invalid"
    fmt.Println(trafficSignal(2)) // "Invalid"
    fmt.Println(trafficSignal(8)) // "Invalid"
    fmt.Println(trafficSignal(30)) // "Orange"
    fmt.Println(trafficSignal(64)) // "Red"
    fmt.Println(trafficSignal(99)) // "Invalid"
    fmt.Println(trafficSignal(100)) // "Invalid"
    fmt.Println(trafficSignal(999)) // "Invalid"
    fmt.Println(trafficSignal(1000)) // "Invalid"
}