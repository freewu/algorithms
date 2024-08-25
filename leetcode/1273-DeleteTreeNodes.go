package main

// 1273. Delete Tree Nodes
// A tree rooted at node 0 is given as follows:
//     The number of nodes is nodes;
//     The value of the ith node is value[i];
//     The parent of the ith node is parent[i].

// Remove every subtree whose sum of values of nodes is zero.
// Return the number of the remaining nodes in the tree.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/07/02/1421_sample_1.PNG" />
// Input: nodes = 7, parent = [-1,0,0,1,2,2,2], value = [1,-2,4,0,-2,-1,-1]
// Output: 2

// Example 2:
// Input: nodes = 7, parent = [-1,0,0,1,2,2,2], value = [1,-2,4,0,-2,-1,-2]
// Output: 6

// Constraints:
//     1 <= nodes <= 10^4
//     parent.length == nodes
//     0 <= parent[i] <= nodes - 1
//     parent[0] == -1 which indicates that 0 is the root.
//     value.length == nodes
//     -10^5 <= value[i] <= 10^5
//     The given input is guaranteed to represent a valid tree.

import "fmt"

func deleteTreeNodes(nodes int, parent []int, value []int) int {
    g := make(map[int][]int)
    for i, p := range parent {
        if p != -1 {
            g[p] = append(g[p], i)
        }
    }
    counter := make([]int, nodes)
    for i:= range counter {
        counter[i] =1 
    }
    var dfs func(n int)
    dfs = func(n int) {
        if vs, ok := g[n]; ok {
            for _, v := range vs {
                dfs(v)
                value[n] += value[v]
                counter[n] += counter[v]
            }
        }
        if value[n] == 0 {
            counter[n] = 0
        }
    }
    dfs(0)
    return counter[0]
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/07/02/1421_sample_1.PNG" />
    // Input: nodes = 7, parent = [-1,0,0,1,2,2,2], value = [1,-2,4,0,-2,-1,-1]
    // Output: 2
    fmt.Println(deleteTreeNodes(7,[]int{-1,0,0,1,2,2,2}, []int{1,-2,4,0,-2,-1,-1})) // 2
    // Example 2:
    // Input: nodes = 7, parent = [-1,0,0,1,2,2,2], value = [1,-2,4,0,-2,-1,-2]
    // Output: 6
    fmt.Println(deleteTreeNodes(7,[]int{-1,0,0,1,2,2,2}, []int{1,-2,4,0,-2,-1,-2})) // 6
}