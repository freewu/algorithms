package main 

// LCR 056. 两数之和 IV - 输入二叉搜索树
// 给定一个二叉搜索树的 根节点 root 和一个整数 k , 请判断该二叉搜索树中是否存在两个节点它们的值之和等于 k 。
// 假设二叉搜索树中节点的值均唯一。

// 示例 1：
// 输入: root = [8,6,10,5,7,9,11], k = 12
// 输出: true
// 解释: 节点 5 和节点 7 之和等于 12

// 示例 2：
// 输入: root = [8,6,10,5,7,9,11], k = 22
// 输出: false
// 解释: 不存在两个节点值之和为 22 的节点
 
// 提示：
//     二叉树的节点个数的范围是  [1, 10^4].
//     -10^4 <= Node.val <= 10^4
//     root 为二叉搜索树
//     -10^5 <= k <= 10^5
 
import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    mp := make(map[int]int)
    var dfs func(*TreeNode) bool 
    dfs = func(node *TreeNode) bool {
        if node == nil {
            return false
        }
        if _, ok := mp[node.Val]; ok { // 如果发现存在返回 true
            return true
        }
        mp[k - node.Val]++ // 遍历树 把  k - node.Val 存在 map 中
        return dfs(node.Left) || dfs(node.Right)
    }
    return dfs(root)
}

func findTarget1(root *TreeNode, k int) bool {
    arr := make([]int, 0)
    var inorder func(root *TreeNode)
    inorder = func(root *TreeNode) {
        if root == nil {
            return
        }
        inorder(root.Left)
        arr = append(arr, root.Val)
        inorder(root.Right)
    }
    inorder(root) // // 中序遍历得到有序数组
    left, right := 0, len(arr)-1
    for left < right {
        if arr[left] + arr[right] < k {
            left++
        } else if arr[left] + arr[right] > k {
            right--
        } else {
            return true
        }
    }
    return false
}

func main() {
    // Example 1:
    //            {5}
    //           /   \
    //         (3)   (6)
    //        /   \     \
    //      [2]    {4}   [7]
    // <img src="https://assets.leetcode.com/uploads/2020/09/21/sum_tree_1.jpg" />
    // Input: root = [5,3,6,2,4,null,7], k = 9
    // Output: true
    tree1 := &TreeNode {
        5,
        &TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}, },
        &TreeNode{6, nil,                    &TreeNode{7, nil, nil}, },
    }
    fmt.Println(findTarget(tree1, 9)) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/09/21/sum_tree_2.jpg" />
    // Input: root = [5,3,6,2,4,null,7], k = 28
    // Output: false
    fmt.Println(findTarget(tree1, 28)) // false

    fmt.Println(findTarget1(tree1, 9)) // true
    fmt.Println(findTarget1(tree1, 28)) // false
}