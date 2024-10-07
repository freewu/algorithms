package main

// 2791. Count Paths That Can Form a Palindrome in a Tree
// You are given a tree (i.e. a connected, undirected graph that has no cycles) rooted at node 0 consisting of n nodes numbered from 0 to n - 1. 
// The tree is represented by a 0-indexed array parent of size n, where parent[i] is the parent of node i. 
// Since node 0 is the root, parent[0] == -1.

// You are also given a string s of length n, where s[i] is the character assigned to the edge between i and parent[i]. 
// s[0] can be ignored.

// Return the number of pairs of nodes (u, v) such that u < v 
// and the characters assigned to edges on the path from u to v can be rearranged to form a palindrome.

// A string is a palindrome when it reads the same backwards as forwards.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/07/15/treedrawio-8drawio.png" />
// Input: parent = [-1,0,0,1,1,2], s = "acaabc"
// Output: 8
// Explanation: The valid pairs are:
// - All the pairs (0,1), (0,2), (1,3), (1,4) and (2,5) result in one character which is always a palindrome.
// - The pair (2,3) result in the string "aca" which is a palindrome.
// - The pair (1,5) result in the string "cac" which is a palindrome.
// - The pair (3,5) result in the string "acac" which can be rearranged into the palindrome "acca".

// Example 2:
// Input: parent = [-1,0,0,0,0], s = "aaaaa"
// Output: 10
// Explanation: Any pair of nodes (u,v) where u < v is valid.
 
// Constraints:
//     n == parent.length == s.length
//     1 <= n <= 10^5
//     0 <= parent[i] <= n - 1 for all i >= 1
//     parent[0] == -1
//     parent represents a valid tree.
//     s consists of only lowercase English letters.

import "fmt"

func countPalindromePaths(parent []int, s string) int64 {
    n := len(s)
    con := make([][]int, n)
    for i := 0; i < n; i++ {
        con[i] = []int{}
    }
    for i := 1; i < n; i++ {
        con[parent[i]] = append(con[parent[i]], i)
    }
    have := make(map[int]int)
    have[0] = 1
    var dfs func(x, mask int) int64
    dfs = func(x, mask int) int64 {
        res := int64(0)
        if x != 0 {
            mask ^= 1 << (s[x] - 'a')
            for i := 1 << 25; i > 0; i >>= 1 {
                if val, ok := have[mask^i]; ok {
                    res += int64(val)
                }
            }
            res += int64(have[mask])
            have[mask]++
        }
        for _, y := range con[x] {
            res += dfs(y, mask)
        }
        return res
    }
    return dfs(0, 0)
}

func countPalindromePaths1(parent []int, s string) int64 {
    res, n := 0, len(parent)
    graph := make([][]int, n) // graph[i]表示i的所有子节点
    for i := 1; i < n; i++{
        p := parent[i]
        graph[p] = append(graph[p], i)
    }
    cnt := map[int]int{0:1}
    var dfs func(int, int)
    dfs = func(v, xor int) {
        for _, w := range graph[v] {
            x := xor ^ (1 << (s[w] - 'a'))
            res += cnt[x]
            for i := 0; i < 26; i++ {
                res += cnt[x^(1<<i)]
            }
            cnt[x]++
            dfs(w, x)
        }
    }
    dfs(0,0)
    return int64(res)
}

func countPalindromePaths2(parent []int, s string) int64 {
    res, n := int64(0), len(parent)
    codes, count:= make([]int, n, n), make(map[int]int64, n)
    var encodeToRoot func(codes []int, parent []int, n int, s string) int
    encodeToRoot = func(codes []int, parent []int, n int, s string) int {
        if n == 0 { return 0 }
        if codes[n] > 0 { return codes[n] - 1 }
        c := encodeToRoot(codes, parent, parent[n], s) ^ (1 << (s[n] - 'a'))
        codes[n] = 1 + c
        return c
    }
    for i := 0; i < n; i++ {
        c := encodeToRoot(codes, parent, i, s)
        res += count[c]
        for j := 0; j < 26; j++ {
            res += count[(1<<j)^c]
        }
        count[c]++
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2023/07/15/treedrawio-8drawio.png" />
    // Input: parent = [-1,0,0,1,1,2], s = "acaabc"
    // Output: 8
    // Explanation: The valid pairs are:
    // - All the pairs (0,1), (0,2), (1,3), (1,4) and (2,5) result in one character which is always a palindrome.
    // - The pair (2,3) result in the string "aca" which is a palindrome.
    // - The pair (1,5) result in the string "cac" which is a palindrome.
    // - The pair (3,5) result in the string "acac" which can be rearranged into the palindrome "acca".
    fmt.Println(countPalindromePaths([]int{-1,0,0,1,1,2}, "acaabc")) // 8
    // Example 2:
    // Input: parent = [-1,0,0,0,0], s = "aaaaa"
    // Output: 10
    // Explanation: Any pair of nodes (u,v) where u < v is valid.
    fmt.Println(countPalindromePaths([]int{-1,0,0,0,0}, "aaaaa")) // 10

    fmt.Println(countPalindromePaths1([]int{-1,0,0,1,1,2}, "acaabc")) // 8
    fmt.Println(countPalindromePaths1([]int{-1,0,0,0,0}, "aaaaa")) // 10

    fmt.Println(countPalindromePaths2([]int{-1,0,0,1,1,2}, "acaabc")) // 8
    fmt.Println(countPalindromePaths2([]int{-1,0,0,0,0}, "aaaaa")) // 10
}