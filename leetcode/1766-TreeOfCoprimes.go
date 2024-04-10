package main

// 1766. Tree of Coprimes
// There is a tree (i.e., a connected, undirected graph that has no cycles) consisting of n nodes numbered from 0 to n - 1 and exactly n - 1 edges. 
// Each node has a value associated with it, and the root of the tree is node 0.

// To represent this tree, you are given an integer array nums and a 2D array edges.
// Each nums[i] represents the ith node's value, and each edges[j] = [uj, vj] represents an edge between nodes uj and vj in the tree.

// Two values x and y are coprime if gcd(x, y) == 1 where gcd(x, y) is the greatest common divisor of x and y.
// An ancestor of a node i is any other node on the shortest path from node i to the root. A node is not considered an ancestor of itself.
// Return an array ans of size n, where ans[i] is the closest ancestor to node i such that nums[i] and nums[ans[i]] are coprime, or -1 if there is no such ancestor.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/01/06/untitled-diagram.png" />
// Input: nums = [2,3,3,2], edges = [[0,1],[1,2],[1,3]]
// Output: [-1,0,0,1]
// Explanation: In the above figure, each node's value is in parentheses.
// - Node 0 has no coprime ancestors.
// - Node 1 has only one ancestor, node 0. Their values are coprime (gcd(2,3) == 1).
// - Node 2 has two ancestors, nodes 1 and 0. Node 1's value is not coprime (gcd(3,3) == 3), but node 0's
//   value is (gcd(2,3) == 1), so node 0 is the closest valid ancestor.
// - Node 3 has two ancestors, nodes 1 and 0. It is coprime with node 1 (gcd(3,2) == 1), so node 1 is its
//   closest valid ancestor.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/01/06/untitled-diagram1.png" />
// Input: nums = [5,6,10,2,3,6,15], edges = [[0,1],[0,2],[1,3],[1,4],[2,5],[2,6]]
// Output: [-1,0,-1,0,0,0,-1]
 
// Constraints:
//     nums.length == n
//     1 <= nums[i] <= 50
//     1 <= n <= 10^5
//     edges.length == n - 1
//     edges[j].length == 2
//     0 <= uj, vj < n
//     uj != vj
import "fmt"

// dfs
func getCoprimes(nums []int, edges [][]int) []int {
    n := len(nums)
    // nums[i] 的范围是 [1,50]，预处理 gcds[j] 表示 [1,50] 中与 j 互质的元素的集合。
    //用 tmp[i] 表示在搜索过程中 i=nums[x] 的节点坐标集合，显然该集合的末尾元素离当前节点最近。
    gcds, tmp := make([][]int, 51), make([][]int, 51)
    res,dep, g := make([]int, n), make([]int, n), make([][]int, n)
    // 初始化
    for i := 0; i < 51; i++ {
        gcds[i] = []int{}
        tmp[i] = []int{}
    }
    for i := 0; i < n; i++ {
        g[i] = []int{}
        res[i], dep[i] = -1, -1
    }
    gcd := func (a, b int) int { for b != 0 { a, b = b, a % b; }; return a; }
    var dfs func(x, depth int)
    dfs = func(x, depth int) {
        dep[x] = depth
        for _, val := range gcds[nums[x]] {
            if len(tmp[val]) == 0 {
                continue
            }
            las := tmp[val][len(tmp[val]) - 1]
            if res[x] == -1 || dep[las] > dep[res[x]] {
                res[x] = las
            }
        }
        tmp[nums[x]] = append(tmp[nums[x]], x)
        for _, val := range g[x] {
            if dep[val] == -1 { // 被访问过的点dep不为-1
                dfs(val, depth + 1)
            }
        }
        tmp[nums[x]] = tmp[nums[x]][:len(tmp[nums[x]]) - 1]
    }

    for i := 1; i <= 50; i++ {
        for j := 1; j <= 50; j++ {
            if gcd(i, j) == 1 {
                gcds[i] = append(gcds[i], j)
            }
        }
    }
    for _, edge := range edges {
        x := edge[0]
        y := edge[1]
        g[x] = append(g[x], y)
        g[y] = append(g[y], x)
    }
    dfs(0, 1)
    return res
}



func main() {
    // Explanation: In the above figure, each node's value is in parentheses.
    // - Node 0 has no coprime ancestors.
    // - Node 1 has only one ancestor, node 0. Their values are coprime (gcd(2,3) == 1).
    // - Node 2 has two ancestors, nodes 1 and 0. Node 1's value is not coprime (gcd(3,3) == 3), but node 0's
    //   value is (gcd(2,3) == 1), so node 0 is the closest valid ancestor.
    // - Node 3 has two ancestors, nodes 1 and 0. It is coprime with node 1 (gcd(3,2) == 1), so node 1 is its
    //   closest valid ancestor.
    fmt.Println(getCoprimes([]int{2,3,3,2},[][]int{{0,1},{1,2},{1,3}})) // [-1,0,0,1]
    fmt.Println(getCoprimes([]int{5,6,10,2,3,6,15},[][]int{{0,1},{0,2},{1,3},{1,4},{2,5},{2,6}})) // [-1,0,-1,0,0,0,-1]
}