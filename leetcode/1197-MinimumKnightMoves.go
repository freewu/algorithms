package main

// 1197. Minimum Knight Moves
// In an infinite chess board with coordinates from -infinity to +infinity, you have a knight at square [0, 0].

// A knight has 8 possible moves it can make, as illustrated below. 
// Each move is two squares in a cardinal direction, then one square in an orthogonal direction.

// Return the minimum number of steps needed to move the knight to the square [x, y]. 
// It is guaranteed the answer exists.

// Example 1:
// Input: x = 2, y = 1
// Output: 1
// Explanation: [0, 0] → [2, 1]

// Example 2:
// Input: x = 5, y = 5
// Output: 4
// Explanation: [0, 0] → [2, 1] → [4, 2] → [3, 4] → [5, 5]

//     Constraints:
//     -300 <= x, y <= 300
//     0 <= |x| + |y| <= 300

import "fmt"

func minKnightMoves(x int, y int) int {
    if x == 0 && y == 0 {
        return 0
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    x, y = abs(x), abs(y)
    directs := [...][2]int{
        {2, 1}, {1, 2}, {-2, 1}, {-1, 2},
        {2, -1}, {1, -2}, {-2, -1}, {-1, -2},
    }
    visited, step, q  := make(map[[2]int]bool), 0, [][2]int{{0, 0}}
    visited[[2]int{0, 0}] = true
    for len(q) > 0 {
        step += 1
        nq := [][2]int{}
        for _, p := range q {
            for _, direct := range directs {
                nx, ny := p[0] + direct[0], p[1]+ direct[1]
                if visited[[2]int{nx, ny}] {
                    continue
                }
                if nx >= -5 && nx <= x + 5 && ny >= -5 && ny <= y + 5 {
                    if nx == x && ny == y {
                        return step
                    }
                    visited[[2]int{nx, ny}] = true 
                    nq = append(nq, [2]int{nx, ny})
                }
            }
        }
        q = nq
    }
    return -1
}

func main() {
    // Example 1:
    // Input: x = 2, y = 1
    // Output: 1
    // Explanation: [0, 0] → [2, 1]
    fmt.Println(minKnightMoves(2, 1)) // 1
    // Example 2:
    // Input: x = 5, y = 5
    // Output: 4
    // Explanation: [0, 0] → [2, 1] → [4, 2] → [3, 4] → [5, 5]
    fmt.Println(minKnightMoves(5, 5)) // 4
}