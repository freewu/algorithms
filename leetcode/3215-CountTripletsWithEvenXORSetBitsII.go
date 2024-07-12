package main

// 3215. Count Triplets with Even XOR Set Bits II
// Given three integer arrays a, b, and c, return the number of triplets (a[i], b[j], c[k]), 
// such that the bitwise XOR between the elements of each triplet has an even number of set bits.

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
//     1 <= a.length, b.length, c.length <= 10^5
//     0 <= a[i], b[i], c[i] <= 10^9

import "fmt"

func tripletCount(a []int, b []int, c []int) int64 {
    onesCount := func(v int) int { // 统计 1 出现的次数
        res := 0
        for v != 0 {
            res++
            v = v & (v - 1)
        }
        return res
    }
    getEvenOddCounts := func(nums []int) [2]int {
        res := [2]int{}
        for _, v := range nums {
            res[onesCount(v) % 2]++
        }
        return res
    }
    ra, rb, rc := getEvenOddCounts(a), getEvenOddCounts(b), getEvenOddCounts(c)
    return int64(ra[0] * rb[0] * rc[0]) +  int64(ra[0] * rb[1] * rc[1]) + 
           int64(ra[1] * rb[0] * rc[1]) +  int64(ra[1] * rb[1] * rc[0])
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
}