package main

// 3535. Unit Conversion II
// There are n types of units indexed from 0 to n - 1.

// You are given a 2D integer array conversions of length n - 1, where conversions[i] = [sourceUniti, targetUniti, conversionFactori]. 
// This indicates that a single unit of type sourceUniti is equivalent to conversionFactori units of type targetUniti.

// You are also given a 2D integer array queries of length q, where queries[i] = [unitAi, unitBi].

// Return an array answer of length q where answer[i] is the number of units of type unitBi equivalent to 1 unit of type unitAi. 
// Return each answer[i] as pq-1 modulo 10^9 + 7, where q-1 represents the multiplicative inverse of q modulo 10^9 + 7.

// Example 1:
// Input: conversions = [[0,1,2],[0,2,6]], queries = [[1,2],[1,0]]
// Output: [3,500000004]
// Explanation:
// In the first query, we can convert unit 1 into 3 units of type 2 using the inverse of conversions[0], then conversions[1].
// In the second query, we can convert unit 1 into 1/2 units of type 0 using the inverse of conversions[0]. We return 500000004 since it is the multiplicative inverse of 2.
// <img src="https://assets.leetcode.com/uploads/2025/03/13/example1.png" />

// Example 2:
// Input: conversions = [[0,1,2],[0,2,6],[0,3,8],[2,4,2],[2,5,4],[3,6,3]], queries = [[1,2],[0,4],[6,5],[4,6],[6,1]]
// Output: [3,12,1,2,83333334]
// Explanation:
// In the first query, we can convert unit 1 into 3 units of type 2 using the inverse of conversions[0], then conversions[1].
// In the second query, we can convert unit 0 into 12 units of type 4 using conversions[1], then conversions[3].
// In the third query, we can convert unit 6 into 1 unit of type 5 using the inverse of conversions[5], the inverse of conversions[2], conversions[1], then conversions[4].
// In the fourth query, we can convert unit 4 into 2 units of type 6 using the inverse of conversions[3], the inverse of conversions[1], conversions[2], then conversions[5].
// In the fifth query, we can convert unit 6 into 1/12 units of type 1 using the inverse of conversions[5], the inverse of conversions[2], then conversions[0]. We return 83333334 since it is the multiplicative inverse of 12.
// <img src="https://assets.leetcode.com/uploads/2025/03/13/example2.png" />

// Constraints:
//     2 <= n <= 10^5
//     conversions.length == n - 1
//     0 <= sourceUniti, targetUniti < n
//     1 <= conversionFactori <= 10^9
//     1 <= q <= 10^5
//     queries.length == q
//     0 <= unitAi, unitBi < n
//     It is guaranteed that unit 0 can be uniquely converted into any other unit through a combination of forward or backward conversions.

import "fmt"

func queryConversions(conversions [][]int, queries [][]int) []int {
    n, mod := len(conversions), 1_000_000_007
    graph := make([][][]int, n + 1)
    powMod := func(base, exp, mod int) int { // 快速幂取模
        res := 1
        for exp > 0 {
            if exp%2 == 1 {
                res = (res * base) % mod
            }
            base = (base * base) % mod
            exp /= 2
        }
        return res
    }
    modInverse := func(a, m int) int { return powMod(a, m-2, m) } // 计算模逆元
    for _, v := range conversions {
        x, y, w := v[0], v[1], v[2]
        graph[x] = append(graph[x], []int{ y, w })
        graph[y] = append(graph[y], []int{ x, modInverse(w, mod) })
    }
    arr := make([]int, n + 1)
    arr[0] = 1
    var dfs func(x, fa int)
    dfs = func(x, fa int) {
        for _, edge := range graph[x] {
            y, w := edge[0], edge[1]
            if y == fa { continue }
            arr[y] = (arr[x] * w) % mod
            dfs(y, x)
        }
    }
    dfs(0, -1)
    res := make([]int, len(queries))
    for i, q := range queries {
        x, y := q[0], q[1]
        res[i] = (arr[y] * modInverse(arr[x], mod)) % mod
    }
    return res
}

func main() {
    // Example 1:
    // Input: conversions = [[0,1,2],[0,2,6]], queries = [[1,2],[1,0]]
    // Output: [3,500000004]
    // Explanation:
    // In the first query, we can convert unit 1 into 3 units of type 2 using the inverse of conversions[0], then conversions[1].
    // In the second query, we can convert unit 1 into 1/2 units of type 0 using the inverse of conversions[0]. We return 500000004 since it is the multiplicative inverse of 2.
    // <img src="https://assets.leetcode.com/uploads/2025/03/13/example1.png" />
    fmt.Println(queryConversions([][]int{{0,1,2},{0,2,6}}, [][]int{{1,2},{1,0}})) // [3,500000004] 
    // Example 2:
    // Input: conversions = [[0,1,2],[0,2,6],[0,3,8],[2,4,2],[2,5,4],[3,6,3]], queries = [[1,2],[0,4],[6,5],[4,6],[6,1]]
    // Output: [3,12,1,2,83333334]
    // Explanation:
    // In the first query, we can convert unit 1 into 3 units of type 2 using the inverse of conversions[0], then conversions[1].
    // In the second query, we can convert unit 0 into 12 units of type 4 using conversions[1], then conversions[3].
    // In the third query, we can convert unit 6 into 1 unit of type 5 using the inverse of conversions[5], the inverse of conversions[2], conversions[1], then conversions[4].
    // In the fourth query, we can convert unit 4 into 2 units of type 6 using the inverse of conversions[3], the inverse of conversions[1], conversions[2], then conversions[5].
    // In the fifth query, we can convert unit 6 into 1/12 units of type 1 using the inverse of conversions[5], the inverse of conversions[2], then conversions[0]. We return 83333334 since it is the multiplicative inverse of 12.
    // <img src="https://assets.leetcode.com/uploads/2025/03/13/example2.png" />
    fmt.Println(queryConversions([][]int{{0,1,2},{0,2,6},{0,3,8},{2,4,2},{2,5,4},{3,6,3}}, [][]int{{1,2},{0,4},{6,5},{4,6},{6,1}})) // [3,12,1,2,83333334]
}

