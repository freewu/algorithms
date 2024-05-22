package main

// 335. Self Crossing
// You are given an array of integers distance.
// You start at the point (0, 0) on an X-Y plane, and you move distance[0] meters to the north, then distance[1] meters to the west, distance[2] meters to the south, distance[3] meters to the east, and so on. 
// In other words, after each move, your direction changes counter-clockwise.

// Return true if your path crosses itself or false if it does not.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/12/21/11.jpg" />
// Input: distance = [2,1,1,2]
// Output: true
// Explanation: The path crosses itself at the point (0, 1).

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/12/21/22.jpg" />
// Input: distance = [1,2,3,4]
// Output: false
// Explanation: The path does not cross itself at any point.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2022/12/21/33.jpg" />
// Input: distance = [1,1,1,2,1]
// Output: true
// Explanation: The path crosses itself at the point (0, 0).
 
// Constraints:
//     1 <= distance.length <= 10^5
//     1 <= distance[i] <= 10^5

import "fmt"

func isSelfCrossing(distance []int) bool {
    distance = append([]int{0, 0, 0, 0}, distance...)
    n, i := len(distance), 4
    for i < n && distance[i] > distance[i-2] {
        i++
    }
    if i == n {
        return false
    }
    if distance[i] >= distance[i-2]-distance[i-4] {
        distance[i-1] -= distance[i-3]
    }
    i = i + 1
    for i < n && distance[i] < distance[i-2] {
        i++
    }
    return i != n
}

func isSelfCrossing1(distance []int) bool {
    n := len(distance)
    for i := 3; i < n; i++ {
        if distance[i] >= distance[i-2] && distance[i-1] <= distance[i-3] {
            return true
        }
        if i < 4 {
            continue
        }
        if distance[i-1] == distance[i-3] && distance[i]+distance[i-4] >= distance[i-2] {
            return true
        }
        if i < 5 {
            continue
        }
        if distance[i]   >= distance[i-2] - distance[i-4] && 
           distance[i-1] >= distance[i-3] - distance[i-5] && 
           distance[i-1] <= distance[i-3] && 
           distance[i-2] >= distance[i-4] && 
           distance[i-3] >= distance[i-5] {
            return true
        }
    }
    return false
}

func isSelfCrossing2(d []int) bool {
    for i := 3; i < len(d); i++ {
        if d[i] >= d[i-2] && d[i-1] <= d[i-3] {
            return true
        }
        if i >= 4 && d[i-1] == d[i-3] && d[i]+d[i-4] >= d[i-2] {
            return true
        }
        if i >= 5 && d[i-2] >= d[i-4] && d[i-1] <= d[i-3] && d[i] >= d[i-2]-d[i-4] && d[i-1]+d[i-5] >= d[i-3] {
            return true
        }
    }
    return false
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/12/21/11.jpg" />
    // Input: distance = [2,1,1,2]
    // Output: true
    // Explanation: The path crosses itself at the point (0, 1).
    fmt.Println(isSelfCrossing([]int{ 2,1,1,2 })) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/12/21/22.jpg" />
    // Input: distance = [1,2,3,4]
    // Output: false
    // Explanation: The path does not cross itself at any point.
    fmt.Println(isSelfCrossing([]int{ 1,2,3,4 })) // false
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2022/12/21/33.jpg" />
    // Input: distance = [1,1,1,2,1]
    // Output: true
    // Explanation: The path crosses itself at the point (0, 0).
    fmt.Println(isSelfCrossing([]int{ 1,1,1,2,1 })) // true

    fmt.Println(isSelfCrossing1([]int{ 2,1,1,2 })) // true
    fmt.Println(isSelfCrossing1([]int{ 1,2,3,4 })) // false
    fmt.Println(isSelfCrossing1([]int{ 1,1,1,2,1 })) // true

    fmt.Println(isSelfCrossing2([]int{ 2,1,1,2 })) // true
    fmt.Println(isSelfCrossing2([]int{ 1,2,3,4 })) // false
    fmt.Println(isSelfCrossing2([]int{ 1,1,1,2,1 })) // true
}