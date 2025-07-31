package main

// 3632. Subarrays with XOR at Least K
// Given an array of positive integers nums of length n and a non‑negative integer k.

// Return the number of contiguous subarrays whose bitwise XOR of all elements is greater than or equal to k.

// Example 1:
// Input: nums = [3,1,2,3], k = 2
// Output: 6
// Explanation:
// The valid subarrays with XOR >= 2 are [3] at index 0, [3, 1] at indices 0 - 1, [3, 1, 2, 3] at indices 0 - 3, [1, 2] at indices 1 - 2, [2] at index 2, and [3] at index 3; there are 6 in total.

// Example 2:
// Input: nums = [0,0,0], k = 0
// Output: 6
// Explanation:
// Every contiguous subarray yields XOR = 0, which meets k = 0. There are 6 such subarrays in total.

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^9
//     0 <= k <= 10^9

import "fmt"

// 字典树节点
type TrieNode struct {
    children [2]*TrieNode
    count    int // 该节点的计数
}

// 字典树
type Trie struct {
    root *TrieNode
}

func NewTrie() *Trie {
    return &Trie{ root: &TrieNode{} }
}

// 插入一个数到字典树中
func (t *Trie) Insert(num int) {
    node := t.root
    // 从最高位开始插入（假设数字不超过2^30）
    for i := 30; i >= 0; i-- {
        bit := (num >> i) & 1
        if node.children[bit] == nil {
            node.children[bit] = &TrieNode{}
        }
        node = node.children[bit]
        node.count++
    }
}

// 查询与num异或结果 >= k的数量
func (t *Trie) Query(num, k int) int {
    res, node := 0, t.root
    if node == nil { return 0 }
    for bit := 30; bit >= 0; bit-- {
        if node == nil { break }
        numBit, kBit := (num >> bit) & 1, (k >> bit) & 1
        if kBit == 0 {
            // 如果k的当前位是0，那么异或结果的当前位是1的数都满足条件
            if node.children[1 - numBit] != nil {
                res += node.children[1 - numBit].count
            }
            // 继续检查异或结果当前位是0的数
            node = node.children[numBit]
        } else {
            // 如果k的当前位是1，那么只有异或结果的当前位是1的数才可能满足条件
            node = node.children[1 - numBit]
        }
    }
    // 如果遍历完所有位后node不为空，说明这些数与num的异或结果等于k
    if node != nil {
        res += node.count
    }
    return res
}

func countXorSubarrays(nums []int, k int) int64 {
    res, n := 0, len(nums)
    if n == 0 { return 0 }
    // 计算前缀异或数组
    // prefixXor[i] 表示 nums[0] 到 nums[i-1] 的异或结果
    prefixXor := make([]int, n+1)
    for i := 0; i < n; i++ {
        prefixXor[i+1] = prefixXor[i] ^ nums[i]
    }
    // 使用字典树存储前缀异或值
    trie := NewTrie()
    trie.Insert(0) // 插入prefixXor[0] = 0
    for j := 1; j <= n; j++ {
        // 查找与prefixXor[j]异或结果 >= k的数量
        res += trie.Query(prefixXor[j], k)
        // 插入当前前缀异或值
        trie.Insert(prefixXor[j])
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [3,1,2,3], k = 2
    // Output: 6
    // Explanation:
    // The valid subarrays with XOR >= 2 are [3] at index 0, [3, 1] at indices 0 - 1, [3, 1, 2, 3] at indices 0 - 3, [1, 2] at indices 1 - 2, [2] at index 2, and [3] at index 3; there are 6 in total.
    fmt.Println(countXorSubarrays([]int{3,1,2,3}, 2)) // 6
    // Example 2:
    // Input: nums = [0,0,0], k = 0
    // Output: 6
    // Explanation:
    // Every contiguous subarray yields XOR = 0, which meets k = 0. There are 6 such subarrays in total.
    fmt.Println(countXorSubarrays([]int{0,0,0}, 0)) // 6

    fmt.Println(countXorSubarrays([]int{1,2,3,4,5,6,7,8,9}, 0)) // 45
    fmt.Println(countXorSubarrays([]int{9,8,7,6,5,4,3,2,1}, 0)) // 45
}