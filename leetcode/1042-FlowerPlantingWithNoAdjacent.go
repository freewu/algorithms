package main

// 1042. Flower Planting With No Adjacent
// You have n gardens, labeled from 1 to n, 
// and an array paths where paths[i] = [xi, yi] describes a bidirectional path between garden xi to garden yi. 
// In each garden, you want to plant one of 4 types of flowers.

// All gardens have at most 3 paths coming into or leaving it.

// Your task is to choose a flower type for each garden such that, 
// for any two gardens connected by a path, they have different types of flowers.

// Return any such a choice as an array answer, where answer[i] is the type of flower planted in the (i+1)th garden. 
// The flower types are denoted 1, 2, 3, or 4. It is guaranteed an answer exists.

// Example 1:
// Input: n = 3, paths = [[1,2],[2,3],[3,1]]
// Output: [1,2,3]
// Explanation:
// Gardens 1 and 2 have different types.
// Gardens 2 and 3 have different types.
// Gardens 3 and 1 have different types.
// Hence, [1,2,3] is a valid answer. Other valid answers include [1,2,4], [1,4,2], and [3,2,1].

// Example 2:
// Input: n = 4, paths = [[1,2],[3,4]]
// Output: [1,2,1,2]

// Example 3:
// Input: n = 4, paths = [[1,2],[2,3],[3,4],[4,1],[1,3],[2,4]]
// Output: [1,2,3,4]

// Constraints:
//     1 <= n <= 10^4
//     0 <= paths.length <= 2 * 10^4
//     paths[i].length == 2
//     1 <= xi, yi <= n
//     xi != yi
//     Every garden has at most 3 paths coming into or leaving it.

import "fmt"

// bfs
func gardenNoAdj(n int, paths [][]int) []int {
    res, edges := make([]int, n), map[int][]int{}
    for i := range paths {
        edges[paths[i][0]] = append(edges[paths[i][0]], paths[i][1])
        edges[paths[i][1]] = append(edges[paths[i][1]], paths[i][0])
    }
    for i := 1; i <= n; i++ {
        color := [4]bool{}
        for _, v := range edges[i] {
            if res[v-1] != 0 {
                color[ res[v-1] - 1 ] = true
            }
        }
        switch {
            case !color[0]: res[i-1] = 1
            case !color[1]: res[i-1] = 2
            case !color[2]: res[i-1] = 3
            default: res[i-1] = 4
        }
    }
    return res
}

func gardenNoAdj1(n int, paths [][]int) []int {
    mp := make([][]int, n)
    for i := 0; i < n; i++ {
        mp[i] = []int{}
    }
    for _, v := range paths {
        mp[v[0]-1] = append(mp[v[0]-1], v[1]-1)
        mp[v[1]-1] = append(mp[v[1]-1], v[0]-1)
    }
    res := make([]int, n)
    for i, v := range mp {
        t := make([]bool, 5)
        for j := 0; j < len(v); j++ {
            if res[v[j]] != 0 {
                t[res[v[j]]] = true
            }
        }
        for k := 1; k < 5; k++ {
            if !t[k] {
                res[i] = k
                break
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 3, paths = [[1,2],[2,3],[3,1]]
    // Output: [1,2,3]
    // Explanation:
    // Gardens 1 and 2 have different types.
    // Gardens 2 and 3 have different types.
    // Gardens 3 and 1 have different types.
    // Hence, [1,2,3] is a valid answer. Other valid answers include [1,2,4], [1,4,2], and [3,2,1].
    fmt.Println(gardenNoAdj(3, [][]int{{1,2},{2,3},{3,1}})) // [1,2,3]
    // Example 2:
    // Input: n = 4, paths = [[1,2],[3,4]]
    // Output: [1,2,1,2]
    fmt.Println(gardenNoAdj(4, [][]int{{1,2},{3,4}})) // [1,2,1,2]
    // Example 3:
    // Input: n = 4, paths = [[1,2],[2,3],[3,4],[4,1],[1,3],[2,4]]
    // Output: [1,2,3,4]
    fmt.Println(gardenNoAdj(4, [][]int{{1,2},{2,3},{3,4},{4,1},{1,3},{2,4}})) // [1,2,3,4]

    fmt.Println(gardenNoAdj1(3, [][]int{{1,2},{2,3},{3,1}})) // [1,2,3]
    fmt.Println(gardenNoAdj1(4, [][]int{{1,2},{3,4}})) // [1,2,1,2]
    fmt.Println(gardenNoAdj1(4, [][]int{{1,2},{2,3},{3,4},{4,1},{1,3},{2,4}})) // [1,2,3,4]
}