package main

// 1916. Count Ways to Build Rooms in an Ant Colony
// You are an ant tasked with adding n new rooms numbered 0 to n-1 to your colony. 
// You are given the expansion plan as a 0-indexed integer array of length n, prevRoom, 
// where prevRoom[i] indicates that you must build room prevRoom[i] before building room i, 
// and these two rooms must be connected directly. 
// Room 0 is already built, so prevRoom[0] = -1. 
// The expansion plan is given such that once all the rooms are built, every room will be reachable from room 0.

// You can only build one room at a time, and you can travel freely between rooms you have already built only if they are connected. 
// You can choose to build any room as long as its previous room is already built.

// Return the number of different orders you can build all the rooms in. 
// Since the answer may be large, return it modulo 10^9 + 7.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/06/19/d1.JPG" />
// Input: prevRoom = [-1,0,1]
// Output: 1
// Explanation: There is only one way to build the additional rooms: 0 → 1 → 2

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/06/19/d2.JPG" />
// Input: prevRoom = [-1,0,0,1,2]
// Output: 6
// Explanation:
// The 6 ways are:
// 0 → 1 → 3 → 2 → 4
// 0 → 2 → 4 → 1 → 3
// 0 → 1 → 2 → 3 → 4
// 0 → 1 → 2 → 4 → 3
// 0 → 2 → 1 → 3 → 4
// 0 → 2 → 1 → 4 → 3

// Constraints:
//     n == prevRoom.length
//     2 <= n <= 10^5
//     prevRoom[0] == -1
//     0 <= prevRoom[i] < n for all 1 <= i < n
//     Every room is reachable from room 0 once all the rooms are built.

import "fmt"

func waysToBuildRooms(prevRoom []int) int {
    n, mod := len(prevRoom), 1_000_000_007
    fac, inv, graph, f := make([]int, n + 1), make([]int, n + 1), make(map[int][]int), 1
    power := func(x, y int) int {
        res := 1
        for ; y > 0; y >>= 1 {
            if (y & 1) == 1 {
                res = res * x % mod
            }
            x = x * x % mod
        }
        return res
    }
    for i := 1; i <= n; i++ {
        f = f * i % mod
        fac[i] = f
        inv[i] = power(f, mod - 2)
        graph[prevRoom[i-1]] = append(graph[prevRoom[i-1]], i-1)
    }
    var dfs func(i int) []int
    dfs = func(i int) []int {
        size, res := 0, 1
        for _, child := range graph[i] {
            cur := dfs(child)
            size += cur[0]
            if size > cur[0] {
                res = (((res * cur[1] % mod) * fac[size] % mod) * inv[cur[0]] % mod) * inv[size - cur[0]] % mod
            } else {
                res = res * cur[1] % mod
            }
        }
        return []int{size + 1, res }
    }
    res := dfs(0)
    return res[1]
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/06/19/d1.JPG" />
    // Input: prevRoom = [-1,0,1]
    // Output: 1
    // Explanation: There is only one way to build the additional rooms: 0 → 1 → 2
    fmt.Println(waysToBuildRooms([]int{-1,0,1})) // 1
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/06/19/d2.JPG" />
    // Input: prevRoom = [-1,0,0,1,2]
    // Output: 6
    // Explanation:
    // The 6 ways are:
    // 0 → 1 → 3 → 2 → 4
    // 0 → 2 → 4 → 1 → 3
    // 0 → 1 → 2 → 3 → 4
    // 0 → 1 → 2 → 4 → 3
    // 0 → 2 → 1 → 3 → 4
    // 0 → 2 → 1 → 4 → 3
    fmt.Println(waysToBuildRooms([]int{-1,0,0,1,2})) // 6
}