package main

// 1652. Defuse the Bomb
// You have a bomb to defuse, and your time is running out! 
// Your informer will provide you with a circular array code of length of n and a key k.

// To decrypt the code, you must replace every number. 
// All the numbers are replaced simultaneously.
//     If k > 0, replace the ith number with the sum of the next k numbers.
//     If k < 0, replace the ith number with the sum of the previous k numbers.
//     If k == 0, replace the ith number with 0.

// As code is circular, the next element of code[n-1] is code[0], and the previous element of code[0] is code[n-1].
// Given the circular array code and an integer key k, return the decrypted code to defuse the bomb!

// Example 1:
// Input: code = [5,7,1,4], k = 3
// Output: [12,10,16,13]
// Explanation: Each number is replaced by the sum of the next 3 numbers. The decrypted code is [7+1+4, 1+4+5, 4+5+7, 5+7+1]. Notice that the numbers wrap around.

// Example 2:
// Input: code = [1,2,3,4], k = 0
// Output: [0,0,0,0]
// Explanation: When k is zero, the numbers are replaced by 0. 

// Example 3:
// Input: code = [2,4,9,3], k = -2
// Output: [12,5,6,13]
// Explanation: The decrypted code is [3+9, 2+3, 4+2, 9+4]. Notice that the numbers wrap around again. If k is negative, the sum is of the previous numbers.

// Constraints:
//     n == code.length
//     1 <= n <= 100
//     1 <= code[i] <= 100
//     -(n - 1) <= k <= n - 1

import "fmt"

func decrypt(code []int, k int) []int {
    res := make([]int,len(code))
    if k == 0 {
        return res
    }
    start, end, sum := 1, k, 0
    if k < 0 {
        k = -k
        start, end = len(code) - k, len(code) - 1
    }
    for i := start; i <= end; i++ {
        sum += code[i]
    }
    for i := 0; i < len(code); i++ {
        end++
        res[i] = sum
        sum -= code[start % len(code)]
        sum += code[end % len(code)]
        start++
    }
    return res
}

func decrypt1(code []int, k int) []int {
    n := len(code)
    res := make([]int, n)
    if k == 0 {
        return res
    }
    l, r := 1, k
    if k < 0 {
        l, r = n+k, n-1
    }
    sum := 0
    for _, v := range code[l : r+1] {
        sum += v
    }
    for i := range res {
        res[i] = sum
        sum = sum - code[l] + code[(r+1)%n]
        l, r = (l+1)%n, (r+1)%n
    }
    return res
}

func main() {
    // Example 1:
    // Input: code = [5,7,1,4], k = 3
    // Output: [12,10,16,13]
    // Explanation: Each number is replaced by the sum of the next 3 numbers. The decrypted code is [7+1+4, 1+4+5, 4+5+7, 5+7+1]. Notice that the numbers wrap around.
    fmt.Println(decrypt([]int{5,7,1,4}, 3)) // [12,10,16,13]
    // Example 2:
    // Input: code = [1,2,3,4], k = 0
    // Output: [0,0,0,0]
    // Explanation: When k is zero, the numbers are replaced by 0. 
    fmt.Println(decrypt([]int{1,2,3,4}, 0)) // [0,0,0,0]
    // Example 3:
    // Input: code = [2,4,9,3], k = -2
    // Output: [12,5,6,13]
    // Explanation: The decrypted code is [3+9, 2+3, 4+2, 9+4]. Notice that the numbers wrap around again. If k is negative, the sum is of the previous numbers.
    fmt.Println(decrypt([]int{2,4,9,3}, -2)) // [12,5,6,13]

    fmt.Println(decrypt1([]int{5,7,1,4}, 3)) // [12,10,16,13]
    fmt.Println(decrypt1([]int{1,2,3,4}, 0)) // [0,0,0,0]
    fmt.Println(decrypt1([]int{2,4,9,3}, -2)) // [12,5,6,13]
}