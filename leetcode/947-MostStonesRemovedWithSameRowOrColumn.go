package main

// 947. Most Stones Removed with Same Row or Column
// On a 2D plane, we place n stones at some integer coordinate points. 
// Each coordinate point may have at most one stone.

// A stone can be removed if it shares either the same row or the same column as another stone that has not been removed.

// Given an array stones of length n where stones[i] = [xi, yi] represents the location of the ith stone, 
// return the largest possible number of stones that can be removed.

// Example 1:
// Input: stones = [[0,0],[0,1],[1,0],[1,2],[2,1],[2,2]]
// Output: 5
// Explanation: One way to remove 5 stones is as follows:
// 1. Remove stone [2,2] because it shares the same row as [2,1].
// 2. Remove stone [2,1] because it shares the same column as [0,1].
// 3. Remove stone [1,2] because it shares the same row as [1,0].
// 4. Remove stone [1,0] because it shares the same column as [0,0].
// 5. Remove stone [0,1] because it shares the same row as [0,0].
// Stone [0,0] cannot be removed since it does not share a row/column with another stone still on the plane.

// Example 2:
// Input: stones = [[0,0],[0,2],[1,1],[2,0],[2,2]]
// Output: 3
// Explanation: One way to make 3 moves is as follows:
// 1. Remove stone [2,2] because it shares the same row as [2,0].
// 2. Remove stone [2,0] because it shares the same column as [0,0].
// 3. Remove stone [0,2] because it shares the same row as [0,0].
// Stones [0,0] and [1,1] cannot be removed since they do not share a row/column with another stone still on the plane.

// Example 3:
// Input: stones = [[0,0]]
// Output: 0
// Explanation: [0,0] is the only stone on the plane, so you cannot remove it.

// Constraints:
//     1 <= stones.length <= 1000
//     0 <= xi, yi <= 10^4
//     No two stones are at the same coordinate point.

import "fmt"

// dfs
func removeStones(stones [][]int) int {
    rows, cols := make(map[int][]int), make(map[int][]int)
    for _, stone := range stones {
        rows[stone[0]] = append(rows[stone[0]], stone[1])
        cols[stone[1]] = append(cols[stone[1]], stone[0])
    }
    count, visited := 0, make(map[[2]int]bool)
    var dfs func(x, y int)
    dfs = func(x, y int) {
        visited[[2]int{x, y}] = true
        for _, col := range rows[x] {
            if !visited[[2]int{x, col}] { dfs(x, col) }
        }
        for _, row := range cols[y] {
            if !visited[[2]int{row, y}] { dfs(row, y) }
        }
    }
    for _, stone := range stones {
        if _, ok := visited[[2]int{stone[0], stone[1]}]; !ok {
            count ++
            dfs(stone[0], stone[1])
        } 
    }
    return len(stones) - count
}

// 并查集
func removeStones1(stones [][]int) int {
    // 记录每一行和每一列的第一块石头
    // 如果遍历到的石头是一行且一列的第一块石头 不用消除
    // 否则合并（随便和行合并还是列合并）
    cols, rows := make(map[int]int), make(map[int]int)
    n := len(stones)
    parent := make([]int, n)
    count := n
    for i := 0; i < n; i++ {
        parent[i] = i
    }
    var find func(x int) int
    find = func(x int) int {
        if parent[x] != x {
            parent[x] = find(parent[x])
        }
        return parent[x]
    }
    union := func(x, y int) {
        if find(x) != find(y) {
            parent[find(x)] = find(y)
            count--
        }
    }
    for i, stone := range stones {
        r, c := stone[0], stone[1]
        if _, ok := rows[r]; !ok { // 如果是一行的第一块石头
            rows[r] = i
        } else {
            union(rows[r], i)
        }
        if _, ok := cols[c]; !ok { // 如果是一列的第一块石头
            cols[c] = i
        } else {
            union(cols[c], i)
        }
    }
    return n - count
}

func main() {
    // Example 1:
    // Input: stones = [[0,0],[0,1],[1,0],[1,2],[2,1],[2,2]]
    // Output: 5
    // Explanation: One way to remove 5 stones is as follows:
    // 1. Remove stone [2,2] because it shares the same row as [2,1].
    // 2. Remove stone [2,1] because it shares the same column as [0,1].
    // 3. Remove stone [1,2] because it shares the same row as [1,0].
    // 4. Remove stone [1,0] because it shares the same column as [0,0].
    // 5. Remove stone [0,1] because it shares the same row as [0,0].
    // Stone [0,0] cannot be removed since it does not share a row/column with another stone still on the plane.
    fmt.Println(removeStones([][]int{{0,0},{0,1},{1,0},{1,2},{2,1},{2,2}})) // 5
    // Example 2:
    // Input: stones = [[0,0],[0,2],[1,1],[2,0],[2,2]]
    // Output: 3
    // Explanation: One way to make 3 moves is as follows:
    // 1. Remove stone [2,2] because it shares the same row as [2,0].
    // 2. Remove stone [2,0] because it shares the same column as [0,0].
    // 3. Remove stone [0,2] because it shares the same row as [0,0].
    // Stones [0,0] and [1,1] cannot be removed since they do not share a row/column with another stone still on the plane.
    fmt.Println(removeStones([][]int{{0,0},{0,2},{1,1},{2,0},{2,2}})) // 3
    // Example 3:
    // Input: stones = [[0,0]]
    // Output: 0
    // Explanation: [0,0] is the only stone on the plane, so you cannot remove it.
    fmt.Println(removeStones([][]int{{0,0}})) // 0

    fmt.Println(removeStones1([][]int{{0,0},{0,1},{1,0},{1,2},{2,1},{2,2}})) // 5
    fmt.Println(removeStones1([][]int{{0,0},{0,2},{1,1},{2,0},{2,2}})) // 3
    fmt.Println(removeStones1([][]int{{0,0}})) // 0
}