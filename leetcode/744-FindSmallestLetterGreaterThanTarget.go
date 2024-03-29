package main

// 744. Find Smallest Letter Greater Than Target
// You are given an array of characters letters that is sorted in non-decreasing order, and a character target. There are at least two different characters in letters.
// Return the smallest character in letters that is lexicographically greater than target. 
// If such a character does not exist, return the first character in letters.

// Example 1:
// Input: letters = ["c","f","j"], target = "a"
// Output: "c"
// Explanation: The smallest character that is lexicographically greater than 'a' in letters is 'c'.

// Example 2:
// Input: letters = ["c","f","j"], target = "c"
// Output: "f"
// Explanation: The smallest character that is lexicographically greater than 'c' in letters is 'f'.

// Example 3:
// Input: letters = ["x","x","y","y"], target = "z"
// Output: "x"
// Explanation: There are no characters in letters that is lexicographically greater than 'z' so we return letters[0].
 
// Constraints:
//     2 <= letters.length <= 10^4
//     letters[i] is a lowercase English letter.
//     letters is sorted in non-decreasing order.
//     letters contains at least two different characters.
//     target is a lowercase English letter

import "fmt"

// 二分 O(logN)
func nextGreatestLetter(letters []byte, target byte) byte {
    left, right := 0, len(letters) - 1
    for left < right {
        mid := left + (right - left) / 2
        if letters[mid] > target {
            right = mid // 这里不 -1是关键
        } else {
            left = mid + 1
        }
    }
    if letters[left] <= target { // If such a character does not exist, return the first character in letters.
        return letters[0]
    }
    return letters[left]
}

// 遍历大法好 O(n)
func nextGreatestLetter1(letters []byte, target byte) byte {
    for _,v := range letters {
        // 找到第一个大于 target 的字符就返回
        if v > target {
            return v
        }
    }
    return letters[0]
}

func main() {
    // Explanation: The smallest character that is lexicographically greater than 'a' in letters is 'c'.
    fmt.Printf("%c\n",nextGreatestLetter([]byte{'c','f','j'},'a')) // c
    // Explanation: The smallest character that is lexicographically greater than 'c' in letters is 'f'.
    fmt.Printf("%c\n",nextGreatestLetter([]byte{'c','f','j'},'c')) // f
    // Explanation: There are no characters in letters that is lexicographically greater than 'z' so we return letters[0].
    fmt.Printf("%c\n",nextGreatestLetter([]byte{'x','x','y','y'},'z')) // x

    fmt.Printf("%c\n",nextGreatestLetter1([]byte{'c','f','j'},'a')) // c
    fmt.Printf("%c\n",nextGreatestLetter1([]byte{'c','f','j'},'c')) // f
    fmt.Printf("%c\n",nextGreatestLetter1([]byte{'x','x','y','y'},'z')) // x
}
    