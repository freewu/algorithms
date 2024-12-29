package main

// 3383. Minimum Runes to Add to Cast Spell
// Alice has just graduated from wizard school, and wishes to cast a magic spell to celebrate. 
// The magic spell contains certain focus points where magic needs to be concentrated, and some of these focus points contain magic crystals which serve as the spell's energy source. 
// Focus points can be linked through directed runes, which channel magic flow from one focus point to another.

// You are given a integer n denoting the number of focus points and an array of integers crystals where crystals[i] indicates a focus point which holds a magic crystal. 
// You are also given two integer arrays flowFrom and flowTo, which represent the existing directed runes. 
// The ith rune allows magic to freely flow from focus point flowFrom[i] to focus point flowTo[i].

// You need to find the number of directed runes Alice must add to her spell, such that each focus point either:
//     Contains a magic crystal.
//     Receives magic flow from another focus point.

// Return the minimum number of directed runes that she should add.

// Example 1:
// Input: n = 6, crystals = [0], flowFrom = [0,1,2,3], flowTo = [1,2,3,0]
// Output: 2
// Explanation: 
// <img src="https://assets.leetcode.com/uploads/2024/11/08/runesexample0.png" />
// Add two directed runes:
// From focus point 0 to focus point 4.
// From focus point 0 to focus point 5.

// Example 2:
// Input: n = 7, crystals = [3,5], flowFrom = [0,1,2,3,5], flowTo = [1,2,0,4,6]
// Output: 1
// Explanation: 
// <img src="https://assets.leetcode.com/uploads/2024/11/08/runesexample1.png" />
// Add a directed rune from focus point 4 to focus point 2.

// Constraints:
//     2 <= n <= 10^5
//     1 <= crystals.length <= n
//     0 <= crystals[i] <= n - 1
//     1 <= flowFrom.length == flowTo.length <= min(2 * 10^5, (n * (n - 1)) / 2)
//     0 <= flowFrom[i], flowTo[i] <= n - 1
//     flowFrom[i] != flowTo[i]
//     All pre-existing directed runes are distinct.

import "fmt"

func minRunesToAdd(n int, crystals []int, flowFrom []int, flowTo []int) int {
    res, graph := 0, make([][]int, n)
    for i := 0; i < len(flowFrom); i++ {
        graph[flowFrom[i]] = append(graph[flowFrom[i]], flowTo[i])
    }
    visited := make([]int, n)
    for _, x := range crystals {
        visited[x] = 1
    }
    bfs := func(queue []int) {
        for len(queue) > 0 {
            x := queue[0]
            queue = queue[1:]
            for _, v := range graph[x] {
                if visited[v] == 1 { continue }
                visited[v] = 1
                queue = append(queue, v)
            }
        }
    }
    seq := []int{}
    var dfs func(x int)
    dfs = func(x int) {
        visited[x] = 2
        for _, v := range graph[x] {
            if visited[v] > 0 { continue }
            dfs(v)
        }
        seq = append(seq, x)
    }
    queue := crystals
    bfs(queue)
    for i := 0; i < n; i++ {
        if visited[i] == 0 {
            dfs(i)
        }
    }
    for i, j := 0, len(seq)-1; i < j; i, j = i+1, j-1 {
        seq[i], seq[j] = seq[j], seq[i]
    }
    for _, i := range seq {
        if visited[i] == 2 {
            queue = []int{ i }
            visited[i] = 1
            bfs(queue)
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 6, crystals = [0], flowFrom = [0,1,2,3], flowTo = [1,2,3,0]
    // Output: 2
    // Explanation: 
    // <img src="https://assets.leetcode.com/uploads/2024/11/08/runesexample0.png" />
    // Add two directed runes:
    // From focus point 0 to focus point 4.
    // From focus point 0 to focus point 5.
    fmt.Println(minRunesToAdd(6, []int{0}, []int{0,1,2,3}, []int{1,2,3,0})) // 2
    // Example 2:
    // Input: n = 7, crystals = [3,5], flowFrom = [0,1,2,3,5], flowTo = [1,2,0,4,6]
    // Output: 1
    // Explanation: 
    // <img src="https://assets.leetcode.com/uploads/2024/11/08/runesexample1.png" />
    // Add a directed rune from focus point 4 to focus point 2.
    fmt.Println(minRunesToAdd(7, []int{3,5}, []int{0,1,2,3,5}, []int{1,2,0,4,6})) // 1
}