package main

// 821. Shortest Distance to a Character
// Given a string s and a character c that occurs in s, 
// return an array of integers answer where answer.length == s.length and answer[i] is the distance from index i to the closest occurrence of character c in s.

// The distance between two indices i and j is abs(i - j), where abs is the absolute value function.

// Example 1:
// Input: s = "loveleetcode", c = "e"
// Output: [3,2,1,0,1,0,0,1,2,2,1,0]
// Explanation: The character 'e' appears at indices 3, 5, 6, and 11 (0-indexed).
// The closest occurrence of 'e' for index 0 is at index 3, so the distance is abs(0 - 3) = 3.
// The closest occurrence of 'e' for index 1 is at index 3, so the distance is abs(1 - 3) = 2.
// For index 4, there is a tie between the 'e' at index 3 and the 'e' at index 5, but the distance is still the same: abs(4 - 3) == abs(4 - 5) = 1.
// The closest occurrence of 'e' for index 8 is at index 6, so the distance is abs(8 - 6) = 2.

// Example 2:
// Input: s = "aaab", c = "b"
// Output: [3,2,1,0]

// Constraints:
//     1 <= s.length <= 10^4
//     s[i] and c are lowercase English letters.
//     It is guaranteed that c occurs at least once in s.

import "fmt"

func shortestToChar(s string, c byte) []int {
    n := len(s)
    res, pos := make([]int, n), -n
    for i := range res { 
        res[i] = n 
    }
    for i := 0; i < n; i++ {
        if s[i] == c { 
            pos = i 
        }
        res[i] = i - pos
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := pos - 1; i >= 0; i-- {
        if s[i] == c { 
            pos = i 
        }
        res[i] = min(res[i], pos - i)
    }
    return res
}

func shortestToChar1(s string, c byte) []int {
    idx, n := []int{}, len(s)
    for i := range s { // 找出所有 c 出现的位置
        if s[i] == c {
            idx = append(idx, i)
        }
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    res := make([]int, n)
    for i := range s {
        if s[i] != c {
            t := n
            for _, j := range idx { // 找离 c 最近的距离
                l := abs(i - j)
                if l < t {
                    t = l
                }
            }
            res[i] = t
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "loveleetcode", c = "e"
    // Output: [3,2,1,0,1,0,0,1,2,2,1,0]
    // Explanation: The character 'e' appears at indices 3, 5, 6, and 11 (0-indexed).
    // The closest occurrence of 'e' for index 0 is at index 3, so the distance is abs(0 - 3) = 3.
    // The closest occurrence of 'e' for index 1 is at index 3, so the distance is abs(1 - 3) = 2.
    // For index 4, there is a tie between the 'e' at index 3 and the 'e' at index 5, but the distance is still the same: abs(4 - 3) == abs(4 - 5) = 1.
    // The closest occurrence of 'e' for index 8 is at index 6, so the distance is abs(8 - 6) = 2.
    fmt.Println(shortestToChar("loveleetcode", 'e')) // [3,2,1,0,1,0,0,1,2,2,1,0]
    // Example 2:
    // Input: s = "aaab", c = "b"
    // Output: [3,2,1,0]
    fmt.Println(shortestToChar("aaab", 'b')) // [3,2,1,0]

    fmt.Println(shortestToChar1("loveleetcode", 'e')) // [3,2,1,0,1,0,0,1,2,2,1,0]
    fmt.Println(shortestToChar1("aaab", 'b')) // [3,2,1,0]
}