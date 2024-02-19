package main

// 589. N-ary Tree Preorder Traversal
// Given the root of an n-ary tree, return the preorder traversal of its nodes' values.
// Nary-Tree input serialization is represented in their level order traversal. 
// Each group of children is separated by the null value (See examples)

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2018/10/12/narytreeexample.png" />
// Input: root = [1,null,3,2,4,null,5,6]
// Output: [1,3,5,6,2,4]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/11/08/sample_4_964.png" />
// Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
// Output: [1,2,3,6,7,11,14,4,8,12,5,9,13,10]

// Constraints:
//         The number of nodes in the tree is in the range [0, 10^4].
//         0 <= Node.val <= 104
//         The height of the n-ary tree is less than or equal to 1000.

// Follow up: Recursive solution is trivial, could you do it iteratively?

//  Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

// 迭代
func preorder(root *Node) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		r := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, r.Val)
		tmp := []*Node{}
        // 将父节点的所有孩子节点逆序入栈，逆序的目的是为了让前序节点永远在栈顶。依次循环直到栈里所有元素都出栈
		for _, v := range r.Children {
			tmp = append([]*Node{v}, tmp...) // 逆序存点
		}
		stack = append(stack, tmp...)
	}
	return res
}

// 递归
func preorder1(root *Node) []int {
	res := []int{}
	preorderdfs(root, &res)
	return res
}

func preorderdfs(root *Node, res *[]int) {
	if root != nil {
		*res = append(*res, root.Val)
		for i := 0; i < len(root.Children); i++ {
			preorderdfs(root.Children[i], res)
		}
	}
}

// best solution
func preorder2(root *Node) []int {
    ret := []int{}
    if root != nil {
        stack := []*Node{root}
        for len(stack) != 0 {
            cur := stack[len(stack) - 1]
            stack = stack[:len(stack)-1]
            ret = append(ret, cur.Val)
            for i := len(cur.Children) - 1;i >= 0;i-= 1 {
                stack = append(stack, cur.Children[i])
            }
        }
    }
    return ret
}

// best solution
func preorder3(root *Node) []int {
    res := []int{}
    if root == nil {
        return res
    }
    res = append(res, root.Val)
    for _, child := range root.Children {
        res = append(res, preorder3(child)...)
    }
    return res
}