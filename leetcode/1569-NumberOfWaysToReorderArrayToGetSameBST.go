package main

// 1569. Number of Ways to Reorder Array to Get Same BST
// Given an array nums that represents a permutation of integers from 1 to n. 
// We are going to construct a binary search tree (BST) by inserting the elements of nums in order into an initially empty BST. 
// Find the number of different ways to reorder nums so that the constructed BST is identical to that formed from the original array nums.
//     For example, given nums = [2,1,3], we will have 2 as the root, 1 as a left child, and 3 as a right child. 
//     The array [2,3,1] also yields the same BST but [3,2,1] yields a different BST.

// Return the number of ways to reorder nums such that the BST formed is identical to the original BST formed from nums.
// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
//      2
//    /   \
//   1     3
// <img src="https://assets.leetcode.com/uploads/2020/08/12/bb.png" />
// Input: nums = [2,1,3]
// Output: 1
// Explanation: We can reorder nums to be [2,3,1] which will yield the same BST. There are no other ways to reorder nums which will yield the same BST.

// Example 2:
//         3
//       /   \
//      1     4
//       \     \ 
//        2      3
// <img src="https://assets.leetcode.com/uploads/2020/08/12/ex1.png" />
// Input: nums = [3,4,5,1,2]
// Output: 5
// Explanation: The following 5 arrays will yield the same BST: 
// [3,1,2,4,5]
// [3,1,4,2,5]
// [3,1,4,5,2]
// [3,4,1,2,5]
// [3,4,1,5,2]

// Example 3:
//     1
//      \ 
//       2 
//        \ 
//         3
// <img src="https://assets.leetcode.com/uploads/2020/08/12/ex4.png" />
// Input: nums = [1,2,3]
// Output: 0
// Explanation: There are no other orderings of nums that will yield the same BST.
 
// Constraints:
//     1 <= nums.length <= 1000
//     1 <= nums[i] <= nums.length
//     All integers in nums are distinct.

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func numOfWays(nums []int) int {
    binom, mod := make([][]int,len(nums)),1_000_000_007
    for i:=0; i < len(nums); i++ {
        binom[i] = make([]int,i+1)
        binom[i][0], binom[i][i] = 1, 1
        for j := 1; j < i; j++ {
            binom[i][j]= (binom[i-1][j]+binom[i-1][j-1]) % mod
        }
    }
    var count func( nums []int) int
    count = func ( nums []int) int {
        n := len(nums)
        if n <= 2 {
            return 1
        }
        root,left, right := nums[0], []int{}, []int{}
        for i := 1; i < n; i++ {
            if nums[i] < root {
                left = append(left, nums[i])
            } else {
                right = append(right, nums[i])
            }
        }
        lc := count(left) % mod
        rc := count(right) % mod
        return (((lc * rc) % mod) * binom[n-1][len(left)] ) % mod
    }
    return count(nums) - 1
}

func numOfWays1(nums []int) int {
    type node struct {
        l, r  int
        i     int
        left  *node
        right *node
    }    
    n, mod := len(nums), 1_000_000_007
    caches := make([]int, n + 1)
    caches[0] = 1
    for i := 1; i < n + 1; i++ {
        caches[i] = caches[i-1] * i % mod
    }
    cacheDivs := make([]int, n + 1)
    getDiv := func(i int) int {
        if cacheDivs[i] != 0 {
            return cacheDivs[i]
        }
        res, r := 1, caches[i]
        for k := int64(mod - 2); k != 0; k >>= 1 {
            if k & 1 == 1 {
                res = (res * r) % mod
            }
            r = r * r % mod
        }
        cacheDivs[i] = res
        return res
    }
    root := node{ l: 1, r: len(nums), i: nums[0] }

    for i := 1; i < len(nums); i++ {
        cur := &root
        for {
            if cur.i > nums[i] {
                if cur.left != nil {
                    cur = cur.left
                    continue
                }
                cur.left = &node{ l: cur.l, r: cur.i - 1, i: nums[i], }
                break
            } else {
                if cur.right != nil {
                    cur = cur.right
                    continue
                }
                cur.right = &node{ l: cur.i + 1, r: cur.r, i: nums[i], }
                break
            }
        }
    }
    var getRes func(cur *node) int
    getRes = func(cur *node) int {
        if cur == nil {
            return 1
        }
        res := 1
        res = res * getRes(cur.right) % mod
        res = res * getRes(cur.left) % mod
        n, r := cur.r - cur.l, cur.i - cur.l
        res = res * caches[n] % mod * getDiv(n-r) % mod * getDiv(r) % mod
        return res
    }
    return (getRes(&root) - 1 + mod) % mod
}

func main() {
    // Example 1:
    //      2
    //    /   \
    //   1     3
    // <img src="https://assets.leetcode.com/uploads/2020/08/12/bb.png" />
    // Input: nums = [2,1,3]
    // Output: 1
    // Explanation: We can reorder nums to be [2,3,1] which will yield the same BST. There are no other ways to reorder nums which will yield the same BST.
    tree1 := &TreeNode{
        2, 
        &TreeNode{1, nil, nil},
        &TreeNode{3, nil, nil},
    }
    fmt.Println(tree1)
    fmt.Println(numOfWays([]int{2,1,3})) // 1
    // Example 2:
    //         3
    //       /   \
    //      1     4
    //       \     \ 
    //        2     5
    // <img src="https://assets.leetcode.com/uploads/2020/08/12/ex1.png" />
    // Input: nums = [3,4,5,1,2]
    // Output: 5
    // Explanation: The following 5 arrays will yield the same BST: 
    // [3,1,2,4,5]
    // [3,1,4,2,5]
    // [3,1,4,5,2]
    // [3,4,1,2,5]
    // [3,4,1,5,2]
    tree2 := &TreeNode{
        3, 
        &TreeNode{1, nil, &TreeNode{2, nil, nil}, },
        &TreeNode{4, nil, &TreeNode{5, nil, nil}, },
    }
    fmt.Println(tree2)
    fmt.Println(numOfWays([]int{3,4,5,1,2})) // 5
    // Example 3:
    //     1
    //      \ 
    //       2 
    //        \ 
    //         3
    // <img src="https://assets.leetcode.com/uploads/2020/08/12/ex4.png" />
    // Input: nums = [1,2,3]
    // Output: 0
    // Explanation: There are no other orderings of nums that will yield the same BST.
    tree3 := &TreeNode{
        1, 
        nil,
        &TreeNode{2, nil,  &TreeNode{1, nil, nil}, },
    }
    fmt.Println(tree3)
    fmt.Println(numOfWays([]int{1,2,3})) // 0

    fmt.Println(numOfWays1([]int{2,1,3})) // 1
    fmt.Println(numOfWays1([]int{3,4,5,1,2})) // 5
    fmt.Println(numOfWays1([]int{1,2,3})) // 0
}