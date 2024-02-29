package main

// 2581. Count Number of Possible Root Nodes
// Alice has an undirected tree with n nodes labeled from 0 to n - 1. 
// The tree is represented as a 2D integer array edges of length n - 1 where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the tree.

// Alice wants Bob to find the root of the tree. She allows Bob to make several guesses about her tree. 
// In one guess, he does the following:

// Chooses two distinct integers u and v such that there exists an edge [u, v] in the tree.
// He tells Alice that u is the parent of v in the tree.
// Bob's guesses are represented by a 2D integer array guesses where guesses[j] = [uj, vj] indicates Bob guessed uj to be the parent of vj.

// Alice being lazy, does not reply to each of Bob's guesses, but just says that at least k of his guesses are true.

// Given the 2D integer arrays edges, guesses and the integer k, return the number of possible nodes that can be the root of Alice's tree. If there is no such tree, return 0.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/12/19/ex-1.png" />
// Input: edges = [[0,1],[1,2],[1,3],[4,2]], guesses = [[1,3],[0,1],[1,0],[2,4]], k = 3
// Output: 3
// Explanation: 
// Root = 0, correct guesses = [1,3], [0,1], [2,4]
// Root = 1, correct guesses = [1,3], [1,0], [2,4]
// Root = 2, correct guesses = [1,3], [1,0], [2,4]
// Root = 3, correct guesses = [1,0], [2,4]
// Root = 4, correct guesses = [1,3], [1,0]
// Considering 0, 1, or 2 as root node leads to 3 correct guesses.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/12/19/ex-2.png" />
// Input: edges = [[0,1],[1,2],[2,3],[3,4]], guesses = [[1,0],[3,4],[2,1],[3,2]], k = 1
// Output: 5
// Explanation: 
// Root = 0, correct guesses = [3,4]
// Root = 1, correct guesses = [1,0], [3,4]
// Root = 2, correct guesses = [1,0], [2,1], [3,4]
// Root = 3, correct guesses = [1,0], [2,1], [3,2], [3,4]
// Root = 4, correct guesses = [1,0], [2,1], [3,2]
// Considering any node as root will give at least 1 correct guess. 

// Constraints:
//         edges.length == n - 1
//         2 <= n <= 10^5
//         1 <= guesses.length <= 10^5
//         0 <= ai, bi, uj, vj <= n - 1
//         ai != bi
//         uj != vj
//         edges represents a valid tree.
//         guesses[j] is an edge of the tree.
//         guesses is unique.
//         0 <= k <= guesses.length

import "fmt"

func rootCount(edges [][]int, guesses [][]int, k int) int {
    type pair struct {
        first int
        second int
    }
    res := 0
    adjList := make([][]int, len(edges)+1)
    // Make the adjancency list
    for _, edge := range edges {
        adjList[edge[0]] = append(adjList[edge[0]], edge[1])
        adjList[edge[1]] = append(adjList[edge[1]], edge[0])
    }
    
    visited := map[int]bool{0 : true}
    guessSet := make(map[pair]bool)
    // Add the guesses to the guess set
    for _, guess := range guesses {
        guessSet[pair{first: guess[0], second: guess[1]}] = true
    }

    var findCount func (node int, adjList [][]int, visited map[int]bool, guesses map[pair]bool) int
    findCount = func (node int, adjList [][]int, visited map[int]bool, guesses map[pair]bool) int {
        var count int
        for _, nei := range adjList[node] {
            if !visited[nei] {
                visited[nei] = true
                if guesses[pair{first: node, second: nei}] {
                    count++
                }
                count += findCount(nei, adjList, visited, guesses)
            }
        }
        return count
    }
    // Assume 0 is the root, find how many guesses are correct.
    count := findCount(0, adjList, visited, guessSet)

    var traverse func (node, count, reverse, k int, adjList [][]int, visited map[int]bool, guesses map[pair]bool)
    traverse = func (node, count, reverse, k int, adjList [][]int, visited map[int]bool, guesses map[pair]bool) {
        for _, nei := range adjList[node] {
            if !visited[nei] {
                visited[nei] = true
                rev, c := reverse, count
                if guesses[pair{first: nei, second: node}] {
                    rev++
                }
                if guesses[pair{first: node, second: nei}] {
                    c--
                }
                traverse(nei, c, rev, k, adjList, visited, guesses)
            }
        }
        if count + reverse >= k {
            res++
        }
    }
    // Starting from 0 traverse the tree and check if the guess for reverted edge exists.
    visited = map[int]bool{0 : true}
    traverse(0, count, 0, k, adjList, visited, guessSet)

    return res
}

