package main

// 1561. Maximum Number of Coins You Can Get
// There are 3n piles of coins of varying size, you and your friends will take piles of coins as follows:
//     In each step, you will choose any 3 piles of coins (not necessarily consecutive).
//     Of your choice, Alice will pick the pile with the maximum number of coins.
//     You will pick the next pile with the maximum number of coins.
//     Your friend Bob will pick the last pile.
//     Repeat until there are no more piles of coins.

// Given an array of integers piles where piles[i] is the number of coins in the ith pile.

// Return the maximum number of coins that you can have.

// Example 1:
// Input: piles = [2,4,1,2,7,8]
// Output: 9
// Explanation: Choose the triplet (2, 7, 8), Alice Pick the pile with 8 coins, you the pile with 7 coins and Bob the last one.
// Choose the triplet (1, 2, 4), Alice Pick the pile with 4 coins, you the pile with 2 coins and Bob the last one.
// The maximum number of coins which you can have are: 7 + 2 = 9.
// On the other hand if we choose this arrangement (1, 2, 8), (2, 4, 7) you only get 2 + 4 = 6 coins which is not optimal.

// Example 2:
// Input: piles = [2,4,5]
// Output: 4

// Example 3:
// Input: piles = [9,8,7,6,5,1,2,3,4]
// Output: 18

// Constraints:
//     3 <= piles.length <= 10^5
//     piles.length % 3 == 0
//     1 <= piles[i] <= 10^4

import "fmt"
import "sort"

func maxCoins(piles []int) int {
    sort.Ints(piles)
    res, n := 0, len(piles)
    for i := n - 2; i >= n / 3; i -= 2 {
        res += piles[i]
    }
    return res
}

func maxCoins1(piles []int) int {
    // 每次最小的配两个最大的
    sort.Ints(piles)
    res, i, j := 0, 0, len(piles) - 1
    for i < j {
        j--
        res += piles[j]
        j--
        i++
    }
    return res
}

func main() {
    // Example 1:
    // Input: piles = [2,4,1,2,7,8]
    // Output: 9
    // Explanation: Choose the triplet (2, 7, 8), Alice Pick the pile with 8 coins, you the pile with 7 coins and Bob the last one.
    // Choose the triplet (1, 2, 4), Alice Pick the pile with 4 coins, you the pile with 2 coins and Bob the last one.
    // The maximum number of coins which you can have are: 7 + 2 = 9.
    // On the other hand if we choose this arrangement (1, 2, 8), (2, 4, 7) you only get 2 + 4 = 6 coins which is not optimal.
    fmt.Println(maxCoins([]int{2,4,1,2,7,8})) // 9
    // Example 2:
    // Input: piles = [2,4,5]
    // Output: 4
    fmt.Println(maxCoins([]int{2,4,5})) // 4
    // Example 3:
    // Input: piles = [9,8,7,6,5,1,2,3,4]
    // Output: 18
    fmt.Println(maxCoins([]int{9,8,7,6,5,1,2,3,4})) // 18

    fmt.Println(maxCoins1([]int{2,4,1,2,7,8})) // 9
    fmt.Println(maxCoins1([]int{2,4,5})) // 4
    fmt.Println(maxCoins1([]int{9,8,7,6,5,1,2,3,4})) // 18
}