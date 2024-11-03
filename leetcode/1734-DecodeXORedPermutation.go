package main

// 1734. Decode XORed Permutation
// There is an integer array perm that is a permutation of the first n positive integers, where n is always odd.

// It was encoded into another integer array encoded of length n - 1, such that encoded[i] = perm[i] XOR perm[i + 1]. 
// For example, if perm = [1,3,2], then encoded = [2,1].

// Given the encoded array, return the original array perm. It is guaranteed that the answer exists and is unique.

// Example 1:
// Input: encoded = [3,1]
// Output: [1,2,3]
// Explanation: If perm = [1,2,3], then encoded = [1 XOR 2,2 XOR 3] = [3,1]

// Example 2:
// Input: encoded = [6,5,4,6]
// Output: [2,4,1,5,3]

// Constraints:
//     3 <= n < 10^5
//     n is odd.
//     encoded.length == n - 1

import "fmt"

func decode(encoded []int) []int {
    sum, n := 0, len(encoded)
    for i := 1; i <= n + 1; i++ {
        sum ^= i
    }
    start := sum
    for i := 1; i < n; i+= 2 {
        start ^= encoded[i]
    }
    res := []int{ start }
    for _, v := range encoded {
        res = append(res, res[len(res) - 1] ^ v)
    }
    return res
}

func decode1(encoded []int) []int {
    n, sum, odd := len(encoded), 0, 0
    res := make([]int, n + 1)
    for i := 1; i <= n + 1; i++ {
        sum ^= i
    }
    for i := 1; i < n; i+=2 {
        odd ^= encoded[i]
    }
    res[0] = sum ^ odd
    for i := 1; i < n + 1; i++ {
        res[i] = res[i - 1] ^ encoded[i - 1]
    }
    return res
}

func main() {
    // Example 1:
    // Input: encoded = [3,1]
    // Output: [1,2,3]
    // Explanation: If perm = [1,2,3], then encoded = [1 XOR 2,2 XOR 3] = [3,1]
    fmt.Println(decode([]int{3,1})) // [1,2,3]
    // Example 2:
    // Input: encoded = [6,5,4,6]
    // Output: [2,4,1,5,3]
    fmt.Println(decode([]int{6,5,4,6})) // [2,4,1,5,3]

    fmt.Println(decode1([]int{3,1})) // [1,2,3]
    fmt.Println(decode1([]int{6,5,4,6})) // [2,4,1,5,3]
}