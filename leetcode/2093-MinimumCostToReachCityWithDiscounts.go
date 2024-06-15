package main

// 2093. Minimum Cost to Reach City With Discounts
// A series of highways connect n cities numbered from 0 to n - 1. 
// You are given a 2D integer array highways where highways[i] = [city1i, city2i, tolli] indicates 
// that there is a highway that connects city1i and city2i, allowing a car to go from city1i to city2i and vice versa for a cost of tolli.

// You are also given an integer discounts which represents the number of discounts you have. 
// You can use a discount to travel across the ith highway for a cost of tolli / 2 (integer division). 
// Each discount may only be used once, and you can only use at most one discount per highway.

// Return the minimum total cost to go from city 0 to city n - 1, or -1 if it is not possible to go from city 0 to city n - 1.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/11/29/image-20211129222429-1.png" />
// Input: n = 5, highways = [[0,1,4],[2,1,3],[1,4,11],[3,2,3],[3,4,2]], discounts = 1
// Output: 9
// Explanation:
// Go from 0 to 1 for a cost of 4.
// Go from 1 to 4 and use a discount for a cost of 11 / 2 = 5.
// The minimum cost to go from 0 to 4 is 4 + 5 = 9.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/11/29/image-20211129222650-4.png" />
// Input: n = 4, highways = [[1,3,17],[1,2,7],[3,2,5],[0,1,6],[3,0,20]], discounts = 20
// Output: 8
// Explanation:
// Go from 0 to 1 and use a discount for a cost of 6 / 2 = 3.
// Go from 1 to 2 and use a discount for a cost of 7 / 2 = 3.
// Go from 2 to 3 and use a discount for a cost of 5 / 2 = 2.
// The minimum cost to go from 0 to 3 is 3 + 3 + 2 = 8.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/11/29/image-20211129222531-3.png" />
// Input: n = 4, highways = [[0,1,3],[2,3,2]], discounts = 0
// Output: -1
// Explanation:
// It is impossible to go from 0 to 3 so return -1.

// Constraints:
//     2 <= n <= 1000
//     1 <= highways.length <= 1000
//     highways[i].length == 3
//     0 <= city1i, city2i <= n - 1
//     city1i != city2i
//     0 <= tolli <= 10^5
//     0 <= discounts <= 500
//     There are no duplicate highways.

import "fmt"
import "container/heap"

func minimumCost(n int, highways [][]int, discounts int) int {
    edges, inf := make([][][]int, n), 1 << 32 - 1
    for _, highway := range highways {
        u, v, c := highway[0], highway[1], highway[2]
        edges[u] = append(edges[u], []int{v, c})
        edges[v] = append(edges[v], []int{u, c})
    }
    dis, state, queue := make([][]int, n), make([][]bool, n), [][]int{}
    for i := 0; i < n; i++ {
        dis[i] = make([]int, discounts+1)
        for j := 0; j <= discounts; j++ {
            dis[i][j] = inf
        }
    }
    for i := 0; i < n; i++ {
        state[i] = make([]bool, discounts + 1)
    }
    dis[0][discounts] = 0
    state[0][discounts] = true
    queue = append(queue, []int{0, discounts})
    for len(queue) > 0 {
        cur, cnt := queue[0][0], queue[0][1]
        queue = queue[1:]
        state[cur][cnt] = false
        for _, ints := range edges[cur] {
            next, c := ints[0], ints[1]
            if dis[next][cnt] > dis[cur][cnt]+c {
                dis[next][cnt] = dis[cur][cnt] + c
                if !state[next][cnt] {
                    state[next][cnt] = true
                    queue = append(queue, []int{next, cnt})
                }
            }
        }
        if cnt > 0 {
            for _, ints := range edges[cur] {
                next, c := ints[0], ints[1]
                if dis[next][cnt-1] > dis[cur][cnt]+c /2{
                    dis[next][cnt-1] =dis[cur][cnt] + c / 2
                    if !state[next][cnt-1] {
                        state[next][cnt-1] = true
                        queue = append(queue, []int{next, cnt - 1})
                    }
                }
            }
        }
    }
    res := inf
    for i := 0; i < discounts+1; i++ {
        if res > dis[n-1][i] {
            res = dis[n-1][i]
        }
    }
    if res == inf {
        return -1
    }
    return res
}

