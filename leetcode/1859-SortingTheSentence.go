package main

// 1859. Sorting the Sentence
// A sentence is a list of words that are separated by a single space with no leading or trailing spaces. 
// Each word consists of lowercase and uppercase English letters.

// A sentence can be shuffled by appending the 1-indexed word position to each word then rearranging the words in the sentence.
//     For example, the sentence "This is a sentence" can be shuffled as "sentence4 a3 is2 This1" or "is2 sentence4 This1 a3".

// Given a shuffled sentence s containing no more than 9 words, reconstruct and return the original sentence.

// Example 1:
// Input: s = "is2 sentence4 This1 a3"
// Output: "This is a sentence"
// Explanation: Sort the words in s to their original positions "This1 is2 a3 sentence4", then remove the numbers.

// Example 2:
// Input: s = "Myself2 Me1 I4 and3"
// Output: "Me Myself and I"
// Explanation: Sort the words in s to their original positions "Me1 Myself2 and3 I4", then remove the numbers.

// Constraints:
//     2 <= s.length <= 200
//     s consists of lowercase and uppercase English letters, spaces, and digits from 1 to 9.
//     The number of words in s is between 1 and 9.
//     The words in s are separated by a single space.
//     s contains no leading or trailing spaces.

import "fmt"
import "strings"

func sortSentence(s string) string {
    indexes, arr := []int{}, strings.Split(s, " ")
    for _, word := range arr {
        indexes = append(indexes, int(word[len(word) - 1] - '0'))
    }
    res := make([]string, len(indexes))
    for i, index := range indexes {
        res[index - 1] = arr[i][:len(arr[i]) - 1]
    }
    return strings.Join(res, " ")
}

func sortSentence1(s string) string {
    mp, arr := make(map[rune]string), strings.Split(s, " ")
    for _, v := range arr {
        mp[rune(v[len(v) - 1])] = v[:len(v) - 1]
    }
    res := []string{}
    for i := '1' ; i < rune(len(arr) + '1') ; i++ {
        res = append(res, mp[i])
    }
    return  strings.Join(res, " ")
}

func main() {
    // Example 1:
    // Input: s = "is2 sentence4 This1 a3"
    // Output: "This is a sentence"
    // Explanation: Sort the words in s to their original positions "This1 is2 a3 sentence4", then remove the numbers.
    fmt.Println(sortSentence("is2 sentence4 This1 a3")) // "This is a sentence"
    // Example 2:
    // Input: s = "Myself2 Me1 I4 and3"
    // Output: "Me Myself and I"
    // Explanation: Sort the words in s to their original positions "Me1 Myself2 and3 I4", then remove the numbers.
    fmt.Println(sortSentence("Myself2 Me1 I4 and3")) // "Me Myself and I"

    fmt.Println(sortSentence1("is2 sentence4 This1 a3")) // "This is a sentence"
    fmt.Println(sortSentence1("Myself2 Me1 I4 and3")) // "Me Myself and I"
}