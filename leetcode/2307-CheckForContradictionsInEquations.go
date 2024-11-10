package main

// 2307. Check for Contradictions in Equations
// You are given a 2D array of strings equations and an array of real numbers values, 
// where equations[i] = [Ai, Bi] and values[i] means that Ai / Bi = values[i].

// Determine if there exists a contradiction in the equations. 
// Return true if there is a contradiction, or false otherwise.

// Note:
//     1. When checking if two numbers are equal, check that their absolute difference is less than 10^-5.
//     2. The testcases are generated such that there are no cases targeting precision, 
//        i.e. using double is enough to solve the problem.

// Example 1:
// Input: equations = [["a","b"],["b","c"],["a","c"]], values = [3,0.5,1.5]
// Output: false
// Explanation:
// The given equations are: a / b = 3, b / c = 0.5, a / c = 1.5
// There are no contradictions in the equations. One possible assignment to satisfy all equations is:
// a = 3, b = 1 and c = 2.

// Example 2:
// Input: equations = [["le","et"],["le","code"],["code","et"]], values = [2,5,0.5]
// Output: true
// Explanation:
// The given equations are: le / et = 2, le / code = 5, code / et = 0.5
// Based on the first two equations, we get code / et = 0.4.
// Since the third equation is code / et = 0.5, we get a contradiction.

// Constraints:
//     1 <= equations.length <= 100
//     equations[i].length == 2
//     1 <= Ai.length, Bi.length <= 5
//     Ai, Bi consist of lowercase English letters.
//     equations.length == values.length
//     0.0 < values[i] <= 10.0
//     values[i] has a maximum of 2 decimal places.

import "fmt"

// // 带权值的并查集
// func checkContradictions(equations [][]string, values []float64) bool {
//     n, count, mp := len(equations), 0, make(map[string]int)
//     for _, v := range equations { // 离散化
//         if _, ok := mp[v[0]]; !ok {
//             count++
//             mp[v[0]] = count
//         }
//         if _, ok := mp[v[1]]; !ok {
//             count++
//             mp[v[1]] = count
//         }
//     }
//     parents, weights := make([]int, count + 1),  make([]float64, count + 1)
//     fmt.Println(count)
//     for i := 0; i < count; i++ {
//         parents[i], weights[i] = i, 1.0
//     }
//     var find func(x int) int
//     find = func(x int) int {
//         if parents[x] != x {
//             f := find(parents[x])
//             weights[x] *= weights[parents[x]]
//             parents[x] = f
//         }
//         return parents[x]
//     }
//     union := func(x, y int, w float64) {
//         xroot, yroot := find(x), find(y)
//         if xroot != yroot {
//             parents[xroot] = yroot
//             weights[xroot] = w * weights[y] / weights[x]
//         }
//     }
//     query := func(x, y int) float64 {
//         return weights[x] / weights[y]
//     }
//     abs := func(x float64) float64 { if x < 0 { return -x; }; return x; }
//     for i := 0; i < n; i++ {
//         q := equations[i]
//         if find(mp[q[0]]) != find(mp[q[1]]) {
//             union(mp[q[0]], mp[q[1]], values[i])
//         } else {
//             if abs(query(mp[q[0]], mp[q[1]]) - values[i]) >= 1e-5 { return true }
//         }
//     }
//     return false
// }

// dfs
func checkContradictions(equations [][]string, values []float64) bool {
    n, mp := len(equations), make(map[string]map[string]float64)
    abs := func(x float64) float64 { if x < 0 { return -x; }; return x; }
    var dfs func(start string, end string, visited map[string]bool) float64
    dfs = func(start string, end string, visited map[string]bool) float64 {
        if mp[start] == nil { return -1 }
        if start == end { return 1.0 }
        visited[start] = true
        for k, v := range mp[start] {
            if !visited[k] {
                subResult := dfs(k, end, visited)
                if subResult != -1 { return subResult * v }
            }
        }
        delete(visited, start)
        return -1
    }
    for i := 0; i < n; i++ {
        node0, node1, value := equations[i][0], equations[i][1], values[i]
        if node0 == node1 && abs(value - 1.0) > 1e-5 { return true } // 自身矛盾
        existingValue := dfs(node0, node1, make(map[string]bool))
        if (existingValue != -1 && abs(existingValue - value) > 1e-5) { return true } // 与现有路径矛盾
        if mp[node0] == nil {
            mp[node0] = make(map[string]float64)
        }
        mp[node0][node1] = value
        if mp[node1] == nil {
            mp[node1] = make(map[string]float64)
        }
        mp[node1][node0] = 1.0 / value
    }
    return false
}

