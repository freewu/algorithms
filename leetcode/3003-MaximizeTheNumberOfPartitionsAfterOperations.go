package main

// 3003. Maximize the Number of Partitions After Operations
// You are given a string s and an integer k.

// First, you are allowed to change at most one index in s to another lowercase English letter.

// After that, do the following partitioning operation until s is empty:
//     1. Choose the longest prefix of s containing at most k distinct characters.
//     2. Delete the prefix from s and increase the number of partitions by one. 
//        The remaining characters (if any) in s maintain their initial order.

// Return an integer denoting the maximum number of resulting partitions 
// after the operations by optimally choosing at most one index to change.

// Example 1:
// Input: s = "accca", k = 2
// Output: 3
// Explanation:
// The optimal way is to change s[2] to something other than a and c, for example, b. then it becomes "acbca".
// Then we perform the operations:
// The longest prefix containing at most 2 distinct characters is "ac", we remove it and s becomes "bca".
// Now The longest prefix containing at most 2 distinct characters is "bc", so we remove it and s becomes "a".
// Finally, we remove "a" and s becomes empty, so the procedure ends.
// Doing the operations, the string is divided into 3 partitions, so the answer is 3.

// Example 2:
// Input: s = "aabaab", k = 3
// Output: 1
// Explanation:
// Initially s contains 2 distinct characters, so whichever character we change, it will contain at most 3 distinct characters, so the longest prefix with at most 3 distinct characters would always be all of it, therefore the answer is 1.

// Example 3:
// Input: s = "xxyz", k = 1
// Output: 4
// Explanation:
// The optimal way is to change s[0] or s[1] to something other than characters in s, for example, to change s[0] to w.
// Then s becomes "wxyz", which consists of 4 distinct characters, so as k is 1, it will divide into 4 partitions.

// Constraints:
//     1 <= s.length <= 10^4
//     s consists only of lowercase English letters.
//     1 <= k <= 26

import "fmt"
import "math/bits"

// dfs
func maxPartitionsAfterOperations(s string, k int) int {
    n := len(s)
    type Tuple struct{ i, cur, t int }
    mp := make(map[Tuple]int)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(i, cur, t int) int
    dfs = func(i, cur, t int) int {
        if i >= n { return 1 }
        key := Tuple{i, cur, t}
        if v, ok := mp[key]; ok { return v }
        v := 1 << (s[i] - 'a')
        res, next := 0, cur | v
        if bits.OnesCount(uint(next)) > k {
            res = dfs(i + 1, v, t) + 1
        } else {
            res = dfs(i + 1, next, t)
        }
        if t > 0 {
            for j := 0; j < 26; j++ {
                next = cur | (1 << j)
                if bits.OnesCount(uint(next)) > k {
                    res = max(res, dfs(i + 1, 1 << j, 0) + 1)
                } else {
                    res = max(res, dfs(i + 1, next, 0))
                }
            }
        }
        mp[key] = res
        return res
    }
    return dfs(0, 0, 1)
}

func maxPartitionsAfterOperations1(s string, k int) int {
    if k == 26 { return 1 }
    seg, mask, size := 1, 0, 0
    update := func(i int) {
        bit := 1 << (s[i] - 'a')
        if mask&bit > 0 { return }
        size++
        if size > k {
            seg++ // s[i] 在新的一段中
            mask = bit
            size = 1
        } else {
            mask |= bit
        }
    }
    n := len(s)
    type Pair struct{ seg, mask int }
    suffix := make([]Pair, n + 1)
    for i := n - 1; i >= 0; i-- {
        update(i)
        suffix[i] = Pair{ seg, mask }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res := seg // 不修改任何字母
    seg, mask, size = 1, 0, 0
    for i := range s {
        p := suffix[i+1]
        v := seg + p.seg // 情况 3
        unionSize := bits.OnesCount(uint(mask | p.mask))
        if unionSize < k {
            v-- // 情况 1
        } else if unionSize < 26 && size == k && bits.OnesCount(uint(p.mask)) == k {
            v++ // 情况 2
        }
        res = max(res, v)
        update(i)
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "accca", k = 2
    // Output: 3
    // Explanation:
    // The optimal way is to change s[2] to something other than a and c, for example, b. then it becomes "acbca".
    // Then we perform the operations:
    // The longest prefix containing at most 2 distinct characters is "ac", we remove it and s becomes "bca".
    // Now The longest prefix containing at most 2 distinct characters is "bc", so we remove it and s becomes "a".
    // Finally, we remove "a" and s becomes empty, so the procedure ends.
    // Doing the operations, the string is divided into 3 partitions, so the answer is 3.
    fmt.Println(maxPartitionsAfterOperations("accca", 2)) // 3
    // Example 2:
    // Input: s = "aabaab", k = 3
    // Output: 1
    // Explanation:
    // Initially s contains 2 distinct characters, so whichever character we change, it will contain at most 3 distinct characters, so the longest prefix with at most 3 distinct characters would always be all of it, therefore the answer is 1.
    fmt.Println(maxPartitionsAfterOperations("aabaab", 3)) // 1
    // Example 3:
    // Input: s = "xxyz", k = 1
    // Output: 4
    // Explanation:
    // The optimal way is to change s[0] or s[1] to something other than characters in s, for example, to change s[0] to w.
    // Then s becomes "wxyz", which consists of 4 distinct characters, so as k is 1, it will divide into 4 partitions.
    fmt.Println(maxPartitionsAfterOperations("xxyz", 1)) // 4

    fmt.Println(maxPartitionsAfterOperations("bluefrog", 1)) // 8
    fmt.Println(maxPartitionsAfterOperations("leetcode", 1)) // 8

    fmt.Println(maxPartitionsAfterOperations1("accca", 2)) // 3
    fmt.Println(maxPartitionsAfterOperations1("aabaab", 3)) // 1
    fmt.Println(maxPartitionsAfterOperations1("xxyz", 1)) // 4
    fmt.Println(maxPartitionsAfterOperations1("bluefrog", 1)) // 8
    fmt.Println(maxPartitionsAfterOperations1("leetcode", 1)) // 8
}