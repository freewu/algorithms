package main

// 3342. Find Minimum Time to Reach Last Room II
// There is a dungeon with n x m rooms arranged as a grid.

// You are given a 2D array moveTime of size n x m, 
// where moveTime[i][j] represents the minimum time in seconds when you can start moving to that room. 
// You start from the room (0, 0) at time t = 0 and can move to an adjacent room. 
// Moving between adjacent rooms takes one second for one move and two seconds for the next, alternating between the two.

// Return the minimum time to reach the room (n - 1, m - 1).

// Two rooms are adjacent if they share a common wall, either horizontally or vertically.

// Example 1:
// Input: moveTime = [[0,4],[4,4]]
// Output: 7
// Explanation:
// The minimum time required is 7 seconds.
// At time t == 4, move from room (0, 0) to room (1, 0) in one second.
// At time t == 5, move from room (1, 0) to room (1, 1) in two seconds.

// Example 2:
// Input: moveTime = [[0,0,0,0],[0,0,0,0]]
// Output: 6
// Explanation:
// The minimum time required is 6 seconds.
// At time t == 0, move from room (0, 0) to room (1, 0) in one second.
// At time t == 1, move from room (1, 0) to room (1, 1) in two seconds.
// At time t == 3, move from room (1, 1) to room (1, 2) in one second.
// At time t == 4, move from room (1, 2) to room (1, 3) in two seconds.

// Example 3:
// Input: moveTime = [[0,1],[1,2]]
// Output: 4

// Constraints:
//     2 <= n == moveTime.length <= 750
//     2 <= m == moveTime[i].length <= 750
//     0 <= moveTime[i][j] <= 10^9

import "fmt"
import "container/heap"

type Tuple struct{ dis, x, y int }
type MinHeap []Tuple
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(v any)        { *h = append(*h, v.(Tuple)) }
func (h *MinHeap) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func minTimeToReach(moveTime [][]int) int {
    directions := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    inf, n, m := 1 << 31, len(moveTime), len(moveTime[0])
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, m)
        for j := range dp[i] {
            dp[i][j] = inf
        }
    }
    dp[0][0] = 0
    h := MinHeap{{}}
    for {
        top := heap.Pop(&h).(Tuple)
        i, j := top.x, top.y
        if i == n - 1 && j == m - 1 {
            return top.dis
        }
        if top.dis > dp[i][j] { continue }
        for _, d := range directions {
            x, y := i + d.x, j + d.y
            if 0 <= x && x < n && 0 <= y && y < m {
                mx := max(top.dis, moveTime[x][y]) + (i + j) % 2 + 1
                if mx < dp[x][y] {
                    dp[x][y] = mx
                    heap.Push(&h, Tuple{mx, x, y})
                }
            }
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: moveTime = [[0,4],[4,4]]
    // Output: 7
    // Explanation:
    // The minimum time required is 7 seconds.
    // At time t == 4, move from room (0, 0) to room (1, 0) in one second.
    // At time t == 5, move from room (1, 0) to room (1, 1) in two seconds.
    fmt.Println(minTimeToReach([][]int{{0,4},{4,4}})) // 7
    // Example 2:
    // Input: moveTime = [[0,0,0,0],[0,0,0,0]]
    // Output: 6
    // Explanation:
    // The minimum time required is 6 seconds.
    // At time t == 0, move from room (0, 0) to room (1, 0) in one second.
    // At time t == 1, move from room (1, 0) to room (1, 1) in two seconds.
    // At time t == 3, move from room (1, 1) to room (1, 2) in one second.
    // At time t == 4, move from room (1, 2) to room (1, 3) in two seconds.
    fmt.Println(minTimeToReach([][]int{{0,0,0,0},{0,0,0,0}})) // 6
}