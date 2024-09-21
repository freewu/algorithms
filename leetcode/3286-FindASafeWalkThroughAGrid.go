package main

// 3286. Find a Safe Walk Through a Grid
// You are given an m x n binary matrix grid and an integer health.

// You start on the upper-left corner (0, 0) and would like to get to the lower-right corner (m - 1, n - 1).

// You can move up, down, left, or right from one cell to another adjacent cell as long as your health remains positive.

// Cells (i, j) with grid[i][j] = 1 are considered unsafe and reduce your health by 1.

// Return true if you can reach the final cell with a health value of 1 or more, and false otherwise.

// Example 1:
// Input: grid = [[0,1,0,0,0],[0,1,0,1,0],[0,0,0,1,0]], health = 1
// Output: true
// Explanation:
// The final cell can be reached safely by walking along the gray cells below.
// <img src="https://assets.leetcode.com/uploads/2024/08/04/3868_examples_1drawio.png" />

// Example 2:
// Input: grid = [[0,1,1,0,0,0],[1,0,1,0,0,0],[0,1,1,1,0,1],[0,0,1,0,1,0]], health = 3
// Output: false
// Explanation:
// A minimum of 4 health points is needed to reach the final cell safely.
// <img scr="https://assets.leetcode.com/uploads/2024/08/04/3868_examples_2drawio.png" />

// Example 3:
// Input: grid = [[1,1,1],[1,0,1],[1,1,1]], health = 5
// Output: true
// Explanation:
// The final cell can be reached safely by walking along the gray cells below.
// <img scr="https://assets.leetcode.com/uploads/2024/08/04/3868_examples_3drawio.png" />
// Any path that does not go through the cell (1, 1) is unsafe since your health will drop to 0 when reaching the final cell.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 50
//     2 <= m * n
//     1 <= health <= m + n
//     grid[i][j] is either 0 or 1.

import "fmt"
import "container/heap"

// 定义优先队列类型
type item struct {
    x, y, cost int
}
type minHeap []*item
func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
    *h = append(*h, x.(*item))
}
func (h *minHeap) Pop() interface{} {
    old := *h
    n := len(old)
    lastItem := old[n-1]
    *h = old[0 : n-1]
    return lastItem
}

// dijkstra+小顶堆优化 使用 dijkstra 算法寻找从 (0,0) 到 (m-1,n-1) 的安全路径
func findSafeWalk(grid [][]int, health int) bool {
    m, n, inf := len(grid), len(grid[0]), 1 << 31
    cost := make([][]int, m)
    for i := range cost {
        cost[i] = make([]int, n)
        for j := range cost[i] {
            cost[i][j] = inf // const inf = int(^uint(0) >> 1) // 最大整数值，表示无穷大
        }
    }
    dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    h := &minHeap{}
    heap.Init(h)
    cost[0][0] = grid[0][0]
    heap.Push(h, &item{x: 0, y: 0, cost: cost[0][0]})
    for h.Len() > 0 {
        curItem := heap.Pop(h).(*item)
        x, y, curCost := curItem.x, curItem.y, curItem.cost
        if x == m-1 && y == n-1 {
            return curCost < health
        }
        for _, d := range dirs {
            nx, ny := x+d[0], y+d[1]
            if nx >= 0 && nx < m && ny >= 0 && ny < n && curCost+grid[nx][ny] < cost[nx][ny] {
                cost[nx][ny] = curCost+grid[nx][ny]
                heap.Push(h, &item{x: nx, y: ny, cost: cost[nx][ny]})
            }
        }
    }
    return false // 无法到达
}

// 双端队列（from: 灵茶山艾府）
// 用两个 slice 头对头拼在一起实现
// 在知道数据量的情况下，也可以直接创建个两倍数据量大小的 slice，然后用两个下标表示头尾，初始化在 slice 正中
// l-1,...1,0,0,1...,r-1
// Deque 数据结构实现
type Deque struct{ l, r [][3]int }

func (q Deque) Empty() bool {
    return len(q.l) == 0 && len(q.r) == 0
}

func (q Deque) Len() int {
    return len(q.l) + len(q.r)
}

func (q *Deque) PushFront(v interface{}) {
    q.l = append(q.l, v.([3]int))
}

func (q *Deque) PushBack(v interface{}) {
	q.r = append(q.r, v.([3]int))
}

func (q *Deque) PopFront() (v interface{}) {
    if len(q.l) > 0 {
        q.l, v = q.l[:len(q.l)-1], q.l[len(q.l)-1]
    } else {
        v, q.r = q.r[0], q.r[1:]
    }
    return
}

func (q *Deque) PopBack() (v interface{}) {
    if len(q.r) > 0 {
        q.r, v = q.r[:len(q.r)-1], q.r[len(q.r)-1]
    } else {
        v, q.l = q.l[0], q.l[1:]
    }
    return
}

func (q Deque) Front() interface{} {
    if len(q.l) > 0 {
        return q.l[len(q.l)-1]
    }
    return q.r[0]
}

func (q Deque) Back() interface{} {
    if len(q.r) > 0 {
        return q.r[len(q.r)-1]
    }
    return q.l[0]
}

