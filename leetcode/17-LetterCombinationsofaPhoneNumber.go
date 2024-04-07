package main

// 17. Letter Combinations of a Phone Number
// Given a string containing digits from 2-9 inclusive, 
// return all possible letter combinations that the number could represent. Return the answer in any order.

// A mapping of digits to letters (just like on the telephone buttons) is given below. 
// Note that 1 does not map to any letters.
//     2 abc
//     3 def
//     4 ghi
//     5 jkl
//     6 mno
//     7 pqs
//     8 tuv
//     9 xyz
//     <img src="https://assets.leetcode.com/uploads/2022/03/15/1200px-telephone-keypad2svg.png" /> 

// Example 1:
// Input: digits = "23"
// Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

// Example 2:
// Input: digits = ""
// Output: []

// Example 3:
// Input: digits = "2"
// Output: ["a","b","c"]

// Constraints:
//     0 <= digits.length <= 4
//     digits[i] is a digit in the range ['2', '9'].

import "fmt"
import "strings"

func letterCombinations(digits string) []string {
    dict := map[rune]string{
        50: "abc",
        51: "def",
        52: "ghi",
        53: "jkl",
        54: "mno",
        55: "pqrs",
        56: "tuv",
        57: "wxyz",
    }
    if len(digits) == 0 {
        return []string{}
    }
    res := []string{""}
    for _, digit := range digits {
        curr := []string{}
        for  _, letter := range dict[digit] {
            for _, pre := range res {
                curr = append(curr, pre + string(letter))
            }
        }
        res = curr
    }
    return res
}

// dfs
func letterCombinations1(digits string) []string {
    res := []string{}
    if len(digits) == 0 {
        return res
    }
    var dfs func(digits string, res *[]string, path []string, index int)
    dfs = func(digits string, res *[]string, path []string, index int)  {
        letterMap := map[string][]string{
            "0": {},
            "1": {},
            "2": {"a", "b", "c"},
            "3": {"d", "e", "f"},
            "4": {"g", "h", "i"},
            "5": {"j", "k", "l"},
            "6": {"m", "n", "o"},
            "7": {"p", "q", "r", "s"},
            "8": {"t", "u", "v"},
            "9": {"w", "x", "y", "z"},
        }
        if index == len(digits) {
            *res = append(*res, strings.Join(path, ""))
            return
        }
        digit := int(digits[index])
        for _, letter := range letterMap[string(digit)] {
            path = append(path, letter)
            dfs(digits, res, path, index+1)
            path = path[:len(path)-1]
        }
    }
    dfs(digits, &res, []string{}, 0)
    return res
}

func main() {
    fmt.Printf("letterCombinations(\"23\") = %v\n", letterCombinations("23")) //  ["ad","ae","af","bd","be","bf","cd","ce","cf"]
    fmt.Printf("letterCombinations(\"1\") = %v\n", letterCombinations("1")) // []
    fmt.Printf("letterCombinations(\"2\") = %v\n", letterCombinations("2")) //  ["a","b","c"]

    fmt.Printf("letterCombinations1(\"23\") = %v\n", letterCombinations1("23")) //  ["ad","ae","af","bd","be","bf","cd","ce","cf"]
    fmt.Printf("letterCombinations1(\"1\") = %v\n", letterCombinations1("1")) // []
    fmt.Printf("letterCombinations1(\"2\") = %v\n", letterCombinations1("2")) //  ["a","b","c"]
}
