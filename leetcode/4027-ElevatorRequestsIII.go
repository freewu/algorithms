package main

// 4027. Elevator Requests III
// You are given an integer n denoting the number of floors in a building, where the floors are numbered from 0 to n - 1.

// You are also given an integer start and a 2D integer array requests, where requests[i] = [arrivali, floori] indicates that a request for floori is made at time arrivali.

// At time 0, the elevator is at floor start.

// At each second, the elevator may move up by 1 floor, move down by 1 floor, or remain on its current floor.

// A request can be fulfilled only at or after its arrival time; it is fulfilled instantly when the elevator is on its requested floor at any time from its arrival time onward.

// Return the minimum time needed to fulfill all requests.

// Example 1:
// Input: n = 9, start = 0, requests = [[0,8],[6,5]]
// Output: 9
// Explanation:
// Move from floor 0 (start) to floor 5 (requests[1][1]) in 5 seconds, reaching at time 5. Since requests[1][0] = 6, wait until time 6 to fulfill it.
// Move from floor 5 to floor 8 (requests[0][1]) in 3 seconds, fulfilling it at time 9.
// Thus, all requests are fulfilled by time 9.

// Example 2:
// Input: n = 8, start = 5, requests = [[1,7],[7,3]]
// Output: 7
// Explanation:
// Move from floor 5 (start) to floor 7 (requests[0][1]) in 2 seconds, reaching at time 2. Since requests[0][0] = 1 has already passed, floor 7 is fulfilled at time 2.
// Move from floor 7 to floor 3 (requests[1][1]) in 4 seconds, reaching at time 6. Since requests[1][0] = 7, wait until time 7.
// Thus, all requests are fulfilled by time 7.

// Example 3:
// Input: n = 7, start = 3, requests = [[0,5],[0,1],[6,3]]
// Output: 8
// Explanation:
// Move from floor 3 (start) to floor 5 (requests[0][1]) in 2 seconds, fulfilling it at time 2.
// Move from floor 5 to floor 1 (requests[1][1]) in 4 seconds, fulfilling it at time 6.
// Move from floor 1 to floor 3 (requests[2][1]) in 2 seconds, reaching at time 8. Its request arrived at requests[2][0] = 6, so floor 3 is fulfilled at time 8.
// Thus, all requests are fulfilled by time 8.

// Constraints:
//     1 <= n <= 10^9
//     1 <= requests.length <= 16
//     requests[i] == [arrivali, floori]
//     0 <= arrivali <= 10^9
//     0 <= start, floori <= n - 1

import "fmt"
import "slices"

