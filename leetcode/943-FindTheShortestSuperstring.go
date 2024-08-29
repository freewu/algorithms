package main

// 943. Find the Shortest Superstring
// Given an array of strings words, return the smallest string that contains each string in words as a substring. 
// If there are multiple valid strings of the smallest length, return any of them.

// You may assume that no string in words is a substring of another string in words.

// Example 1:
// Input: words = ["alex","loves","leetcode"]
// Output: "alexlovesleetcode"
// Explanation: All permutations of "alex","loves","leetcode" would also be accepted.

// Example 2:
// Input: words = ["catg","ctaagt","gcta","ttca","atgcatc"]
// Output: "gctaagttcatgcatc"

// Constraints:
//     1 <= words.length <= 12
//     1 <= words[i].length <= 20
//     words[i] consists of lowercase English letters.
//     All the strings of words are unique.

import "fmt"
import "strings"

func shortestSuperstring(words []string) string {
    n, inf := len(words), 1 << 31
    graph := make([][]int, n)
    calculateAdditionalLength := func (a, b string) int {
        for i := 0; i < len(a); i++ {
            if strings.HasPrefix(b, a[i:]) {
                return len(b) - len(a) + i
            }
        }
        return len(b)
    }
    for i := 0; i < n; i++ {
        graph[i] = make([]int, n)
        for j := 0; j < n; j++ {
            if i == j {
                continue
            }
            graph[i][j] = calculateAdditionalLength(words[i], words[j])
        }
    }
    pow2n := 1 << n
    dp, path := make([][]int, pow2n), make([][]int, pow2n)
    for i := 0; i < pow2n; i++ {
        dp[i], path[i] = make([]int, n), make([]int, n)
    }
    last, mn := -1, inf
    for i := 1; i < pow2n; i++ {
        for j := 0; j < n; j++ {
            dp[i][j] = inf
            pow2j := 1 << j
            if i&pow2j > 0 {
                prev := i - pow2j
                if prev == 0 {
                    dp[i][j] = len(words[j])
                } else {
                    for k := 0; k < n; k++ {
                        if dp[prev][k] < inf && dp[prev][k]+graph[k][j] < dp[i][j] {
                            dp[i][j] = dp[prev][k] + graph[k][j]
                            path[i][j] = k
                        }
                    }
                }
            }
            if i == pow2n-1 && dp[i][j] < mn {
                mn = dp[i][j]
                last = j
            }
        }
    }
    cur := pow2n - 1
    stack := make([]int, 0, n)
    for cur > 0 {
        stack = append(stack, last)
        cur, last = cur-(1<<last), path[cur][last]
    }
    i := stack[len(stack)-1]
    stack = stack[:len(stack)-1]
    var sb strings.Builder
    sb.WriteString(words[i])
    for len(stack) > 0 {
        j := stack[len(stack)-1]
        sb.WriteString(words[j][len(words[j]) - graph[i][j]:])
        i, stack = j, stack[:len(stack)-1]
    }
    return sb.String()
}

func shortestSuperstring1(words []string) string {
    n,inf := len(words), 1 << 31
    // adjacent map
    // words[i] = "ab", Awords[j] = "abcde"
    // adj_map[i][j] = 3, it means that we can only add 3 char in words[j]
    // after words[i]
    adj_map := make([][]int, n)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i:=0; i<n; i++ {
        adj_map[i] = make([]int, n)
        for j:=0; j<n; j++ {
            aj, ai := []byte(words[j]), []byte(words[i]) 
            adj_map[i][j] = len(aj)
            for k := 1; k <= min(len(ai), len(aj)); k++ {
                if string(ai[len(ai)-k:]) == string(aj[:k]) {
                    adj_map[i][j] = len(aj) - k
                }
            }
        }
    }
    // dp[s][i] := min distance to visit nodes 
    // (represented as a binary state s) once and only once 
    // and the path ends with node i.
    dp := make([][]int, 1 << n)
    // record the path
    parent := make([][]int, 1 << n)
    // dp[7][1] is the min distance to visit nodes (0, 1, 2) 
    // and ends with node 1, the possible paths could be 
    // (0, 2, 1), (2, 0, 1). parent[7][1] could be 2 or 0
    // init dp and parent
    for i:=0; i < (1 << n); i++ {
        dp[i] = make([]int, n)
        parent[i] = make([]int, n)
        for j:=0; j<n; j++ {
            dp[i][j] = inf
            parent[i][j] = -1
        }
    }
    // Init
    for i := 0; i < n; i++ {   
        dp[1<< i ][i] = len([]byte(words[i]))
    }
    // Transition
    for s := 1; s < (1 << n ); s++ {
        for j:=0; j<n; j++ {
        	// if node j is not in s, skip
            if s & (1 << j ) == 0 { continue }
            // remove node j from s
            ps := s & ^ (1 << j )
            // test all the parent of j
            for i := 0; i < n; i++ {
            	// find out the i with min(dp[ps][i] + adj_map[i][j])
            	// record it in dp[s][j] and parent[s][j]
                if dp[ps][i] + adj_map[i][j] < dp[s][j] {
                    dp[s][j] = dp[ps][i] + adj_map[i][j]
                    parent[s][j] = i;
                }
            }
        }
    }
    
    // find out the last node in shortest path
    j, mn := 0, inf
    for i := 0; i < n; i++ {
        if dp[(1 << n) - 1][i] < mn {
            mn = dp[(1 << n) - 1][i]
            j = i
        }
    }
    // get the whole path by parent[s][j] 
    // and build the Shortest Superstring
    s := (1 << n) - 1
    res := ""
    for s != 0 {
        i := parent[s][j]
        if i < 0 {
            res = words[j] + res
        } else {
            adj := []byte(words[j])
            res = string(adj[len(adj) - adj_map[i][j]:]) + res
        }
        s &= ^(1 << j )
        j = i
    }
    return res
}

func main() {
    // Example 1:
    // Input: words = ["alex","loves","leetcode"]
    // Output: "alexlovesleetcode"
    // Explanation: All permutations of "alex","loves","leetcode" would also be accepted.
    fmt.Println(shortestSuperstring([]string{"alex","loves","leetcode"})) // "alexlovesleetcode"
    // Example 2:
    // Input: words = ["catg","ctaagt","gcta","ttca","atgcatc"]
    // Output: "gctaagttcatgcatc"
    fmt.Println(shortestSuperstring([]string{"catg","ctaagt","gcta","ttca","atgcatc"})) //"gctaagttcatgcatc"

    fmt.Println(shortestSuperstring1([]string{"alex","loves","leetcode"})) // "alexlovesleetcode"
    fmt.Println(shortestSuperstring1([]string{"catg","ctaagt","gcta","ttca","atgcatc"})) //"gctaagttcatgcatc"
}