// 前往目标城市的最小费用
// 一组公路连接 n 个城市，城市编号为从 0 到 n - 1 。 输入包含一个二维数组 highways ，
// 其中 highways[i] = [city1i, city2i, tolli] 表示有一条连接城市 city1i 和 city2i 的双向公路，允许汽车缴纳值为 tolli 的费用从  city1i 前往 city2i 或 从  city2i 前往 city1i 。
// 另给你一个整数 discounts 表示你最多可以使用折扣的次数。你可以使用一次折扣使通过第 ith 条公路的费用降低至 tolli / 2（向下取整）。
// 最多只可使用 discounts 次折扣， 且 每条公路最多只可使用一次折扣 。
// 返回从城市0 前往城市 n - 1 的 最小费用 。如果不存在从城市0 前往城市 n - 1 的路径，返回 -1 。
func minimumCost1(n int, highways [][]int, discounts int) int {
    // 分层最短路, 两个资源 折扣次数和 cost
    const inf = 1 << 32 - 1
    type pair struct{ to, toll int }
    g := make([][]pair, n)
    for _, e := range highways {
        u, v, wt := e[0], e[1], e[2]
        g[u] = append(g[u], pair{v, wt})
        g[v] = append(g[v], pair{u, wt})
    }
    tolls := make([][]int, n) // 一维:Vertex 二维:discount
    for i := range tolls {
        tolls[i] = make([]int, discounts+1) // [0,discounts]
        for j := range tolls[i] {
            tolls[i][j] = inf
        }
    }
    hp := &dijkstraTollHeap{}
    heap.Init(hp)
    hp.push(dijkstraTollPair{0, 0, 0})
    tolls[0][0] = 0
    for hp.Len() > 0 {
        cur := hp.pop()
        u, toll, cnt := cur.id, cur.toll, cur.discount
        if toll > tolls[u][cnt] {
            continue
        }
        if u == n-1 { // trick!! 即使是分层,使用折扣最多的层一定先计算完毕(某点的值小,先弹出,促使其后继也先弹出),所以end第一次弹出时,一定是花费最小
            return toll
        }
        for _, nx := range g[u] {
            v, cost := nx.to, nx.toll
            if cnt < discounts { // 尝试使用折扣
                newCnt := cnt + 1
                newToll := toll + cost/2
                if newToll < tolls[v][newCnt] {
                    tolls[v][newCnt] = newToll
                    hp.push(dijkstraTollPair{v, newToll, newCnt})
                }
            }
            // 尝试不使用折扣
            newToll := toll + cost // 注意!! 重新计算,防止被上面的分支的相同变量值影响
            if newToll < tolls[v][cnt] {
                tolls[v][cnt] = newToll
                hp.push(dijkstraTollPair{v, newToll, cnt})
            }
        }
    }
    return -1
}

type dijkstraTollPair struct{ id, toll, discount int }
type dijkstraTollHeap []dijkstraTollPair

func (h dijkstraTollHeap) Len() int                   { return len(h) }
func (h dijkstraTollHeap) Less(i, j int) bool         { return h[i].toll < h[j].toll }
func (h dijkstraTollHeap) Swap(i, j int)              { h[i], h[j] = h[j], h[i] }
func (h *dijkstraTollHeap) Push(v any)                { *h = append(*h, v.(dijkstraTollPair)) }
func (h *dijkstraTollHeap) Pop() (v any)              { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *dijkstraTollHeap) push(v dijkstraTollPair)   { heap.Push(h, v) }
func (h *dijkstraTollHeap) pop() (v dijkstraTollPair) { return heap.Pop(h).(dijkstraTollPair) }

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/11/29/image-20211129222429-1.png" />
    // Input: n = 5, highways = [[0,1,4],[2,1,3],[1,4,11],[3,2,3],[3,4,2]], discounts = 1
    // Output: 9
    // Explanation:
    // Go from 0 to 1 for a cost of 4.
    // Go from 1 to 4 and use a discount for a cost of 11 / 2 = 5.
    // The minimum cost to go from 0 to 4 is 4 + 5 = 9.
    highways1 := [][]int{
        {0,1,4},
        {2,1,3},
        {1,4,11},
        {3,2,3},
        {3,4,2},
    }
    fmt.Println(minimumCost(5,highways1, 1)) // 9
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/11/29/image-20211129222650-4.png" />
    // Input: n = 4, highways = [[1,3,17],[1,2,7],[3,2,5],[0,1,6],[3,0,20]], discounts = 20
    // Output: 8
    // Explanation:
    // Go from 0 to 1 and use a discount for a cost of 6 / 2 = 3.
    // Go from 1 to 2 and use a discount for a cost of 7 / 2 = 3.
    // Go from 2 to 3 and use a discount for a cost of 5 / 2 = 2.
    // The minimum cost to go from 0 to 3 is 3 + 3 + 2 = 8.
    highways2:= [][]int{
        {1,3,17},
        {1,2,7},
        {3,2,5},
        {0,1,6},
        {3,0,20},
    }
    fmt.Println(minimumCost(4,highways2, 20)) // 8
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/11/29/image-20211129222531-3.png" />
    // Input: n = 4, highways = [[0,1,3],[2,3,2]], discounts = 0
    // Output: -1
    // Explanation:
    // It is impossible to go from 0 to 3 so return -1.
    highways3:= [][]int{
        {0,1,3},
        {2,3,2},
    }
    fmt.Println(minimumCost(4,highways3, 0)) // -1

    fmt.Println(minimumCost1(5,highways1, 1)) // 9
    fmt.Println(minimumCost1(4,highways2, 20)) // 8
    fmt.Println(minimumCost1(4,highways3, 0)) // -1
}