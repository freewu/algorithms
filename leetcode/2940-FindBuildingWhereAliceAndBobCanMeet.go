package main

// 2940. Find Building Where Alice and Bob Can Meet
// You are given a 0-indexed array heights of positive integers,
// where heights[i] represents the height of the ith building.

// If a person is in building i, they can move to any other building j if and only if i < j and heights[i] < heights[j].

// You are also given another array queries where queries[i] = [ai, bi]. On the ith query, Alice is in building ai while Bob is in building bi.

// Return an array ans where ans[i] is the index of the leftmost building where Alice and Bob can meet on the ith query. 
// If Alice and Bob cannot move to a common building on query i, set ans[i] to -1.

// Example 1:
// Input: heights = [6,4,8,5,2,7], queries = [[0,1],[0,3],[2,4],[3,4],[2,2]]
// Output: [2,5,-1,5,2]
// Explanation: In the first query, Alice and Bob can move to building 2 since heights[0] < heights[2] and heights[1] < heights[2]. 
// In the second query, Alice and Bob can move to building 5 since heights[0] < heights[5] and heights[3] < heights[5]. 
// In the third query, Alice cannot meet Bob since Alice cannot move to any other building.
// In the fourth query, Alice and Bob can move to building 5 since heights[3] < heights[5] and heights[4] < heights[5].
// In the fifth query, Alice and Bob are already in the same building.  
// For ans[i] != -1, It can be shown that ans[i] is the leftmost building where Alice and Bob can meet.
// For ans[i] == -1, It can be shown that there is no building where Alice and Bob can meet.

// Example 2:
// Input: heights = [5,3,8,2,6,1,4,6], queries = [[0,7],[3,5],[5,2],[3,0],[1,6]]
// Output: [7,6,-1,4,6]
// Explanation: In the first query, Alice can directly move to Bob's building since heights[0] < heights[7].
// In the second query, Alice and Bob can move to building 6 since heights[3] < heights[6] and heights[5] < heights[6].
// In the third query, Alice cannot meet Bob since Bob cannot move to any other building.
// In the fourth query, Alice and Bob can move to building 4 since heights[3] < heights[4] and heights[0] < heights[4].
// In the fifth query, Alice can directly move to Bob's building since heights[1] < heights[6].
// For ans[i] != -1, It can be shown that ans[i] is the leftmost building where Alice and Bob can meet.
// For ans[i] == -1, It can be shown that there is no building where Alice and Bob can meet.

// Constraints:
//     1 <= heights.length <= 5 * 10^4
//     1 <= heights[i] <= 10^9
//     1 <= queries.length <= 5 * 10^4
//     queries[i] = [ai, bi]
//     0 <= ai, bi <= heights.length - 1

import "fmt"
import "sort"
import "container/heap"
import "math/bits"

// 最小堆
func leftmostBuildingQueries(heights []int, queries [][]int) []int {
    res := make([]int, len(queries))
    for i := range res {
        res[i] = -1
    }
    qs := make([][]pair, len(heights))
    for i, q := range queries {
        a, b := q[0], q[1]
        if a > b {
            a, b = b, a // 保证 a <= b
        }
        if a == b || heights[a] < heights[b] {
            res[i] = b // a 直接跳到 b
        } else {
            qs[b] = append(qs[b], pair{heights[a], i}) // 离线询问
        }
    }
    h := hp{}
    for i, x := range heights {
        for h.Len() > 0 && h[0].h < x {
            // 堆顶的 heights[a] 可以跳到 heights[i]
            res[heap.Pop(&h).(pair).i] = i
        }
        for _, p := range qs[i] {
            heap.Push(&h, p) // 后面再回答
        }
    }
    return res
}

type pair struct{ h, i int }
type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].h < h[j].h }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

// 单调栈
func leftmostBuildingQueries1(heights []int, queries [][]int) []int {
    res := make([]int, len(queries))
    type pair struct{ h, i int }
    qs := make([][]pair, len(heights))
    for i, q := range queries {
        a, b := q[0], q[1]
        if a > b {
            a, b = b, a // 保证 a <= b
        }
        if a == b || heights[a] < heights[b] {
            res[i] = b // a 直接跳到 b
        } else {
            qs[b] = append(qs[b], pair{heights[a], i}) // 离线询问
        }
    }
    stack := []int{}
    for i := len(heights) - 1; i >= 0; i-- {
        for _, q := range qs[i] {
            j := sort.Search(len(stack), func(i int) bool { return heights[stack[i]] <= q.h }) - 1
            if j >= 0 {
                res[q.i] = stack[j]
            } else {
                res[q.i] = -1
            }
        }
        for len(stack) > 0 && heights[i] >= heights[stack[len(stack)-1]] {
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, i)
    }
    return res
}

