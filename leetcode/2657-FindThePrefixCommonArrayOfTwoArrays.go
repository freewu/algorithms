package main

// 2657. Find the Prefix Common Array of Two Arrays
// You are given two 0-indexed integer permutations A and B of length n.

// A prefix common array of A and B is an array C such that C[i] is equal to the count of numbers that are present at or before the index i in both A and B.

// Return the prefix common array of A and B.

// A sequence of n integers is called a permutation if it contains all integers from 1 to n exactly once.

// Example 1:
// Input: A = [1,3,2,4], B = [3,1,2,4]
// Output: [0,2,3,4]
// Explanation: At i = 0: no number is common, so C[0] = 0.
// At i = 1: 1 and 3 are common in A and B, so C[1] = 2.
// At i = 2: 1, 2, and 3 are common in A and B, so C[2] = 3.
// At i = 3: 1, 2, 3, and 4 are common in A and B, so C[3] = 4.

// Example 2:
// Input: A = [2,3,1], B = [3,1,2]
// Output: [0,1,3]
// Explanation: At i = 0: no number is common, so C[0] = 0.
// At i = 1: only 3 is common in A and B, so C[1] = 1.
// At i = 2: 1, 2, and 3 are common in A and B, so C[2] = 3.

// Constraints:
//     1 <= A.length == B.length == n <= 50
//     1 <= A[i], B[i] <= n
//     It is guaranteed that A and B are both a permutation of n integers.

import "fmt"
import "math/bits"

func findThePrefixCommonArray1(A []int, B []int) []int {
    res := make([]int, len(A))
    var a, b uint
    for i, v := range A {
        a |= 1 << v
        b |= 1 << B[i]
        res[i] = bits.OnesCount(a & b)
    }
    return res
}

func findThePrefixCommonArray(A []int, B []int) []int {
    n := len(A)
    res, count1, count2 := make([]int, n), make([]int, n + 1), make([]int, n + 1)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i, a := range A {
        b := B[i]
        count1[a]++
        count2[b]++
        for j := 1; j <= n; j++ {
            res[i] += min(count1[j], count2[j])
        }
    }
    return res
}

func findThePrefixCommonArray2(A []int, B []int) []int {
    n := len(A)
    mp := make(map[int]int, n) // A 1 B 2 AB 3
    res, mpab := make([]int, n), make(map[int]struct{}, n)
    for i := 0; i < n; i++ {
        if v, ok := mp[A[i]]; ok {
            if v == 2 { // B 也存在
                mp[A[i]] = 3
                mpab[A[i]] = struct{}{}
            }
        } else {
            mp[A[i]] = 1
        }
        if v, ok := mp[B[i]]; ok {
            if v == 1 { // A 中存在
                mp[B[i]] = 3
                mpab[B[i]] = struct{}{}
            }
        } else {
            mp[B[i]] = 2
        }
        res[i] = len(mpab)
    }
    return res
}

func main() {
    // Example 1:
    // Input: A = [1,3,2,4], B = [3,1,2,4]
    // Output: [0,2,3,4]
    // Explanation: At i = 0: no number is common, so C[0] = 0.
    // At i = 1: 1 and 3 are common in A and B, so C[1] = 2.
    // At i = 2: 1, 2, and 3 are common in A and B, so C[2] = 3.
    // At i = 3: 1, 2, 3, and 4 are common in A and B, so C[3] = 4.
    fmt.Println(findThePrefixCommonArray([]int{1,3,2,4}, []int{3,1,2,4})) // [0,2,3,4]
    // Example 2:
    // Input: A = [2,3,1], B = [3,1,2]
    // Output: [0,1,3]
    // Explanation: At i = 0: no number is common, so C[0] = 0.
    // At i = 1: only 3 is common in A and B, so C[1] = 1.
    // At i = 2: 1, 2, and 3 are common in A and B, so C[2] = 3.
    fmt.Println(findThePrefixCommonArray([]int{2,3,1}, []int{3,1,2})) // [0,1,3]

    fmt.Println(findThePrefixCommonArray([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // [0 0 0 0 1 3 5 7 9]
    fmt.Println(findThePrefixCommonArray([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // [1 2 3 4 5 6 7 8 9]
    fmt.Println(findThePrefixCommonArray([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // [0 0 0 0 1 3 5 7 9]

    fmt.Println(findThePrefixCommonArray1([]int{1,3,2,4}, []int{3,1,2,4})) // [0,2,3,4]
    fmt.Println(findThePrefixCommonArray1([]int{2,3,1}, []int{3,1,2})) // [0,1,3]
    fmt.Println(findThePrefixCommonArray1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // [0 0 0 0 1 3 5 7 9]
    fmt.Println(findThePrefixCommonArray1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // [1 2 3 4 5 6 7 8 9]
    fmt.Println(findThePrefixCommonArray1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // [0 0 0 0 1 3 5 7 9]

    fmt.Println(findThePrefixCommonArray2([]int{1,3,2,4}, []int{3,1,2,4})) // [0,2,3,4]
    fmt.Println(findThePrefixCommonArray2([]int{2,3,1}, []int{3,1,2})) // [0,1,3]
    fmt.Println(findThePrefixCommonArray2([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // [0 0 0 0 1 3 5 7 9]
    fmt.Println(findThePrefixCommonArray2([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // [1 2 3 4 5 6 7 8 9]
    fmt.Println(findThePrefixCommonArray2([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // [0 0 0 0 1 3 5 7 9]
}