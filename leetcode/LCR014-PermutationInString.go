package main

// LCR 014. 字符串的排列
// 给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的某个变位词。
// 换句话说，第一个字符串的排列之一是第二个字符串的 子串 。

// 示例 1：
// 输入: s1 = "ab" s2 = "eidbaooo"
// 输出: True
// 解释: s2 包含 s1 的排列之一 ("ba").

// 示例 2：
// 输入: s1= "ab" s2 = "eidboaoo"
// 输出: False

// 提示：
//     1 <= s1.length, s2.length <= 10^4
//     s1 和 s2 仅包含小写字母

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
    str1, str2 := make([]int, 26), make([]int, 26)
    for _, char := range(s1) {
        str1[char - 'a'] += 1
    }

    isEqual := func (arr1, arr2 []int) bool {
        if len(arr1) != len(arr2) {
            return false
        }
        for k, v := range arr1 {
            if v != arr2[k] {
                return false
            }
        }
        return true
    }
    
    for i := 0; i < len(s2); i++ {
        str2[s2[i] - 'a'] += 1
        
        if i >= len(s1) {
            str2[s2[i - len(s1)] - 'a'] -= 1
        }
        if isEqual(str1, str2) {
            return true
        }
    }
    return false
}

// 滑动窗口
func checkInclusion1(s1 string, s2 string) bool {
    var freq [256]int
    if len(s2) == 0 || len(s2) < len(s1) {
        return false
    }
    for i := 0; i < len(s1); i++ {
        freq[s1[i]-'a']++
    }
    left, right, count := 0, 0, len(s1)
    for right < len(s2) {
        if freq[s2[right]-'a'] >= 1 {
            count--
        }
        freq[s2[right]-'a']--
        right++
        if count == 0 {
            return true
        }
        if right-left == len(s1) {
            if freq[s2[left]-'a'] >= 0 {
                count++
            }
            freq[s2[left]-'a']++
            left++
        }
    }
    return false
}

func main() {
    // Explanation: s2 contains one permutation of s1 ("ba").
    fmt.Println(checkInclusion("ab","eidbaooo")) // true
    fmt.Println(checkInclusion("ab","eidboaoo")) // false

    fmt.Println(checkInclusion1("ab","eidbaooo")) // true
    fmt.Println(checkInclusion1("ab","eidboaoo")) // false
}