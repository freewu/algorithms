package main

// 3389. Minimum Operations to Make Character Frequencies Equal
// You are given a string s.

// A string t is called good if all characters of t occur the same number of times.

// You can perform the following operations any number of times:
//     1. Delete a character from s.
//     2. Insert a character in s.
//     3. Change a character in s to its next letter in the alphabet.

// Note that you cannot change 'z' to 'a' using the third operation.

// Return the minimum number of operations required to make s good.

// Example 1:
// Input: s = "acab"
// Output: 1
// Explanation:
// We can make s good by deleting one occurrence of character 'a'.

// Example 2:
// Input: s = "wddw"
// Output: 0
// Explanation:
// We do not need to perform any operations since s is initially good.

// Example 3:
// Input: s = "aaabc"
// Output: 2
// Explanation:
// We can make s good by applying these operations:
// Change one occurrence of 'a' to 'b'
// Insert one occurrence of 'c' into s

// Constraints:
//     3 <= s.length <= 2 * 10^4
//     s contains only lowercase English letters.

import "fmt"

func makeStringGood(s string) int {
    n := len(s)
    freq := make([]int, 26)
    for i := 0; i < n; i++ {
        freq[s[i] - 'a']++
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res := n
    for m := 1; m <= n; m++ {
        dp1, dp2, free := 0, 0, 0
        for i := 0; i < 26; i++ {
            ndp1, ndp2, nfree := 0, 0, 0
            if freq[i] >= m {
                ndp1, ndp2, nfree = min(dp1, dp2) + (freq[i] - m), n, freq[i] - m
            } else{
                ndp1, ndp2, nfree = min(dp1, dp2) + freq[i], min(dp1 + max(0, m - freq[i] - free), dp2 + m - freq[i]), freq[i]
            }
            dp1, dp2, free = ndp1, ndp2, nfree
        }
        res = min(res, min(dp1, dp2))
    }
    return res
}

func makeStringGood1(s string) int {
    freq := make([]int, 26)
    for _, v := range s {
        freq[v - 'a']++
    }
    mp := make(map[int]bool)
    mp[freq[0]] = true
    for i := 1; i < 26; i++ {
        v, w := freq[i - 1], freq[i]
        mp[w] = true
        mp[v + w] = true
        mp[(v + w) / 2] = true
        mp[(v + w + 1) / 2] = true
    }
    res := len(s) // target = 0 时的答案
    f := [27]int{}
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for target := range mp {
        f[25] = min(freq[25], abs(freq[25] - target))
        for i := 24; i >= 0; i-- {
            x := freq[i]
            if x == 0 {
                f[i] = f[i + 1]
                continue
            }
            // 单独操作 x（变成 target 或 0）
            f[i] = f[i+1] + min(x, abs(x-target))
            // x 变成 target 或 0，y 变成 target
            y := freq[i+1]
            if 0 < y && y < target { // 只有当 y 需要变大时，才去执行第三种操作
                if x > target { // x 变成 target
                    f[i] = min(f[i], f[i+2] + max(x-target, target-y))
                } else { // x 变成 0
                    f[i] = min(f[i], f[i+2] + max(x, target-y))
                }
            }
        }
        res = min(res, f[0])
    }
    return res
}


func main() {
    // Example 1:
    // Input: s = "acab"
    // Output: 1
    // Explanation:
    // We can make s good by deleting one occurrence of character 'a'.
    fmt.Println(makeStringGood("acab")) // 1
    // Example 2:
    // Input: s = "wddw"
    // Output: 0
    // Explanation:
    // We do not need to perform any operations since s is initially good.
    fmt.Println(makeStringGood("wddw")) // 0
    // Example 3:
    // Input: s = "aaabc"
    // Output: 2
    // Explanation:
    // We can make s good by applying these operations:
    // Change one occurrence of 'a' to 'b'
    // Insert one occurrence of 'c' into s
    fmt.Println(makeStringGood("aaabc")) // 2

    fmt.Println(makeStringGood("leetcode")) // 2
    fmt.Println(makeStringGood("bluefrog")) // 0

    fmt.Println(makeStringGood1("acab")) // 1
    fmt.Println(makeStringGood1("wddw")) // 0
    fmt.Println(makeStringGood1("aaabc")) // 2
    fmt.Println(makeStringGood1("leetcode")) // 2
    fmt.Println(makeStringGood1("bluefrog")) // 0
}