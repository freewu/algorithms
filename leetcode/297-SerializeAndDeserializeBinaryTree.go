package main

// 297. Serialize and Deserialize Binary Tree
// Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, 
// or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

// Design an algorithm to serialize and deserialize a binary tree. 
// There is no restriction on how your serialization/deserialization algorithm should work. 
// You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

// Clarification: 
//     The input/output format is the same as how LeetCode serializes a binary tree. 
//     You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/09/15/serdeser.jpg" />
// Input: root = [1,2,3,null,null,4,5]
// Output: [1,2,3,null,null,4,5]

// Example 2:
// Input: root = []
// Output: []
 
// Constraints:
//     The number of nodes in the tree is in the range [0, 10^4].
//     -1000 <= Node.val <= 1000

import "fmt"
import "bytes"
import "strconv"
import "strings"

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

type Codec struct {

}
 
func Constructor() Codec {
    return Codec{}
}

// dfs
func (this *Codec) serialize(root *TreeNode) string {
    var buffer bytes.Buffer
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            buffer.WriteString("#,")
        } else {
            buffer.WriteString(strconv.Itoa(node.Val))
            buffer.WriteString(",")
            dfs(node.Left)
            dfs(node.Right)
        }
    }
    dfs(root)
    return buffer.String()
}

// dfs
func (this *Codec) deserialize(data string) *TreeNode {    
    tokens := strings.Split(data, ",")
    var dfs func() *TreeNode
    dfs = func() *TreeNode {
        token := tokens[0]
        tokens = tokens[1:]
        if token == "#" { return nil }
        val, _ := strconv.Atoi(token)
        return &TreeNode{val, dfs(), dfs() }
    }
    return dfs()
}
 
 /**
  * Your Codec object will be instantiated and called as such:
  * ser := Constructor();
  * deser := Constructor();
  * data := ser.serialize(root);
  * ans := deser.deserialize(data);
  */
// bfs
type Codec1 struct {
}

func Constructor1() Codec1 {
	return Codec1{}
}

// Serializes a tree to a single string.
func (this *Codec1) serialize(root *TreeNode) string {
    res := ""
    if root == nil {
        return res
    }
    queue := []*TreeNode{}
    queue = append(queue, root)
    for len(queue) > 0 {
        temp := queue[0]
        queue = queue[1:]
        if temp == nil {
            res += "#,"
        } else {
            res += strconv.Itoa(temp.Val) + ","
        }
        if temp != nil{
            queue = append(queue, temp.Left)
            queue = append(queue, temp.Right)
        }
    }
    return res
}

// Deserializes your encoded data to tree.
func (this *Codec1) deserialize(data string) *TreeNode {
    if data == "" {
        return nil
    }
    nodes := strings.Split(data, ",")
    rootVal, _ := strconv.Atoi(nodes[0])
    root, queue, index := &TreeNode{Val: rootVal}, []*TreeNode{}, 1
    queue = append(queue, root)

    for len(queue) > 0 {
        temp := queue[0]
        queue = queue[1:]
        if nodes[index] != "#" {
            val, _ := strconv.Atoi(nodes[index])
            temp.Left = &TreeNode{Val: val}
            queue = append(queue, temp.Left)
        }
        index++
        if nodes[index] != "#" {
            val, _ := strconv.Atoi(nodes[index])
            temp.Right = &TreeNode{Val: val}
            queue = append(queue, temp.Right)
        }
        index++
    }
    return root
}

func main() {
    tree1 := &TreeNode {
        1,
        &TreeNode{2, nil, nil},
        &TreeNode {
            3,
            &TreeNode{4, nil, nil},
            &TreeNode{5, nil, nil},
        },
    }
    obj := Constructor()
    data := obj.serialize(tree1)
    fmt.Println(data)
    t := obj.deserialize(data)
    fmt.Println(t.Val)
    fmt.Println(t.Left.Val)
    fmt.Println(t.Right.Val)

    obj1 := Constructor1()
    data1 := obj1.serialize(tree1)
    fmt.Println(data1)
    t1 := obj1.deserialize(data1)
    fmt.Println(t1.Val)
    fmt.Println(t1.Left.Val)
    fmt.Println(t1.Right.Val)
}