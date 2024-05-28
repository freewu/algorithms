package main

// 1601. Maximum Number of Achievable Transfer Requests
// We have n buildings numbered from 0 to n - 1. 
// Each building has a number of employees. 
// It's transfer season, and some employees want to change the building they reside in.

// You are given an array requests where requests[i] = [fromi, toi] represents an employee's request to transfer from building fromi to building toi.

// All buildings are full, so a list of requests is achievable only if for each building, the net change in employee transfers is zero. 
// This means the number of employees leaving is equal to the number of employees moving in. 
// For example if n = 3 and two employees are leaving building 0, one is leaving building 1, 
// and one is leaving building 2, there should be two employees moving to building 0, 
// one employee moving to building 1, and one employee moving to building 2.

// Return the maximum number of achievable requests.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/09/10/move1.jpg" />
// Input: n = 5, requests = [[0,1],[1,0],[0,1],[1,2],[2,0],[3,4]]
// Output: 5
// Explantion: Let's see the requests:
// From building 0 we have employees x and y and both want to move to building 1.
// From building 1 we have employees a and b and they want to move to buildings 2 and 0 respectively.
// From building 2 we have employee z and they want to move to building 0.
// From building 3 we have employee c and they want to move to building 4.
// From building 4 we don't have any requests.
// We can achieve the requests of users x and b by swapping their places.
// We can achieve the requests of users y, a and z by swapping the places in the 3 buildings.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/09/10/move2.jpg" />
// Input: n = 3, requests = [[0,0],[1,2],[2,1]]
// Output: 3
// Explantion: Let's see the requests:
// From building 0 we have employee x and they want to stay in the same building 0.
// From building 1 we have employee y and they want to move to building 2.
// From building 2 we have employee z and they want to move to building 1.
// We can achieve all the requests. 

// Example 3:
// Input: n = 4, requests = [[0,3],[3,1],[1,2],[2,0]]
// Output: 4

// Constraints:
//     1 <= n <= 20
//     1 <= requests.length <= 16
//     requests[i].length == 2
//     0 <= fromi, toi < n

import "fmt"
import "math/bits"

func maximumRequests(n int, requests [][]int) int {
    buildingState := map[int]int{}
    requestsFulfilled, maxRequestFulfilled := 0, 0
    var backtracking func(int)
    backtracking = func(idxToConsider int) {
        isBalanced := true
        for _, state := range buildingState {
            if state != 0 {
                isBalanced = false
            }
        }
        if isBalanced {
            if requestsFulfilled > maxRequestFulfilled {
                maxRequestFulfilled = requestsFulfilled
            }
        }
        for idx := idxToConsider; idx < len(requests); idx++ {
            req := requests[idx]
            buildingState[req[0]]--
            buildingState[req[1]]++
            requestsFulfilled++
            backtracking(idx+1)
            buildingState[req[0]]++
            buildingState[req[1]]--
            requestsFulfilled--

        }
    }
    backtracking(0)
    return maxRequestFulfilled
}

func maximumRequests1(n int, requests [][]int) int {
    cache := make([][2]int, n) // n 栋楼： 0: 出  1:入
    for i := 0; i < n; i++ {
        cache[i] = [2]int{0, 0}
    }
    res, m := 0, len(requests) // 请求的长度
    var dfs func(i, cnt int) // i 请求的下标 cnt当前满足了多少个请求
    dfs = func(i, cnt int) {
        if i == m {
            for j := range cache { // 每栋楼
                if cache[j][0] + cache[j][1] != 0 {
                    return
                }
            }
            if cnt > res {
                res = cnt
            }
            return
        }
        dfs(i + 1, cnt) // 当前请求不满足
        // 满足当前请求 楼 requests[i][0]--  楼 requests[i][1]++
        cache[requests[i][0]][0]--
        cache[requests[i][1]][1]++
        dfs(i+1, cnt+1)
        cache[requests[i][0]][0]++
        cache[requests[i][1]][1]--
    }
    dfs(0, 0)
    return res
}

func maximumRequests2(n int, requests [][]int) int {
    numRequests := len(requests)
    numCombinations := 1 << numRequests
    indegree := make([]int, n)
    var curMask, nextMask, diff uint
    res, countValid := 0, n
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < numCombinations; i++ {
        nextMask = uint(i ^ (i>>1))
        diff = curMask ^ nextMask
        idx := bits.TrailingZeros(diff)
        from, to := requests[idx][0], requests[idx][1]
        fromBefore, toBefore := indegree[from], indegree[to]
        if (curMask>>idx) & 1 == 0 {
            indegree[from]--
            indegree[to]++
        } else {
            indegree[from]++
            indegree[to]--
        }
        if indegree[from] == 0 && fromBefore != 0 {
            countValid++
        } else if indegree[from] != 0 && fromBefore == 0 {
            countValid--
        }
        if indegree[to] == 0 && toBefore != 0 {
            countValid++
        } else if indegree[to] != 0 && toBefore == 0 {
            countValid--
        }
        curMask = nextMask
        if countValid != n {
            continue
        }
        res = max(res, bits.OnesCount(nextMask)) // count number of 1 in nextMask
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/09/10/move1.jpg" />
    // Input: n = 5, requests = [[0,1],[1,0],[0,1],[1,2],[2,0],[3,4]]
    // Output: 5
    // Explantion: Let's see the requests:
    // From building 0 we have employees x and y and both want to move to building 1.
    // From building 1 we have employees a and b and they want to move to buildings 2 and 0 respectively.
    // From building 2 we have employee z and they want to move to building 0.
    // From building 3 we have employee c and they want to move to building 4.
    // From building 4 we don't have any requests.
    // We can achieve the requests of users x and b by swapping their places.
    // We can achieve the requests of users y, a and z by swapping the places in the 3 buildings.
    fmt.Println(maximumRequests(5,[][]int{{0,1},{1,0},{0,1},{1,2},{2,0},{3,4}})) // 5
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/09/10/move2.jpg" />
    // Input: n = 3, requests = [[0,0],[1,2],[2,1]]
    // Output: 3
    // Explantion: Let's see the requests:
    // From building 0 we have employee x and they want to stay in the same building 0.
    // From building 1 we have employee y and they want to move to building 2.
    // From building 2 we have employee z and they want to move to building 1.
    // We can achieve all the requests. 
    fmt.Println(maximumRequests(5,[][]int{{0,0},{1,2},{2,1}})) // 3
    // Example 3:
    // Input: n = 4, requests = [[0,3],[3,1],[1,2],[2,0]]
    // Output: 4
    fmt.Println(maximumRequests(5,[][]int{{0,3},{3,1},{1,2},{2,0}})) // 4

    fmt.Println(maximumRequests1(5,[][]int{{0,1},{1,0},{0,1},{1,2},{2,0},{3,4}})) // 5
    fmt.Println(maximumRequests1(5,[][]int{{0,0},{1,2},{2,1}})) // 3
    fmt.Println(maximumRequests1(5,[][]int{{0,3},{3,1},{1,2},{2,0}})) // 4

    fmt.Println(maximumRequests2(5,[][]int{{0,1},{1,0},{0,1},{1,2},{2,0},{3,4}})) // 5
    fmt.Println(maximumRequests2(5,[][]int{{0,0},{1,2},{2,1}})) // 3
    fmt.Println(maximumRequests2(5,[][]int{{0,3},{3,1},{1,2},{2,0}})) // 4
}