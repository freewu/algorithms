package main

// 1263. Minimum Moves to Move a Box to Their Target Location
// A storekeeper is a game in which the player pushes boxes around in a warehouse trying to get them to target locations.

// The game is represented by an m x n grid of characters grid where each element is a wall, floor, or box.

// Your task is to move the box 'B' to the target position 'T' under the following rules:
//     The character 'S' represents the player. The player can move up, down, left, right in grid if it is a floor (empty cell).
//     The character '.' represents the floor which means a free cell to walk.
//     The character '#' represents the wall which means an obstacle (impossible to walk there).
//     There is only one box 'B' and one target cell 'T' in the grid.
//     The box can be moved to an adjacent free cell by standing next to the box and then moving in the direction of the box. This is a push.
//     The player cannot walk through the box.

// Return the minimum number of pushes to move the box to the target. 
// If there is no way to reach the target, return -1.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/11/06/sample_1_1620.png" />
// Input: grid = [["#","#","#","#","#","#"],
//                ["#","T","#","#","#","#"],
//                ["#",".",".","B",".","#"],
//                ["#",".","#","#",".","#"],
//                ["#",".",".",".","S","#"],
//                ["#","#","#","#","#","#"]]
// Output: 3
// Explanation: We return only the number of times the box is pushed.

// Example 2:
// Input: grid = [["#","#","#","#","#","#"],
//                ["#","T","#","#","#","#"],
//                ["#",".",".","B",".","#"],
//                ["#","#","#","#",".","#"],
//                ["#",".",".",".","S","#"],
//                ["#","#","#","#","#","#"]]
// Output: -1

// Example 3:
// Input: grid = [["#","#","#","#","#","#"],
//                ["#","T",".",".","#","#"],
//                ["#",".","#","B",".","#"],
//                ["#",".",".",".",".","#"],
//                ["#",".",".",".","S","#"],
//                ["#","#","#","#","#","#"]]
// Output: 5
// Explanation: push the box down, left, left, up and up.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 20
//     grid contains only characters '.', '#', 'S', 'T', or 'B'.
//     There is only one character 'S', 'B', and 'T' in the grid.

import "fmt"
import "container/heap"
import "math"

type Items struct {
    box, man            [2]int
    heuristic, distance int
}

type PriorityQueue []Items

func (pq PriorityQueue)  Len() int { return len(pq) }
func (pq PriorityQueue)  Less(i, j int) bool { return pq[i].heuristic < pq[j].heuristic }
func (pq PriorityQueue)  Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(Items)) }
func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    x := old[n-1]
    *pq = old[:n-1]
    return x
}

func checkPosition(grid [][]byte, point [2]int) bool {
    return point[0] >= 0 && point[0] < len(grid) && point[1] >= 0 && point[1] < len(grid[0]) && grid[point[0]][point[1]] != '#'
}

func getPoints(grid [][]byte) ([2]int, [2]int, [2]int) {
    target, box, start := [2]int{}, [2]int{}, [2]int{}
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            switch grid[i][j] {
            case 'S':
                start = [2]int{i, j}
            case 'T':
                target = [2]int{i, j}
            case 'B':
                box = [2]int{i, j}
            }
        }
    }
    return target, box, start
}

func heuristic(a, b [2]int) int {
    return int(math.Abs(float64(a[0]-b[0])) + math.Abs(float64(a[1]-b[1])))
}

func minPushBox(grid [][]byte) int {
    target, box, start := getPoints(grid)
    visited := make(map[[2][2]int]bool)
    pq := new(PriorityQueue)
    heap.Init(pq)
    heap.Push(pq, Items{box, start, heuristic(box, target), 0})
    for pq.Len() > 0 {
        item := heap.Pop(pq).(Items)
        if item.box == target {
            return item.distance
        }
        if visited[[2][2]int{item.box, item.man}] {
            continue
        }
        visited[[2][2]int{item.box, item.man}] = true
        for _, dir := range [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
            nextMan := [2]int{item.man[0] + dir[0], item.man[1] + dir[1]}
            if !checkPosition(grid, nextMan) {
                continue
            }
            if nextMan == item.box {
                nextBox := [2]int{nextMan[0] + dir[0], nextMan[1] + dir[1]}
                if !checkPosition(grid, nextBox) {
                    continue
                }
                heap.Push(pq, Items{
                    nextBox,
                    nextMan,
                    heuristic(nextBox, target) + item.distance + 1,
                    item.distance + 1,
                })
            } else {
                heap.Push(pq, Items{
                    item.box,
                    nextMan,
                    item.heuristic,
                    item.distance,
                })
            }
        }
    }
    return -1
}

