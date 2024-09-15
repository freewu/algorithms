package main

// 1104. Path In Zigzag Labelled Binary Tree
// In an infinite binary tree where every node has two children, the nodes are labelled in row order.

// In the odd numbered rows (ie., the first, third, fifth,...), the labelling is left to right, 
// while in the even numbered rows (second, fourth, sixth,...), the labelling is right to left.
// <img src="https://assets.leetcode.com/uploads/2019/06/24/tree.png" />

// Given the label of a node in this tree, 
// return the labels in the path from the root of the tree to the node with that label.

// Example 1:
// Input: label = 14
// Output: [1,3,4,14]

// Example 2:
// Input: label = 26
// Output: [1,2,6,10,26]

// Constraints:
//     1 <= label <= 10^6

import "fmt"
import "math"

func pathInZigZagTree(label int) []int {
    res, arr := []int{}, []int{}
    index := int(math.Log2(float64(label))) + 1
    arr = append(arr, label)
    tmp := label
    for tmp > 1 {
        tmp /= 2
        index--
        tmp = int(math.Pow(2, float64(index)) - 1 + math.Pow(2, float64(index-1))) - tmp
        arr = append(arr, tmp)
    }
    for i := len(arr) - 1; i >= 0; i-- {
        res = append(res, arr[i])
    }
    return res
}

func main() {
    // Example 1:
    // Input: label = 14
    // Output: [1,3,4,14]
    fmt.Println(pathInZigZagTree(14)) // [1,3,4,14]
    // Example 2:
    // Input: label = 26
    // Output: [1,2,6,10,26]
    fmt.Println(pathInZigZagTree(26)) // [1,2,6,10,26]
}