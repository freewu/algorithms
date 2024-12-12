package main

// 3109. Find the Index of Permutation
// Given an array perm of length n which is a permutation of [1, 2, ..., n], 
// return the index of perm in the lexicographically sorted array of all of the permutations of [1, 2, ..., n].

// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: perm = [1,2]
// Output: 0
// Explanation:
// There are only two permutations in the following order:
// [1,2], [2,1]
// And [1,2] is at index 0.

// Example 2:
// Input: perm = [3,1,2]
// Output: 4
// Explanation:
// There are only six permutations in the following order:
// [1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1]
// And [3,1,2] is at index 4.

// Constraints:
//     1 <= n == perm.length <= 10^5
//     perm is a permutation of [1, 2, ..., n].

import "fmt"

type BinaryIndexedTree struct {
    capacity int
    tree     []int
}

func NewBinaryIndexedTree(n int) *BinaryIndexedTree {
    return &BinaryIndexedTree{  n,  make([]int, n + 1) }
}

func (bit *BinaryIndexedTree) Update(x, delta int) {
    for ; x <= bit.capacity; x += x & -x {
        bit.tree[x] += delta
    }
}

func (bit *BinaryIndexedTree) Query(x int) int {
    sum := 0
    for ; x > 0; x -= x & -x {
        sum += bit.tree[x]
    }
    return sum
}

func getPermutationIndex(perm []int) int {
    res, n, mod := 0, len(perm), 1_000_000_007
    tree := NewBinaryIndexedTree(n + 1)
    facts := make([]int, n)
    facts[0] = 1
    for i := 1; i < n; i++ {
        facts[i] = facts[i-1] * i % mod
    }
    for i, v := range perm {
        count := v - 1 - tree.Query(v)
        res += count * facts[n - i - 1] % mod
        tree.Update(v, 1)
    }
    return res % mod
}

func main() {
    // Example 1:
    // Input: perm = [1,2]
    // Output: 0
    // Explanation:
    // There are only two permutations in the following order:
    // [1,2], [2,1]
    // And [1,2] is at index 0.
    fmt.Println(getPermutationIndex([]int{1,2})) // 0
    // Example 2:
    // Input: perm = [3,1,2]
    // Output: 4
    // Explanation:
    // There are only six permutations in the following order:
    // [1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1]
    // And [3,1,2] is at index 4.
    fmt.Println(getPermutationIndex([]int{3,1,2})) // 4
}