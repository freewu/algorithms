package main

// 789. Escape The Ghosts
// You are playing a simplified PAC-MAN game on an infinite 2-D grid. 
// You start at the point [0, 0], and you are given a destination point target = [xtarget, ytarget] that you are trying to get to. 
// There are several ghosts on the map with their starting positions given as a 2D array ghosts, where ghosts[i] = [xi, yi] represents the starting position of the ith ghost. 
// All inputs are integral coordinates.

// Each turn, you and all the ghosts may independently choose to either move 1 unit in any of the four cardinal directions: north, east, south, or west, or stay still. All actions happen simultaneously.

// You escape if and only if you can reach the target before any ghost reaches you. 
// If you reach any square (including the target) at the same time as a ghost, it does not count as an escape.

// Return true if it is possible to escape regardless of how the ghosts move, otherwise return false.

// Example 1:
// Input: ghosts = [[1,0],[0,3]], target = [0,1]
// Output: true
// Explanation: You can reach the destination (0, 1) after 1 turn, while the ghosts located at (1, 0) and (0, 3) cannot catch up with you.

// Example 2:
// Input: ghosts = [[1,0]], target = [2,0]
// Output: false
// Explanation: You need to reach the destination (2, 0), but the ghost at (1, 0) lies between you and the destination.

// Example 3:
// Input: ghosts = [[2,0]], target = [1,0]
// Output: false
// Explanation: The ghost can reach the target at the same time as you.

// Constraints:
//     1 <= ghosts.length <= 100
//     ghosts[i].length == 2
//     -10^4 <= xi, yi <= 10^4
//     There can be multiple ghosts in the same location.
//     target.length == 2
//     -10^4 <= xtarget, ytarget <= 10^4

import "fmt"
import "math"

// O(n) time, O(1) space,
// Approach: math, manhattan distance
func escapeGhosts(ghosts [][]int, target []int) bool {
    manhattanDistance := func (x int, y int, x0 int, y0 int) int { return int(math.Abs(float64(x-x0)) + math.Abs(float64(y-y0))); }
    pacman_dist := manhattanDistance(0, 0, target[0], target[1]) 
    for _, ghost := range ghosts {
        ghost_dist := manhattanDistance(ghost[0], ghost[1], target[0], target[1])
        if ghost_dist <= pacman_dist {
            return false
        }
    }
    return true
}

func escapeGhosts1(ghosts [][]int, target []int) bool {
    tx, ty := target[0], target[1]
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    dis := abs(tx) + abs(ty)
    for _, ghost := range ghosts {
        x, y := ghost[0], ghost[1]
        d := abs(tx-x) + abs(ty-y)
        if d <= dis {
            return false
        } 
    }
    return true
}

func main() {
    // Example 1:
    // Input: ghosts = [[1,0],[0,3]], target = [0,1]
    // Output: true
    // Explanation: You can reach the destination (0, 1) after 1 turn, while the ghosts located at (1, 0) and (0, 3) cannot catch up with you.
    fmt.Println(escapeGhosts([][]int{{1,0},{0,3}},[]int{0, 1})) // true
    // Example 2:
    // Input: ghosts = [[1,0]], target = [2,0]
    // Output: false
    // Explanation: You need to reach the destination (2, 0), but the ghost at (1, 0) lies between you and the destination.
    fmt.Println(escapeGhosts([][]int{{1,0}},[]int{2,0})) // false
    // Example 3:
    // Input: ghosts = [[2,0]], target = [1,0]
    // Output: false
    // Explanation: The ghost can reach the target at the same time as you.
    fmt.Println(escapeGhosts([][]int{{2,0}},[]int{1,0})) // false

    fmt.Println(escapeGhosts1([][]int{{1,0},{0,3}},[]int{0, 1})) // true
    fmt.Println(escapeGhosts1([][]int{{1,0}},[]int{2,0})) // false
    fmt.Println(escapeGhosts1([][]int{{2,0}},[]int{1,0})) // false
}