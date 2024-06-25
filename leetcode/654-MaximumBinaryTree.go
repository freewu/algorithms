package main

// 654. Maximum Binary Tree
// You are given an integer array nums with no duplicates. 
// A maximum binary tree can be built recursively from nums using the following algorithm:
//     Create a root node whose value is the maximum value in nums.
//     Recursively build the left subtree on the subarray prefix to the left of the maximum value.
//     Recursively build the right subtree on the subarray suffix to the right of the maximum value.

// Return the maximum binary tree built from nums.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/12/24/tree1.jpg" />
// Input: nums = [3,2,1,6,0,5]
// Output: [6,3,5,null,2,0,null,null,1]
// Explanation: The recursive calls are as follow:
// - The largest value in [3,2,1,6,0,5] is 6. Left prefix is [3,2,1] and right suffix is [0,5].
//     - The largest value in [3,2,1] is 3. Left prefix is [] and right suffix is [2,1].
//         - Empty array, so no child.
//         - The largest value in [2,1] is 2. Left prefix is [] and right suffix is [1].
//             - Empty array, so no child.
//             - Only one element, so child is a node with value 1.
//     - The largest value in [0,5] is 5. Left prefix is [0] and right suffix is [].
//         - Only one element, so child is a node with value 0.
//         - Empty array, so no child.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/12/24/tree2.jpg" />
// Input: nums = [3,2,1]
// Output: [3,null,2,null,1]

// Constraints:
//     1 <= nums.length <= 1000
//     0 <= nums[i] <= 1000
//     All integers in nums are unique.

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
    // 找到左边最大的数作为左子树的根节点
    // 找到最大的数作为根节点
    // 以此类推
    if len(nums) == 0 {
        return nil
    }
    var root TreeNode
    mx, index := nums[0], 0
    for i, v := range nums { // 找到最大的数作为根节点
        if mx < v {
            mx = v
            index = i
        }
    }
    root.Val = mx
    root.Left = constructMaximumBinaryTree(nums[:index]) // 找到左边最大的数作为左子树的根节点
    root.Right = constructMaximumBinaryTree(nums[index + 1:])
    return &root
}

func constructMaximumBinaryTree1(nums []int) *TreeNode {
    var build func(left, right int) *TreeNode
    build = func(left, right int) *TreeNode {
        if right < left {
            return nil
        }
        mx, index := -1, 0
        for i := left; i <= right; i++ {
            if nums[i] > mx {
                mx = nums[i]
                index = i
            }
        }
        root := &TreeNode{ Val: nums[index], }
        root.Left = build(left, index - 1)
        root.Right = build(index + 1, right)
        return root
    }
    return build(0, len(nums) - 1)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/12/24/tree1.jpg" />
    // Input: nums = [3,2,1,6,0,5]
    // Output: [6,3,5,null,2,0,null,null,1]
    // Explanation: The recursive calls are as follow:
    // - The largest value in [3,2,1,6,0,5] is 6. Left prefix is [3,2,1] and right suffix is [0,5].
    //     - The largest value in [3,2,1] is 3. Left prefix is [] and right suffix is [2,1].
    //         - Empty array, so no child.
    //         - The largest value in [2,1] is 2. Left prefix is [] and right suffix is [1].
    //             - Empty array, so no child.
    //             - Only one element, so child is a node with value 1.
    //     - The largest value in [0,5] is 5. Left prefix is [0] and right suffix is [].
    //         - Only one element, so child is a node with value 0.
    //         - Empty array, so no child.
    fmt.Println(constructMaximumBinaryTree([]int{3,2,1,6,0,5})) // &{6 0xc000008060 0xc0000080a8}
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/12/24/tree2.jpg" />
    // Input: nums = [3,2,1]
    // Output: [3,null,2,null,1]
    fmt.Println(constructMaximumBinaryTree([]int{3,2,1})) // &{3 <nil> 0xc000008108}

    fmt.Println(constructMaximumBinaryTree1([]int{3,2,1,6,0,5})) // &{6 0xc000008060 0xc0000080a8}
    fmt.Println(constructMaximumBinaryTree1([]int{3,2,1})) // &{3 <nil> 0xc000008108}
}