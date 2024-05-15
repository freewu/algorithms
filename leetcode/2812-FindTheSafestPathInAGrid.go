package main

// 2812. Find the Safest Path in a Grid
// You are given a 0-indexed 2D matrix grid of size n x n, where (r, c) represents:
//     A cell containing a thief if grid[r][c] = 1
//     An empty cell if grid[r][c] = 0

// You are initially positioned at cell (0, 0). In one move, you can move to any adjacent cell in the grid, including cells containing thieves.
// The safeness factor of a path on the grid is defined as the minimum manhattan distance from any cell in the path to any thief in the grid.
// Return the maximum safeness factor of all paths leading to cell (n - 1, n - 1).
// An adjacent cell of cell (r, c), is one of the cells (r, c + 1), (r, c - 1), (r + 1, c) and (r - 1, c) if it exists.
// The Manhattan distance between two cells (a, b) and (x, y) is equal to |a - x| + |b - y|, where |val| denotes the absolute value of val.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/07/02/example1.png" />
// Input: grid = [[1,0,0],[0,0,0],[0,0,1]]
// Output: 0
// Explanation: All paths from (0, 0) to (n - 1, n - 1) go through the thieves in cells (0, 0) and (n - 1, n - 1).

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/07/02/example2.png" />
// Input: grid = [[0,0,1],[0,0,0],[0,0,0]]
// Output: 2
// Explanation: The path depicted in the picture above has a safeness factor of 2 since:
// - The closest cell of the path to the thief at cell (0, 2) is cell (0, 0). The distance between them is | 0 - 0 | + | 0 - 2 | = 2.
// It can be shown that there are no other paths with a higher safeness factor.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2023/07/02/example3.png" />
// Input: grid = [[0,0,0,1],[0,0,0,0],[0,0,0,0],[1,0,0,0]]
// Output: 2
// Explanation: The path depicted in the picture above has a safeness factor of 2 since:
// - The closest cell of the path to the thief at cell (0, 3) is cell (1, 2). The distance between them is | 0 - 1 | + | 3 - 2 | = 2.
// - The closest cell of the path to the thief at cell (3, 0) is cell (3, 2). The distance between them is | 3 - 3 | + | 0 - 2 | = 2.
// It can be shown that there are no other paths with a higher safeness factor.

// Constraints:
//     1 <= grid.length == n <= 400
//     grid[i].length == n
//     grid[i][j] is either 0 or 1.
//     There is at least one thief in the grid.

import "fmt"

// dfs
func maximumSafenessFactor(grid [][]int) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    type pair struct { i, j int }
    n, dirs := len(grid), []pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    dis := make([][]int, n)
    for i := range dis {
        dis[i] = make([]int, n)
        for j := range dis[i] {
            dis[i][j] = -1
        }
    }
    thiefs := []pair{}
    for i, row := range grid {
        for j, v := range row {
            if v == 1 {
                thiefs = append(thiefs, pair{i, j})
                dis[i][j] = 0
            }
        }
    }
    q := thiefs
    for i := 1; len(q) > 0; i++ {
        tmp := q
        q = []pair{}
        for _, v := range tmp {
            for _, d := range dirs {
                ni, nj := v.i+d.i, v.j+d.j
                if ni >= 0 && ni < n && nj >= 0 && nj < n && dis[ni][nj] < 0 {
                    dis[ni][nj] = i
                    q = append(q, pair{ni, nj})
                }
            }
        }
    }
    res, visited := 0, make([]int, n * n)
    var dfs func(i, j, mi int)
    dfs = func(i, j, mi int) {
        if i == n-1 && j == n-1 {
            res = max(res, mi)
            return
        }
        ha:= i * n + j
        if mi <= res || mi <= visited[ha] {
            return
        }
        visited[ha] = max(visited[ha], mi)
        for _, d := range dirs {
            if ni, nj := i+d.i, j+d.j; ni >= 0 && ni < n && nj >= 0 && nj < n {
                dfs(ni, nj, min(mi, dis[ni][nj]))
            }
        }
    }
    dfs(0, 0, min(dis[0][0], dis[n-1][n-1]))
    return res
}

