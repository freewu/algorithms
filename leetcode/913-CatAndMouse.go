package main

// 913. Cat and Mouse
// A game on an undirected graph is played by two players, Mouse and Cat, who alternate turns.

// The graph is given as follows: graph[a] is a list of all nodes b such that ab is an edge of the graph.

// The mouse starts at node 1 and goes first, the cat starts at node 2 and goes second, and there is a hole at node 0.

// During each player's turn, they must travel along one edge of the graph that meets where they are.  
// For example, if the Mouse is at node 1, it must travel to any node in graph[1].

// Additionally, it is not allowed for the Cat to travel to the Hole (node 0).

// Then, the game can end in three ways:
//     If ever the Cat occupies the same node as the Mouse, the Cat wins.
//     If ever the Mouse reaches the Hole, the Mouse wins.
//     If ever a position is repeated (i.e., the players are in the same position as a previous turn, and it is the same player's turn to move), the game is a draw.

// Given a graph, and assuming both players play optimally, return
//     1 if the mouse wins the game,
//     2 if the cat wins the game, or
//     0 if the game is a draw.
    
// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/11/17/cat1.jpg" />
// Input: graph = [[2,5],[3],[0,4,5],[1,4,5],[2,3],[0,2,3]]
// Output: 0

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/11/17/cat2.jpg" />
// Input: graph = [[1,3],[0],[3],[0,2]]
// Output: 1

// Constraints:
//     3 <= graph.length <= 50
//     1 <= graph[i].length < graph.length
//     0 <= graph[i][j] < graph.length
//     graph[i][j] != i
//     graph[i] is unique.
//     The mouse and the cat can always move. 

import "fmt"

// bfs
func catMouseGame(graph [][]int) int {
    n := len(graph)
    nonZeroNeighbor := make([]int, n)
    for i := 0; i < n; i++ {
        nonZeroNeighbor[i] = len(graph[i])
        for _, neighbor := range graph[i] {
            if neighbor == 0 {
                nonZeroNeighbor[i]--
                break
            }
        }
    }
    nextStateCount := make([][][]int, n)
    for i := 0; i < n; i++ {
        nextStateCount[i] = make([][]int, n)
        for j := 0; j < n; j++ {
            nextStateCount[i][j] = make([]int, 2)
            nextStateCount[i][j][0] = len(graph[i])
            nextStateCount[i][j][1] = nonZeroNeighbor[j]
        }
    }
    dp := make([][][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([][]int, n)
        for j := 0; j < n; j++ {
            dp[i][j] = make([]int, 2)
        }
    }
    queue := make([][]int, 0)
    for catPos := 1; catPos < n; catPos++ {
        for turn := 0; turn <= 1; turn++ {
            dp[0][catPos][turn] = 1
            queue = append(queue, []int{0, catPos, turn, 1})
            dp[catPos][catPos][turn] = 2
            queue = append(queue, []int{catPos, catPos, turn, 2})
        }
    }
    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:]
        mousePos, catPos, turn, result := curr[0], curr[1], curr[2], curr[3]
        if turn == 0 {
            if mousePos == 1 && catPos == 2 {
                return result
            }
            for _, prev := range graph[catPos] {
                if prev == 0 || dp[mousePos][prev][1] != 0 {
                    continue
                }
                if result == 2 || nextStateCount[mousePos][prev][1] == 1 {
                    nextStateCount[mousePos][prev][1]--
                    dp[mousePos][prev][1] = result
                    queue = append(queue, []int{mousePos, prev, 1, result})
                } else if nextStateCount[mousePos][prev][1] > 1 {
                    nextStateCount[mousePos][prev][1]--
                }
            }
        } else {
            for _, prev := range graph[mousePos] {
                if dp[prev][catPos][0] != 0 {
                    continue
                }
                if result == 1 || nextStateCount[prev][catPos][0] == 1 {
                    nextStateCount[prev][catPos][0]--
                    dp[prev][catPos][0] = result
                    queue = append(queue, []int{prev, catPos, 0, result})
                } else if nextStateCount[prev][catPos][0] > 1 {
                    nextStateCount[prev][catPos][0]--
                }
            }
        }
    }
    return dp[1][2][0]
}

func catMouseGame1(graph [][]int) int {
    DRAW, MOUSE_WIN, CAT_WIN := 0, 1, 2
    MOUSE_MOVE, CAT_MOVE := 0, 1
    n := len(graph)
    dp := make([][][2]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([][2]int, n)
    }
    degree := make([][][2]int, n)
    for i := 0; i < n; i++ {
        degree[i] = make([][2]int, n)
        for j := 0; j < n; j++ {
            cnt := 0
            for _, to := range graph[j] {
                if to == 0 {
                    cnt++
                }
            }
            degree[i][j][MOUSE_MOVE] = len(graph[i])
            degree[i][j][CAT_MOVE] = len(graph[j]) - cnt
        }
    }
    q := make([][3]int, 0, 16)
    for i := 1; i < n; i++ {
        dp[0][i][CAT_MOVE], dp[0][i][MOUSE_MOVE] = MOUSE_WIN, MOUSE_WIN
        dp[i][i][CAT_MOVE], dp[i][i][MOUSE_MOVE] = CAT_WIN, CAT_WIN
        q = append(q, [3]int{0, i, CAT_MOVE}, [3]int{0, i, MOUSE_MOVE}, [3]int{i, i, CAT_MOVE}, [3]int{i, i, MOUSE_MOVE})
    }
    for len(q) > 0 {
        mouse, cat, move := q[0][0], q[0][1], q[0][2]
        q = q[1:]
        if mouse == 1 && cat == 2 && move == MOUSE_MOVE {
            break
        }
        prevTurn := move ^ 1 
        moves := graph[mouse] 
        if prevTurn == 1 {
            moves = graph[cat]
        }
        for _, prevMove := range moves {
            prevMouse, prevCat := prevMove, cat
            if prevTurn == 1 {
                prevMouse, prevCat = mouse, prevMove
            } 
            if prevCat == 0 || dp[prevMouse][prevCat][prevTurn] != DRAW {
                continue
            }
            if (prevTurn == 0 && dp[mouse][cat][move] == MOUSE_WIN) ||(prevTurn == 1 && dp[mouse][cat][move] == CAT_WIN) {
                dp[prevMouse][prevCat][prevTurn] = dp[mouse][cat][move]
                q = append(q, [3]int{prevMouse, prevCat, prevTurn})
                continue
            }
            degree[prevMouse][prevCat][prevTurn]--
            if degree[prevMouse][prevCat][prevTurn] == 0 {
                dp[prevMouse][prevCat][prevTurn] = dp[mouse][cat][move]
                q = append(q, [3]int{prevMouse, prevCat, prevTurn})
            }
        }
    }
    return dp[1][2][MOUSE_MOVE]
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/11/17/cat1.jpg" />
    // Input: graph = [[2,5],[3],[0,4,5],[1,4,5],[2,3],[0,2,3]]
    // Output: 0
    fmt.Println(catMouseGame([][]int{{2,5},{3},{0,4,5},{1,4,5},{2,3},{0,2,3}})) // 0
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/11/17/cat2.jpg" />
    // Input: graph = [[1,3],[0],[3],[0,2]]
    // Output: 1
    fmt.Println(catMouseGame([][]int{{1,3},{0},{3},{0,2}})) // 1

    fmt.Println(catMouseGame1([][]int{{2,5},{3},{0,4,5},{1,4,5},{2,3},{0,2,3}})) // 0
    fmt.Println(catMouseGame1([][]int{{1,3},{0},{3},{0,2}})) // 1
}