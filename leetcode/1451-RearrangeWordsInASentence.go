package main

// 1451. Rearrange Words in a Sentence
// Given a sentence text (A sentence is a string of space-separated words) in the following format:
//     First letter is in upper case.
//     Each word in text are separated by a single space.

// Your task is to rearrange the words in text such that all words are rearranged in an increasing order of their lengths. 
// If two words have the same length, arrange them in their original order.

// Return the new text following the format shown above.

// Example 1:
// Input: text = "Leetcode is cool"
// Output: "Is cool leetcode"
// Explanation: There are 3 words, "Leetcode" of length 8, "is" of length 2 and "cool" of length 4.
// Output is ordered by length and the new first word starts with capital letter.

// Example 2:
// Input: text = "Keep calm and code on"
// Output: "On and keep calm code"
// Explanation: Output is ordered as follows:
// "On" 2 letters.
// "and" 3 letters.
// "keep" 4 letters in case of tie order by position in original text.
// "calm" 4 letters.
// "code" 4 letters.

// Example 3:
// Input: text = "To be or not to be"
// Output: "To be or to be not"

// Constraints:
//     text begins with a capital letter and then contains lowercase letters and single space between words.
//     1 <= text.length <= 10^5

import "fmt"
import "sort"
import "strings"

func arrangeWords(text string) string {
    arr := strings.Split(strings.ToLower(text), " ")
    sort.SliceStable(arr, func(a, b int) bool { // 长度从短到长
        return len(arr[a]) < len(arr[b]) 
    })
    res := strings.Join(arr, " ")
    return string(res[0] - 32) + res[1:] // 开头大写
}

func main() {
    // Example 1:
    // Input: text = "Leetcode is cool"
    // Output: "Is cool leetcode"
    // Explanation: There are 3 words, "Leetcode" of length 8, "is" of length 2 and "cool" of length 4.
    // Output is ordered by length and the new first word starts with capital letter.
    fmt.Println(arrangeWords("Leetcode is cool")) // "Is cool leetcode"
    // Example 2:
    // Input: text = "Keep calm and code on"
    // Output: "On and keep calm code"
    // Explanation: Output is ordered as follows:
    // "On" 2 letters.
    // "and" 3 letters.
    // "keep" 4 letters in case of tie order by position in original text.
    // "calm" 4 letters.
    // "code" 4 letters.
    fmt.Println(arrangeWords("On and keep calm code")) // "Is cool leetcode"

    // Example 3:
    // Input: text = "To be or not to be"
    // Output: "To be or to be not"
    fmt.Println(arrangeWords("To be or not to be")) // "To be or to be not"
}