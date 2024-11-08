package main

// 1782. Count Pairs Of Nodes
// You are given an undirected graph defined by an integer n, the number of nodes, and a 2D integer array edges, 
// the edges in the graph, where edges[i] = [ui, vi] indicates that there is an undirected edge between ui and vi. 
// You are also given an integer array queries.

// Let incident(a, b) be defined as the number of edges that are connected to either node a or b.

// The answer to the jth query is the number of pairs of nodes (a, b) that satisfy both of the following conditions:
//     1. a < b
//     2. incident(a, b) > queries[j]

// Return an array answers such that answers.length == queries.length and answers[j] is the answer of the jth query.

// Note that there can be multiple edges between the same two nodes.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/06/08/winword_2021-06-08_00-58-39.png" />
// Input: n = 4, edges = [[1,2],[2,4],[1,3],[2,3],[2,1]], queries = [2,3]
// Output: [6,5]
// Explanation: The calculations for incident(a, b) are shown in the table above.
// The answers for each of the queries are as follows:
// - answers[0] = 6. All the pairs have an incident(a, b) value greater than 2.
// - answers[1] = 5. All the pairs except (3, 4) have an incident(a, b) value greater than 3.

// Example 2:
// Input: n = 5, edges = [[1,5],[1,5],[3,4],[2,5],[1,3],[5,1],[2,3],[2,5]], queries = [1,2,3,4,5]
// Output: [10,10,9,8,6]

// Constraints:
//     2 <= n <= 2 * 10^4
//     1 <= edges.length <= 10^5
//     1 <= ui, vi <= n
//     ui != vi
//     1 <= queries.length <= 20
//     0 <= queries[j] < edges.length

import "fmt"
import "sort"

func countPairs(n int, edges [][]int, queries []int) []int {
    count, sorted, shared := make([]int, n + 1),  make([]int, n + 1), make([]map[int]int, n + 1)
    for i := range shared {
        shared[i] = make(map[int]int)
    }
    for _, v := range edges {
        small, large := v[0], v[1]
        if small > large {
            small, large = large, small
        }
        count[small]++
        count[large]++
        shared[small][large] += 1
    }
    copy(sorted, count)
    sort.Ints(sorted)
    res := make([]int, len(queries))
    for k, q := range queries {
        for i, j := 1, n; i < j; {
            if sorted[i] + sorted[j] > q {
                res[k] += (j - i)
                j--
            } else {
                i++
            }
        }
        for i := 1; i <= n; i++ {
            for j, v := range shared[i] {
                if count[i] + count[j] > q && count[i] + count[j] - v <= q {
                    res[k]--
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/06/08/winword_2021-06-08_00-58-39.png" />
    // Input: n = 4, edges = [[1,2],[2,4],[1,3],[2,3],[2,1]], queries = [2,3]
    // Output: [6,5]
    // Explanation: The calculations for incident(a, b) are shown in the table above.
    // The answers for each of the queries are as follows:
    // - answers[0] = 6. All the pairs have an incident(a, b) value greater than 2.
    // - answers[1] = 5. All the pairs except (3, 4) have an incident(a, b) value greater than 3.
    fmt.Println(countPairs(4, [][]int{{1,2},{2,4},{1,3},{2,3},{2,1}}, []int{2,3})) // [6,5]
    // Example 2:
    // Input: n = 5, edges = [[1,5],[1,5],[3,4],[2,5],[1,3],[5,1],[2,3],[2,5]], queries = [1,2,3,4,5]
    // Output: [10,10,9,8,6]
    fmt.Println(countPairs(5, [][]int{{1,5},{1,5},{3,4},{2,5},{1,3},{5,1},{2,3},{2,5}}, []int{1,2,3,4,5})) // [10,10,9,8,6]
}