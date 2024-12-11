package main

// 2179. Count Good Triplets in an Array
// You are given two 0-indexed arrays nums1 and nums2 of length n, both of which are permutations of [0, 1, ..., n - 1].

// A good triplet is a set of 3 distinct values which are present in increasing order by position both in nums1 and nums2. 
// In other words, if we consider pos1v as the index of the value v in nums1 and pos2v as the index of the value v in nums2, 
// then a good triplet will be a set (x, y, z) where 0 <= x, y, z <= n - 1, 
// such that pos1x < pos1y < pos1z and pos2x < pos2y < pos2z.

// Return the total number of good triplets.

// Example 1:
// Input: nums1 = [2,0,1,3], nums2 = [0,1,2,3]
// Output: 1
// Explanation: 
// There are 4 triplets (x,y,z) such that pos1x < pos1y < pos1z. They are (2,0,1), (2,0,3), (2,1,3), and (0,1,3). 
// Out of those triplets, only the triplet (0,1,3) satisfies pos2x < pos2y < pos2z. Hence, there is only 1 good triplet.

// Example 2:
// Input: nums1 = [4,0,1,3,2], nums2 = [4,1,0,2,3]
// Output: 4
// Explanation: The 4 good triplets are (4,0,3), (4,0,2), (4,1,3), and (4,1,2).

// Constraints:
//     n == nums1.length == nums2.length
//     3 <= n <= 105
//     0 <= nums1[i], nums2[i] <= n - 1
//     nums1 and nums2 are permutations of [0, 1, ..., n - 1].

import "fmt"

type BinaryIndexedTree struct {
    n int
    data []int
}

func NewBinaryIndexedTree(n int) *BinaryIndexedTree {
    return &BinaryIndexedTree{n, make([]int, n + 1)}
}

func (this *BinaryIndexedTree) Lowbit(x int) int {
    return x & -x
}

func (this *BinaryIndexedTree) Update(x, delta int) {
    for x <= this.n {
        this.data[x] += delta
        x += this.Lowbit(x)
    }
}

func (this *BinaryIndexedTree) Query(x int) int {
    sum := 0
    for x > 0 {
        sum += this.data[x]
        x -= this.Lowbit(x)
    }
    return sum
}

func goodTriplets(nums1 []int, nums2 []int) int64 {
    res, n := 0, len(nums1)
    pos := make([]int, n)
    for i, v := range nums2 {
        pos[v] = i + 1
    }
    tree := NewBinaryIndexedTree(n)
    for _, v := range nums1 {
        left := tree.Query(pos[v])
        right := n - pos[v] - (tree.Query(n) - tree.Query(pos[v]))
        res += left * right
        tree.Update(pos[v], 1)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums1 = [2,0,1,3], nums2 = [0,1,2,3]
    // Output: 1
    // Explanation: 
    // There are 4 triplets (x,y,z) such that pos1x < pos1y < pos1z. They are (2,0,1), (2,0,3), (2,1,3), and (0,1,3). 
    // Out of those triplets, only the triplet (0,1,3) satisfies pos2x < pos2y < pos2z. Hence, there is only 1 good triplet.
    fmt.Println(goodTriplets([]int{2,0,1,3}, []int{0,1,2,3})) // 1
    // Example 2:
    // Input: nums1 = [4,0,1,3,2], nums2 = [4,1,0,2,3]
    // Output: 4
    // Explanation: The 4 good triplets are (4,0,3), (4,0,2), (4,1,3), and (4,1,2).
    fmt.Println(goodTriplets([]int{4,0,1,3,2}, []int{4,1,0,2,3})) // 4
}