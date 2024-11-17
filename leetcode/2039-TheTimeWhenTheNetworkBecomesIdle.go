package main

// 2039. The Time When the Network Becomes Idle
// There is a network of n servers, labeled from 0 to n - 1. 
// You are given a 2D integer array edges, where edges[i] = [ui, vi] indicates there is a message channel between servers ui and vi, 
// and they can pass any number of messages to each other directly in one second. 
// You are also given a 0-indexed integer array patience of length n.

// All servers are connected, i.e., a message can be passed from one server to any other server(s) directly or indirectly through the message channels.

// The server labeled 0 is the master server. The rest are data servers. 
// Each data server needs to send its message to the master server for processing and wait for a reply. 
// Messages move between servers optimally, so every message takes the least amount of time to arrive at the master server. 
// The master server will process all newly arrived messages instantly and send a reply to the originating server via the reversed path the message had gone through.

// At the beginning of second 0, each data server sends its message to be processed. 
// Starting from second 1, at the beginning of every second, each data server will check if it has received a reply to the message it sent (including any newly arrived replies) from the master server:
//     1. If it has not, it will resend the message periodically. 
//        The data server i will resend the message every patience[i] second(s), i.e., the data server i will resend the message if patience[i] second(s) have elapsed since the last time the message was sent from this server.
//     2. Otherwise, no more resending will occur from this server.

// The network becomes idle when there are no messages passing between servers or arriving at servers.

// Return the earliest second starting from which the network becomes idle.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/09/22/quiet-place-example1.png" />
// Input: edges = [[0,1],[1,2]], patience = [0,2,1]
// Output: 8
// Explanation:
// At (the beginning of) second 0,
// - Data server 1 sends its message (denoted 1A) to the master server.
// - Data server 2 sends its message (denoted 2A) to the master server.
// At second 1,
// - Message 1A arrives at the master server. Master server processes message 1A instantly and sends a reply 1A back.
// - Server 1 has not received any reply. 1 second (1 < patience[1] = 2) elapsed since this server has sent the message, therefore it does not resend the message.
// - Server 2 has not received any reply. 1 second (1 == patience[2] = 1) elapsed since this server has sent the message, therefore it resends the message (denoted 2B).
// At second 2,
// - The reply 1A arrives at server 1. No more resending will occur from server 1.
// - Message 2A arrives at the master server. Master server processes message 2A instantly and sends a reply 2A back.
// - Server 2 resends the message (denoted 2C).
// ...
// At second 4,
// - The reply 2A arrives at server 2. No more resending will occur from server 2.
// ...
// At second 7, reply 2D arrives at server 2.
// Starting from the beginning of the second 8, there are no messages passing between servers or arriving at servers.
// This is the time when the network becomes idle.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/09/04/network_a_quiet_place_2.png">
// Input: edges = [[0,1],[0,2],[1,2]], patience = [0,10,10]
// Output: 3
// Explanation: Data servers 1 and 2 receive a reply back at the beginning of second 2.
// From the beginning of the second 3, the network becomes idle.

// Constraints:
//     n == patience.length
//     2 <= n <= 10^5
//     patience[0] == 0
//     1 <= patience[i] <= 10^5 for 1 <= i < n
//     1 <= edges.length <= min(10^5, n * (n - 1) / 2)
//     edges[i].length == 2
//     0 <= ui, vi < n
//     ui != vi
//     There are no duplicate edges.
//     Each server can directly or indirectly reach another server.

import "fmt"
import "math"

// dfs
func networkBecomesIdle(edges [][]int, patience []int) int {
    n := len(patience)
    adj, mn, visited := make([][]int, n), make([]int, n), make([]bool, n)
    visited[0] = true
    for _, v := range edges {
        adj[v[0]] = append(adj[v[0]], v[1])
        adj[v[1]] = append(adj[v[1]], v[0])
    }
    var dfs func(i, level int)
    dfs = func(i, level int) {
        for _, v := range adj[i] {
            if visited[v] {
                if mn[v] > level * 2 {
                    mn[v] = level * 2
                    dfs(v, level + 1)
                }
                continue
            }
            visited[v] = true
            mn[v] = level * 2
            dfs(v, level + 1)
        }
    }
    dfs(0, 1)
    mx := 0
    for i := 1; i < n; i++ {
        if len(adj[i]) == 0 { continue }
        if patience[i] >= mn[i] {
            mx = max(mx, mn[i])
        } else {
            mx = max(mx, mn[i] + patience[i] * int(math.Ceil(float64(mn[i]) / float64(patience[i])) - 1))
        }
    }
    return mx + 1
}

// bfs
func networkBecomesIdle1(edges [][]int, patience []int) int {
    n, inf := len(patience), -(1 << 32)
    adj := make([][]int, n)
    for _, v := range edges {
        adj[v[0]] = append(adj[v[0]], v[1])
        adj[v[1]] = append(adj[v[1]], v[0])
    }
    dis := make([]int, n)
    for i := range dis {
        dis[i] = inf
    }
    dis[0] = 0
    queue := []int{ 0 }
    for len(queue) > 0 {
        x := queue[0]
        queue = queue[1:]
        for _, y := range adj[x] {
            if dis[y] >= 0 { continue }
            dis[y] = dis[x] + 1
            queue = append(queue, y)
        }
    }
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < n; i++ {
        d, p := dis[i] * 2, patience[i]
        if p > 0 {
            res = max(res, (d - 1) / p * p + d)
        }
    }
    return res + 1
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/09/22/quiet-place-example1.png" />
    // Input: edges = [[0,1],[1,2]], patience = [0,2,1]
    // Output: 8
    // Explanation:
    // At (the beginning of) second 0,
    // - Data server 1 sends its message (denoted 1A) to the master server.
    // - Data server 2 sends its message (denoted 2A) to the master server.
    // At second 1,
    // - Message 1A arrives at the master server. Master server processes message 1A instantly and sends a reply 1A back.
    // - Server 1 has not received any reply. 1 second (1 < patience[1] = 2) elapsed since this server has sent the message, therefore it does not resend the message.
    // - Server 2 has not received any reply. 1 second (1 == patience[2] = 1) elapsed since this server has sent the message, therefore it resends the message (denoted 2B).
    // At second 2,
    // - The reply 1A arrives at server 1. No more resending will occur from server 1.
    // - Message 2A arrives at the master server. Master server processes message 2A instantly and sends a reply 2A back.
    // - Server 2 resends the message (denoted 2C).
    // ...
    // At second 4,
    // - The reply 2A arrives at server 2. No more resending will occur from server 2.
    // ...
    // At second 7, reply 2D arrives at server 2.
    // Starting from the beginning of the second 8, there are no messages passing between servers or arriving at servers.
    // This is the time when the network becomes idle.
    fmt.Println(networkBecomesIdle([][]int{{0,1},{1,2}}, []int{0,2,1})) // 8
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/09/04/network_a_quiet_place_2.png">
    // Input: edges = [[0,1],[0,2],[1,2]], patience = [0,10,10]
    // Output: 3
    // Explanation: Data servers 1 and 2 receive a reply back at the beginning of second 2.
    // From the beginning of the second 3, the network becomes idle.
    fmt.Println(networkBecomesIdle([][]int{{0,1},{0,2},{1,2}}, []int{0,10,10})) // 3

    fmt.Println(networkBecomesIdle1([][]int{{0,1},{1,2}}, []int{0,2,1})) // 8
    fmt.Println(networkBecomesIdle1([][]int{{0,1},{0,2},{1,2}}, []int{0,10,10})) // 3
}