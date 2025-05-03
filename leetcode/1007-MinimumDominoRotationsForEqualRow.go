package main

// 1007. Minimum Domino Rotations For Equal Row
// In a row of dominoes, tops[i] and bottoms[i] represent the top and bottom halves of the ith domino. 
// (A domino is a tile with two numbers from 1 to 6 - one on each half of the tile.)

// We may rotate the ith domino, so that tops[i] and bottoms[i] swap values.

// Return the minimum number of rotations so that all the values in tops are the same, 
// or all the values in bottoms are the same.

// If it cannot be done, return -1.

// Example 1:
// Input: tops = [2,1,2,4,2,2], bottoms = [5,2,6,2,3,2]
// Output: 2
// Explanation: 
// The first figure represents the dominoes as given by tops and bottoms: before we do any rotations.
// If we rotate the second and fourth dominoes, we can make every value in the top row equal to 2, as indicated by the second figure.

// Example 2:
// Input: tops = [3,5,1,2,3], bottoms = [3,6,3,3,4]
// Output: -1
// Explanation: 
// In this case, it is not possible to rotate the dominoes to make one row of values equal.

// Constraints:
//     2 <= tops.length <= 2 * 10^4
//     bottoms.length == tops.length
//     1 <= tops[i], bottoms[i] <= 6

import "fmt"

func minDominoRotations(tops []int, bottoms []int) int {
    countTop, countBottom, countSame := make([]int, 7), make([]int, 7), make([]int, 7)
    res, n := len(tops), len(tops)
    for i := 0; i < n; i++ {
        t, b := tops[i], bottoms[i]
        countTop[t]++
        countBottom[b]++
        if t == b {
            countSame[t]++
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i <= 6; i++ {
        if countTop[i] + countBottom[i] - countSame[i] == n {
            res = min(res, min(countTop[i], countBottom[i]) - countSame[i])
        }
    }
    if res == n {
        return -1
    }
    return res
}

func minDominoRotations1(tops []int, bottoms []int) int {
    n, x, y := len(tops), tops[0], bottoms[0]
    cnt_x1, cnt_x2, cnt_y1, cnt_y2 := 0, 0, 0, 0
    // top = x
    for i := 0; i < n; i++ {
        if tops[i] == x { continue }
        if bottoms[i] == x {
            cnt_x1++
        } else {
            cnt_x1 = n + 1
            break
        }
    }
    for i := 0; i < n; i++ {
        if bottoms[i] == x { continue }
        if tops[i] == x {
            cnt_x2++
        } else {
            cnt_x2 = n + 1
            break
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    res1 := min(cnt_x1, cnt_x2)
    for i := 0; i < n; i++ {
        if tops[i] == y { continue }
        if bottoms[i] == y {
            cnt_y1++
        } else {
            cnt_y1 = n + 1
            break
        }
    }
    for i := 0; i < n; i++ {
        if bottoms[i] == y { continue }
        if tops[i] == y {
            cnt_y2 ++
        } else {
            cnt_y2 = n + 1
            break
        }
    }
    res2 := min(cnt_y1, cnt_y2)
    if res1 == n + 1 && res2 == n + 1 {
        return -1
    }
    return min(res1, res2)
}

func main() {
    // Example 1:
    // Input: tops = [2,1,2,4,2,2], bottoms = [5,2,6,2,3,2]
    // Output: 2
    // Explanation: 
    // The first figure represents the dominoes as given by tops and bottoms: before we do any rotations.
    // If we rotate the second and fourth dominoes, we can make every value in the top row equal to 2, as indicated by the second figure.
    fmt.Println(minDominoRotations([]int{2,1,2,4,2,2}, []int{5,2,6,2,3,2})) // 2
    // Example 2:
    // Input: tops = [3,5,1,2,3], bottoms = [3,6,3,3,4]
    // Output: -1
    // Explanation: 
    // In this case, it is not possible to rotate the dominoes to make one row of values equal.
    fmt.Println(minDominoRotations([]int{3,5,1,2,3}, []int{3,6,3,3,4})) // -1

    fmt.Println(minDominoRotations([]int{1,2,3,4,5,6}, []int{1,2,3,4,5,6})) // -1
    fmt.Println(minDominoRotations([]int{1,2,3,4,5,6}, []int{6,5,4,3,2,1})) // -1
    fmt.Println(minDominoRotations([]int{6,5,4,3,2,1}, []int{6,5,4,3,2,1})) // -1
    fmt.Println(minDominoRotations([]int{6,5,4,3,2,1}, []int{1,2,3,4,5,6})) // -1

    fmt.Println(minDominoRotations1([]int{2,1,2,4,2,2}, []int{5,2,6,2,3,2})) // 2
    fmt.Println(minDominoRotations1([]int{3,5,1,2,3}, []int{3,6,3,3,4})) // -1
    fmt.Println(minDominoRotations1([]int{1,2,3,4,5,6}, []int{1,2,3,4,5,6})) // -1
    fmt.Println(minDominoRotations1([]int{1,2,3,4,5,6}, []int{6,5,4,3,2,1})) // -1
    fmt.Println(minDominoRotations1([]int{6,5,4,3,2,1}, []int{6,5,4,3,2,1})) // -1
    fmt.Println(minDominoRotations1([]int{6,5,4,3,2,1}, []int{1,2,3,4,5,6})) // -1
}