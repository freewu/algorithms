package main

// 449. Serialize and Deserialize BST
// Serialization is converting a data structure or object into a sequence of bits so 
// that it can be stored in a file or memory buffer, 
// or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

// Design an algorithm to serialize and deserialize a binary search tree. 
// There is no restriction on how your serialization/deserialization algorithm should work. 
// You need to ensure that a binary search tree can be serialized to a string, 
// and this string can be deserialized to the original tree structure.

// The encoded string should be as compact as possible.

// Example 1:
// Input: root = [2,1,3]
// Output: [2,1,3]

// Example 2:
// Input: root = []
// Output: []
 
// Constraints:
//     The number of nodes in the tree is in the range [0, 10^4].
//     0 <= Node.val <= 10^4
//     The input tree is guaranteed to be a binary search tree.

import "fmt"
import "strings"
import "container/list"
import "strconv"
import "sort"

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

type Codec1 struct {

}
 
func Constructor1() Codec1 {
    return Codec1{}
}

// Serializes a tree to a single string.
func (this *Codec1) serialize(root *TreeNode) string {
    result, queue := "", list.New()
    if root != nil {
        queue.PushBack(root)
    }
    for queue.Len() > 0 {
        qSize := queue.Len()
        for i := 0; i < qSize; i++ {
            current := queue.Remove(queue.Front()).(*TreeNode)
            if result != "" {
                result += ","
            }
            if current != nil {
                result += fmt.Sprint(current.Val)
                queue.PushBack(current.Left)
                queue.PushBack(current.Right)
            } else {
                result += "null"
            }
        }
    }

    return result
}

// Deserializes your encoded data to tree.
func (this *Codec1) deserialize(data string) *TreeNode {
    if data == "" {
        return nil
    }
    values := strings.Split(data, ",")
    nodes := []*TreeNode{}
    for _, value := range values {
        if value == "null" {
            nodes = append(nodes, nil)
        } else {
            val, _ := strconv.Atoi(value)
            nodes = append(nodes, &TreeNode{Val: val})
        }
    }
    i := 0
    for _, node := range nodes {
        if node == nil {
            continue
        }
        left := 2 * i + 1
        right := 2 * i + 2
        if left < len(nodes) {
            node.Left = nodes[left]
        }
        if right < len(nodes) {
            node.Right = nodes[right]
        }
        i++
    }
    return nodes[0]
}


type Codec struct {
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return ""
    }
    arr := []string{}
    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        arr = append(arr, strconv.Itoa(root.Val))
        dfs(root.Left)
        dfs(root.Right)
    }
    dfs(root)
    return strings.Join(arr, ",")
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
    if data == "" {
        return nil
    }
    arr := strings.Split(data, ",")
    pre := make([]int, len(arr))
    mid := make([]int, len(arr))
    for i, v := range arr {
        value, _ := strconv.Atoi(v)
        pre[i] = value
        mid[i] = value
    }
    sort.Ints(mid)
    return buildTree(pre, mid)
}

func buildTree(pre, mid []int) *TreeNode {
    if len(pre) == 0 {
        return nil
    }
    v := pre[0]
    index := 0
    for i := 0; i < len(pre); i++ {
        if mid[i] == v {
            index = i
            break
        }
    }
    return &TreeNode{
        Val:   v,
        Left:  buildTree(pre[1:index+1], mid[:index]),
        Right: buildTree(pre[index+1:], mid[index+1:]),
    }
}

 /**
  * Your Codec object will be instantiated and called as such:
  * ser := Constructor()
  * deser := Constructor()
  * tree := ser.serialize(root)
  * ans := deser.deserialize(tree)
  * return ans
  */

func main() {
    obj := Constructor()
    // Example 1:
    // Input: root = [2,1,3]
    // Output: [2,1,3]
    tree1 := &TreeNode {
        2,
        &TreeNode{1, nil, nil},
        &TreeNode{3, nil, nil},
    }
    e1 := obj.serialize(tree1)
    fmt.Println("e1 ", e1)
    d1 := obj.deserialize(e1)
    fmt.Println("d1 ", d1)
    fmt.Println("d1.Left ", d1.Left)
    fmt.Println("d1.Right ", d1.Right)
    // Example 2:
    // Input: root = []
    // Output: []
    tree2 := &TreeNode {}
    e2 := obj.serialize(tree2)
    fmt.Println("e2 ", e2)
    d2 := obj.deserialize(e2)
    fmt.Println("d2 ", d2)


    obj1 := Constructor()
    // Example 1:
    // Input: root = [2,1,3]
    // Output: [2,1,3]
    tree11 := &TreeNode {
        2,
        &TreeNode{1, nil, nil},
        &TreeNode{3, nil, nil},
    }
    e11 := obj1.serialize(tree11)
    fmt.Println("e11 ", e11)
    d11 := obj1.deserialize(e11)
    fmt.Println("d11 ", d11)
    fmt.Println("d11.Left ", d11.Left)
    fmt.Println("d11.Right ", d11.Right)
    // Example 2:
    // Input: root = []
    // Output: []
    tree12 := &TreeNode {}
    e12 := obj1.serialize(tree12)
    fmt.Println("e12 ", e12)
    d12 := obj1.deserialize(e12)
    fmt.Println("d12 ", d12)
}