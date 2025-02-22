package main

// 3039. Apply Operations to Make String Empty
// You are given a string s.

// Consider performing the following operation until s becomes empty:
//     For every alphabet character from 'a' to 'z', remove the first occurrence of that character in s (if it exists).

// For example, let initially s = "aabcbbca". We do the following operations:
//     1. Remove the underlined characters s = "aabcbbca". The resulting string is s = "abbca".
//     2. Remove the underlined characters s = "abbca". The resulting string is s = "ba".
//     3. Remove the underlined characters s = "ba". The resulting string is s = "".

// Return the value of the string s right before applying the last operation. 
// In the example above, answer is "ba".

// Example 1:
// Input: s = "aabcbbca"
// Output: "ba"
// Explanation: Explained in the statement.

// Input: s = "abcd"
// Output: "abcd"
// Explanation: We do the following operation:
// - Remove the underlined characters s = "abcd". The resulting string is s = "".
// The string just before the last operation is "abcd".

// Constraints:
//     1 <= s.length <= 5 * 10^5
//     s consists only of lowercase English letters.

import "fmt"
import "sort"

func lastNonEmptyString(s string) string {
    count, pos := make([]int, 26, 26), make([]int, 26, 26)
    mx, exist := 0, 0
    for i := 0 ; i < 26; i++ {
        pos[i] = -1
    }
    
    for i := len(s) - 1; i >= 0; i-- {
        if pos[s[i]-'a'] == -1 {
            pos[s[i] - 'a'] = i
            exist++
        }
        count[s[i] - 'a']++
        if count[s[i] - 'a'] > mx {
            mx = count[s[i] - 'a']
        }
    }
    str := ""
    for i := 0; i < 26; i++ {
        if mx == count[i] {
            str += string(i + 'a')
        }
    }
    arr := []byte(str)
    sort.Slice(arr, func(i, j int) bool { 
        return pos[arr[i] - 'a'] < pos[arr[j] - 'a']
    })
    return string(arr)
}

func lastNonEmptyString1(s string) string {
    count, pos := make([]int, 26), make([]int, 26)
    mx := 0
    for i := 0; i < len(s); i++ {
        k := int(s[i] - 'a')
        count[k]++
        pos[k] = i
        mx = max(mx, count[k])
    }
    res := []byte{}
    for i := 0; i < 26; i++ {
        if count[i] == mx {
            res = append(res, 'a' + byte(i))
        }
    }
    sort.Slice(res, func (i, j int) bool{
        return pos[int(res[i] - 'a')] < pos[int(res[j] - 'a')]
    })
    return string(res)
}

func main() {
    // Example 1:
    // Input: s = "aabcbbca"
    // Output: "ba"
    // Explanation: Explained in the statement.
    fmt.Println(lastNonEmptyString("aabcbbca")) // "ba"
    // Input: s = "abcd"
    // Output: "abcd"
    // Explanation: We do the following operation:
    // - Remove the underlined characters s = "abcd". The resulting string is s = "".
    // The string just before the last operation is "abcd".
    fmt.Println(lastNonEmptyString("abcd")) // "abcd"

    fmt.Println(lastNonEmptyString("bluefrog")) // "bluefrog"
    fmt.Println(lastNonEmptyString("leetcode")) // "e"

    fmt.Println(lastNonEmptyString1("aabcbbca")) // "ba"
    fmt.Println(lastNonEmptyString1("abcd")) // "abcd"
    fmt.Println(lastNonEmptyString1("bluefrog")) // "bluefrog"
    fmt.Println(lastNonEmptyString1("leetcode")) // "e"
}