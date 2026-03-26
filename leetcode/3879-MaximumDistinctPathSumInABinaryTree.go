package main

// 3879. Maximum Distinct Path Sum in a Binary Tree
// You are given the root of a binary tree, where each node contains an integer value.

// A valid path in the tree is a sequence of connected nodes such that:
//     1, The path can start and end at any node in the tree.
//     2. The path does not need to pass through the root.
//     3. All node values along the path are distinct.

// Return an integer denoting the maximum possible sum of node values among all valid paths.

// Example 1:
//         2
//      /    \
//     2      1
// ​​​​​​<img src="https://assets.leetcode.com/uploads/2026/01/28/screenshot-2026-01-29-at-12940am.png" />
// Input: root = [2,2,1]
// Output: 3
// Explanation:
// The path 2 → 2 is invalid because the value 2 is not distinct.
// The maximum-sum valid path is 2 → 1, with a sum = 2 + 1 = 3.

// Example 2:
//         1
//        /    \
//       -2      5
//             /   \
//            3      5
// ​​​​​​<img src="https://assets.leetcode.com/uploads/2026/01/28/screenshot-2026-01-29-at-15149am.png" />
// Input: root = [1,-2,5,null,null,3,5]
// Output: 9
// Explanation:
// The path 3 → 5 → 5 is invalid due to duplicate value 5.
// The maximum-sum valid path is 1 → 5 → 3, with a sum = 1 + 5 + 3 = 9.

// Example 3:
//         4
//        /    \
//       6      6
//                \
//                 9
// ​​​​​​<img src="https://assets.leetcode.com/uploads/2026/01/28/screenshot-2026-01-29-at-15555am.png" />
// Input: root = [4,6,6,null,null,null,9]
// Output: 19
// Explanation:
// The path 6 → 4 → 6 → 9 is invalid because the value 6 appears more than once.
// The maximum-sum valid path is 4 → 6 → 9, with a sum = 4 + 6 + 9 = 19.

// Constraints:
//     The number of nodes in the tree is in the range [1, 1000].
//     -1000 <= Node.val <= 1000​​​​​​​

