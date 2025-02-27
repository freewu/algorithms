package main

// 面试题 10.01. Sorted Merge LCCI
// You are given two sorted arrays, A and B, where A has a large enough buffer at the end to hold B. 
// Write a method to merge B into A in sorted order.

// Initially the number of elements in A and B are m and n respectively.

// Example:
// Input:
// A = [1,2,3,0,0,0], m = 3
// B = [2,5,6],       n = 3
// Output: [1,2,2,3,5,6]

// Note:
//     A.length == n + m

import "fmt"

func merge(A []int, m int, B []int, n int)  {
    i, j, k := m - 1, n - 1, m + n - 1
    for i >= 0 && j >= 0 {
        if A[i] > B[j] {
            A[k] = A[i]
            i--
        } else {
            A[k] = B[j]
            j--
        }
        k--
    }
    if j >= 0 {
        copy(A[:k+1], B[:j+1])
    }
}

func main() {
    // Example:
    // Input:
    // A = [1,2,3,0,0,0], m = 3
    // B = [2,5,6],       n = 3
    // Output: [1,2,2,3,5,6]
    A := []int{1,2,3,0,0,0}
    B := []int{2,5,6}
    fmt.Println("A: ", A)
    fmt.Println("B: ", B)
    merge(A, 3, B, 3)
    fmt.Println("A: ", A)
}