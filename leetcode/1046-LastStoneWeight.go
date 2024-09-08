package main

// 1046. Last Stone Weight
// You are given an array of integers stones where stones[i] is the weight of the ith stone.

// We are playing a game with the stones. On each turn, 
// we choose the heaviest two stones and smash them together. 
// Suppose the heaviest two stones have weights x and y with x <= y. 
// The result of this smash is:
//     If x == y, both stones are destroyed, and
//     If x != y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.

// At the end of the game, there is at most one stone left.

// Return the weight of the last remaining stone. 
// If there are no stones left, return 0.

// Example 1:
// Input: stones = [2,7,4,1,8,1]
// Output: 1
// Explanation: 
// We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
// we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
// we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
// we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of the last stone.

// Example 2:
// Input: stones = [1]
// Output: 1

// Constraints:
//     1 <= stones.length <= 30
//     1 <= stones[i] <= 1000

import "fmt"
import "sort"

func lastStoneWeight(stones []int) int {
    sort.Ints(stones)
    for len(stones) > 1 {
        if stones[len(stones) - 1] == stones[len(stones) - 2] { // 如果 x == y，那么两块石头都会被完全粉碎
            stones = stones[:len(stones) - 2]
        } else { // 如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x
            stones[len(stones) - 2] = stones[len(stones) - 1] - stones[len(stones) - 2]
            stones = stones[:len(stones)-1]
            sort.Ints(stones) // 重新排序
        }
    }
    if len(stones) > 0 {
        return stones[0]
    }
    return 0
}

func main() {
    // Example 1:
    // Input: stones = [2,7,4,1,8,1]
    // Output: 1
    // Explanation: 
    // We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
    // we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
    // we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
    // we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of the last stone.
    fmt.Println(lastStoneWeight([]int{2,7,4,1,8,1})) // 1
    // Example 2:
    // Input: stones = [1]
    // Output: 1
    fmt.Println(lastStoneWeight([]int{1})) // 1
}