package main

// LCR 152. 验证二叉搜索树的后序遍历序列
// 请实现一个函数来判断整数数组 postorder 是否为二叉搜索树的后序遍历结果。

// 示例 1：
// <img src="https://pic.leetcode.cn/1706665328-rfvWhs-%E6%88%AA%E5%B1%8F2024-01-31%2009.41.48.png" />
// 输入: postorder = [4,9,6,5,8]
// 输出: false 
// 解释：从上图可以看出这不是一颗二叉搜索树

// 示例 2：
// <img src="https://pic.leetcode.cn/1694762510-vVpTic-%E5%89%91%E6%8C%8733.png" />
// 输入: postorder = [4,6,5,9,8]
// 输出: true 
// 解释：可构建的二叉搜索树如上图

// 提示：
//     数组长度 <= 1000
//     postorder 中无重复数字

import "fmt"

func verifyTreeOrder(postorder []int) bool {
    if len(postorder) == 0 {
        return true
    }
    root := postorder[len(postorder)-1]
    right, left := make([]int, 0), make([]int, 0)
    for i := 0; i < len(postorder)-1; i++ {
        if postorder[i] < root {
            left = append(left, postorder[i])
        } else {
            right = postorder[i : len(postorder)-1]
            break
        }
    }
    for i := 0; i < len(right); i++ {
        if right[i] <= root {
            return false
        }
    }
    if len(left) + len(right) != len(postorder) - 1 {
        return false
    }
    return verifyTreeOrder(left) && verifyTreeOrder(right)
}

// 递归分治
func verifyTreeOrder1(postorder []int) bool {
    var helper func(postorder []int, i int, j int) bool 
    helper = func(postorder []int, i int, j int) bool {
        if i >= j { return true }
        p := i 
        for postorder[p] < postorder[j] { p++ }
        m := p 
        for postorder[p] > postorder[j] { p++ }
        return p == j && helper(postorder, i, m - 1) && helper(postorder, m, j - 1)
    }
    return helper(postorder,0, len(postorder) - 1)
}

// 辅助单调栈
func verifyTreeOrder2(postorder []int) bool {
    stack, root := []int{}, 1 << 32 - 1
    for i := len(postorder) - 1; i >= 0; i-- {
        if postorder[i] > root { 
            return false 
        }
        for len(stack) != 0 && stack[len(stack) - 1] > postorder[i] {
            root = stack[len(stack) - 1]
            stack = stack[0:len(stack) - 1] // pop
        }
        stack = append(stack, postorder[i]) // add
    }
    return true
}

func main() {
    // 示例 1：
    // <img src="https://pic.leetcode.cn/1706665328-rfvWhs-%E6%88%AA%E5%B1%8F2024-01-31%2009.41.48.png" />
    // 输入: postorder = [4,9,6,5,8]
    // 输出: false 
    // 解释：从上图可以看出这不是一颗二叉搜索树
    fmt.Println(verifyTreeOrder([]int{4,9,6,5,8})) // false
    // 示例 2：
    // <img src="https://pic.leetcode.cn/1694762510-vVpTic-%E5%89%91%E6%8C%8733.png" />
    // 输入: postorder = [4,6,5,9,8]
    // 输出: true 
    // 解释：可构建的二叉搜索树如上图
    fmt.Println(verifyTreeOrder([]int{4,6,5,9,8})) // true

    fmt.Println(verifyTreeOrder1([]int{4,9,6,5,8})) // false
    fmt.Println(verifyTreeOrder1([]int{4,6,5,9,8})) // true

    fmt.Println(verifyTreeOrder2([]int{4,9,6,5,8})) // false
    fmt.Println(verifyTreeOrder2([]int{4,6,5,9,8})) // true
}