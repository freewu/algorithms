package main

// 1737. Change Minimum Characters to Satisfy One of Three Conditions
// You are given two strings a and b that consist of lowercase letters. 
// In one operation, you can change any character in a or b to any lowercase letter.

// Your goal is to satisfy one of the following three conditions:
//     Every letter in a is strictly less than every letter in b in the alphabet.
//     Every letter in b is strictly less than every letter in a in the alphabet.
//     Both a and b consist of only one distinct letter.

// Return the minimum number of operations needed to achieve your goal.

// Example 1:
// Input: a = "aba", b = "caa"
// Output: 2
// Explanation: Consider the best way to make each condition true:
// 1) Change b to "ccc" in 2 operations, then every letter in a is less than every letter in b.
// 2) Change a to "bbb" and b to "aaa" in 3 operations, then every letter in b is less than every letter in a.
// 3) Change a to "aaa" and b to "aaa" in 2 operations, then a and b consist of one distinct letter.
// The best way was done in 2 operations (either condition 1 or condition 3).

// Example 2:
// Input: a = "dabadd", b = "cda"
// Output: 3
// Explanation: The best way is to make condition 1 true by changing b to "eee".

// Constraints:
//     1 <= a.length, b.length <= 10^5
//     a and b consist only of lowercase letters.

import "fmt"

func minCharacters(a string, b string) int {
    m, n := len(a), len(b)
    res, count1, count2 := m + n, [26]int{}, [26]int{}
    for _, v := range a { count1[v - 'a']++ }
    for _, v := range b { count2[v - 'a']++ }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < 26; i++ {
        res = min(res, m + n - count1[i] - count2[i])
        if (i != 0) {
            count1[i] += count1[i-1]
            count2[i] += count2[i-1]
        }
        if (i != 25) {
            res = min(min(res, n - count2[i] + count1[i]), m - count1[i] + count2[i])
        }
    }
    return res
}

func minCharacters1(a string, b string) int {
    m, n := len(a), len(b)
    count1, count2 := [26]int{}, [26]int{}
    for _, v := range a { count1[v - 'a']++ }
    for _, v := range b { count2[v - 'a']++ }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    res, x1, y1, x2, y2 := m + n - count1[0] - count2[0], m, 0, 0, n
    for i := 1; i < 26; i++ {
        x1 -= count1[i-1]
        y1 += count2[i-1]
        x2 += count1[i-1]
        y2 -= count2[i-1]
        res = min(min(res, m + n - count1[i] - count2[i]), min(x1 + y1, x2 + y2))
    }
    return res
}

func main() {
    // Example 1:
    // Input: a = "aba", b = "caa"
    // Output: 2
    // Explanation: Consider the best way to make each condition true:
    // 1) Change b to "ccc" in 2 operations, then every letter in a is less than every letter in b.
    // 2) Change a to "bbb" and b to "aaa" in 3 operations, then every letter in b is less than every letter in a.
    // 3) Change a to "aaa" and b to "aaa" in 2 operations, then a and b consist of one distinct letter.
    // The best way was done in 2 operations (either condition 1 or condition 3).
    fmt.Println(minCharacters("aba", "caa")) // 2
    // Example 2:
    // Input: a = "dabadd", b = "cda"
    // Output: 3
    // Explanation: The best way is to make condition 1 true by changing b to "eee".
    fmt.Println(minCharacters("dabadd", "cda")) // 3

    fmt.Println(minCharacters1("aba", "caa")) // 2
    fmt.Println(minCharacters1("dabadd", "cda")) // 3
}