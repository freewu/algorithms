package main

// 1592. Rearrange Spaces Between Words
// You are given a string text of words that are placed among some number of spaces. 
// Each word consists of one or more lowercase English letters and are separated by at least one space. 
// It's guaranteed that text contains at least one word.

// Rearrange the spaces so that there is an equal number of spaces between every pair of adjacent words and that number is maximized. 
// If you cannot redistribute all the spaces equally, place the extra spaces at the end, meaning the returned string should be the same length as text.

// Return the string after rearranging the spaces.

// Example 1:
// Input: text = "  this   is  a sentence "
// Output: "this   is   a   sentence"
// Explanation: There are a total of 9 spaces and 4 words. We can evenly divide the 9 spaces between the words: 9 / (4-1) = 3 spaces.

// Example 2:
// Input: text = " practice   makes   perfect"
// Output: "practice   makes   perfect "
// Explanation: There are a total of 7 spaces and 3 words. 7 / (3-1) = 3 spaces plus 1 extra space. We place this extra space at the end of the string.

// Constraints:
//     1 <= text.length <= 100
//     text consists of lowercase English letters and ' '.
//     text contains at least one word.

import "fmt"
import "strings"

func reorderSpaces(text string) string {
    words, spaces := strings.Fields(text), strings.Count(text, " ")
    if len(words) == 1 {
        return words[0] + strings.Repeat(" ", spaces)
    }
    spacesBetween, rem := spaces / (len(words) - 1), spaces % (len(words) - 1)
    sep := strings.Repeat(" ", spacesBetween)
    return strings.Join(words, sep) + strings.Repeat(" ", rem)
}

func reorderSpaces1(text string) string {
    n, nums, blanks := len(text), 0, 0
    arr := []byte(text)
    for i := 0; i < n; {
        for i < n && arr[i] == ' ' {
            blanks++
            i++
        }
        if i == n { break }
        nums++
        for i < n && arr[i] != ' ' {
            i++
        }
    }
    avg := blanks
    if nums > 1 {
        avg = blanks / (nums - 1)
    }
    left, right, count := 0, 0,  0
    for left < n && right < n {
        for right < n && text[right] == ' ' {
            right++
        }
        if right == n { break }
        for right < n && text[right] != ' ' {
            arr[left] = byte(text[right])
            left++
            right++
        }
        count++
        if left >= n { break }
        if count == nums {
            for left < n {
                arr[left] = ' '
                left++
            }
            break
        } else {
            for i := 0; i < avg; i++ {
                arr[left] = ' '
                left++
            }
        }
    }
    return string(arr)
}

func main() {
    // Example 1:
    // Input: text = "  this   is  a sentence "
    // Output: "this   is   a   sentence"
    // Explanation: There are a total of 9 spaces and 4 words. We can evenly divide the 9 spaces between the words: 9 / (4-1) = 3 spaces.
    fmt.Println(reorderSpaces("  this   is  a sentence ")) // "this   is   a   sentence"
    // Example 2:
    // Input: text = " practice   makes   perfect"
    // Output: "practice   makes   perfect "
    // Explanation: There are a total of 7 spaces and 3 words. 7 / (3-1) = 3 spaces plus 1 extra space. We place this extra space at the end of the string.
    fmt.Println(reorderSpaces(" practice   makes   perfect")) // "practice   makes   perfect "

    fmt.Println(reorderSpaces1("  this   is  a sentence ")) // "this   is   a   sentence"
    fmt.Println(reorderSpaces1(" practice   makes   perfect")) // "practice   makes   perfect "
}