package main

// 4030. Elevator Requests IV
// You are given an integer n denoting the number of floors in a building, where the floors are numbered from 0 to n - 1.

// You are also given an integer start and a 2D integer array requests, where requests[i] = [arrivali, floori] indicates that a request for floori is made at time arrivali.

// At time 0, the elevator is at floor start.

// At each second, the elevator may move up by 1 floor, move down by 1 floor, or remain on its current floor.

// A request can be fulfilled only at or after its arrival time; 
// it is fulfilled instantly when the elevator is on its requested floor at any time from its arrival time onward.

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
//     1 <= requests.length <= 500
//     requests[i] == [arrivali, floori]
//     0 <= arrivali <= 10^9
//     0 <= start, floori <= n - 1

import "fmt"
import "slices"

// 执行出错 511 / 550 oom
func elevatorRequests(n int, start int, requests [][]int) int64 {
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

func elevatorRequests1(n int, start int, requests [][]int) int64 {
    requests = append(requests, []int{0, -1}, []int{0, n}) // 插入两个哨兵
    slices.SortFunc(requests, func(a, b []int) int { // 按楼层排序
         return a[1] - b[1] 
    }) 
    m := len(requests) // 不含哨兵的下标范围是 [1, m-2]
    f := make([][2]int, m)
    for j := range f {
        f[j] = [2]int{ 1 << 61, 1 << 61 }
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 1; i < m-1; i++ {
        t, x := requests[i][0], requests[i][1]
        for j := m - 2; j >= i; j-- {
            t2, y := requests[j][0], requests[j][1]
            if i == 1 && j == m-2 { // 当前请求是第一个请求
                // 从 start 到当前楼层
                f[j][0] = max(abs(x-start), t)
                f[j][1] = max(abs(y-start), t2)
                continue
            }
            f[j][1] = min(max(f[j][0]+y-requests[i-1][1], t2), // 从 floor[i-1] 到当前楼层
                        max(f[j+1][1]+requests[j+1][1]-y, t2)) // 从 floor[j+1] 到当前楼层
            f[j][0] = min(max(f[j][0]+x-requests[i-1][1], t), // 从 floor[i-1] 到当前楼层
                        max(f[j+1][1]+requests[j+1][1]-x, t)) // 从 floor[j+1] 到当前楼层
        }
    }
    // 枚举最后一个完成的请求
    res := 1 << 61
    for i := 1; i < m-1; i++ {
        res = min(res, f[i][0])
    }
    return int64(res)
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

    //fmt.Println(elevatorRequests(41, 32, [][]int{{11,3},{17,14},{10,5},{0,17},{2,25},{2,28},{16,12},{15,36},{13,22},{6,13},{6,2},{3,16},{12,33},{4,38},{17,9},{8,29},{18,26},{18,37},{7,32},{17,6},{13,10},{2,31},{0,40},{16,39},{9,23},{0,0},{10,34},{20,35},{4,30},{21,19},{10,4},{14,11},{4,18}})) // 55

    fmt.Println(elevatorRequests1(9, 0, [][]int{{0,8},{6,5}})) // 9
    fmt.Println(elevatorRequests1(8, 5, [][]int{{1,7},{7,3}})) // 7
    fmt.Println(elevatorRequests1(7, 3, [][]int{{0,5},{0,1},{6,3}})) // 8
    fmt.Println(elevatorRequests1(41, 32, [][]int{{11,3},{17,14},{10,5},{0,17},{2,25},{2,28},{16,12},{15,36},{13,22},{6,13},{6,2},{3,16},{12,33},{4,38},{17,9},{8,29},{18,26},{18,37},{7,32},{17,6},{13,10},{2,31},{0,40},{16,39},{9,23},{0,0},{10,34},{20,35},{4,30},{21,19},{10,4},{14,11},{4,18}})) // 55

}