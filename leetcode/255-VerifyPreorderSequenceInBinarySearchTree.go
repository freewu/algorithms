package main 

// 255. Verify Preorder Sequence in Binary Search Tree
// Given an array of unique integers preorder, 
// return true if it is the correct preorder traversal sequence of a binary search tree.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/03/12/preorder-tree.jpg" />
// Input: preorder = [5,2,1,3,6]
// Output: true

// Example 2:
// Input: preorder = [5,2,6,1,3]
// Output: false

// Constraints:
//     1 <= preorder.length <= 10^4
//     1 <= preorder[i] <= 10^4
//     All the elements of preorder are unique.
 
// Follow up: Could you do it using only constant space complexity?

import "fmt"

// 递归
func verifyPreorder(preorder []int) bool {
    if len(preorder) <= 1 {
        return true
    }
    root, i := preorder[0], 1
    for ; i < len(preorder) && preorder[i] < root; i++ {} // 遍历数组，找到大于 preorder[0] 的第一个下标 i。那么下标 1 到 i-1 就是左子树
    for j := i + 1; j < len(preorder); j++ { // 然后继续向后遍历，如果有小于 preorder[0] 的数字，说明不成立
        if preorder[j] < root {
            return false
        }
    }
    return verifyPreorder(preorder[1:i]) && verifyPreorder(preorder[i:]) // 使用递归继续处理左右子树即可
}

// stack 单调栈
func verifyPreorder1(preorder []int) bool {
    stack, prev := make([]int, 0, len(preorder)), -1 << 32 - 1
    stack = append(stack, preorder[0])
    for i := 1; i < len(preorder); i++ {
        if preorder[i] < prev {
            return false
        }
        for len(stack) > 0 && preorder[i] > stack[len(stack)-1] {
            prev = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, preorder[i])
    }
    return true
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/03/12/preorder-tree.jpg" />
    // Input: preorder = [5,2,1,3,6]
    // Output: true
    fmt.Println(verifyPreorder([]int{5,2,1,3,6})) // true
    // Example 2:
    // Input: preorder = [5,2,6,1,3]
    // Output: false
    fmt.Println(verifyPreorder([]int{5,2,6,1,3})) // false

    fmt.Println(verifyPreorder1([]int{5,2,1,3,6})) // true
    fmt.Println(verifyPreorder1([]int{5,2,6,1,3})) // false
}