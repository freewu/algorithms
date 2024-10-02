package main

// 1496. Path Crossing
// Given a string path, where path[i] = 'N', 'S', 'E' or 'W', each representing moving one unit north, south, east, or west, respectively. 
// You start at the origin (0, 0) on a 2D plane and walk on the path specified by path.

// Return true if the path crosses itself at any point, that is, if at any time you are on a location you have previously visited. 
// Return false otherwise.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/06/10/screen-shot-2020-06-10-at-123929-pm.png" />
// Input: path = "NES"
// Output: false 
// Explanation: Notice that the path doesn't cross any point more than once.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/06/10/screen-shot-2020-06-10-at-123843-pm.png" />
// Input: path = "NESWW"
// Output: true
// Explanation: Notice that the path visits the origin twice.

// Constraints:
//     1 <= path.length <= 10^4
//     path[i] is either 'N', 'S', 'E', or 'W'.

import "fmt"

func isPathCrossing(path string) bool {
    mp, point := make(map[[2]int]bool), [2]int{0,0}
    mp[point] = true
    for i := 0 ; i < len(path) ; i++ {
        switch path[i] {
            case 'S': point[1]++ 
            case 'N': point[1]--
            case 'E': point[0]++
            case 'W': point[0]--
        }
        if mp[point] {
            return true
        }
        mp[point] = true
    }
    return false
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/06/10/screen-shot-2020-06-10-at-123929-pm.png" />
    // Input: path = "NES"
    // Output: false 
    // Explanation: Notice that the path doesn't cross any point more than once.
    fmt.Println(isPathCrossing("NES")) // false
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/06/10/screen-shot-2020-06-10-at-123843-pm.png" />
    // Input: path = "NESWW"
    // Output: true
    // Explanation: Notice that the path visits the origin twice.
    fmt.Println(isPathCrossing("NESWW")) // true
}