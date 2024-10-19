package main

// 1520. Maximum Number of Non-Overlapping Substrings
// Given a string s of lowercase letters, you need to find the maximum number of non-empty substrings of s that meet the following conditions:
//     The substrings do not overlap, that is for any two substrings s[i..j] and s[x..y], either j < x or i > y is true.
//     A substring that contains a certain character c must also contain all occurrences of c.

// Find the maximum number of substrings that meet the above conditions. 
// If there are multiple solutions with the same number of substrings, return the one with minimum total length. 
// It can be shown that there exists a unique solution of minimum total length.

// Notice that you can return the substrings in any order.

// Example 1:
// Input: s = "adefaddaccc"
// Output: ["e","f","ccc"]
// Explanation: The following are all the possible substrings that meet the conditions:
// [
//   "adefaddaccc"
//   "adefadda",
//   "ef",
//   "e",
//   "f",
//   "ccc",
// ]
// If we choose the first string, we cannot choose anything else and we'd get only 1. If we choose "adefadda", we are left with "ccc" which is the only one that doesn't overlap, thus obtaining 2 substrings. Notice also, that it's not optimal to choose "ef" since it can be split into two. Therefore, the optimal way is to choose ["e","f","ccc"] which gives us 3 substrings. No other solution of the same number of substrings exist.

// Example 2:
// Input: s = "abbaccd"
// Output: ["d","bb","cc"]
// Explanation: Notice that while the set of substrings ["d","abba","cc"] also has length 3, it's considered incorrect since it has larger total length.

// Constraints:
//     1 <= s.length <= 10^5
//     s contains only lowercase English letters.

import "fmt"

func maxNumOfSubstrings(s string) []string {
    n, left, right := len(s), make([]int, 26), make([]int, 26)
    for i := 0; i < 26; i++ { 
        left[i], right[i] = 1 << 31, -1 << 31
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        left[s[i]-'a']  = min(i, left[s[i]-'a'])
        right[s[i]-'a'] = max(i, right[s[i]-'a'])
    }
    // 关键函数 extend
    //计算从第i个字母开始，满足条件的子串的右边边界,left 表示某个字母第一次出现的位置,right 表示某个字母最后一次出现的位置
    //保证下一次迭代，如果有相交，肯定是上一次迭代的子串
    extend := func(i int) int {
        p := right[s[i]-'a'] //最后出现的位置
        pos := i
        for pos < p {
            if left[s[pos]-'a'] < i { return -1 }
            if right[s[pos]-'a'] > p {
                p = right[s[pos]-'a']
            }
            pos++
        }
        return p
    }
    res, last := make([]string, 0), -1
    for i := 0; i < n; i++ {
        if i != left[s[i]-'a'] { continue } // 只计算当前字母第一次出现的位置
        p := extend(i)
        if p == -1 { continue }
        if i > last {
            res = append(res, s[i:p+1])
        } else {
            res[len(res)-1] = s[i : p+1]
        }
        last = p
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "adefaddaccc"
    // Output: ["e","f","ccc"]
    // Explanation: The following are all the possible substrings that meet the conditions:
    // [
    //   "adefaddaccc"
    //   "adefadda",
    //   "ef",
    //   "e",
    //   "f",
    //   "ccc",
    // ]
    // If we choose the first string, we cannot choose anything else and we'd get only 1. If we choose "adefadda", we are left with "ccc" which is the only one that doesn't overlap, thus obtaining 2 substrings. Notice also, that it's not optimal to choose "ef" since it can be split into two. Therefore, the optimal way is to choose ["e","f","ccc"] which gives us 3 substrings. No other solution of the same number of substrings exist.
    fmt.Println(maxNumOfSubstrings("adefaddaccc")) // ["e","f","ccc"]
    // Example 2:
    // Input: s = "abbaccd"
    // Output: ["d","bb","cc"]
    // Explanation: Notice that while the set of substrings ["d","abba","cc"] also has length 3, it's considered incorrect since it has larger total length.
    fmt.Println(maxNumOfSubstrings("abbaccd")) //  ["d","bb","cc"]
}