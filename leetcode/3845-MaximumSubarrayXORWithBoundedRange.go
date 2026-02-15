package main

// 3845. Maximum Subarray XOR with Bounded Range
// You are given a non-negative integer array nums and an integer k.

// You must select a subarray of nums such that the difference between its maximum and minimum elements is at most k. 
// The value of this subarray is the bitwise XOR of all elements in the subarray.

// Return an integer denoting the maximum possible value of the selected subarray.
 
// Example 1:
// Input: nums = [5,4,5,6], k = 2
// Output: 7
// Explanation:
// Select the subarray [5, 4, 5, 6].
// The difference between its maximum and minimum elements is 6 - 4 = 2 <= k.
// The value is 4 XOR 5 XOR 6 = 7.

// Example 2:
// Input: nums = [5,4,5,6], k = 1
// Output: 6
// Explanation:
// Select the subarray [5, 4, 5, 6].
// The difference between its maximum and minimum elements is 6 - 6 = 0 <= k.
// The value is 6.
 
// Constraints:
//     1 <= nums.length <= 4 * 10^4
//     0 <= nums[i] < 2^15
//     0 <= k < 2^15

import "fmt"

const WIDTH = 15 // nums[i] 二进制长度的最大值

type Node struct {
    son  [2]*Node
    leaf int // 子树叶子个数
}

type Trie struct {
    root *Node
}

func NewTrie() *Trie {
    return &Trie{&Node{}}
}

func (t *Trie) put(val int) {
    cur := t.root
    for i := WIDTH - 1; i >= 0; i-- {
        bit := val >> i & 1
        if cur.son[bit] == nil {
            cur.son[bit] = &Node{}
        }
        cur = cur.son[bit]
        cur.leaf++
    }
}

func (t *Trie) del(val int) {
    cur := t.root
    for i := WIDTH - 1; i >= 0; i-- {
        cur = cur.son[val>>i&1]
        cur.leaf-- // 如果减成 0 了，说明子树是空的，可以理解成 cur == nil
    }
}

func (t *Trie) maxXor(val int) int {
    res, cur := 0,t.root
    for i := WIDTH - 1; i >= 0; i-- {
        bit := val >> i & 1
        if cur.son[bit^1] != nil && cur.son[bit^1].leaf > 0 {
            res |= 1 << i
            bit ^= 1
        }
        cur = cur.son[bit]
    }
    return res
}

func maxXor(nums []int, k int) int {
    res, sum := 0, make([]int, len(nums) + 1)
    for i, x := range nums {
        sum[i+1] = sum[i] ^ x
    }
    t := NewTrie()
    var minQ, maxQ []int
    left := 0
    for right, x := range nums {
        // 1. 入
        t.put(sum[right])
        for len(minQ) > 0 && x <= nums[minQ[len(minQ)-1]] {
            minQ = minQ[:len(minQ)-1]
        }
        minQ = append(minQ, right)
        for len(maxQ) > 0 && x >= nums[maxQ[len(maxQ)-1]] {
            maxQ = maxQ[:len(maxQ)-1]
        }
        maxQ = append(maxQ, right)
        // 2. 出
        for nums[maxQ[0]]-nums[minQ[0]] > k {
            t.del(sum[left])
            left++
            if minQ[0] < left {
                minQ = minQ[1:]
            }
            if maxQ[0] < left {
                maxQ = maxQ[1:]
            }
        }
        // 3. 更新答案
        res = max(res, t.maxXor(sum[right+1]))
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [5,4,5,6], k = 2
    // Output: 7
    // Explanation:
    // Select the subarray [5, 4, 5, 6].
    // The difference between its maximum and minimum elements is 6 - 4 = 2 <= k.
    // The value is 4 XOR 5 XOR 6 = 7.
    fmt.Println(maxXor([]int{5,4,5,6}, 2)) // 7
    // Example 2:
    // Input: nums = [5,4,5,6], k = 1
    // Output: 6
    // Explanation:
    // Select the subarray [5, 4, 5, 6].
    // The difference between its maximum and minimum elements is 6 - 6 = 0 <= k.
    // The value is 6.
    fmt.Println(maxXor([]int{5,4,5,6}, 1)) // 6

    fmt.Println(maxXor([]int{1,2,3,4,5,6,7,8,9}, 2)) // 15
    fmt.Println(maxXor([]int{9,8,7,6,5,4,3,2,1}, 2)) // 15
}