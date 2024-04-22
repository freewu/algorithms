package main

// 68. Text Justification
// Given an array of strings words and a width maxWidth, 
// format the text such that each line has exactly maxWidth characters and is fully (left and right) justified.

// You should pack your words in a greedy approach; that is, pack as many words as you can in each line. 
// Pad extra spaces ' ' when necessary so that each line has exactly maxWidth characters.

// Extra spaces between words should be distributed as evenly as possible. 
// If the number of spaces on a line does not divide evenly between words, 
// the empty slots on the left will be assigned more spaces than the slots on the right.

// For the last line of text, it should be left-justified, and no extra space is inserted between words.

// Note:
//     A word is defined as a character sequence consisting of non-space characters only.
//     Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
//     The input array words contains at least one word.

// Example 1:
// Input: words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16
// Output:
// [
//    "This    is    an",
//    "example  of text",
//    "justification.  "
// ]

// Example 2:
// Input: words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
// Output:
// [
//   "What   must   be",
//   "acknowledgment  ",
//   "shall be        "
// ]
// Explanation: Note that the last line is "shall be    " instead of "shall     be", because the last line must be left-justified instead of fully-justified.
// Note that the second line is also left-justified because it contains only one word.

// Example 3:
// Input: words = ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"], maxWidth = 20
// Output:
// [
//   "Science  is  what we",
//   "understand      well",
//   "enough to explain to",
//   "a  computer.  Art is",
//   "everything  else  we",
//   "do                  "
// ]

// Constraints:
//     1 <= words.length <= 300
//     1 <= words[i].length <= 20
//     words[i] consists of only English letters and symbols.
//     1 <= maxWidth <= 100
//     words[i].length <= maxWidth

import "fmt"
import "strings"

func fullJustify(words []string, maxWidth int) []string {
    res, cur, num_of_letters := []string{}, []string{}, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, word := range words {
        if len(word) + len(cur) + num_of_letters > maxWidth {
            for i := 0; i < maxWidth - num_of_letters; i++ {
                cur[i % max(1, len(cur) - 1)] += " "
            }
            res = append(res, strings.Join(cur, ""))
            cur = cur[:0]
            num_of_letters = 0
        }
        cur = append(cur, word)
        num_of_letters += len(word)
    }
    lastLine := strings.Join(cur, " ")
    for len(lastLine) < maxWidth {
        lastLine += " "
    }
    res = append(res, lastLine)
    return res
}

func main() {
    // Input: words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16
    // Output:
    // [
    //    "This    is    an",
    //    "example  of text",
    //    "justification.  "
    // ]
    fmt.Println(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
    // Input: words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
    // Output:
    // [
    //   "What   must   be",
    //   "acknowledgment  ",
    //   "shall be        "
    // ]
    // Explanation: Note that the last line is "shall be    " instead of "shall     be", because the last line must be left-justified instead of fully-justified.
    // Note that the second line is also left-justified because it contains only one word.
    fmt.Println(fullJustify([]string{"What","must","be","acknowledgment","shall","be"}, 16))
    // Example 3:
    // Input: words = ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"], maxWidth = 20
    // Output:
    // [
    //   "Science  is  what we",
    //   "understand      well",
    //   "enough to explain to",
    //   "a  computer.  Art is",
    //   "everything  else  we",
    //   "do                  "
    // ]
    fmt.Println(fullJustify([]string{"Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"}, 20))
}