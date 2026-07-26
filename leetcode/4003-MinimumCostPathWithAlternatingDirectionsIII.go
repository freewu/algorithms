package main

// 4003. Minimum Cost Path with Alternating Directions III
// You are given two integers m and n representing the number of rows and columns of a grid. 
// Your goal is to reach cell (m - 1, n - 1). 
// You are also given a 2D integer array penalty.

// The cost to enter cell (i, j) is (i + 1) * (j + 1).

// You begin at cell (0, 0) and initially pay its entrance cost. 
// Actions performed after entering (0, 0) are numbered starting from 1.

// On each action, you may move to an adjacent cell or wait in the current cell. 
// A move follows the parity rule if:
//     1. On an odd-numbered action, you move right or down.
//     2. On an even-numbered action, you move left or up.

// The cost of an action is determined as follows:
//     1. If you move according to the parity rule, pay only the entrance cost of the destination cell.
//     2. If you move in a direction that violates the parity rule, pay the entrance cost of the destination cell plus penalty[i][j], where (i, j) is the cell you move from.
//     3. If you wait in cell (i, j), pay penalty[i][j].

// After every move or wait, the action number increases by 1. 
// Therefore, the required parity alternates after every action, regardless of whether a penalty was paid.

// Return the minimum total cost required to reach (m - 1, n - 1).

// Example 1:
// Input: m = 2, n = 2, penalty = [[5,3],[1,4]]
// Output: 8
// Explanation:
// The optimal path is:
// Start at cell (0, 0) with entry cost (0 + 1) * (0 + 1) = 1.
// Move 1: Move down to cell (1, 0) with entry cost (1 + 1) * (0 + 1) = 2.
// Move 2: Move right to cell (1, 1) with entry cost (1 + 1) * (1 + 1) = 4 and an extra cost of penalty[1][0] = 1 for violating the even parity rule.
// Thus, the total cost is 1 + 2 + 4 + 1 = 8.

// Example 2:
// Input: m = 2, n = 2, penalty = [[0,7],[3,2]]
// Output: 7
// Explanation:
// The optimal path is:
// Start at cell (0, 0) with entry cost (0 + 1) * (0 + 1) = 1.
// Move 1: Wait at cell (0, 0) with an extra cost of penalty[0][0] = 0 to flip to even parity.
// Move 2: Move right to cell (0, 1) with entry cost (0 + 1) * (1 + 1) = 2 and an extra cost of penalty[0][0] = 0 for violating the even parity rule.
// Move 3: Move down to cell (1, 1) with entry cost (1 + 1) * (1 + 1) = 4.
// Thus, the total cost is 1 + 0 + 2 + 0 + 4 = 7.

// Example 3:
// Input: m = 2, n = 3, penalty = [[8,0,9],[7,4,1]]
// Output: 12
// Explanation:
// The optimal path is:
// Start at cell (0, 0) with entry cost (0 + 1) * (0 + 1) = 1.
// Move 1: Move right to cell (0, 1) with entry cost (0 + 1) * (1 + 1) = 2.
// Move 2: Move right to cell (0, 2) with entry cost (0 + 1) * (2 + 1) = 3 and an extra cost of penalty[0][1] = 0 for violating the even parity rule.
// Move 3: Move down to cell (1, 2) with entry cost (1 + 1) * (2 + 1) = 6.
// Thus, the total cost is 1 + 2 + 3 + 0 + 6 = 12.

// Constraints:
//     1 <= m, n <= 10^5
//     2 <= m * n <= 10^5
//     penalty.length == m
//     penalty[i].length == n
//     0 <= penalty[i][j] <= 10^5

import "fmt"
import "container/heap"

type Tuple struct{ distance, x, y, k int }
type MinHeap []Tuple
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(v any)        { *h = append(*h, v.(Tuple)) }  
func (h *MinHeap) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

// 奇数下标 1,3 对应向右或向下
// 偶数下标 0,2 对应向左或向上
var dirs = []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func minCost(m, n int, penalty [][]int) int64 {
    dis := make([][][2]int, m)
    for i := range dis {
        dis[i] = make([][2]int, n)
        for j := range dis[i] {
            dis[i][j] = [2]int{ 1 << 61, 1 << 61}
        }
    }
    // 支付 1 的入口代价
    dis[0][0][1] = 1
    h := &MinHeap{{1, 0, 0, 1}}
    for {
        top := heap.Pop(h).(Tuple)
        d, i, j, k := top.distance, top.x, top.y, top.k
        if i == m - 1 && j == n - 1 {
            return int64(d)
        }
        if d > dis[i][j][k] {
            continue
        }
        p := penalty[i][j]
        // 原地不动
        newDis := d + p
        if newDis < dis[i][j][k^1] {
            dis[i][j][k^1] = newDis
            heap.Push(h, Tuple{newDis, i, j, k ^ 1}) // k^1 切换行动编号的奇偶性
        }
        // 移动一步
        for idx, dir := range dirs {
            x, y := i+dir.x, j+dir.y
            if 0 <= x && x < m && 0 <= y && y < n {
                // 如果 k 和 idx 的奇偶性不同，那么违反了奇偶性规则，需要额外支付 p 的代价
                newDis := d + (x + 1) * (y + 1) + (idx % 2 ^ k) * p
                if newDis < dis[x][y][k^1] {
                    dis[x][y][k^1] = newDis
                    heap.Push(h, Tuple{newDis, x, y, k ^ 1}) // k^1 切换行动编号的奇偶性    
                }
            }
        }
    }
}

