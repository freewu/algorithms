package main

// 2849. Determine if a Cell Is Reachable at a Given Time
// You are given four integers sx, sy, fx, fy, and a non-negative integer t.

// In an infinite 2D grid, you start at the cell (sx, sy). 
// Each second, you must move to any of its adjacent cells.

// Return true if you can reach cell (fx, fy) after exactly t seconds, or false otherwise.

// A cell's adjacent cells are the 8 cells around it that share at least one corner with it. 
// You can visit the same cell several times.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/08/05/example2.svg" />
// Input: sx = 2, sy = 4, fx = 7, fy = 7, t = 6
// Output: true
// Explanation: Starting at cell (2, 4), we can reach cell (7, 7) in exactly 6 seconds by going through the cells depicted in the picture above. 

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/08/05/example1.svg" />
// Input: sx = 3, sy = 1, fx = 7, fy = 3, t = 3
// Output: false
// Explanation: Starting at cell (3, 1), it takes at least 4 seconds to reach cell (7, 3) by going through the cells depicted in the picture above. Hence, we cannot reach cell (7, 3) at the third second.

// Constraints:
//     1 <= sx, sy, fx, fy <= 10^9
//     0 <= t <= 10^9

import "fmt"

func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    if sx == fx && sy == fy { return t == 0 || t > 1 }
    return t >= max(abs(fx - sx), abs(fy - sy))
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2023/08/05/example2.svg" />
    // Input: sx = 2, sy = 4, fx = 7, fy = 7, t = 6
    // Output: true
    // Explanation: Starting at cell (2, 4), we can reach cell (7, 7) in exactly 6 seconds by going through the cells depicted in the picture above. 
    fmt.Println(isReachableAtTime(2, 4, 7, 7, 6)) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2023/08/05/example1.svg" />
    // Input: sx = 3, sy = 1, fx = 7, fy = 3, t = 3
    // Output: false
    // Explanation: Starting at cell (3, 1), it takes at least 4 seconds to reach cell (7, 3) by going through the cells depicted in the picture above. Hence, we cannot reach cell (7, 3) at the third second.
    fmt.Println(isReachableAtTime(3, 1, 7, 3, 3)) // false

    fmt.Println(isReachableAtTime(1, 1, 1, 1, 0)) // true
    fmt.Println(isReachableAtTime(1_000_000_000, 1_000_000_000, 1_000_000_000, 1_000_000_000, 1_000_000_000)) // true
    fmt.Println(isReachableAtTime(1, 1_000_000_000, 1_000_000_000, 1_000_000_000, 1_000_000_000)) // true
    fmt.Println(isReachableAtTime(1_000_000_000, 1, 1_000_000_000, 1_000_000_000, 1_000_000_000)) // true
    fmt.Println(isReachableAtTime(1_000_000_000, 1_000_000_000, 1, 1_000_000_000, 1_000_000_000)) // true
    fmt.Println(isReachableAtTime(1_000_000_000, 1_000_000_000, 1_000_000_000, 1, 1_000_000_000)) // true
    fmt.Println(isReachableAtTime(1_000_000_000, 1_000_000_000, 1_000_000_000, 1_000_000_000, 0)) // true
    fmt.Println(isReachableAtTime(1, 1, 1_000_000_000, 1_000_000_000, 1_000_000_000)) // true
    fmt.Println(isReachableAtTime(1_000_000_000, 1, 1, 1_000_000_000, 1_000_000_000)) // true
    fmt.Println(isReachableAtTime(1_000_000_000, 1_000_000_000, 1, 1, 1_000_000_000)) // true
    fmt.Println(isReachableAtTime(1_000_000_000, 1_000_000_000, 1_000_000_000, 1, 0)) // false
    fmt.Println(isReachableAtTime(1, 1_000_000_000, 1_000_000_000, 1_000_000_000, 0)) // false
    fmt.Println(isReachableAtTime(1, 1, 1, 1_000_000_000, 1_000_000_000)) // true
    fmt.Println(isReachableAtTime(1_000_000_000, 1, 1, 1, 1_000_000_000)) // true
    fmt.Println(isReachableAtTime(1_000_000_000, 1_000_000_000, 1, 1, 0)) // false
}