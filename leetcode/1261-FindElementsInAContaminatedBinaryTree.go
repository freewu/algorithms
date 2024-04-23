package main

// 1261. Find Elements in a Contaminated Binary Tree
// Given a binary tree with the following rules:
//         root.val == 0
//         If treeNode.val == x and treeNode.left != null, then treeNode.left.val == 2 * x + 1
//         If treeNode.val == x and treeNode.right != null, then treeNode.right.val == 2 * x + 2

// Now the binary tree is contaminated, which means all treeNode.val have been changed to -1.
// Implement the FindElements class:
//         FindElements(TreeNode* root) Initializes the object with a contaminated binary tree and recovers it.
//         bool find(int target) Returns true if the target value exists in the recovered binary tree.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/11/06/untitled-diagram-4-1.jpg" />
// Input
// ["FindElements","find","find"]
// [[[-1,null,-1]],[1],[2]]
// Output
// [null,false,true]
// Explanation
// FindElements findElements = new FindElements([-1,null,-1]); 
// findElements.find(1); // return False 
// findElements.find(2); // return True 

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/11/06/untitled-diagram-4.jpg" />
// Input
// ["FindElements","find","find","find"]
// [[[-1,-1,-1,-1,-1]],[1],[3],[5]]
// Output
// [null,true,true,false]
// Explanation
// FindElements findElements = new FindElements([-1,-1,-1,-1,-1]);
// findElements.find(1); // return True
// findElements.find(3); // return True
// findElements.find(5); // return False

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2019/11/07/untitled-diagram-4-1-1.jpg" />
// Input
// ["FindElements","find","find","find","find"]
// [[[-1,null,-1,-1,null,-1]],[2],[3],[4],[5]]
// Output
// [null,true,false,false,true]
// Explanation
// FindElements findElements = new FindElements([-1,null,-1,-1,null,-1]);
// findElements.find(2); // return True
// findElements.find(3); // return False
// findElements.find(4); // return False
// findElements.find(5); // return True
 
// Constraints:
//         TreeNode.val == -1
//         The height of the binary tree is less than or equal to 20
//         The total number of nodes is between [1, 10^4]
//         Total calls of find() is between [1, 10^4]
//         0 <= target <= 10^6

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
type FindElements struct {
    m map[int]struct{}
}

func Constructor(root *TreeNode) FindElements {
    m := make(map[int]struct{})
    var recoverTree func(*TreeNode, int)
    recoverTree = func(root *TreeNode, val int)  {
        if root == nil {
            return
        }
        // 把二叉树的值存入 map 中 用于 find
        m[val] = struct{}{}
        // 如果 treeNode.val == x 且 treeNode.left != null，那么 treeNode.left.val == 2 * x + 1
        recoverTree(root.Left, 2 * val + 1)
        // 如果 treeNode.val == x 且 treeNode.right != null，那么 treeNode.right.val == 2 * x + 2
        recoverTree(root.Right, 2 * val + 2)
    }
    // root.val == 0
    recoverTree(root, 0)
    return FindElements {
        m: m,
    }
}

func (this *FindElements) Find(target int) bool {
    if _, ok := this.m[target]; ok {
        return true
    }
    return false
}

type FindElements1 struct {
    root *TreeNode
    nodeCache map[int]struct{}
}

func Constructor1(root *TreeNode) FindElements1 {
    var recover func(*TreeNode, int)
    cache := make(map[int]struct{})
    recover = func(root *TreeNode, v int) {
        root.Val = v
        cache[v] = struct{}{}
        if root.Left != nil {
            recover(root.Left, 2*v+1)
        }
        if root.Right != nil {
            recover(root.Right, 2*v+2)
        }
    }
    recover(root, 0)
    return FindElements1{
        root: root,
        nodeCache: cache,
    }
}

func (this *FindElements1) Find(target int) bool {
    _, ok := this.nodeCache[target]
    return ok
}
 
 
/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
func main() {
    tree1 := &TreeNode {
        -1,
        nil,
        &TreeNode { -1, nil, nil },
    }
    obj := Constructor(tree1);
    fmt.Println(obj.Find(1)) // false
    fmt.Println(obj.Find(2)) // true

    tree2 := &TreeNode {
        -1,
        &TreeNode {
            -1,
            &TreeNode { -1 ,nil, nil },
            &TreeNode { -1 ,nil, nil },
        },
        &TreeNode { -1 ,nil, nil },
    }
    obj = Constructor(tree2);
    fmt.Println(obj.Find(1)) // true
    fmt.Println(obj.Find(3)) // true
    fmt.Println(obj.Find(5)) // false

    tree3 := &TreeNode {
        -1,
        nil,
        &TreeNode {
            -1,
            &TreeNode { 
                -1,
                &TreeNode { -1, nil, nil }, 
                nil,
            },
            nil,
        },
    }
    obj = Constructor(tree3);
    fmt.Println(obj.Find(2)) // true
    fmt.Println(obj.Find(3)) // false
    fmt.Println(obj.Find(4)) // false
    fmt.Println(obj.Find(5)) // true


    obj1 := Constructor1(tree1);
    fmt.Println(obj1.Find(1)) // false
    fmt.Println(obj1.Find(2)) // true

    obj1 = Constructor1(tree2);
    fmt.Println(obj1.Find(1)) // true
    fmt.Println(obj1.Find(3)) // true
    fmt.Println(obj1.Find(5)) // false

    obj1 = Constructor1(tree3);
    fmt.Println(obj1.Find(2)) // true
    fmt.Println(obj1.Find(3)) // false
    fmt.Println(obj1.Find(4)) // false
    fmt.Println(obj1.Find(5)) // true
}