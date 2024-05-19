package main

// 1535. Find the Winner of an Array Game
// Given an integer array arr of distinct integers and an integer k.
// A game will be played between the first two elements of the array (i.e. arr[0] and arr[1]). 
// In each round of the game, we compare arr[0] with arr[1], the larger integer wins and remains at position 0, 
// and the smaller integer moves to the end of the array. 
// The game ends when an integer wins k consecutive rounds.

// Return the integer which will win the game.
// It is guaranteed that there will be a winner of the game.

// Example 1:
// Input: arr = [2,1,3,5,4,6,7], k = 2
// Output: 5
// Explanation: Let's see the rounds of the game:
// Round |       arr       | winner | win_count
//   1   | [2,1,3,5,4,6,7] | 2      | 1
//   2   | [2,3,5,4,6,7,1] | 3      | 1
//   3   | [3,5,4,6,7,1,2] | 5      | 1
//   4   | [5,4,6,7,1,2,3] | 5      | 2
// So we can see that 4 rounds will be played and 5 is the winner because it wins 2 consecutive games.

// Example 2:
// Input: arr = [3,2,1], k = 10
// Output: 3
// Explanation: 3 will win the first 10 rounds consecutively.

// Constraints:
//     2 <= arr.length <= 10^5
//     1 <= arr[i] <= 10^6
//     arr contains distinct integers.
//     1 <= k <= 10^9

import "fmt"

func getWinner(arr []int, k int) int {
    count, leaderPos := 0, 0
    for i := 0; i < len(arr)-1; i++ {
        if count == k { // 当一个整数赢得 k 个连续回合时，游戏结束，该整数就是比赛的 赢家
            return arr[leaderPos]
        }
        if arr[leaderPos] > arr[i+1] { // 较大的整数将会取得这一回合的胜利并保留在位置 0 ，较小的整数移至数组的末尾
            count++
        } else {
            leaderPos = i + 1
            count = 1
        }
    }
    return arr[leaderPos]
}

func getWinner1(arr []int, k int) int {
    winner,cnt := arr[0], 0
    for i := 1; i < len(arr); i++ {
        curr := arr[i]
        if winner > curr {
            cnt++
        } else {
            winner = curr
            cnt = 1
        }
        if cnt == k {
            break
        }
    }
    return winner
}

func getWinner2(arr []int, k int) int {
    count, winner := 0, arr[0] // Initialize consecutive win count and current winner
    for i := 1; i < len(arr); i++ { // Simulate the game rounds
        if arr[i] > winner {
            winner = arr[i]
            count = 1
        } else {
            count++
        }
        if count == k { // Check if the current winner has won k consecutive rounds
            return winner
        }
    }
    return winner // If k consecutive wins are not achieved, return the current winner
}

func main() {
    // Example 1:
    // Input: arr = [2,1,3,5,4,6,7], k = 2
    // Output: 5
    // Explanation: Let's see the rounds of the game:
    // Round |       arr       | winner | win_count
    //   1   | [2,1,3,5,4,6,7] | 2      | 1
    //   2   | [2,3,5,4,6,7,1] | 3      | 1
    //   3   | [3,5,4,6,7,1,2] | 5      | 1
    //   4   | [5,4,6,7,1,2,3] | 5      | 2
    // So we can see that 4 rounds will be played and 5 is the winner because it wins 2 consecutive games.
    fmt.Println(getWinner([]int{2,1,3,5,4,6,7}, 2)) // 5
    // Example 2:
    // Input: arr = [3,2,1], k = 10
    // Output: 3
    // Explanation: 3 will win the first 10 rounds consecutively.
    fmt.Println(getWinner([]int{3,2,1}, 10)) // 3

    fmt.Println(getWinner1([]int{2,1,3,5,4,6,7}, 2)) // 5
    fmt.Println(getWinner1([]int{3,2,1}, 10)) // 3

    fmt.Println(getWinner2([]int{2,1,3,5,4,6,7}, 2)) // 5
    fmt.Println(getWinner2([]int{3,2,1}, 10)) // 3
}