package main

// LCP 34. 二叉树染色
// 小扣有一个根结点为 root 的二叉树模型，初始所有结点均为白色，可以用蓝色染料给模型结点染色，模型的每个结点有一个 val 价值。
// 小扣出于美观考虑，希望最后二叉树上每个蓝色相连部分的结点个数不能超过 k 个，求所有染成蓝色的结点价值总和最大是多少？

// 示例 1：
// 输入：root = [5,2,3,4], k = 2
// 输出：12
// 解释：结点 5、3、4 染成蓝色，获得最大的价值 5+3+4=12
// <img src="https://pic.leetcode-cn.com/1616126267-BqaCRj-image.png" />

// 示例 2：
// 输入：root = [4,1,3,9,null,null,2], k = 2
// 输出：16
// 解释：结点 4、3、9 染成蓝色，获得最大的价值 4+3+9=16
// <img src="https://pic.leetcode-cn.com/1616126301-gJbhba-image.png" />

// 提示：
//     1 <= k <= 10
//     1 <= val <= 10000
//     1 <= 结点数量 <= 10000

import "fmt"
import "slices"

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
func maxValue(root *TreeNode, k int) int {
    // 每个节点可以有两种处理结果：染色或者不染色
    // 根据题意显然是要用后序遍历，自底向上，由下层给出结果，供上层使用和选择
    // 定义：dp[0] 表示：以某个节点为根节点，根节点不染色的情况
    // 定义：dp[i] 表示：以某个节点为根节点，（根节点染色的情况），蓝色相连节点为i个的最大染色值， 1 <= i <= k， 
    // 所以dp数组的长度应为（k+1）
    // 根据每个节点的染色情况：
    // 不染色 ： dp[0] = 左子节点处理的最大值+右子节点处理的最大值
    // 染色： dp[i] = left_dp[j] + right_dp[i-1-j] + root.Val     (i = [1,k] , j=[0,i）)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(root *TreeNode, k int) []int
    dfs = func(root *TreeNode, k int) []int {
        dp := make([]int, k + 1)
        if root == nil { return dp }
        left_dp, right_dp := dfs(root.Left, k), dfs(root.Right, k)
        dp[0] = slices.Max(left_dp) + slices.Max(right_dp)
        for i := 1; i <= k; i++ {
            for j := 0; j < i; j++ {
                // 一共有 i， 左子树分配 j个， 右子树分配 i - j- 1 个 （根结点一个）
                dp[i] = max(dp[i], (left_dp[j] + right_dp[i - j - 1] + root.Val))
            }
        }
        return dp
    }
    return slices.Max(dfs(root, k))
}

func main() {
    // 示例 1：
    // 输入：root = [5,2,3,4], k = 2
    // 输出：12
    // 解释：结点 5、3、4 染成蓝色，获得最大的价值 5+3+4=12
    // <img src="https://pic.leetcode-cn.com/1616126267-BqaCRj-image.png" />
    tree1 := &TreeNode {
        5,
        &TreeNode{2, &TreeNode{4, nil, nil}, nil, },
        &TreeNode{3, nil, nil, },
    }
    fmt.Println(maxValue(tree1, 2)) // 12
    // 示例 2：
    // 输入：root = [4,1,3,9,null,null,2], k = 2
    // 输出：16
    // 解释：结点 4、3、9 染成蓝色，获得最大的价值 4+3+9=16
    // <img src="https://pic.leetcode-cn.com/1616126301-gJbhba-image.png" />
    tree2 := &TreeNode {
        4,
        &TreeNode{1, &TreeNode{9, nil, nil}, nil, },
        &TreeNode{3, nil, &TreeNode{2, nil, nil}, },
    }
    fmt.Println(maxValue(tree2, 2)) // 16
}