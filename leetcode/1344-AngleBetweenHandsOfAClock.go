package main

// 1344. Angle Between Hands of a Clock
// Given two numbers, hour and minutes, 
// return the smaller angle (in degrees) formed between the hour and the minute hand.

// Answers within 10-5 of the actual value will be accepted as correct.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/12/26/sample_1_1673.png" />
// Input: hour = 12, minutes = 30
// Output: 165

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/12/26/sample_2_1673.png" />
// Input: hour = 3, minutes = 30
// Output: 75

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2019/12/26/sample_3_1673.png" />
// Input: hour = 3, minutes = 15
// Output: 7.5

// Constraints:
//     1 <= hour <= 12
//     0 <= minutes <= 59

import "fmt"
import "math"

func angleClock(hour int, minutes int) float64 {
    h := (float64(hour % 12) + float64(minutes) / 60) * 30
    m := float64(minutes) * 6
    angle := math.Abs(h - m)
    if angle > 180 { 
        return 360 - angle 
    }
    return angle
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/12/26/sample_1_1673.png" />
    // Input: hour = 12, minutes = 30
    // Output: 165
    fmt.Println(angleClock(12, 30)) // 165
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2019/12/26/sample_2_1673.png" />
    // Input: hour = 3, minutes = 30
    // Output: 75
    fmt.Println(angleClock(3, 30)) // 75
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2019/12/26/sample_3_1673.png" />
    // Input: hour = 3, minutes = 15
    // Output: 7.5
    fmt.Println(angleClock(3, 15)) // 7.5
}