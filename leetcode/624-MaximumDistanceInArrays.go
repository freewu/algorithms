package main

// 624. Maximum Distance in Arrays
// You are given m arrays, where each array is sorted in ascending order.
// You can pick up two integers from two different arrays (each array picks one) and calculate the distance. 
// We define the distance between two integers a and b to be their absolute difference |a - b|.
// Return the maximum distance.

// Example 1:
// Input: arrays = [[1,2,3],[4,5],[1,2,3]]
// Output: 4
// Explanation: One way to reach the maximum distance 4 is to pick 1 in the first or third array and pick 5 in the second array.

// Example 2:
// Input: arrays = [[1],[1]]
// Output: 0

// Constraints:
//     m == arrays.length
//     2 <= m <= 10^5
//     1 <= arrays[i].length <= 500
//     -10^4 <= arrays[i][j] <= 10^4
//     arrays[i] is sorted in ascending order.
//     There will be at most 10^5 integers in all the arrays.

import "fmt"

func maxDistance(arrays [][]int) int {
    if len(arrays) <= 0 {
        return 0
    }
    res, mx, mn := 0, arrays[0][len(arrays[0]) - 1], arrays[0][0]
    max := func(x, y int) int { if x > y { return x; }; return y; }
    min := func(x, y int) int { if x < y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i:=1; i<len(arrays); i++ {
        n := len(arrays[i])
        // 因为数据是有序的，每次比较的时候只用比较: 数组里面的第一个和最后一个值
        res = max(res, max(abs(arrays[i][n-1] - mn), abs(mx - arrays[i][0])))
        mx = max(mx, arrays[i][n-1])
        mn = min(mn, arrays[i][0])
    }
    return res
}

func main() {
    // Example 1:
    // Input: arrays = [[1,2,3],[4,5],[1,2,3]]
    // Output: 4
    // Explanation: One way to reach the maximum distance 4 is to pick 1 in the first or third array and pick 5 in the second array.
    fmt.Println(maxDistance([][]int{{1,2,3},{4,5},{1,2,3}})) // 4
    // Example 2:
    // Input: arrays = [[1],[1]]
    // Output: 0
    fmt.Println(maxDistance([][]int{{1},{1}})) // 0
}