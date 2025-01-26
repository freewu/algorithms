package main

// 2651. Calculate Delayed Arrival Time
// You are given a positive integer arrivalTime denoting the arrival time of a train in hours, 
// and another positive integer delayedTime denoting the amount of delay in hours.

// Return the time when the train will arrive at the station.

// Note that the time in this problem is in 24-hours format.

// Example 1:
// Input: arrivalTime = 15, delayedTime = 5 
// Output: 20 
// Explanation: Arrival time of the train was 15:00 hours. It is delayed by 5 hours. Now it will reach at 15+5 = 20 (20:00 hours).

// Example 2:
// Input: arrivalTime = 13, delayedTime = 11
// Output: 0
// Explanation: Arrival time of the train was 13:00 hours. It is delayed by 11 hours. Now it will reach at 13+11=24 (Which is denoted by 00:00 in 24 hours format so return 0).

// Constraints:
//     1 <= arrivaltime < 24
//     1 <= delayedTime <= 24

import "fmt"

func findDelayedArrivalTime(arrivalTime int, delayedTime int) int {
    if arrivalTime + delayedTime > 24  { return (arrivalTime + delayedTime) - 24 }
    if arrivalTime + delayedTime == 24 { return 0 }
    return arrivalTime + delayedTime
}

func findDelayedArrivalTime1(arrivalTime int, delayedTime int) int {
    return (arrivalTime + delayedTime) % 24
}

func main() {
    // Example 1:
    // Input: arrivalTime = 15, delayedTime = 5 
    // Output: 20 
    // Explanation: Arrival time of the train was 15:00 hours. It is delayed by 5 hours. Now it will reach at 15+5 = 20 (20:00 hours).
    fmt.Println(findDelayedArrivalTime(15, 5)) // 20
    // Example 2:
    // Input: arrivalTime = 13, delayedTime = 11
    // Output: 0
    // Explanation: Arrival time of the train was 13:00 hours. It is delayed by 11 hours. Now it will reach at 13+11=24 (Which is denoted by 00:00 in 24 hours format so return 0).
    fmt.Println(findDelayedArrivalTime(13, 11)) // 0

    fmt.Println(findDelayedArrivalTime(1, 1)) // 2
    fmt.Println(findDelayedArrivalTime(24, 24)) // 24
    fmt.Println(findDelayedArrivalTime(1, 24)) // 1
    fmt.Println(findDelayedArrivalTime(24, 1)) // 1

    fmt.Println(findDelayedArrivalTime1(15, 5)) // 20
    fmt.Println(findDelayedArrivalTime1(13, 11)) // 0
    fmt.Println(findDelayedArrivalTime1(1, 1)) // 2
    fmt.Println(findDelayedArrivalTime1(24, 24)) // 24
    fmt.Println(findDelayedArrivalTime1(1, 24)) // 1
    fmt.Println(findDelayedArrivalTime1(24, 1)) // 1
}