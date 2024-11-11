package main

// 1830. Minimum Number of Operations to Make String Sorted
// You are given a string s (0-indexed)​​​​​​. 
// You are asked to perform the following operation on s​​​​​​ until you get a sorted string:
//     1. Find the largest index i such that 1 <= i < s.length and s[i] < s[i - 1].
//     2. Find the largest index j such that i <= j < s.length and s[k] < s[i - 1] 
//        for all the possible values of k in the range [i, j] inclusive.
//     3. Swap the two characters at indices i - 1​​​​ and j​​​​​.
//     4. Reverse the suffix starting at index i​​​​​​.

// Return the number of operations needed to make the string sorted. 
// Since the answer can be too large, return it modulo 10^9 + 7.

// Example 1:
// Input: s = "cba"
// Output: 5
// Explanation: The simulation goes as follows:
// Operation 1: i=2, j=2. Swap s[1] and s[2] to get s="cab", then reverse the suffix starting at 2. Now, s="cab".
// Operation 2: i=1, j=2. Swap s[0] and s[2] to get s="bac", then reverse the suffix starting at 1. Now, s="bca".
// Operation 3: i=2, j=2. Swap s[1] and s[2] to get s="bac", then reverse the suffix starting at 2. Now, s="bac".
// Operation 4: i=1, j=1. Swap s[0] and s[1] to get s="abc", then reverse the suffix starting at 1. Now, s="acb".
// Operation 5: i=2, j=2. Swap s[1] and s[2] to get s="abc", then reverse the suffix starting at 2. Now, s="abc".

// Example 2:
// Input: s = "aabaa"
// Output: 2
// Explanation: The simulation goes as follows:
// Operation 1: i=3, j=4. Swap s[2] and s[4] to get s="aaaab", then reverse the substring starting at 3. Now, s="aaaba".
// Operation 2: i=4, j=4. Swap s[3] and s[4] to get s="aaaab", then reverse the substring starting at 4. Now, s="aaaab".

// Constraints:
//     1 <= s.length <= 3000
//     s​​​​​​ consists only of lowercase English letters.

import "fmt"

func makeStringSorted(s string) int {
    res, n, mod := 0, len(s), 1_000_000_007
    count, factMemo, inverseMemo := [26]int{}, [3001]int{0: 1}, [3001]int{0: 1}
    modPow := func(base, exp, mod int) int {
        res := 1
        base %= mod
        if base == 0 { return 0 }
        for ; exp > 0; exp >>= 1 {
            if (exp & 1) == 1 { res = (res * base) % mod }
            base = (base * base) % mod
        }
        return res
    }
    for i := 1; i <= n; i++ {
        factMemo[i] = (i * factMemo[i-1]) % mod
        inverseMemo[i] = modPow(factMemo[i], mod - 2, mod)
    }
    for i := n - 1; i >= 0; i-- {
        count[s[i]-'a']++
        less := 0
        for j := byte(0); j < s[i]-'a'; j++ {
            less += count[j]
        }
        perms := (less * factMemo[n-1-i]) % mod
        for _, c := range count {
            perms = (perms * inverseMemo[c]) % mod
        }
        res = (res + perms) % mod
    }
    return res
}

func makeStringSorted1(s string) int {
    res, n, mod := 0, len(s), 1_000_000_007
    count, cache := [26]int{}, []int{1, 1}
    inverse := func (x int) int {
        for i := len(cache); i <= x; i++ {
            cache = append(cache, mod - ((mod / i) * cache[mod % i]) % mod )
        }
        return cache[x]
    }
    for i, comb := n - 1, 1; i >= 0; i-- {
        j, add := s[i]-'a', 0
        count[j]++
        comb = (comb * (n - i)) % mod
        comb = (comb * inverse(count[j])) % mod
        for _, x := range count[:j] {
            add += x
        }
        add = (add * comb) % mod
        add = (add * inverse(n-i)) % mod
        res = (res + add) % mod
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "cba"
    // Output: 5
    // Explanation: The simulation goes as follows:
    // Operation 1: i=2, j=2. Swap s[1] and s[2] to get s="cab", then reverse the suffix starting at 2. Now, s="cab".
    // Operation 2: i=1, j=2. Swap s[0] and s[2] to get s="bac", then reverse the suffix starting at 1. Now, s="bca".
    // Operation 3: i=2, j=2. Swap s[1] and s[2] to get s="bac", then reverse the suffix starting at 2. Now, s="bac".
    // Operation 4: i=1, j=1. Swap s[0] and s[1] to get s="abc", then reverse the suffix starting at 1. Now, s="acb".
    // Operation 5: i=2, j=2. Swap s[1] and s[2] to get s="abc", then reverse the suffix starting at 2. Now, s="abc".
    fmt.Println(makeStringSorted("cba")) // 5
    // Example 2:
    // Input: s = "aabaa"
    // Output: 2
    // Explanation: The simulation goes as follows:
    // Operation 1: i=3, j=4. Swap s[2] and s[4] to get s="aaaab", then reverse the substring starting at 3. Now, s="aaaba".
    // Operation 2: i=4, j=4. Swap s[3] and s[4] to get s="aaaab", then reverse the substring starting at 4. Now, s="aaaab".
    fmt.Println(makeStringSorted("aabaa")) // 2

    fmt.Println(makeStringSorted1("cba")) // 5
    fmt.Println(makeStringSorted1("aabaa")) // 2

}