package main

// 3598. Longest Common Prefix Between Adjacent Strings After Removals
// You are given an array of strings words. 
// For each index i in the range [0, words.length - 1], perform the following steps:
//     1. Remove the element at index i from the words array.
//     2. Compute the length of the longest common prefix among all adjacent pairs in the modified array.

// Return an array answer, where answer[i] is the length of the longest common prefix between the adjacent pairs after removing the element at index i. 
// If no adjacent pairs remain or if none share a common prefix, then answer[i] should be 0.

// Example 1:
// Input: words = ["jump","run","run","jump","run"]
// Output: [3,0,0,3,3]
// Explanation:
// Removing index 0:
// words becomes ["run", "run", "jump", "run"]
// Longest adjacent pair is ["run", "run"] having a common prefix "run" (length 3)
// Removing index 1:
// words becomes ["jump", "run", "jump", "run"]
// No adjacent pairs share a common prefix (length 0)
// Removing index 2:
// words becomes ["jump", "run", "jump", "run"]
// No adjacent pairs share a common prefix (length 0)
// Removing index 3:
// words becomes ["jump", "run", "run", "run"]
// Longest adjacent pair is ["run", "run"] having a common prefix "run" (length 3)
// Removing index 4:
// words becomes ["jump", "run", "run", "jump"]
// Longest adjacent pair is ["run", "run"] having a common prefix "run" (length 3)

// Example 2:
// Input: words = ["dog","racer","car"]
// Output: [0,0,0]
// Explanation:
// Removing any index results in an answer of 0.

// Constraints:
//     1 <= words.length <= 10^5
//     1 <= words[i].length <= 10^4
//     words[i] consists of lowercase English letters.
//     The sum of words[i].length is smaller than or equal 10^5.

import "fmt"

func longestCommonPrefix(words []string) []int {
    n := len(words)
    res := make([]int, n)
    if n == 1 { return res } // 不存在相邻对
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    lcp := func (s, t string) int {
        count, n := 0, min(len(s), len(t))
        for i := 0; i < n && s[i] == t[i]; i++ {
            count++
        }
        return count
    }
    suffix := make([]int, n) // 后缀 [i,n-1] 中的相邻 LCP 长度的最大值
    for i := n - 2; i > 0; i-- {
        suffix[i] = max(suffix[i+1],lcp(words[i], words[i+1]))
    }
    res[0] = suffix[1]
    prefix := 0 // 前缀 [0,i-1] 中的相邻 LCP 长度的最大值
    for i := 1; i < n-1; i++ {
        res[i] = max(prefix, max(lcp(words[i-1], words[i+1]), suffix[i+1]))
        prefix = max(prefix, lcp(words[i-1], words[i])) // 为下一轮循环做准备
    }
    res[n-1] = prefix
    return res
}

func longestCommonPrefix1(words []string) []int {
    n := len(words)
    res := make([]int, n-1)
    type Pair struct{ val, left, right int } // 左边界  右边界
    var first, second, third Pair
    for i := range res {
        a, b := words[i], words[i+1]
        j, k := 0, 0
        for j < len(a) && k < len(b) && a[j] == b[k] {
            j++
            k++
        }
        res[i] = min(j, k)
        if res[i] > first.val {
            second, third = first, second
            first.val = res[i]
            first.left = i
            first.right = i + 1
        } else if res[i] > second.val {
            third = second
            second.val = res[i]
            second.left = i
            second.right = i + 1
        } else if res[i] > third.val {
            third.val = res[i]
            third.left = i
            third.right = i + 1
        }
    }
    res = make([]int, n)
    for i := range res {
        if first.left != i && first.right != i {
            res[i] = first.val
        } else if second.left != i && second.right != i {
            res[i] = second.val
        } else {
            res[i] = third.val
        }
        if i-1 >= 0 && i+1 < n {
            a, b := words[i-1], words[i+1]
            j, k := 0, 0
            for j < len(a) && k < len(b) && a[j] == b[k] {
                j++
                k++
            }
            res[i] = max(res[i], min(j, k))
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: words = ["jump","run","run","jump","run"]
    // Output: [3,0,0,3,3]
    // Explanation:
    // Removing index 0:
    // words becomes ["run", "run", "jump", "run"]
    // Longest adjacent pair is ["run", "run"] having a common prefix "run" (length 3)
    // Removing index 1:
    // words becomes ["jump", "run", "jump", "run"]
    // No adjacent pairs share a common prefix (length 0)
    // Removing index 2:
    // words becomes ["jump", "run", "jump", "run"]
    // No adjacent pairs share a common prefix (length 0)
    // Removing index 3:
    // words becomes ["jump", "run", "run", "run"]
    // Longest adjacent pair is ["run", "run"] having a common prefix "run" (length 3)
    // Removing index 4:
    // words becomes ["jump", "run", "run", "jump"]
    // Longest adjacent pair is ["run", "run"] having a common prefix "run" (length 3)
    fmt.Println(longestCommonPrefix([]string{"jump","run","run","jump","run"})) // [3,0,0,3,3]
    // Example 2:
    // Input: words = ["dog","racer","car"]
    // Output: [0,0,0]
    // Explanation:
    // Removing any index results in an answer of 0.
    fmt.Println(longestCommonPrefix([]string{"dog","racer","car"})) // [0,0,0]

    fmt.Println(longestCommonPrefix([]string{"leetcode","bluefrog","freewu"})) // [0,0,0]

    fmt.Println(longestCommonPrefix1([]string{"jump","run","run","jump","run"})) // [3,0,0,3,3]
    fmt.Println(longestCommonPrefix1([]string{"dog","racer","car"})) // [0,0,0]
    fmt.Println(longestCommonPrefix1([]string{"leetcode","bluefrog","freewu"})) // [0,0,0]
}