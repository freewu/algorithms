package main

// 2476. Closest Nodes Queries in a Binary Search Tree
// You are given the root of a binary search tree and an array queries of size n consisting of positive integers.
// Find a 2D array answer of size n where answer[i] = [mini, maxi]:
//     mini is the largest value in the tree that is smaller than or equal to queries[i]. 
//     If a such value does not exist, add -1 instead.
//     maxi is the smallest value in the tree that is greater than or equal to queries[i]. 
//     If a such value does not exist, add -1 instead.

// Return the array answer.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/09/28/bstreeedrawioo.png" />
// Input: root = [6,2,13,1,4,9,15,null,null,null,null,null,null,14], queries = [2,5,16]
// Output: [[2,2],[4,6],[15,-1]]
// Explanation: We answer the queries in the following way:
// - The largest number that is smaller or equal than 2 in the tree is 2, and the smallest number that is greater or equal than 2 is still 2. So the answer for the first query is [2,2].
// - The largest number that is smaller or equal than 5 in the tree is 4, and the smallest number that is greater or equal than 5 is 6. So the answer for the second query is [4,6].
// - The largest number that is smaller or equal than 16 in the tree is 15, and the smallest number that is greater or equal than 16 does not exist. So the answer for the third query is [15,-1].

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/09/28/bstttreee.png" />
// Input: root = [4,null,9], queries = [3]
// Output: [[-1,4]]
// Explanation: The largest number that is smaller or equal to 3 in the tree does not exist, and the smallest number that is greater or equal to 3 is 4. So the answer for the query is [-1,4].
 
// Constraints:
//         The number of nodes in the tree is in the range [2, 10^5].
//         1 <= Node.val <= 10^6
//         n == queries.length
//         1 <= n <= 10^5
//         1 <= queries[i] <= 10^6

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

// Time Limit Exceeded
func closestNodes(root *TreeNode, queries []int) [][]int {
    res := [][]int{}
    for _ , n := range queries {
        // answer[i] = [mini, maxi] ：
        // mini 是树中小于等于 queries[i] 的 最大值 。如果不存在这样的值，则使用 -1 代替。
        // maxi 是树中大于等于 queries[i] 的 最小值 。如果不存在这样的值，则使用 -1 代替。
        curr := []int{-1,-1}
        traverseTree(root, n, &curr)
        res = append(res, curr)
    }
    return res
}

func traverseTree(root *TreeNode, n int, res *[]int) {
    if root != nil {
        if root.Val < n { // 树中小于等于 queries[i] 的 最大值
            (*res)[0] = root.Val
            traverseTree(root.Right, n, res)
        } else if root.Val > n { // 树中大于等于 queries[i] 的 最小值
            (*res)[1] = root.Val
            traverseTree(root.Left, n, res)
        } else {
            (*res)[0] = n
            (*res)[1] = n
        }
    }
}

// 由于该二叉搜索树并不是平衡的，则最坏情况该树可能形成一条链，直接在树上查找可能存在超时。
// 我们可以保存树中所有的节点值，并将其排序，每次给定查询值 val 时，利用二分查找直接在树中找到大于等于 val 的最小值与小于等于 val 的最小值
func closestNodes1(root *TreeNode, queries []int) [][]int {
    arr := []int{}
    var dfs func(*TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        dfs(root.Left)
        arr = append(arr, root.Val)
        dfs(root.Right)
    }
    
    dfs(root)
    res := make([][]int, len(queries))
    for i, val := range queries {
        maxVal, minVal := -1, -1
        index := sort.SearchInts(arr, val)
        if index < len(arr) {
            maxVal = arr[index]
            if arr[index] == val {
                minVal = arr[index]
                res[i] = []int{minVal, maxVal}
                continue
            }
        }
        if index != 0 {
            minVal = arr[index - 1]
        }
        res[i] = []int{minVal, maxVal}
    }
    return res
}

// best solution
func closestNodes2(root *TreeNode, queries []int) [][]int {
	a := []int{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		a = append(a, node.Val)
		dfs(node.Right)
	}
	dfs(root)

	ans := make([][]int, len(queries))
	for i, q := range queries {
		mn, mx := -1, -1
		j, ok := slices.BinarySearch(a, q)
		if j < len(a) {
			mx = a[j]
		}
		if !ok { // a[j]>q, a[j-1]<q
			j--
		}
		if j >= 0 {
			mn = a[j]
		}
		ans[i] = []int{mn, mx}
	}
	return ans
}