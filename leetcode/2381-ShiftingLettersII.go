package main

// 2381. Shifting Letters II
// You are given a string s of lowercase English letters and a 2D integer array shifts where shifts[i] = [starti, endi, directioni]. 
// For every i, shift the characters in s from the index starti to the index endi (inclusive) forward if directioni = 1, or shift the characters backward if directioni = 0.

// Shifting a character forward means replacing it with the next letter in the alphabet (wrapping around so that 'z' becomes 'a'). 
// Similarly, shifting a character backward means replacing it with the previous letter in the alphabet (wrapping around so that 'a' becomes 'z').

// Return the final string after all such shifts to s are applied.

// Example 1:
// Input: s = "abc", shifts = [[0,1,0],[1,2,1],[0,2,1]]
// Output: "ace"
// Explanation: Firstly, shift the characters from index 0 to index 1 backward. Now s = "zac".
// Secondly, shift the characters from index 1 to index 2 forward. Now s = "zbd".
// Finally, shift the characters from index 0 to index 2 forward. Now s = "ace".

// Example 2:
// Input: s = "dztz", shifts = [[0,0,0],[1,1,1]]
// Output: "catz"
// Explanation: Firstly, shift the characters from index 0 to index 0 backward. Now s = "cztz".
// Finally, shift the characters from index 1 to index 1 forward. Now s = "catz".

// Constraints:
//     1 <= s.length, shifts.length <= 5 * 10^4
//     shifts[i].length == 3
//     0 <= starti <= endi < s.length
//     0 <= directioni <= 1
//     s consists of lowercase English letters.

import "fmt"

func shiftingLetters(s string, shifts [][]int) string {
    n := len(s)
    res, prefix := make([]byte, n), make([]int, n)
    for _, shift := range shifts {
        val := -1
        if shift[2] == 1 { val = 1 }
        prefix[shift[0]] += val
        if shift[1] + 1 < n {
            prefix[shift[1] + 1] -= val
        }
    }
    prefix[0] %= 26
    for i := 0; i < n; i++ {
        if i > 0 {
             prefix[i] = (prefix[i] + prefix[i - 1]) % 26
        }
        val := int(s[i]) + prefix[i]
        if val > 'z' { 
            val -= 26 
        } else if val < 'a' { 
            val += 26 
        }
        res[i] = byte(val)
    }
    return string(res)
}

func shiftingLetters1(s string, shifts [][]int) string {
    n, sum := len(s), 0
    res, prefix := []byte(s), make([]int, n + 1)
    for _, v := range shifts {
        if v[2] == 1 {
            prefix[v[0]]++
            prefix[v[1]+1]--
        } else {
            prefix[v[0]]--
            prefix[v[1]+1]++
        }
    }
    for i, v := range res {
        sum += prefix[i]
        t := (int(v - 'a') + sum) % 26
        if t < 0 {
            t += 26
        }
        res[i] = byte(t) + 'a'
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: s = "abc", shifts = [[0,1,0],[1,2,1],[0,2,1]]
    // Output: "ace"
    // Explanation: Firstly, shift the characters from index 0 to index 1 backward. Now s = "zac".
    // Secondly, shift the characters from index 1 to index 2 forward. Now s = "zbd".
    // Finally, shift the characters from index 0 to index 2 forward. Now s = "ace".
    fmt.Println(shiftingLetters("abc", [][]int{{0,1,0},{1,2,1},{0,2,1}})) // "ace"
    // Example 2:
    // Input: s = "dztz", shifts = [[0,0,0],[1,1,1]]
    // Output: "catz"
    // Explanation: Firstly, shift the characters from index 0 to index 0 backward. Now s = "cztz".
    // Finally, shift the characters from index 1 to index 1 forward. Now s = "catz".
    fmt.Println(shiftingLetters("dztz", [][]int{{0,0,0},{1,1,1}})) // "catz"

    fmt.Println(shiftingLetters1("abc", [][]int{{0,1,0},{1,2,1},{0,2,1}})) // "ace"
    fmt.Println(shiftingLetters1("dztz", [][]int{{0,0,0},{1,1,1}})) // "catz"
}