package main

// 1791. Find Center of Star Graph
// There is an undirected star graph consisting of n nodes labeled from 1 to n. 
// A star graph is a graph where there is one center node and exactly n - 1 edges that connect the center node with every other node.

// You are given a 2D integer array edges where each edges[i] = [ui, vi] indicates that there is an edge between the nodes ui and vi. 
// Return the center of the given star graph.

// Example 1:
// Input: edges = [[1,2],[2,3],[4,2]]
// Output: 2
// Explanation: As shown in the figure above, node 2 is connected to every other node, so 2 is the center.

// Example 2:
// Input: edges = [[1,2],[5,1],[1,3],[1,4]]
// Output: 1

// Constraints:
//     3 <= n <= 10^5
//     edges.length == n - 1
//     edges[i].length == 2
//     1 <= ui, vi <= n
//     ui != vi
//     The given edges represent a valid star graph.

import "fmt"

func findCenter(edges [][]int) int {
    n, mp := len(edges), make(map[int]int)
    for i := range edges {
        mp[edges[i][0]]++
        mp[edges[i][1]]++
    }
    for i, c := range mp {
        if c == n {
            return i
        }
    }
    return -1
}

func findCenter1(edges [][]int) int {
    if edges[1][0] == edges[0][0] || edges[1][0] == edges[0][1] {
        return edges[1][0]
    }
    return edges[1][1]
}

func main() {
    // Example 1:
    // Input: edges = [[1,2],[2,3],[4,2]]
    // Output: 2
    // Explanation: As shown in the figure above, node 2 is connected to every other node, so 2 is the center.
    fmt.Println(findCenter([][]int{{1,2},{2,3},{4,2}})) // 2
    // Example 2:
    // Input: edges = [[1,2],[5,1],[1,3],[1,4]]
    // Output: 1
    fmt.Println(findCenter([][]int{{1,2},{5,1},{1,3},{1,4}})) // 1

    fmt.Println(findCenter1([][]int{{1,2},{2,3},{4,2}})) // 2
    fmt.Println(findCenter1([][]int{{1,2},{5,1},{1,3},{1,4}})) // 1
}