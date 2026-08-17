package main

// 4023. Elevator Requests II
// You are given an integer n denoting the number of floors in a building, where the floors are numbered from 0 to n - 1.

// You are also given an integer start, representing the floor where the elevator begins, and an integer array requests, where requests[i] is a floor that the elevator is requested to reach. 
// All floors in requests are distinct.

// At time 0, the elevator is on floor start, and all requests are made simultaneously.

// During each second before all requests are fulfilled, the elevator moves exactly one floor, either up or down.
// A request is fulfilled instantly when the elevator reaches its requested floor. 
// If start appears in requests, that request is fulfilled at time 0.

// For each second that a request remains unfulfilled, you receive 1 penalty. 
// Equivalently, a request fulfilled at time t contributes t to the total penalty.

// Return the minimum total penalty required to fulfill all requests.

// Example 1:
// Input: n = 6, start = 4, requests = [1,5]
// Output: 6
// Explanation:
// Move from floor 4 (start) to floor 5 in 1 second. Penalty for floor 5 is 1.
// Move from floor 5 to floor 1 in 4 seconds. Penalty for floor 1 is 5.
// Thus, the total penalty is 1 + 5 = 6.

// Example 2:
// Input: n = 8, start = 3, requests = [3,7,1]
// Output: 10
// Explanation:
// Floor 3 (start) is fulfilled instantly. Penalty for floor 3 is 0.
// Move from floor 3 to floor 1 in 2 seconds. Penalty for floor 1 is 2.
// Move from floor 1 to floor 7 in 6 seconds. Penalty for floor 7 is 8.
// Thus, the total penalty is 0 + 2 + 8 = 10.

// Example 3:
// Input: n = 10, start = 5, requests = [0,2,9]
// Output: 22
// Explanation:
// Move from floor 5 (start) to floor 2 in 3 seconds. Penalty for floor 2 is 3.
// Move from floor 2 to floor 0 in 2 seconds. Penalty for floor 0 is 5.
// Move from floor 0 to floor 9 in 9 seconds. Penalty for floor 9 is 14.
// Thus, the total penalty is 3 + 5 + 14 = 22.

// Constraints:
//     1 <= n <= 10^9
//     1 <= requests.length <= 1500
//     0 <= start, requests[i] <= n - 1
//     All values in requests are distinct.

import "fmt"
import "sort"

func elevatorRequests(n int, start int, requests []int) int64 {
    requests = append(requests, start, -1, n) // 插入 start 和两个哨兵
    sort.Ints(requests)
    m := len(requests)
    memo := make([][][2]int, m - 1)
    for i := range memo {
        memo[i] = make([][2]int, m-1)
        for j := range memo[i] {
            memo[i][j] = [2]int{-1, -1} // -1 表示该状态没有计算过
        }
    }
    // 已处理完 requests 的子数组 [i, j]
    // isRight = 0 表示电梯在 requests[i]
    // isRight = 1 表示电梯在 requests[j]
    var dfs func(i, j, isRight int) int 
    dfs = func(i, j, isRight int) int {
        if i == 0 || j == m-1 { // 出界
            return 1 << 61
        }
        if i == 1 && j == m-2 { // 已处理完所有请求
            return 0
        }
        p := &memo[i][j][isRight]
        if *p != -1 { // 之前计算过
            return *p
        }
        v := 0
        if isRight > 0 {
            v = requests[j]
        } else {
            v = requests[i]
        }
        remain := m - 3 - j + i
        *p = min(dfs(i - 1, j, 0) + (v - requests[i - 1]) * remain, // 往左
                dfs(i, j + 1, 1) + (requests[j + 1] - v) * remain) // 往右
        return *p
    }
    index := sort.SearchInts(requests, start)
    return int64(dfs(index, index, 0)) // 这里 0 和 1 是一样的
}

