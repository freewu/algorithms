package main

// 3921. Score Validator
// You are given a string array events.

// Initially, score = 0 and counter = 0. Each element in events is one of the following:
//     1. "0", "1", "2", "3", "4", "6": Add that value to the total score.
//     2. "W": Increase the counter by 1. No score is added.
//     3. "WD": Add 1 to the total score.
//     4. "NB": Add 1 to the total score.

// Process the array from left to right. Stop processing when either:
//     1. All elements in events have been processed, or
//     2. The counter becomes 10.

// Return an integer array [score, counter], where:
//     1. score is the final total score.
//     2. counter is the final counter value.

// Example 1:
// Input: events = ["1","4","W","6","WD"]
// Output: [12,1]
// Explanation:
// Event | Score | Counter
// "1"	  |  1	  | 0
// "4"	  |  5	  | 0
// "W"	  |  5	  | 1
// "6"	  |  11	  | 1
// "WD"  |  12	  | 1
// Final result: [12, 1].

// Example 2:
// Input: events = ["WD","NB","0","4","4"]
// Output: [10,0]
// Explanation:
// Event | Score | Counter
// "WD"  |  1	  | 0
// "NB"  |  2	  | 0
// "0"	  |  2	  | 0
// "4"	  |  6	  | 0
// "4"	  |  10	  | 0
// Final result: [10, 0].

// Example 3:
// Input: events = ["W","W","W","W","W","W","W","W","W","W","W"]
// Output: [0,10]
// Explanation:
// After 10 occurrences of "W", the counter reaches 10, so processing stops. The remaining events are ignored.

// Constraints:
//     1 <= events.length <= 1000
//     events[i] is one of "0", "1", "2", "3", "4", "6", "W", "WD", or "NB".

import "fmt"

func scoreValidator(events []string) []int {
    count, total := 0, 0
    for _, event := range events {
        if event == "W" { // "W"：计数器加 1。不增加得分。
            count++
            if count == 10 { // 停止处理： events 中的所有元素都已处理完毕，或计数器变为 10。
                break
            }
        } else if event == "WD" { // "WD"：总得分加 1。
            total++
        } else if event == "NB" { // "NB"：总得分加 1。
            total++
        } else { // "0", "1", "2", "3", "4", "6"：将该值加到总得分中。
            total += int(event[0] - '0')
        }
    }
    return []int{ total, count } // 整数数组 [总得分, 计数器数值]
}

func main() {
    // Example 1:
    // Input: events = ["1","4","W","6","WD"]
    // Output: [12,1]
    // Explanation:
    // Event | Score | Counter
    // "1"	  |  1	  | 0
    // "4"	  |  5	  | 0
    // "W"	  |  5	  | 1
    // "6"	  |  11	  | 1
    // "WD"  |  12	  | 1
    // Final result: [12, 1].
    fmt.Println(scoreValidator([]string{"1","4","W","6","WD"})) // [12,1]
    // Example 2:
    // Input: events = ["WD","NB","0","4","4"]
    // Output: [10,0]
    // Explanation:
    // Event | Score | Counter
    // "WD"  |  1	  | 0
    // "NB"  |  2	  | 0
    // "0"	  |  2	  | 0
    // "4"	  |  6	  | 0
    // "4"	  |  10	  | 0
    // Final result: [10, 0].
    fmt.Println(scoreValidator([]string{"WD","NB","0","4","4"})) // [10,0]
    // Example 3:
    // Input: events = ["W","W","W","W","W","W","W","W","W","W","W"]
    // Output: [0,10]
    // Explanation:
    // After 10 occurrences of "W", the counter reaches 10, so processing stops. The remaining events are ignored.
    fmt.Println(scoreValidator([]string{"W","W","W","W","W","W","W","W","W","W","W"})) // [0,10]
}