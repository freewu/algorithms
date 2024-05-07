package main

// 443. String Compression
// Given an array of characters chars, compress it using the following algorithm:
// Begin with an empty string s. For each group of consecutive repeating characters in chars:
//     If the group's length is 1, append the character to s.
//     Otherwise, append the character followed by the group's length.

// The compressed string s should not be returned separately, but instead, be stored in the input character array chars. 
// Note that group lengths that are 10 or longer will be split into multiple characters in chars.
// After you are done modifying the input array, return the new length of the array.
// You must write an algorithm that uses only constant extra space.

// Example 1:
// Input: chars = ["a","a","b","b","c","c","c"]
// Output: Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]
// Explanation: The groups are "aa", "bb", and "ccc". This compresses to "a2b2c3".

// Example 2:
// Input: chars = ["a"]
// Output: Return 1, and the first character of the input array should be: ["a"]
// Explanation: The only group is "a", which remains uncompressed since it's a single character.

// Example 3:
// Input: chars = ["a","b","b","b","b","b","b","b","b","b","b","b","b"]
// Output: Return 4, and the first 4 characters of the input array should be: ["a","b","1","2"].
// Explanation: The groups are "a" and "bbbbbbbbbbbb". This compresses to "ab12".

// Constraints:
//     1 <= chars.length <= 2000
//     chars[i] is a lowercase English letter, uppercase English letter, digit, or symbol.

import "fmt"
import "strconv"

func compress(chars []byte) int {
    anchor, res := 0, 0
    for read, ch := range chars {
        if read + 1 == len(chars) || chars[read + 1] != ch { // 如果字符相同
            chars[res] = chars[anchor]
            res++
            if read > anchor {
                for _, digit := range strconv.Itoa(read - anchor + 1) {
                    chars[res] = byte(digit)
                    res++
                }
            }
            anchor = read + 1
        }
    }
    return res
}

func compress1(chars []byte) int {
    res, left := 0, 0
    for read, ch := range chars {
        if read == len(chars)-1 || ch != chars[read+1] {
            chars[res] = ch
            res++
            num := read - left + 1
            if num > 1 {
                anchor := res
                for ; num > 0; num /= 10 {
                    chars[res] = '0' + byte(num%10)
                    res++
                }
                s := chars[anchor:res]
                for i, n := 0, len(s); i < n/2; i++ {
                    s[i], s[n-1-i] = s[n-1-i], s[i]
                }
            }
            left = read + 1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: chars = ["a","a","b","b","c","c","c"]
    // Output: Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]
    // Explanation: The groups are "aa", "bb", and "ccc". This compresses to "a2b2c3".
    fmt.Println(compress([]byte("aabbccc"))) // 6
    // Example 2:
    // Input: chars = ["a"]
    // Output: Return 1, and the first character of the input array should be: ["a"]
    // Explanation: The only group is "a", which remains uncompressed since it's a single character.
    fmt.Println(compress([]byte("a"))) // 1
    // Example 3:
    // Input: chars = ["a","b","b","b","b","b","b","b","b","b","b","b","b"]
    // Output: Return 4, and the first 4 characters of the input array should be: ["a","b","1","2"].
    // Explanation: The groups are "a" and "bbbbbbbbbbbb". This compresses to "ab12".
    fmt.Println(compress([]byte("abbbbbbbbbbbbbbbbbbbbbbbb"))) // 4

    fmt.Println(compress1([]byte("aabbccc"))) // 6
    fmt.Println(compress1([]byte("a"))) // 1
    fmt.Println(compress1([]byte("abbbbbbbbbbbbbbbbbbbbbbbb"))) // 4
}