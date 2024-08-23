package main

// LCR 111. 除法求值
// 给定一个变量对数组 equations 和一个实数值数组 values 作为已知条件，
// 其中 equations[i] = [Ai, Bi] 和 values[i] 共同表示等式 Ai / Bi = values[i] 。
// 每个 Ai 或 Bi 是一个表示单个变量的字符串。

// 另有一些以数组 queries 表示的问题，其中 queries[j] = [Cj, Dj] 表示第 j 个问题，
// 请你根据已知条件找出 Cj / Dj = ? 的结果作为答案。

// 返回 所有问题的答案 。如果存在某个无法确定的答案，则用 -1.0 替代这个答案。
// 如果问题中出现了给定的已知条件中没有出现的字符串，也需要用 -1.0 替代这个答案。

// 注意：输入总是有效的。可以假设除法运算中不会出现除数为 0 的情况，且不存在任何矛盾的结果。

// 示例 1：
// 输入：equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
// 输出：[6.00000,0.50000,-1.00000,1.00000,-1.00000]
// 解释：
// 条件：a / b = 2.0, b / c = 3.0
// 问题：a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
// 结果：[6.0, 0.5, -1.0, 1.0, -1.0 ]

// 示例 2：
// 输入：equations = [["a","b"],["b","c"],["bc","cd"]], values = [1.5,2.5,5.0], queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
// 输出：[3.75000,0.40000,5.00000,0.20000]

// 示例 3：
// 输入：equations = [["a","b"]], values = [0.5], queries = [["a","b"],["b","a"],["a","c"],["x","y"]]
// 输出：[0.50000,2.00000,-1.00000,-1.00000]

// 提示：
//     1 <= equations.length <= 20
//     equations[i].length == 2
//     1 <= Ai.length, Bi.length <= 5
//     values.length == equations.length
//     0.0 < values[i] <= 20.0
//     1 <= queries.length <= 20
//     queries[i].length == 2
//     1 <= Cj.length, Dj.length <= 5
//     Ai, Bi, Cj, Dj 由小写英文字母与数字组成

import "fmt"

// dfs
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    var dfs func(x, y string, visited map[string]bool, g map[string]map[string]float64) float64
    dfs = func(x, y string, visited map[string]bool, g map[string]map[string]float64) float64 {
        // 1. if x or y is not in the graph return -1
        _, okX := g[x]
        _, okY := g[y]
        if !okX || ! okY {
            return -1.0
        }
        // 2. if len of x edges is 0; retturn -1  
        if len(g[x]) == 0 {
            return -1.0
        }
        // 3. if y in x return the result
        val, okResult := g[x][y] 
        if okResult {
            return val
        }
        // 4. do dfs now, and multiply the product if answer is not -1 (which means it is not in the path of query)
        for k, _ := range g[x]{
            if !visited[k]{
                visited[k] = true
                tmp := dfs(k, y, visited, g)
                if tmp == -1.0 {
                    continue
                } else {
                    return tmp * g[x][k]
                }
            }
        }
        return -1.0
    }
    // 1. let's build the graph
    g := make(map[string]map[string]float64)
    for i, v := range equations {
        if g[v[0]] == nil {
            g[v[0]] = make(map[string]float64)
        }
        if g[v[1]] == nil {
            g[v[1]] = make(map[string]float64)
        }
        g[v[0]][v[1]] = values[i]
        g[v[1]][v[0]] = 1 / values[i]
    }
    results := make([]float64, 0, len(queries))
    // 2. answer queries
    for _, v := range queries {
        x := v[0]
        y := v[1]
        v := make(map[string]bool)
        result := dfs(x, y, v, g)
        results = append(results, result)
    }
    return results
}

func main() {
    // Example 1:
    // Input: equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
    // Output: [6.00000,0.50000,-1.00000,1.00000,-1.00000]
    // Explanation: 
    // Given: a / b = 2.0, b / c = 3.0
    // queries are: a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ? 
    // return: [6.0, 0.5, -1.0, 1.0, -1.0 ]
    // note: x is undefined => -1.0
    fmt.Println(calcEquation(
        [][]string{{"a","b"},{"b","c"}},
        []float64{2.0,3.0},
        [][]string{{"a","c"},{"b","a"},{"a","e"},{"a","a"},{"x","x"}},
    )) // [6.00000,0.50000,-1.00000,1.00000,-1.00000]
    // Example 2:
    // Input: equations = [["a","b"],["b","c"],["bc","cd"]], values = [1.5,2.5,5.0], queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
    // Output: [3.75000,0.40000,5.00000,0.20000]
    fmt.Println(calcEquation(
        [][]string{{"a","b"},{"b","c"},{"bc","cd"}},
        []float64{1.5,2.5,5.0}, 
        [][]string{{"a","c"},{"c","b"},{"bc","cd"},{"cd","bc"}},
    )) // [3.75000,0.40000,5.00000,0.20000] 
    // Example 3:
    // Input: equations = [["a","b"]], values = [0.5], queries = [["a","b"],["b","a"],["a","c"],["x","y"]]
    // Output: [0.50000,2.00000,-1.00000,-1.00000]
    fmt.Println(calcEquation(
        [][]string{{"a","b"}},
        []float64{0.5}, 
        [][]string{{"a","b"},{"b","a"},{"a","c"},{"x","y"}},
    )) // [0.50000,2.00000,-1.00000,-1.00000]
}