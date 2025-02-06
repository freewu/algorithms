package main

// 面试题 16.20. T9 LCCI
// On old cell phones, users typed on a numeric keypad and the phone would provide a list of words that matched these numbers. 
// Each digit mapped to a set of 0 - 4 letters. 
// Implement an algo­rithm to return a list of matching words, given a sequence of digits. 
// You are provided a list of valid words. 
// The mapping is shown in the diagram below:

// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/original_images/17_telephone_keypad.png" />

// Example 1:
// Input: num = "8733", words = ["tree", "used"]
// Output: ["tree", "used"]

// Example 2:
// Input: num = "2", words = ["a", "b", "c", "d"]
// Output: ["a", "b", "c"]

// Note:
//     num.length <= 1000
//     words.length <= 500
//     words[i].length == num.length
//     There are no number 0 and 1 in num.

import "fmt"

func getValidT9Words(num string, words []string) []string {
    mp := map[byte]byte{
        'a': '2', 'b': '2', 'c': '2',
        'd': '3', 'e': '3', 'f': '3',
        'g': '4', 'h': '4', 'i': '4',
        'j': '5', 'k': '5', 'l': '5',
        'm': '6', 'n': '6', 'o': '6',
        'p': '7', 'q': '7', 'r': '7', 's': '7',
        't': '8', 'u': '8', 'v': '8',
        'w': '9', 'x': '9', 'y': '9', 'z': '9',
    }
    res, n, m := make([]string,0,10), len(num), len(words)
    for i := 0; i < m; i++ {
        if len(words[i]) != n { continue }
        for j := 0; j < n; j++ {
            if mp[words[i][j]] != num[j] { break }
            if j == n - 1 {
                res = append(res, words[i])
            }
        }
    }
    return res
}

func getValidT9Words1(num string, words []string) []string {
    n, res := len(num), []string{}
    mp := [26]int{ 2,2,2,3,3,3,4,4,4,5,5,5,6,6,6,7,7,7,7,8,8,8,9,9,9,9 }
    for _, word := range words {
        i := 0
        for i < n {
            if mp[word[i] - 'a'] == int(num[i] - '0') {
                i++
            } else {
                break
            }
        }
        if i == n {
            res = append(res, word)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: num = "8733", words = ["tree", "used"]
    // Output: ["tree", "used"]
    fmt.Println(getValidT9Words("8733", []string{"tree", "used"})) // ["tree", "used"]
    // Example 2:
    // Input: num = "2", words = ["a", "b", "c", "d"]
    // Output: ["a", "b", "c"]
    fmt.Println(getValidT9Words("2", []string{"a", "b", "c", "d"})) // ["a", "b", "c"]

    fmt.Println(getValidT9Words("25833764", []string{"bluefrog", "leetcode"})) // ["bluefrog"]

    fmt.Println(getValidT9Words1("8733", []string{"tree", "used"})) // ["tree", "used"]
    fmt.Println(getValidT9Words1("2", []string{"a", "b", "c", "d"})) // ["a", "b", "c"]
    fmt.Println(getValidT9Words1("25833764", []string{"bluefrog", "leetcode"})) // ["bluefrog"]
}