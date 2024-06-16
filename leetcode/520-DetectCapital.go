package main

// 520. Detect Capital
// We define the usage of capitals in a word to be right when one of the following cases holds:
//     All letters in this word are capitals, like "USA".
//     All letters in this word are not capitals, like "leetcode".
//     Only the first letter in this word is capital, like "Google".

// Given a string word, return true if the usage of capitals in it is right.

// Example 1:
// Input: word = "USA"
// Output: true

// Example 2:
// Input: word = "FlaG"
// Output: false

// Constraints:
//     1 <= word.length <= 100
//     word consists of lowercase and uppercase English letters.

import "fmt"

func detectCapitalUse(word string) bool {
    if len(word) <= 1 {
        return true
    }
    isUpper := func(c byte) bool {
        return c >= 'A' && c <= 'Z'
    }
    first, second:= isUpper(word[0]), isUpper(word[1]) // 判断首字母是否为大写, 第二个字母是否为大写
    if first { // 首字母为写出现在不为大写字母的返回 false
        for i := 2; i < len(word); i++ {
            if second { // 都必须为大写
                if !isUpper(word[i]) {
                    return false
                }
            } else {
                if isUpper(word[i]) { // 首字母为大写,其它必须为小写
                    return false
                }
            }
        }
    } else {
        for i := 1; i < len(word); i++ {
            if isUpper(word[i]) { // 首字母为小写 其它出现大写字母时
                return false
            }
        }
    }
    return true
}

func detectCapitalUse1(word string) bool {
    n, lowCase, upperCase := len(word), 0, 0
    isUpper := func(c byte) bool {
        return c >= 'A' && c <= 'Z'
    }
    if isUpper(word[0]) { // 首字母为大写
        for i := 1; i < n; i++ {
            if isUpper(word[i]) {
                upperCase++
            } else {
                lowCase++
            }
        }
        return lowCase == n - 1 || upperCase == n - 1
    }
    for i := 1; i < n; i++ {
        if !isUpper(word[i]) {
            lowCase++
        }
    }
    return lowCase == n - 1
}

func main() {
    // Example 1:
    // Input: word = "USA"
    // Output: true
    fmt.Println(detectCapitalUse("USA")) // true
    // Example 2:
    // Input: word = "FlaG"
    // Output: false
    fmt.Println(detectCapitalUse("FlaG")) // false
    fmt.Println(detectCapitalUse("leetcode")) // true
    fmt.Println(detectCapitalUse("Google")) // true
    fmt.Println(detectCapitalUse("mL")) // false


    fmt.Println(detectCapitalUse1("USA")) // true
    fmt.Println(detectCapitalUse1("FlaG")) // false
    fmt.Println(detectCapitalUse1("leetcode")) // true
    fmt.Println(detectCapitalUse1("Google")) // true
    fmt.Println(detectCapitalUse1("mL")) // false
}