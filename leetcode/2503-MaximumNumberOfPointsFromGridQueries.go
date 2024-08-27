package main

// 2503. Maximum Number of Points From Grid Queries
// You are given an m x n integer matrix grid and an array queries of size k.

// Find an array answer of size k such that for each integer queries[i] you start in the top left cell of the matrix and repeat the following process:
//     1. If queries[i] is strictly greater than the value of the current cell that you are in, 
//        then you get one point if it is your first time visiting this cell, 
//        and you can move to any adjacent cell in all 4 directions: up, down, left, and right.
//     2. Otherwise, you do not get any points, and you end this process.

// After the process, answer[i] is the maximum number of points you can get. 
// Note that for each query you are allowed to visit the same cell multiple times.

// Return the resulting array answer.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/10/19/yetgriddrawio.png" />
// Input: grid = [[1,2,3],[2,5,7],[3,5,1]], queries = [5,6,2]
// Output: [5,8,1]
// Explanation: The diagrams above show which cells we visit to get points for each query.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/10/20/yetgriddrawio-2.png" />
// Input: grid = [[5,2,1],[1,1,2]], queries = [3]
// Output: [0]
// Explanation: We can not get any points because the value of the top left cell is already greater than or equal to 3.
 
// Constraints:
//     m == grid.length
//     n == grid[i].length
//     2 <= m, n <= 1000
//     4 <= m * n <= 10^5
//     k == queries.length
//     1 <= k <= 10^4
//     1 <= grid[i][j], queries[i] <= 10^6

import "fmt"
import "sort"
import "container/heap"

// 并查集
func maxPoints(grid [][]int, queries []int) []int {
    dirs := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    m, n := len(grid), len(grid[0])
    mn := m * n

    // 并查集模板
    fa := make([]int, mn)
    size := make([]int, mn)
    for i := range fa {
        fa[i], size[i] = i, 1
    }
    var find func(int) int
    find = func(x int) int {
        if fa[x] != x {
            fa[x] = find(fa[x])
        }
        return fa[x]
    }
    merge := func(from, to int) {
        from = find(from)
        to = find(to)
        if from != to {
            fa[from] = to
            size[to] += size[from]
        }
    }

    // 矩阵元素从小到大排序，方便离线
    type Tuple struct{ x, i, j int }
    arr := make([]Tuple, 0, mn)
    for i, row := range grid {
        for j, x := range row {
            arr = append(arr, Tuple{x, i, j})
        }
    }
    sort.Slice(arr, func(i, j int) bool { 
        return arr[i].x < arr[j].x 
    })

    // 查询的下标按照查询值从小到大排序，方便离线
    id := make([]int, len(queries))
    for i := range id {
        id[i] = i
    }
    sort.Slice(id, func(i, j int) bool { 
        return queries[id[i]] < queries[id[j]] 
    })

    res, j := make([]int, len(queries)), 0
    for _, i := range id {
        q := queries[i]
        for ; j < mn && arr[j].x < q; j++ {
            x, y := arr[j].i, arr[j].j
            for _, d := range dirs { // 枚举周围四个格子，值小于 q 才可以合并
                x2, y2 := x + d.x, y + d.y
                if 0 <= x2 && x2 < m && 0 <= y2 && y2 < n && grid[x2][y2] < q {
                    merge(x * n + y, x2 * n + y2) // 把坐标压缩成一维的编号
                }
            }
        }
        if grid[0][0] < q {
            res[i] = size[find(0)] // 左上角的连通块的大小
        }
    }
    return res
}

func maxPoints1(grid [][]int, queries []int) []int {
    dirs := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    m, n := len(grid), len(grid[0])
    // 查询的下标按照查询值从小到大排序，方便离线
    id := make([]int, len(queries))
    for i := range id {
        id[i] = i
    }
    sort.Slice(id, func(i, j int) bool { 
        return queries[id[i]] < queries[id[j]] 
    })
    res := make([]int, len(queries))
    h, cnt := hp{{grid[0][0], 0, 0}}, 0
    grid[0][0] = 0 // 充当 vis 数组的作用
    for _, i := range id {
        q := queries[i]
        for len(h) > 0 && h[0].val < q {
            cnt++
            p := heap.Pop(&h).(Tuple)
            for _, d := range dirs { // 枚举周围四个格子
                x, y := p.i+d.x, p.j+d.y
                if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] > 0 {
                    heap.Push(&h, Tuple{grid[x][y], x, y})
                    grid[x][y] = 0 // 充当 vis 数组的作用
                }
            }
        }
        res[i] = cnt
    }
    return res
}

type Tuple struct{ val, i, j int }
type hp []Tuple
func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].val < h[j].val }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(Tuple)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/10/19/yetgriddrawio.png" />
    // Input: grid = [[1,2,3],[2,5,7],[3,5,1]], queries = [5,6,2]
    // Output: [5,8,1]
    // Explanation: The diagrams above show which cells we visit to get points for each query.
    grid1 := [][]int{
        {1,2,3},
        {2,5,7},
        {3,5,1},
    }
    fmt.Println(maxPoints(grid1, []int{5,6,2})) // [5,8,1]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/10/20/yetgriddrawio-2.png" />
    // Input: grid = [[5,2,1],[1,1,2]], queries = [3]
    // Output: [0]
    // Explanation: We can not get any points because the value of the top left cell is already greater than or equal to 3.
    grid2 := [][]int{
        {5,2,1},
        {1,1,2},
    }
    fmt.Println(maxPoints(grid2, []int{3})) // [0]

    fmt.Println(maxPoints1(grid1, []int{5,6,2})) // [5,8,1]
    fmt.Println(maxPoints1(grid2, []int{3})) // [0]
}