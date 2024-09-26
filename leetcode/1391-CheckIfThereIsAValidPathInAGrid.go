package main

// 1391. Check if There is a Valid Path in a Grid
// You are given an m x n grid. Each cell of grid represents a street. 
// The street of grid[i][j] can be:
//     1 which means a street connecting the left cell and the right cell.
//     2 which means a street connecting the upper cell and the lower cell.
//     3 which means a street connecting the left cell and the lower cell.
//     4 which means a street connecting the right cell and the lower cell.
//     5 which means a street connecting the left cell and the upper cell.
//     6 which means a street connecting the right cell and the upper cell.

// <img src="https://assets.leetcode.com/uploads/2020/03/05/main.png" />

// You will initially start at the street of the upper-left cell (0, 0). 
// A valid path in the grid is a path that starts from the upper left cell (0, 0) and ends at the bottom-right cell (m - 1, n - 1). The path should only follow the streets.

// Notice that you are not allowed to change any street.

// Return true if there is a valid path in the grid or false otherwise.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/03/05/e1.png" />
// Input: grid = [[2,4,3],[6,5,2]]
// Output: true
// Explanation: As shown you can start at cell (0, 0) and visit all the cells of the grid to reach (m - 1, n - 1).

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/03/05/e2.png" />
// Input: grid = [[1,2,1],[1,2,1]]
// Output: false
// Explanation: As shown you the street at cell (0, 0) is not connected with any street of any other cell and you will get stuck at cell (0, 0)

// Example 3:
// Input: grid = [[1,1,2]]
// Output: false
// Explanation: You will get stuck at cell (0, 1) and you cannot reach cell (0, 2).

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 300
//     1 <= grid[i][j] <= 6

import "fmt"


func hasValidPath(grid [][]int) bool {
    type cell struct {
        x int
        y int
    }
    directions := map[int][][]int{
        1: [][]int{{0,-1}, {0,1}},
        2: [][]int{{-1,0}, {1,0}},
        3: [][]int{{0,-1}, {1,0}},
        4: [][]int{{0,1}, {1,0}},
        5: [][]int{{0,-1}, {-1,0}},
        6: [][]int{{0,1}, {-1,0}},
    }
    accept := map[cell]map[int]bool{
        cell{x:0, y:1}: map[int]bool{1:true,3:true, 5:true},
        cell{x:1, y:0}: map[int]bool{2:true,5:true, 6:true},
        cell{x:0, y:-1}: map[int]bool{4:true,1:true, 6:true},
        cell{x:-1, y:0}: map[int]bool{2:true,4:true, 3:true},
    }
    res, visited := false, make(map[cell]bool)
    var dfs func(i, j int)
    dfs = func(i, j int){
        if i == len(grid)-1 && j == len(grid[0])-1 {
            res = true
        }
        if !res {
            ds := directions[grid[i][j]]
            for _, d := range ds {
                nx, ny :=i + d[0], j + d[1]
                if nx >= 0 && ny >= 0 && nx < len(grid) && ny <len(grid[0]) {
                    if !visited[cell{x:nx, y:ny}] && accept[cell{x:d[0],y:d[1]}][grid[nx][ny]] {
                        visited[cell{x:nx, y:ny}] = true
                        dfs(nx,ny)
                    }
                }
            }
        }
    }
    visited[cell{x:0,y:0}] = true
    dfs(0,0)
    return res
}

