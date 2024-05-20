package main

// 267. Palindrome Permutation II
// Given a string s, return all the palindromic permutations (without duplicates) of it.
// You may return the answer in any order. If s has no palindromic permutation, return an empty list.

// Example 1:
// Input: s = "aabb"
// Output: ["abba","baab"]

// Example 2:
// Input: s = "abc"
// Output: []

// Constraints:
//     1 <= s.length <= 16
//     s consists of only lowercase English letters.

import "fmt"

// backtrack
func generatePalindromes(s string) []string {
    if len(s) == 1 {
        return []string{s}
    }
    cnt := map[byte]int{} // 计算每个字符的数量
    for _, b := range []byte(s) {
        cnt[b]++
    }
    b := []byte{}
    odd_b, odd_cnt := "", 0 // 奇数个字符及数量
    for k, v := range cnt {
        if v % 2 != 0 {
            odd_b = string(k)
            odd_cnt++
            if odd_cnt > 1 { // 如果数量为奇数，并且超过1个，肯定不能满足，只能有一个放中间
                return []string{}
            }
        }
        cnt[k] /= 2 // 将数量/2，只考虑一半字符串的所有组合情况
        for i:=0; i<cnt[k]; i++ {
            b = append(b, k)
        }
    }
    res, visit := []string{}, map[string]bool{} // 存储访问过的字符串
    reverse := func (s string) string {
        str := []byte(s)
        left, right := 0, len(str)-1
        for left < right {
            str[left], str[right] = str[right], str[left]
            left++
            right--
        }
        return string(str)
    }
    var backtrack func(b []byte, s string, odd_b string, cnt *map[byte]int, visit *map[string]bool, ans *[]string)
    backtrack = func(b []byte, s string, odd_b string, cnt *map[byte]int, visit *map[string]bool, ans *[]string) {
        if len(s) == len(b) && !(*visit)[s] { // 如果长度一致，并且没有访问过
            *ans = append(*ans, s+odd_b + reverse(s)) // 因为是一般的字符串，所以是s + 奇数个字符 + 倒序s
            (*visit)[s] = true
            return
        }
        for i := 0; i < len(b); i++ {
            if (*cnt)[b[i]] == 0 { // 如果字符没有了，则继续
                continue
            }
            s += string(b[i])
            (*cnt)[b[i]]--
            backtrack(b, s, odd_b, cnt, visit, ans)
            s = s[:len(s)-1]
            (*cnt)[b[i]]++
        }
    }
    backtrack(b, "", odd_b, &cnt, &visit, &res)
    return res
}

// dfs
func generatePalindromes1(s string) []string {
    res, ch, dedupMap, l, odd := []string{}, [26]int{}, map[string]bool{}, len(s), 0
    for i := 0; i < l; i++ {
        c := s[i]
        ch[c-'a']++
        odd ^= int(c - 'a')
    }
    cs := make([]byte, l)
    if l % 2 > 0 { // 放到中间 用掉不恢复
        cs[l/2] = byte(odd) + 'a'
    }
    var dfs func(pos int)
    dfs = func(pos int) {
        if pos >= l/2 {
            ss := string(cs)
            if !dedupMap[ss] {
                res = append(res, ss)
                dedupMap[ss] = true
            }
            return
        }
        for i, cnt := range ch {
            if cnt >= 2 {
                cs[pos] = 'a' + byte(i)
                cs[l-pos-1] = cs[pos]
                ch[i] -= 2
                dfs(pos + 1)
                ch[i] += 2 // 撤回
            } 
        }
    }
    dfs(0)
    return res
}

func main() {
    // Example 1:
    // Input: s = "aabb"
    // Output: ["abba","baab"]
    fmt.Println(generatePalindromes("aabb")) // ["abba","baab"]
    // Example 2:
    // Input: s = "abc"
    // Output: []
    fmt.Println(generatePalindromes("abc")) // []

    fmt.Println(generatePalindromes1("aabb")) // ["abba","baab"]
    fmt.Println(generatePalindromes1("abc")) // []
}