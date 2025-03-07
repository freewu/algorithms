package main

// 2279. Maximum Bags With Full Capacity of Rocks
// You have n bags numbered from 0 to n - 1. 
// You are given two 0-indexed integer arrays capacity and rocks. 
// The ith bag can hold a maximum of capacity[i] rocks and currently contains rocks[i] rocks. 
// You are also given an integer additionalRocks, the number of additional rocks you can place in any of the bags.

// Return the maximum number of bags that could have full capacity after placing the additional rocks in some bags.

// Example 1:
// Input: capacity = [2,3,4,5], rocks = [1,2,4,4], additionalRocks = 2
// Output: 3
// Explanation:
// Place 1 rock in bag 0 and 1 rock in bag 1.
// The number of rocks in each bag are now [2,3,4,4].
// Bags 0, 1, and 2 have full capacity.
// There are 3 bags at full capacity, so we return 3.
// It can be shown that it is not possible to have more than 3 bags at full capacity.
// Note that there may be other ways of placing the rocks that result in an answer of 3.

// Example 2:
// Input: capacity = [10,2,2], rocks = [2,2,0], additionalRocks = 100
// Output: 3
// Explanation:
// Place 8 rocks in bag 0 and 2 rocks in bag 2.
// The number of rocks in each bag are now [10,2,2].
// Bags 0, 1, and 2 have full capacity.
// There are 3 bags at full capacity, so we return 3.
// It can be shown that it is not possible to have more than 3 bags at full capacity.
// Note that we did not use all of the additional rocks.

// Constraints:
//     n == capacity.length == rocks.length
//     1 <= n <= 5 * 10^4
//     1 <= capacity[i] <= 10^9
//     0 <= rocks[i] <= capacity[i]
//     1 <= additionalRocks <= 10^9

import "fmt"
import "sort"

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
    cost := capacity
    for i := range cost {
        cost[i] -= rocks[i]
    }
    sort.Ints(cost)
    rest := additionalRocks
    for i := range cost {
        if cost[i] > rest {
            return i
        }
        rest -= cost[i]
    }
    return len(capacity)
}

func maximumBags1(capacity []int, rocks []int, additionalRocks int) int {
    for i := range capacity {
        capacity[i] -= rocks[i]
    }
    res := 0
    sort.Ints(capacity)
    for _, v := range capacity {
        if v > additionalRocks { break }
        res++
        additionalRocks -= v
    }
    return res
}

func maximumBags2(capacity []int, rocks []int, additionalRocks int) int {
    res, n := 0, len(capacity)
    can := make([]int, n)
    for i := range can {
        can[i] = capacity[i] - rocks[i]
    }
    sort.Ints(can)
    for i := 0; i < n; i++ {
        if can[i] == 0 {
            res++
        } else if can[i] <= additionalRocks {
            res++
            additionalRocks -= can[i]
        } else if can[i] > additionalRocks {
            i++
        } else {
            break
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: capacity = [2,3,4,5], rocks = [1,2,4,4], additionalRocks = 2
    // Output: 3
    // Explanation:
    // Place 1 rock in bag 0 and 1 rock in bag 1.
    // The number of rocks in each bag are now [2,3,4,4].
    // Bags 0, 1, and 2 have full capacity.
    // There are 3 bags at full capacity, so we return 3.
    // It can be shown that it is not possible to have more than 3 bags at full capacity.
    // Note that there may be other ways of placing the rocks that result in an answer of 3.
    fmt.Println(maximumBags([]int{2,3,4,5}, []int{1,2,4,4}, 2)) // 3
    // Example 2:
    // Input: capacity = [10,2,2], rocks = [2,2,0], additionalRocks = 100
    // Output: 3
    // Explanation:
    // Place 8 rocks in bag 0 and 2 rocks in bag 2.
    // The number of rocks in each bag are now [10,2,2].
    // Bags 0, 1, and 2 have full capacity.
    // There are 3 bags at full capacity, so we return 3.
    // It can be shown that it is not possible to have more than 3 bags at full capacity.
    // Note that we did not use all of the additional rocks.
    fmt.Println(maximumBags([]int{10,2,2}, []int{2,2,0}, 100)) // 3

    fmt.Println(maximumBags1([]int{2,3,4,5}, []int{1,2,4,4}, 2)) // 3
    fmt.Println(maximumBags1([]int{10,2,2}, []int{2,2,0}, 100)) // 3

    fmt.Println(maximumBags2([]int{2,3,4,5}, []int{1,2,4,4}, 2)) // 3
    fmt.Println(maximumBags2([]int{10,2,2}, []int{2,2,0}, 100)) // 3
}