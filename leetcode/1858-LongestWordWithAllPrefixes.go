package main

// 1858. Longest Word With All Prefixes
// Given an array of strings words, find the longest string in words such that every prefix of it is also in words.
//     For example, let words = ["a", "app", "ap"].
//     The string "app" has prefixes "ap" and "a", all of which are in words.

// Return the string described above. 
// If there is more than one string with the same length, 
// return the lexicographically smallest one, and if no string exists, return "".

// Example 1:
// Input: words = ["k","ki","kir","kira", "kiran"]
// Output: "kiran"
// Explanation: "kiran" has prefixes "kira", "kir", "ki", and "k", and all of them appear in words.

// Example 2:
// Input: words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
// Output: "apple"
// Explanation: Both "apple" and "apply" have all their prefixes in words.
// However, "apple" is lexicographically smaller, so we return that.

// Example 3:
// Input: words = ["abc", "bc", "ab", "qwe"]
// Output: ""

// Constraints:
//     1 <= words.length <= 10^5
//     1 <= words[i].length <= 10^5
//     1 <= sum(words[i].length) <= 10^5
//     words[i] consists only of lowercase English letters.

import "fmt"
import "slices"

// 排序 + 剪枝 + trie
// 排序保证其前缀在它之前都已经处理
// 剪枝! 如果一个串s已经不满足条件,那么以s为前缀的剩余串都不满足条件,不需要插入trie
func longestWord(words []string) string {
    type Trie struct {
        son  [26]*Trie // only of lowercase English letters
    }
    res, root := "", &Trie{}
    wordInsert := func(s string) bool { // 如果插入过程中,所有前缀都存在,返回true
        cur, n := root, len(s)
        for _, ch := range s[:n-1] {
            if cur.son[ch - 'a'] == nil { // 剪枝! 途中少了一环,后续以此字符串为前缀的都不可能成为答案, 无需插入了
                return false
            }
            cur = cur.son[ch - 'a'] 
        }
        if cur.son[s[n-1] - 'a'] == nil { // 前面都存在,单独插入最后一个字符
            cur.son[s[n-1] - 'a'] = &Trie{}
        }
        return true
    }
    // 业务逻辑, 按照字符串长度从小到大处理,保证处理s时,其前缀已经全部处理完毕
    slices.SortFunc(words, func(a, b string) int {
        return len(a) - len(b)
    })
    for _, word := range words {
        if wordInsert(word) {
            if len(word) > len(res) || len(word) == len(res) && word < res { // 注意!! 字典序比较放在这里比放在排序中好(字典序比较耗时,放在这比较次数少)
                res = word
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: words = ["k","ki","kir","kira", "kiran"]
    // Output: "kiran"
    // Explanation: "kiran" has prefixes "kira", "kir", "ki", and "k", and all of them appear in words.
    fmt.Println(longestWord([]string{"k","ki","kir","kira", "kiran"})) // "kiran"
    // Example 2:
    // Input: words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
    // Output: "apple"
    // Explanation: Both "apple" and "apply" have all their prefixes in words.
    // However, "apple" is lexicographically smaller, so we return that.
    fmt.Println(longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"})) // "apple"
    // Example 3:
    // Input: words = ["abc", "bc", "ab", "qwe"]
    // Output: ""
    fmt.Println(longestWord([]string{"abc", "bc", "ab", "qwe"})) // ""
}