package main

// 1345. Jump Game IV
// Given an array of integers arr, you are initially positioned at the first index of the array.

// In one step you can jump from index i to index:
//     i + 1 where: i + 1 < arr.length.
//     i - 1 where: i - 1 >= 0.
//     j where: arr[i] == arr[j] and i != j.

// Return the minimum number of steps to reach the last index of the array.

// Notice that you can not jump outside of the array at any time.

// Example 1:
// Input: arr = [100,-23,-23,404,100,23,23,23,3,404]
// Output: 3
// Explanation: You need three jumps from index 0 --> 4 --> 3 --> 9. Note that index 9 is the last index of the array.

// Example 2:
// Input: arr = [7]
// Output: 0
// Explanation: Start index is the last index. You do not need to jump.

// Example 3:
// Input: arr = [7,6,9,6,9,6,9,7]
// Output: 1
// Explanation: You can jump directly from index 0 to index 7 which is last index of the array.

// Constraints:
//     1 <= arr.length <= 5 * 10^4
//     -10^8 <= arr[i] <= 10^8

import "fmt"

func minJumps(arr []int) int {
    mp := make(map[int][]int)
    for i, v:=range arr {
        mp[v] = append(mp[v], i)
    }
    numVisited, visited, step := make(map[int]bool), make([]bool, len(arr)), 0
    queue := []int{ 0 }
    visited[0] = true
    for len(queue) > 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            cur := queue[i]
            if cur == len(arr) - 1 { return step }
            /* if the current number has been visited, then all the same number are already in queue */
            if !numVisited[arr[cur]] {
                numVisited[arr[cur]] = true
                for _, j := range mp[arr[cur]] {
                    if j == cur { continue }
                    queue = append(queue, j)
                    visited[j] = true
                }
            }
            if cur -1 >= 0 && !visited[cur-1] {
                queue = append(queue, cur-1)
                visited[cur-1] = true
            }
            if cur + 1 <len(arr) && !visited[cur+1] {
                queue = append(queue, cur+1)
                visited[cur+1] = true
            }
        }
        queue = queue[size:]
        step++
    }
    return -1
}

func minJumps1(arr []int) int {
    n := len(arr)
    mp := make(map[int][]int)
    for i, v := range arr { // 应付等值的情况
        mp[v] = append(mp[v], i)
    }
    queue, visit, step := []int{ 0 }, make([]bool, n), 0
    visit[0] = true
    for len(queue) > 0 {
        lev := make([]int, 0)
        for _, v := range queue {
            if v == n-1 { return step }
            visit[v] = true
            if v - 1 >= 0 && !visit[v - 1] {
                visit[v - 1] = true
                lev = append(lev, v - 1)
            }
            if v + 1 < n && !visit[v + 1] {
                visit[v + 1] = true
                lev = append(lev, v + 1)
            }
            // 等值
            for _, next:= range mp[arr[v]] {
                if !visit[next] {
                    visit[next] = true
                    lev = append(lev, next)
                }
            }
            delete(mp,arr[v])
        }
        if len(lev) > 0 {
            step++
        }
        queue = lev
    }
    return -1
}

func main() {
    // Example 1:
    // Input: arr = [100,-23,-23,404,100,23,23,23,3,404]
    // Output: 3
    // Explanation: You need three jumps from index 0 --> 4 --> 3 --> 9. Note that index 9 is the last index of the array.
    fmt.Println(minJumps([]int{100,-23,-23,404,100,23,23,23,3,404})) // 3
    // Example 2:
    // Input: arr = [7]
    // Output: 0
    // Explanation: Start index is the last index. You do not need to jump.
    fmt.Println(minJumps([]int{7})) // 0
    // Example 3:
    // Input: arr = [7,6,9,6,9,6,9,7]
    // Output: 1
    // Explanation: You can jump directly from index 0 to index 7 which is last index of the array.
    fmt.Println(minJumps([]int{7,6,9,6,9,6,9,7})) // 1

    fmt.Println(minJumps1([]int{100,-23,-23,404,100,23,23,23,3,404})) // 3
    fmt.Println(minJumps1([]int{7})) // 0
    fmt.Println(minJumps1([]int{7,6,9,6,9,6,9,7})) // 1
}