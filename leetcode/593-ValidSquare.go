package main

// 593. Valid Square
// Given the coordinates of four points in 2D space p1, p2, p3 and p4, 
// return true if the four points construct a square.

// The coordinate of a point pi is represented as [xi, yi]. The input is not given in any order.
// A valid square has four equal sides with positive length and four equal angles (90-degree angles).

// Example 1:
// Input: p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,1]
// Output: true

// Example 2:
// Input: p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,12]
// Output: false

// Example 3:
// Input: p1 = [1,0], p2 = [-1,0], p3 = [0,1], p4 = [0,-1]
// Output: true

// Constraints:
//     p1.length == p2.length == p3.length == p4.length == 2
//     -10^4 <= xi, yi <= 10^4

import "fmt"

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
    sums := map[int]int{}
    points := [][]int{ p1, p2, p3, p4 }
    for i := range points {
        for k:= i+1; k < 4; k++ {
            xl, yl := points[i][0] - points[k][0], points[i][1] - points[k][1]
            sqSum := xl * xl + yl * yl
            if sqSum == 0 {
                return false
            }
            sums[sqSum] = 1
            if len(sums) > 2 {
                return false
            }
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,1]
    // Output: true
    fmt.Println(validSquare([]int{0,0},[]int{1,1},[]int{1,0},[]int{0,1})) // true
    // Example 2:
    // Input: p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,12]
    // Output: false
    fmt.Println(validSquare([]int{0,0},[]int{1,1},[]int{1,0},[]int{0,12})) // false
    // Example 3:
    // Input: p1 = [1,0], p2 = [-1,0], p3 = [0,1], p4 = [0,-1]
    // Output: true
    fmt.Println(validSquare([]int{1,0},[]int{-1,0},[]int{0,1},[]int{0,-1})) // true
}