func elevatorRequests(n int, start int, requests [][]int) int64 {
    res, m := 1 << 61, len(requests)
    memo := make([][]int, 1 << m)
    for i := range memo {
        memo[i] = make([]int, m)
        for j := range memo[i] {
            memo[i][j] = -1 // -1 表示没有计算过
        }
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    // 返回处理完请求集合 mask，且电梯停在 requests[i][1]，所需的最短时间
    var dfs func(mask, i int) int
    dfs = func(mask, i int) int {
        mask ^= 1 << i // 这里去掉了 i
        req := requests[i]
        t, x := req[0], req[1]
        if mask == 0 {
            // i 是第一个被处理的请求
            return max(abs(x-start), t)
        }
        p := &memo[mask][i]
        if *p != -1 { // 之前计算过
            return *p
        }
        res := 1 << 61
        for j, r := range requests {
            if mask >> j & 0x1 > 0 {
                // 处理完请求 j 的时间 + 从 j 到 i 的时间
                res = min(res, dfs(mask, j)+abs(x-r[1]))
            }
        }
        // 处理完请求 i 的时间不能早于 t
        res = max(res, t)
        *p = res // 记忆化
        return res
    }
    for i := range m {
        res = min(res, dfs(1 << m - 1, i))
    }
    return int64(res)
}

func elevatorRequests1(n int, start int, requests [][]int) int64 {
    m := len(requests)
    f := make([][]int, 1 << m)
    for i := range f {
        f[i] = make([]int, m)
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i, req := range requests {
        f[1<<i][i] = max(abs(req[1]-start), req[0])
    }
    for mask := 1; mask < 1 << m; mask++ {
        if mask&(mask-1) == 0 { // mask 只有一个元素
            continue
        }
        for i, req := range requests {
            if mask>>i&1 == 0 {
                continue
            }
            res, msk := 1 << 61, mask ^ 1 << i
            t, x := req[0], req[1]
            for j, r := range requests {
                if msk>>j&1 > 0 {
                    res = min(res, f[msk][j] + abs(x - r[1]))
                }
            }
            f[mask][i] = max(res, t)
        }
    }
    return int64(slices.Min(f[1<<m-1]))
}

func elevatorRequests2(n int, start int, requests [][]int) int64 {
    m, inf := len(requests), int64(1 << 61)
    dp := make([][]int64, 1 << m)
    for mask := range dp {
        dp[mask] = make([]int64, m)
        for i := range dp[mask] {
            dp[mask][i] = inf
        }
    }
    abs := func(x int64) int64 { if x < 0 { return -x; }; return x; }
    for i, r := range requests {
        arrival, floor := int64(r[0]), int64(r[1])
        dist := abs(int64(start) - floor)
        dp[1<<i][i]=max(arrival, dist)
    }
    for mask := 1; mask < 1<<m; mask++ {
        for i := 0; i < m; i++ {
            if mask >> i & 1 == 0 || dp[mask][i] == inf {
                continue
            }
            currTime, currFloor := dp[mask][i], int64(requests[i][1])
            for j := 0; j < m;j++ {
                if mask >> j & 1 == 1 {
                    continue
                }
                arrival, nextFloor := int64(requests[j][0]), int64(requests[j][1])
                dist := abs(currFloor - nextFloor)
                nextTime := max(currTime + dist, arrival)
                nextMast := mask | 1<<j
                dp[nextMast][j] = min(dp[nextMast][j], nextTime)
            }
        }
    }
    res, all := inf, (1 << m) -1
    for i :=0;i < m; i++ {
        res = min(res, dp[all][i])
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 9, start = 0, requests = [[0,8],[6,5]]
    // Output: 9
    // Explanation:
    // Move from floor 0 (start) to floor 5 (requests[1][1]) in 5 seconds, reaching at time 5. Since requests[1][0] = 6, wait until time 6 to fulfill it.
    // Move from floor 5 to floor 8 (requests[0][1]) in 3 seconds, fulfilling it at time 9.
    // Thus, all requests are fulfilled by time 9.
    fmt.Println(elevatorRequests(9, 0, [][]int{{0,8},{6,5}})) // 9 
    // Example 2:
    // Input: n = 8, start = 5, requests = [[1,7],[7,3]]
    // Output: 7
    // Explanation:
    // Move from floor 5 (start) to floor 7 (requests[0][1]) in 2 seconds, reaching at time 2. Since requests[0][0] = 1 has already passed, floor 7 is fulfilled at time 2.
    // Move from floor 7 to floor 3 (requests[1][1]) in 4 seconds, reaching at time 6. Since requests[1][0] = 7, wait until time 7.
    // Thus, all requests are fulfilled by time 7.
    fmt.Println(elevatorRequests(8, 5, [][]int{{1,7},{7,3}})) // 7
    // Example 3:
    // Input: n = 7, start = 3, requests = [[0,5],[0,1],[6,3]]
    // Output: 8
    // Explanation:
    // Move from floor 3 (start) to floor 5 (requests[0][1]) in 2 seconds, fulfilling it at time 2.
    // Move from floor 5 to floor 1 (requests[1][1]) in 4 seconds, fulfilling it at time 6.
    // Move from floor 1 to floor 3 (requests[2][1]) in 2 seconds, reaching at time 8. Its request arrived at requests[2][0] = 6, so floor 3 is fulfilled at time 8.
    // Thus, all requests are fulfilled by time 8.
    fmt.Println(elevatorRequests(7, 3, [][]int{{0,5},{0,1},{6,3}})) // 8

    fmt.Println(elevatorRequests1(9, 0, [][]int{{0,8},{6,5}})) // 9 
    fmt.Println(elevatorRequests1(8, 5, [][]int{{1,7},{7,3}})) // 7
    fmt.Println(elevatorRequests1(7, 3, [][]int{{0,5},{0,1},{6,3}})) // 8

    fmt.Println(elevatorRequests2(9, 0, [][]int{{0,8},{6,5}})) // 9 
    fmt.Println(elevatorRequests2(8, 5, [][]int{{1,7},{7,3}})) // 7
    fmt.Println(elevatorRequests2(7, 3, [][]int{{0,5},{0,1},{6,3}})) // 8
}