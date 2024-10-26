package main

// 1624. Largest Substring Between Two Equal Characters
// Given a string s, return the length of the longest substring between two equal characters, excluding the two characters. 
// If there is no such substring return -1.

// A substring is a contiguous sequence of characters within a string.

// Example 1:
// Input: s = "aa"
// Output: 0
// Explanation: The optimal substring here is an empty substring between the two 'a's.

// Example 2:
// Input: s = "abca"
// Output: 2
// Explanation: The optimal substring here is "bc".

// Example 3:
// Input: s = "cbzxy"
// Output: -1
// Explanation: There are no characters that appear twice in s.

// Constraints:
//     1 <= s.length <= 300
//     s contains only lowercase English letters.

import "fmt"

func maxLengthBetweenEqualCharacters(s string) int {
    for i := len(s); i > 1; i-- {
        for j := 0; j < len(s) - i + 1; j++ {
            if s[j] == s[j + i - 1] {
                return i - 2
            }
        }
    }
    return -1
}

func maxLengthBetweenEqualCharacters1(s string) int {
    res, mp := -1, make(map[byte][]int)
    for i := 0; i < len(s); i++ {
        mp[s[i]] = append(mp[s[i]], i)
    }
    for _, v := range mp {
        if len(v) >= 2 {
            diff := v[len(v) - 1] - v[0] - 1
            if diff > res {
                res = diff
            } 
        }
    }
    return res
}

func maxLengthBetweenEqualCharacters2(s string) int {
    res, mp := -1, make(map[byte]int)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < len(s); i++ {
        if _, ok := mp[s[i]]; !ok {
            mp[s[i]] = i // 只记录第一次出现的位置
        } else {
            res = max(res, i - mp[s[i]] - 1)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "aa"
    // Output: 0
    // Explanation: The optimal substring here is an empty substring between the two 'a's.
    fmt.Println(maxLengthBetweenEqualCharacters("aa")) // 0
    // Example 2:
    // Input: s = "abca"
    // Output: 2
    // Explanation: The optimal substring here is "bc".
    fmt.Println(maxLengthBetweenEqualCharacters("abca")) // 2
    // Example 3:
    // Input: s = "cbzxy"
    // Output: -1
    // Explanation: There are no characters that appear twice in s.
    fmt.Println(maxLengthBetweenEqualCharacters("cbzxy")) // -1
    fmt.Println(maxLengthBetweenEqualCharacters("cabbac")) // 4
    fmt.Println(maxLengthBetweenEqualCharacters("mgntdygtxrvxjnwksqhxuxtrv")) // 18

    fmt.Println(maxLengthBetweenEqualCharacters1("aa")) // 0
    fmt.Println(maxLengthBetweenEqualCharacters1("abca")) // 2
    fmt.Println(maxLengthBetweenEqualCharacters1("cbzxy")) // -1
    fmt.Println(maxLengthBetweenEqualCharacters1("cabbac")) // 4
    fmt.Println(maxLengthBetweenEqualCharacters1("mgntdygtxrvxjnwksqhxuxtrv")) // 18

    fmt.Println(maxLengthBetweenEqualCharacters2("aa")) // 0
    fmt.Println(maxLengthBetweenEqualCharacters2("abca")) // 2
    fmt.Println(maxLengthBetweenEqualCharacters2("cbzxy")) // -1
    fmt.Println(maxLengthBetweenEqualCharacters2("cabbac")) // 4
    fmt.Println(maxLengthBetweenEqualCharacters2("mgntdygtxrvxjnwksqhxuxtrv")) // 18
}