package main

// 3437. Permutations III
// Given an integer n, an alternating permutation is a permutation of the first n positive integers 
// such that no two adjacent elements are both odd or both even.

// Return all such alternating permutations sorted in lexicographical order.

// Example 1:
// Input: n = 4
// Output: [[1,2,3,4],[1,4,3,2],[2,1,4,3],[2,3,4,1],[3,2,1,4],[3,4,1,2],[4,1,2,3],[4,3,2,1]]

// Example 2:
// Input: n = 2
// Output: [[1,2],[2,1]]

// Example 3:
// Input: n = 3
// Output: [[1,2,3],[3,2,1]]

// Constraints:
//     1 <= n <= 10

import "fmt"
import "slices"

func permute(n int) [][]int {
    res, used := [][]int{}, make([]bool, n)
    var backtrack func(path []int)
    backtrack = func(path []int) {
        if len(path) == n {
            tmp := make([]int, n)
            copy(tmp, path)
            res = append(res, tmp)
            return
        }
        for i := 1; i <= n; i++ {
            if used[i-1] { continue }
            if len(path) == 0 || (path[len(path)-1] % 2 != i % 2) {
                used[i-1] = true
                backtrack(append(path, i))
                used[i-1] = false
            }
        }
    }
    backtrack([]int{})
    return res
}

func permute1(n int) [][]int {
    res, visited := [][]int{}, make([]bool, n+1)
    t := make([]int, n)
    var dfs func(i int)
    dfs = func(i int) {
        if i >= n {
            res = append(res, slices.Clone(t))
            return
        }
        for j := 1; j <= n; j++ {
            if !visited[j] && (i == 0 || t[i-1]%2 != j%2) {
                visited[j] = true
                t[i] = j
                dfs(i + 1)
                visited[j] = false
            }
        }
    }
    dfs(0)
    return res
}

func main() {
    // Example 1:
    // Input: n = 4
    // Output: [[1,2,3,4],[1,4,3,2],[2,1,4,3],[2,3,4,1],[3,2,1,4],[3,4,1,2],[4,1,2,3],[4,3,2,1]]
    fmt.Println(permute(4)) // [[1,2,3,4],[1,4,3,2],[2,1,4,3],[2,3,4,1],[3,2,1,4],[3,4,1,2],[4,1,2,3],[4,3,2,1]]
    // Example 2:
    // Input: n = 2
    // Output: [[1,2],[2,1]]
    fmt.Println(permute(2)) // [[1,2],[2,1]]
    // Example 3:
    // Input: n = 3
    // Output: [[1,2,3],[3,2,1]]
    fmt.Println(permute(3)) // [[1,2,3],[3,2,1]]

    fmt.Println(permute(1)) // [[1]]
    fmt.Println(permute(5)) // [[1 2 3 4 5] [1 2 5 4 3] [1 4 3 2 5] [1 4 5 2 3] [3 2 1 4 5] [3 2 5 4 1] [3 4 1 2 5] [3 4 5 2 1] [5 2 1 4 3] [5 2 3 4 1] [5 4 1 2 3] [5 4 3 2 1]]
    // fmt.Println(permute(6)) // 
    // fmt.Println(permute(7)) // 
    // fmt.Println(permute(8)) // 
    // fmt.Println(permute(9)) // 
    // fmt.Println(permute(10)) // 

    fmt.Println(permute1(4)) // [[1,2,3,4],[1,4,3,2],[2,1,4,3],[2,3,4,1],[3,2,1,4],[3,4,1,2],[4,1,2,3],[4,3,2,1]]
    fmt.Println(permute1(2)) // [[1,2],[2,1]]
    fmt.Println(permute1(3)) // [[1,2,3],[3,2,1]]
    fmt.Println(permute1(1)) // [[1]]
    fmt.Println(permute1(5)) // [[1 2 3 4 5] [1 2 5 4 3] [1 4 3 2 5] [1 4 5 2 3] [3 2 1 4 5] [3 2 5 4 1] [3 4 1 2 5] [3 4 5 2 1] [5 2 1 4 3] [5 2 3 4 1] [5 4 1 2 3] [5 4 3 2 1]]
    // fmt.Println(permute1(6)) // 
    // fmt.Println(permute1(7)) // 
    // fmt.Println(permute1(8)) // 
    // fmt.Println(permute1(9)) // 
    // fmt.Println(permute1(10)) // 
}