package main

// 3121. Count the Number of Special Characters II
// You are given a string word. 
// A letter c is called special if it appears both in lowercase and uppercase in word, 
// and every lowercase occurrence of c appears before the first uppercase occurrence of c.

// Return the number of special letters in word.

// Example 1:
// Input: word = "aaAbcBC"
// Output: 3
// Explanation:
// The special characters are 'a', 'b', and 'c'.

// Example 2:
// Input: word = "abc"
// Output: 0
// Explanation:
// There are no special characters in word.

// Example 3:
// Input: word = "AbBCab"
// Output: 0
// Explanation:
// There are no special characters in word.

// Constraints:
//     1 <= word.length <= 2 * 10^5
//     word consists of only lowercase and uppercase English letters

import "fmt"
import "math/bits"

func numberOfSpecialChars1(word string) int {
    res, mp := 0, make(map[byte]int)
    for i := range word {
        if word[i] >= 'A' && word[i] <= 'Z' {
            if _, ok := mp[word[i]]; !ok { // 只记录第一个出现的位置
                mp[word[i]] = i
            }
        } else {
            mp[word[i]] = i
        }
    }
    for i := 0; i < 26; i++ {
        v1, ok1 := mp[byte(i + 'a')]
        v2, ok2 := mp[byte(i + 'A')]
        if ok1 && ok2 && v2 > v1 { // 小写必须出现在大写之前
            res++
        }
    }
    return res
}

func numberOfSpecialChars(word string) int {
    res, mpLower, mpUpper := 0, make(map[rune]int), make(map[rune]int)
    for i, char := range word {
        if char >= 'A' && char <= 'Z' {
            if _, ok := mpUpper[char]; ok { // 记录第一个大写
                continue
            } else {
                mpUpper[char] = i
            }
        } else {
            mpLower[char] = i
        }
    }
    for k, v := range mpLower {
        if v < mpUpper[k - 32] {
            res++
        }
    }
    return res
}

func numberOfSpecialChars2(word string) int {
    res, mp := 0, [52]bool{}
    for i := 0; i < len(word); i++ {
        b := word[i]
        if b >= 'a' && b <= 'z' {
            mp[b-'a'] = !mp[b - 'a' + 26]
        }
        if b >= 'A' && b <= 'Z' {
            mp[b - 'A' + 26] = true
        }
    }
    for i := 0; i < 26; i++ {
        if mp[i] && mp[i + 26] {
            res++
        }
    }
    return res
}

func numberOfSpecialChars3(word string) int {
    lower, upper, invalid := uint(0), uint(0), uint(0)
    for _, c := range word {
        bit := uint(1) << (c & 31)
        if c & 32 > 0 {
            lower |= bit
            if upper&bit > 0 {
                invalid |= bit
            }
        } else {
            upper |= bit
        }
    }
    return bits.OnesCount(lower & upper &^ invalid)
}

func main() {
    // Example 1:
    // Input: word = "aaAbcBC"
    // Output: 3
    // Explanation:
    // The special characters are 'a', 'b', and 'c'.
    fmt.Println(numberOfSpecialChars("aaAbcBC")) // 3
    // Example 2:
    // Input: word = "abc"
    // Output: 0
    // Explanation:
    // There are no special characters in word.
    fmt.Println(numberOfSpecialChars("abc")) // 0
    // Example 3:
    // Input: word = "AbBCab"
    // Output: 0
    // Explanation:
    // There are no special characters in word.
    fmt.Println(numberOfSpecialChars("AbBCab")) // 0

    fmt.Println(numberOfSpecialChars("cCceDC")) // 0

    fmt.Println(numberOfSpecialChars1("aaAbcBC")) // 3
    fmt.Println(numberOfSpecialChars1("abc")) // 0
    fmt.Println(numberOfSpecialChars1("AbBCab")) // 0
    fmt.Println(numberOfSpecialChars1("cCceDC")) // 0

    fmt.Println(numberOfSpecialChars2("aaAbcBC")) // 3
    fmt.Println(numberOfSpecialChars2("abc")) // 0
    fmt.Println(numberOfSpecialChars2("AbBCab")) // 0
    fmt.Println(numberOfSpecialChars2("cCceDC")) // 0

    fmt.Println(numberOfSpecialChars3("aaAbcBC")) // 3
    fmt.Println(numberOfSpecialChars3("abc")) // 0
    fmt.Println(numberOfSpecialChars3("AbBCab")) // 0
    fmt.Println(numberOfSpecialChars3("cCceDC")) // 0
}