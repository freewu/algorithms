package main

// 1641. Count Sorted Vowel Strings
// Given an integer n, return the number of strings of length n 
// that consist only of vowels (a, e, i, o, u) and are lexicographically sorted.

// A string s is lexicographically sorted if for all valid i, s[i] is the same as or comes before s[i+1] in the alphabet.

// Example 1:
// Input: n = 1
// Output: 5
// Explanation: The 5 sorted strings that consist of vowels only are ["a","e","i","o","u"].

// Example 2:
// Input: n = 2
// Output: 15
// Explanation: The 15 sorted strings that consist of vowels only are
// ["aa","ae","ai","ao","au","ee","ei","eo","eu","ii","io","iu","oo","ou","uu"].
// Note that "ea" is not a valid string since 'e' comes after 'a' in the alphabet.

// Example 3:
// Input: n = 33
// Output: 66045

// Constraints:
//     1 <= n <= 50 

import "fmt"

func countVowelStrings(n int) int {
    res := []int{1, 1, 1, 1, 1}
    sum := func(arr []int) int {
        res := 0
        for _, v := range arr { res += v }
        return res
    }
    for ; n > 1; n-- {
        for i := range res {
            res[i] = sum(res[i:])
        }
    }
    return sum(res)
}

func main() {
    // Example 1:
    // Input: n = 1
    // Output: 5
    // Explanation: The 5 sorted strings that consist of vowels only are ["a","e","i","o","u"].
    fmt.Println(countVowelStrings(1)) // 5 ["a","e","i","o","u"]
    // Example 2:
    // Input: n = 2
    // Output: 15
    // Explanation: The 15 sorted strings that consist of vowels only are
    // ["aa","ae","ai","ao","au","ee","ei","eo","eu","ii","io","iu","oo","ou","uu"].
    // Note that "ea" is not a valid string since 'e' comes after 'a' in the alphabet.
    fmt.Println(countVowelStrings(2)) // 15 ["aa","ae","ai","ao","au","ee","ei","eo","eu","ii","io","iu","oo","ou","uu"].
    // Example 3:
    // Input: n = 33
    // Output: 66045
    fmt.Println(countVowelStrings(33)) // 66045

    fmt.Println(countVowelStrings(50)) // 316251
}