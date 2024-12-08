package main

// 2156. Find Substring With Given Hash Value
// The hash of a 0-indexed string s of length k, given integers p and m, is computed using the following function:
//     hash(s, p, m) = (val(s[0]) * p0 + val(s[1]) * p1 + ... + val(s[k-1]) * pk-1) mod m.

// Where val(s[i]) represents the index of s[i] in the alphabet from val('a') = 1 to val('z') = 26.

// You are given a string s and the integers power, modulo, k, and hashValue. 
// Return sub, the first substring of s of length k such that hash(sub, power, modulo) == hashValue.

// The test cases will be generated such that an answer always exists.

// A substring is a contiguous non-empty sequence of characters within a string.

// Example 1:
// Input: s = "leetcode", power = 7, modulo = 20, k = 2, hashValue = 0
// Output: "ee"
// Explanation: The hash of "ee" can be computed to be hash("ee", 7, 20) = (5 * 1 + 5 * 7) mod 20 = 40 mod 20 = 0. 
// "ee" is the first substring of length 2 with hashValue 0. Hence, we return "ee".

// Example 2:
// Input: s = "fbxzaad", power = 31, modulo = 100, k = 3, hashValue = 32
// Output: "fbx"
// Explanation: The hash of "fbx" can be computed to be hash("fbx", 31, 100) = (6 * 1 + 2 * 31 + 24 * 312) mod 100 = 23132 mod 100 = 32. 
// The hash of "bxz" can be computed to be hash("bxz", 31, 100) = (2 * 1 + 24 * 31 + 26 * 312) mod 100 = 25732 mod 100 = 32. 
// "fbx" is the first substring of length 3 with hashValue 32. Hence, we return "fbx".
// Note that "bxz" also has a hash of 32 but it appears later than "fbx".

// Constraints:
//     1 <= k <= s.length <= 2 * 10^4
//     1 <= power, modulo <= 10^9
//     0 <= hashValue < modulo
//     s consists of lowercase English letters only.
//     The test cases are generated such that an answer always exists.

import "fmt"

// sliding window
func subStrHash(s string, power int, modulo int, k int, hashValue int) string {
    res, p, n := "", 1, len(s)
    hash := (int(s[n-1]) - 96) % modulo
    for i := n - 1; i - k + 1 >= 0; i-- {
        if i == n-1 {
            for j := 1; j < k; j++ {
                hash = (hash * power % modulo + int(s[i-j]) - 96) % modulo
                p = p * power % modulo
            }
        } else {
            hash = ((hash - (int(s[i + 1]) - 96) * p % modulo + modulo) * power + int(s[i - k + 1]) -  96) % modulo
        }
        if hash == hashValue {
            res = s[i-k+1 : i+1]
        }
    }
    return res
}

func subStrHash1(s string, power int, modulo int, k int, hashValue int) string {
    h, p, n := 0, 1, len(s)
    for i := n - 1; i >= n-k; i-- {
        val := int(s[i] - 'a' + 1)
        h = (h * power % modulo + val) % modulo
        if i != n - k {
            p = p * power % modulo
        }
    }
    j := n - k
    for i := n - k - 1; i >= 0; i-- {
        pre, cur := int(s[i+k] - 'a' + 1), int(s[i] - 'a' + 1)
        h = ((h - pre * p % modulo + modulo) * power % modulo + cur) % modulo
        if h == hashValue {
            j = i
        }
    }
    return s[j : j+k]
}

func main() {
    // Example 1:
    // Input: s = "leetcode", power = 7, modulo = 20, k = 2, hashValue = 0
    // Output: "ee"
    // Explanation: The hash of "ee" can be computed to be hash("ee", 7, 20) = (5 * 1 + 5 * 7) mod 20 = 40 mod 20 = 0. 
    // "ee" is the first substring of length 2 with hashValue 0. Hence, we return "ee".
    fmt.Println(subStrHash("leetcode", 7 ,20, 2, 0)) // "ee"
    // Example 2:
    // Input: s = "fbxzaad", power = 31, modulo = 100, k = 3, hashValue = 32
    // Output: "fbx"
    // Explanation: The hash of "fbx" can be computed to be hash("fbx", 31, 100) = (6 * 1 + 2 * 31 + 24 * 312) mod 100 = 23132 mod 100 = 32. 
    // The hash of "bxz" can be computed to be hash("bxz", 31, 100) = (2 * 1 + 24 * 31 + 26 * 312) mod 100 = 25732 mod 100 = 32. 
    // "fbx" is the first substring of length 3 with hashValue 32. Hence, we return "fbx".
    // Note that "bxz" also has a hash of 32 but it appears later than "fbx".
    fmt.Println(subStrHash("fbxzaad", 31 ,100, 3, 32)) // "fbx"

    fmt.Println(subStrHash("leetcode", 7 ,20, 2, 0)) // "ee"
    fmt.Println(subStrHash("fbxzaad", 31 ,100, 3, 32)) // "fbx"
}