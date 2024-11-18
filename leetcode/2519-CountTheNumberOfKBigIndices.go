package main

// 2519. Count the Number of K-Big Indices
// You are given a 0-indexed integer array nums and a positive integer k.

// We call an index i k-big if the following conditions are satisfied:
//     There exist at least k different indices idx1 such that idx1 < i and nums[idx1] < nums[i].
//     There exist at least k different indices idx2 such that idx2 > i and nums[idx2] < nums[i].

// Return the number of k-big indices.

// Example 1:
// Input: nums = [2,3,6,5,2,3], k = 2
// Output: 2
// Explanation: There are only two 2-big indices in nums:
// - i = 2 --> There are two valid idx1: 0 and 1. There are three valid idx2: 2, 3, and 4.
// - i = 3 --> There are two valid idx1: 0 and 1. There are two valid idx2: 3 and 4.

// Example 2:
// Input: nums = [1,1,1], k = 3
// Output: 0
// Explanation: There are no 3-big indices in nums.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i], k <= nums.length

import "fmt"
import "slices"

type BinaryIndexedTree struct {
    n int
    c []int
}

func NewBinaryIndexedTree(n int) *BinaryIndexedTree {
    return &BinaryIndexedTree{n,  make([]int, n + 1)}
}

func (this *BinaryIndexedTree) update(x, delta int) {
    for x <= this.n {
        this.c[x] += delta
        x += x & -x
    }
}

func (this *BinaryIndexedTree) query(x int) int {
    s := 0
    for x > 0 {
        s += this.c[x]
        x -= x & -x
    }
    return s
}

func kBigIndices(nums []int, k int) int {
    res, n := 0, len(nums)
    tree1, tree2 := NewBinaryIndexedTree(n), NewBinaryIndexedTree(n)
    for _, v := range nums {
        tree2.update(v, 1)
    }
    for _, v := range nums {
        tree2.update(v, -1)
        if tree1.query(v-1) >= k && tree2.query(v-1) >= k {
            res++
        }
        tree1.update(v, 1)
    }
    return res
}

func kBigIndices1(nums []int, k int) int {
    // fenwick
    add := func(tree []int, i int, val int) {
        for ; i < len(tree); i += i & -i {
            tree[i] += val
        }
    }
    preSum := func(tree []int, i int) (sum int) {
        for ; i > 0; i &= i - 1 {
            sum += tree[i]
        }
        return sum
    }
    res, mx := 0, slices.Max(nums)
    leftTree, rightTree := make([]int, mx + 1), make([]int, mx + 1) // nums[i] > 0 坐标无需偏移
    for _, v := range nums {
        add(rightTree, v, 1)
    }
    n := len(leftTree)
    for i, v := range nums {
        add(rightTree, v, -1) // 剩下的为 [i+1:end]的数量
        if i >= k && n-i-1 >= k { // 剪枝! 左边只有i个数, 右边有 n-i-1个数,必须都满足k个才有可能
            if preSum(leftTree, v - 1) >= k && preSum(rightTree, v - 1) >= k {
                res++
            }
        }
        add(leftTree, v, 1)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,3,6,5,2,3], k = 2
    // Output: 2
    // Explanation: There are only two 2-big indices in nums:
    // - i = 2 --> There are two valid idx1: 0 and 1. There are three valid idx2: 2, 3, and 4.
    // - i = 3 --> There are two valid idx1: 0 and 1. There are two valid idx2: 3 and 4.
    fmt.Println(kBigIndices([]int{2,3,6,5,2,3}, 2)) // 2
    // Example 2:
    // Input: nums = [1,1,1], k = 3
    // Output: 0
    // Explanation: There are no 3-big indices in nums.
    fmt.Println(kBigIndices([]int{1,1,1}, 3)) // 0

    fmt.Println(kBigIndices1([]int{2,3,6,5,2,3}, 2)) // 2
    fmt.Println(kBigIndices1([]int{1,1,1}, 3)) // 0
}