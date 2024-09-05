package main

// 1217. Minimum Cost to Move Chips to The Same Position
// We have n chips, where the position of the ith chip is position[i].

// We need to move all the chips to the same position. 
// In one step, we can change the position of the ith chip from position[i] to:
//     position[i] + 2 or position[i] - 2 with cost = 0.
//     position[i] + 1 or position[i] - 1 with cost = 1.

// Return the minimum cost needed to move all the chips to the same position.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/08/15/chips_e1.jpg" />
// Input: position = [1,2,3]
// Output: 1
// Explanation: First step: Move the chip at position 3 to position 1 with cost = 0.
// Second step: Move the chip at position 2 to position 1 with cost = 1.
// Total cost is 1.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/08/15/chip_e2.jpg" />
// Input: position = [2,2,2,3,3]
// Output: 2
// Explanation: We can move the two chips at position  3 to position 2. Each move has cost = 1. The total cost = 2.

// Example 3:
// Input: position = [1,1000000000]
// Output: 1

// Constraints:
//     1 <= position.length <= 100
//     1 <= position[i] <= 10^9

import "fmt"

func minCostToMoveChips(position []int) int {
    res := [2]int{} // [even, odd]
    for _, v := range position {
        res[v % 2]++
    }
    if res[0] < res[1] {
        return res[0]
    }
    return res[1]
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/08/15/chips_e1.jpg" />
    // Input: position = [1,2,3]
    // Output: 1
    // Explanation: First step: Move the chip at position 3 to position 1 with cost = 0.
    // Second step: Move the chip at position 2 to position 1 with cost = 1.
    // Total cost is 1.
    fmt.Println(minCostToMoveChips([]int{1,2,3})) // 1
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/08/15/chip_e2.jpg" />
    // Input: position = [2,2,2,3,3]
    // Output: 2
    // Explanation: We can move the two chips at position  3 to position 2. Each move has cost = 1. The total cost = 2.
    fmt.Println(minCostToMoveChips([]int{2,2,2,3,3})) // 2
    // Example 3:
    // Input: position = [1,1000000000]
    // Output: 1
    fmt.Println(minCostToMoveChips([]int{1,1000000000})) // 1
}