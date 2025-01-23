package main

// 2944. Minimum Number of Coins for Fruits
// You are given an 1-indexed integer array prices where prices[i] denotes the number of coins needed to purchase the ith fruit.

// The fruit market has the following reward for each fruit:
//     If you purchase the ith fruit at prices[i] coins, you can get any number of the next i fruits for free.

// Note that even if you can take fruit j for free, you can still purchase it for prices[j] coins to receive its reward.

// Return the minimum number of coins needed to acquire all the fruits.

// Example 1:
// Input: prices = [3,1,2]
// Output: 4
// Explanation:
// Purchase the 1st fruit with prices[0] = 3 coins, you are allowed to take the 2nd fruit for free.
// Purchase the 2nd fruit with prices[1] = 1 coin, you are allowed to take the 3rd fruit for free.
// Take the 3rd fruit for free.
// Note that even though you could take the 2nd fruit for free as a reward of buying 1st fruit, you purchase it to receive its reward, which is more optimal.

// Example 2:
// Input: prices = [1,10,1,1]
// Output: 2
// Explanation:
// Purchase the 1st fruit with prices[0] = 1 coin, you are allowed to take the 2nd fruit for free.
// Take the 2nd fruit for free.
// Purchase the 3rd fruit for prices[2] = 1 coin, you are allowed to take the 4th fruit for free.
// Take the 4th fruit for free.

// Example 3:
// Input: prices = [26,18,6,12,49,7,45,45]
// Output: 39
// Explanation:
// Purchase the 1st fruit with prices[0] = 26 coin, you are allowed to take the 2nd fruit for free.
// Take the 2nd fruit for free.
// Purchase the 3rd fruit for prices[2] = 6 coin, you are allowed to take the 4th, 5th and 6th (the next three) fruits for free.
// Take the 4th fruit for free.
// Take the 5th fruit for free.
// Purchase the 6th fruit with prices[5] = 7 coin, you are allowed to take the 8th and 9th fruit for free.
// Take the 7th fruit for free.
// Take the 8th fruit for free.
// Note that even though you could take the 6th fruit for free as a reward of buying 3rd fruit, you purchase it to receive its reward, which is more optimal.

// Constraints:
//     1 <= prices.length <= 1000
//     1 <= prices[i] <= 10^5

import "fmt"
import "slices"

func minimumCoins(prices []int) int {
    n := len(prices)
    dp := make([]int,n + 1)
    for i := 0; i < n; i++ {
        dp[i] = 1 << 31
    }
    dp[n] = 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := n - 1; i >= 0; i-- {
        for j := i + 1; j < min(i + i + 3, n + 1); j++ {
            dp[i] = min(dp[i], dp[j] + prices[i])
        }
    }
    return dp[0]
}

func minimumCoins1(prices []int) int {
    n := len(prices)
    for i := (n + 1) / 2 - 1; i > 0; i-- {
        prices[i-1] += slices.Min(prices[i : i * 2 + 1])
    }
    return prices[0]
}

func main() {
    // Example 1:
    // Input: prices = [3,1,2]
    // Output: 4
    // Explanation:
    // Purchase the 1st fruit with prices[0] = 3 coins, you are allowed to take the 2nd fruit for free.
    // Purchase the 2nd fruit with prices[1] = 1 coin, you are allowed to take the 3rd fruit for free.
    // Take the 3rd fruit for free.
    // Note that even though you could take the 2nd fruit for free as a reward of buying 1st fruit, you purchase it to receive its reward, which is more optimal.
    fmt.Println(minimumCoins([]int{3,1,2})) // 4
    // Example 2:
    // Input: prices = [1,10,1,1]
    // Output: 2
    // Explanation:
    // Purchase the 1st fruit with prices[0] = 1 coin, you are allowed to take the 2nd fruit for free.
    // Take the 2nd fruit for free.
    // Purchase the 3rd fruit for prices[2] = 1 coin, you are allowed to take the 4th fruit for free.
    // Take the 4th fruit for free.
    fmt.Println(minimumCoins([]int{1,10,1,1})) // 2
    // Example 3:
    // Input: prices = [26,18,6,12,49,7,45,45]
    // Output: 39
    // Explanation:
    // Purchase the 1st fruit with prices[0] = 26 coin, you are allowed to take the 2nd fruit for free.
    // Take the 2nd fruit for free.
    // Purchase the 3rd fruit for prices[2] = 6 coin, you are allowed to take the 4th, 5th and 6th (the next three) fruits for free.
    // Take the 4th fruit for free.
    // Take the 5th fruit for free.
    // Purchase the 6th fruit with prices[5] = 7 coin, you are allowed to take the 8th and 9th fruit for free.
    // Take the 7th fruit for free.
    // Take the 8th fruit for free.
    // Note that even though you could take the 6th fruit for free as a reward of buying 3rd fruit, you purchase it to receive its reward, which is more optimal.
    fmt.Println(minimumCoins([]int{26,18,6,12,49,7,45,45})) // 39

    fmt.Println(minimumCoins([]int{1,2,3,4,5,6,7,8,9})) // 8
    fmt.Println(minimumCoins([]int{9,8,7,6,5,4,3,2,1})) // 19

    fmt.Println(minimumCoins1([]int{3,1,2})) // 4
    fmt.Println(minimumCoins1([]int{1,10,1,1})) // 2
    fmt.Println(minimumCoins1([]int{26,18,6,12,49,7,45,45})) // 39
    fmt.Println(minimumCoins1([]int{1,2,3,4,5,6,7,8,9})) // 8
    fmt.Println(minimumCoins1([]int{9,8,7,6,5,4,3,2,1})) // 19
}