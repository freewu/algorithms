package main

// 1707. Maximum XOR With an Element From Array
// You are given an array nums consisting of non-negative integers. 
// You are also given a queries array, where queries[i] = [xi, mi].

// The answer to the ith query is the maximum bitwise XOR value of xi and any element of nums that does not exceed mi. 
// In other words, the answer is max(nums[j] XOR xi) for all j such that nums[j] <= mi. 
// If all elements in nums are larger than mi, then the answer is -1.

// Return an integer array answer where answer.length == queries.length 
// and answer[i] is the answer to the ith query.

// Example 1:
// Input: nums = [0,1,2,3,4], queries = [[3,1],[1,3],[5,6]]
// Output: [3,3,7]
// Explanation:
// 1) 0 and 1 are the only two integers not greater than 1. 0 XOR 3 = 3 and 1 XOR 3 = 2. The larger of the two is 3.
// 2) 1 XOR 2 = 3.
// 3) 5 XOR 2 = 7.

// Example 2:
// Input: nums = [5,2,4,6,6,3], queries = [[12,4],[8,1],[6,3]]
// Output: [15,-1,5]

// Constraints:
//     1 <= nums.length, queries.length <= 10^5
//     queries[i].length == 2
//     0 <= nums[j], xi, mi <= 10^9

import "fmt"

func maximizeXor(nums []int, queries [][]int) []int {
    type TrieNode struct {
        child    [2]*TrieNode // child[0]: bit 0 trie,  child[1]: bit 1 trie
        minValue int // minValue under this trie
    }
    insert := func(x int, root *TrieNode) {
        t := root
        for i := 29; i >= 0; i-- {
            k := ((x >> i) & 1)
            if t.child[k] == nil {
                t.child[k] = &TrieNode{}
                t.child[k].minValue = x
            }
            t = t.child[k]
            if t.minValue > x {
                t.minValue = x
            }
        }
    }
    find := func(x int, m int, root *TrieNode) int {
        res, t := 0, root
        for i := 29; i >= 0; i-- {
            k := ((x >> i) & 1)
            if t.child[k^1] != nil && t.child[k^1].minValue <= m {
                res += (1 << i)
                t = t.child[k^1]
            } else if t.child[k] != nil && t.child[k].minValue <= m {
                t = t.child[k]
            } else {
                return -1
            }
        }
        return res
    }
    root := &TrieNode{}
    for i := 0; i < len(nums); i++ { // build trie
        insert(nums[i], root)
    }
    res := make([]int, len(queries))
    for i := 0; i < len(queries); i++ {
        res[i] = find(queries[i][0], queries[i][1], root)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [0,1,2,3,4], queries = [[3,1],[1,3],[5,6]]
    // Output: [3,3,7]
    // Explanation:
    // 1) 0 and 1 are the only two integers not greater than 1. 0 XOR 3 = 3 and 1 XOR 3 = 2. The larger of the two is 3.
    // 2) 1 XOR 2 = 3.
    // 3) 5 XOR 2 = 7.
    fmt.Println(maximizeXor([]int{0,1,2,3,4}, [][]int{{3,1},{1,3},{5,6}})) // [3,3,7]
    // Example 2:
    // Input: nums = [5,2,4,6,6,3], queries = [[12,4],[8,1],[6,3]]
    // Output: [15,-1,5]
    fmt.Println(maximizeXor([]int{5,2,4,6,6,3}, [][]int{{12,4},{8,1},{6,3}})) // [15,-1,5]
}