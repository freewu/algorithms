package main

// 4100. Minimum Cost Path With Even XOR Parity
// You are given a tree with n nodes numbered from 0 to n - 1, represented by a 2D integer array edges.
// Each edges[i] = [u_{i}, v_{i}, c_{i}] represents an edge between nodes u_{i} and v_{i} with cost c_{i}.

// You are also given an integer array values of length n, where values[i] is the value of node i.

// A path between two distinct nodes is Balanced-XOR if the bitwise XOR of the values of all nodes on the path, including both endpoints, has an even number of set bits.

// The cost of a path is the sum of its edge costs.

// Return the minimum cost of a Balanced-XOR path. If no such path exists, return -1.

// The number 0 has zero set bits.

// Example 1:
// <img alt="" src="https://assets.leetcode.com/uploads/2026/08/31/4763e1ffdrawio.png" style="width: 300px; height: 300px;">
// Input: edges = [[0,1,4],[0,2,6],[0,3,3],[3,4,5]], values = [3,1,7,1,3]
// Output: 7
// Explanation:
//  Balanced-XOR Path XOR Value Total Cost
//  1 - 0 - 2 1 XOR 3 XOR 7 = 5 4 + 6 = 10
//  1 - 0 - 3 1 XOR 3 XOR 1 = 3 4 + 3 = 7
//  1 - 0 - 3 - 4 1 XOR 3 XOR 1 XOR 3 = 0 4 + 3 + 5 = 12
//  2 - 0 - 3 7 XOR 3 XOR 1 = 5 6 + 3 = 9
//  2 - 0 - 3 - 4 7 XOR 3 XOR 1 XOR 3 = 6 6 + 3 + 5 = 14
// The minimum cost is 7, achieved by the path 1 - 0 - 3. Thus, the answer is 7.

// Example 2:
// <img alt="" src="https://assets.leetcode.com/uploads/2026/08/31/4763e2drawio.png" style="width: 300px; height: 132px;">
// Input: edges = [[0,1,5],[1,2,10]], values = [1,7,3]
// Output: 5
// Explanation:
//  Balanced-XOR Path XOR Value Total Cost
//  0 - 1 1 XOR 7 = 6 5
//  0 - 1 - 2 1 XOR 7 XOR 3 = 5 5 + 10 = 15
// The minimum cost is 5, achieved by the path 0 - 1. Thus, the answer is 5.

// Example 3:
// <img alt="" src="https://assets.leetcode.com/uploads/2026/08/31/4743e3drawio.png" style="width: 300px; height: 215px;">
// Input: edges = [[0,1,5],[0,2,7],[0,3,9]], values = [1,3,5,9]
// Output: -1
// Explanation:
// - Every path between two distinct nodes has an XOR value with an odd number of set bits.
// - Therefore, no Balanced-XOR path exists, so the answer is -1.

// Constraints:
//     - 2 <= n <= 10^{5}
//     - edges.length == n - 1
//     - edges[i] = [u_{i}, v_{i}, c_{i}]
//     - 0 <= u_{i}, v_{i} <= n - 1
//     - 1 <= c_{i} <= 10^{9}
//     - values.length == n
//     - 1 <= values[i] <= 10^{9}
//     - edges represents a valid tree.

import (
	"fmt"
	"math/bits"
)

func minActivationCost(edges [][]int, values []int) int {
	const inf = 1 << 61
	res := inf
	arr := make([]int, len(values))
	for i := range arr {
		arr[i] = inf
	}
	for _, e := range edges {
		x, y, c := e[0], e[1], e[2]
		vx := bits.OnesCount(uint(values[x])) & 1
		vy := bits.OnesCount(uint(values[y])) & 1
		if vx == vy {
			res = min(res, c)
		} else if vx == 0 { // 此时 vy = 1
			res = min(res, arr[x]+c)
			arr[x] = min(arr[x], c)
		} else { // 此时 vy = 0 且 vx = 1
			res = min(res, arr[y]+c)
			arr[y] = min(arr[y], c)
		}
	}
	if res < inf {
		return res
	}
	return -1
}

func main() {
	// Example 1:
	// <img alt="" src="https://assets.leetcode.com/uploads/2026/08/31/4763e1ffdrawio.png" style="width: 300px; height: 300px;">
	// Input: edges = [[0,1,4],[0,2,6],[0,3,3],[3,4,5]], values = [3,1,7,1,3]
	// Output: 7
	// Explanation:
	//  Balanced-XOR Path XOR Value Total Cost
	//  1 - 0 - 2 1 XOR 3 XOR 7 = 5 4 + 6 = 10
	//  1 - 0 - 3 1 XOR 3 XOR 1 = 3 4 + 3 = 7
	//  1 - 0 - 3 - 4 1 XOR 3 XOR 1 XOR 3 = 0 4 + 3 + 5 = 12
	//  2 - 0 - 3 7 XOR 3 XOR 1 = 5 6 + 3 = 9
	//  2 - 0 - 3 - 4 7 XOR 3 XOR 1 XOR 3 = 6 6 + 3 + 5 = 14
	// The minimum cost is 7, achieved by the path 1 - 0 - 3. Thus, the answer is 7.
	fmt.Println(minActivationCost([][]int{{0, 1, 4}, {0, 2, 6}, {0, 3, 3}, {3, 4, 5}}, []int{3, 1, 7, 1, 3})) // 7
	// Example 2:
	// <img alt="" src="https://assets.leetcode.com/uploads/2026/08/31/4763e2drawio.png" style="width: 300px; height: 132px;">
	// Input: edges = [[0,1,5],[1,2,10]], values = [1,7,3]
	// Output: 5
	// Explanation:
	//  Balanced-XOR Path XOR Value Total Cost
	//  0 - 1 1 XOR 7 = 6 5
	//  0 - 1 - 2 1 XOR 7 XOR 3 = 5 5 + 10 = 15
	// The minimum cost is 5, achieved by the path 0 - 1. Thus, the answer is 5.
	fmt.Println(minActivationCost([][]int{{0, 1, 5}, {1, 2, 10}}, []int{1, 7, 3})) // 5
	// Example 3:
	// <img alt="" src="https://assets.leetcode.com/uploads/2026/08/31/4743e3drawio.png" style="width: 300px; height: 215px;">
	// Input: edges = [[0,1,5],[0,2,7],[0,3,9]], values = [1,3,5,9]
	// Output: -1
	// Explanation:
	// - Every path between two distinct nodes has an XOR value with an odd number of set bits.
	// - Therefore, no Balanced-XOR path exists, so the answer is -1.
	fmt.Println(minActivationCost([][]int{{0, 1, 5}, {0, 2, 7}, {0, 3, 9}}, []int{1, 3, 5, 9})) // -1
}
