package main

// 3394. Check if Grid can be Cut into Sections
// You are given an integer n representing the dimensions of an n x n grid, 
// with the origin at the bottom-left corner of the grid. 
// You are also given a 2D array of coordinates rectangles, 
// where rectangles[i] is in the form [startx, starty, endx, endy], representing a rectangle on the grid. 
// Each rectangle is defined as follows:
//     1. (startx, starty): The bottom-left corner of the rectangle.
//     2. (endx, endy): The top-right corner of the rectangle.

// Note that the rectangles do not overlap. 
// Your task is to determine if it is possible to make either two horizontal or two vertical cuts on the grid such that:
//     1. Each of the three resulting sections formed by the cuts contains at least one rectangle.
//     2. Every rectangle belongs to exactly one section.

// Return true if such cuts can be made; otherwise, return false.

// Example 1:
// Input: n = 5, rectangles = [[1,0,5,2],[0,2,2,4],[3,2,5,3],[0,4,4,5]]
// Output: true
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/10/23/tt1drawio.png" />
// The grid is shown in the diagram. We can make horizontal cuts at y = 2 and y = 4. Hence, output is true.

// Example 2:
// Input: n = 4, rectangles = [[0,0,1,1],[2,0,3,4],[0,2,2,3],[3,0,4,3]]
// Output: true
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/10/23/tc2drawio.png" />
// We can make vertical cuts at x = 2 and x = 3. Hence, output is true.

// Example 3:
// Input: n = 4, rectangles = [[0,2,2,4],[1,0,3,2],[2,2,3,4],[3,0,4,2],[3,2,4,4]]
// Output: false
// Explanation:
// We cannot make two horizontal or two vertical cuts that satisfy the conditions. Hence, output is false.

// Constraints:
//     3 <= n <= 10^9
//     3 <= rectangles.length <= 10^5
//     0 <= rectangles[i][0] < rectangles[i][2] <= n
//     0 <= rectangles[i][1] < rectangles[i][3] <= n
//     No two rectangles overlap.

import "fmt"
import "sort"
import "slices"

func checkValidCuts(n int, rectangles [][]int) bool {
    horizontalArea, verticalArea := make(map[float64]int), make(map[float64]int)
    for _, rect := range rectangles {
        sx, sy, ex, ey := float64(rect[0]), float64(rect[1]), float64(rect[2]), float64(rect[3])
        verticalArea[sx + 0.1]++
        verticalArea[ex]--
        horizontalArea[sy + 0.1]++
        horizontalArea[ey]--
    }
    helper := func(areas map[float64]int) bool {
        keys := make([]float64, 0, len(areas))
        for k := range areas {
            keys = append(keys, k)
        }
        sort.Float64s(keys)
        sortedAreas := make([]struct{ K float64; V int }, len(areas))
        for i, k := range keys {
            sortedAreas[i] = struct{ K float64; V int }{k, areas[k]}
        }
        areaAcc, prevAreaAcc, count := 0, 0, 0
        for _, v := range sortedAreas {
            prevAreaAcc = areaAcc
            areaAcc += v.V
            if prevAreaAcc != 0 && areaAcc == 0 {
                count++
            }
        }
        return count > 2
    }
    return helper(horizontalArea) || helper(verticalArea)
}

func checkValidCuts1(n int, rectangles [][]int) bool {
    type Pair struct{ l, r int }
    m := len(rectangles)
    arr1, arr2 := make([]Pair, m), make([]Pair, m)
    for i, rect := range rectangles {
        arr1[i], arr2[i] = Pair{ rect[0], rect[2] }, Pair{ rect[1], rect[3] }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    check := func(intervals []Pair) bool {
        // 按照左端点从小到大排序
        slices.SortFunc(intervals, func(a, b Pair) int { 
            return a.l - b.l 
        })
        count, mx := 0, 0
        for _, p := range intervals {
            if p.l >= mx { // 新区间
                count++
            }
            mx = max(mx, p.r) // 更新右端点最大值
        }
        return count >= 3
    }
    return check(arr1) || check(arr2)
}

func main() {
    // Example 1:
    // Input: n = 5, rectangles = [[1,0,5,2],[0,2,2,4],[3,2,5,3],[0,4,4,5]]
    // Output: true
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/10/23/tt1drawio.png" />
    // The grid is shown in the diagram. We can make horizontal cuts at y = 2 and y = 4. Hence, output is true.
    fmt.Println(checkValidCuts(5, [][]int{{1,0,5,2},{0,2,2,4},{3,2,5,3},{0,4,4,5}})) // true
    // Example 2:
    // Input: n = 4, rectangles = [[0,0,1,1],[2,0,3,4],[0,2,2,3],[3,0,4,3]]
    // Output: true
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/10/23/tc2drawio.png" />
    // We can make vertical cuts at x = 2 and x = 3. Hence, output is true.
    fmt.Println(checkValidCuts(4, [][]int{{0,0,1,1},{2,0,3,4},{0,2,2,3},{3,0,4,3}})) // true
    // Example 3:
    // Input: n = 4, rectangles = [[0,2,2,4],[1,0,3,2],[2,2,3,4],[3,0,4,2],[3,2,4,4]]
    // Output: false
    // Explanation:
    // We cannot make two horizontal or two vertical cuts that satisfy the conditions. Hence, output is false.
    fmt.Println(checkValidCuts(4, [][]int{{0,2,2,4},{1,0,3,2},{2,2,3,4},{3,0,4,2},{3,2,4,4}})) // false

    fmt.Println(checkValidCuts1(5, [][]int{{1,0,5,2},{0,2,2,4},{3,2,5,3},{0,4,4,5}})) // true
    fmt.Println(checkValidCuts1(4, [][]int{{0,0,1,1},{2,0,3,4},{0,2,2,3},{3,0,4,3}})) // true
    fmt.Println(checkValidCuts1(4, [][]int{{0,2,2,4},{1,0,3,2},{2,2,3,4},{3,0,4,2},{3,2,4,4}})) // false
}