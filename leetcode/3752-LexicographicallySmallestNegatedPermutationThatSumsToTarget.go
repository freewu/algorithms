package main

// 3752. Lexicographically Smallest Negated Permutation that Sums to Target
// You are given a positive integer n and an integer target.

// Return the lexicographically smallest array of integers of size n such that:
//     1. The sum of its elements equals target.
//     2. The absolute values of its elements form a permutation of size n.
//     3. If no such array exists, return an empty array.

// A permutation of size n is a rearrangement of integers 1, 2, ..., n.

// Example 1:
// Input: n = 3, target = 0
// Output: [-3,1,2]
// Explanation:
// The arrays that sum to 0 and whose absolute values form a permutation of size 3 are:
// [-3, 1, 2]
// [-3, 2, 1]
// [-2, -1, 3]
// [-2, 3, -1]
// [-1, -2, 3]
// [-1, 3, -2]
// [1, -3, 2]
// [1, 2, -3]
// [2, -3, 1]
// [2, 1, -3]
// [3, -2, -1]
// [3, -1, -2]
// The lexicographically smallest one is [-3, 1, 2].

// Example 2:
// Input: n = 1, target = 10000000000
// Output: []
// Explanation:
// There are no arrays that sum to 10000000000 and whose absolute values form a permutation of size 1. Therefore, the answer is [].

// Constraints:
//     1 <= n <= 10^5
//     -10^10 <= target <= 10^10

import "fmt"
import "sort"

func lexSmallestNegatedPerm(n int, target int64) []int {
    res, s := []int{}, int64(n * (n + 1) / 2)  // 计算1到n的绝对值和
    abs := func(x int64) int64 { if x < 0 { return -x; }; return x; }
    if abs(target) > s || (s + target) % 2 != 0 { // 可行性判断
        return res
    }
    diff := (s - target) / 2
    if diff < 0 || diff > s {
        return res
    }
    // 从大到小遍历
    for i := n; i >= 1; i-- {
        if int64(i) <= diff {
            res = append(res, -i)
            diff -= int64(i)
        } else {
            res = append(res, i)
        }
    }
    sort.Ints(res)
    return res
}

func lexSmallestNegatedPerm1(n int, target int64) []int {
    abs := func(x int64) int64 { if x < 0 { return -x; }; return x; }
    s := int64(n) * int64(n+1) / 2
    if abs(target) > s { return nil }
    if ((s - target) & 1) != 0 { return nil }
    u := (s - target) / 2
    neg := make([]bool, n + 1)
    for i := n; i >= 1; i-- {
        if int64(i) <= u {
            neg[i] = true
            u -= int64(i)
        }
    }
    if u != 0 { return nil }
    res := make([]int, 0, n)
    for i := n; i >= 1; i-- {
        if neg[i] { res = append(res, -i) }
    }
    for i := 1; i <= n; i++ {
        if !neg[i] { res = append(res, i) }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 3, target = 0
    // Output: [-3,1,2]
    // Explanation:
    // The arrays that sum to 0 and whose absolute values form a permutation of size 3 are:
    // [-3, 1, 2]
    // [-3, 2, 1]
    // [-2, -1, 3]
    // [-2, 3, -1]
    // [-1, -2, 3]
    // [-1, 3, -2]
    // [1, -3, 2]
    // [1, 2, -3]
    // [2, -3, 1]
    // [2, 1, -3]
    // [3, -2, -1]
    // [3, -1, -2]
    // The lexicographically smallest one is [-3, 1, 2].
    fmt.Println(lexSmallestNegatedPerm(3, 0)) // [-3, 1, 2]
    // Example 2:
    // Input: n = 1, target = 10000000000
    // Output: []
    // Explanation:
    // There are no arrays that sum to 10000000000 and whose absolute values form a permutation of size 1. Therefore, the answer is [].
    fmt.Println(lexSmallestNegatedPerm(1, 10000000000)) // []

    fmt.Println(lexSmallestNegatedPerm(1, 1)) // [1]
    fmt.Println(lexSmallestNegatedPerm(1, 0)) // []
    fmt.Println(lexSmallestNegatedPerm(1, 10_000_000_000)) // []
    fmt.Println(lexSmallestNegatedPerm(1, -10_000_000_000)) // []
    fmt.Println(lexSmallestNegatedPerm(100_000, 1)) // []
    //fmt.Println(lexSmallestNegatedPerm(100_000, 0)) // [
    fmt.Println(lexSmallestNegatedPerm(100_000, 10_000_000_000)) // []
    fmt.Println(lexSmallestNegatedPerm(100_000, -10_000_000_000)) // []

    fmt.Println(lexSmallestNegatedPerm1(3, 0)) // [-3, 1, 2]
    fmt.Println(lexSmallestNegatedPerm1(1, 10000000000)) // []
    fmt.Println(lexSmallestNegatedPerm1(1, 1)) // [1]
    fmt.Println(lexSmallestNegatedPerm1(1, 0)) // []
    fmt.Println(lexSmallestNegatedPerm1(1, 10_000_000_000)) // []
    fmt.Println(lexSmallestNegatedPerm1(1, -10_000_000_000)) // []
    fmt.Println(lexSmallestNegatedPerm1(100_000, 1)) // []
    //fmt.Println(lexSmallestNegatedPerm1(100_000, 0)) // [
    fmt.Println(lexSmallestNegatedPerm1(100_000, 10_000_000_000)) // []
    fmt.Println(lexSmallestNegatedPerm1(100_000, -10_000_000_000)) // []
}