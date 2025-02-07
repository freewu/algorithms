package main

// 3378. Count Connected Components in LCM Graph
// You are given an array of integers nums of size n and a positive integer threshold.

// There is a graph consisting of n nodes with the ith node having a value of nums[i]. 
// Two nodes i and j in the graph are connected via an undirected edge if lcm(nums[i], nums[j]) <= threshold.

// Return the number of connected components in this graph.

// A connected component is a subgraph of a graph in which there exists a path between any two vertices, 
// and no vertex of the subgraph shares an edge with a vertex outside of the subgraph.

// The term lcm(a, b) denotes the least common multiple of a and b.

// Example 1:
// Input: nums = [2,4,8,3,9], threshold = 5
// Output: 4
// Explanation: 
// <img src="https://assets.leetcode.com/uploads/2024/10/31/example0.png" />
// The four connected components are (2, 4), (3), (8), (9).

// Example 2:
// Input: nums = [2,4,8,3,9,12], threshold = 10
// Output: 2
// Explanation: 
// <img src="https://assets.leetcode.com/uploads/2024/10/31/example1.png" />
// The two connected components are (2, 3, 4, 8, 9), and (12).

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     All elements of nums are unique.
//     1 <= threshold <= 2 * 10^5

import "fmt"

type DSU struct {
    parent map[int]int
    rank   map[int]int
}

func NewDSU(n int) *DSU {
    dsu := &DSU{ make(map[int]int), make(map[int]int) }
    for i := 0; i <= n; i++ {
        dsu.parent[i] = i
        dsu.rank[i] = 0
    }
    return dsu
}

func (dsu *DSU) Find(x int) int {
    if dsu.parent[x] != x {
        dsu.parent[x] = dsu.Find(dsu.parent[x])
    }
    return dsu.parent[x]
}

func (dsu *DSU) Union(u, v int) {
    uRoot, vRoot := dsu.Find(u), dsu.Find(v)
    if uRoot != vRoot {
        if dsu.rank[uRoot] < dsu.rank[vRoot] {
            uRoot, vRoot = vRoot, uRoot
        }
        dsu.parent[vRoot] = uRoot
        if dsu.rank[uRoot] == dsu.rank[vRoot] {
            dsu.rank[uRoot]++
        }
    }
}

func countComponents(nums []int, threshold int) int {
    dsu := NewDSU(threshold)
    for _, num := range nums {
        for j := num; j <= threshold; j += num {
            dsu.Union(num, j)
        }
    }
    uniqueParents := make(map[int]struct{})
    for _, num := range nums {
        if num > threshold {
            uniqueParents[num] = struct{}{}
        } else {
            uniqueParents[dsu.Find(num)] = struct{}{}
        }
    }
    return len(uniqueParents)
}

func countComponents1(nums []int, threshold int) int {
    n := len(nums)
    fa := make([]int, n)
    for i := range fa {
        fa[i] = i
    }
    // 非递归并查集
    find := func(x int) int {
        res := x
        for fa[res] != res {
            res = fa[res]
        }
        for fa[x] != res {
            fa[x], x = res, fa[x]
        }
        return res
    }
    // 记录每个数的下标
    indexs := make([]int, threshold + 1)
    for i, v := range nums {
        if v <= threshold {
            indexs[v] = i + 1 // 这里 +1 了，下面减掉
        }
    }
    for i := 1; i <= threshold; i++ {
        mn := -1
        for j := i; j <= threshold; j +=  i {
            if indexs[j] > 0 { // indexs[j] == 0 表示不存在
                mn = j
                break
            }
        }
        if mn < 0 { continue }
        fi := find(indexs[mn] - 1)
        for j := mn + i; j <= threshold && j <= i * threshold / mn; j += i {
            if indexs[j] > 0 {
                fj := find(indexs[j] - 1)
                if fj != fi {
                    fa[fj] = fi // 合并 indexs[i] 和 indexs[j]
                    n--
                }
            }
        }
    }
    return n
}

func main() {
    // Example 1:
    // Input: nums = [2,4,8,3,9], threshold = 5
    // Output: 4
    // Explanation: 
    // <img src="https://assets.leetcode.com/uploads/2024/10/31/example0.png" />
    // The four connected components are (2, 4), (3), (8), (9).
    fmt.Println(countComponents([]int{2,4,8,3,9}, 5)) // 4
    // Example 2:
    // Input: nums = [2,4,8,3,9,12], threshold = 10
    // Output: 2
    // Explanation: 
    // <img src="https://assets.leetcode.com/uploads/2024/10/31/example1.png" />
    // The two connected components are (2, 3, 4, 8, 9), and (12).
    fmt.Println(countComponents([]int{2,4,8,3,9,12}, 10)) // 2

    fmt.Println(countComponents([]int{1,2,3,4,5,6,7,8,9}, 10)) // 1
    fmt.Println(countComponents([]int{9,8,7,6,5,4,3,2,1}, 10)) // 1

    fmt.Println(countComponents1([]int{2,4,8,3,9}, 5)) // 4
    fmt.Println(countComponents1([]int{2,4,8,3,9,12}, 10)) // 2
    fmt.Println(countComponents1([]int{1,2,3,4,5,6,7,8,9}, 10)) // 1
    fmt.Println(countComponents1([]int{9,8,7,6,5,4,3,2,1}, 10)) // 1
}