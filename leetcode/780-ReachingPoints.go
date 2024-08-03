package main

// 780. Reaching Points
// Given four integers sx, sy, tx, and ty, return true if it is possible to convert the point (sx, sy) to the point (tx, ty) through some operations, or false otherwise.
// The allowed operation on some point (x, y) is to convert it to either (x, x + y) or (x + y, y).

// Example 1:
// Input: sx = 1, sy = 1, tx = 3, ty = 5
// Output: true
// Explanation:
// One series of moves that transforms the starting point to the target is:
// (1, 1) -> (1, 2)
// (1, 2) -> (3, 2)
// (3, 2) -> (3, 5)

// Example 2:
// Input: sx = 1, sy = 1, tx = 2, ty = 2
// Output: false

// Example 3:
// Input: sx = 1, sy = 1, tx = 1, ty = 1
// Output: true

// Constraints:
//     1 <= sx, sy, tx, ty <= 10^9

import "fmt"

// if tx > ty then tx = x + y, ty = y , else if ty > tx then tx = x , ty = x + y. 
// trace backwards until we find a solution.
func reachingPoints1(sx int, sy int, tx int, ty int) bool {
    if sx == tx && sy == ty { return true }
    if sx == tx {
        return ty - sy > 0 && (ty - sy) % sx == 0 // ty has to be bigger than sy
    } else if sy == ty {
        return tx - sx > 0 && (tx - sx) % sy == 0 // tx has to be bigger than sx
    }
    for tx >= sx && ty >= sy { // search backward
        if sx == tx && sy == ty { return true }
        if tx > ty {
            tx = tx - ty
        } else {
            ty = ty - tx
        }
    }
    return false
}

func reachingPoints(sx int, sy int, tx int, ty int) bool {
    if tx < sx || ty < sy{  return false }
    if sx == tx{
        return (ty - sy) % sx == 0
    } else if sy == ty {
        return (tx - sx) % sy == 0
    }
    if tx > ty { return reachingPoints(sx, sy, tx % ty, ty) } 
    return reachingPoints(sx, sy, tx, ty % tx)
}

func main() {
    // Example 1:
    // Input: sx = 1, sy = 1, tx = 3, ty = 5
    // Output: true
    // Explanation:
    // One series of moves that transforms the starting point to the target is:
    // (1, 1) -> (1, 2)
    // (1, 2) -> (3, 2)
    // (3, 2) -> (3, 5)
    fmt.Println(reachingPoints(1,1,3,5)) // true
    // Example 2:
    // Input: sx = 1, sy = 1, tx = 2, ty = 2
    // Output: false
    fmt.Println(reachingPoints(1,1,2,2)) // false
    // Example 3:
    // Input: sx = 1, sy = 1, tx = 1, ty = 1
    // Output: true
    fmt.Println(reachingPoints(1,1,1,1)) // true

    fmt.Println(reachingPoints1(1,1,3,5)) // true
    fmt.Println(reachingPoints1(1,1,2,2)) // false
    fmt.Println(reachingPoints1(1,1,1,1)) // true
}