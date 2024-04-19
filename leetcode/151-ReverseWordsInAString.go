package main

// 151. Reverse Words in a String
// Given an input string s, reverse the order of the words.
// A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.
// Return a string of the words in reverse order concatenated by a single space.
// Note that s may contain leading or trailing spaces or multiple spaces between two words. 
// The returned string should only have a single space separating the words. Do not include any extra spaces.

// Example 1:
// Input: s = "the sky is blue"
// Output: "blue is sky the"

// Example 2:
// Input: s = "  hello world  "
// Output: "world hello"
// Explanation: Your reversed string should not contain leading or trailing spaces.

// Example 3:
// Input: s = "a good   example"
// Output: "example good a"
// Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.
 
// Constraints:
//     1 <= s.length <= 10^4
//     s contains English letters (upper-case and lower-case), digits, and spaces ' '.
//     There is at least one word in s.

// Follow-up: If the string data type is mutable in your language, can you solve it in-place with O(1) extra space?

import "fmt"
import "strings"

func reverseWords(s string) string {
    // 先把字符串按照空格分隔成每个小单词 返回数组
    words := strings.Fields(s)
    reverse := func (m *[]string, i int, j int) {
        for i <= j {
            // 单词前后翻转
            (*m)[i], (*m)[j] = (*m)[j], (*m)[i]
            i++
            j--
        }
    }
    reverse(&words, 0, len(words) - 1)
    return strings.Join(words, " ") // 重新组合成字符串
}


// 双指针
func reverseWords1(s string) string {
    words := strings.Fields(s)
    first, last := 0, len(words)-1
    for first < last {
        words[first], words[last] =  words[last], words[first]
        first++
        last--
    }
    return strings.Join(words, " ")
}

func main() {
    fmt.Println(reverseWords("the sky is blue")) // blue is sky the
    fmt.Println(reverseWords("  hello world  ")) // world hello
    fmt.Println(reverseWords("a good   example")) // example good a

    fmt.Println(reverseWords1("the sky is blue")) // blue is sky the
    fmt.Println(reverseWords1("  hello world  ")) // world hello
    fmt.Println(reverseWords1("a good   example")) // example good a
}