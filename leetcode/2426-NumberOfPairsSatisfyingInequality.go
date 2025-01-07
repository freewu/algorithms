package main

// 2426. Number of Pairs Satisfying Inequality
// You are given two 0-indexed integer arrays nums1 and nums2, each of size n, and an integer diff. 
// Find the number of pairs (i, j) such that:
//     0 <= i < j <= n - 1 and
//     nums1[i] - nums1[j] <= nums2[i] - nums2[j] + diff.

// Return the number of pairs that satisfy the conditions.

// Example 1:
// Input: nums1 = [3,2,5], nums2 = [2,2,1], diff = 1
// Output: 3
// Explanation:
// There are 3 pairs that satisfy the conditions:
// 1. i = 0, j = 1: 3 - 2 <= 2 - 2 + 1. Since i < j and 1 <= 1, this pair satisfies the conditions.
// 2. i = 0, j = 2: 3 - 5 <= 2 - 1 + 1. Since i < j and -2 <= 2, this pair satisfies the conditions.
// 3. i = 1, j = 2: 2 - 5 <= 2 - 1 + 1. Since i < j and -3 <= 2, this pair satisfies the conditions.
// Therefore, we return 3.

// Example 2:
// Input: nums1 = [3,-1], nums2 = [-2,2], diff = -1
// Output: 0
// Explanation:
// Since there does not exist any pair that satisfies the conditions, we return 0.

// Constraints:
//     n == nums1.length == nums2.length
//     2 <= n <= 10^5
//     -10^4 <= nums1[i], nums2[i] <= 10^4
//     -10^4 <= diff <= 10^4

import "fmt"


// brute force
func numberOfPairs(nums1 []int, nums2 []int, diff int) int64 {
    res := 0
    for i := 0; i < len(nums1); i++ {
        for j := i + 1; j < len(nums1); j++ {
            if nums1[i] - nums1[j] <= nums2[i] - nums2[j] + diff {
                res++
            }
        }
    }
    return int64(res)
}

func numberOfPairs1(nums1 []int, nums2 []int, diff int) int64 {
    res, n := 0, len(nums1)
    sortedDiffs := make([]int, 0)
    // Find the proper position to insert the value
    getInsertPos := func(arr []int, v int) int {
        l, r := 0, len(arr)
        for l < r {
            mid := l + (r - l) / 2
            if v < arr[mid] {
                r = mid
            } else {
                l = mid + 1
            }
        }
        return l
    }
    for i := 0; i < n; i++ {
        num := nums1[i] - nums2[i]
        // check what position will be for the `num+diff` in the
        // sorted array. It will be the number of the values to be added to the result
        res += getInsertPos(sortedDiffs, num + diff)
        pos := getInsertPos(sortedDiffs, num)
        // insert the value to the position
        sortedDiffs = append(sortedDiffs, 0)
        copy(sortedDiffs[pos + 1:], sortedDiffs[pos:])
        sortedDiffs[pos] = nums1[i] - nums2[i]
    }
    return int64(res)
}

type BinaryIndexedTree struct {
    n int
    c []int
}

func NewBinaryIndexedTree(n int) *BinaryIndexedTree {
    return &BinaryIndexedTree{ n, make([]int, n + 1) }
}

func (this *BinaryIndexedTree) lowbit(x int) int {
    return x & -x
}

func (this *BinaryIndexedTree) update(x, delta int) {
    for x <= this.n {
        this.c[x] += delta
        x += this.lowbit(x)
    }
}

func (this *BinaryIndexedTree) query(x int) int {
    s := 0
    for x > 0 {
        s += this.c[x]
        x -= this.lowbit(x)
    }
    return s
}

func numberOfPairs2(nums1 []int, nums2 []int, diff int) int64 {
    tree := NewBinaryIndexedTree(100000)
    res := 0
    for i := range nums1 {
        v := nums1[i] - nums2[i]
        res += tree.query(v + diff + 40000)
        tree.update(v + 40000, 1)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums1 = [3,2,5], nums2 = [2,2,1], diff = 1
    // Output: 3
    // Explanation:
    // There are 3 pairs that satisfy the conditions:
    // 1. i = 0, j = 1: 3 - 2 <= 2 - 2 + 1. Since i < j and 1 <= 1, this pair satisfies the conditions.
    // 2. i = 0, j = 2: 3 - 5 <= 2 - 1 + 1. Since i < j and -2 <= 2, this pair satisfies the conditions.
    // 3. i = 1, j = 2: 2 - 5 <= 2 - 1 + 1. Since i < j and -3 <= 2, this pair satisfies the conditions.
    // Therefore, we return 3.
    fmt.Println(numberOfPairs([]int{3,2,5}, []int{2,2,1}, 1)) // 3
    // Example 2:
    // Input: nums1 = [3,-1], nums2 = [-2,2], diff = -1
    // Output: 0
    // Explanation:
    // Since there does not exist any pair that satisfies the conditions, we return 0.
    fmt.Println(numberOfPairs([]int{3,-1}, []int{-2,2}, -1)) // 0

    fmt.Println(numberOfPairs1([]int{3,2,5}, []int{2,2,1}, 1)) // 3
    fmt.Println(numberOfPairs1([]int{3,-1}, []int{-2,2}, -1)) // 0

    fmt.Println(numberOfPairs2([]int{3,2,5}, []int{2,2,1}, 1)) // 3
    fmt.Println(numberOfPairs2([]int{3,-1}, []int{-2,2}, -1)) // 0
}