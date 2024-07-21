package main

// 1049. Last Stone Weight II
// You are given an array of integers stones where stones[i] is the weight of the ith stone.

// We are playing a game with the stones. On each turn, we choose any two stones and smash them together. 
// Suppose the stones have weights x and y with x <= y. The result of this smash is:
//     If x == y, both stones are destroyed, and
//     If x != y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.

// At the end of the game, there is at most one stone left.
// Return the smallest possible weight of the left stone. If there are no stones left, return 0.

// Example 1:
// Input: stones = [2,7,4,1,8,1]
// Output: 1
// Explanation:
// We can combine 2 and 4 to get 2, so the array converts to [2,7,1,8,1] then,
// we can combine 7 and 8 to get 1, so the array converts to [2,1,1,1] then,
// we can combine 2 and 1 to get 1, so the array converts to [1,1,1] then,
// we can combine 1 and 1 to get 0, so the array converts to [1], then that's the optimal value.

// Example 2:
// Input: stones = [31,26,33,21,40]
// Output: 5

// Constraints:
//     1 <= stones.length <= 30
//     1 <= stones[i] <= 100

import "fmt"

func lastStoneWeightII(stones []int) int {
    if len(stones) == 1 {
        return stones[0]
    }
    sum, dp := 0,make([]bool, 1501)
    dp[0] = true
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for _, v := range stones {
        sum += v
        for i := min(1500, sum); i >= v; i-- {
            dp[i] = dp[i] || dp[i - v]
        }
    }
    for i := sum / 2; i > 0; i-- {
        if dp[i] { 
            return sum - i - i 
        }
    }
    return 0
}

func lastStoneWeightII1(stones []int) int {
    sum := 0
    for _, v := range stones{
        sum += v
    }
    target := sum / 2
    dp := make([]int,target+1)
    for _, v := range stones {
        for j := target ;j >= v; j-- {
            if dp[j] < dp[j - v] + v {
                dp[j] = dp[j - v] + v
            }
        }
    }
    res := (target - dp[target])*2 
    if sum % 2 == 1 {
        res++
    }
    return res
}

func main() {
    // Example 1:
    // Input: stones = [2,7,4,1,8,1]
    // Output: 1
    // Explanation:
    // We can combine 2 and 4 to get 2, so the array converts to [2,7,1,8,1] then,
    // we can combine 7 and 8 to get 1, so the array converts to [2,1,1,1] then,
    // we can combine 2 and 1 to get 1, so the array converts to [1,1,1] then,
    // we can combine 1 and 1 to get 0, so the array converts to [1], then that's the optimal value.
    fmt.Println(lastStoneWeightII([]int{2,7,4,1,8,1})) // 1
    // Example 2:
    // Input: stones = [31,26,33,21,40]
    // Output: 5
    fmt.Println(lastStoneWeightII([]int{31,26,33,21,40})) // 5

    fmt.Println(lastStoneWeightII([]int{91})) // 91

    fmt.Println(lastStoneWeightII1([]int{2,7,4,1,8,1})) // 1
    fmt.Println(lastStoneWeightII1([]int{31,26,33,21,40})) // 5
    fmt.Println(lastStoneWeightII1([]int{91})) // 91
}