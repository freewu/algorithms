package main

// LCR 032. 有效的字母异位词
// 给定两个字符串 s 和 t ，编写一个函数来判断它们是不是一组变位词（字母异位词）。
// 注意：若 s 和 t 中每个字符出现的次数都相同且字符顺序不完全相同，则称 s 和 t 互为变位词（字母异位词）。

// 示例 1:
// 输入: s = "anagram", t = "nagaram"
// 输出: true

// 示例 2:
// 输入: s = "rat", t = "car"
// 输出: false

// 示例 3:
// 输入: s = "a", t = "a"
// 输出: false

// 提示:
//     1 <= s.length, t.length <= 5 * 10^4

// 进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

import "fmt"

// 支持 utf8
func isAnagram(s string, t string) bool {
    if len(t) != len(s) || t == s {
        return false
    }
    // 逐字写到 map中
    ms, mt := make(map[rune]int), make(map[rune]int)
    for _, v := range s {
        ms[v]++
    }
    for _, v := range t {
        mt[v]++
    }
    for k,v := range ms {
        if v != mt[k] { // 数量出现不一,返回 false
            return false
        }
    }
    return true
}

// best solution
func isAnagram1(s string, t string) bool {
    if len(s) != len(t) || t == s {
        return false
    }
    cnt := make([]int, 26)
    for i := range s {
        cnt[s[i] - 'a'] ++
        cnt[t[i] - 'a'] --
    }
    for i := range cnt {
        if cnt[i] != 0 {
            return false
        }
    }
    return true
}

func main() {
    fmt.Println(isAnagram("anagram","nagaram")) // true
    fmt.Println(isAnagram("rat","car")) // false
    fmt.Println(isAnagram("a","a")) // false

    fmt.Println(isAnagram1("anagram","nagaram")) // true
    fmt.Println(isAnagram1("rat","car")) // false
    fmt.Println(isAnagram1("a","a")) // false
}