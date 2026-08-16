package main

// 4020. Elevator Requests I
// You are given an integer n denoting the number of floors in a building, where the floors are numbered from 0 to n - 1.

// You are also given an integer array requests, where requests represents the sequence of floor requests.

// An elevator starts at floor 0, and follows these rules:
//     1. The elevator moves one floor per second.
//     2. The elevator serves requests in the given order.
//     3. If the elevator is already on the requested floor, no movement is needed.
//     4. After serving a request, the elevator immediately starts moving toward the next request.

// Return the total time (in seconds) required to serve all requests.

// Example 1:
// Input: n = 5, requests = [2,1,4,3]
// Output: 7
// Explanation:
// requests[0] = 2: Moving from floor 0 to floor 2 takes 2 seconds.
// requests[1] = 1: Moving from floor 2 to floor 1 takes 1 second.
// requests[2] = 4: Moving from floor 1 to floor 4 takes 3 seconds.
// requests[3] = 3: Moving from floor 4 to floor 3 takes 1 second.
// The total time required is 2 + 1 + 3 + 1 = 7 seconds.

// Example 2:
// Input: n = 3, requests = [2,0,0]
// Output: 4
// Explanation:
// requests[0] = 2: Moving from floor 0 to floor 2 takes 2 seconds.
// requests[1] = 0: Moving from floor 2 to floor 0 takes 2 seconds.
// requests[2] = 0: No movement is needed.
// The total time required is 2 + 2 + 0 = 4 seconds.

// Constraints:
//     1 <= n <= 100
//     1 <= requests.length <= 100
//     0 <= requests[i] <= n - 1

import "fmt"

func elevatorRequests(n int, requests []int) int {
    res, curr := 0, 0
    for _, floor := range requests {
        if curr > floor {
            res += curr - floor
        } else {
            res += floor - curr
        }
        curr = floor
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 5, requests = [2,1,4,3]
    // Output: 7
    // Explanation:
    // requests[0] = 2: Moving from floor 0 to floor 2 takes 2 seconds.
    // requests[1] = 1: Moving from floor 2 to floor 1 takes 1 second.
    // requests[2] = 4: Moving from floor 1 to floor 4 takes 3 seconds.
    // requests[3] = 3: Moving from floor 4 to floor 3 takes 1 second.
    // The total time required is 2 + 1 + 3 + 1 = 7 seconds.
    fmt.Println(elevatorRequests(5, []int{2,1,4,3})) // 7
    // Example 2:
    // Input: n = 3, requests = [2,0,0]
    // Output: 4
    // Explanation:
    // requests[0] = 2: Moving from floor 0 to floor 2 takes 2 seconds.
    // requests[1] = 0: Moving from floor 2 to floor 0 takes 2 seconds.
    // requests[2] = 0: No movement is needed.
    // The total time required is 2 + 2 + 0 = 4 seconds.
    fmt.Println(elevatorRequests(3, []int{2,0,0})) // 4

    fmt.Println(elevatorRequests(3, []int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(elevatorRequests(3, []int{9,8,7,6,5,4,3,2,1})) // 17
}