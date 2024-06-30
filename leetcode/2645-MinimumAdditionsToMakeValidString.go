package main

// 2645. Minimum Additions to Make Valid String
// Given a string word to which you can insert letters "a", "b" or "c" anywhere and any number of times, 
// return the minimum number of letters that must be inserted so that word becomes valid.

// A string is called valid if it can be formed by concatenating the string "abc" several times.

// Example 1:
// Input: word = "b"
// Output: 2
// Explanation: Insert the letter "a" right before "b", and the letter "c" right next to "b" to obtain the valid string "abc".

// Example 2:
// Input: word = "aaa"
// Output: 6
// Explanation: Insert letters "b" and "c" next to each "a" to obtain the valid string "abcabcabc".

// Example 3:
// Input: word = "abc"
// Output: 0
// Explanation: word is already valid. No modifications are needed. 

// Constraints:
//     1 <= word.length <= 50
//     word consists of letters "a", "b" and "c" only. 

import "fmt"
import "strings"

// func addMinimum(word string) int {
//     res, mp := 0, [3]int{}
//     for _, c := range word { // 统计 a b c 各出现的次数
//         mp[c - 'a']++
//     }
//     mx := mp[0] 
//     for i := 1; i < len(mp); i++ { // 找出出现最多次数和字母
//         if mp[i] > mx {
//             mx = mp[i]
//         }
//     }
//     abs := func(x int) int { if x < 0 { return -x; }; return x; }
//     for i := 0; i < len(mp); i++ { // 计算需要补全的字符数
//         res += abs(mx - mp[i])
//     }
//     return res
// }

func addMinimum(word string) int {
    word = strings.ReplaceAll(word, "abc", "***")
    res := 0
    pattern := []string{"ab", "bc", "ac", "bc"}
    for _, p := range pattern {
        temp := word
        temp = strings.ReplaceAll(temp, p, "")
        word = strings.ReplaceAll(word, p, "**")
        res += (len(word) - len(temp)) / 2
    }
    word = strings.ReplaceAll(word, "*", "")
    res += 2 * len(word)
    return res
}

func main() {
    // Example 1:
    // Input: word = "b"
    // Output: 2
    // Explanation: Insert the letter "a" right before "b", and the letter "c" right next to "b" to obtain the valid string "abc".
    fmt.Println(addMinimum("b")) // 2
    // Example 2:
    // Input: word = "aaa"
    // Output: 6
    // Explanation: Insert letters "b" and "c" next to each "a" to obtain the valid string "abcabcabc".
    fmt.Println(addMinimum("aaa")) // 6
    // Example 3:
    // Input: word = "abc"
    // Output: 0
    // Explanation: word is already valid. No modifications are needed. 
    fmt.Println(addMinimum("abc")) // 0

    fmt.Println(addMinimum("aaaabb")) // 9
}