package main

// 3288. Length of the Longest Increasing Path
// You are given a 2D array of integers coordinates of length n and an integer k, where 0 <= k < n.

// coordinates[i] = [xi, yi] indicates the point (xi, yi) in a 2D plane.

// An increasing path of length m is defined as a list of points (x1, y1), (x2, y2), (x3, y3), ..., (xm, ym) such that:
//     xi < xi + 1 and yi < yi + 1 for all i where 1 <= i < m.
//     (xi, yi) is in the given coordinates for all i where 1 <= i <= m.

// Return the maximum length of an increasing path that contains coordinates[k].

// Example 1:
// Input: coordinates = [[3,1],[2,2],[4,1],[0,0],[5,3]], k = 1
// Output: 3
// Explanation:
// (0, 0), (2, 2), (5, 3) is the longest increasing path that contains (2, 2).

// Example 2:
// Input: coordinates = [[2,1],[7,0],[5,6]], k = 2
// Output: 2
// Explanation:
// (2, 1), (5, 6) is the longest increasing path that contains (5, 6).

// Constraints:
//     1 <= n == coordinates.length <= 10^5
//     coordinates[i].length == 2
//     0 <= coordinates[i][0], coordinates[i][1] <= 10^9
//     All elements in coordinates are distinct.
//     0 <= k <= n - 1

import "fmt"
import "sort"

func maxPathLength(coordinates [][]int, k int) int {
    arr := coordinates[k]
    sort.Slice(coordinates, func(i, j int) bool {
        if coordinates[i][0] == coordinates[j][0] { return coordinates[i][1] > coordinates[j][1] }
        return coordinates[i][0] < coordinates[j][0]
    })

    nc := make([][]int, 0)
    for i := 0; i < len(coordinates); i++ {
        if coordinates[i][0] == arr[0] && coordinates[i][1] != arr[1] { continue }
        nc = append(nc, coordinates[i])
    }
    coordinates = nc
    res, cantChange := make([]int, 0), -1
    for i := 0; i < len(coordinates); i++ {
        index := sort.SearchInts(res, coordinates[i][1])
        if index == len(res) {
            res = append(res, coordinates[i][1])
        } else if coordinates[i][1] > cantChange {
            res[index] = coordinates[i][1]
        }
        if coordinates[i][0] == arr[0] && coordinates[i][1] == arr[1] {
            res = res[:index+1]
            cantChange = arr[1]
        } 
    }
    return len(res)
}

func maxPathLength1(coordinates [][]int, k int) int {
    x, y := coordinates[k][0], coordinates[k][1]
    sort.Slice(coordinates, func(i, j int) bool {
        if coordinates[i][0] == coordinates[j][0] { return coordinates[i][1] > coordinates[j][1] }
        return coordinates[i][0] < coordinates[j][0]
    })
    res := []int{}
    for _, v := range coordinates {
        if (v[0] > x && v[1] > y) || (v[0] < x && v[1] < y) {
            t := sort.SearchInts(res, v[1])
            if t == len(res) {
                res = append(res, v[1])
            } else {
                res[t] = v[1]
            }
        }
    }
    return len(res) + 1
}

func main() {
    // Example 1:
    // Input: coordinates = [[3,1],[2,2],[4,1],[0,0],[5,3]], k = 1
    // Output: 3
    // Explanation:
    // (0, 0), (2, 2), (5, 3) is the longest increasing path that contains (2, 2).
    fmt.Println(maxPathLength([][]int{{3,1},{2,2},{4,1},{0,0},{5,3}}, 1)) // 3
    // Example 2:
    // Input: coordinates = [[2,1],[7,0],[5,6]], k = 2
    // Output: 2
    // Explanation:
    // (2, 1), (5, 6) is the longest increasing path that contains (5, 6).
    fmt.Println(maxPathLength([][]int{{2,1},{7,0},{5,6}}, 2)) // 2

    fmt.Println(maxPathLength1([][]int{{3,1},{2,2},{4,1},{0,0},{5,3}}, 1)) // 3
    fmt.Println(maxPathLength1([][]int{{2,1},{7,0},{5,6}}, 2)) // 2
}