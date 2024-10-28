package main

// 2935. Maximum Strong Pair XOR II
// You are given a 0-indexed integer array nums. 
// A pair of integers x and y is called a strong pair if it satisfies the condition:
//     |x - y| <= min(x, y)

// You need to select two integers from nums such that they form a strong pair 
// and their bitwise XOR is the maximum among all strong pairs in the array.

// Return the maximum XOR value out of all possible strong pairs in the array nums.

// Note that you can pick the same integer twice to form a pair.

// Example 1:
// Input: nums = [1,2,3,4,5]
// Output: 7
// Explanation: There are 11 strong pairs in the array nums: (1, 1), (1, 2), (2, 2), (2, 3), (2, 4), (3, 3), (3, 4), (3, 5), (4, 4), (4, 5) and (5, 5).
// The maximum XOR possible from these pairs is 3 XOR 4 = 7.

// Example 2:
// Input: nums = [10,100]
// Output: 0
// Explanation: There are 2 strong pairs in the array nums: (10, 10) and (100, 100).
// The maximum XOR possible from these pairs is 10 XOR 10 = 0 since the pair (100, 100) also gives 100 XOR 100 = 0.

// Example 3:
// Input: nums = [500,520,2500,3000]
// Output: 1020
// Explanation: There are 6 strong pairs in the array nums: (500, 500), (500, 520), (520, 520), (2500, 2500), (2500, 3000) and (3000, 3000).
// The maximum XOR possible from these pairs is 500 XOR 520 = 1020 since the only other non-zero XOR value is 2500 XOR 3000 = 636.

// Constraints:
//     1 <= nums.length <= 5 * 10^4
//     1 <= nums[i] <= 2^20 - 1

import "fmt"
import "sort"
import "math"

type Trie struct {
    Head   *Node
    MaxBit int
}

type Node struct {
    Child0   *Node
    Child1   *Node
    LastIdx  int
}

func NewTrie(maxBit int) *Trie {
    return &Trie{ &Node{}, maxBit, }
}

func (t *Trie) Insert(num, index int) {
    temp, i := t.Head, t.MaxBit
    for i >= 0 {
        bit := (num & (1 << i))
        if bit == 0 {
            if temp.Child0 == nil {
                temp.Child0 = &Node{}
            }
            temp = temp.Child0
        } else {
            if temp.Child1 == nil {
                temp.Child1 = &Node{}
            }
            temp = temp.Child1
        }
        temp.LastIdx = index
        i--
    }
}

func (t *Trie) GetMaxXor(num, minIdx int) int {
    temp, i, maxXor := t.Head, t.MaxBit, 0
    for i >= 0 {
        bit := (num & (1 << i))
        if bit == 0 {
            if temp.Child1 != nil && temp.Child1.LastIdx >= minIdx {
                temp = temp.Child1
                maxXor |= (1 << i)
            } else {
                temp = temp.Child0
            }
        } else {
            if temp.Child0 != nil && temp.Child0.LastIdx >= minIdx {
                temp = temp.Child0
                maxXor |= (1 << i)
            } else {
                temp = temp.Child1
            }
        }
        i--
    }
    return maxXor
}

func maximumStrongPairXor(nums []int) int {
    sort.Ints(nums)
    n, i, j, res := len(nums), 0, 0, 0
    t := NewTrie(int(math.Log2(float64(nums[n-1]))))
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for j < n {
        t.Insert(nums[j], j)
        half := int(math.Ceil(float64(nums[j]) / 2.0))
        for nums[i] < half {
            i++
        }
        res = max(res, t.GetMaxXor(nums[j], i))
        j++
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4,5]
    // Output: 7
    // Explanation: There are 11 strong pairs in the array nums: (1, 1), (1, 2), (2, 2), (2, 3), (2, 4), (3, 3), (3, 4), (3, 5), (4, 4), (4, 5) and (5, 5).
    // The maximum XOR possible from these pairs is 3 XOR 4 = 7.
    fmt.Println(maximumStrongPairXor([]int{1,2,3,4,5})) // 7
    // Example 2:
    // Input: nums = [10,100]
    // Output: 0
    // Explanation: There are 2 strong pairs in the array nums: (10, 10) and (100, 100).
    // The maximum XOR possible from these pairs is 10 XOR 10 = 0 since the pair (100, 100) also gives 100 XOR 100 = 0.
    fmt.Println(maximumStrongPairXor([]int{10,100})) // 0
    // Example 3:
    // Input: nums = [500,520,2500,3000]
    // Output: 1020
    // Explanation: There are 6 strong pairs in the array nums: (500, 500), (500, 520), (520, 520), (2500, 2500), (2500, 3000) and (3000, 3000).
    // The maximum XOR possible from these pairs is 500 XOR 520 = 1020 since the only other non-zero XOR value is 2500 XOR 3000 = 636.
    fmt.Println(maximumStrongPairXor([]int{500,520,2500,3000})) // 1020
}