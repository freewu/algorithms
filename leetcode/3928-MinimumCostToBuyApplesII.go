package main

// 3928. Minimum Cost to Buy Apples II
// You are given an integer n and an integer array prices of length n, where prices[i] is the price of apples at shop i.

// You are also given a 2D integer array roads, where roads[i] = [ui, vi, costi, taxi] represents a bidirectional road:
//     1. ui and vi are the shops connected by the road.
//     2. costi is the cost to travel the road without carrying apples.
//     3. taxi is the multiplier applied to costi when traveling with apples.

// For each shop i, you can either:
//     1. Buy apples locally at shop i for prices[i].
//     2. Travel empty to any shop j using any number of roads, buy apples for prices[j], 
//        and return to shop i while carrying apples, paying cost * tax on each road used for the return trip.

// The forward path, where you travel empty, and the return path may be different.

// Return an integer array ans of length n, where ans[i] is the minimum total cost to buy apples starting from shop i.

// Example 1:
// Input: n = 2, prices = [8,3], roads = [[0,1,1,2]]
// Output: [6,3]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/08/22/screenshot-2025-08-23-at-23341-am.png" />
// Shop i | prices[i] | Shop j | prices[j]	| costi | taxi | Travel cost | Return cost  | Total           | Minimum
// 0	   | 8         |	1	| 3	        | 1	    | 2    |	1	     | 1 * 2 = 2	| 1 + 2 + 3 = 6   | min(8, 6) = 6
// 1	   | 3         |	0	| 8	        | 1	    | 2    |	1	     | 1 * 2 = 2	| 1 + 2 + 8 = 11  | min(3, 11) = 3
// Thus, the answer is [6, 3].

// Example 2:
// Input: n = 3, prices = [9,4,6], roads = [[0,1,1,3],[1,2,4,2]]
// Output: [8,4,6]
// Explanation:
// ​​​​​​<img src="https://assets.leetcode.com/uploads/2025/08/22/screenshot-2025-08-23-at-23736-am.png" />
// Shop i  | prices[i] | Shop j | prices[j]	| costi | taxi  | Travel cost    | Return cost    | Total           | Minimum
// 0	    | 9 	    | 1	     | 4	        | 1	    | 3	    | 1	             | 1 * 3 = 3	  | 1 + 3 + 4 = 8	| min(9, 8) = 8
// 1	    | 4	        | 2	     | 6	        | 4	    | 2	    | 4	             | 4 * 2 = 8	  | 4 + 8 + 6 = 18	| min(4, 18) = 4
// 2	    | 6	        | 1	     | 4	        | 2	    | 4	    | 4	             | 4 * 2 = 8	  | 4 + 8 + 4 = 16	| min(6, 16) = 6
// Thus, the answer is [8, 4, 6].

// Example 3:
// Input: n = 3, prices = [10,11,1], roads = [[0,2,1,3],[1,2,3,4],[0,1,5,2]]
// Output: [5,11,1]
// Explanation:
// ​​​​​​​​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2025/08/22/screenshot-2025-08-23-at-24644-am.png" />
// Shop i  | prices[i] | Shop j | prices[j]	| costi | taxi  | Travel cost    | Return cost    | Total           | Minimum
// 0	    | 10	    | 2	     | 1	        | 1	    | 3	    | 1	             | 1 * 3 = 3	  | 1 + 3 + 1 = 5	| min(10, 5) = 5
// 1	    | 11	    | 2	     | 3	        | 1	    | 3	    | 4	             | 3 * 4 = 12	  | 3 + 12 + 1 = 16	| min(11, 16) = 11
// 2	    | 1	        | 10	 | 3	        | 1	    | 3	    | 1	             | 1 * 3 = 3	  | 1 + 3 + 10 = 14	| min(1, 14) = 1
// Thus, the answer is [5, 11, 1].

// Constraints:
//     1 <= n <= 1000
//     prices.length == n
//     1 <= prices[i] <= 10^9
//     0 <= roads.length <= min(n × (n - 1) / 2, 2000)
//     roads[i] = [ui, vi, costi, taxi]
//     0 <= ui, vi <= n - 1
//     ui != vi
//     1 <= costi <= 10^9
//     ​​​​​​​1 <= tax​​​​​​​i <= 100​​​​​​​
//     There are no repeated edges.

import "fmt"
import "container/heap"

