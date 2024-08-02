package main

// 3143. Maximum Points Inside the Square
// You are given a 2D array points and a string s where, points[i] represents the coordinates of point i, 
// and s[i] represents the tag of point i.

// A valid square is a square centered at the origin (0, 0), has edges parallel to the axes, 
// and does not contain two points with the same tag.

// Return the maximum number of points contained in a valid square.

// Note:
//     A point is considered to be inside the square if it lies on or within the square's boundaries.
//     The side length of the square can be zero.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2024/03/29/3708-tc1.png" />
// Input: points = [[2,2],[-1,-2],[-4,4],[-3,1],[3,-3]], s = "abdca"
// Output: 2
// Explanation:
// The square of side length 4 covers two points points[0] and points[1].

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2024/03/29/3708-tc2.png" />
// Input: points = [[1,1],[-2,-2],[-2,2]], s = "abb"
// Output: 1
// Explanation:
// The square of side length 2 covers one point, which is points[0].

// Example 3:
// Input: points = [[1,1],[-1,-1],[2,-2]], s = "ccd"
// Output: 0
// Explanation:
// It's impossible to make any valid squares centered at the origin such that it covers only one point among points[0] and points[1].

// Constraints:
//     1 <= s.length, points.length <= 10^5
//     points[i].length == 2
//     -10^9 <= points[i][0], points[i][1] <= 10^9
//     s.length == points.length
//     points consists of distinct coordinates.
//     s consists only of lowercase English letters.

import "fmt"
import "sort"
import "math/bits"

func maxPointsInsideSquare(points [][]int, s string) int {
    minLens, secondMin, count := make(map[rune]int), 1 << 32 - 1, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 0; i < len(points); i++ {
        n := max(abs(points[i][0]), abs(points[i][1]))
        c := rune(s[i])
        if _, ok := minLens[c]; !ok {
            minLens[c] = n
        } else if n < minLens[c] {
            if minLens[c] < secondMin {
                secondMin = minLens[c]
            }
            minLens[c] = n
        } else if n < secondMin {
            secondMin = n
        }
    }
    for _, v := range minLens {
        if v < secondMin {
            count++
        }
    }
    return count
}

func maxPointsInsideSquare1(points [][]int, s string) int {
    res := 0
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    sort.Search(1_000_000_001, func(size int) bool {
        vis := 0
        for i, p := range points {
            if abs(p[0]) <= size && abs(p[1]) <= size {
                c := s[i] - 'a'
                if vis >> c & 1 == 1 {
                    return true
                }
                vis |= 1 << c
            }
        }
        res = bits.OnesCount(uint(vis))
        return false
    })
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2024/03/29/3708-tc1.png" />
    // Input: points = [[2,2],[-1,-2],[-4,4],[-3,1],[3,-3]], s = "abdca"
    // Output: 2
    // Explanation:
    // The square of side length 4 covers two points points[0] and points[1].
    fmt.Println(maxPointsInsideSquare([][]int{{2,2},{-1,-2},{-4,4},{-3,1},{3,-3}}, "abdca")) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2024/03/29/3708-tc2.png" />
    // Input: points = [[1,1],[-2,-2],[-2,2]], s = "abb"
    // Output: 1
    // Explanation:
    // The square of side length 2 covers one point, which is points[0].
    fmt.Println(maxPointsInsideSquare([][]int{{1,1},{-2,-2},{-2,2}}, "abb")) // 1
    // Example 3:
    // Input: points = [[1,1],[-1,-1],[2,-2]], s = "ccd"
    // Output: 0
    // Explanation:
    // It's impossible to make any valid squares centered at the origin such that it covers only one point among points[0] and points[1].
    fmt.Println(maxPointsInsideSquare([][]int{{1,1},{-1,-1},{2,-2}}, "ccd")) // 0

    fmt.Println(maxPointsInsideSquare1([][]int{{2,2},{-1,-2},{-4,4},{-3,1},{3,-3}}, "abdca")) // 2
    fmt.Println(maxPointsInsideSquare1([][]int{{1,1},{-2,-2},{-2,2}}, "abb")) // 1
    fmt.Println(maxPointsInsideSquare1([][]int{{1,1},{-1,-1},{2,-2}}, "ccd")) // 0
}