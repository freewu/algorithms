package main

// 3301. Maximize the Total Height of Unique Towers
// You are given an array maximumHeight, where maximumHeight[i] denotes the maximum height the ith tower can be assigned.

// Your task is to assign a height to each tower so that:
//     1. The height of the ith tower is a positive integer and does not exceed maximumHeight[i].
//     2. No two towers have the same height.

// Return the maximum possible total sum of the tower heights. 
// If it's not possible to assign heights, return -1.

// Example 1:
// Input: maximumHeight = [2,3,4,3]
// Output: 10
// Explanation:
// We can assign heights in the following way: [1, 2, 4, 3].

// Example 2:
// Input: maximumHeight = [15,10]
// Output: 25
// Explanation:
// We can assign heights in the following way: [15, 10].

// Example 3:
// Input: maximumHeight = [2,2,1]
// Output: -1
// Explanation:
// It's impossible to assign positive heights to each index so that no two towers have the same height.

// Constraints:
//     1 <= maximumHeight.length <= 10^5
//     1 <= maximumHeight[i] <= 10^9

import "fmt"
import "sort"

func maximumTotalSum(maximumHeight []int) int64 {
    sort.Ints(maximumHeight)
    res, last := maximumHeight[len(maximumHeight) - 1], maximumHeight[len(maximumHeight) - 1]
    for i := len(maximumHeight) -  2; i >= 0; i-- {
        if maximumHeight[i] >= last {
            last--
        } else {
            last = maximumHeight[i]
        }
        if last <= 0 {
            return -1
        }
        res += last
    }
    return int64(res)
}

func maximumTotalSum1(maximumHeight []int) int64 {
    res, last := 0, 1 << 31
    sort.Ints(maximumHeight)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := len(maximumHeight) - 1; i >= 0; i-- {
        cur := min(maximumHeight[i], last - 1)
        if cur < 1 {
            return -1
        }
        res += cur
        last = cur
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: maximumHeight = [2,3,4,3]
    // Output: 10
    // Explanation:
    // We can assign heights in the following way: [1, 2, 4, 3].
    fmt.Println(maximumTotalSum([]int{2,3,4,3})) // 10
    // Example 2:
    // Input: maximumHeight = [15,10]
    // Output: 25
    // Explanation:
    // We can assign heights in the following way: [15, 10].
    fmt.Println(maximumTotalSum([]int{15,10})) // 25
    // Example 3:
    // Input: maximumHeight = [2,2,1]
    // Output: -1
    // Explanation:
    // It's impossible to assign positive heights to each index so that no two towers have the same height.
    fmt.Println(maximumTotalSum([]int{2,2,1})) // -1

    fmt.Println(maximumTotalSum([]int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maximumTotalSum([]int{9,8,7,6,5,4,3,2,1})) // 45

    fmt.Println(maximumTotalSum1([]int{2,3,4,3})) // 10
    fmt.Println(maximumTotalSum1([]int{15,10})) // 25
    fmt.Println(maximumTotalSum1([]int{2,2,1})) // -1
    fmt.Println(maximumTotalSum1([]int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maximumTotalSum1([]int{9,8,7,6,5,4,3,2,1})) // 45
}