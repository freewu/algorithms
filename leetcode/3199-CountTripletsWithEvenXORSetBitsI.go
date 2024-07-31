package main

// 3199. Count Triplets with Even XOR Set Bits I
// Given three integer arrays a, b, and c, return the number of triplets (a[i], b[j], c[k]), 
// such that the bitwise XOR of the elements of each triplet has an even number of set bits.

// Example 1:
// Input: a = [1], b = [2], c = [3]
// Output: 1
// Explanation:
// The only triplet is (a[0], b[0], c[0]) and their XOR is: 1 XOR 2 XOR 3 = 002.

// Example 2:
// Input: a = [1,1], b = [2,3], c = [1,5]
// Output: 4
// Explanation:
// Consider these four triplets:
// (a[0], b[1], c[0]): 1 XOR 3 XOR 1 = 0112
// (a[1], b[1], c[0]): 1 XOR 3 XOR 1 = 0112
// (a[0], b[0], c[1]): 1 XOR 2 XOR 5 = 1102
// (a[1], b[0], c[1]): 1 XOR 2 XOR 5 = 1102
 
// Constraints:
//     1 <= a.length, b.length, c.length <= 100
//     0 <= a[i], b[i], c[i] <= 100

import "fmt"
import "math/bits"

func tripletCount(a []int, b []int, c []int) int {
    res := 0
    for i := 0; i < len(a); i++ {
        for j := 0; j < len(b); j++ {
            for k := 0; k < len(c); k++ {
                if bits.OnesCount(uint(a[i] ^ b[j] ^ c[k])) % 2 == 0 {
                    res++
                }
            } 
        } 
    } 
    return res
}

func tripletCount1(a []int, b []int, c []int) int {
    res, cnt1, cnt2, cnt3 := 0, [2]int{}, [2]int{}, [2]int{}
    for _, v := range a { cnt1[bits.OnesCount(uint(v)) % 2]++ }
    for _, v := range b { cnt2[bits.OnesCount(uint(v)) % 2]++ }
    for _, v := range c { cnt3[bits.OnesCount(uint(v)) % 2]++ }
    for i := 0; i < 2; i++ {
        for j := 0; j < 2; j++ {
            for k := 0; k < 2; k++ {
                if (i+j+k) % 2 == 0 {
                    res += cnt1[i] * cnt2[j] * cnt3[k]
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: a = [1], b = [2], c = [3]
    // Output: 1
    // Explanation:
    // The only triplet is (a[0], b[0], c[0]) and their XOR is: 1 XOR 2 XOR 3 = 002.
    fmt.Println(tripletCount([]int{1},[]int{2},[]int{3})) // 1
    // Example 2:
    // Input: a = [1,1], b = [2,3], c = [1,5]
    // Output: 4
    // Explanation:
    // Consider these four triplets:
    // (a[0], b[1], c[0]): 1 XOR 3 XOR 1 = 0112
    // (a[1], b[1], c[0]): 1 XOR 3 XOR 1 = 0112
    // (a[0], b[0], c[1]): 1 XOR 2 XOR 5 = 1102
    // (a[1], b[0], c[1]): 1 XOR 2 XOR 5 = 1102
    fmt.Println(tripletCount([]int{1,1},[]int{2,3},[]int{1,5})) // 4

    fmt.Println(tripletCount1([]int{1},[]int{2},[]int{3})) // 1
    fmt.Println(tripletCount1([]int{1,1},[]int{2,3},[]int{1,5})) // 4
}