package main

// 804. Unique Morse Code Words
// International Morse Code defines a standard encoding where each letter is mapped to a series of dots and dashes, 
// as follows:
//     'a' maps to ".-",
//     'b' maps to "-...",
//     'c' maps to "-.-.", and so on.

// For convenience, the full table for the 26 letters of the English alphabet is given below:
//     [".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]

// Given an array of strings words where each word can be written as a concatenation of the Morse code of each letter.
//     For example, "cab" can be written as "-.-..--...", which is the concatenation of "-.-.", ".-", and "-...". 
//     We will call such a concatenation the transformation of a word.

// Return the number of different transformations among all words we have.

// Example 1:
// Input: words = ["gin","zen","gig","msg"]
// Output: 2
// Explanation: The transformation of each word is:
// "gin" -> "--...-."
// "zen" -> "--...-."
// "gig" -> "--...--."
// "msg" -> "--...--."
// There are 2 different transformations: "--...-." and "--...--.".

// Example 2:
// Input: words = ["a"]
// Output: 1

// Constraints:
//     1 <= words.length <= 100
//     1 <= words[i].length <= 12
//     words[i] consists of lowercase English letters.

import "fmt"
import "strings"

func uniqueMorseRepresentations(words []string) int {
    morses := [] string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
    transformations := make(map[string]bool)
    for _, s := range words {
        transformation := ""
        runes := [] rune(s)
        for _, r := range runes {
            index := r - 'a'
            transformation += morses[index]
        }
        transformations[transformation] = true
    }
    return len(transformations)
}

func uniqueMorseRepresentations1(words []string) int {
    morses := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.","---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
    s := make(map[string]bool)
    for _, word := range words {
        t := &strings.Builder{}
        for _, c := range word {
            t.WriteString(morses[c-'a'])
        }
        s[t.String()] = true
    }
    return len(s)
}

func main() {
    // Example 1:
    // Input: words = ["gin","zen","gig","msg"]
    // Output: 2
    // Explanation: The transformation of each word is:
    // "gin" -> "--...-."
    // "zen" -> "--...-."
    // "gig" -> "--...--."
    // "msg" -> "--...--."
    // There are 2 different transformations: "--...-." and "--...--.".
    fmt.Println(uniqueMorseRepresentations([]string{"gin","zen","gig","msg"})) // 2
    // Example 2:
    // Input: words = ["a"]
    // Output: 1
    fmt.Println(uniqueMorseRepresentations([]string{"a"})) // 1

    fmt.Println(uniqueMorseRepresentations1([]string{"gin","zen","gig","msg"})) // 2
    fmt.Println(uniqueMorseRepresentations1([]string{"a"})) // 1
}