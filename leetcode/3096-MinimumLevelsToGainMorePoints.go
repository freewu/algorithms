package main

// 3096. Minimum Levels to Gain More Points
// You are given a binary array possible of length n.

// Alice and Bob are playing a game that consists of n levels. 
// Some of the levels in the game are impossible to clear while others can always be cleared. 
// In particular, if possible[i] == 0, then the ith level is impossible to clear for both the players. 
// A player gains 1 point on clearing a level and loses 1 point if the player fails to clear it.

// At the start of the game, Alice will play some levels in the given order starting from the 0th level, 
// after which Bob will play for the rest of the levels.

// Alice wants to know the minimum number of levels she should play to gain more points than Bob, 
// if both players play optimally to maximize their points.

// Return the minimum number of levels Alice should play to gain more points. If this is not possible, return -1.
// Note that each player must play at least 1 level.

// Example 1:
// Input: possible = [1,0,1,0]
// Output: 1
// Explanation:
// Let's look at all the levels that Alice can play up to:
// If Alice plays only level 0 and Bob plays the rest of the levels, Alice has 1 point, while Bob has -1 + 1 - 1 = -1 point.
// If Alice plays till level 1 and Bob plays the rest of the levels, Alice has 1 - 1 = 0 points, while Bob has 1 - 1 = 0 points.
// If Alice plays till level 2 and Bob plays the rest of the levels, Alice has 1 - 1 + 1 = 1 point, while Bob has -1 point.
// Alice must play a minimum of 1 level to gain more points.

// Example 2:
// Input: possible = [1,1,1,1,1]
// Output: 3
// Explanation:
// Let's look at all the levels that Alice can play up to:
// If Alice plays only level 0 and Bob plays the rest of the levels, Alice has 1 point, while Bob has 4 points.
// If Alice plays till level 1 and Bob plays the rest of the levels, Alice has 2 points, while Bob has 3 points.
// If Alice plays till level 2 and Bob plays the rest of the levels, Alice has 3 points, while Bob has 2 points.
// If Alice plays till level 3 and Bob plays the rest of the levels, Alice has 4 points, while Bob has 1 point.
// Alice must play a minimum of 3 levels to gain more points.

// Example 3:
// Input: possible = [0,0]
// Output: -1
// Explanation:
// The only possible way is for both players to play 1 level each. 
// Alice plays level 0 and loses 1 point. 
// Bob plays level 1 and loses 1 point. 
// As both players have equal points, Alice can't gain more points than Bob.

// Constraints:
//     2 <= n == possible.length <= 10^5
//     possible[i] is either 0 or 1.

import "fmt"

func minimumLevels(possible []int) int {
    n := len(possible)
    if possible[0] == 0 { possible[0] = -1 } // 第一局是困难模式直接 game over
    for i := 1; i < n; i++ {
        if possible[i] == 1 {
            possible[i] = possible[i - 1] + 1 // 通过一个简单模式的关卡可以获得 1 分
        } else {
            possible[i] = possible[i - 1] - 1 // 通过困难模式的关卡将失去 1 分
        }
    }
    //fmt.Println(possible)
    for i := 0; i < n - 1; i++ {
        if possible[i] > possible[n - 1] - possible[i] { // possible[i] 最后得分
            return i + 1
        }
    }
    return -1
}

func minimumLevels1(possible []int) int {
    n, sum, pre := len(possible), 0, 0
    for _, v := range possible {
        sum += v
    }
    sum = sum * 2 - n
    for i, v := range possible[:n-1] {
        pre += v * 4 - 2
        if pre > sum {
            return i + 1
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: possible = [1,0,1,0]
    // Output: 1
    // Explanation:
    // Let's look at all the levels that Alice can play up to:
    // If Alice plays only level 0 and Bob plays the rest of the levels, Alice has 1 point, while Bob has -1 + 1 - 1 = -1 point.
    // If Alice plays till level 1 and Bob plays the rest of the levels, Alice has 1 - 1 = 0 points, while Bob has 1 - 1 = 0 points.
    // If Alice plays till level 2 and Bob plays the rest of the levels, Alice has 1 - 1 + 1 = 1 point, while Bob has -1 point.
    // Alice must play a minimum of 1 level to gain more points.
    fmt.Println(minimumLevels([]int{1,0,1,0})) // 1
    // Example 2:
    // Input: possible = [1,1,1,1,1]
    // Output: 3
    // Explanation:
    // Let's look at all the levels that Alice can play up to:
    // If Alice plays only level 0 and Bob plays the rest of the levels, Alice has 1 point, while Bob has 4 points.
    // If Alice plays till level 1 and Bob plays the rest of the levels, Alice has 2 points, while Bob has 3 points.
    // If Alice plays till level 2 and Bob plays the rest of the levels, Alice has 3 points, while Bob has 2 points.
    // If Alice plays till level 3 and Bob plays the rest of the levels, Alice has 4 points, while Bob has 1 point.
    // Alice must play a minimum of 3 levels to gain more points.
    fmt.Println(minimumLevels([]int{1,1,1,1,1})) // 3
    // Example 3:
    // Input: possible = [0,0]
    // Output: -1
    // Explanation:
    // The only possible way is for both players to play 1 level each. 
    // Alice plays level 0 and loses 1 point. 
    // Bob plays level 1 and loses 1 point. 
    // As both players have equal points, Alice can't gain more points than Bob.
    fmt.Println(minimumLevels([]int{0,0})) // -1

    fmt.Println(minimumLevels1([]int{1,0,1,0})) // 1
    fmt.Println(minimumLevels1([]int{1,1,1,1,1})) // 3
    fmt.Println(minimumLevels1([]int{0,0})) // -1
}