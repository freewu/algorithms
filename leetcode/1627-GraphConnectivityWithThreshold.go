package main

// 1627. Graph Connectivity With Threshold
// We have n cities labeled from 1 to n. 
// Two different cities with labels x and y are directly connected by a bidirectional road if and only if x and y share a common divisor strictly greater than some threshold. 
// More formally, cities with labels x and y have a road between them if there exists an integer z such that all of the following are true:
//     x % z == 0,
//     y % z == 0, and
//     z > threshold.

// Given the two integers, n and threshold, and an array of queries, 
// you must determine for each queries[i] = [ai, bi] if cities ai and bi are connected directly or indirectly. 
// (i.e. there is some path between them).

// Return an array answer, where answer.length == queries.length 
// and answer[i] is true if for the ith query, there is a path between ai and bi, 
// or answer[i] is false if there is no path.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/10/09/ex1.jpg" />
// Input: n = 6, threshold = 2, queries = [[1,4],[2,5],[3,6]]
// Output: [false,false,true]
// Explanation: The divisors for each number:
// 1:   1
// 2:   1, 2
// 3:   1, 3
// 4:   1, 2, 4
// 5:   1, 5
// 6:   1, 2, 3, 6
// Using the underlined divisors above the threshold, only cities 3 and 6 share a common divisor, so they are the
// only ones directly connected. The result of each query:
// [1,4]   1 is not connected to 4
// [2,5]   2 is not connected to 5
// [3,6]   3 is connected to 6 through path 3--6

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/10/10/tmp.jpg" />
// Input: n = 6, threshold = 0, queries = [[4,5],[3,4],[3,2],[2,6],[1,3]]
// Output: [true,true,true,true,true]
// Explanation: The divisors for each number are the same as the previous example. However, since the threshold is 0,
// all divisors can be used. Since all numbers share 1 as a divisor, all cities are connected.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2020/10/17/ex3.jpg" />
// Input: n = 5, threshold = 1, queries = [[4,5],[4,5],[3,2],[2,3],[3,4]]
// Output: [false,false,false,false,false]
// Explanation: Only cities 2 and 4 share a common divisor 2 which is strictly greater than the threshold 1, so they are the only ones directly connected.
// Please notice that there can be multiple queries for the same pair of nodes [x, y], and that the query [x, y] is equivalent to the query [y, x].

// Constraints:
//     2 <= n <= 10^4
//     0 <= threshold <= n
//     1 <= queries.length <= 10^5
//     queries[i].length == 2
//     1 <= ai, bi <= cities
//     ai != bi

import "fmt"

func areConnected(n int, threshold int, queries [][]int) []bool {
    mp, arr := make(map[int]int), make([]int, n - threshold)
    for i := 0; i < len(arr); i ++ {
        index := i + threshold + 1
        tag := len(mp) + 1
        mp[tag] = tag
        for j := i; j < len(arr); j += index {
            if arr[j] == 0 {
                arr[j] = tag
                continue
            }
            t := arr[j]
            for t != mp[t] {
                t = mp[t]
            }
            mp[tag] = t
            tag = t
        }
    }
    for i := 0; i < len(arr); i++ {
        t := arr[i]
        for t != mp[t] {
            t = mp[t]
        }
        arr[i] = t
    }
    res := make([]bool, len(queries))
    for i := 0; i < len(queries); i++ {
        ia, ib := queries[i][0] - 1 - threshold, queries[i][1] - 1 - threshold
        if ia < 0 || ib < 0 { continue }
        res[i] = (arr[ia] == arr[ib])
    }
    return res
}

type UnionFound struct {
    Parent, Rank []int
}

func newUnionFound(n int) *UnionFound {
    parent, rank := make([]int, n), make([]int, n)
    for i := range parent {
        parent[i], rank[i] = i, 1
    }
    return &UnionFound{parent, rank}
}

func (u *UnionFound) Find(x int) int {
    res := x
    for u.Parent[res] != res {
        res = u.Parent[res]
    }
    for u.Parent[x] != res {
        u.Parent[x], x = res, u.Parent[x]
    }
    return res
}

func (u *UnionFound) Union(from, to int) {
    ff, ft := u.Find(from), u.Find(to)
    if ff == ft { return }
    if u.Rank[ff] > u.Rank[ft] {
        ff, ft = ft, ff
    }
    u.Parent[ff] = ft
    u.Rank[ft] += u.Rank[ff]
}

func (u *UnionFound) Same(x, y int) bool {
    return u.Find(x) == u.Find(y)
}

func areConnected1(n int, threshold int, queries [][]int) []bool {
    uf := newUnionFound(n + 1)
    for i := threshold + 1; i <= n; i++ {
        for p, q := i, i*2; q <= n; p += i {
            uf.Union(p, q)
            q += i
        }
    }
    res := make([]bool, len(queries))
    for i, q := range queries {
        res[i] = uf.Same(q[0], q[1])
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/10/09/ex1.jpg" />
    // Input: n = 6, threshold = 2, queries = [[1,4],[2,5],[3,6]]
    // Output: [false,false,true]
    // Explanation: The divisors for each number:
    // 1:   1
    // 2:   1, 2
    // 3:   1, 3
    // 4:   1, 2, 4
    // 5:   1, 5
    // 6:   1, 2, 3, 6
    // Using the underlined divisors above the threshold, only cities 3 and 6 share a common divisor, so they are the
    // only ones directly connected. The result of each query:
    // [1,4]   1 is not connected to 4
    // [2,5]   2 is not connected to 5
    // [3,6]   3 is connected to 6 through path 3--6
    fmt.Println(areConnected(6, 2, [][]int{{1,4},{2,5},{3,6}})) // [false,false,true]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/10/10/tmp.jpg" />
    // Input: n = 6, threshold = 0, queries = [[4,5],[3,4],[3,2],[2,6],[1,3]]
    // Output: [true,true,true,true,true]
    // Explanation: The divisors for each number are the same as the previous example. However, since the threshold is 0,
    // all divisors can be used. Since all numbers share 1 as a divisor, all cities are connected.
    fmt.Println(areConnected(6, 0, [][]int{{4,5},{3,4},{3,2},{2,6},{1,3}})) // [true,true,true,true,true]
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2020/10/17/ex3.jpg" />
    // Input: n = 5, threshold = 1, queries = [[4,5],[4,5],[3,2],[2,3],[3,4]]
    // Output: [false,false,false,false,false]
    // Explanation: Only cities 2 and 4 share a common divisor 2 which is strictly greater than the threshold 1, so they are the only ones directly connected.
    // Please notice that there can be multiple queries for the same pair of nodes [x, y], and that the query [x, y] is equivalent to the query [y, x].
    fmt.Println(areConnected(5, 1, [][]int{{4,5},{4,5},{3,2},{2,3},{3,4}})) // [false,false,false,false,false]

    fmt.Println(areConnected1(6, 2, [][]int{{1,4},{2,5},{3,6}})) // [false,false,true]
    fmt.Println(areConnected1(6, 0, [][]int{{4,5},{3,4},{3,2},{2,6},{1,3}})) // [true,true,true,true,true]
    fmt.Println(areConnected1(5, 1, [][]int{{4,5},{4,5},{3,2},{2,3},{3,4}})) // [false,false,false,false,false]
}