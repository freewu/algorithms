package main

// 1309. Decrypt String from Alphabet to Integer Mapping
// You are given a string s formed by digits and '#'. 
// We want to map s to English lowercase characters as follows:
//     Characters ('a' to 'i') are represented by ('1' to '9') respectively.
//     Characters ('j' to 'z') are represented by ('10#' to '26#') respectively.

// Return the string formed after mapping.

// The test cases are generated so that a unique mapping will always exist.

// Example 1:
// Input: s = "10#11#12"
// Output: "jkab"
// Explanation: "j" -> "10#" , "k" -> "11#" , "a" -> "1" , "b" -> "2".

// Example 2:
// Input: s = "1326#"
// Output: "acz"

// Constraints:
//     1 <= s.length <= 1000
//     s consists of digits and the '#' letter.
//     s will be a valid string such that mapping is always possible.

import "fmt"
import "strconv"

func freqAlphabets(s string) string {
    res, index := []byte{}, len(s) - 1
    for index >= 0 {
        newRes := make([]byte, len(res) + 1)
        copy(newRes[1:], res)
        res = newRes 
        if index > 1 && s[index] == '#' { // 取两位 ('j' to 'z') ('10#' to '26#')
            num, _ := strconv.Atoi(string(s[index - 2:index]))
            res[0] = byte(num + 96)
            index -= 3
        } else { // ('a' to 'i') ('1' to '9')
            num, _ := strconv.Atoi(string(s[index]))
            res[0] = byte(num + 96)
            index--
        }
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: s = "10#11#12"
    // Output: "jkab"
    // Explanation: "j" -> "10#" , "k" -> "11#" , "a" -> "1" , "b" -> "2".
    fmt.Println(freqAlphabets("10#11#12")) // "jkab"
    // Example 2:
    // Input: s = "1326#"
    // Output: "acz"
    fmt.Println(freqAlphabets("1326#")) // "acz"
}