// union find
func hasValidPath1(grid [][]int) bool {
    n, m := len(grid), len(grid[0])
    k := n * m
    u := newUnionFind(k)
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            x := u.find(i*m + j)
            switch grid[i][j] {
            case 1:
                if j > 0 && (grid[i][j-1] == 1 || grid[i][j-1] == 6 || grid[i][j-1] == 4) {
                    y := u.find(i*m + j - 1)
                    u.union(x, y)
                }
                if j < m-1 && (grid[i][j+1] == 1 || grid[i][j+1] == 3 || grid[i][j+1] == 5) {
                    y := u.find(i*m + j + 1)
                    u.union(x, y)
                }
            case 2:
                if i > 0 && (grid[i-1][j] == 2 || grid[i-1][j] == 3 || grid[i-1][j] == 4) {
                    y := u.find((i-1)*m + j)
                    u.union(x, y)
                }
                if i < n-1 && (grid[i+1][j] == 2 || grid[i+1][j] == 5 || grid[i+1][j] == 6) {
                    y := u.find((i+1)*m + j)
                    u.union(x, y)
                }
            case 3:
                if j > 0 && (grid[i][j-1] == 1 || grid[i][j-1] == 4 || grid[i][j-1] == 6) {
                    y := u.find(i*m + j - 1)
                    u.union(x, y)
                }
                if i < n-1 && (grid[i+1][j] == 2 || grid[i+1][j] == 5 || grid[i+1][j] == 6) {
                    y := u.find((i+1)*m + j)
                    u.union(x, y)
                }
            case 4:
                if j < m-1 && (grid[i][j+1] == 1 || grid[i][j+1] == 3 || grid[i][j+1] == 5) {
                    y := u.find(i*m + j + 1)
                    u.union(x, y)
                }
                if i < n-1 && (grid[i+1][j] == 2 || grid[i+1][j] == 5 || grid[i+1][j] == 6) {
                    y := u.find((i+1)*m + j)
                    u.union(x, y)
                }
            case 5:
                if j > 0 && (grid[i][j-1] == 1 || grid[i][j-1] == 4 || grid[i][j-1] == 6) {
                    y := u.find(i*m + j - 1)
                    u.union(x, y)
                }
                if i > 0 && (grid[i-1][j] == 2 || grid[i-1][j] == 3 || grid[i-1][j] == 4) {
                    y := u.find((i-1)*m + j)
                    u.union(x, y)
                }
            case 6:
                if i > 0 && (grid[i-1][j] == 2 || grid[i-1][j] == 3 || grid[i-1][j] == 4) {
                    y := u.find((i-1)*m + j)
                    u.union(x, y)
                }
                if j < m-1 && (grid[i][j+1] == 1 || grid[i][j+1] == 3 || grid[i][j+1] == 5) {
                    y := u.find(i*m + j + 1)
                    u.union(x, y)
                }
            }
        }
    }
    return u.find(0) == u.find(k-1)
}


type UnionFind struct {
    fa    []int
    rank  []int
    count int
}

func newUnionFind(n int) *UnionFind {
    u := &UnionFind{
        fa:    make([]int, n),
        rank:  make([]int, n),
        count: n,
    }
    for i := 0; i < n; i++ {
        u.fa[i] = i
    }
    return u
}

func (u *UnionFind) find(x int) int {
    if u.fa[x] == x { return x }
    u.fa[x] = u.find(u.fa[x])
    return u.fa[x]
}

func (u *UnionFind) union(x, y int) {
    if x == y { return }
    if u.rank[x] <= u.rank[y] {
        u.fa[x] = y
    } else {
        u.fa[y] = x
    }
    if u.rank[x] == u.rank[y] {
        u.rank[x]++
    }
    u.count--
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/03/05/e1.png" />
    // Input: grid = [[2,4,3],[6,5,2]]
    // Output: true
    // Explanation: As shown you can start at cell (0, 0) and visit all the cells of the grid to reach (m - 1, n - 1).
    fmt.Println(hasValidPath([][]int{{2,4,3},{6,5,2}})) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/03/05/e2.png" />
    // Input: grid = [[1,2,1],[1,2,1]]
    // Output: false
    // Explanation: As shown you the street at cell (0, 0) is not connected with any street of any other cell and you will get stuck at cell (0, 0)
    fmt.Println(hasValidPath([][]int{{1,2,1},{1,2,1}})) // false
    // Example 3:
    // Input: grid = [[1,1,2]]
    // Output: false
    // Explanation: You will get stuck at cell (0, 1) and you cannot reach cell (0, 2).
    fmt.Println(hasValidPath([][]int{{1,1,2}})) // false

    fmt.Println(hasValidPath([][]int{{1,1,3}})) // true
    fmt.Println(hasValidPath([][]int{{1,1,1}})) // true

    fmt.Println(hasValidPath1([][]int{{2,4,3},{6,5,2}})) // true
    fmt.Println(hasValidPath1([][]int{{1,2,1},{1,2,1}})) // false
    fmt.Println(hasValidPath1([][]int{{1,1,2}})) // false
    fmt.Println(hasValidPath1([][]int{{1,1,3}})) // true
    fmt.Println(hasValidPath1([][]int{{1,1,1}})) // true
}