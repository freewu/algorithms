package main

// 919. Complete Binary Tree Inserter
// A complete binary tree is a binary tree in which every level, except possibly the last, is completely filled, and all nodes are as far left as possible.
// Design an algorithm to insert a new node to a complete binary tree keeping it complete after the insertion.

// Implement the CBTInserter class:
//     CBTInserter(TreeNode root) 
//         Initializes the data structure with the root of the complete binary tree.
//     int insert(int v) 
//         Inserts a TreeNode into the tree with value Node.val == val so that the tree remains complete, and returns the value of the parent of the inserted TreeNode.
//     TreeNode get_root() 
//         Returns the root node of the tree.

// Example 1:
// <img src="" />
// Input
// ["CBTInserter", "insert", "insert", "get_root"]
// [[[1, 2]], [3], [4], []]
// Output
// [null, 1, 2, [1, 2, 3, 4]]
// Explanation
// CBTInserter cBTInserter = new CBTInserter([1, 2]);
// cBTInserter.insert(3);  // return 1
// cBTInserter.insert(4);  // return 2
// cBTInserter.get_root(); // return [1, 2, 3, 4]

// Constraints:
//     The number of nodes in the tree will be in the range [1, 1000].
//     0 <= Node.val <= 5000
//     root is a complete binary tree.
//     0 <= val <= 5000
//     At most 10^4 calls will be made to insert and get_root.

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