func minCost1(m int, n int, penalty [][]int) int64 {
    inf := int64(1) << 60
    dist := make([]int64, m*n*2)
    for i := range dist {
        dist[i] = inf
    }
    dist[0] = 1
    q := []int64{1 << 20} // encoded cost<<20 | state, state = (i*n+j)*2 + parity
    push := func(x int64) {
        q = append(q, x)
        for i := len(q) - 1; i > 0; {
            p := (i - 1) / 2
            if q[p] <= q[i] {
                break
            }
            q[p], q[i] = q[i], q[p]
            i = p
        }
    }
    pop := func() int64 {
        t := q[0]
        q[0] = q[len(q)-1]
        q = q[:len(q)-1]
        for i := 0; ; {
            l, s := 2*i+1, i
            if l < len(q) && q[l] < q[s] {
                s = l
            }
            if l+1 < len(q) && q[l+1] < q[s] {
                s = l + 1
            }
            if s == i {
                break
            }
            q[i], q[s] = q[s], q[i]
            i = s
        }
        return t
    }
    relax := func(st, nd int64) {
        if nd < dist[st] {
            dist[st] = nd
            push(nd<<20 | st)
        }
    }
    dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // right,down,left,up
    for len(q) > 0 {
        t := pop()
        d, st := t>>20, t&(1<<20-1)
        if d > dist[st] {
            continue
        }
        p := int(st & 1)
        c := int(st >> 1)
        i, j := c/n, c%n
        pen := int64(penalty[i][j])
        relax(int64(c*2+1-p), d+pen) // wait: pay penalty, flip parity
        for dd, dr := range dirs {
            ni, nj := i+dr[0], j+dr[1]
            if ni < 0 || ni >= m || nj < 0 || nj >= n {
                continue
            }
            cost := int64((ni + 1) * (nj + 1))
            if (p == 0) != (dd < 2) { // odd action: right/down follow; even: left/up
                cost += pen
            }
            relax(int64((ni*n+nj)*2+1-p), d + cost)
        }
    }
    g := (m*n - 1) * 2
    return min(dist[g], dist[g+1])
}

func main() {
    // Example 1:
    // Input: m = 2, n = 2, penalty = [[5,3],[1,4]]
    // Output: 8
    // Explanation:
    // The optimal path is:
    // Start at cell (0, 0) with entry cost (0 + 1) * (0 + 1) = 1.
    // Move 1: Move down to cell (1, 0) with entry cost (1 + 1) * (0 + 1) = 2.
    // Move 2: Move right to cell (1, 1) with entry cost (1 + 1) * (1 + 1) = 4 and an extra cost of penalty[1][0] = 1 for violating the even parity rule.
    // Thus, the total cost is 1 + 2 + 4 + 1 = 8.
    fmt.Println(minCost(2, 2, [][]int{{5,3},{1,4}})) // 8
    // Example 2:
    // Input: m = 2, n = 2, penalty = [[0,7],[3,2]]
    // Output: 7
    // Explanation:
    // The optimal path is:
    // Start at cell (0, 0) with entry cost (0 + 1) * (0 + 1) = 1.
    // Move 1: Wait at cell (0, 0) with an extra cost of penalty[0][0] = 0 to flip to even parity.
    // Move 2: Move right to cell (0, 1) with entry cost (0 + 1) * (1 + 1) = 2 and an extra cost of penalty[0][0] = 0 for violating the even parity rule.
    // Move 3: Move down to cell (1, 1) with entry cost (1 + 1) * (1 + 1) = 4.
    // Thus, the total cost is 1 + 0 + 2 + 0 + 4 = 7.
    fmt.Println(minCost(2, 2, [][]int{{0,7},{3,2}})) // 7
    // Example 3:
    // Input: m = 2, n = 3, penalty = [[8,0,9],[7,4,1]]
    // Output: 12
    // Explanation:
    // The optimal path is:
    // Start at cell (0, 0) with entry cost (0 + 1) * (0 + 1) = 1.
    // Move 1: Move right to cell (0, 1) with entry cost (0 + 1) * (1 + 1) = 2.
    // Move 2: Move right to cell (0, 2) with entry cost (0 + 1) * (2 + 1) = 3 and an extra cost of penalty[0][1] = 0 for violating the even parity rule.
    // Move 3: Move down to cell (1, 2) with entry cost (1 + 1) * (2 + 1) = 6.
    // Thus, the total cost is 1 + 2 + 3 + 0 + 6 = 12.
    fmt.Println(minCost(2, 3, [][]int{{8,0,9},{7,4,1}})) // 12

    fmt.Println(minCost1(2, 2, [][]int{{5,3},{1,4}})) // 8
    fmt.Println(minCost1(2, 2, [][]int{{0,7},{3,2}})) // 7
    fmt.Println(minCost1(2, 3, [][]int{{8,0,9},{7,4,1}})) // 12
}