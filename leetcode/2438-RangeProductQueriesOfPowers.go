package main

// 2438. Range Product Queries of Powers
// Given a positive integer n, there exists a 0-indexed array called powers, composed of the minimum number of powers of 2 that sum to n. 
// The array is sorted in non-decreasing order, and there is only one way to form the array.

// You are also given a 0-indexed 2D integer array queries, where queries[i] = [lefti, righti]. 
// Each queries[i] represents a query where you have to find the product of all powers[j] with lefti <= j <= righti.

// Return an array answers, equal in length to queries, where answers[i] is the answer to the ith query. 
// Since the answer to the ith query may be too large, each answers[i] should be returned modulo 10^9 + 7.

// Example 1:
// Input: n = 15, queries = [[0,1],[2,2],[0,3]]
// Output: [2,4,64]
// Explanation:
// For n = 15, powers = [1,2,4,8]. It can be shown that powers cannot be a smaller size.
// Answer to 1st query: powers[0] * powers[1] = 1 * 2 = 2.
// Answer to 2nd query: powers[2] = 4.
// Answer to 3rd query: powers[0] * powers[1] * powers[2] * powers[3] = 1 * 2 * 4 * 8 = 64.
// Each answer modulo 109 + 7 yields the same answer, so [2,4,64] is returned.

// Example 2:
// Input: n = 2, queries = [[0,0]]
// Output: [2]
// Explanation:
// For n = 2, powers = [2].
// The answer to the only query is powers[0] = 2. The answer modulo 109 + 7 is the same, so [2] is returned.

// Constraints:
//     1 <= n <= 10^9
//     1 <= queries.length <= 10^5
//     0 <= starti <= endi < powers.length

import "fmt"
import "math"

func productQueries(n int, queries [][]int) []int {
    // The powers array can be created using the binary representation of n.
    powers, x := []int{}, float64(0)
    for n > 0 {
        if n & 1 == 1 {
            powers = append(powers, int(math.Pow(2, x)))
        }
        x++
        n >>= 1
    }
    res := []int{}
    for _, q := range queries {
        l, r := q[0], q[1]
        val := powers[r]
        for i := l; i < r; i++ {
            val = (val * powers[i]) % 1_000_000_007
        }
        res = append(res, val)
    }
    return res
}

func productQueries1(n int, queries [][]int) []int {
    arr := []int{}
    for n > 0{
        v := n& -n
        arr = append(arr, v)
        n ^= v
    }
    m := len(arr)
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, m)
        dp[i][i] = arr[i]
        for j := i + 1; j < m; j++ {
            dp[i][j] = (dp[i][j-1] * arr[j]) % 1_000_000_007
        }
    }
    res := make([]int, len(queries))
    for i, q := range queries {
        res[i] = dp[q[0]][q[1]]
    }
    return res
}

func productQueries2(n int, queries [][]int) []int {
    const mod = 1_000_000_007
    rep, bins := 1, []int{}
    for n > 0 {
        if n % 2 == 1 {
            bins = append(bins, rep)
        }
        n /= 2
        rep *= 2
    }
    m := len(bins)
    arr := make([][]int, m)
    for i := range arr {
        arr[i] = make([]int, m)
        cur := 1
        for j := i; j < m; j++ {
            cur = (cur * bins[j]) % mod
            arr[i][j] = cur
        }
    }
    res := make([]int, len(queries))
    for i, query := range queries {
        res[i] = arr[query[0]][query[1]]
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 15, queries = [[0,1],[2,2],[0,3]]
    // Output: [2,4,64]
    // Explanation:
    // For n = 15, powers = [1,2,4,8]. It can be shown that powers cannot be a smaller size.
    // Answer to 1st query: powers[0] * powers[1] = 1 * 2 = 2.
    // Answer to 2nd query: powers[2] = 4.
    // Answer to 3rd query: powers[0] * powers[1] * powers[2] * powers[3] = 1 * 2 * 4 * 8 = 64.
    // Each answer modulo 109 + 7 yields the same answer, so [2,4,64] is returned.
    fmt.Println(productQueries(15,[][]int{{0,1},{2,2},{0,3}})) // [2,4,64]
    // Example 2:
    // Input: n = 2, queries = [[0,0]]
    // Output: [2]
    // Explanation:
    // For n = 2, powers = [2].
    // The answer to the only query is powers[0] = 2. The answer modulo 109 + 7 is the same, so [2] is returned.
    fmt.Println(productQueries(2,[][]int{{0,0}})) // [2]

    fmt.Println(productQueries1(15,[][]int{{0,1},{2,2},{0,3}})) // [2,4,64]
    fmt.Println(productQueries1(2,[][]int{{0,0}})) // [2]

    fmt.Println(productQueries2(15,[][]int{{0,1},{2,2},{0,3}})) // [2,4,64]
    fmt.Println(productQueries2(2,[][]int{{0,0}})) // [2]
}