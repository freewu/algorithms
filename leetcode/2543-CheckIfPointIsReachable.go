package main

// 2543. Check if Point Is Reachable
// There exists an infinitely large grid. 
// You are currently at point (1, 1), and you need to reach the point (targetX, targetY) using a finite number of steps.

// In one step, you can move from point (x, y) to any one of the following points:
//     (x, y - x)
//     (x - y, y)
//     (2 * x, y)
//     (x, 2 * y)

// Given two integers targetX and targetY representing the X-coordinate and Y-coordinate of your final position, 
// return true if you can reach the point from (1, 1) using some number of steps, and false otherwise.

// Example 1:
// Input: targetX = 6, targetY = 9
// Output: false
// Explanation: It is impossible to reach (6,9) from (1,1) using any sequence of moves, so false is returned.

// Example 2:
// Input: targetX = 4, targetY = 7
// Output: true
// Explanation: You can follow the path (1,1) -> (1,2) -> (1,4) -> (1,8) -> (1,7) -> (2,7) -> (4,7).

// Constraints:
//     1 <= targetX, targetY <= 10^9

import "fmt"

func isReachable(targetX int, targetY int) bool {
    for targetY != 0 { // Find greatest common divisor (GCD) using the Euclidean algorithm
        targetX, targetY = targetY, targetX % targetY
    }
    return targetX & (targetX - 1) == 0 // Check if GCD is a power of 2
}

func isReachable1(targetX int, targetY int) bool {
    for targetX % 2 == 0 { targetX /= 2 }
    for targetY % 2 == 0 { targetY /= 2 }
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    return gcd(targetX, targetY) == 1
}

func main() {
    // Example 1:
    // Input: targetX = 6, targetY = 9
    // Output: false
    // Explanation: It is impossible to reach (6,9) from (1,1) using any sequence of moves, so false is returned.
    fmt.Println(isReachable(6, 9)) // false
    // Example 2:
    // Input: targetX = 4, targetY = 7
    // Output: true
    // Explanation: You can follow the path (1,1) -> (1,2) -> (1,4) -> (1,8) -> (1,7) -> (2,7) -> (4,7).
    fmt.Println(isReachable(4, 7)) // true

    fmt.Println(isReachable(1, 1)) // true
    fmt.Println(isReachable(1, 1_000_000_000)) // true
    fmt.Println(isReachable(1_000_000_000, 1)) // true
    fmt.Println(isReachable(1_000_000_000, 1_000_000_000)) // false

    fmt.Println(isReachable1(6, 9)) // false
    fmt.Println(isReachable1(4, 7)) // true
    fmt.Println(isReachable1(1, 1)) // true
    fmt.Println(isReachable1(1, 1_000_000_000)) // true
    fmt.Println(isReachable1(1_000_000_000, 1)) // true
    fmt.Println(isReachable1(1_000_000_000, 1_000_000_000)) // false
}