// 
func maximumSafenessFactor1(grid [][]int) int {
    n := len(grid)
    if grid[0][0]==1 || grid[n-1][n-1]==1 {
        return 0
    }
    dis := make([][]int, n)
    // 多源BFS
    type pair struct {x,y int}
    q:=[]pair{}
    for i,row:=range grid {
        if dis[i]==nil {
            dis[i]=make([]int, n)
        }
        for j,num:=range row {
            if num==0 {
                dis[i][j]=-1
            } else {
                q=append(q, pair{i,j})
            }
        }
    }
    moves:=[][]int{{1,0},{-1,0},{0,1},{0,-1}}
    groups:=[][]pair{q}
    for len(q)>0 {
        l:=len(q)
        for i:=0; i<l; i++ {
            dis[q[i].x][q[i].y]=len(groups)-1
            for _,move:=range moves {
                ni:=q[i].x+move[0]
                nj:=q[i].y+move[1]
                if ni>=0 && ni<n && nj>=0 && nj<n && grid[ni][nj]==0 && dis[ni][nj]<0 {
                    q=append(q, pair{ni, nj})
                    grid[ni][nj]=2
                }
            }
            
        }
        q=q[l:]
        if len(q)>0 {
            groups=append(groups, q)
        }
    }
    // 并查集
    father := make([]int, n*n)
    for i := range father {
        father[i]=i
    }
    var find func(u int) int
    find=func(u int) int {
        if father[u]==u {
            return u
        }
        father[u]=find(father[u])
        return father[u]
    }
    var union func(u,v int)
    union=func(u,v int) {
        u=find(u)
        v=find(v)
        if u==v {
            return
        }
        father[v]=u
    } 
    var same func(u,v int) bool
    same = func(u,v int) bool {
        u = find(u)
        v = find(v)
        return u == v
    }
    for i := len(groups)-1; i>=0; i-- {
        for _,v := range groups[i] {
            for _, move := range moves {
                ni := v.x+move[0]
                nj := v.y+move[1]
                if ni >= 0 && ni < n && nj >= 0 && nj < n && dis[ni][nj] >= dis[v.x][v.y] {
                    union(v.x*n+v.y, ni*n+nj)
                }
            }
        }
        if same(0, n*n-1) {
            return i
        }
    }
    return 0
}

func main() {
    // Example 1:
    // Input: grid = [[1,0,0],[0,0,0],[0,0,1]]
    // Output: 0
    // Explanation: All paths from (0, 0) to (n - 1, n - 1) go through the thieves in cells (0, 0) and (n - 1, n - 1).
    fmt.Println(maximumSafenessFactor([][]int{{1,0,0},{0,0,0},{0,0,1}})) // 0
    // Example 2:
    // Input: grid = [[0,0,1],[0,0,0],[0,0,0]]
    // Output: 2
    // Explanation: The path depicted in the picture above has a safeness factor of 2 since:
    // - The closest cell of the path to the thief at cell (0, 2) is cell (0, 0). The distance between them is | 0 - 0 | + | 0 - 2 | = 2.
    // It can be shown that there are no other paths with a higher safeness factor.
    fmt.Println(maximumSafenessFactor([][]int{{0,0,1},{0,0,0},{0,0,0}})) // 2
    // Example 3:
    // Input: grid = [[0,0,0,1],[0,0,0,0],[0,0,0,0],[1,0,0,0]]
    // Output: 2
    // Explanation: The path depicted in the picture above has a safeness factor of 2 since:
    // - The closest cell of the path to the thief at cell (0, 3) is cell (1, 2). The distance between them is | 0 - 1 | + | 3 - 2 | = 2.
    // - The closest cell of the path to the thief at cell (3, 0) is cell (3, 2). The distance between them is | 3 - 3 | + | 0 - 2 | = 2.
    // It can be shown that there are no other paths with a higher safeness factor.
    fmt.Println(maximumSafenessFactor([][]int{{0,0,0,1},{0,0,0,0},{0,0,0,0},{1,0,0,0}})) // 2

    fmt.Println(maximumSafenessFactor1([][]int{{1,0,0},{0,0,0},{0,0,1}})) // 0
    fmt.Println(maximumSafenessFactor1([][]int{{0,0,1},{0,0,0},{0,0,0}})) // 2
    fmt.Println(maximumSafenessFactor1([][]int{{0,0,0,1},{0,0,0,0},{0,0,0,0},{1,0,0,0}})) // 2
}