// best solution
func rootCount1(edges [][]int, guesses [][]int, k int) int {
    ans := 0
	g := make([][]int, len(edges)+1) // 建树
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v) // 建图
	}

	type pair struct{ x, y int }
	s := make(map[pair]int, len(guesses))
	for _, p := range guesses { // guesses 转成哈希表
		s[pair{p[0], p[1]}] = 1
	}

	cnt0 := 0
	var dfs func(int, int)
	dfs = func(x, fa int) {
		for _, y := range g[x] {
			if y != fa {
				if s[pair{x, y}] == 1 { // 以 0 为根时，猜对了
					cnt0++
				}
				dfs(y, x)
			}
		}
	}
	dfs(0, -1) // 以0 为根，猜对的总次数

	var reroot func(int, int, int)
	reroot = func(x, fa, cnt int) {
		if cnt >= k { // 此时 cnt 就是以 x 为根时的猜对次数
			ans++
		}
		for _, y := range g[x] { // 现在要变化
			if y != fa {
				reroot(y, x, cnt-s[pair{x, y}]+s[pair{y, x}])
			}
		}
	}
	reroot(0, -1, cnt0) // 从 0 开始
	return ans
}

func main() {
    // 根为节点 0 ，正确的猜测为 [1,3], [0,1], [2,4]
    // 根为节点 1 ，正确的猜测为 [1,3], [1,0], [2,4]
    // 根为节点 2 ，正确的猜测为 [1,3], [1,0], [2,4]
    // 根为节点 3 ，正确的猜测为 [1,0], [2,4]
    // 根为节点 4 ，正确的猜测为 [1,3], [1,0]
    // 节点 0 ，1 或 2 为根时，可以得到 3 个正确的猜测。
    fmt.Println(rootCount(
        [][]int{ []int{0,1},[]int{1,2},[]int{1,3},[]int{4,2}},
        [][]int{ []int{1,3},[]int{0,1},[]int{1,0},[]int{2,4}},
        3,
    )) // 3

    // 根为节点 0 ，正确的猜测为 [3,4]
    // 根为节点 1 ，正确的猜测为 [1,0], [3,4]
    // 根为节点 2 ，正确的猜测为 [1,0], [2,1], [3,4]
    // 根为节点 3 ，正确的猜测为 [1,0], [2,1], [3,2], [3,4]
    // 根为节点 4 ，正确的猜测为 [1,0], [2,1], [3,2]
    // 任何节点为根，都至少有 1 个正确的猜测。
    fmt.Println(rootCount(
        [][]int{ []int{0,1},[]int{1,2},[]int{2,3},[]int{3,4}},
        [][]int{ []int{1,0},[]int{3,4},[]int{2,1},[]int{3,2}},
        1,
    )) // 5

    fmt.Println(rootCount1(
        [][]int{ []int{0,1},[]int{1,2},[]int{1,3},[]int{4,2}},
        [][]int{ []int{1,3},[]int{0,1},[]int{1,0},[]int{2,4}},
        3,
    )) // 3

    fmt.Println(rootCount1(
        [][]int{ []int{0,1},[]int{1,2},[]int{2,3},[]int{3,4}},
        [][]int{ []int{1,0},[]int{3,4},[]int{2,1},[]int{3,2}},
        1,
    )) // 5
}

