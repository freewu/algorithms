package main

// 1040. Moving Stones Until Consecutive II
// There are some stones in different positions on the X-axis. 
// You are given an integer array stones, the positions of the stones.

// Call a stone an endpoint stone if it has the smallest or largest position. 
// In one move, you pick up an endpoint stone and move it to an unoccupied position so that it is no longer an endpoint stone.
//     In particular, if the stones are at say, stones = [1,2,5], 
//     you cannot move the endpoint stone at position 5, 
//     since moving it to any position (such as 0, or 3) will still keep that stone as an endpoint stone.

// The game ends when you cannot make any more moves (i.e., the stones are in three consecutive positions).

// Return an integer array answer of length 2 where:
//     answer[0] is the minimum number of moves you can play, and
//     answer[1] is the maximum number of moves you can play.

// Example 1:
// Input: stones = [7,4,9]
// Output: [1,2]
// Explanation: We can move 4 -> 8 for one move to finish the game.
// Or, we can move 9 -> 5, 4 -> 6 for two moves to finish the game.

// Example 2:
// Input: stones = [6,5,4,3,10]
// Output: [2,3]
// Explanation: We can move 3 -> 8 then 10 -> 7 to finish the game.
// Or, we can move 3 -> 7, 4 -> 8, 5 -> 9 to finish the game.
// Notice we cannot move 10 -> 2 to finish the game, because that would be an illegal move.

// Constraints:
//     3 <= stones.length <= 10^4
//     1 <= stones[i] <= 10^9
//     All the values of stones are unique.

import "fmt"
import "sort"

func numMovesStonesII(stones []int) []int {
    sort.Ints(stones)
    i, n := 0, len(stones)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    low, high := n, max(stones[n - 1] - n + 2 - stones[1], stones[n - 2] - stones[0] - n + 2)
    for j := 0; j < n; j++ {
        for stones[j] - stones[i] >= n { i++ }
        if j - i + 1 == n - 1 && stones[j] - stones[i] == n - 2 {
            low = min(low, 2)
        } else {
            low = min(low, n - (j - i + 1))
        }
    }
    return []int{ low, high }
}

func numMovesStonesII1(nums []int) []int {
    sort.Ints(nums)
    n := len(nums)
    if n <= 2 {  return []int{0, 0} }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    e1, e2 := nums[n-2] - nums[0] - (n - 2), nums[n-1] - nums[1] - (n - 2)
    maxMove := max(e2, e1)
    if e1 == 0 || e2 == 0 { return []int{ min(2, maxMove), maxMove } }
    maxCnt, left := 0, 0
    for ri, sr := range nums {
        for sr-nums[left]+1 > n {
            left++
        }
        maxCnt = max(maxCnt, ri - left + 1)
    }
    return []int{ n - maxCnt, maxMove }
}

func main() {
    // Example 1:
    // Input: stones = [7,4,9]
    // Output: [1,2]
    // Explanation: We can move 4 -> 8 for one move to finish the game.
    // Or, we can move 9 -> 5, 4 -> 6 for two moves to finish the game.
    fmt.Println(numMovesStonesII([]int{7,4,9})) // [1,2]
    // Example 2:
    // Input: stones = [6,5,4,3,10]
    // Output: [2,3]
    // Explanation: We can move 3 -> 8 then 10 -> 7 to finish the game.
    // Or, we can move 3 -> 7, 4 -> 8, 5 -> 9 to finish the game.
    // Notice we cannot move 10 -> 2 to finish the game, because that would be an illegal move.
    fmt.Println(numMovesStonesII([]int{6,5,4,3,10})) // [2,3]

    fmt.Println(numMovesStonesII1([]int{7,4,9})) // [1,2]
    fmt.Println(numMovesStonesII1([]int{6,5,4,3,10})) // [2,3]
}