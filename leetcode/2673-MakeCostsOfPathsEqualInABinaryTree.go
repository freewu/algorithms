package main

// 2673. Make Costs of Paths Equal in a Binary Tree
// You are given an integer n representing the number of nodes in a perfect binary tree consisting of nodes numbered from 1 to n. 
// The root of the tree is node 1 and each node i in the tree has two children where the left child is the node 2 * i and the right child is 2 * i + 1.
// Each node in the tree also has a cost represented by a given 0-indexed integer array cost of size n where cost[i] is the cost of node i + 1. 
// You are allowed to increment the cost of any node by 1 any number of times.
// Return the minimum number of increments you need to make the cost of paths from the root to each leaf node equal.
// Note:
//         A perfect binary tree is a tree where each node, except the leaf nodes, has exactly 2 children.
//         The cost of a path is the sum of costs of nodes in the path.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/04/04/binaryytreeedrawio-4.png" />
// Input: n = 7, cost = [1,5,2,2,3,3,1]
// Output: 6
// Explanation: We can do the following increments:
// - Increase the cost of node 4 one time.
// - Increase the cost of node 3 three times.
// - Increase the cost of node 7 two times.
// Each path from the root to a leaf will have a total cost of 9.
// The total increments we did is 1 + 3 + 2 = 6.
// It can be shown that this is the minimum answer we can achieve.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/04/04/binaryytreee2drawio.png" />
// Input: n = 3, cost = [5,3,3]
// Output: 0
// Explanation: The two paths already have equal total costs, so no increments are needed.
 
// Constraints:
//         3 <= n <= 10^5
//         n + 1 is a power of 2
//         cost.length == n
//         1 <= cost[i] <= 10^4

import "fmt"

func minIncrements(n int, cost []int) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, sums := 0, make([]int, n / 2) // 满二叉树 路径结果一半
    for i := n - 1; i > 1; i -= 2 {
        leftSum, rightSum := cost[i - 1], cost[i]
        if i < len(sums) {
            leftSum += sums[i - 1]
            rightSum += sums[i]
        }
        sums[i / 2 - 1] = max(leftSum, rightSum)
        res += sums[i / 2 - 1] - min(leftSum, rightSum)
    }
    return res
}

func minIncrements1(n int, cost []int) int {
    res := 0
    for i := 0; i < (n - 1) / 2; i++ { // 累加路径和
        cost[2 * i + 1] += cost[i]
        cost[2 * i + 2] += cost[i]
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := (n - 3)/2; i >= 0; i-- { // 反向构建，根为孩子中最大值
        cost[i] = max(cost[2 * i + 1], cost[2 * i + 2])
    }
    for i := 1; i < n; i++ { // 计算当前节点和父亲节点的差值就是答案,
        res += cost[(i-1)/2] - cost[i]
    }
    return res
}

// best solution
func minIncrements2(n int, cost []int) int {
    res := 0
    for i := n / 2; i >= 1; i-- { // 从最后一个非叶节点开始算
        left, right := cost[i*2-1], cost[i*2]
        if left > right { // 保证 left <= right
            left, right = right, left
        }
        res += right - left // 两个子节点变成一样的
        cost[i-1] += right // 累加路径和
    }
    return res
}

func main() {
    fmt.Println(minIncrements(7,[]int{1,5,2,2,3,3,1})) // 6
    fmt.Println(minIncrements(3,[]int{5,3,3})) // 0

    fmt.Println(minIncrements1(7,[]int{1,5,2,2,3,3,1})) // 6
    fmt.Println(minIncrements1(3,[]int{5,3,3})) // 0

    fmt.Println(minIncrements2(7,[]int{1,5,2,2,3,3,1})) // 6
    fmt.Println(minIncrements2(3,[]int{5,3,3})) // 0
}