type seg []int

// 初始化线段树，维护区间最大值
func (t seg) build(a []int, o, l, r int) {
    if l == r {
        t[o] = a[l]
        return
    }
    m := (l + r) >> 1
    t.build(a, o<<1, l, m)
    t.build(a, o<<1|1, m+1, r)
    t[o] = max(t[o<<1], t[o<<1|1])
}

// 返回 [L,n-1] 中第一个 > v 的值的下标
// 如果不存在，返回 -1
func (t seg) query(o, l, r, L, v int) int {
    if t[o] <= v { // 区间最大值 <= v
        return -1 // 没有 > v 的数
    }
    if l == r { // 找到了
        return l
    }
    m := (l + r) >> 1
    if L <= m {
        pos := t.query(o<<1, l, m, L, v) // 递归左子树
        if pos >= 0 { // 找到了
            return pos
        }
    }
    return t.query(o<<1|1, m+1, r, L, v) // 递归右子树
}

func leftmostBuildingQueries2(heights []int, queries [][]int) []int {
    n := len(heights)
    t := make(seg, 2<<bits.Len(uint(n-1)))
    t.build(heights, 1, 0, n-1)
    res := make([]int, len(queries))
    for i, q := range queries {
        a, b := q[0], q[1]
        if a > b {
            a, b = b, a // 保证 a <= b
        }
        if a == b || heights[a] < heights[b] {
            res[i] = b // a 直接跳到 b
        } else {
            // 线段树二分，找 [b+1,n-1] 中的第一个 > heights[a] 的位置
            res[i] = t.query(1, 0, n-1, b+1, heights[a])
        }
    }
    return res
}


func main() {
    // Example 1:
    // Input: heights = [6,4,8,5,2,7], queries = [[0,1],[0,3],[2,4],[3,4],[2,2]]
    // Output: [2,5,-1,5,2]
    // Explanation: In the first query, Alice and Bob can move to building 2 since heights[0] < heights[2] and heights[1] < heights[2]. 
    // In the second query, Alice and Bob can move to building 5 since heights[0] < heights[5] and heights[3] < heights[5]. 
    // In the third query, Alice cannot meet Bob since Alice cannot move to any other building.
    // In the fourth query, Alice and Bob can move to building 5 since heights[3] < heights[5] and heights[4] < heights[5].
    // In the fifth query, Alice and Bob are already in the same building.  
    // For ans[i] != -1, It can be shown that ans[i] is the leftmost building where Alice and Bob can meet.
    // For ans[i] == -1, It can be shown that there is no building where Alice and Bob can meet.
    fmt.Println(leftmostBuildingQueries([]int{6,4,8,5,2,7},[][]int{{0,1},{0,3},{2,4},{3,4},{2,2}})) // [2,5,-1,5,2]
    // Example 2:
    // Input: heights = [5,3,8,2,6,1,4,6], queries = [[0,7],[3,5],[5,2],[3,0],[1,6]]
    // Output: [7,6,-1,4,6]
    // Explanation: In the first query, Alice can directly move to Bob's building since heights[0] < heights[7].
    // In the second query, Alice and Bob can move to building 6 since heights[3] < heights[6] and heights[5] < heights[6].
    // In the third query, Alice cannot meet Bob since Bob cannot move to any other building.
    // In the fourth query, Alice and Bob can move to building 4 since heights[3] < heights[4] and heights[0] < heights[4].
    // In the fifth query, Alice can directly move to Bob's building since heights[1] < heights[6].
    // For ans[i] != -1, It can be shown that ans[i] is the leftmost building where Alice and Bob can meet.
    // For ans[i] == -1, It can be shown that there is no building where Alice and Bob can meet.
    fmt.Println(leftmostBuildingQueries([]int{5,3,8,2,6,1,4,6},[][]int{{0,7},{3,5},{5,2},{3,0},{1,6}})) // [7,6,-1,4,6]

    fmt.Println(leftmostBuildingQueries1([]int{6,4,8,5,2,7},[][]int{{0,1},{0,3},{2,4},{3,4},{2,2}})) // [2,5,-1,5,2]
    fmt.Println(leftmostBuildingQueries1([]int{5,3,8,2,6,1,4,6},[][]int{{0,7},{3,5},{5,2},{3,0},{1,6}})) // [7,6,-1,4,6]

    fmt.Println(leftmostBuildingQueries2([]int{6,4,8,5,2,7},[][]int{{0,1},{0,3},{2,4},{3,4},{2,2}})) // [2,5,-1,5,2]
    fmt.Println(leftmostBuildingQueries2([]int{5,3,8,2,6,1,4,6},[][]int{{0,7},{3,5},{5,2},{3,0},{1,6}})) // [7,6,-1,4,6]
}