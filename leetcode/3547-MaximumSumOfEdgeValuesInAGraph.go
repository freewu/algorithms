package main

// 3547. Maximum Sum of Edge Values in a Graph
// You are given an undirected graph of n nodes, numbered from 0 to n - 1. 
// Each node is connected to at most 2 other nodes.

// The graph consists of m edges, represented by a 2D array edges, where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi.

// You have to assign a unique value from 1 to n to each node. 
// The value of an edge will be the product of the values assigned to the two nodes it connects.

// Your score is the sum of the values of all edges in the graph.

// Return the maximum score you can achieve.

// Example 1:
// <img scr="https://assets.leetcode.com/uploads/2025/03/23/graphproblemex1drawio.png" />
// Input: n = 7, edges = [[0,1],[1,2],[2,0],[3,4],[4,5],[5,6]]
// Output: 130
// Explanation:
// The diagram above illustrates an optimal assignment of values to nodes. 
// The sum of the values of the edges is: (7 * 6) + (7 * 5) + (6 * 5) + (1 * 3) + (3 * 4) + (4 * 2) = 130.

// Example 2:
// <img scr="https://assets.leetcode.com/uploads/2025/03/23/graphproblemex2drawio.png" />
// Input: n = 6, edges = [[0,3],[4,5],[2,0],[1,3],[2,4],[1,5]]
// Output: 82
// Explanation:
// The diagram above illustrates an optimal assignment of values to nodes. 
// The sum of the values of the edges is: (1 * 2) + (2 * 4) + (4 * 6) + (6 * 5) + (5 * 3) + (3 * 1) = 82.

// Constraints:
//     1 <= n <= 5 * 10^4
//     m == edges.length
//     1 <= m <= n
//     edges[i].length == 2
//     0 <= ai, bi < n
//     ai != bi
//     There are no repeated edges.
//     Each node is connected to at most 2 other nodes.

import "fmt"
import "sort"

func maxScore(n int, edges [][]int) int64 {
    path, cluster, cyclic := make([][]int, n), make([]int, n), make([]bool, 1)
    for i := 0; i < len(edges); i++ {
        from, to := edges[i][0], edges[i][1]
        path[from] = append(path[from], to)
        path[to] = append(path[to], from)
    }
    visit, local,localPath   := make([]bool, n), make([]bool, n), make(map[[2]int]bool)
    // topological sort
    // detect cycle
    // assign node to cluster
    var dfs func(curr int, c int) bool
    dfs = func(curr, c int) (cycle bool) {
        if visit[curr] { return false }
        if local[curr] { return true }
        local[curr] = true
        for i := 0; i < len(path[curr]); i++ {
            if localPath[[2]int{curr, path[curr][i]}] { continue }
            localPath[[2]int{curr, path[curr][i]}] = true
            localPath[[2]int{path[curr][i], curr}] = true
            cycle = dfs(path[curr][i], c) || cycle
        }
        local[curr] = false; visit[curr] = true
        cluster[curr] = c
        return cycle
    }
    res, count, mul := int64(0), 1, n
    for i := 0; i < n; i++ {
        if visit[i] { continue }
        c := dfs(i, count)
        cyclic = append(cyclic, c)
        count++
    }
    freq := make(map[int]int)
    for i := 0; i < len(cluster); i++ {
        freq[cluster[i]]++
    }
    sum := make([][2]int, 0)
    for k, v := range freq {
        sum = append(sum, [2]int{k, v})
    }
    // prioritize cyclic first 
    // prioritize cluster with larger member
    sort.Slice(sum, func(i, j int) bool {
        c1, c2 := cyclic[sum[i][0]], cyclic[sum[j][0]]
        if c1 && !c2 { return true }
        if c2 && !c1 { return false }
        return sum[i][1] > sum[j][1]
    })
    // assign the mul number to each node
    // based on pattern, from the midle to left and right differ 1
    // after it differ 2
    // additonal step if cyclic, multiplied most left & most right
    calc := func(size int, mul int, cyclic bool) {
        if size == 1 { return }
        if size == 2 {
            res += int64(mul * (mul - 1))
            return 
        }
        arr := make([]int, size)
        arr[size / 2], arr[size / 2 + 1], arr[size / 2 - 1] = mul,  mul - 2, mul - 1
        for i := 0; i < size/2-1; i++ {
            arr[size/2-2-i] = (mul - 3) - i * 2
        }
        for i := 0; i < size - (size / 2 + 1); i++ {
            arr[size/2+1+i] = (mul - 2) - i * 2
        }
        for i := 0; i < len(arr) - 1; i++ {
            res += int64(arr[i]*arr[i+1])
        }
        if cyclic {res += int64(arr[0]*arr[len(arr)-1])}
    }
    // traverse every cluster
    for i := 0; i < len(sum); i++ {
        calc(sum[i][1], mul, cyclic[sum[i][0]])
        mul -= sum[i][1]
    }
    return res
}