func elevatorRequests1(n int, start int, requests []int) int64 {
    inf := int64(1 << 61)
    pos := make([]int, 0, len(requests) + 1)
    pos = append(pos, start)
    for _, floor := range requests {
        if floor != start {
            pos = append(pos, floor)
        }
    }
    if len(pos) == 1 {
        return 0
    }
    for i := 1; i < len(pos); i++ {
        val := pos[i]
        j := i - 1
        for j >= 0 && pos[j] > val {
            pos[j+1] = pos[j]
            j--
        }
        pos[j+1] = val
    }
    startIndex := 0
    for pos[startIndex] != start {
        startIndex++
    }
    leftCount, rightCount := startIndex, len(pos) - startIndex - 1
    requestCount := leftCount + rightCount
    prevLeft, prevRight := make([]int64, rightCount + 1), make([]int64, rightCount + 1)
    prevLeft[0], prevRight[0] = 0, 0
    for j := 1; j <= rightCount; j++ {
        remaining := int64(requestCount - j + 1)
        newRight := pos[startIndex+j]
        fromLeft := prevLeft[j-1] + int64(newRight-pos[startIndex])*remaining
        fromRight := prevRight[j-1] + int64(newRight-pos[startIndex+j-1])*remaining
        prevLeft[j] = inf
        if fromLeft < fromRight {
            prevRight[j] = fromLeft
        } else {
            prevRight[j] = fromRight
        }
    }
    for i := 1; i <= leftCount; i++ {
        currentLeft, currentRight := make([]int64, rightCount + 1), make([]int64, rightCount + 1)
        newLeft, remaining := pos[startIndex-i], int64(requestCount - i + 1)
        fromLeft := prevLeft[0] + int64(pos[startIndex-i+1]-newLeft)*remaining
        fromRight := prevRight[0] + int64(pos[startIndex]-newLeft)*remaining
        if fromLeft < fromRight {
            currentLeft[0] = fromLeft
        } else {
            currentLeft[0] = fromRight
        }
        currentRight[0] = inf
        for j := 1; j <= rightCount; j++ {
            remaining = int64(requestCount - i - j + 1)
            newRight := pos[startIndex+j]
            fromLeft = prevLeft[j] + int64(pos[startIndex-i+1]-newLeft)*remaining
            fromRight = prevRight[j] + int64(pos[startIndex+j]-newLeft)*remaining
            if fromLeft < fromRight {
                currentLeft[j] = fromLeft
            } else {
                currentLeft[j] = fromRight
            }
            fromLeft = currentLeft[j-1] + int64(newRight-newLeft)*remaining
            fromRight = currentRight[j-1] + int64(newRight-pos[startIndex+j-1])*remaining
            if fromLeft < fromRight {
                currentRight[j] = fromLeft
            } else {
                currentRight[j] = fromRight
            }
        }
        prevLeft, prevRight = currentLeft, currentRight
    }
    if prevLeft[rightCount] < prevRight[rightCount] {
        return prevLeft[rightCount]
    }
    return prevRight[rightCount]
}

func main() {
    // Example 1:
    // Input: n = 6, start = 4, requests = [1,5]
    // Output: 6
    // Explanation:
    // Move from floor 4 (start) to floor 5 in 1 second. Penalty for floor 5 is 1.
    // Move from floor 5 to floor 1 in 4 seconds. Penalty for floor 1 is 5.
    // Thus, the total penalty is 1 + 5 = 6.
    fmt.Println(elevatorRequests(6, 4, []int{1,5})) // 6
    // Example 2:
    // Input: n = 8, start = 3, requests = [3,7,1]
    // Output: 10
    // Explanation:
    // Floor 3 (start) is fulfilled instantly. Penalty for floor 3 is 0.
    // Move from floor 3 to floor 1 in 2 seconds. Penalty for floor 1 is 2.
    // Move from floor 1 to floor 7 in 6 seconds. Penalty for floor 7 is 8.
    // Thus, the total penalty is 0 + 2 + 8 = 10.
    fmt.Println(elevatorRequests(8, 3, []int{3,7,1})) // 10
    // Example 3:
    // Input: n = 10, start = 5, requests = [0,2,9]
    // Output: 22
    // Explanation:
    // Move from floor 5 (start) to floor 2 in 3 seconds. Penalty for floor 2 is 3.
    // Move from floor 2 to floor 0 in 2 seconds. Penalty for floor 0 is 5.
    // Move from floor 0 to floor 9 in 9 seconds. Penalty for floor 9 is 14.
    // Thus, the total penalty is 3 + 5 + 14 = 22.
    fmt.Println(elevatorRequests(10, 5, []int{0,2,9})) // 22

    fmt.Println(elevatorRequests(6, 4, []int{1,2,3,4,5,6,7,8,9})) // 42
    fmt.Println(elevatorRequests(6, 4, []int{9,8,7,6,5,4,3,2,1})) // 42

    fmt.Println(elevatorRequests1(6, 4, []int{1,5})) // 6
    fmt.Println(elevatorRequests1(8, 3, []int{3,7,1})) // 10
    fmt.Println(elevatorRequests1(10, 5, []int{0,2,9})) // 22
    fmt.Println(elevatorRequests1(6, 4, []int{1,2,3,4,5,6,7,8,9})) // 42
    fmt.Println(elevatorRequests1(6, 4, []int{9,8,7,6,5,4,3,2,1})) // 42
}