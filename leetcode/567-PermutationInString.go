package main

// 567. Permutation in String
// Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.
// In other words, return true if one of s1's permutations is the substring of s2.

// Example 1:
// Input: s1 = "ab", s2 = "eidbaooo"
// Output: true
// Explanation: s2 contains one permutation of s1 ("ba").

// Example 2:
// Input: s1 = "ab", s2 = "eidboaoo"
// Output: false

// Constraints:
//     1 <= s1.length, s2.length <= 10^4
//     s1 and s2 consist of lowercase English letters.

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