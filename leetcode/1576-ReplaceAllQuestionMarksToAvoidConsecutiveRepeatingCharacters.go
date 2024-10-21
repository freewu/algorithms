package main

// 1576. Replace All ?'s to Avoid Consecutive Repeating Characters
// Given a string s containing only lowercase English letters and the '?' character, 
// convert all the '?' characters into lowercase letters 
// such that the final string does not contain any consecutive repeating characters. 
// You cannot modify the non '?' characters.

// It is guaranteed that there are no consecutive repeating characters in the given string except for '?'.

// Return the final string after all the conversions (possibly zero) have been made. 
// If there is more than one solution, return any of them. 
// It can be shown that an answer is always possible with the given constraints.

// Example 1:
// Input: s = "?zs"
// Output: "azs"
// Explanation: There are 25 solutions for this problem. From "azs" to "yzs", all are valid. Only "z" is an invalid modification as the string will consist of consecutive repeating characters in "zzs".

// Example 2:
// Input: s = "ubv?w"
// Output: "ubvaw"
// Explanation: There are 24 solutions for this problem. Only "v" and "w" are invalid modifications as the strings will consist of consecutive repeating characters in "ubvvw" and "ubvww".

// Constraints:
//     1 <= s.length <= 100
//     s consist of lowercase English letters and '?'.

import "fmt"

func modifyString(s string) string {
    res := []byte(s)
    for i := 0; i < len(res); i++ {
        if res[i] == '?' {
            for j := byte(97); j < 100; j++ {
                if (i == 0 || res[i - 1] != j) && (i == len(s) - 1 || res[i + 1] != j) {
                    res[i] = j
                    break
                }
            }
        }
    }
    return string(res)
}

func modifyString1(s string) string {
    res := []byte(s)
    get := func(i int) byte {
        j := 0
        for ; j < 26; {
            ch := byte(j)+'a'
            if (i-1 >= 0 && res[i-1] == ch ) || (i+1 < len(s) && s[i+1] == ch ) {
                j++
            } else {
                break
            }
        }
        return byte(j) + 'a'
    }
    for i := range res {
        if res[i] == '?' {
            res[i] = get(i)
        }
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: s = "?zs"
    // Output: "azs"
    // Explanation: There are 25 solutions for this problem. From "azs" to "yzs", all are valid. Only "z" is an invalid modification as the string will consist of consecutive repeating characters in "zzs".
    fmt.Println(modifyString("?zs")) // "azs"
    // Example 2:
    // Input: s = "ubv?w"
    // Output: "ubvaw"
    // Explanation: There are 24 solutions for this problem. Only "v" and "w" are invalid modifications as the strings will consist of consecutive repeating characters in "ubvvw" and "ubvww".
    fmt.Println(modifyString("ubv?w")) // "ubvaw"

    fmt.Println(modifyString("ubv?w???")) // "ubvawaba"

    fmt.Println(modifyString1("?zs")) // "azs"
    fmt.Println(modifyString1("ubv?w")) // "ubvaw"
    fmt.Println(modifyString1("ubv?w???")) // "ubvawaba"
}

