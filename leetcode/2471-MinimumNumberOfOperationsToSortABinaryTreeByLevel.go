package main

// 2471. Minimum Number of Operations to Sort a Binary Tree by Level
// You are given the root of a binary tree with unique values.

// In one operation, you can choose any two nodes at the same level and swap their values.

// Return the minimum number of operations needed to make the values at each level sorted in a strictly increasing order.

// The level of a node is the number of edges along the path between it and the root node.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/09/18/image-20220918174006-2.png" />
//             1
//           /    \
//          4      3
//        /   \   /   \
//       7    6  8     5
//              /     /
//             9     10
// Input: root = [1,4,3,7,6,8,5,null,null,null,null,9,null,10]
// Output: 3
// Explanation:
// - Swap 4 and 3. The 2nd level becomes [3,4].
// - Swap 7 and 5. The 3rd level becomes [5,6,8,7].
// - Swap 8 and 7. The 3rd level becomes [5,6,7,8].
// We used 3 operations so return 3.
// It can be proven that 3 is the minimum number of operations needed.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/09/18/image-20220918174026-3.png" />
//             1
//          /     \
//         3       2
//       /   \   /    \
//      7     6 5      4
// Input: root = [1,3,2,7,6,5,4]
// Output: 3
// Explanation:
// - Swap 3 and 2. The 2nd level becomes [2,3].
// - Swap 7 and 4. The 3rd level becomes [4,6,5,7].
// - Swap 6 and 5. The 3rd level becomes [4,5,6,7].
// We used 3 operations so return 3.
// It can be proven that 3 is the minimum number of operations needed.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2022/09/18/image-20220918174052-4.png" />
//            1
//          /    \
//         2      3
//       /   \    / 
//      4     5  6
// Input: root = [1,2,3,4,5,6]
// Output: 0
// Explanation: Each level is already sorted in increasing order so return 0.

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^5].
//     1 <= Node.val <= 10^5
//     All the values of the tree are unique.

import "fmt"
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
func minimumOperations(root *TreeNode) int {
    if root == nil { return 0 }
    helper := func(nums []int) int {
        sorted := make([]int, len(nums))
        copy(sorted, nums)
        sort.Ints(sorted)
        res, index := 0, make(map[int]int)
        for i, v := range nums {
            index[v] = i
        }
        for i := range sorted {
            if sorted[i] != nums[i] {
                res++
                newPos := index[sorted[i]]
                index[nums[i]] = newPos
                index[nums[newPos]] = i
                nums[i], nums[newPos] = nums[newPos], nums[i]
            }
        }
        return res
    }
    res, queue := 0, []*TreeNode{root}
    for len(queue) > 0 {
        n := len(queue)
        arr := make([]int, 0, n)
        
        for i := 0; i < n; i++ {
            top := queue[0]
            queue = queue[1:] // pop
            arr = append(arr, top.Val)
            if top.Left != nil  { queue = append(queue, top.Left)  }
            if top.Right != nil { queue = append(queue, top.Right) }
        }
        res += helper(arr)
    }
    return res
}

func minimumOperations1(root *TreeNode) int {
    res, queue := 0, []*TreeNode{root}
    for len(queue) > 0 {
        n := len(queue)
        arr := make([]int, n)
        tmp := queue
        queue = nil
        for i, node := range tmp {
            arr[i] = node.Val
            if node.Left != nil  { queue = append(queue, node.Left) }
            if node.Right != nil { queue = append(queue, node.Right) }
        }
        index := make([]int, n)
        for i := range index {
            index[i] = i
        }
        sort.Slice(index, func(i, j int) bool {
            return arr[index[i]] < arr[index[j]]
        })
        res += n
        visited := make([]bool, n)
        for _, v := range index {
            if !visited[v] {
                for ; !visited[v]; v = index[v] {
                    visited[v] = true
                }
                res--
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/09/18/image-20220918174006-2.png" />
    //             1
    //           /    \
    //          4      3
    //        /   \   /   \
    //       7    6  8     5
    //              /     /
    //             9     10
    // Input: root = [1,4,3,7,6,8,5,null,null,null,null,9,null,10]
    // Output: 3
    // Explanation:
    // - Swap 4 and 3. The 2nd level becomes [3,4].
    // - Swap 7 and 5. The 3rd level becomes [5,6,8,7].
    // - Swap 8 and 7. The 3rd level becomes [5,6,7,8].
    // We used 3 operations so return 3.
    // It can be proven that 3 is the minimum number of operations needed.
    tree1 := &TreeNode {
        1,
        &TreeNode { 4, &TreeNode { 7, nil, nil }, &TreeNode { 6, nil, nil }, },
        &TreeNode { 3, &TreeNode { 8, &TreeNode { 9, nil, nil }, nil }, &TreeNode { 5, &TreeNode { 10, nil, nil }, nil }, },
    }
    fmt.Println(minimumOperations(tree1)) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/09/18/image-20220918174026-3.png" />
    //             1
    //          /     \
    //         3       2
    //       /   \   /    \
    //      7     6 5      4
    // Input: root = [1,3,2,7,6,5,4]
    // Output: 3
    // Explanation:
    // - Swap 3 and 2. The 2nd level becomes [2,3].
    // - Swap 7 and 4. The 3rd level becomes [4,6,5,7].
    // - Swap 6 and 5. The 3rd level becomes [4,5,6,7].
    // We used 3 operations so return 3.
    // It can be proven that 3 is the minimum number of operations needed.
    tree2 := &TreeNode {
        1,
        &TreeNode { 3, &TreeNode { 7, nil, nil }, &TreeNode { 6, nil, nil }, },
        &TreeNode { 2, &TreeNode { 5, nil, nil }, &TreeNode { 4, nil, nil }, },
    }
    fmt.Println(minimumOperations(tree2)) // 3
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2022/09/18/image-20220918174052-4.png" />
    //            1
    //          /    \
    //         2      3
    //       /   \    / 
    //      4     5  6
    // Input: root = [1,2,3,4,5,6]
    // Output: 0
    // Explanation: Each level is already sorted in increasing order so return 0.
    tree3 := &TreeNode {
        1,
        &TreeNode { 2, &TreeNode { 4, nil, nil }, &TreeNode { 5, nil, nil }, },
        &TreeNode { 3, &TreeNode { 6, nil, nil }, nil, },
    }
    fmt.Println(minimumOperations(tree3)) // 0

    fmt.Println(minimumOperations1(tree1)) // 3
    fmt.Println(minimumOperations1(tree2)) // 3
    fmt.Println(minimumOperations1(tree3)) // 0
}