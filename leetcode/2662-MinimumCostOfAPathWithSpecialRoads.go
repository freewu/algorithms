package main

// 2662. Minimum Cost of a Path With Special Roads
// You are given an array start where start = [startX, startY] represents your initial position (startX, startY) in a 2D space. 
// You are also given the array target where target = [targetX, targetY] represents your target position (targetX, targetY).

// The cost of going from a position (x1, y1) to any other position in the space (x2, y2) is |x2 - x1| + |y2 - y1|.

// There are also some special roads. 
// You are given a 2D array specialRoads where specialRoads[i] = [x1i, y1i, x2i, y2i, costi] indicates that the ith special road goes in one direction from (x1i, y1i) to (x2i, y2i) with a cost equal to costi. 
// You can use each special road any number of times.

// Return the minimum cost required to go from (startX, startY) to (targetX, targetY).

// Example 1:
// Input: start = [1,1], target = [4,5], specialRoads = [[1,2,3,3,2],[3,4,4,5,1]]
// Output: 5
// Explanation:
// (1,1) to (1,2) with a cost of |1 - 1| + |2 - 1| = 1.
// (1,2) to (3,3). Use specialRoads[0] with the cost 2.
// (3,3) to (3,4) with a cost of |3 - 3| + |4 - 3| = 1.
// (3,4) to (4,5). Use specialRoads[1] with the cost 1.
// So the total cost is 1 + 2 + 1 + 1 = 5.

// Example 2:
// Input: start = [3,2], target = [5,7], specialRoads = [[5,7,3,2,1],[3,2,3,4,4],[3,3,5,5,5],[3,4,5,6,6]]
// Output: 7
// Explanation:
// It is optimal not to use any special edges and go directly from the starting to the ending position with a cost |5 - 3| + |7 - 2| = 7.
// Note that the specialRoads[0] is directed from (5,7) to (3,2).

// Example 3:
// Input: start = [1,1], target = [10,4], specialRoads = [[4,2,1,1,3],[1,2,7,4,4],[10,3,6,1,2],[6,1,1,2,3]]
// Output: 8
// Explanation:
// (1,1) to (1,2) with a cost of |1 - 1| + |2 - 1| = 1.
// (1,2) to (7,4). Use specialRoads[1] with the cost 4.
// (7,4) to (10,4) with a cost of |10 - 7| + |4 - 4| = 3.

// Constraints:
//     start.length == target.length == 2
//     1 <= startX <= targetX <= 10^5
//     1 <= startY <= targetY <= 10^5
//     1 <= specialRoads.length <= 200
//     specialRoads[i].length == 5
//     startX <= x1i, x2i <= targetX
//     startY <= y1i, y2i <= targetY
//     1 <= costi <= 10^5

import "fmt"
import "container/heap"

type Tuple struct {
    d, x, y int
}
type MinHeap []Tuple

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].d < h[j].d }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(v any)        { *h = append(*h, v.(Tuple)) }
func (h *MinHeap) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func minimumCost(start []int, target []int, specialRoads [][]int) int {
    res, n := 1 << 31, 100_000
    pq := MinHeap{ {0, start[0], start[1]} }
    visited := map[int]bool{}
    min := func(x, y int) int { if x < y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    dist := func(x1, y1, x2, y2 int) int { return abs(x1 - x2) + abs(y1 - y2) }
    for len(pq) > 0 {
        p := pq[0]
        heap.Pop(&pq)
        d, x, y := p.d, p.x, p.y
        if visited[x * n + y] { continue }
        visited[x * n + y] = true
        res = min(res, d + dist(x, y, target[0], target[1]))
        for _, r := range specialRoads {
            x1, y1, x2, y2, cost := r[0], r[1], r[2], r[3], r[4]
            heap.Push(&pq, Tuple{d + dist(x, y, x1, y1) + cost, x2, y2})
        }
    }
    return res
}

func minimumCost1(start, target []int, specialRoads [][]int) int {
    type Pair struct{ x, y int }
    t := Pair{target[0], target[1]}
    dis := make(map[Pair]int, len(specialRoads)+2)
    dis[t] = 1 << 31
    dis[Pair{start[0], start[1]}] = 0
    visited := make(map[Pair]bool, len(specialRoads) + 1)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for {
        point, distance := Pair{}, -1
        for p, d := range dis {
            if !visited[p] && (distance < 0 || d < distance) {
                point, distance = p, d
            }
        }
        if point == t {
            return distance
        }
        visited[point] = true
        dis[t] = min(dis[t], distance + t.x - point.x + t.y - point.y)
        for _, r := range specialRoads {
            w := Pair{ r[2], r[3] }
            d := distance + abs(r[0]-point.x) + abs(r[1]-point.y) + r[4]
            if dw, ok := dis[w]; !ok || d < dw {
                dis[w] = d
            }
        }
    }
}

func main() {
    // Example 1:
    // Input: start = [1,1], target = [4,5], specialRoads = [[1,2,3,3,2],[3,4,4,5,1]]
    // Output: 5
    // Explanation:
    // (1,1) to (1,2) with a cost of |1 - 1| + |2 - 1| = 1.
    // (1,2) to (3,3). Use specialRoads[0] with the cost 2.
    // (3,3) to (3,4) with a cost of |3 - 3| + |4 - 3| = 1.
    // (3,4) to (4,5). Use specialRoads[1] with the cost 1.
    // So the total cost is 1 + 2 + 1 + 1 = 5.
    fmt.Println(minimumCost([]int{1,1}, []int{4,5}, [][]int{{1,2,3,3,2},{3,4,4,5,1}})) // 5
    // Example 2:
    // Input: start = [3,2], target = [5,7], specialRoads = [[5,7,3,2,1],[3,2,3,4,4],[3,3,5,5,5],[3,4,5,6,6]]
    // Output: 7
    // Explanation:
    // It is optimal not to use any special edges and go directly from the starting to the ending position with a cost |5 - 3| + |7 - 2| = 7.
    // Note that the specialRoads[0] is directed from (5,7) to (3,2).
    fmt.Println(minimumCost([]int{3,2}, []int{5,7}, [][]int{{5,7,3,2,1},{3,2,3,4,4},{3,3,5,5,5},{3,4,5,6,6}})) // 7
    // Example 3:
    // Input: start = [1,1], target = [10,4], specialRoads = [[4,2,1,1,3],[1,2,7,4,4],[10,3,6,1,2],[6,1,1,2,3]]
    // Output: 8
    // Explanation:
    // (1,1) to (1,2) with a cost of |1 - 1| + |2 - 1| = 1.
    // (1,2) to (7,4). Use specialRoads[1] with the cost 4.
    // (7,4) to (10,4) with a cost of |10 - 7| + |4 - 4| = 3.
    fmt.Println(minimumCost([]int{1,1}, []int{10,4}, [][]int{{4,2,1,1,3},{1,2,7,4,4},{10,3,6,1,2},{6,1,1,2,3}})) // 8

    fmt.Println(minimumCost1([]int{1,1}, []int{4,5}, [][]int{{1,2,3,3,2},{3,4,4,5,1}})) // 5
    fmt.Println(minimumCost1([]int{3,2}, []int{5,7}, [][]int{{5,7,3,2,1},{3,2,3,4,4},{3,3,5,5,5},{3,4,5,6,6}})) // 7
    fmt.Println(minimumCost1([]int{1,1}, []int{10,4}, [][]int{{4,2,1,1,3},{1,2,7,4,4},{10,3,6,1,2},{6,1,1,2,3}})) // 8
}