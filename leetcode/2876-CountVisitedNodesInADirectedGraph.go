package main

// 2876. Count Visited Nodes in a Directed Graph
// There is a directed graph consisting of n nodes numbered from 0 to n - 1 and n directed edges.

// You are given a 0-indexed array edges where edges[i] indicates 
// that there is an edge from node i to node edges[i].

// Consider the following process on the graph:
//     You start from a node x and keep visiting other nodes through edges 
//     until you reach a node that you have already visited before on this same process.

// Return an array answer where answer[i] is the number of different nodes 
// that you will visit if you perform the process starting from node i.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/08/31/graaphdrawio-1.png" />
// Input: edges = [1,2,0,0]
// Output: [3,3,3,4]
// Explanation: We perform the process starting from each node in the following way:
// - Starting from node 0, we visit the nodes 0 -> 1 -> 2 -> 0. The number of different nodes we visit is 3.
// - Starting from node 1, we visit the nodes 1 -> 2 -> 0 -> 1. The number of different nodes we visit is 3.
// - Starting from node 2, we visit the nodes 2 -> 0 -> 1 -> 2. The number of different nodes we visit is 3.
// - Starting from node 3, we visit the nodes 3 -> 0 -> 1 -> 2 -> 0. The number of different nodes we visit is 4.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/08/31/graaph2drawio.png" />
// Input: edges = [1,2,3,4,0]
// Output: [5,5,5,5,5]
// Explanation: Starting from any node we can visit every node in the graph in the process.

// Constraints:
//     n == edges.length
//     2 <= n <= 10^5
//     0 <= edges[i] <= n - 1
//     edges[i] != i

import "fmt"

func countVisitedNodes(edges []int) []int {
    type Node struct {
        index int
    }
    n := len(edges)
    res := make([]int, n)
    for i := 0; i < n; i++ {
        stack := []int{}
        nodes := make(map[int]*Node)
        j := i
        for nodes[j] == nil && res[j] == 0 { // while node has not been seen before
            nodes[j] = &Node{len(stack)}
            stack = append(stack, j)
            j = edges[j]
        }
        if nodes[j] != nil { // if node was seen that means theres a cycle
            start := nodes[j].index // all nodes from the first encounter of this node
            cycleLength := len(stack) - start // are in the cycle and have the cycle length
            for j := start; j < len(stack); j++ {
                res[stack[j]] = cycleLength
            }
            stack = stack[:start] // pop remove these cycle nodes from the stack
        }
        for j := len(stack) - 1; j >= 0; j-- { // the last node in the stack is connected to a cycle
            res[stack[j]] = res[edges[stack[j]]] + 1 // retrieve length until cycle of node it is directed to and add 1
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2023/08/31/graaphdrawio-1.png" />
    // Input: edges = [1,2,0,0]
    // Output: [3,3,3,4]
    // Explanation: We perform the process starting from each node in the following way:
    // - Starting from node 0, we visit the nodes 0 -> 1 -> 2 -> 0. The number of different nodes we visit is 3.
    // - Starting from node 1, we visit the nodes 1 -> 2 -> 0 -> 1. The number of different nodes we visit is 3.
    // - Starting from node 2, we visit the nodes 2 -> 0 -> 1 -> 2. The number of different nodes we visit is 3.
    // - Starting from node 3, we visit the nodes 3 -> 0 -> 1 -> 2 -> 0. The number of different nodes we visit is 4.
    fmt.Println(countVisitedNodes([]int{1,2,0,0})) // [3,3,3,4]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2023/08/31/graaph2drawio.png" />
    // Input: edges = [1,2,3,4,0]
    // Output: [5,5,5,5,5]
    // Explanation: Starting from any node we can visit every node in the graph in the process.
    fmt.Println(countVisitedNodes([]int{1,2,3,4,0})) // [5,5,5,5,5]
}