type Edge struct{ to, weight int }
type Pair struct{ distance, x int }
type MinHeap []Pair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(v any)        { *h = append(*h, v.(Pair)) }
func (h *MinHeap) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func dijkstra(g [][]Edge, start int, price int) []int {
    dis := make([]int, len(g))
    for i := range dis {
        dis[i] = price
    }
    dis[start] = 0
    h := MinHeap{{0, start}}

    for len(h) > 0 {
        top := heap.Pop(&h).(Pair)
        d, x := top.distance, top.x
        if d > dis[x] {
            continue
        }
        for _, e := range g[x] {
            y := e.to
            newD := d + e.weight
            if newD < dis[y] {
                dis[y] = newD
                heap.Push(&h, Pair{newD, y})
            }
        }
    }
    return dis
}

func minCost(n int, prices []int, roads [][]int) []int {
    g1 := make([][]Edge, n)
    g2 := make([][]Edge, n)
    for _, e := range roads {
        x, y, cost, tax := e[0], e[1], e[2], e[3]
        g1[x] = append(g1[x], Edge{y, cost})
        g1[y] = append(g1[y], Edge{x, cost})
        g2[x] = append(g2[x], Edge{y, cost * tax})
        g2[y] = append(g2[y], Edge{x, cost * tax})
    }
    res := make([]int, n)
    for i, price := range prices {
        dis1 := dijkstra(g1, i, price)
        dis2 := dijkstra(g2, i, price)
        val := 1 << 61
        for j, p := range prices {
            val = min(val, p + dis1[j] + dis2[j])
        }
        res[i] = val
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 2, prices = [8,3], roads = [[0,1,1,2]]
    // Output: [6,3]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/08/22/screenshot-2025-08-23-at-23341-am.png" />
    // Shop i | prices[i] | Shop j | prices[j]	| costi | taxi | Travel cost | Return cost  | Total           | Minimum
    // 0	   | 8         |	1	| 3	        | 1	    | 2    |	1	     | 1 * 2 = 2	| 1 + 2 + 3 = 6   | min(8, 6) = 6
    // 1	   | 3         |	0	| 8	        | 1	    | 2    |	1	     | 1 * 2 = 2	| 1 + 2 + 8 = 11  | min(3, 11) = 3
    // Thus, the answer is [6, 3].
    fmt.Println(minCost(2, []int{8,3}, [][]int{{0,1,1,2}})) // [6,3]
    // Example 2:
    // Input: n = 3, prices = [9,4,6], roads = [[0,1,1,3],[1,2,4,2]]
    // Output: [8,4,6]
    // Explanation:
    // ​​​​​​<img src="https://assets.leetcode.com/uploads/2025/08/22/screenshot-2025-08-23-at-23736-am.png" />
    // Shop i  | prices[i] | Shop j | prices[j]	| costi | taxi  | Travel cost    | Return cost    | Total           | Minimum
    // 0	    | 9 	    | 1	     | 4	        | 1	    | 3	    | 1	             | 1 * 3 = 3	  | 1 + 3 + 4 = 8	| min(9, 8) = 8
    // 1	    | 4	        | 2	     | 6	        | 4	    | 2	    | 4	             | 4 * 2 = 8	  | 4 + 8 + 6 = 18	| min(4, 18) = 4
    // 2	    | 6	        | 1	     | 4	        | 2	    | 4	    | 4	             | 4 * 2 = 8	  | 4 + 8 + 4 = 16	| min(6, 16) = 6
    // Thus, the answer is [8, 4, 6].
    fmt.Println(minCost(3, []int{9,4,6}, [][]int{{0,1,1,3},{1,2,4,2}})) // [8,4,6]
    // Example 3:
    // Input: n = 3, prices = [10,11,1], roads = [[0,2,1,3],[1,2,3,4],[0,1,5,2]]
    // Output: [5,11,1]
    // Explanation:
    // ​​​​​​​​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2025/08/22/screenshot-2025-08-23-at-24644-am.png" />
    // Shop i  | prices[i] | Shop j | prices[j]	| costi | taxi  | Travel cost    | Return cost    | Total           | Minimum
    // 0	    | 10	    | 2	     | 1	        | 1	    | 3	    | 1	             | 1 * 3 = 3	  | 1 + 3 + 1 = 5	| min(10, 5) = 5
    // 1	    | 11	    | 2	     | 3	        | 1	    | 3	    | 4	             | 3 * 4 = 12	  | 3 + 12 + 1 = 16	| min(11, 16) = 11
    // 2	    | 1	        | 10	 | 3	        | 1	    | 3	    | 1	             | 1 * 3 = 3	  | 1 + 3 + 10 = 14	| min(1, 14) = 1
    // Thus, the answer is [5, 11, 1].
    fmt.Println(minCost(3, []int{10,11,1}, [][]int{{0,2,1,3},{1,2,3,4},{0,1,5,2}})) // [5,11,1]
}