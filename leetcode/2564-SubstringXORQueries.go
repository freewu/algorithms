package main

// 2564. Substring XOR Queries
// You are given a binary string s, and a 2D integer array queries where queries[i] = [firsti, secondi].

// For the ith query, find the shortest substring of s whose decimal value, val, yields secondi when bitwise XORed with firsti. 
// In other words, val ^ firsti == secondi.

// The answer to the ith query is the endpoints (0-indexed) of the substring [lefti, righti] or [-1, -1] if no such substring exists. 
// If there are multiple answers, choose the one with the minimum lefti.

// Return an array ans where ans[i] = [lefti, righti] is the answer to the ith query.

// A substring is a contiguous non-empty sequence of characters within a string.

// Example 1:
// Input: s = "101101", queries = [[0,5],[1,2]]
// Output: [[0,2],[2,3]]
// Explanation: For the first query the substring in range [0,2] is "101" which has a decimal value of 5, and 5 ^ 0 = 5, hence the answer to the first query is [0,2]. In the second query, the substring in range [2,3] is "11", and has a decimal value of 3, and 3 ^ 1 = 2. So, [2,3] is returned for the second query. 

// Example 2:
// Input: s = "0101", queries = [[12,8]]
// Output: [[-1,-1]]
// Explanation: In this example there is no substring that answers the query, hence [-1,-1] is returned.

// Example 3:
// Input: s = "1", queries = [[4,5]]
// Output: [[0,0]]
// Explanation: For this example, the substring in range [0,0] has a decimal value of 1, and 1 ^ 4 = 5. So, the answer is [0,0].

// Constraints:
//     1 <= s.length <= 10^4
//     s[i] is either '0' or '1'.
//     1 <= queries.length <= 10^5
//     0 <= firsti, secondi <= 10^9

import "fmt"
import "strconv"
import "strings"

func substringXorQueries(s string, queries [][]int) [][]int {
    res := make([][]int, len(queries))
    for i, v := range queries {
        arr := make([]int, 2)
        xor := strconv.FormatInt(int64(v[0] ^ v[1]), 2)
        v := strings.Index(s, xor)
        if v >= 0 {
            arr[0], arr[1] = v, v + len(xor) - 1
        } else {
            arr[0], arr[1] = -1, -1
        }
        res[i] = arr
    }
    return res
}

func substringXorQueries1(s string, queries [][]int) [][]int {
    type Pair struct{ l, r int }
    mp := make(map[int]Pair)
    if i := strings.IndexByte(s, '0'); i >= 0 {
        mp[0] = Pair{i, i}
    }
    for l, c := range s {
        if c == '0' { continue }
        for r, x := l, 0; r < l + 30 && r < len(s); r++ {
            x = x << 1 | int(s[r] & 1)
            if _, ok := mp[x]; !ok {
                mp[x] = Pair{ l, r }
            }
        }
    }
    res := make([][]int, len(queries))
    for i, q := range queries {
        if v, ok := mp[q[0]^q[1]]; ok {
            res[i] = []int{ v.l, v.r }
        } else {
            res[i] = []int{-1, -1}
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "101101", queries = [[0,5],[1,2]]
    // Output: [[0,2],[2,3]]
    // Explanation: For the first query the substring in range [0,2] is "101" which has a decimal value of 5, and 5 ^ 0 = 5, hence the answer to the first query is [0,2]. In the second query, the substring in range [2,3] is "11", and has a decimal value of 3, and 3 ^ 1 = 2. So, [2,3] is returned for the second query. 
    fmt.Println(substringXorQueries("101101", [][]int{{0,5},{1,2}})) // [[0,2],[2,3]] 
    // Example 2:
    // Input: s = "0101", queries = [[12,8]]
    // Output: [[-1,-1]]
    // Explanation: In this example there is no substring that answers the query, hence [-1,-1] is returned.
    fmt.Println(substringXorQueries("0101", [][]int{{12,8}})) // [[-1,-1]]
    // Example 3:
    // Input: s = "1", queries = [[4,5]]
    // Output: [[0,0]]
    // Explanation: For this example, the substring in range [0,0] has a decimal value of 1, and 1 ^ 4 = 5. So, the answer is [0,0].
    fmt.Println(substringXorQueries("1", [][]int{{4,5}})) // [[0,0]]

    fmt.Println(substringXorQueries1("101101", [][]int{{0,5},{1,2}})) // [[0,2],[2,3]] 
    fmt.Println(substringXorQueries1("0101", [][]int{{12,8}})) // [[-1,-1]]
    fmt.Println(substringXorQueries1("1", [][]int{{4,5}})) // [[0,0]]
}