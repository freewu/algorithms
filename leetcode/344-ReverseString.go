package main

// 344. Reverse String
// Write a function that reverses a string. 
// The input string is given as an array of characters s.

// You must do this by modifying the input array in-place with O(1) extra memory.

// Example 1:
// Input: s = ["h","e","l","l","o"]
// Output: ["o","l","l","e","h"]

// Example 2:
// Input: s = ["H","a","n","n","a","h"]
// Output: ["h","a","n","n","a","H"]
 
// Constraints:
//     1 <= s.length <= 10^5
//     s[i] is a printable ascii character.

import "fmt"

// func reverseString(s string) string {
//     var l = len(s)
//     if l <= 1 {
//         return s
//     }
//     var i = 0
//     var t = make([]byte, l)
//     l = l - 1
//     for {
//         if i > l {
//             break
//         }
//         // syntax sugar
//         t[l], t[i] = s[i], s[l]
//         i++
//         l--
//     }
//     return string(t)
// }

// func reverseString1(s string) string {
//     l, r := 0, len(s) - 1
//     for l < r {
//         t := s[l]
//         s[l] = s[r]
//         s[r] = t
//         // s[l],s[r] = s[r],s[l]
//         l++
//         r--
//     }
//     return s 
// }

func reverseString(s []byte)  {
    l, r := 0, len(s) - 1
    for l < r {
        s[l],s[r] = s[r],s[l]
        l++
        r--
    }
}

func main() {
    s1 := []byte{'a','b'}
    reverseString(s1)
    fmt.Println(string(s1)) // ba
    s2 := []byte{'1','2','3'}
    reverseString(s2)
    fmt.Println(string(s2))  // 321
    s3 := []byte{'H','e','l','l','o'}
    reverseString(s3)
    fmt.Println(string(s3)) // olleh
}
