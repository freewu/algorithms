package main

// 1805. Number of Different Integers in a String
// You are given a string word that consists of digits and lowercase English letters.

// You will replace every non-digit character with a space. 
// For example, "a123bc34d8ef34" will become " 123  34 8  34". 
// Notice that you are left with some integers that are separated by at least one space: "123", "34", "8", and "34".

// Return the number of different integers after performing the replacement operations on word.

// Two integers are considered different if their decimal representations without any leading zeros are different.

// Example 1:
// Input: word = "a123bc34d8ef34"
// Output: 3
// Explanation: The three different integers are "123", "34", and "8". Notice that "34" is only counted once.

// Example 2:
// Input: word = "leet1234code234"
// Output: 2

// Example 3:
// Input: word = "a1b01c001"
// Output: 1
// Explanation: The three integers "1", "01", and "001" all represent the same integer because
// the leading zeros are ignored when comparing their decimal values.

// Constraints:
//     1 <= word.length <= 1000
//     word consists of digits and lowercase English letters.

import "fmt"
import "regexp"

func numDifferentIntegers(word string) int {
    re := regexp.MustCompile(`0*(\d+)`)
    mp := make(map[string]bool)
    for _, v := range re.FindAllStringSubmatch(word, -1) {
        mp[v[1]] = true  
    }
    return len(mp)
}

func numDifferentIntegers1(word string) int {
    n, start, mp := len(word), 0, make(map[string]bool)
    isDigit := func(ch byte) bool { return ch >= '0' && ch <= '9' }
    for {
        for start < n && !isDigit(word[start]) {
            start++
        }
        if start == n { break }
        end := start
        for end < n && isDigit(word[end]) {
            end++
        }
        for end - start > 1 && word[start] == '0' { // 去除前导 0
            start++
        }
        mp[word[start:end]] = true
        start = end
    }
    return len(mp)
}

func main() {
    // Example 1:
    // Input: word = "a123bc34d8ef34"
    // Output: 3
    // Explanation: The three different integers are "123", "34", and "8". Notice that "34" is only counted once.
    fmt.Println(numDifferentIntegers("a123bc34d8ef34")) // 3
    // Example 2:
    // Input: word = "leet1234code234"
    // Output: 2
    fmt.Println(numDifferentIntegers("leet1234code234")) // 2
    // Example 3:
    // Input: word = "a1b01c001"
    // Output: 1
    // Explanation: The three integers "1", "01", and "001" all represent the same integer because
    // the leading zeros are ignored when comparing their decimal values.
    fmt.Println(numDifferentIntegers("a1b01c001")) // 1

    fmt.Println(numDifferentIntegers1("a123bc34d8ef34")) // 3
    fmt.Println(numDifferentIntegers1("leet1234code234")) // 2
    fmt.Println(numDifferentIntegers1("a1b01c001")) // 1
}