package main 

// 1182. Shortest Distance to Target Color
// You are given an array colors, in which there are three colors: 1, 2 and 3.

// You are also given some queries. 
// Each query consists of two integers i and c, return the shortest distance between the given index i and the target color c.
// If there is no solution return -1.

// Example 1:
// Input: colors = [1,1,2,1,3,2,2,3,3], queries = [[1,3],[2,2],[6,1]]
// Output: [3,0,3]
// Explanation: 
// The nearest 3 from index 1 is at index 4 (3 steps away).
// The nearest 2 from index 2 is at index 2 itself (0 steps away).
// The nearest 1 from index 6 is at index 3 (3 steps away).

// Example 2:
// Input: colors = [1,2], queries = [[0,3]]
// Output: [-1]
// Explanation: There is no 3 in the array.

// Constraints:
//     1 <= colors.length <= 5*10^4
//     1 <= colors[i] <= 3
//     1 <= queries.length <= 5*10^4
//     queries[i].length == 2
//     0 <= queries[i][0] < colors.length
//     1 <= queries[i][1] <= 3

import "fmt"

func shortestDistanceColor(colors []int, queries [][]int) []int {
    // 先将颜色分组，分别保存颜色的索引，便于后续二分获取最近的目标颜色
    m := [3][]int{}  // 使用数组代替哈希
    for i,n := range colors {
        m[n-1] = append(m[n-1], i)
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    res := make([]int, len(queries))
    for i,item := range queries{
        pos, target := item[0],item[1] - 1 // 目标位置, 目标颜色
        if colors[pos] - 1 == target { // 目标颜色和目标位置的颜色相同
            res[i] = 0
            continue
        }
        if len(m[target]) == 0 { // 目标颜色不存在
            res[i] = -1
            continue
        }
        res[i] = len(colors)
        l, r := 0, len(m[target]) - 1 // 二分获取素插入位置
        for l < r {
            mid := (l + r) >> 1
            if m[target][mid] < pos {
                l = mid + 1
            } else {
                r = mid
            }
        }
        if l != len(m[target]) {
            res[i] = min(res[i], abs(m[target][l] - pos))
        }
        if l != 0 {
            res[i] = min(res[i], abs(pos - m[target][l-1]))
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: colors = [1,1,2,1,3,2,2,3,3], queries = [[1,3],[2,2],[6,1]]
    // Output: [3,0,3]
    // Explanation: 
    // The nearest 3 from index 1 is at index 4 (3 steps away).
    // The nearest 2 from index 2 is at index 2 itself (0 steps away).
    // The nearest 1 from index 6 is at index 3 (3 steps away).
    fmt.Println(shortestDistanceColor([]int{1,1,2,1,3,2,2,3,3},[][]int{{1,3},{2,2},{6,1}})) // [3,0,3]
    // Example 2:
    // Input: colors = [1,2], queries = [[0,3]]
    // Output: [-1]
    // Explanation: There is no 3 in the array.
    fmt.Println(shortestDistanceColor([]int{1,2},[][]int{{0,3}})) // [-1]
}