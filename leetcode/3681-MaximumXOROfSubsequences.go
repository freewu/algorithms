package main

// 3681. Maximum XOR of Subsequences
// You are given an integer array nums of length n where each element is a non-negative integer.

// Select two subsequences of nums (they may be empty and are allowed to overlap), 
// each preserving the original order of elements, and let:
//     1. X be the bitwise XOR of all elements in the first subsequence.
//     2. Y be the bitwise XOR of all elements in the second subsequence.

// Return the maximum possible value of X XOR Y.

// Note: The XOR of an empty subsequence is 0.

// Example 1:
// Input: nums = [1,2,3]
// Output: 3
// Explanation:
// Choose subsequences:
//     First subsequence [2], whose XOR is 2.
//     Second subsequence [2,3], whose XOR is 1.
//     Then, XOR of both subsequences = 2 XOR 1 = 3.
// This is the maximum XOR value achievable from any two subsequences.

// Example 2:
// Input: nums = [5,2]
// Output: 7
// Explanation:
// Choose subsequences:
//     First subsequence [5], whose XOR is 5.
//     Second subsequence [2], whose XOR is 2.
//     Then, XOR of both subsequences = 5 XOR 2 = 7.
// This is the maximum XOR value achievable from any two subsequences.

// Constraints:
//     2 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^9

import "fmt"
import "slices"
import "math/bits"

type XorBasis []int

// n 为值域最大值 U 的二进制长度，例如 U=1e9 时 n=30
func newXorBasis(n int) XorBasis {
    return make(XorBasis, n)
}

func (b XorBasis) insert(x int) {
    // 从高到低遍历，保证计算 maxXor 的时候，参与 XOR 的基的最高位（或者说二进制长度）是互不相同的
    for i := len(b) - 1; i >= 0; i-- {
        if x>>i == 0 { // 由于大于 i 的位都被我们异或成了 0，所以 x>>i 的结果只能是 0 或 1
            continue
        }
        if b[i] == 0 { // x 和之前的基是线性无关的
            b[i] = x // 新增一个基，最高位为 i
            return
        }
        x ^= b[i] // 保证每个基的二进制长度互不相同
    }
    // 正常循环结束，此时 x=0，说明一开始的 x 可以被已有基表出，不是一个线性无关基
}

func (b XorBasis) maxXor() (res int) {
    // 从高到低贪心：越高的位，越必须是 1
    // 由于每个位的基至多一个，所以每个位只需考虑异或一个基，若能变大，则异或之
    for i := len(b) - 1; i >= 0; i-- {
        res = max(res, res^b[i])
    }
    return
}

// 线性基
func maxXorSubsequences(nums []int) int {
    n := bits.Len(uint(slices.Max(nums)))
    b := newXorBasis(n)
    for _, v := range nums {
        b.insert(v)
    }
    return b.maxXor()
}

func maxXorSubsequences1(nums []int) int {
    res, basis := 0, []int{}
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, x := range nums {
        for _, b := range basis {
            x = min(x, x^b)
        }
        if x > 0 {
            basis = append(basis, x)
            slices.SortFunc(basis, func(i, j int) int {
                return j - i
            })
        }
    }
    for _, b := range basis {
        res = max(res, res^b)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3]
    // Output: 3
    // Explanation:
    // Choose subsequences:
    //     First subsequence [2], whose XOR is 2.
    //     Second subsequence [2,3], whose XOR is 1.
    //     Then, XOR of both subsequences = 2 XOR 1 = 3.
    // This is the maximum XOR value achievable from any two subsequences.
    fmt.Println(maxXorSubsequences([]int{1,2,3})) // 3
    // Example 2:
    // Input: nums = [5,2]
    // Output: 7
    // Explanation:
    // Choose subsequences:
    //     First subsequence [5], whose XOR is 5.
    //     Second subsequence [2], whose XOR is 2.
    //     Then, XOR of both subsequences = 5 XOR 2 = 7.
    // This is the maximum XOR value achievable from any two subsequences.
    fmt.Println(maxXorSubsequences([]int{5,2})) // 7

    fmt.Println(maxXorSubsequences([]int{1,2,3,4,5,6,7,8,9})) // 15
    fmt.Println(maxXorSubsequences([]int{9,8,7,6,5,4,3,2,1})) // 15

    fmt.Println(maxXorSubsequences1([]int{1,2,3})) // 3
    fmt.Println(maxXorSubsequences1([]int{5,2})) // 7
    fmt.Println(maxXorSubsequences1([]int{1,2,3,4,5,6,7,8,9})) // 15
    fmt.Println(maxXorSubsequences1([]int{9,8,7,6,5,4,3,2,1})) // 15
}