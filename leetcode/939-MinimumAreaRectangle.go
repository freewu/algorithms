package main

// 939. Minimum Area Rectangle
// You are given an array of points in the X-Y plane points where points[i] = [xi, yi].

// Return the minimum area of a rectangle formed from these points, with sides parallel to the X and Y axes. 
// If there is not any such rectangle, return 0.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/08/03/rec1.JPG" />
// Input: points = [[1,1],[1,3],[3,1],[3,3],[2,2]]
// Output: 4

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/08/03/rec2.JPG" />
// Input: points = [[1,1],[1,3],[3,1],[3,3],[4,1],[4,3]]
// Output: 2

// Constraints:
//     1 <= points.length <= 500
//     points[i].length == 2
//     0 <= xi, yi <= 4 * 10^4
//     All the given points are unique.

import "fmt"

func minAreaRect(points [][]int) int {
    mp, res := map[[2]int]bool{}, 1 << 31
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for _, point := range points {
        corner1 := [2]int{point[0], point[1]}
        for corner2 := range mp {
            corner3 := [2]int{corner1[0], corner2[1]}
            if !mp[corner3] {
                continue
            }
            corner4 := [2]int{corner2[0], corner1[1]}
            if !mp[corner4] {
                continue
            }
            res = min(res, abs(corner1[0] - corner2[0]) * abs(corner1[1] - corner2[1]) )
        }
        mp[corner1] = true
    }
    if res == 1 << 31 {
        return 0
    }
    return res
}

func minAreaRect1(points [][]int) int {
    if len(points) < 4 { return 0 }
    contains := func (mp map[int]struct{}, key int) bool { _, ok := mp[key]; return ok; }
    genKey := func (x, y int) int { return 40001*x + y }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    res, mp := 1 << 31, make(map[int]struct{})
    for _, point := range points {
        mp[genKey(point[0], point[1])] = struct{}{}
    }
    for i := 0; i < len(points); i++ {
        for j := i + 1; j < len(points); j++ {
            if points[i][0] != points[j][0] && points[i][1] != points[j][1] { // 不是同一个点
                if contains(mp, genKey(points[i][0], points[j][1])) && contains(mp, genKey(points[j][0], points[i][1])) {
                    res = min(res, abs(points[i][0] - points[j][0]) * abs(points[i][1] - points[j][1]))
                }
            }
        }
    }
    if res == 1 << 31 {
        return 0
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/08/03/rec1.JPG" />
    // Input: points = [[1,1],[1,3],[3,1],[3,3],[2,2]]
    // Output: 4
    fmt.Println(minAreaRect([][]int{{1,1},{1,3},{3,1},{3,3},{2,2}})) // 4
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/08/03/rec2.JPG" />
    // Input: points = [[1,1],[1,3],[3,1],[3,3],[4,1],[4,3]]
    // Output: 2
    fmt.Println(minAreaRect([][]int{{1,1},{1,3},{3,1},{3,3},{4,1},{4,3}})) // 2

    fmt.Println(minAreaRect1([][]int{{1,1},{1,3},{3,1},{3,3},{2,2}})) // 4
    fmt.Println(minAreaRect1([][]int{{1,1},{1,3},{3,1},{3,3},{4,1},{4,3}})) // 2
}