import "fmt"
import "math"
import "bytes"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// buildBinaryTree 根据数组构建 完全二叉树
// 规则：数组下标 i 的左孩子 2i+1，右孩子 2i+2，-1 代表空节点
func buildBinaryTree(arr []int) *TreeNode {
	if len(arr) == 0 || arr[0] == -1 {
		return nil
	}
	root := &TreeNode{Val: arr[0]}
	queue := []*TreeNode{root}
	i := 1
	// 队列层序构建二叉树
	for len(queue) > 0 && i < len(arr) {
		node := queue[0]
		queue = queue[1:]
		// 构建左子树
		if i < len(arr) && arr[i] != math.MinInt32 {
			node.Left = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Left)
		}
		i++
		// 构建右子树
		if i < len(arr) && arr[i] != math.MinInt32 {
			node.Right = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// printBinaryTree 美观打印二叉树：根居中、/ 左、\ 右
func printBinaryTree(root *TreeNode) {
    if root == nil {
        fmt.Println("二叉树为空")
        return
    }
    // getTreeHeight 递归获取树高度
    var getTreeHeight func(root *TreeNode) int
    getTreeHeight =func (root *TreeNode) int {
        if root == nil { return 0 }
        leftHeight := getTreeHeight(root.Left)
        rightHeight := getTreeHeight(root.Right)
        if leftHeight > rightHeight {
            return leftHeight + 1
        }
        return rightHeight + 1
    }
	// 1. 计算树的高度
	height := getTreeHeight(root)
	// 2. 最后一层宽度 = 2^(h-1)
	lastLevelWidth := int(math.Pow(2, float64(height-1)))
	// 3. 总打印宽度（每个节点占 4 字符）
	totalWidth := lastLevelWidth * 4
	// 队列层序遍历，存储节点 + 当前层级
	queue := []struct {
		node  *TreeNode
		level int
	}{{root, 1}}
	for len(queue) > 0 {
		levelSize := len(queue)
		// 存储当前行节点、连接线
		var nodeBuf bytes.Buffer
		var linkBuf bytes.Buffer
		for i := 0; i < levelSize; i++ {
			curr := queue[0]
			queue = queue[1:]
			// 当前层级的间距
			gap := totalWidth / int(math.Pow(2, float64(curr.level)))
			// 打印节点
			if curr.node != nil {
				fmt.Fprintf(&nodeBuf, "%*s%d%*s", gap, "", curr.node.Val, gap, "")
				// 打印连接线 / \
				if curr.level < height {
					fmt.Fprintf(&linkBuf, "%*s/%*s\\%*s", gap-1, "", gap*2-2, "", gap-1, "")
				}
				// 入队子节点
				queue = append(queue, struct {
					node  *TreeNode
					level int
				}{curr.node.Left, curr.level + 1})
				queue = append(queue, struct {
					node  *TreeNode
					level int
				}{curr.node.Right, curr.level + 1})
			} else {
				// 空节点占位
				fmt.Fprintf(&nodeBuf, "%*s%*s", gap, "", gap, "")
				if curr.level < height {
					fmt.Fprintf(&linkBuf, "%*s%*s%*s", gap, "", gap*2, "", gap, "")
				}
				// 空节点也入队占位
				queue = append(queue, struct {
					node  *TreeNode
					level int
				}{nil, curr.level + 1})
				queue = append(queue, struct {
					node  *TreeNode
					level int
				}{nil, curr.level + 1})
			}
		}
		// 输出当前行
		fmt.Println(nodeBuf.String())
		// 输出连接线（最后一层无连接线）
		if linkBuf.Len() > 0 {
			fmt.Println(linkBuf.String())
		}
	}
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// Time Limit Exceeded 985 / 1000 testcases passed
func maxSum(root *TreeNode) int {
    res := -1 << 61
    type Node struct { // 给节点加上父指针
        Val    int
        Left   *Node
        Right  *Node
        Parent *Node
    }
    var build func(*TreeNode) *Node
    build = func(tn *TreeNode) *Node { // 重建带父指针的树
        if tn == nil { return nil }
        n := &Node{Val: tn.Val}
        n.Left = build(tn.Left)
        if n.Left != nil {
            n.Left.Parent = n
        }
        n.Right = build(tn.Right)
        if n.Right != nil {
            n.Right.Parent = n
        }
        return n
    }
    newRoot := build(root)
    // 从每个点出发，DFS 三个方向，值不能重复
    var dfs func(*Node)
    dfs = func(node *Node) {
        if node == nil { return }
        var back func(cur *Node, valMap map[int]bool, sum int)
        back = func(cur *Node, valMap map[int]bool, sum int) {
            if cur == nil || valMap[cur.Val] { return } // 已访问过跳出
            // 复制 map，保证每条分支独立
            mp := make(map[int]bool)
            for k := range valMap {
                mp[k] = true
            }
            mp[cur.Val] = true
            newSum := sum + cur.Val
            if newSum > res {
                res = newSum
            }
            back(cur.Left, mp, newSum)
            back(cur.Right, mp, newSum)
            back(cur.Parent, mp, newSum)
        }
        back(node, make(map[int]bool), 0)
        dfs(node.Left)
        dfs(node.Right)
    }
    dfs(newRoot)
    return res
}

func maxSum1(root *TreeNode) int {
    res := -1 << 61
    type Node struct {
        Val    int
        Left   *Node
        Right  *Node
        Parent *Node
    }
    var build func(*TreeNode) *Node  // 构建带父指针的树，只执行1次 O(n)
    build = func(tn *TreeNode) *Node { 
        if tn == nil { return nil }
        n := &Node{ Val: tn.Val }
        n.Left = build(tn.Left)
        if n.Left != nil {
            n.Left.Parent = n
        }
        n.Right = build(tn.Right)
        if n.Right != nil {
            n.Right.Parent = n
        }
        return n
    }
    newRoot := build(root)
    // 核心优化：全局复用 map + 回溯（无复制！）
    var traverse func(*Node)
    traverse = func(node *Node) {
        if node == nil { return }
        visited := make(map[int]bool, 100) // 预分配容量
        var dfs func(cur *Node, sum int)
        dfs = func(cur *Node, sum int) {
            // 终止条件：空节点 / 值重复
            if cur == nil || visited[cur.Val] { return }
            // 标记
            visited[cur.Val] = true
            sum += cur.Val
            if sum > res {
                res = sum
            }
            // 三个方向搜索
            dfs(cur.Left, sum)
            dfs(cur.Right, sum)
            dfs(cur.Parent, sum)
            // 回溯：撤销标记（关键！避免map复制，性能爆炸提升）
            visited[cur.Val] = false
        }
        dfs(node, 0)
        traverse(node.Left)
        traverse(node.Right)
    }
    traverse(newRoot)
    return res
}

func main() {
    // Example 1:
    //         2
    //      /    \
    //     2      1
    // ​​​​​​<img src="https://assets.leetcode.com/uploads/2026/01/28/screenshot-2026-01-29-at-12940am.png" />
    // Input: root = [2,2,1]
    // Output: 3
    // Explanation:
    // The path 2 → 2 is invalid because the value 2 is not distinct.
    // The maximum-sum valid path is 2 → 1, with a sum = 2 + 1 = 3.
    tree1 := &TreeNode{
        Val: 2, 
        Left:  &TreeNode{ Val: 2}, 
        Right: &TreeNode{ Val: 1},
    }
    fmt.Println(maxSum(tree1)) // 3
    // Example 2:
    //         1
    //        /    \
    //       -2      5
    //             /   \
    //            3      5
    // ​​​​​​<img src="https://assets.leetcode.com/uploads/2026/01/28/screenshot-2026-01-29-at-15149am.png" />
    // Input: root = [1,-2,5,null,null,3,5]
    // Output: 9
    // Explanation:
    // The path 3 → 5 → 5 is invalid due to duplicate value 5.
    // The maximum-sum valid path is 1 → 5 → 3, with a sum = 1 + 5 + 3 = 9.
    tree2 := &TreeNode{
        Val: 1, 
        Left: &TreeNode{ Val: -2}, 
        Right: &TreeNode{
                Val: 5, 
                Left:  &TreeNode{ Val: 3}, 
                Right: &TreeNode{Val: 5},
        },
    }
    fmt.Println(maxSum(tree2)) // 9
    // Example 3:
    //         4
    //        /    \
    //       6      6
    //                \
    //                 9
    // ​​​​​​<img src="https://assets.leetcode.com/uploads/2026/01/28/screenshot-2026-01-29-at-15555am.png" />
    // Input: root = [4,6,6,null,null,null,9]
    // Output: 19
    // Explanation:
    // The path 6 → 4 → 6 → 9 is invalid because the value 6 appears more than once.
    // The maximum-sum valid path is 4 → 6 → 9, with a sum = 4 + 6 + 9 = 19. 
    tree3 := &TreeNode{
        Val: 4, 
        Left:  &TreeNode{ Val: 6}, 
        Right: &TreeNode{ 
            Val: 6, 
            Right: &TreeNode{ Val: 9},
        },
    }
    // printBinaryTree(tree3)
    fmt.Println(maxSum(tree3)) // 19
    // Example 4:
    // Input: root = [900,245,794,627,853,858,-581,null,null,null,165,null,-392,null,72]
    // Output: 3815
    tree4 := buildBinaryTree([]int{900,245,794,627,853,858,-581,math.MinInt32 ,math.MinInt32 ,math.MinInt32 ,165,math.MinInt32 ,-392,math.MinInt32 ,72})
    fmt.Println(maxSum(tree4)) // 3815
    //printBinaryTree(tree4)

    fmt.Println(maxSum1(tree1)) // 3
    fmt.Println(maxSum1(tree2)) // 9
    fmt.Println(maxSum1(tree3)) // 19
    fmt.Println(maxSum1(tree4)) // 3815
}

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

// 使用 golang 实现，传入 数组，返回通过数组创建的二叉树, 方法定义如下:
// func buildBinaryTree(arr []int) *TreeNode
// 使用 golang 实现，传入 二叉树， 层级显示打印显示二叉树，方法签名如下:
// func printBinaryTree(root *TreeNode)  
//     1 root 需要在顶部居中
//     2 左边使用 /
//     3 右边使用 \
   
