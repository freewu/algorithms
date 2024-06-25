package main

// 953. Verifying an Alien Dictionary
// In an alien language, surprisingly, they also use English lowercase letters, but possibly in a different order. 
// The order of the alphabet is some permutation of lowercase letters.

// Given a sequence of words written in the alien language, and the order of the alphabet, 
// return true if and only if the given words are sorted lexicographically in this alien language.

// Example 1:
// Input: words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
// Output: true
// Explanation: As 'h' comes before 'l' in this language, then the sequence is sorted.

// Example 2:
// Input: words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
// Output: false
// Explanation: As 'd' comes after 'l' in this language, then words[0] > words[1], hence the sequence is unsorted.

// Example 3:
// Input: words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
// Output: false
// Explanation: The first three characters "app" match, and the second string is shorter (in size.) According to lexicographical rules "apple" > "app", because 'l' > '∅', where '∅' is defined as the blank character which is less than any other character (More info).

// Constraints:
//     1 <= words.length <= 100
//     1 <= words[i].length <= 20
//     order.length == 26
//     All characters in words[i] and order are English lowercase letters.

import "fmt"

func isAlienSorted(words []string, order string) bool {
    alphabet := [26]int{}
    for i:=0; i < len(order); i++ {	
        alphabet[order[i] - 'a'] = i
    }
    prev := words[0]
    for i := 1; i < len(words); i++ {
        j := 0
        for j = 0; j < len(prev)-1 && j < len(words[i])-1 && prev[j] == words[i][j]; j++ {}
        if  alphabet[prev[j] - 'a'] >  alphabet[words[i][j] - 'a'] || 
            alphabet[prev[j] - 'a'] == alphabet[words[i][j] - 'a'] && 
            len(prev) > len(words[i]) {
            return false
        }
        prev = words[i]
    }
    return true
}

func main() {
    // Example 1:
    // Input: words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
    // Output: true
    // Explanation: As 'h' comes before 'l' in this language, then the sequence is sorted.
    fmt.Println(isAlienSorted([]string{"hello","leetcode"}, "hlabcdefgijkmnopqrstuvwxyz")) // true
    // Example 2:
    // Input: words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
    // Output: false
    // Explanation: As 'd' comes after 'l' in this language, then words[0] > words[1], hence the sequence is unsorted.
    fmt.Println(isAlienSorted([]string{"word","world","row"}, "worldabcefghijkmnpqstuvxyz")) // false
    // Example 3:
    // Input: words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
    // Output: false
    // Explanation: The first three characters "app" match, and the second string is shorter (in size.) According to lexicographical rules "apple" > "app", because 'l' > '∅', where '∅' is defined as the blank character which is less than any other character (More info).
    fmt.Println(isAlienSorted([]string{"apple","app"}, "abcdefghijklmnopqrstuvwxyz")) // false
}