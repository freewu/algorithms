package main

// 2003. Smallest Missing Genetic Value in Each Subtree
// There is a family tree rooted at 0 consisting of n nodes numbered 0 to n - 1. 
// You are given a 0-indexed integer array parents, where parents[i] is the parent for node i. 
// Since node 0 is the root, parents[0] == -1.

// There are 105 genetic values, each represented by an integer in the inclusive range [1, 10^5].
// You are given a 0-indexed integer array nums, where nums[i] is a distinct genetic value for node i.

// Return an array ans of length n where ans[i] is the smallest genetic value that is missing from the subtree rooted at node i.
// The subtree rooted at a node x contains node x and all of its descendant nodes.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/08/23/case-1.png" />
// Input: parents = [-1,0,0,2], nums = [1,2,3,4]
// Output: [5,1,1,1]
// Explanation: The answer for each subtree is calculated as follows:
// - 0: The subtree contains nodes [0,1,2,3] with values [1,2,3,4]. 5 is the smallest missing value.
// - 1: The subtree contains only node 1 with value 2. 1 is the smallest missing value.
// - 2: The subtree contains nodes [2,3] with values [3,4]. 1 is the smallest missing value.
// - 3: The subtree contains only node 3 with value 4. 1 is the smallest missing value.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/08/23/case-2.png" />
// Input: parents = [-1,0,1,0,3,3], nums = [5,4,6,2,1,3]
// Output: [7,1,1,4,2,1]
// Explanation: The answer for each subtree is calculated as follows:
// - 0: The subtree contains nodes [0,1,2,3,4,5] with values [5,4,6,2,1,3]. 7 is the smallest missing value.
// - 1: The subtree contains nodes [1,2] with values [4,6]. 1 is the smallest missing value.
// - 2: The subtree contains only node 2 with value 6. 1 is the smallest missing value.
// - 3: The subtree contains nodes [3,4,5] with values [2,1,3]. 4 is the smallest missing value.
// - 4: The subtree contains only node 4 with value 1. 2 is the smallest missing value.
// - 5: The subtree contains only node 5 with value 3. 1 is the smallest missing value.

// Example 3:
// Input: parents = [-1,2,3,0,2,4,1], nums = [2,3,4,5,6,7,8]
// Output: [1,1,1,1,1,1,1]
// Explanation: The value 1 is missing from all the subtrees.

// Constraints:
//     n == parents.length == nums.length
//     2 <= n <= 10^5
//     0 <= parents[i] <= n - 1 for i != 0
//     parents[0] == -1
//     parents represents a valid tree.
//     1 <= nums[i] <= 10^5
//     Each nums[i] is distinct.

import "fmt"

func smallestMissingValueSubtree(parents []int, nums []int) []int {
    n := len(parents)
    // geneSet为当前根节点包含的所有基因值的集合；visited 来记录已经访问过的节点，避免重复搜索
    children, geneSet, visited := make([][]int, n), make(map[int]bool), make([]bool, n)
    for i := 1; i < n; i++ {
        children[parents[i]] = append(children[parents[i]], i) // 记录每个节点的子节点
    }
    var dfs func(int)
    dfs = func(node int) {
        if visited[node] { return }// 已访问
        visited[node] = true // 标记为已访问
        geneSet[nums[node]] = true // 录入当前节点
        for _, child := range children[node] { // 继续深度优先搜索
            dfs(child)
        }
    }
    res, node, iNode := make([]int, n), -1, 1 // 结果数组，当前节点，当前子树内缺失的最小基因值
    for i := 0; i < n; i++ {
        res[i] = 1 // 初始化res
        if nums[i] == 1 { // 找到基因值为1的节点，记录下标
            node = i
        }
    }
    for node != -1 {
        dfs(node) // 统计以node为根节点的子树，并将所含基因值记录到geneSet中
        for geneSet[iNode] { // geneSet[iNode]==false，说明当前子树不含iNode，iNode为当前子树内缺失的最小基因值
            iNode++
        }
        res[node], node = iNode, parents[node] // 更新当前节点的res，并移步到祖先节点，去更新祖先节点的res
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/08/23/case-1.png" />
    // Input: parents = [-1,0,0,2], nums = [1,2,3,4]
    // Output: [5,1,1,1]
    // Explanation: The answer for each subtree is calculated as follows:
    // - 0: The subtree contains nodes [0,1,2,3] with values [1,2,3,4]. 5 is the smallest missing value.
    // - 1: The subtree contains only node 1 with value 2. 1 is the smallest missing value.
    // - 2: The subtree contains nodes [2,3] with values [3,4]. 1 is the smallest missing value.
    // - 3: The subtree contains only node 3 with value 4. 1 is the smallest missing value.
    fmt.Println(smallestMissingValueSubtree([]int{-1,0,0,2}, []int{1,2,3,4})) // [5,1,1,1]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/08/23/case-2.png" />
    // Input: parents = [-1,0,1,0,3,3], nums = [5,4,6,2,1,3]
    // Output: [7,1,1,4,2,1]
    // Explanation: The answer for each subtree is calculated as follows:
    // - 0: The subtree contains nodes [0,1,2,3,4,5] with values [5,4,6,2,1,3]. 7 is the smallest missing value.
    // - 1: The subtree contains nodes [1,2] with values [4,6]. 1 is the smallest missing value.
    // - 2: The subtree contains only node 2 with value 6. 1 is the smallest missing value.
    // - 3: The subtree contains nodes [3,4,5] with values [2,1,3]. 4 is the smallest missing value.
    // - 4: The subtree contains only node 4 with value 1. 2 is the smallest missing value.
    // - 5: The subtree contains only node 5 with value 3. 1 is the smallest missing value.
    fmt.Println(smallestMissingValueSubtree([]int{-1,0,1,0,3,3}, []int{5,4,6,2,1,3})) // [7,1,1,4,2,1]
    // Example 3:
    // Input: parents = [-1,2,3,0,2,4,1], nums = [2,3,4,5,6,7,8]
    // Output: [1,1,1,1,1,1,1]
    // Explanation: The value 1 is missing from all the subtrees.
    fmt.Println(smallestMissingValueSubtree([]int{-1,2,3,0,2,4,1}, []int{2,3,4,5,6,7,8})) // [1,1,1,1,1,1,1]
}