func minPushBox1(grid [][]byte) int {
    m, n := len(grid), len(grid[0])
    var sx, sy, bx, by int // 玩家、箱子的初始位置
    for x := 0; x < m; x++ {
        for y := 0; y < n; y++ {
            if grid[x][y] == 'S' {
                sx, sy = x, y
            } else if grid[x][y] == 'B' {
                bx, by = x, y
            }
        }
    }
    check := func(x, y int) bool { // 不越界且不在墙上
        return x >= 0 && x < m && y >= 0 && y < n && grid[x][y] != '#'
    }
    d := []int{0, -1, 0, 1, 0}
    dp := make([][]int, m*n)
    for i := 0; i < m*n; i++ {
        dp[i] = make([]int, m*n)
        for j := 0; j < m*n; j++ {
            dp[i][j] = 0x3f3f3f3f
        }
    }
    dp[sx*n+sy][bx*n+by] = 0 // 初始状态的推动次数为 0
    q := [][2]int{{sx*n + sy, bx*n + by}}
    for len(q) > 0 {
        q1 := [][2]int{}
        for len(q) > 0 {
            s1, b1 := q[0][0], q[0][1]
            q = q[1:]
            sx1, sy1, bx1, by1 := s1/n, s1%n, b1/n, b1%n
            if grid[bx1][by1] == 'T' { // 箱子已被推到目标处
                return dp[s1][b1]
            }
            for i := 0; i < 4; i++ { // 玩家向四个方向移动到另一个状态
                sx2, sy2 := sx1+d[i], sy1+d[i+1]
                s2 := sx2*n + sy2
                if !check(sx2, sy2) { // 玩家位置不合法
                    continue
                }
                if bx1 == sx2 && by1 == sy2 { // 推动箱子
                    bx2, by2 := bx1+d[i], by1+d[i+1]
                    b2 := bx2*n + by2
                    if !check(bx2, by2) || dp[s2][b2] <= dp[s1][b1]+1 { // 箱子位置不合法 或 状态已访问
                        continue
                    }
                    dp[s2][b2] = dp[s1][b1] + 1
                    q1 = append(q1, [2]int{s2, b2})
                } else {
                    if dp[s2][b1] <= dp[s1][b1] { // 状态已访问
                        continue
                    }
                    dp[s2][b1] = dp[s1][b1]
                    q = append(q, [2]int{s2, b1})
                }
            }
        }
        q = q1
    }
    return -1
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/11/06/sample_1_1620.png" />
    // Input: grid = [["#","#","#","#","#","#"],
    //                ["#","T","#","#","#","#"],
    //                ["#",".",".","B",".","#"],
    //                ["#",".","#","#",".","#"],
    //                ["#",".",".",".","S","#"],
    //                ["#","#","#","#","#","#"]]
    // Output: 3
    // Explanation: We return only the number of times the box is pushed.
    grid1 := [][]byte{
        {'#','#','#','#','#','#'},
        {'#','T','#','#','#','#'},
        {'#','.','.','B','.','#'},
        {'#','.','#','#','.','#'},
        {'#','.','.','.','S','#'},
        {'#','#','#','#','#','#'},
    }
    fmt.Println(minPushBox(grid1)) // 3
    // Example 2:
    // Input: grid = [["#","#","#","#","#","#"],
    //                ["#","T","#","#","#","#"],
    //                ["#",".",".","B",".","#"],
    //                ["#","#","#","#",".","#"],
    //                ["#",".",".",".","S","#"],
    //                ["#","#","#","#","#","#"]]
    // Output: -1
    grid2 := [][]byte{
        {'#','#','#','#','#','#'},
        {'#','T','#','#','#','#'},
        {'#','.','.','B','.','#'},
        {'#','#','#','#','.','#'},
        {'#','.','.','.','S','#'},
        {'#','#','#','#','#','#'},
    }
    fmt.Println(minPushBox(grid2)) // -1
    // Example 3:
    // Input: grid = [["#","#","#","#","#","#"],
    //                ["#","T",".",".","#","#"],
    //                ["#",".","#","B",".","#"],
    //                ["#",".",".",".",".","#"],
    //                ["#",".",".",".","S","#"],
    //                ["#","#","#","#","#","#"]]
    // Output: 5
    // Explanation: push the box down, left, left, up and up.
    grid3 := [][]byte{
        {'#','#','#','#','#','#'},
        {'#','T','.','.','#','#'},
        {'#','.','#','B','.','#'},
        {'#','.','.','.','.','#'},
        {'#','.','.','.','S','#'},
        {'#','#','#','#','#','#'},
    }
    fmt.Println(minPushBox(grid3)) // 5

    fmt.Println(minPushBox1(grid1)) // 3
    fmt.Println(minPushBox1(grid2)) // -1
    fmt.Println(minPushBox1(grid3)) // 5
}