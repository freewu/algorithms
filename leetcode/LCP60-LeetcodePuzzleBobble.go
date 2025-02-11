package main

// LCP 60. 力扣泡泡龙
// 欢迎各位勇者来到力扣城，本次试炼主题为「力扣泡泡龙」。

// 游戏初始状态的泡泡形如二叉树 root，每个节点值对应了该泡泡的分值。
// 勇者最多可以击破一个节点泡泡，要求满足：
//     1. 被击破的节点泡泡 至多 只有一个子节点泡泡
//     2. 当被击破的节点泡泡有子节点泡泡时，则子节点泡泡将取代被击破泡泡的位置

// 注：即整棵子树泡泡上移

// 请问在击破一个节点泡泡操作或无击破操作后，二叉泡泡树的最大「层和」是多少。

// 注意：
//     「层和」为同一高度的所有节点的分值之和

// 示例 1：
// 输入：root = [6,0,3,null,8]
// 输出：11
// 解释：勇者的最佳方案如图所示
// <img src="https://pic.leetcode-cn.com/1648180809-XSWPLu-image.png" />

// 示例 2：
// 输入：root = [5,6,2,4,null,null,1,3,5]
// 输出：9
// 解释：勇者击破 6 节点，此时「层和」最大为 3+5+1 = 9
// <img src="https://pic.leetcode-cn.com/1648180769-TLpYop-image.png" />

// 示例 3：
// 输入：root = [-5,1,7]
// 输出：8
// 解释：勇者不击破节点，「层和」最大为 1+7 = 8

// 提示：
//     2 <= 树中节点个数 <= 10^5
//     -10000 <= 树中节点的值 <= 10000

import "fmt"

type TreeNode struct {
    Val   int
    Left  *TreeNode
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
func getMaxLayerSum(root *TreeNode) int {
    type LevelInfo struct { PrefixSum, IsOccupied, LeftPtr, RightPtr int }   // 用于存储每层的信息，包括前缀和、是否被占用以及左、右子节点的指针
    type RemoveNode struct { NodeIndex, Level int } // 用于存储需要被移除的节点信息
    levelInfos := make([][]LevelInfo, 0) // 用于存储每一层的信息
    removeNodes := make([]RemoveNode, 0) // 用于存储需要被移除的节点信息
    // collectNodes 执行DFS，并收集树的每一层的信息和需要被移除的节点。
    var collectNodes func(level int, node *TreeNode) int
    collectNodes = func(level int, node *TreeNode) int {
        if node == nil { return 0 }
        // 扩展 levelInfos 列表以适应新层。
        for len(levelInfos) <= level+1 {
            levelInfos = append(levelInfos, []LevelInfo{{ PrefixSum: 0, IsOccupied: -1, LeftPtr: -1, RightPtr: -1 }})
        }
        // 更新当前层的信息。
        lastInfo := levelInfos[level][len(levelInfos[level]) - 1]
        newInfo := LevelInfo{
            PrefixSum:  lastInfo.PrefixSum + node.Val, // 计算前缀和
            IsOccupied: -1,                            // 初始化为未占用
            LeftPtr:    len(levelInfos[level + 1]),      // 设置左子节点的指针
            RightPtr:   -1,                            // 初始化右子节点的指针为-1
        }
        levelInfos[level] = append(levelInfos[level], newInfo)
        node.Val = len(levelInfos[level]) - 1 // 用于识别节点的索引
        // 如果节点有两个子节点，则添加到 removeNodes 列表中。
        if collectNodes(level+1, node.Left) + collectNodes(level+1, node.Right) != 2 {
            removeNodes = append(removeNodes, RemoveNode{NodeIndex: len(levelInfos[level]) - 1, Level: level})
        }
        // 更新右子节点的指针。
        levelInfos[level][len(levelInfos[level])-1].RightPtr = len(levelInfos[level+1]) - 1
        return 1
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    collectNodes(0, root)
    res, treeHeight := -1 << 31, len(levelInfos) - 1 // 获取树的高度
    for level := 0; level < treeHeight; level++ { // 计算没有移除任何节点时的最大层和
        res = max(res, levelInfos[level][len(levelInfos[level]) - 1].PrefixSum)
    }
    for _, removeInfo := range removeNodes { // 遍历 removeNodes 列表，并尝试移除每一个节点以计算新的最大层和
        nodeIndex, startLevel := removeInfo.NodeIndex, removeInfo.Level
        left, right := nodeIndex, nodeIndex
        // 计算被移除节点的值。
        lostVal := levelInfos[startLevel][left].PrefixSum - levelInfos[startLevel][left - 1].PrefixSum
        for level := startLevel; level < treeHeight; level++ {
            if left > right { break }
            leftInfo, rightInfo := &levelInfos[level][left], &levelInfos[level][right]
            if leftInfo.IsOccupied != -1 && leftInfo.IsOccupied == rightInfo.IsOccupied { break } // 如果节点已被标记为占用，则跳过
            leftInfo.IsOccupied, rightInfo.IsOccupied = nodeIndex, nodeIndex // 标记节点为占用
            addVal := 0
            if leftInfo.LeftPtr <= rightInfo.RightPtr { // 计算被移除节点的子节点和。
                addVal = levelInfos[level+1][rightInfo.RightPtr].PrefixSum - levelInfos[level+1][leftInfo.LeftPtr-1].PrefixSum
            }
            newSum := levelInfos[level][len(levelInfos[level])-1].PrefixSum - lostVal + addVal // 计算新的层和
            if newSum != 0 || (right-left + 1 != len(levelInfos[level]) - 1) { // 如果新的层和不为0，或者该层的节点数不等于总节点数，则更新最大层和
                res = max(res, newSum)
            }
            // 更新 left 和 right 指针，以及需要减去的值。
            left, right, lostVal = leftInfo.LeftPtr, rightInfo.RightPtr, addVal
        }
    }
    return res
}

func main() {
    // 示例 1：
    // 输入：root = [6,0,3,null,8]
    // 输出：11
    // 解释：勇者的最佳方案如图所示
    // <img src="https://pic.leetcode-cn.com/1648180809-XSWPLu-image.png" />
    tree1 := &TreeNode {
        6,
        &TreeNode{0, nil, &TreeNode{8, nil, nil}, },
        &TreeNode{3, nil, nil},
    }
    fmt.Println(getMaxLayerSum(tree1)) // 11
    // 示例 2：
    // 输入：root = [5,6,2,4,null,null,1,3,5]
    // 输出：9
    // 解释：勇者击破 6 节点，此时「层和」最大为 3+5+1 = 9
    // <img src="https://pic.leetcode-cn.com/1648180769-TLpYop-image.png" />
    tree2 := &TreeNode {
        5,
        &TreeNode{6, &TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}, }, nil, },
        &TreeNode{2, nil, &TreeNode{1, nil, nil}, },
    }
    fmt.Println(getMaxLayerSum(tree2)) // 9
    // 示例 3：
    // 输入：root = [-5,1,7]
    // 输出：8
    // 解释：勇者不击破节点，「层和」最大为 1+7 = 8
    tree3 := &TreeNode {
        -5,
        &TreeNode{1, nil, nil, },
        &TreeNode{7, nil, nil, },
    }
    fmt.Println(getMaxLayerSum(tree3)) // 8
}