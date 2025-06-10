package main

// 3577. Count the Number of Computer Unlocking Permutations
// You are given an array complexity of length n.

// There are n locked computers in a room with labels from 0 to n - 1, each with its own unique password. 
// The password of the computer i has a complexity complexity[i].

// The password for the computer labeled 0 is already decrypted and serves as the root. 
// All other computers must be unlocked using it or another previously unlocked computer, following this information:
//     1. You can decrypt the password for the computer i using the password for computer j, where j is any integer less than i with a lower complexity. (i.e. j < i and complexity[j] < complexity[i])
//     2. To decrypt the password for computer i, you must have already unlocked a computer j such that j < i and complexity[j] < complexity[i].

// Find the number of permutations of [0, 1, 2, ..., (n - 1)] that represent a valid order in which the computers can be unlocked, starting from computer 0 as the only initially unlocked one.

// Since the answer may be large, return it modulo 10^9 + 7.

// Note that the password for the computer with label 0 is decrypted, and not the computer with the first position in the permutation.

// Example 1:
// Input: complexity = [1,2,3]
// Output: 2
// Explanation:
// The valid permutations are:
// [0, 1, 2]
// Unlock computer 0 first with root password.
// Unlock computer 1 with password of computer 0 since complexity[0] < complexity[1].
// Unlock computer 2 with password of computer 1 since complexity[1] < complexity[2].
// [0, 2, 1]
// Unlock computer 0 first with root password.
// Unlock computer 2 with password of computer 0 since complexity[0] < complexity[2].
// Unlock computer 1 with password of computer 0 since complexity[0] < complexity[1].

// Example 2:
// Input: complexity = [3,3,3,4,4,4]
// Output: 0
// Explanation:
// There are no possible permutations which can unlock all computers.

// Constraints:
//     2 <= complexity.length <= 10^5
//     1 <= complexity[i] <= 10^9

import "fmt"

func countPermutations(complexity []int) int {
    res, n, mod := 1, len(complexity), 1_000_000_007
    for i := 1; i < n; i++ {
        if complexity[i] <= complexity[0] {
            return 0
        }
    }
    for i := 1; i < n; i++ {
        res *= i
        res %= mod
    }
    return res
}

func countPermutations1(complexity []int) int {
    res, n, mod, mn := 1, len(complexity), 1_000_000_007, complexity[0]
    for i := 1; i < n; i++ {
        if mn >= complexity[i] { // 如果不是从小到大的序列
            return 0
        }
        if complexity[i] < mn {
            mn = complexity[i]
        }
    }
    for i := 1; i < n; i++ {
        res = (res * i) % mod
    }
    return res
}

func main() {
    // Example 1:
    // Input: complexity = [1,2,3]
    // Output: 2
    // Explanation:
    // The valid permutations are:
    // [0, 1, 2]
    // Unlock computer 0 first with root password.
    // Unlock computer 1 with password of computer 0 since complexity[0] < complexity[1].
    // Unlock computer 2 with password of computer 1 since complexity[1] < complexity[2].
    // [0, 2, 1]
    // Unlock computer 0 first with root password.
    // Unlock computer 2 with password of computer 0 since complexity[0] < complexity[2].
    // Unlock computer 1 with password of computer 0 since complexity[0] < complexity[1].
    fmt.Println(countPermutations([]int{1,2,3})) // 2
    // Example 2:
    // Input: complexity = [3,3,3,4,4,4]
    // Output: 0
    // Explanation:
    // There are no possible permutations which can unlock all computers.
    fmt.Println(countPermutations([]int{3,3,3,4,4,4})) // 0

    fmt.Println(countPermutations([]int{1,2,3,4,5,6,7,8,9})) // 40320
    fmt.Println(countPermutations([]int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(countPermutations1([]int{1,2,3})) // 2
    fmt.Println(countPermutations1([]int{3,3,3,4,4,4})) // 0
    fmt.Println(countPermutations1([]int{1,2,3,4,5,6,7,8,9})) // 40320
    fmt.Println(countPermutations1([]int{9,8,7,6,5,4,3,2,1})) // 0
}