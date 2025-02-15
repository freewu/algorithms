package main

// 3435. Frequencies of Shortest Supersequences
// You are given an array of strings words. 
// Find all shortest common supersequences (SCS) of words that are not permutations of each other.

// A shortest common supersequence is a string of minimum length that contains each string in words as a subsequence.

// Return a 2D array of integers freqs that represent all the SCSs. 
// Each freqs[i] is an array of size 26, representing the frequency of each letter in the lowercase English alphabet for a single SCS. 
// You may return the frequency arrays in any order.

// Example 1:
// Input: words = ["ab","ba"]
// Output: [[1,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[2,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]]
// Explanation:
// The two SCSs are "aba" and "bab". The output is the letter frequencies for each one.

// Example 2:
// Input: words = ["aa","ac"]
// Output: [[2,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]]
// Explanation:
// The two SCSs are "aac" and "aca". Since they are permutations of each other, keep only "aac".

// Example 3:
// Input: words = ["aa","bb","cc"]
// Output: [[2,2,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]]
// Explanation:
// "aabbcc" and all its permutations are SCSs.

// Constraints:
//     1 <= words.length <= 256
//     words[i].length == 2
//     All strings in words will altogether be composed of no more than 16 unique lowercase letters.
//     All strings in words are unique.

import "fmt"
import "math/bits"

func supersequences(words []string) [][]int {
    // 收集有哪些字母，同时建图
    all, must2 := 0, 0
    g := [26][]int{}
    for _, s := range words {
        x, y := int(s[0]-'a'), int(s[1]-'a')
        all |= 1<<x | 1<<y
        if x == y {
            must2 |= 1 << x
        }
        g[x] = append(g[x], y)
    }
    // 判断是否有环
    hasCycle := func(sub int) bool {
        color := [26]int8{}
        var dfs func(int) bool
        dfs = func(x int) bool {
            color[x] = 1
            for _, y := range g[x] {
                // 只遍历在 sub 中的字母
                if sub>>y&1 == 0 {
                    continue
                }
                if color[y] == 1 || color[y] == 0 && dfs(y) {
                    return true
                }
            }
            color[x] = 2
            return false
        }
        for i, c := range color {
            // 只遍历在 sub 中的字母
            if c == 0 && sub>>i&1 > 0 && dfs(i) {
                return true
            }
        }
        return false
    }
    set := map[int]struct{}{}
    maxSize := 0
    allM := all ^ must2
    for sub, ok := allM, true; ok; ok = sub != allM {
        size := bits.OnesCount(uint(sub))
        // 剪枝：如果 size < maxSize 就不需要判断了
        if size >= maxSize && !hasCycle(sub) {
            if size > maxSize {
                maxSize = size
                clear(set)
            }
            set[sub] = struct{}{}
        }
        sub = (sub - 1) & allM
    }
    res := make([][]int, 0, len(set)) // 预分配空间
    for sub := range set {
        count := make([]int, 26)
        for i := range count {
            count[i] = all >> i & 1 + (all ^ sub) >> i & 1
        }
        res = append(res, count)
    }
    return res
}

func main() {
    // Example 1:
    // Input: words = ["ab","ba"]
    // Output: [[1,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[2,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]]
    // Explanation:
    // The two SCSs are "aba" and "bab". The output is the letter frequencies for each one.
    fmt.Println(supersequences([]string{"ab","ba"})) // [[1,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[2,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]]
    // Example 2:
    // Input: words = ["aa","ac"]
    // Output: [[2,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]]
    // Explanation:
    // The two SCSs are "aac" and "aca". Since they are permutations of each other, keep only "aac".
    fmt.Println(supersequences([]string{"aa","ac"})) // [[2,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]]
    // Example 3:
    // Input: words = ["aa","bb","cc"]
    // Output: [[2,2,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]]
    // Explanation:
    // "aabbcc" and all its permutations are SCSs.
    fmt.Println(supersequences([]string{"aa","bb","cc"})) // [[2,2,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]]

    fmt.Println(supersequences([]string{"bluefrog","leetcode"})) // [[0 1 0 0 1 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0]]
}
