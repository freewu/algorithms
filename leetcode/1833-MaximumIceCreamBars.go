package main

// 1833. Maximum Ice Cream Bars
// It is a sweltering summer day, and a boy wants to buy some ice cream bars.

// At the store, there are n ice cream bars. 
// You are given an array costs of length n, where costs[i] is the price of the ith ice cream bar in coins. 
// The boy initially has coins coins to spend, and he wants to buy as many ice cream bars as possible. 

// Note: The boy can buy the ice cream bars in any order.

// Return the maximum number of ice cream bars the boy can buy with coins coins.

// You must solve the problem by counting sort.

// Example 1:
// Input: costs = [1,3,2,4,1], coins = 7
// Output: 4
// Explanation: The boy can buy ice cream bars at indices 0,1,2,4 for a total price of 1 + 3 + 2 + 1 = 7.

// Example 2:
// Input: costs = [10,6,8,7,7,8], coins = 5
// Output: 0
// Explanation: The boy cannot afford any of the ice cream bars.

// Example 3:
// Input: costs = [1,6,3,1,2,5], coins = 20
// Output: 6
// Explanation: The boy can buy all the ice cream bars for a total price of 1 + 6 + 3 + 1 + 2 + 5 = 18.

// Constraints:
//     costs.length == n
//     1 <= n <= 10^5
//     1 <= costs[i] <= 10^5
//     1 <= coins <= 10^8

import "fmt"
import "sort"
import "slices"

func maxIceCream(costs []int, coins int) int {
    sort.Ints(costs)
    res, n := 0, len(costs)
    for i := 0; i < n; i++ { // 价格最小的买起
        if costs[i] > coins { break } // 金币不足
        coins -= costs[i]
        res++
    }
    return res
}

// O(n)
func maxIceCream1(costs []int, coins int) int {
    res, mx := 0, 100_000
    freq := make([]int, mx + 1)
    for _, v := range costs {
        freq[v]++
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i <= mx && coins >= i; i++ {
        if freq[i] > 0 {
            c := min(freq[i], coins / i)
            res += c 
            coins -= (i * c) 
        }
    }
    return res
}

func maxIceCream2(costs []int, coins int) int {
    res, freq := 0, make([]int, slices.Max(costs) + 1)
    for _, v := range costs {
        freq[v]++
    }
    for i, v := range freq {
        if v > 0 {
            if i > coins { return res }
            if coins/i >= v {
                coins -= (i * v)
                res += v
            } else {
                canBuy := coins / i
                coins -= (canBuy * i)
                res += canBuy
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: costs = [1,3,2,4,1], coins = 7
    // Output: 4
    // Explanation: The boy can buy ice cream bars at indices 0,1,2,4 for a total price of 1 + 3 + 2 + 1 = 7.
    fmt.Println(maxIceCream([]int{1,3,2,4,1}, 7)) // 4
    // Example 2:
    // Input: costs = [10,6,8,7,7,8], coins = 5
    // Output: 0
    // Explanation: The boy cannot afford any of the ice cream bars.
    fmt.Println(maxIceCream([]int{10,6,8,7,7,8}, 5)) // 0
    // Example 3:
    // Input: costs = [1,6,3,1,2,5], coins = 20
    // Output: 6
    // Explanation: The boy can buy all the ice cream bars for a total price of 1 + 6 + 3 + 1 + 2 + 5 = 18.
    fmt.Println(maxIceCream([]int{1,6,3,1,2,5}, 20)) // 6

    fmt.Println(maxIceCream1([]int{1,3,2,4,1}, 7)) // 4
    fmt.Println(maxIceCream1([]int{10,6,8,7,7,8}, 5)) // 0
    fmt.Println(maxIceCream1([]int{1,6,3,1,2,5}, 20)) // 6

    fmt.Println(maxIceCream2([]int{1,3,2,4,1}, 7)) // 4
    fmt.Println(maxIceCream2([]int{10,6,8,7,7,8}, 5)) // 0
    fmt.Println(maxIceCream2([]int{1,6,3,1,2,5}, 20)) // 6
}