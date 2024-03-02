package main

// 2368. Reachable Nodes With Restrictions  
// There is an undirected tree with n nodes labeled from 0 to n - 1 and n - 1 edges.
// You are given a 2D integer array edges of length n - 1 where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the tree. 
// You are also given an integer array restricted which represents restricted nodes.
// Return the maximum number of nodes you can reach from node 0 without visiting a restricted node.
// Note that node 0 will not be a restricted node.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/06/15/ex1drawio.png" />
// Input: n = 7, edges = [[0,1],[1,2],[3,1],[4,0],[0,5],[5,6]], restricted = [4,5]
// Output: 4
// Explanation: The diagram above shows the tree.
// We have that [0,1,2,3] are the only nodes that can be reached from node 0 without visiting a restricted node.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/06/15/ex2drawio.png" />
// Input: n = 7, edges = [[0,1],[0,2],[0,5],[0,4],[3,2],[6,5]], restricted = [4,2,1]
// Output: 3
// Explanation: The diagram above shows the tree.
// We have that [0,5,6] are the only nodes that can be reached from node 0 without visiting a restricted node.
 
// Constraints:
//         2 <= n <= 10^5
//         edges.length == n - 1
//         edges[i].length == 2
//         0 <= ai, bi < n
//         ai != bi
//         edges represents a valid tree.
//         1 <= restricted.length < n
//         1 <= restricted[i] < n
//         All the values of restricted are unique.

import "fmt"

func reachableNodes(n int, edges [][]int, restricted []int) int {
    adjList := make([][]int, n)
    for _, v := range edges {
        adjList[v[0]] = append(adjList[v[0]], v[1])
        adjList[v[1]] = append(adjList[v[1]], v[0])
    }
    //fmt.Println(adjList)
    
    seen := make(map[int]bool)
    seen[0] = true
    for _, v := range restricted {
        seen[v] = true
    }
    stack := make([]int, 0)
    stack = append(stack, adjList[0]...)
    
    ans := 1

    //fmt.Println("seen: ", seen)
    
    var popLeft int
    for len(stack) > 0 {
        stack, popLeft = stack[1:], stack[0]
        if seen[popLeft] {
            continue
        }
        seen[popLeft] = true
        ans+=1
        stack = append(stack, adjList[popLeft]...)
        //fmt.Println("stack: ", stack)
    }
    return ans
}

// 使用并查集
func reachableNodes1(n int, edges [][]int, restricted []int) int {
	restrictedMap := make(map[int]bool)
	for _, r := range restricted {
		restrictedMap[r] = true
	}
	uf := newUnionFind(n)
	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		if restrictedMap[u] || restrictedMap[v] {
			continue
		}
		uf.union(u, v)
	}
	return uf.size[0]
}

type unionFind struct {
	parent []int
	size   []int
	count  int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	count := n
	return &unionFind{
		parent: parent,
		size:   size,
		count:  count,
	}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) bool {
	xRoot := uf.find(x)
	yRoot := uf.find(y)
	if xRoot == yRoot {
		return false
	}
	if xRoot > yRoot {
		xRoot, yRoot = yRoot, xRoot
	}
	uf.parent[yRoot] = xRoot
	uf.size[xRoot] += uf.size[yRoot]
	uf.count--
	return true
}

func main() {
    fmt.Println(reachableNodes(
        7,
        [][]int{[]int{0,1},[]int{1,2},[]int{3,1},[]int{4,0},[]int{0,5},[]int{5,6}},
        []int{ 4, 5},
    )) // 4

    fmt.Println(reachableNodes(
        7,
        [][]int{[]int{0,1},[]int{0,2},[]int{0,5},[]int{0,4},[]int{3,2},[]int{6,5}},
        []int{ 4,2,1 },
    )) // 3

    fmt.Println(reachableNodes1(
        7,
        [][]int{[]int{0,1},[]int{1,2},[]int{3,1},[]int{4,0},[]int{0,5},[]int{5,6}},
        []int{ 4, 5},
    )) // 4

    fmt.Println(reachableNodes1(
        7,
        [][]int{[]int{0,1},[]int{0,2},[]int{0,5},[]int{0,4},[]int{3,2},[]int{6,5}},
        []int{ 4,2,1 },
    )) // 3
}