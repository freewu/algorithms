package main

// 1803. Count Pairs With XOR in a Range
// Given a (0-indexed) integer array nums and two integers low and high, return the number of nice pairs.

// A nice pair is a pair (i, j) where 0 <= i < j < nums.length and low <= (nums[i] XOR nums[j]) <= high.

// Example 1:
// Input: nums = [1,4,2,7], low = 2, high = 6
// Output: 6
// Explanation: All nice pairs (i, j) are as follows:
//     - (0, 1): nums[0] XOR nums[1] = 5 
//     - (0, 2): nums[0] XOR nums[2] = 3
//     - (0, 3): nums[0] XOR nums[3] = 6
//     - (1, 2): nums[1] XOR nums[2] = 6
//     - (1, 3): nums[1] XOR nums[3] = 3
//     - (2, 3): nums[2] XOR nums[3] = 5

// Example 2:
// Input: nums = [9,8,4,2,1], low = 5, high = 14
// Output: 8
// Explanation: All nice pairs (i, j) are as follows:
// ​​​​​    - (0, 2): nums[0] XOR nums[2] = 13
//     - (0, 3): nums[0] XOR nums[3] = 11
//     - (0, 4): nums[0] XOR nums[4] = 8
//     - (1, 2): nums[1] XOR nums[2] = 12
//     - (1, 3): nums[1] XOR nums[3] = 10
//     - (1, 4): nums[1] XOR nums[4] = 9
//     - (2, 3): nums[2] XOR nums[3] = 6
//     - (2, 4): nums[2] XOR nums[4] = 5

// Constraints:
//     1 <= nums.length <= 2 * 10^4
//     1 <= nums[i] <= 2 * 10^4
//     1 <= low <= high <= 2 * 10^4

import "fmt"

type TrieNode struct {
    count int
    children [2]*TrieNode
}

func (this *TrieNode) Insert(num int)  {
    for i := 14; i >= 0; i-- {
        bit := (num >> i) & 1
        if this.children[bit] == nil {
            this.children[bit] = &TrieNode{}
        }
        this = this.children[bit]
        this.count++
    }
}

func countPairs(nums []int, low int, high int) int {
    root, count := &TrieNode{}, 0
    countLessThan := func(root *TrieNode, num, limit int) int {
        node, count := root, 0
        for i := 14; i >= 0; i-- {
            if node == nil { break }
            numBit, limitBit := (num >> i) & 1, (limit >> i) & 1
            if limitBit == 1 {
                if node.children[numBit] != nil {
                    count += node.children[numBit].count
                }
                node = node.children[numBit ^ 1]
            } else {
                node = node.children[numBit]
            }
        }
        return count
    }
    for _, v := range nums {
        count += countLessThan(root, v, high + 1) - countLessThan(root, v, low)
        root.Insert(v)
    }
    return count
}

func main() {
    // Example 1:
    // Input: nums = [1,4,2,7], low = 2, high = 6
    // Output: 6
    // Explanation: All nice pairs (i, j) are as follows:
    //     - (0, 1): nums[0] XOR nums[1] = 5 
    //     - (0, 2): nums[0] XOR nums[2] = 3
    //     - (0, 3): nums[0] XOR nums[3] = 6
    //     - (1, 2): nums[1] XOR nums[2] = 6
    //     - (1, 3): nums[1] XOR nums[3] = 3
    //     - (2, 3): nums[2] XOR nums[3] = 5
    fmt.Println(countPairs([]int{1,4,2,7}, 2, 6)) // 6
    // Example 2:
    // Input: nums = [9,8,4,2,1], low = 5, high = 14
    // Output: 8
    // Explanation: All nice pairs (i, j) are as follows:
    // ​​​​​    - (0, 2): nums[0] XOR nums[2] = 13
    //     - (0, 3): nums[0] XOR nums[3] = 11
    //     - (0, 4): nums[0] XOR nums[4] = 8
    //     - (1, 2): nums[1] XOR nums[2] = 12
    //     - (1, 3): nums[1] XOR nums[3] = 10
    //     - (1, 4): nums[1] XOR nums[4] = 9
    //     - (2, 3): nums[2] XOR nums[3] = 6
    //     - (2, 4): nums[2] XOR nums[4] = 5
    fmt.Println(countPairs([]int{9,8,4,2,1}, 5, 14)) // 8
}