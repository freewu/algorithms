package main

// 886. Possible Bipartition
// We want to split a group of n people (labeled from 1 to n) into two groups of any size. 
// Each person may dislike some other people, and they should not go into the same group.

// Given the integer n and the array dislikes where dislikes[i] = [ai, bi] indicates 
// that the person labeled ai does not like the person labeled bi, 
// return true if it is possible to split everyone into two groups in this way.

// Example 1:
// Input: n = 4, dislikes = [[1,2],[1,3],[2,4]]
// Output: true
// Explanation: The first group has [1,4], and the second group has [2,3].

// Example 2:
// Input: n = 3, dislikes = [[1,2],[1,3],[2,3]]
// Output: false
// Explanation: We need at least 3 groups to divide them. We cannot put them in two groups.

// Constraints:
//     1 <= n <= 2000
//     0 <= dislikes.length <= 10^4
//     dislikes[i].length == 2
//     1 <= ai < bi <= n
//     All the pairs of dislikes are unique.

import "fmt"

// bfs
func possibleBipartition(n int, dislikes [][]int) bool {
    edges := make([][]int, n + 1)
    for _, dislike := range dislikes {
        edges[dislike[0]] = append(edges[dislike[0]], dislike[1])
        edges[dislike[1]] = append(edges[dislike[1]], dislike[0]) // 因为要标记，所以需要一次遍历完有关系的人，所以要逆关系
    }
    colors := make([]int, n + 1)
    bfs := func (start int, edges [][]int, colors []int) bool {
        colors[start] = 1 // 标记
        for queue := []int{start}; len(queue) > 0; queue = queue[1:] {
            node := queue[0]
            for _, neighbor := range edges[node] {
                if colors[neighbor] == colors[node] {
                    return false
                }
                if colors[neighbor] == 0 {
                    colors[neighbor] = -colors[node]
                    queue = append(queue, neighbor)
                }
            }
        }
        return true
    }
    for i := 1; i <= n; i++ {
        if colors[i] == 0 && !bfs(i, edges, colors) { //
            return false
        }
    }
    return true
}

// 并查集
func possibleBipartition1(n int, dislikes [][]int) bool {
    p := make([]int, ( n + 1) << 1)
    for i := range p {
        p[i] = i
    }
    var find func(int) int
    find = func(i int) int {
        if p[i] != i {
            p[i] = find(p[i])
        }
        return p[i]
    }
    union := func(i, j int) {
        p[find(i)] = p[find(j)]
    }
    for _, d := range dislikes {
        i, j := d[0], d[1]
        if find(i) == find(j) {
            return false
        }
        union(i, j+n)
        union(j, i+n)
    }
    return true
}

func main() {
    // Example 1:
    // Input: n = 4, dislikes = [[1,2],[1,3],[2,4]]
    // Output: true
    // Explanation: The first group has [1,4], and the second group has [2,3].
    fmt.Println(possibleBipartition(4, [][]int{{1,2},{1,3},{2,4}})) // true
    // Example 2:
    // Input: n = 3, dislikes = [[1,2],[1,3],[2,3]]
    // Output: false
    // Explanation: We need at least 3 groups to divide them. We cannot put them in two groups.
    fmt.Println(possibleBipartition(3, [][]int{{1,2},{1,3},{2,3}})) // false

    fmt.Println(possibleBipartition1(4, [][]int{{1,2},{1,3},{2,4}})) // true
    fmt.Println(possibleBipartition1(3, [][]int{{1,2},{1,3},{2,3}})) // false
}