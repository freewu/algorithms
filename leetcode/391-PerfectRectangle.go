package main

// 391. Perfect Rectangle
// Given an array rectangles where rectangles[i] = [xi, yi, ai, bi] represents an axis-aligned rectangle. 
// The bottom-left point of the rectangle is (xi, yi) and the top-right point of it is (ai, bi).

// Return true if all the rectangles together form an exact cover of a rectangular region.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/03/27/perectrec1-plane.jpg" />
// Input: rectangles = [[1,1,3,3],[3,1,4,2],[3,2,4,4],[1,3,2,4],[2,3,3,4]]
// Output: true
// Explanation: All 5 rectangles together form an exact cover of a rectangular region.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/03/27/perfectrec2-plane.jpg" />
// Input: rectangles = [[1,1,2,3],[1,3,2,4],[3,1,4,2],[3,2,4,4]]
// Output: false
// Explanation: Because there is a gap between the two rectangular regions.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/03/27/perfecrrec4-plane.jpg" />
// Input: rectangles = [[1,1,3,3],[3,1,4,2],[1,3,2,4],[2,2,4,4]]
// Output: false
// Explanation: Because two of the rectangles overlap with each other.

// Constraints:
//     1 <= rectangles.length <= 2 * 10^4
//     rectangles[i].length == 4
//     -10^5 <= xi, yi, ai, bi <= 10^5

import "fmt"

func isRectangleCover(rectangles [][]int) bool {
    type Pair struct {
        x, y int
    }
    count, m := 0, make(map[Pair]int)
    for i := range(rectangles) {
        x1, y1 := rectangles[i][0], rectangles[i][1]
        x2, y2 := rectangles[i][2], rectangles[i][3]
        m[Pair{x: x1, y: y1}]--
        m[Pair{x: x1, y: y2}]++
        m[Pair{x: x2, y: y1}]++
        m[Pair{x: x2, y: y2}]--
    }
    for _, v := range(m) {
        if (v != 0) {
            if (v < 0) {
                count+= (0 - v)
            } else {
                count+= (v - 0)
            }
        }
    }
    return count == 4 // 最后只剩下4个点即为完美的长方型
}

func isRectangleCover1(rectangles [][]int) bool {
    type point struct{ x, y int }
    area, minX, minY, maxX, maxY := 0, rectangles[0][0], rectangles[0][1], rectangles[0][2], rectangles[0][3]
    cnt := map[point]int{}
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, rect := range rectangles {
        x, y, a, b := rect[0], rect[1], rect[2], rect[3]
        area += (a - x) * (b - y)
        minX = min(minX, x)
        minY = min(minY, y)
        maxX = max(maxX, a)
        maxY = max(maxY, b)
        cnt[point{x, y}]++
        cnt[point{x, b}]++
        cnt[point{a, y}]++
        cnt[point{a, b}]++
    }
    if area != (maxX - minX) * (maxY-minY) || 
       cnt[point{minX, minY}] != 1 || cnt[point{minX, maxY}] != 1 || 
       cnt[point{maxX, minY}] != 1 || cnt[point{maxX, maxY}] != 1 {
        return false
    }
    delete(cnt, point{minX, minY})
    delete(cnt, point{minX, maxY})
    delete(cnt, point{maxX, minY})
    delete(cnt, point{maxX, maxY})
    for _, c := range cnt {
        if c != 2 && c != 4 {
            return false
        }
    }
    return true
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/03/27/perectrec1-plane.jpg" />
    // Input: rectangles = [[1,1,3,3],[3,1,4,2],[3,2,4,4],[1,3,2,4],[2,3,3,4]]
    // Output: true
    // Explanation: All 5 rectangles together form an exact cover of a rectangular region.
    fmt.Println(isRectangleCover([][]int{{1,1,3,3},{3,1,4,2},{3,2,4,4},{1,3,2,4},{2,3,3,4}})) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/03/27/perfectrec2-plane.jpg" />
    // Input: rectangles = [[1,1,2,3],[1,3,2,4],[3,1,4,2],[3,2,4,4]]
    // Output: false
    // Explanation: Because there is a gap between the two rectangular regions.
    fmt.Println(isRectangleCover([][]int{{1,1,2,3},{1,3,2,4},{3,1,4,2},{3,2,4,4}})) // false
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/03/27/perfecrrec4-plane.jpg" />
    // Input: rectangles = [[1,1,3,3],[3,1,4,2],[1,3,2,4],[2,2,4,4]]
    // Output: false
    // Explanation: Because two of the rectangles overlap with each other.
    fmt.Println(isRectangleCover([][]int{{1,1,3,3},{3,1,4,2},{1,3,2,4},{2,2,4,4}})) // false

    fmt.Println(isRectangleCover1([][]int{{1,1,3,3},{3,1,4,2},{3,2,4,4},{1,3,2,4},{2,3,3,4}})) // true
    fmt.Println(isRectangleCover1([][]int{{1,1,2,3},{1,3,2,4},{3,1,4,2},{3,2,4,4}})) // false
    fmt.Println(isRectangleCover1([][]int{{1,1,3,3},{3,1,4,2},{1,3,2,4},{2,2,4,4}})) // false
}