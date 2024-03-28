package main

// 859. Buddy Strings
// Given two strings s and goal, return true if you can swap two letters in s so the result is equal to goal, otherwise, return false.
// Swapping letters is defined as taking two indices i and j (0-indexed) such that i != j and swapping the characters at s[i] and s[j].

// For example, swapping at indices 0 and 2 in "abcd" results in "cbad".
 
// Example 1:
// Input: s = "ab", goal = "ba"
// Output: true
// Explanation: You can swap s[0] = 'a' and s[1] = 'b' to get "ba", which is equal to goal.

// Example 2:
// Input: s = "ab", goal = "ab"
// Output: false
// Explanation: The only letters you can swap are s[0] = 'a' and s[1] = 'b', which results in "ba" != goal.

// Example 3:
// Input: s = "aa", goal = "aa"
// Output: true
// Explanation: You can swap s[0] = 'a' and s[1] = 'a' to get "aa", which is equal to goal.
 
// Constraints:
//     1 <= s.length, goal.length <= 2 * 10^4
//     s and goal consist of lowercase letters.

import "fmt"

func buddyStrings(s string, goal string) bool {
    if len(s) != len(goal) {
        return false
    }
    if s == goal {
        m := make(map[byte]int)
        // 有同一个字符出现两次
        for i := 0; i < len(s); i++ {
            m[s[i]]++
            if m[s[i]] == 2 {
                return true
            }
        }
        return false
    } else { // 出现差异的字符为两个字母
        r := []int{}
        for i := 0; i < len(s); i++ {
            if s[i] != goal[i] {
                r = append(r,i)
            }
        }
        if len(r) != 2 { // 异位超过2个 说明不是
            return false
        }
        // 判断交换后的字符是否相等
        if s[r[0]] != goal[r[1]] || s[r[1]] != goal[r[0]] {
            return false
        }
    }
    return true
}

// best solution
func buddyStrings1(s string, goal string) bool {
    if len(s) == 1 || len(s) != len(goal) {
        return false
    }
    index := 0
    ls := [2]int{0, 0}
    res := false
    ha := make(map[byte]bool)
    for i := 0; i < len(s); i++ {
        if s[i] != goal[i] {
            if index > 1 {
                return false
            }
            ls[index] = i
            index++
        } else {
            if ha[s[i]] {
                res = true
            } else {
                ha[s[i]] = true
            }
        }
    }
    if index == 1 {
        return false
    }
    if index == 0 {
        return res
    }
    return s[ls[0]] == goal[ls[1]] && s[ls[1]] == goal[ls[0]]
}

func main() {
    // Explanation: You can swap s[0] = 'a' and s[1] = 'b' to get "ba", which is equal to goal.
    fmt.Println(buddyStrings("ab","ba")) // true
    // Explanation: The only letters you can swap are s[0] = 'a' and s[1] = 'b', which results in "ba" != goal.
    fmt.Println(buddyStrings("ab","ab")) // false
    // Explanation: You can swap s[0] = 'a' and s[1] = 'a' to get "aa", which is equal to goal.
    fmt.Println(buddyStrings("aa","aa")) // true
    // For example, swapping at indices 0 and 2 in "abcd" results in "cbad".
    fmt.Println(buddyStrings("abcd","cbad")) // true

    // Explanation: You can swap s[0] = 'a' and s[1] = 'b' to get "ba", which is equal to goal.
    fmt.Println(buddyStrings1("ab","ba")) // true
    // Explanation: The only letters you can swap are s[0] = 'a' and s[1] = 'b', which results in "ba" != goal.
    fmt.Println(buddyStrings1("ab","ab")) // false
    // Explanation: You can swap s[0] = 'a' and s[1] = 'a' to get "aa", which is equal to goal.
    fmt.Println(buddyStrings1("aa","aa")) // true
    // For example, swapping at indices 0 and 2 in "abcd" results in "cbad".
    fmt.Println(buddyStrings1("abcd","cbad")) // true
}