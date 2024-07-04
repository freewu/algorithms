package main

// LCR 043. 完全二叉树插入器
// 完全二叉树是每一层（除最后一层外）都是完全填充（即，节点数达到最大，第 n 层有 2n-1 个节点）的，并且所有的节点都尽可能地集中在左侧。

// 设计一个用完全二叉树初始化的数据结构 CBTInserter，它支持以下几种操作：
//     CBTInserter(TreeNode root) 
//         使用根节点为 root 的给定树初始化该数据结构；
//     CBTInserter.insert(int v)  
//         向树中插入一个新节点，节点类型为 TreeNode，值为 v 。
//         使树保持完全二叉树的状态，并返回插入的新节点的父节点的值；
//     CBTInserter.get_root() 
//         将返回树的根节点。

// 示例 1：
// 输入：inputs = ["CBTInserter","insert","get_root"], inputs = [[[1]],[2],[]]
// 输出：[null,1,[1,2]]

// 示例 2：
// 输入：inputs = ["CBTInserter","insert","insert","get_root"], inputs = [[[1,2,3,4,5,6]],[7],[8],[]]
// 输出：[null,3,4,[1,2,3,4,5,6,7,8]]

// 提示：
//     最初给定的树是完全二叉树，且包含 1 到 1000 个节点。
//     每个测试用例最多调用 CBTInserter.insert  操作 10000 次。
//     给定节点或插入节点的每个值都在 0 到 5000 之间。

import "fmt"
import "math/bits"

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
type CBTInserter struct { 
    treeArr []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
    return CBTInserter{treeArr: bfs(root)} 
}
  
func bfs(root *TreeNode) []*TreeNode {
    queue, res := []*TreeNode{}, []*TreeNode{}
    queue = append(queue, root)
    for {
        node := queue[0]
        queue = queue[1:]
        res = append(res, node)
        if node.Left != nil {
            queue = append(queue, node.Left)   
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
        }
        if len(queue) == 0 {
            break
        }
    }
    return res 
}

func (this *CBTInserter) Insert(v int) int {
    length := len(this.treeArr)
    insert_loc := ((float64(length) - 1) / 2.0)
    loc := int64(insert_loc)
    parent := this.treeArr[loc]
    node := &TreeNode{Val: v}
    if parent.Left == nil {
        parent.Left = node
    } else {
        parent.Right = node
    }
    this.treeArr = append(this.treeArr, node)
    return parent.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
    return this.treeArr[0] 
}

type CBTInserter1 struct {
    root *TreeNode
    count  int
}

func Constructor1(root *TreeNode) CBTInserter1 {
    queue := []*TreeNode{root}
    count := 0
    for len(queue) > 0 { // 得到树的节点个数
        count++
        node := queue[0]
        queue = queue[1:]
        if node.Left != nil {
            queue = append(queue, node.Left)
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
        }
    }
    return CBTInserter1{root, count}
}

func (this *CBTInserter1) Insert(val int) int {
    this.count++
    child := &TreeNode{ Val: val }
    node := this.root
    for i := bits.Len(uint(this.count)) - 2; i > 0; i-- {
        if this.count >> i & 1 == 0 {
            node = node.Left
        } else {
            node = node.Right
        }
    }
    if this.count & 1 == 0 {
        node.Left = child
    } else {
        node.Right = child
    }
    return node.Val
}

func (this *CBTInserter1) Get_root() *TreeNode {
    return this.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */

func main() {
    // CBTInserter cBTInserter = new CBTInserter([1, 2]);
    tree1 := &TreeNode{
        1, 
        &TreeNode{2, nil, nil},
        nil,
    }
    obj := Constructor(tree1)
    fmt.Println(obj)
    // cBTInserter.insert(3);  // return 1
    fmt.Println(obj.Insert(3)) // 1
    fmt.Println(obj)
    // cBTInserter.insert(4);  // return 2
    fmt.Println(obj.Insert(4)) // 2
    fmt.Println(obj)
    // cBTInserter.get_root(); // return [1, 2, 3, 4]
    fmt.Println(obj.Get_root()) // [1, 2, 3, 4]

    tree11 := &TreeNode{
        1, 
        &TreeNode{2, nil, nil},
        nil,
    }
    obj1 := Constructor1(tree11)
    fmt.Println(obj1)
    // cBTInserter.insert(3);  // return 1
    fmt.Println(obj1.Insert(3)) // 1
    fmt.Println(obj1)
    // cBTInserter.insert(4);  // return 2
    fmt.Println(obj1.Insert(4)) // 2
    fmt.Println(obj1)
    // cBTInserter.get_root(); // return [1, 2, 3, 4]
    fmt.Println(obj1.Get_root()) // [1, 2, 3, 4]
}