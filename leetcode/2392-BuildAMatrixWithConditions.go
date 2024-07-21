package main

// 2392. Build a Matrix With Conditions
// You are given a positive integer k. You are also given:
//     a 2D integer array rowConditions of size n where rowConditions[i] = [abovei, belowi], and
//     a 2D integer array colConditions of size m where colConditions[i] = [lefti, righti].

// The two arrays contain integers from 1 to k.

// You have to build a k x k matrix that contains each of the numbers from 1 to k exactly once. 
// The remaining cells should have the value 0.

// The matrix should also satisfy the following conditions:
//     1. The number abovei should appear in a row that is strictly above the row at which the number belowi appears for all i from 0 to n - 1.
//     2. The number lefti should appear in a column that is strictly left of the column at which the number righti appears for all i from 0 to m - 1.

// Return any matrix that satisfies the conditions. If no answer exists, return an empty matrix.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/07/06/gridosdrawio.png" />
// Input: k = 3, rowConditions = [[1,2],[3,2]], colConditions = [[2,1],[3,2]]
// Output: [[3,0,0],[0,0,1],[0,2,0]]
// Explanation: The diagram above shows a valid example of a matrix that satisfies all the conditions.
// The row conditions are the following:
// - Number 1 is in row 1, and number 2 is in row 2, so 1 is above 2 in the matrix.
// - Number 3 is in row 0, and number 2 is in row 2, so 3 is above 2 in the matrix.
// The column conditions are the following:
// - Number 2 is in column 1, and number 1 is in column 2, so 2 is left of 1 in the matrix.
// - Number 3 is in column 0, and number 2 is in column 1, so 3 is left of 2 in the matrix.
// Note that there may be multiple correct answers.

// Example 2:
// Input: k = 3, rowConditions = [[1,2],[2,3],[3,1],[2,3]], colConditions = [[2,1]]
// Output: []
// Explanation: From the first two conditions, 3 has to be below 1 but the third conditions needs 3 to be above 1 to be satisfied.
// No matrix can satisfy all the conditions, so we return the empty matrix.

// Constraints:
//     2 <= k <= 400
//     1 <= rowConditions.length, colConditions.length <= 10^4
//     rowConditions[i].length == colConditions[i].length == 2
//     1 <= abovei, belowi, lefti, righti <= k
//     abovei != belowi
//     lefti != righti

import "fmt"

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
    getTopSort := func(k int, conditions [][]int) []int {
        res, graph, ind := make([]int, 0), make([][]int, k + 1),make([]int, k + 1)
        for _, cv := range conditions {
            ind[cv[1]]++
            graph[cv[0]] = append(graph[cv[0]], cv[1])
        }
        queue := make([]int, 0)
        for i, iv := range ind {
            if iv == 0 && i > 0{
                queue = append(queue, i)
            }
        }
        for len(queue) > 0 {
            cur := queue[0]
            queue = queue[1:]
            res = append(res, cur)
            for _, nextNode := range graph[cur] {
                ind[nextNode]--
                if ind[nextNode] == 0 {
                    queue = append(queue, nextNode)
                }
            }
        }
        return res
    }
    rsort, csort := getTopSort(k, rowConditions), getTopSort(k, colConditions)
    if len(rsort) < k || len(csort) < k {
        return [][]int{}
    }
    res, rowidx, colidx := make([][]int, k), make(map[int]int, 0), make(map[int]int, 0)
    for i := 0; i < k; i++ {
        res[i] = make([]int, k)
    }
    for i := 0; i < len(rsort); i++ {
        rowidx[rsort[i]] = i
        colidx[csort[i]] = i
    }
    for i := 1; i <= k; i++ {
        res[rowidx[i]][colidx[i]] = i
    }
    return res
}

func buildMatrix1(k int, rowConditions, colConditions [][]int) [][]int {
    topoSort := func(k int, edges [][]int) []int {
        g, inDeg := make([][]int, k), make([]int, k)
        for _, e := range edges {
            x, y := e[0]-1, e[1]-1 // 顶点编号从 0 开始，方便计算
            g[x] = append(g[x], y)
            inDeg[y]++
        }
        queue := make([]int, 0, k)
        orders := queue // 复用队列作为拓扑序
        for i, d := range inDeg {
            if d == 0 {
                queue = append(queue, i)
            }
        }
        for len(queue) > 0 {
            x := queue[0]
            queue = queue[1:]
            for _, y := range g[x] {
                if inDeg[y]--; inDeg[y] == 0 {
                    queue = append(queue, y)
                }
            }
        }
        if cap(queue) > 0 {
            return nil
        }
        return orders[:k]
    }
    row, col:= topoSort(k, rowConditions), topoSort(k, colConditions)
    if row == nil || col == nil {
        return nil
    }
    pos := make([]int, k)
    for i, v := range col {
        pos[v] = i
    }
    res := make([][]int, k)
    for i, x := range row {
        res[i] = make([]int, k)
        res[i][pos[x]] = x + 1
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/07/06/gridosdrawio.png" />
    // Input: k = 3, rowConditions = [[1,2],[3,2]], colConditions = [[2,1],[3,2]]
    // Output: [[3,0,0],[0,0,1],[0,2,0]]
    // Explanation: The diagram above shows a valid example of a matrix that satisfies all the conditions.
    // The row conditions are the following:
    // - Number 1 is in row 1, and number 2 is in row 2, so 1 is above 2 in the matrix.
    // - Number 3 is in row 0, and number 2 is in row 2, so 3 is above 2 in the matrix.
    // The column conditions are the following:
    // - Number 2 is in column 1, and number 1 is in column 2, so 2 is left of 1 in the matrix.
    // - Number 3 is in column 0, and number 2 is in column 1, so 3 is left of 2 in the matrix.
    // Note that there may be multiple correct answers.
    fmt.Println(buildMatrix(3,[][]int{{1,2},{3,2}}, [][]int{{2,1},{3,2}})) // [[3,0,0],[0,0,1],[0,2,0]]
    // Example 2:
    // Input: k = 3, rowConditions = [[1,2],[2,3],[3,1],[2,3]], colConditions = [[2,1]]
    // Output: []
    // Explanation: From the first two conditions, 3 has to be below 1 but the third conditions needs 3 to be above 1 to be satisfied.
    // No matrix can satisfy all the conditions, so we return the empty matrix.
    fmt.Println(buildMatrix(3,[][]int{{1,2},{2,3},{3,1},{2,3}}, [][]int{{2,1}})) // []

    fmt.Println(buildMatrix1(3,[][]int{{1,2},{3,2}}, [][]int{{2,1},{3,2}})) // [[3,0,0],[0,0,1],[0,2,0]]
    fmt.Println(buildMatrix1(3,[][]int{{1,2},{2,3},{3,1},{2,3}}, [][]int{{2,1}})) // []
}