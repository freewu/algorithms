package main

import "fmt"

/**
77. Combinations
Given two integers n and k, return all possible combinations of k numbers out of the range [1, n].
You may return the answer in any order.

Constraints:

	1 <= n <= 20
	1 <= k <= n


Example 1:

	Input: n = 4, k = 2
	Output:
	[
	  [2,4],
	  [3,4],
	  [2,3],
	  [1,2],
	  [1,3],
	  [1,4],
	]

Example 2:

	Input: n = 1, k = 1
	Output: [[1]]

解题思路:
	给定两个整数 n 和 k，返回 1 … n 中所有可能的 k 个数的组合。
	DFS
 */

func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 || k > n {
		return [][]int{}
	}
	var res [][]int
	var c []int
	generateCombinations(n, k, 1, c, &res)
	return res
}

func generateCombinations(n, k, start int, c []int, res *[][]int) {
	// fmt.Printf("start =  %v\n",start)
	// fmt.Printf("c = %v\n",c)
	if len(c) == k { // 如果生成了符合要求长度的值 写入到结果集中
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	for i := start; i <= n-(k-len(c))+1; i++ {
		c = append(c, i)
		generateCombinations(n, k, i+1, c, res)
		fmt.Printf("before c = %v\n",c)
		c = c[:len(c)-1]
		fmt.Printf("after c = %v\n",c)
	}
	return
}

func main() {
	fmt.Printf("combine(4,2) = %v\n",combine(4,2))
	fmt.Printf("combine(1,1) = %v\n",combine(1,1))
}
