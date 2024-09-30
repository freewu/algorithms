package main

// 1361. Validate Binary Tree Nodes
// You have n binary tree nodes numbered from 0 to n - 1 where node i has two children leftChild[i] and rightChild[i], 
// return true if and only if all the given nodes form exactly one valid binary tree.

// If node i has no left child then leftChild[i] will equal -1, similarly for the right child.

// Note that the nodes have no values and that we only use the node numbers in this problem.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/08/23/1503_ex1.png" />
// Input: n = 4, leftChild = [1,-1,3,-1], rightChild = [2,-1,-1,-1]
// Output: true

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/08/23/1503_ex2.png" />
// Input: n = 4, leftChild = [1,-1,3,-1], rightChild = [2,3,-1,-1]
// Output: false

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2019/08/23/1503_ex3.png" />
// Input: n = 2, leftChild = [1,0], rightChild = [-1,-1]
// Output: false

// Constraints:
//     n == leftChild.length == rightChild.length
//     1 <= n <= 10^4
//     -1 <= leftChild[i], rightChild[i] <= n - 1

import "fmt"

// // 解答错误 45 / 51
// func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
//     checker := make([]int, n)
//     // check parent count: [0, 1]
//     for i := 0; i < n; i++ {
//         if leftChild[i] != -1 {
//             checker[leftChild[i]]++
//             if checker[leftChild[i]] > 1 {
//                 return false
//             }
//         }
//         if rightChild[i] != -1 {
//             checker[rightChild[i]]++
//             if checker[rightChild[i]] > 1 {
//                 return false
//             }
//         }
//     }
//     // check root count == 1
//     rootCount := 0
//     for _, v := range checker {
//         if v == 0 {
//             rootCount++
//             if rootCount > 1 {
//                 return false
//             }
//         }
//     }
//     return rootCount == 1
// }

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
    treeMap := make(map[int]int)
    for i := 0; i < n; i++ {
        treeMap[leftChild[i]]++
        treeMap[rightChild[i]]++
    }
    // find root node
    root := -1
    for i := 0; i < n; i++ {
        if treeMap[i] == 0 {
            if root != -1 { return false } // 存在两个根节点返回 false
            root = i
            continue
        }
        if treeMap[i] != 1 { return false } // 存在两个父节点返回 false
    }
    if root == -1 { return false } // 不存在根节点，说明是有循环
    // BFS检查树
    queue := []int{ root }
    treeMap[root]++
    for len(queue) != 0 {
        tmp := []int{}
        for _, n := range queue {
            // 被遍历到的点从原来的1变为0
            treeMap[n]--
            if leftChild[n] != -1 {
                tmp = append(tmp, leftChild[n])
            }
            if rightChild[n] != -1 {
                tmp = append(tmp, rightChild[n])
            }
        }
        queue = tmp
    }
    for n := range treeMap { // 检查是否还存在没遍历到的节点
        if n == -1 { continue }
        if treeMap[n] != 0 { return false }
    }
    return true
}

func validateBinaryTreeNodes1(n int, left []int, right []int) bool {
    in := make([]int, n)
    for i := 0; i < n; i++ {
        if left[i] != -1 {
            in[left[i]]++
        }
        if right[i] != -1 {
            in[right[i]]++
        }
    }
    count0, countOther,rootIndex := 0, 0, 0
    for i := 0; i < n; i++ {
        if in[i] == 0 {
            rootIndex = i
            count0++ 
        }
        if in[i] > 1 {
            countOther++
        }
    }
    visited := make([]bool, n)
    // 特别：一个独立的环 加上 一颗正常的树,就得正常判断一下
    var notCircle func(rootIdx int, left, right []int, visit []bool) int
    notCircle = func(rootIdx int, left, right []int, visit []bool) int {
        if rootIdx == -1 || visit[rootIdx] { return 0 }
        visit[rootIdx] = true
        l, r := notCircle(left[rootIdx], left, right, visit), notCircle(right[rootIdx], left, right, visit)
        return 1 + l + r
    }
    return count0 == 1 && countOther == 0 && notCircle(rootIndex, left, right, visited) == n
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/08/23/1503_ex1.png" />
    // Input: n = 4, leftChild = [1,-1,3,-1], rightChild = [2,-1,-1,-1]
    // Output: true
    fmt.Println(validateBinaryTreeNodes(4, []int{1,-1,3,-1},[]int{2,-1,-1,-1})) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2019/08/23/1503_ex2.png" />
    // Input: n = 4, leftChild = [1,-1,3,-1], rightChild = [2,3,-1,-1]
    // Output: false
    fmt.Println(validateBinaryTreeNodes(4, []int{1,-1,3,-1},[]int{2,3,-1,-1})) // false
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2019/08/23/1503_ex3.png" />
    // Input: n = 2, leftChild = [1,0], rightChild = [-1,-1]
    // Output: false
    fmt.Println(validateBinaryTreeNodes(2, []int{1,0},[]int{-1,-1})) // false

    fmt.Println(validateBinaryTreeNodes1(4, []int{1,-1,3,-1},[]int{2,-1,-1,-1})) // true
    fmt.Println(validateBinaryTreeNodes1(4, []int{1,-1,3,-1},[]int{2,3,-1,-1})) // false
    fmt.Println(validateBinaryTreeNodes1(2, []int{1,0},[]int{-1,-1})) // false
}