func maxScore1(n int, edges [][]int) int64 {
    f, en := make([]int, n), make([]int, n)
    for i := range f {
        f[i] = -1 
    }
    var fa func(int) int 
    fa = func(u int) int {
        if f[u] < 0 { return u }
        f[u] = fa(f[u])
        return f[u]
    }
    merge := func(u int, v int) {
        u, v = fa(u), fa(v)
        if f[u] > f[v] { u, v = v, u }
        en[u] ++ 
        if u == v { return }
        f[u] += f[v]
        en[u] += en[v]
        f[v] = u
        return 
    }
    for _, e := range edges {
        merge(e[0], e[1])
    }
    cir, chn := make([]int, 0), make([]int, 0)
    for i := range f {
        if f[i] >= 0  { continue }
        if f[i] == -1 { continue }
        if f[i] == -en[i] { 
            cir = append(cir, en[i]) 
        } else {
            chn = append(chn, en[i] + 1)
        }
    }
    sort.Ints(chn)
    sort.Ints(cir) 
    res, cur := 0, n 
    for i := range cir {
        mi := cur - cir[i] + 1
        for j := 0; j + 2 < cir[i]; j ++ {
            res += (mi + j) * (mi + j + 2)
        }
        res += cur * (cur - 1) + mi * (mi + 1)
        cur -= cir[i]
    }
    for i := len(chn) - 1; i >= 0; i -- {
        mi := cur - chn[i] + 1
        for j := 0; j + 2 < chn[i]; j ++ {
            res += (mi + j) * (mi + j + 2)
        }
        res += cur * (cur - 1) 
        cur -= chn[i]
    }
    return int64(res)
}

func main() {
    // Example 1:
    // <img scr="https://assets.leetcode.com/uploads/2025/03/23/graphproblemex1drawio.png" />
    // Input: n = 7, edges = [[0,1],[1,2],[2,0],[3,4],[4,5],[5,6]]
    // Output: 130
    // Explanation:
    // The diagram above illustrates an optimal assignment of values to nodes. 
    // The sum of the values of the edges is: (7 * 6) + (7 * 5) + (6 * 5) + (1 * 3) + (3 * 4) + (4 * 2) = 130.
    fmt.Println(maxScore(7, [][]int{{0,1},{1,2},{2,0},{3,4},{4,5},{5,6}})) // 130
    // Example 2:
    // <img scr="https://assets.leetcode.com/uploads/2025/03/23/graphproblemex2drawio.png" />
    // Input: n = 6, edges = [[0,3],[4,5],[2,0],[1,3],[2,4],[1,5]]
    // Output: 82
    // Explanation:
    // The diagram above illustrates an optimal assignment of values to nodes. 
    // The sum of the values of the edges is: (1 * 2) + (2 * 4) + (4 * 6) + (6 * 5) + (5 * 3) + (3 * 1) = 82.
    fmt.Println(maxScore(6, [][]int{{0,3},{4,5},{2,0},{1,3},{2,4},{1,5}})) // 82

    fmt.Println(maxScore(11, [][]int{{0,1},{1,2},{2,3},{5,6},{6,7}})) // 366

    fmt.Println(maxScore1(7, [][]int{{0,1},{1,2},{2,0},{3,4},{4,5},{5,6}})) // 130
    fmt.Println(maxScore1(6, [][]int{{0,3},{4,5},{2,0},{1,3},{2,4},{1,5}})) // 82
    fmt.Println(maxScore1(11, [][]int{{0,1},{1,2},{2,3},{5,6},{6,7}})) // 366
}