package main

// 3307. Find the K-th Character in String Game II
// Alice and Bob are playing a game. Initially, Alice has a string word = "a".

// You are given a positive integer k. 
// You are also given an integer array operations, where operations[i] represents the type of the ith operation.

// Now Bob will ask Alice to perform all operations in sequence:
//     1. If operations[i] == 0, append a copy of word to itself.
//     2. If operations[i] == 1, generate a new string by changing each character in word to its next character in the English alphabet, 
//        and append it to the original word. 
//        For example, performing the operation on "c" generates "cd" and performing the operation on "zb" generates "zbac".

// Return the value of the kth character in word after performing all the operations.

// Note that the character 'z' can be changed to 'a' in the second type of operation.

// Example 1:
// Input: k = 5, operations = [0,0,0]
// Output: "a"
// Explanation:
// Initially, word == "a". Alice performs the three operations as follows:
// Appends "a" to "a", word becomes "aa".
// Appends "aa" to "aa", word becomes "aaaa".
// Appends "aaaa" to "aaaa", word becomes "aaaaaaaa".

// Example 2:
// Input: k = 10, operations = [0,1,0,1]
// Output: "b"
// Explanation:
// Initially, word == "a". Alice performs the four operations as follows:
// Appends "a" to "a", word becomes "aa".
// Appends "bb" to "aa", word becomes "aabb".
// Appends "aabb" to "aabb", word becomes "aabbaabb".
// Appends "bbccbbcc" to "aabbaabb", word becomes "aabbaabbbbccbbcc".

// Constraints:
//     1 <= k <= 10^14
//     1 <= operations.length <= 100
//     operations[i] is either 0 or 1.
//     The input is generated such that word has at least k characters after all operations.

import "fmt"

func kthCharacter(k int64, operations []int) byte {
    k--
    c := 0
    for i := len(operations) - 1; i >= 0; i-- {
        if k >> i & 1 == 1 {
            c += operations[i]
        }
    }
    return 'a' + byte(c % 26)
}

func kthCharacter1(k int64, operations []int) byte {
    // 0：复制一份
    // 1：+1复制一份
    incr, e, length := 0, 0, int64(1)
    for length < k {
        e++
        length *= 2
    }
    // length = 2^e
    index := e - 1
    for index >= 0 {
        if k > length / 2 { // 可能发生变化
            if operations[index] == 1 {
                incr++
            }
            k = k - length / 2
        }
        length /= 2
        index--
    }
    return 'a' + byte(incr % 26)
}

func main() {
    // Example 1:
    // Input: k = 5, operations = [0,0,0]
    // Output: "a"
    // Explanation:
    // Initially, word == "a". Alice performs the three operations as follows:
    // Appends "a" to "a", word becomes "aa".
    // Appends "aa" to "aa", word becomes "aaaa".
    // Appends "aaaa" to "aaaa", word becomes "aaaaaaaa".
    fmt.Printf("%c\n",kthCharacter(5,[]int{0,0,0})) // "a"
    // Example 2:
    // Input: k = 10, operations = [0,1,0,1]
    // Output: "b"
    // Explanation:
    // Initially, word == "a". Alice performs the four operations as follows:
    // Appends "a" to "a", word becomes "aa".
    // Appends "bb" to "aa", word becomes "aabb".
    // Appends "aabb" to "aabb", word becomes "aabbaabb".
    // Appends "bbccbbcc" to "aabbaabb", word becomes "aabbaabbbbccbbcc".
    fmt.Printf("%c\n",kthCharacter(10, []int{0,1,0,1})) // "b"

    fmt.Printf("%c\n",kthCharacter(10, []int{1,0,1,0,1,0,1,0,1,0})) // "b"
    fmt.Printf("%c\n",kthCharacter(10, []int{0,0,0,0,0,0,0,0,0,0})) // "a"
    fmt.Printf("%c\n",kthCharacter(10, []int{1,1,1,1,1,1,1,1,1,1})) // "c"
    fmt.Printf("%c\n",kthCharacter(10, []int{1,1,1,1,1,0,0,0,0,0})) // "c"
    fmt.Printf("%c\n",kthCharacter(10, []int{0,0,0,0,0,1,1,1,1,1})) // "a"

    fmt.Printf("%c\n",kthCharacter1(5,[]int{0,0,0})) // "a"
    fmt.Printf("%c\n",kthCharacter1(10, []int{0,1,0,1})) // "b"
    fmt.Printf("%c\n",kthCharacter1(10, []int{1,0,1,0,1,0,1,0,1,0})) // "b"
    fmt.Printf("%c\n",kthCharacter1(10, []int{0,0,0,0,0,0,0,0,0,0})) // "a"
    fmt.Printf("%c\n",kthCharacter1(10, []int{1,1,1,1,1,1,1,1,1,1})) // "c"
    fmt.Printf("%c\n",kthCharacter1(10, []int{1,1,1,1,1,0,0,0,0,0})) // "c"
    fmt.Printf("%c\n",kthCharacter1(10, []int{0,0,0,0,0,1,1,1,1,1})) // "a"
}