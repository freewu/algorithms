package main

// 2467. Most Profitable Path in a Tree
// There is an undirected tree with n nodes labeled from 0 to n - 1, rooted at node 0. 
// You are given a 2D integer array edges of length n - 1 where edges[i] = [ai, bi] indicates 
// that there is an edge between nodes ai and bi in the tree.

// At every node i, there is a gate. You are also given an array of even integers amount, where amount[i] represents:
//     the price needed to open the gate at node i, if amount[i] is negative, or,
//     the cash reward obtained on opening the gate at node i, otherwise.

// The game goes on as follows:
//     1. Initially, Alice is at node 0 and Bob is at node bob.
//     2. At every second, Alice and Bob each move to an adjacent node. 
//        Alice moves towards some leaf node, while Bob moves towards node 0.
//     3. For every node along their path, Alice and Bob either spend money to open the gate at that node, 
//        or accept the reward. Note that:
//             3.1 If the gate is already open, no price will be required, nor will there be any cash reward.
//             3.2 If Alice and Bob reach the node simultaneously, they share the price/reward for opening the gate there. 
//                 In other words, if the price to open the gate is c, then both Alice and Bob pay c / 2 each. 
//                 Similarly, if the reward at the gate is c, both of them receive c / 2 each.
//     4. If Alice reaches a leaf node, she stops moving. Similarly, if Bob reaches node 0, he stops moving. 
//        Note that these events are independent of each other.

// Return the maximum net income Alice can have if she travels towards the optimal leaf node.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/10/29/eg1.png" />
// Input: edges = [[0,1],[1,2],[1,3],[3,4]], bob = 3, amount = [-2,4,2,-4,6]
// Output: 6
// Explanation: 
// The above diagram represents the given tree. The game goes as follows:
// - Alice is initially on node 0, Bob on node 3. They open the gates of their respective nodes.
//   Alice's net income is now -2.
// - Both Alice and Bob move to node 1. 
//   Since they reach here simultaneously, they open the gate together and share the reward.
//   Alice's net income becomes -2 + (4 / 2) = 0.
// - Alice moves on to node 3. Since Bob already opened its gate, Alice's income remains unchanged.
//   Bob moves on to node 0, and stops moving.
// - Alice moves on to node 4 and opens the gate there. Her net income becomes 0 + 6 = 6.
// Now, neither Alice nor Bob can make any further moves, and the game ends.
// It is not possible for Alice to get a higher net income.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/10/29/eg2.png" />
// Input: edges = [[0,1]], bob = 1, amount = [-7280,2350]
// Output: -7280
// Explanation: 
// Alice follows the path 0->1 whereas Bob follows the path 1->0.
// Thus, Alice opens the gate at node 0 only. Hence, her net income is -7280. 

// Constraints:
//     2 <= n <= 10^5
//     edges.length == n - 1
//     edges[i].length == 2
//     0 <= ai, bi < n
//     ai != bi
//     edges represents a valid tree.
//     1 <= bob < n
//     amount.length == n
//     amount[i] is an even integer in the range [-10^4, 10^4].

import "fmt"

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
    n := len(amount)
    visited, graph := make([]bool, n), make([][]int, n)
    for _, v := range edges {
        graph[v[0]] = append(graph[v[0]], v[1])
        graph[v[1]] = append(graph[v[1]], v[0])
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(node, level int) (int, int)
    dfs = func(node, level int) (int, int) {
        visited[node] = true
        mx, bobLevel := -1 << 31, -2
        if node == bob {
            bobLevel = level
        }
        for _, next := range graph[node] {
            if visited[next] { continue }
            nextMax, nextLevel := dfs(next, level + 1)
            mx, bobLevel = max(mx, nextMax), max(bobLevel, nextLevel)
        }
        if mx == -1 << 31 { mx = 0 }
        if level * 2 == bobLevel + 1 {
            mx += amount[node] / 2
        } else if bobLevel == -2 || level * 2 <= bobLevel {
            mx += amount[node];
        }
        return mx, bobLevel 
    }
    res, _ := dfs(0, 1)
    return res
}

func mostProfitablePath1(edges [][]int, bob int, amount []int) int {
    n := len(edges) + 1
    mp, seen := make([][]int, n), make([]bool, n)
    for _, v := range edges {
        mp[v[0]] = append(mp[v[0]], v[1])
        mp[v[1]] = append(mp[v[1]], v[0])
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(node, distination int) (int, int)
    dfs = func(node, distination int) (int, int) {
        seen[node] = true
        mx, steps := -1 << 31, n
        if node == bob {
            steps = 0
        }
        for _, next := range mp[node] {
            if seen[next] {continue }
            cur, k := dfs(next, distination + 1)
            mx, steps = max(mx, cur), min(steps, k)
        }
        if mx == -1 << 31 {
            mx = 0
        }
        if distination == steps {
            mx += amount[node] / 2
        } else if distination < steps {
            mx += amount[node]
        }
        return mx, steps + 1
    }
    res, _ := dfs(0, 0)
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/10/29/eg1.png" />
    // Input: edges = [[0,1],[1,2],[1,3],[3,4]], bob = 3, amount = [-2,4,2,-4,6]
    // Output: 6
    // Explanation: 
    // The above diagram represents the given tree. The game goes as follows:
    // - Alice is initially on node 0, Bob on node 3. They open the gates of their respective nodes.
    //   Alice's net income is now -2.
    // - Both Alice and Bob move to node 1. 
    //   Since they reach here simultaneously, they open the gate together and share the reward.
    //   Alice's net income becomes -2 + (4 / 2) = 0.
    // - Alice moves on to node 3. Since Bob already opened its gate, Alice's income remains unchanged.
    //   Bob moves on to node 0, and stops moving.
    // - Alice moves on to node 4 and opens the gate there. Her net income becomes 0 + 6 = 6.
    // Now, neither Alice nor Bob can make any further moves, and the game ends.
    // It is not possible for Alice to get a higher net income.
    fmt.Println(mostProfitablePath([][]int{{0,1},{1,2},{1,3},{3,4}}, 3, []int{-2,4,2,-4,6})) // 6
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/10/29/eg2.png" />
    // Input: edges = [[0,1]], bob = 1, amount = [-7280,2350]
    // Output: -7280
    // Explanation: 
    // Alice follows the path 0->1 whereas Bob follows the path 1->0.
    // Thus, Alice opens the gate at node 0 only. Hence, her net income is -7280. 
    fmt.Println(mostProfitablePath([][]int{{0,1}}, 1, []int{-7280,2350})) // -7280

    fmt.Println(mostProfitablePath1([][]int{{0,1},{1,2},{1,3},{3,4}}, 3, []int{-2,4,2,-4,6})) // 6
    fmt.Println(mostProfitablePath1([][]int{{0,1}}, 1, []int{-7280,2350})) // -7280
}