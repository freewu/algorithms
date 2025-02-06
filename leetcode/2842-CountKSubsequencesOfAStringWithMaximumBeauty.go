package main

// 2842. Count K-Subsequences of a String With Maximum Beauty
// You are given a string s and an integer k.

// A k-subsequence is a subsequence of s, having length k, and all its characters are unique, i.e., every character occurs once.

// Let f(c) denote the number of times the character c occurs in s.

// The beauty of a k-subsequence is the sum of f(c) for every character c in the k-subsequence.

// For example, consider s = "abbbdd" and k = 2:
//     1. f('a') = 1, f('b') = 3, f('d') = 2
//     2. Some k-subsequences of s are:
//         2.1 "abbbdd" -> "ab" having a beauty of f('a') + f('b') = 4
//         2.2 "abbbdd" -> "ad" having a beauty of f('a') + f('d') = 3
//         2.3 "abbbdd" -> "bd" having a beauty of f('b') + f('d') = 5

// Return an integer denoting the number of k-subsequences whose beauty is the maximum among all k-subsequences. 
// Since the answer may be too large, return it modulo 10^9 + 7.

// A subsequence of a string is a new string formed from the original string by deleting some (possibly none) of the characters without disturbing the relative positions of the remaining characters.

// Notes
//     1. f(c) is the number of times a character c occurs in s, not a k-subsequence.
//     2. Two k-subsequences are considered different if one is formed by an index that is not present in the other. 
//        So, two k-subsequences may form the same string.

// Example 1:
// Input: s = "bcca", k = 2
// Output: 4
// Explanation: From s we have f('a') = 1, f('b') = 1, and f('c') = 2.
// The k-subsequences of s are: 
// bcca having a beauty of f('b') + f('c') = 3 
// bcca having a beauty of f('b') + f('c') = 3 
// bcca having a beauty of f('b') + f('a') = 2 
// bcca having a beauty of f('c') + f('a') = 3
// bcca having a beauty of f('c') + f('a') = 3 
// There are 4 k-subsequences that have the maximum beauty, 3. 
// Hence, the answer is 4. 

// Example 2:
// Input: s = "abbcd", k = 4
// Output: 2
// Explanation: From s we have f('a') = 1, f('b') = 2, f('c') = 1, and f('d') = 1. 
// The k-subsequences of s are: 
// abbcd having a beauty of f('a') + f('b') + f('c') + f('d') = 5
// abbcd having a beauty of f('a') + f('b') + f('c') + f('d') = 5 
// There are 2 k-subsequences that have the maximum beauty, 5. 
// Hence, the answer is 2. 

// Constraints:
//     1 <= s.length <= 2 * 10^5
//     1 <= k <= s.length
//     s consists only of lowercase English letters.

import "fmt"
import "sort"

func countKSubsequencesWithMaxBeauty(s string, k int) int {
    res, count, mod := 1, 0, 1_000_000_007
    freq, arr := [26]int{}, []int{}
    for _, v := range s {
        freq[v - 'a']++
        if freq[v - 'a'] == 1 { count++ }
    }
    if count < k { return 0 }
    for _, v := range freq {
        if v > 0 {
            arr = append(arr, v)
        }
    }
    sort.Slice(arr, func(i, j int) bool {
        return arr[i] > arr[j]
    })
    x, val := 0, arr[k - 1]
    for _, v := range arr {
        if v == val {
            x++
        }
    }
    for _, v := range arr {
        if v == val { break }
        k--
        res = res * v % mod
    }
    c := make([][]int, x + 1)
    for i := range c {
        c[i] = make([]int, x + 1)
        c[i][0] = 1
        for j := 1; j <= i; j++ {
            c[i][j] = (c[i-1][j-1] + c[i-1][j]) % mod
        }
    }
    pow := func(x, y int) int {
        res := 1
        for ; y > 0; y >>= 1 {
            if y & 1 == 1 {
                res = res * x % mod
            }
            x = x * x % mod
        }
        return res
    }
    return (res * c[x][k] % mod) * pow(val, k) % mod
}

func countKSubsequencesWithMaxBeauty1(s string, k int) int {
    count := [26]int{}
    for _, v := range s {
        count[v - 'a']++
    }
    mp := make(map[int]int)
    for _, v := range count {
        if v > 0 {
            mp[v]++
        }
    }
    type Pair struct{ count, num int }
    arr := make([]Pair, 0, len(mp))
    for k, v := range mp {
        arr = append(arr, Pair{ k, v })
    }
    sort.Slice(arr, func(i, j int) bool { 
        return arr[i].count > arr[j].count 
    })
    res, mod := 1, 1_000_000_007
    pow := func(x, n int) int {
        res := 1
        for ; n > 0; n /= 2 {
            if n % 2 > 0 {
                res = res * x % mod
            }
            x = x * x % mod
        }
        return res
    }
    comb := func (n, k int) int {
        res := n
        for i := 2; i <= k; i++ {
            res = res * (n - i + 1) / i // n, n - 1, n - 2,... 中的前 i 个数至少有一个因子 i
        }
        return res % mod
    }
    for _, p := range arr {
        if p.num >= k {
            return res * pow(p.count, k) % mod * comb(p.num, k) % mod
        }
        res = res * pow(p.count, p.num) % mod
        k -= p.num
    }
    return 0 // k 太大，无法选 k 个不一样的字符
}

func main() {
    // Example 1:
    // Input: s = "bcca", k = 2
    // Output: 4
    // Explanation: From s we have f('a') = 1, f('b') = 1, and f('c') = 2.
    // The k-subsequences of s are: 
    // bcca having a beauty of f('b') + f('c') = 3 
    // bcca having a beauty of f('b') + f('c') = 3 
    // bcca having a beauty of f('b') + f('a') = 2 
    // bcca having a beauty of f('c') + f('a') = 3
    // bcca having a beauty of f('c') + f('a') = 3 
    // There are 4 k-subsequences that have the maximum beauty, 3. 
    // Hence, the answer is 4. 
    fmt.Println(countKSubsequencesWithMaxBeauty("bcca", 2)) // 4
    // Example 2:
    // Input: s = "abbcd", k = 4
    // Output: 2
    // Explanation: From s we have f('a') = 1, f('b') = 2, f('c') = 1, and f('d') = 1. 
    // The k-subsequences of s are: 
    // abbcd having a beauty of f('a') + f('b') + f('c') + f('d') = 5
    // abbcd having a beauty of f('a') + f('b') + f('c') + f('d') = 5 
    // There are 2 k-subsequences that have the maximum beauty, 5. 
    // Hence, the answer is 2. 
    fmt.Println(countKSubsequencesWithMaxBeauty("abbcd", 4)) // 2

    fmt.Println(countKSubsequencesWithMaxBeauty("bluefrog", 4)) // 70
    fmt.Println(countKSubsequencesWithMaxBeauty("leetcode", 4)) // 30

    fmt.Println(countKSubsequencesWithMaxBeauty1("bcca", 2)) // 4
    fmt.Println(countKSubsequencesWithMaxBeauty1("abbcd", 4)) // 2
    fmt.Println(countKSubsequencesWithMaxBeauty1("bluefrog", 4)) // 70
    fmt.Println(countKSubsequencesWithMaxBeauty1("leetcode", 4)) // 30
}