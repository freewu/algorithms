package main 

// 366. Find Leaves of Binary Tree
// Given the root of a binary tree, collect a tree's nodes as if you were doing this:
//     Collect all the leaf nodes.
//     Remove all the leaf nodes.
//     Repeat until the tree is empty.

// Example 1:
//         1                 1           (1)
//       /   \              /
//      2    (3)    =>    (2)      => 
//     /  \
//   (4)  (5)
//    [4, 5, 3]             [2]           [1]
// <img src="https://assets.leetcode.com/uploads/2021/03/16/remleaves-tree.jpg" />
// Input: root = [1,2,3,4,5]
// Output: [[4,5,3],[2],[1]]
// Explanation:
// [[3,5,4],[2],[1]] and [[3,4,5],[2],[1]] are also considered correct answers since per each level it does not matter the order on which elements are returned.

// Example 2:
// Input: root = [1]
// Output: [[1]]

// Constraints:
//     The number of nodes in the tree is in the range [1, 100].
//     -100 <= Node.val <= 100


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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findLeaves(root *TreeNode) [][]int {
    // 最先被删除叶子的节点也是下一轮最先被删除的叶子候选者（为什么说是候选，因为一个节点可以有一个叶子孩子和一个非叶子孩子，那么它的叶子孩子被删除后自己并不会立即变成叶子。但是只有叶子孩子的节点肯定是下一批要删除的。
    // 而叶子的删除顺序其实就是先序/中序遍历时遇到叶子的顺序。
    // 所以可以维护一个子节点指向父节点的映射，这样，先遍历一遍即可得到当前要删除的叶子，并且可以在删除后立即修改并检查父亲节点的状态，从而确定下一批要删除的叶子。
    res := [][]int{}
    if root == nil { return res }
    type Ref struct { // 为方便得知当前叶节点是其父亲的哪个孩子，这里定义一个结构体
        parent *TreeNode // 父节点
        leftOrRight bool // 是父节点的左子还是右子
    }
    isLeaves := func(node *TreeNode) bool { // 判断是否是叶子节点
        return node.Left == nil && node.Right == nil 
    }
    leaves,toParent := []*TreeNode{}, map[*TreeNode]Ref{} // 维护一个指向父亲的 map
    var traversal func(node *TreeNode)
    traversal = func(node *TreeNode) { 
        if isLeaves(node) { // 是叶子节点,把节点加入到数组,跳出递归
            leaves = append(leaves, node)
            return 
        }
        if node.Left != nil {
            toParent[node.Left] = Ref{ parent: node, leftOrRight: true}
            traversal(node.Left)
        }
        if node.Right != nil {
            toParent[node.Right] = Ref{ parent: node, leftOrRight: false}
            traversal(node.Right)
        }
    }
    traversal(root) // 1. 遍历，构造toParent映射和第一批叶子
    for len(leaves) > 0 { // 删除leaves直到为空
        nextLeaves := []*TreeNode{} // 下一批叶子
        curVals := make([]int, len(leaves)) //本批次叶子值
        for i, leaf := range leaves {
            curVals[i] = leaf.Val
            ref, ok := toParent[leaf]
            if !ok { // 找不到父节点说明是根了，处理完成
                break
            }
            if ref.leftOrRight { // 删除的是左子叶
                ref.parent.Left = nil
            } else {
                ref.parent.Right = nil
            }
            if isLeaves(ref.parent) {// 父节点变成叶子了
                nextLeaves = append(nextLeaves, ref.parent)
            }
        }
        res = append(res, curVals)
        leaves = nextLeaves
    }
    return res
}

func findLeaves1(root *TreeNode) [][]int {
    res := [][]int{}
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(cur *TreeNode) int
    dfs = func(cur *TreeNode) int {
        if cur == nil {
            return 0
        }
        left, right := dfs(cur.Left), dfs(cur.Right)
        level := max(left, right) + 1
        if level - 1 == len(res) {
            res = append(res, []int{ cur.Val })
        } else {
            res[level-1] = append(res[level-1], cur.Val)
        }
        return level
    }
    dfs(root)
    return res
}

func main() {
    // Example 1:
    //         1                 1           (1)
    //       /   \              /
    //      2    (3)    =>    (2)      => 
    //     /  \
    //   (4)  (5)
    //    [4, 5, 3]             [2]           [1]
    // <img src="https://assets.leetcode.com/uploads/2021/03/16/remleaves-tree.jpg" />
    // Input: root = [1,2,3,4,5]
    // Output: [[4,5,3],[2],[1]]
    // Explanation:
    // [[3,5,4],[2],[1]] and [[3,4,5],[2],[1]] are also considered correct answers since per each level it does not matter the order on which elements are returned.
    tree1 := &TreeNode {
        1,
        &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}, },
        &TreeNode{3, nil, nil },
    }
    fmt.Println(findLeaves(tree1)) // [[4,5,3],[2],[1]]
    // Example 2:
    // Input: root = [1]
    // Output: [[1]]
    tree2 := &TreeNode{1, nil, nil}
    fmt.Println(findLeaves(tree2)) // [[1]]

    tree11 := &TreeNode {
        1,
        &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}, },
        &TreeNode{3, nil, nil },
    }
    fmt.Println(findLeaves(tree11)) // [[4,5,3],[2],[1]]
    tree12 := &TreeNode{1, nil, nil}
    fmt.Println(findLeaves(tree12)) // [[1]]
}