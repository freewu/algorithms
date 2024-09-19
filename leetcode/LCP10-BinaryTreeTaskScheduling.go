package main

// LCP 10. 二叉树任务调度
// 任务调度优化是计算机性能优化的关键任务之一。
// 在任务众多时，不同的调度策略可能会得到不同的总体执行时间，因此寻求一个最优的调度方案是非常有必要的。

// 通常任务之间是存在依赖关系的，即对于某个任务，你需要先完成他的前导任务（如果非空），才能开始执行该任务。
// 我们保证任务的依赖关系是一棵二叉树，其中 root 为根任务，root.left 和 root.right 为他的两个前导任务（可能为空），
// root.val 为其自身的执行时间。

// 在一个 CPU 核执行某个任务时，我们可以在任何时刻暂停当前任务的执行，并保留当前执行进度。
// 在下次继续执行该任务时，会从之前停留的进度开始继续执行。
// 暂停的时间可以不是整数。

// 现在，系统有两个 CPU 核，即我们可以同时执行两个任务，但是同一个任务不能同时在两个核上执行。
// 给定这颗任务树，请求出所有任务执行完毕的最小时间。

// 示例 1：
//         47
//       /    \
//     74     31
// <img src="https://pic.leetcode-cn.com/3522fbf8ce4ebb20b79019124eb9870109fdfe97fe9da99f6c20c07ceb1c60b3-image.png" />
// 输入：root = [47, 74, 31]
// 输出：121
// 解释：根节点的左右节点可以并行执行31分钟，剩下的43+47分钟只能串行执行，因此总体执行时间是121分钟。

// 示例 2：
//             15
//            /
//          12
//         /
//        24
//       /   \
//      27   26
// <img src="https://pic.leetcode-cn.com/13accf172ee4a660d241e25901595d55b759380b090890a17e6e7bd51a143e3f-image.png" />
// 输入：root = [15, 21, null, 24, null, 27, 26]
// 输出：87

// 示例 3：
//         1
//       /   \
//      3     2
//          /    \
//         4      4
// <img src="https://pic.leetcode-cn.com/bef743a12591aafb9047dd95d335b8083dfa66e8fdedc63f50fd406b4a9d163a-image.png" />
// 输入：root = [1,3,2,null,null,4,4]
// 输出：7.5

// 限制：
//     1 <= 节点数量 <= 1000
//     1 <= 单节点执行时间 <= 1000

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
func minimalExecTime(root *TreeNode) float64 {
    var dfs func(node *TreeNode) (mx, par float64)
    dfs = func(node *TreeNode) (mx, par float64) {
        if node == nil {
            return 0.0, 0.0
        }
        mx = float64(node.Val)
        lmx, lpar := dfs(node.Left)
        rmx, rpar := dfs(node.Right)
        mx += (lmx + rmx)
        if lmx < rmx {
            lmx, rmx = rmx, lmx
            lpar, rpar = rpar, lpar
        }
        if lmx - rmx <= 2 * lpar {
            par = (lmx + rmx) / 2
        } else {
            par = lpar + rmx
        }
        return mx, par
    }
    mx, par := dfs(root)
    return mx - par
}

func main() {
    // 示例 1：
    //         47
    //       /    \
    //     74     31
    // <img src="https://pic.leetcode-cn.com/3522fbf8ce4ebb20b79019124eb9870109fdfe97fe9da99f6c20c07ceb1c60b3-image.png" />
    // 输入：root = [47, 74, 31]
    // 输出：121
    // 解释：根节点的左右节点可以并行执行31分钟，剩下的43+47分钟只能串行执行，因此总体执行时间是121分钟。
    tree1 := &TreeNode {
        47,
        &TreeNode{74, nil, nil},
        &TreeNode{31, nil, nil},
    }
    fmt.Println(minimalExecTime(tree1)) // 121  (47 + 74)
    // 示例 2：
    //             15
    //            /
    //          12
    //         /
    //        24
    //       /   \
    //      27   26
    // <img src="https://pic.leetcode-cn.com/13accf172ee4a660d241e25901595d55b759380b090890a17e6e7bd51a143e3f-image.png" />
    // 输入：root = [15, 21, null, 24, null, 27, 26]
    // 输出：87
    tree2 := &TreeNode {
        15,
        &TreeNode{21, &TreeNode{24, &TreeNode{27, nil, nil}, &TreeNode{26, nil, nil},}, nil},
        nil,
    }
    fmt.Println(minimalExecTime(tree2)) // 87 (15 + 12 + 24 + 27)
    // 示例 3：
    //         1
    //       /   \
    //      3     2
    //          /    \
    //         4      4
    // <img src="https://pic.leetcode-cn.com/bef743a12591aafb9047dd95d335b8083dfa66e8fdedc63f50fd406b4a9d163a-image.png" />
    // 输入：root = [1,3,2,null,null,4,4]
    // 输出：7.5
    tree3 := &TreeNode {
        1,
        &TreeNode{3, nil, nil},
        &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{4, nil, nil}, },
    }
    fmt.Println(minimalExecTime(tree3)) // 7.5
}