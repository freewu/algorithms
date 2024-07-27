package main

// 2445. Number of Nodes With Value One
// There is an undirected connected tree with n nodes labeled from 1 to n and n - 1 edges. 
// You are given the integer n. 
// The parent node of a node with a label v is the node with the label floor (v / 2). 
// The root of the tree is the node with the label 1.

//     For example, if n = 7, then the node with the label 3 has the node with the label floor(3 / 2) = 1 as its parent, 
//     and the node with the label 7 has the node with the label floor(7 / 2) = 3 as its parent.

// You are also given an integer array queries. Initially, every node has a value 0 on it. 
// For each query queries[i], you should flip all values in the subtree of the node with the label queries[i].

// Return the total number of nodes with the value 1 after processing all the queries.

// Note that:
//     Flipping the value of a node means that the node with the value 0 becomes 1 and vice versa.
//     floor(x) is equivalent to rounding x down to the nearest integer.

// Example 1:
// <img src="" />
// Input: n = 5 , queries = [1,2,5]
// Output: 3
// Explanation: The diagram above shows the tree structure and its status after performing the queries. The blue node represents the value 0, and the red node represents the value 1.
// After processing the queries, there are three red nodes (nodes with value 1): 1, 3, and 5.

// Example 2:
// <img src="" />
// Input: n = 3, queries = [2,3,3]
// Output: 1
// Explanation: The diagram above shows the tree structure and its status after performing the queries. The blue node represents the value 0, and the red node represents the value 1.
// After processing the queries, there are one red node (node with value 1): 2.

// Constraints:
//     1 <= n <= 10^5
//     1 <= queries.length <= 10^5
//     1 <= queries[i] <= n

import "fmt"

func numberOfNodes(n int, queries []int) int {
    res, nodes := 0, make([]bool, n+1)
    for i := range queries{
        nodes[queries[i]] = !nodes[queries[i]]
    }
    for i := range nodes {
        if i == 0 {
            continue
        }
        if i == 1 {
            if nodes[i] {
                res++
            }
            continue
        }
        if (!nodes[i/2] && nodes[i]) || (nodes[i/2] && !nodes[i]) { // 更新状态
            nodes[i] = true
            res++
        } else{
            nodes[i] = false
        }
    }
    return res
}

func numberOfNodes1(n int, queries []int) int {
    tree := make([]int, n+1)
    for _, v := range queries {
        tree[v] = (1 - tree[v])
    }
    var f func(int,int) int
    f = func(now int,state int) int{
        if now> n {
            return 0
        }
        if state == 1 {
            return (1 - tree[now]) + f(now * 2,1 - tree[now]) + f(now * 2 + 1,1 - tree[now])
        } else {
            return tree[now] + f(now * 2,tree[now]) + f(now * 2 + 1, tree[now])
        }
    }
    return f(1, 0)
}

func main() {
    // Example 1:
    // <img src="" />
    // Input: n = 5 , queries = [1,2,5]
    // Output: 3
    // Explanation: The diagram above shows the tree structure and its status after performing the queries. The blue node represents the value 0, and the red node represents the value 1.
    // After processing the queries, there are three red nodes (nodes with value 1): 1, 3, and 5.
    fmt.Println(numberOfNodes(5,[]int{1,2,5})) // 3
    // Example 2:
    // <img src="" />
    // Input: n = 3, queries = [2,3,3]
    // Output: 1
    // Explanation: The diagram above shows the tree structure and its status after performing the queries. The blue node represents the value 0, and the red node represents the value 1.
    // After processing the queries, there are one red node (node with value 1): 2.
    fmt.Println(numberOfNodes(3,[]int{2,3,3})) // 1

    fmt.Println(numberOfNodes1(5,[]int{1,2,5})) // 3
    fmt.Println(numberOfNodes1(3,[]int{2,3,3})) // 1
}