func (q Deque) Get(i int) interface{} {
    if i < len(q.l) {
        return q.l[len(q.l)-1-i]
    }
    return q.r[i-len(q.l)]
}

// findSafeWalk 使用双端队列在网格上寻找安全路径
func findSafeWalk1(grid [][]int, health int) bool {
    m, n := len(grid), len(grid[0])
    visited := make([][]bool, m)
    for i := range visited {
        visited[i] = make([]bool, n)
    }
    dq := &Deque{}
    dq.PushBack([3]int{0, 0, grid[0][0]})
    visited[0][0] = true
    for !dq.Empty() {
        p := dq.PopFront().([3]int)
        x, y, curCost := p[0], p[1], p[2]
        if x == m-1 && y == n-1 {
            return curCost < health
        }
        for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
            nx, ny := x+d[0], y+d[1]
            if nx < 0 || nx >= m || ny < 0 || ny >= n || visited[nx][ny] {
                continue
            }
            cost := grid[nx][ny]
            if cost == 1 {
                dq.PushBack([3]int{nx, ny, curCost+cost})
            } else {
                dq.PushFront([3]int{nx, ny, curCost})
            }
            visited[nx][ny] = true
        }
    }
    return false // 无法到达
}

func findSafeWalk2(grid [][]int, health int) bool {
    type pair struct{ x, y int }
    dirs := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    m, n := len(grid), len(grid[0])
    dis := make([][]int, m)
    for i := range dis {
        dis[i] = make([]int, n)
        for j := range dis[i] {
            dis[i][j] = m * n
        }
    }
    dis[0][0] = grid[0][0]
    q := [2][]pair{{{}}} // 两个 slice 头对头来实现 deque
    for len(q[0]) > 0 || len(q[1]) > 0 {
        var p pair
        if len(q[0]) > 0 {
            p, q[0] = q[0][len(q[0])-1], q[0][:len(q[0])-1]
        } else {
            p, q[1] = q[1][0], q[1][1:]
        }
        for _, d := range dirs {
            x, y := p.x+d.x, p.y+d.y
            if 0 <= x && x < m && 0 <= y && y < n {
                g := grid[x][y]
                if dis[p.x][p.y]+g < dis[x][y] {
                    dis[x][y] = dis[p.x][p.y] + g
                    q[g] = append(q[g], pair{x, y})
                }
            }
        }
    }
    return dis[m-1][n-1] < health
}

func main() {
    // Example 1:
    // Input: grid = [[0,1,0,0,0],[0,1,0,1,0],[0,0,0,1,0]], health = 1
    // Output: true
    // Explanation:
    // The final cell can be reached safely by walking along the gray cells below.
    // <img src="https://assets.leetcode.com/uploads/2024/08/04/3868_examples_1drawio.png" />
    fmt.Println(findSafeWalk([][]int{{0,1,0,0,0},{0,1,0,1,0},{0,0,0,1,0}}, 1)) // true
    // Example 2:
    // Input: grid = [[0,1,1,0,0,0],[1,0,1,0,0,0],[0,1,1,1,0,1],[0,0,1,0,1,0]], health = 3
    // Output: false
    // Explanation:
    // A minimum of 4 health points is needed to reach the final cell safely.
    // <img scr="https://assets.leetcode.com/uploads/2024/08/04/3868_examples_2drawio.png" />
    fmt.Println(findSafeWalk([][]int{{0,1,1,0,0,0},{1,0,1,0,0,0},{0,1,1,1,0,1},{0,0,1,0,1,0}}, 3)) // false
    // Example 3:
    // Input: grid = [[1,1,1],[1,0,1],[1,1,1]], health = 5
    // Output: true
    // Explanation:
    // The final cell can be reached safely by walking along the gray cells below.
    // <img scr="https://assets.leetcode.com/uploads/2024/08/04/3868_examples_3drawio.png" />
    // Any path that does not go through the cell (1, 1) is unsafe since your health will drop to 0 when reaching the final cell.
    fmt.Println(findSafeWalk([][]int{{1,1,1},{1,0,1},{1,1,1}}, 5)) // true

    fmt.Println(findSafeWalk1([][]int{{0,1,0,0,0},{0,1,0,1,0},{0,0,0,1,0}}, 1)) // true
    fmt.Println(findSafeWalk1([][]int{{0,1,1,0,0,0},{1,0,1,0,0,0},{0,1,1,1,0,1},{0,0,1,0,1,0}}, 3)) // false
    fmt.Println(findSafeWalk1([][]int{{1,1,1},{1,0,1},{1,1,1}}, 5)) // true

    fmt.Println(findSafeWalk2([][]int{{0,1,0,0,0},{0,1,0,1,0},{0,0,0,1,0}}, 1)) // true
    fmt.Println(findSafeWalk2([][]int{{0,1,1,0,0,0},{1,0,1,0,0,0},{0,1,1,1,0,1},{0,0,1,0,1,0}}, 3)) // false
    fmt.Println(findSafeWalk2([][]int{{1,1,1},{1,0,1},{1,1,1}}, 5)) // true
}