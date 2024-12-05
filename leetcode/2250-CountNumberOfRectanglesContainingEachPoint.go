package main

// 2250. Count Number of Rectangles Containing Each Point
// You are given a 2D integer array rectangles where rectangles[i] = [li, hi] indicates 
// that ith rectangle has a length of li and a height of hi. 
// You are also given a 2D integer array points where points[j] = [xj, yj] is a point with coordinates (xj, yj).

// The ith rectangle has its bottom-left corner point at the coordinates (0, 0) 
// and its top-right corner point at (li, hi).

// Return an integer array count of length points.length 
// where count[j] is the number of rectangles that contain the jth point.

// The ith rectangle contains the jth point if 0 <= xj <= li and 0 <= yj <= hi. 
// Note that points that lie on the edges of a rectangle are also considered to be contained by that rectangle.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/03/02/example1.png" />
// Input: rectangles = [[1,2],[2,3],[2,5]], points = [[2,1],[1,4]]
// Output: [2,1]
// Explanation: 
// The first rectangle contains no points.
// The second rectangle contains only the point (2, 1).
// The third rectangle contains the points (2, 1) and (1, 4).
// The number of rectangles that contain the point (2, 1) is 2.
// The number of rectangles that contain the point (1, 4) is 1.
// Therefore, we return [2, 1].

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/03/02/example2.png" />
// Input: rectangles = [[1,1],[2,2],[3,3]], points = [[1,3],[1,1]]
// Output: [1,3]
// Explanation:
// The first rectangle contains only the point (1, 1).
// The second rectangle contains only the point (1, 1).
// The third rectangle contains the points (1, 3) and (1, 1).
// The number of rectangles that contain the point (1, 3) is 1.
// The number of rectangles that contain the point (1, 1) is 3.
// Therefore, we return [1, 3].

// Constraints:
//     1 <= rectangles.length, points.length <= 5 * 10^4
//     rectangles[i].length == points[j].length == 2
//     1 <= li, xj <= 10^9
//     1 <= hi, yj <= 100
//     All the rectangles are unique.
//     All the points are unique.

import "fmt"
import "sort"
import "slices"

func countRectangles(rectangles [][]int, points [][]int) []int {
    n := 101
    res, d := []int{}, make([][]int, n)
    for _, r := range rectangles {
        d[r[1]] = append(d[r[1]], r[0])
    }
    for _, row := range d {
        sort.Ints(row)
    }
    for _, p := range points {
        count, x, y := 0, p[0], p[1]
        for h := y; h < n; h++ {
            xs := d[h]
            left, right := 0, len(xs)
            for left < right {
                mid := (left + right) >> 1
                if xs[mid] >= x {
                    right = mid
                } else {
                    left = mid + 1
                }
            }
            count += len(xs) - left
        }
        res = append(res, count)
    }
    return res
}

func countRectangles1(rectangles [][]int, points [][]int) []int {
    d := make(map[int][]int)
    for _, rectangle := range rectangles {
        d[rectangle[1]] = append(d[rectangle[1]], rectangle[0])
    }
    for _, row := range d {
        slices.Sort(row)
    }
    res := make([]int, len(points))
    for i, point := range points {
        for j := point[1]; j <= 100; j++ {
            index, _ := slices.BinarySearch(d[j], point[0])
            res[i] += len(d[j]) - index
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/03/02/example1.png" />
    // Input: rectangles = [[1,2],[2,3],[2,5]], points = [[2,1],[1,4]]
    // Output: [2,1]
    // Explanation: 
    // The first rectangle contains no points.
    // The second rectangle contains only the point (2, 1).
    // The third rectangle contains the points (2, 1) and (1, 4).
    // The number of rectangles that contain the point (2, 1) is 2.
    // The number of rectangles that contain the point (1, 4) is 1.
    // Therefore, we return [2, 1].
    fmt.Println(countRectangles([][]int{{1,2},{2,3},{2,5}}, [][]int{{2,1},{1,4}})) // [2,1]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/03/02/example2.png" />
    // Input: rectangles = [[1,1],[2,2],[3,3]], points = [[1,3],[1,1]]
    // Output: [1,3]
    // Explanation:
    // The first rectangle contains only the point (1, 1).
    // The second rectangle contains only the point (1, 1).
    // The third rectangle contains the points (1, 3) and (1, 1).
    // The number of rectangles that contain the point (1, 3) is 1.
    // The number of rectangles that contain the point (1, 1) is 3.
    // Therefore, we return [1, 3].
    fmt.Println(countRectangles([][]int{{1,1},{2,2},{3,3}}, [][]int{{1,3},{1,1}})) // [1,3]

    fmt.Println(countRectangles1([][]int{{1,2},{2,3},{2,5}}, [][]int{{2,1},{1,4}})) // [2,1]
    fmt.Println(countRectangles1([][]int{{1,1},{2,2},{3,3}}, [][]int{{1,3},{1,1}})) // [1,3]
}