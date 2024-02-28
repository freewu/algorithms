package main

// 面试题 01.04. Palindrome Permutation LCCI
// Given a string, write a function to check if it is a permutation of a palindrome. 
// A palindrome is a word or phrase that is the same forwards and backwards. 
// A permutation is a rearrangement of letters. The palindrome does not need to be limited to just dictionary words.

// Example1:
// Input: "tactcoa"
// Output: true（permutations: "tacocat"、"atcocta", etc.）

import "fmt"

func canPermutePalindrome(s string) bool {
    m := make(map[rune]int,26)
    // 先转成 map
    for _,v := range s {
        m[v]++
        //fmt.Println(m)
    }
    // 为 0 说明偶数个 出现奇数就可以返回 false
    // 为 1 说明奇数个 出现个奇数先 -1 再次出现奇数返回 false
    flag := len(s) % 2
    for _,v := range m {
        // 出现奇数
        if v % 2 == 1 {
            if flag == 0 {
                return false
            } else {
                flag--
            }
        }
    }
    return true
}

func main() {
    fmt.Println(canPermutePalindrome("tactcoa")) // true
    fmt.Println(canPermutePalindrome("tacocat")) // true
    fmt.Println(canPermutePalindrome("atcocta")) // true
    fmt.Println(canPermutePalindrome("atcoctaa")) // false
}