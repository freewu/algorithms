package main

// 3304. Find the K-th Character in String Game I
// Alice and Bob are playing a game. Initially, Alice has a string word = "a".

// You are given a positive integer k.

// Now Bob will ask Alice to perform the following operation forever:
//     Generate a new string by changing each character in word to its next character in the English alphabet, 
//     and append it to the original word.

// For example, performing the operation on "c" generates "cd" and performing the operation on "zb" generates "zbac".

// Return the value of the kth character in word, 
// after enough operations have been done for word to have at least k characters.

// Note that the character 'z' can be changed to 'a' in the operation.

// Example 1:
// Input: k = 5
// Output: "b"
// Explanation:
// Initially, word = "a". We need to do the operation three times:
// Generated string is "b", word becomes "ab".
// Generated string is "bc", word becomes "abbc".
// Generated string is "bccd", word becomes "abbcbccd".

// Example 2:
// Input: k = 10
// Output: "c"

// Constraints:
//     1 <= k <= 500

import "fmt"

func kthCharacter(k int) byte {
    arr := []byte{'a'}
    for len(arr) < k {
        next := []byte{}
        for _, ch := range arr {
            next = append(next, ch)
            next = append(next, ch+1)
        }
        arr = next
    }
    return arr[k-1]
}

func kthCharacter1(k int) byte {
    n, s := 1, "a"
    for n < k {
        t := []byte{}
        for _, c := range s {
            if c == 'z' {
                t = append(t, 'a')
            } else {
                t = append(t, byte(c+1))
            }
        }
        s += string(t)
        n = len(s)
    }
    return s[k-1]
}

func main() {
    // Example 1:
    // Input: k = 5
    // Output: "b"
    // Explanation:
    // Initially, word = "a". We need to do the operation three times:
    // Generated string is "b", word becomes "ab".
    // Generated string is "bc", word becomes "abbc".
    // Generated string is "bccd", word becomes "abbcbccd".
    fmt.Printf("%c\n",kthCharacter(5)) // "b"
    // Example 2: 
    // Input: k = 10
    // Output: "c"
    fmt.Printf("%c\n",kthCharacter(10)) // "c"

    fmt.Printf("%c\n",kthCharacter(1)) // "a"
    fmt.Printf("%c\n",kthCharacter(2)) // "b"
    fmt.Printf("%c\n",kthCharacter(4)) // "c"
    fmt.Printf("%c\n",kthCharacter(64)) // "g"
    fmt.Printf("%c\n",kthCharacter(499)) // "g"
    fmt.Printf("%c\n",kthCharacter(500)) // "h"

    fmt.Printf("%c\n",kthCharacter1(5)) // "b"
    fmt.Printf("%c\n",kthCharacter1(10)) // "c"
    fmt.Printf("%c\n",kthCharacter1(1)) // "a"
    fmt.Printf("%c\n",kthCharacter1(2)) // "b"
    fmt.Printf("%c\n",kthCharacter1(4)) // "c"
    fmt.Printf("%c\n",kthCharacter1(64)) // "g"
    fmt.Printf("%c\n",kthCharacter1(499)) // "g"
    fmt.Printf("%c\n",kthCharacter1(500)) // "h"
}