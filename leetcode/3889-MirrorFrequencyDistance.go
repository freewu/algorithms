package main

// 3889. Mirror Frequency Distance
// You are given a string s consisting of lowercase English letters and digits.

// For each character, its mirror character is defined by reversing the order of its character set:
//     1. For letters, the mirror of a character is the letter at the same position from the end of the alphabet.
//         1.1 For example, the mirror of 'a' is 'z', and the mirror of 'b' is 'y', and so on.
//     2. For digits, the mirror of a character is the digit at the same position from the end of the range '0' to '9'.
//         2.1 For example, the mirror of '0' is '9', and the mirror of '1' is '8', and so on.

// For each unique character c in the string:
//     1. Let m be its mirror character.
//     2. Let freq(x) denote the number of times character x appears in the string.
//     3. Compute the absolute difference between their frequencies, defined as: |freq(c) - freq(m)|

// The mirror pairs (c, m) and (m, c) are the same and must be counted only once.

// Return an integer denoting the total sum of these values over all such distinct mirror pairs.

// Example 1:
// Input: s = "ab1z9"
// Output: 3
// Explanation:
// For every mirror pair:
// c   | m     | freq(c) | freq(m)	| |freq(c) - freq(m)|   |    
// a	| z	    | 1	      | 1	    | 0                     |
// b	| y	    | 1	      | 1	    | 0                     |
// 1	| 8	    | 1	      | 0       | 1                     |
// 9	| 9	    | 0	      | 1       | 1                     |
// Thus, the answer is 0 + 1 + 1 + 1 = 3.

// Example 2:
// Input: s = "4m7n"
// Output: 2
// Explanation:
// c   | m     | freq(c) | freq(m)	| |freq(c) - freq(m)|    |    
// 4	| 5	    | 1	      | 0	    |  1                     |
// m	| n	    | 1	      | 1	    |  1                     |
// 7	| 2	    | 1	      | 0	    |  1                     |
// Thus, the answer is 1 + 0 + 1 = 2.​​​​​​​

// Example 3:
// Input: s = "byby"
// Output: 0
// Explanation:
// c   | m     | freq(c) | freq(m)	    | |freq(c) - freq(m)|    |    
// b	| y	    | 2	      | 2	        |  0                     |
// Thus, the answer is 0.

// Constraints:
//     1 <= s.length <= 5 * 10^5
//     s consists only of lowercase English letters and digits.

import "fmt"
import "unicode"

func mirrorFrequency(s string) int {
    res, mp := 0, make(map[string]bool)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 0; i < len(s); i++ {
        original, mirror := string(s[i]), ""
        if s[i] >= '0' && s[i] <= '9' {
            mirror = string('9' - (rune(s[i]) - '0'))
        } else if s[i] >= 'a' && s[i] <= 'z' {
            mirror = string('z' - (rune(s[i]) - 'a'))
        }
        if mp[original] || mp[mirror] {
            continue
        }
        x, y := 0, 0
        for j := 0; j < len(s); j++ {
            if string(s[j]) == mirror {
                y++
            }
            if string(s[j]) == original {
                x++
            }
        }
        if x > 0 {
            mp[original] = true
        }
        if y > 0 {
            mp[mirror] = true
        }
        res += abs(x - y)
    }
    return res
}

func mirrorFrequency1(s string) int {
    res, freq1, freq2 := 0, make([]int, 26), make([]int, 10)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for _, v := range s {
        if(unicode.IsDigit(v)){
            freq2[v - '0']++
        } else {
            freq1[v - 'a']++  
        }
    }
    for i, j := 0, 25; i < j; i, j = i+1, j-1 {
        res += abs(freq1[i] - freq1[j])
    }

    for i, j := 0, 9; i < j; i, j = i+1, j-1 {
        res += abs(freq2[i] - freq2[j])
    }
    return res
}

func mirrorFrequency2(s string) int {
    res, mp := 0, ['z' + 1]int{}
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for _, ch := range s {
        mp[ch]++
    }
    for i := range 13 {
        res += abs(mp['a' + i] - mp['z' - i])
    }
    for i := range 5 {
        res += abs(mp['0' + i] - mp['9' - i])
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "ab1z9"
    // Output: 3
    // Explanation:
    // For every mirror pair:
    // c   | m     | freq(c) | freq(m)	| |freq(c) - freq(m)|   |    
    // a	| z	    | 1	      | 1	    | 0                     |
    // b	| y	    | 1	      | 1	    | 0                     |
    // 1	| 8	    | 1	      | 0       | 1                     |
    // 9	| 9	    | 0	      | 1       | 1                     |
    // Thus, the answer is 0 + 1 + 1 + 1 = 3.
    fmt.Println(mirrorFrequency("ab1z9")) // 3
    // Example 2:
    // Input: s = "4m7n"
    // Output: 2
    // Explanation:
    // c   | m     | freq(c) | freq(m)	| |freq(c) - freq(m)|    |    
    // 4	| 5	    | 1	      | 0	    |  1                     |
    // m	| n	    | 1	      | 1	    |  1                     |
    // 7	| 2	    | 1	      | 0	    |  1                     |
    // Thus, the answer is 1 + 0 + 1 = 2.​​​​​​​
    fmt.Println(mirrorFrequency("4m7n")) // 2
    // Example 3:
    // Input: s = "byby"
    // Output: 0
    // Explanation:
    // c   | m     | freq(c) | freq(m)	    | |freq(c) - freq(m)|    |    
    // b	| y	    | 2	      | 2	        |  0                     |
    // Thus, the answer is 0.
    fmt.Println(mirrorFrequency("byby")) // 0

    fmt.Println(mirrorFrequency("bluefrog")) // 4
    fmt.Println(mirrorFrequency("leetcode")) // 6
    fmt.Println(mirrorFrequency("freewu")) // 4

    fmt.Println(mirrorFrequency1("ab1z9")) // 3
    fmt.Println(mirrorFrequency1("4m7n")) // 2
    fmt.Println(mirrorFrequency1("byby")) // 0
    fmt.Println(mirrorFrequency1("bluefrog")) // 4
    fmt.Println(mirrorFrequency1("leetcode")) // 6
    fmt.Println(mirrorFrequency1("freewu")) // 4

    fmt.Println(mirrorFrequency2("ab1z9")) // 3
    fmt.Println(mirrorFrequency2("4m7n")) // 2
    fmt.Println(mirrorFrequency2("byby")) // 0
    fmt.Println(mirrorFrequency2("bluefrog")) // 4
    fmt.Println(mirrorFrequency2("leetcode")) // 6
    fmt.Println(mirrorFrequency2("freewu")) // 4
}