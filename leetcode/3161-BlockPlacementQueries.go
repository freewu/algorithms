package main

// 3161. Block Placement Queries
// There exists an infinite number line, with its origin at 0 and extending towards the positive x-axis.

// You are given a 2D array queries, which contains two types of queries:
//     1. For a query of type 1, queries[i] = [1, x]. Build an obstacle at distance x from the origin. 
//        It is guaranteed that there is no obstacle at distance x when the query is asked.
//     2. For a query of type 2, queries[i] = [2, x, sz]. 
//        Check if it is possible to place a block of size sz anywhere in the range [0, x] on the line, such that the block entirely lies in the range [0, x]. 
//        A block cannot be placed if it intersects with any obstacle, but it may touch it. 
//        Note that you do not actually place the block. Queries are separate.

// Return a boolean array results, 
// where results[i] is true if you can place the block specified in the ith query of type 2, and false otherwise.

// Example 1:
// Input: queries = [[1,2],[2,3,3],[2,3,1],[2,2,2]]
// Output: [false,true,true]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/04/22/example0block.png" />
// For query 0, place an obstacle at x = 2. A block of size at most 2 can be placed before x = 3.

// Example 2:
// Input: queries = [[1,7],[2,7,6],[1,2],[2,7,5],[2,7,6]]
// Output: [true,true,false]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/04/22/example1block.png" />
// Place an obstacle at x = 7 for query 0. A block of size at most 7 can be placed before x = 7.
// Place an obstacle at x = 2 for query 2. Now, a block of size at most 5 can be placed before x = 7, and a block of size at most 2 before x = 2.

// Constraints:
//     1 <= queries.length <= 15 * 10^4
//     2 <= queries[i].length <= 3
//     1 <= queries[i][0] <= 2
//     1 <= x, sz <= min(5 * 10^4, 3 * queries.length)
//     The input is generated such that for queries of type 1, no obstacle exists at distance x when the query is asked.
//     The input is generated such that there is at least one query of type 2.

import "fmt"
import "slices"

// func max(x, y int) int { if x > y { return x; }; return y; }

// type Segment []int

// // 把 i 处的值改成 val
// func (t Segment) Update(o, l, r, i, val int) {
//     if l == r {
//         t[o] = val
//         return
//     }
//     mid := (l + r) >> 1
//     if i <= mid {
//         t.Update(o << 1, l, mid, i, val)
//     } else {
//         t.Update(o << 1 | 1, mid + 1, r, i, val)
//     }
//     t[o] = max(t[o << 1], t[o << 1 | 1])
// }

// // 查询 [0,R] 中的最大值
// func (t Segment) Query(o, l, r, R int) int {
//     if r <= R { return t[o] }
//     mid := (l + r) >> 1
//     if R <= mid { return t.Query(o << 1, l, mid, R) }
//     return max(t[o << 1], t.Query(o << 1 | 1, mid + 1, r, R))
// }

// func getResults(queries [][]int) []bool {
//     mx, res := 0, []bool{}
//     for _, q := range queries {
//         mx = max(mx, q[1])
//     }
//     mx++
//     set := redblacktree.New[int, struct{}]()
//     set.Put(0, struct{}{}) // 哨兵
//     set.Put(mx, struct{}{})
//     t := make(seg, 2 << bits.Len(uint(mx)))
//     for _, q := range queries {
//         x := q[1]
//         pre, _ := set.Floor(x - 1) // x 左侧最近障碍物的位置
//         if q[0] == 1 {
//             nxt, _ := set.Ceiling(x) // x 右侧最近障碍物的位置
//             set.Put(x, struct{}{})
//             t.update(1, 0, mx, x, x-pre.Key)       // 更新 d[x] = x - pre
//             t.update(1, 0, mx, nxt.Key, nxt.Key-x) // 更新 d[nxt] = nxt - x
//         } else {
//             // 最大长度要么是 [0,pre] 中的最大 d，要么是 [pre,x] 这一段的长度
//             maxGap := max(t.query(1, 0, m, pre.Key), x-pre.Key)
//             res = append(res, maxGap >= q[2])
//         }
//     }
//     return res
// }

type Fenwick []int

func (f Fenwick) update(i, val int) {
    for ; i < len(f); i += i & -i {
        f[i] = max(f[i], val)
    }
}

func (f Fenwick) preMax(i int) (res int) {
    for ; i > 0; i &= i - 1 {
        res = max(res, f[i])
    }
    return res
}

type UnionFind []int

func (f UnionFind) find(x int) int {
    if f[x] != x {
        f[x] = f.find(f[x])
    }
    return f[x]
}

func getResults(queries [][]int) []bool {
    res, pos, m := []bool{}, []int{0}, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, q := range queries {
        m = max(m, q[1])
        if q[0] == 1 {
            pos = append(pos, q[1])
        }
    }
    m++
    left, right := make(UnionFind, m + 1), make(UnionFind, m + 1)
    for i := range left {
        left[i] = i
        right[i] = i
    }
    t := make(Fenwick, m)
    slices.Sort(pos)
    for i := 1; i < len(pos); i++ {
        p, q := pos[i-1], pos[i]
        t.update(q, q - p)
        for j := p + 1; j < q; j++ {
            left[j], right[j] = p, q // 删除 j
        }
    }
    for j := pos[len(pos)-1] + 1; j < m; j++ {
        left[j], right[j] = pos[len(pos)-1], m // 删除 j
    }
    for i := len(queries) - 1; i >= 0; i-- {
        q := queries[i]
        x := q[1]
        pre := left.find(x - 1) // x 左侧最近障碍物的位置
        if q[0] == 1 {
            left[x] = x - 1 // 删除 x
            right[x] = x + 1
            next := right.find(x)   // x 右侧最近障碍物的位置
            t.update(next, next - pre) // 更新 d[next] = next - pre
        } else {
            // 最大长度要么是 [0,pre] 中的最大 d，要么是 [pre,x] 这一段的长度
            maxGap := max(t.preMax(pre), x-pre)
            res = append(res, maxGap >= q[2])
        }
    }
    slices.Reverse(res)
    return res
}

func main() {
    // Example 1:
    // Input: queries = [[1,2],[2,3,3],[2,3,1],[2,2,2]]
    // Output: [false,true,true]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/04/22/example0block.png" />
    // For query 0, place an obstacle at x = 2. A block of size at most 2 can be placed before x = 3.
    fmt.Println(getResults([][]int{{1,2},{2,3,3},{2,3,1},{2,2,2}})) // [false,true,true]
    // Example 2:
    // Input: queries = [[1,7],[2,7,6],[1,2],[2,7,5],[2,7,6]]
    // Output: [true,true,false]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/04/22/example1block.png" />
    // Place an obstacle at x = 7 for query 0. A block of size at most 7 can be placed before x = 7.
    // Place an obstacle at x = 2 for query 2. Now, a block of size at most 5 can be placed before x = 7, and a block of size at most 2 before x = 2.
    fmt.Println(getResults([][]int{{1,7},{2,7,6},{1,2},{2,7,5},{2,7,6}})) // [true,true,false]
}