// 带权并查集
func checkContradictions1(equations [][]string, values []float64) bool {
    const loLimit, hiLimit = -1e-5, 1e-5
    ids := map[string]int{}
    for _, equation := range equations { // 给每一个等式的变量分配一个唯一id
        u, v := equation[0], equation[1]
        if _, ok := ids[u]; !ok {
            ids[u] = len(ids) // 使用当前map的大小,恰好可以当一个新的id
        }
        if _, ok := ids[v]; !ok {
            ids[v] = len(ids)
        }
    }
    // 带权并查集模板
    n := len(ids)
    fa := make([]int, n)
    wt := make([]float64, n)
    for i := range fa {
        fa[i] = i
        wt[i] = 1
    }
    var find func(int) int
    find = func(x int) int {
        if fa[x] != x {
            rt := find(fa[x]) // 调用结束后,x的父节点fa[x]已经修正为和rt的关系,那么{x和fa[x]的倍数 * fa[x]和rt的倍数},就是 x和rt的倍数.
            wt[x] *= wt[fa[x]]
            fa[x] = rt
        }
        return fa[x]
    }
    merge := func(from, to int, val float64) {
        rtFrom, rtTo := find(from), find(to)
        if rtFrom != rtTo {
            wt[rtFrom] = val * wt[to] / wt[from] // rtFrom是from的 1/wt[from]倍, from是to的vals倍, to是rtTo的 wt[to]倍,连乘即可
            fa[rtFrom] = rtTo
        }
    }
    for i, equation := range equations {
        u, v, val := equation[0], equation[1], values[i]
        idU, idV := ids[u], ids[v]
        if find(idU) != find(idV) { // 特别注意!! 经过find后, 两者的wt都修改为和root的关系了(如果不主动调用,是没有的)
            merge(idU, idV, val)
        } else {
            w := wt[idU] / wt[idV]
            if delta := w - val; delta < loLimit || delta > hiLimit { // 如果已经在一个set了,那么权重应该一样,
                return true
            }
        }
    }
    return false
}

func main() {
    // Example 1:
    // Input: equations = [["a","b"],["b","c"],["a","c"]], values = [3,0.5,1.5]
    // Output: false
    // Explanation:
    // The given equations are: a / b = 3, b / c = 0.5, a / c = 1.5
    // There are no contradictions in the equations. One possible assignment to satisfy all equations is:
    // a = 3, b = 1 and c = 2.
    fmt.Println(checkContradictions([][]string{{"a","b"},{"b","c"},{"a","c"}}, []float64{3,0.5,1.5})) // false
    // Example 2:
    // Input: equations = [["le","et"],["le","code"],["code","et"]], values = [2,5,0.5]
    // Output: true
    // Explanation:
    // The given equations are: le / et = 2, le / code = 5, code / et = 0.5
    // Based on the first two equations, we get code / et = 0.4.
    // Since the third equation is code / et = 0.5, we get a contradiction.
    fmt.Println(checkContradictions([][]string{{"le","et"},{"le","code"},{"code","et"}}, []float64{2,5,0.5})) // true

    fmt.Println(checkContradictions1([][]string{{"a","b"},{"b","c"},{"a","c"}}, []float64{3,0.5,1.5})) // false
    fmt.Println(checkContradictions1([][]string{{"le","et"},{"le","code"},{"code","et"}}, []float64{2,5,0.5})) // true
}