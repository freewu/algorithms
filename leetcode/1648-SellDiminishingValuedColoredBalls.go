package main

// 1648. Sell Diminishing-Valued Colored Balls
// You have an inventory of different colored balls, and there is a customer that wants orders balls of any color.

// The customer weirdly values the colored balls. 
// Each colored ball's value is the number of balls of that color you currently have in your inventory. 
// For example, if you own 6 yellow balls, the customer would pay 6 for the first yellow ball. 
// After the transaction, there are only 5 yellow balls left, so the next yellow ball is then valued at 5 (i.e., the value of the balls decreases as you sell more to the customer).

// You are given an integer array, inventory, where inventory[i] represents the number of balls of the ith color that you initially own. 
// You are also given an integer orders, which represents the total number of balls that the customer wants. 
// You can sell the balls in any order.

// Return the maximum total value that you can attain after selling orders colored balls. 
// As the answer may be too large, return it modulo 10^9 + 7.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/11/05/jj.gif" />
// Input: inventory = [2,5], orders = 4
// Output: 14
// Explanation: Sell the 1st color 1 time (2) and the 2nd color 3 times (5 + 4 + 3).
// The maximum total value is 2 + 5 + 4 + 3 = 14.

// Example 2:
// Input: inventory = [3,5], orders = 6
// Output: 19
// Explanation: Sell the 1st color 2 times (3 + 2) and the 2nd color 4 times (5 + 4 + 3 + 2).
// The maximum total value is 3 + 2 + 5 + 4 + 3 + 2 = 19.

// Constraints:
//     1 <= inventory.length <= 10^5
//     1 <= inventory[i] <= 10^9
//     1 <= orders <= min(sum(inventory[i]), 10^9)

import "fmt"
import "sort"

func maxProfit(inventory []int, orders int) int {
    sort.Ints(inventory)
    inventory = append([]int{0}, inventory...)
    res, n, mod := 0, len(inventory), 1_000_000_007
    cur := inventory[n - 1]
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := n - 2; i >= 0 && orders > 0; i-- {
        occ := n - i - 1
        can := min(orders / occ, cur - inventory[i])
        n1, n2 := cur, cur - can + 1
        res = (res + ((n1 - n2 + 1) * (n1 + n2) / 2) * occ) % mod
        orders -= can * occ
        cur -= can
    }
    if orders > 0 {
        res = (res + orders * cur) % mod
    }
    return res % mod
}

func maxProfit1(inventory []int, orders int) int {
    mx := 0
    for _, v := range inventory {
        if v > mx { mx = v }
    }
    bound := sort.Search(mx + 1, func(i int) bool {
        res := 0
        for _, v := range inventory {
            if v - i > 0 {
                res += v - i
            }     
        }
        return res <= orders
    })
    res, count, mod := 0, 0, 1_000_000_007
    for _, v := range inventory {
        if v - bound <= 0 { continue }
        tmp := 0
        if (v - bound) % 2 == 0 {
            tmp = (v + bound + 1) % mod * (v - bound) / 2 % mod
        } else {
            tmp = (v + bound + 1) / 2 % mod * (v - bound) % mod
        }           
        res = (res + tmp) % mod
        count += v - bound
    }
    res = (res + bound % mod * (orders - count) ) % mod
    return res
}

func maxProfit2(inventory []int, orders int) int {
    sum := func(lower int) int {
        res := 0
        for _, v := range inventory {
            if v > lower {
                res += v - lower
            }
        }
        return res
    }
    res, mod := 0, 1_000_000_007
    bound := sort.Search(mod, func(i int) bool {
        return sum(i) <= orders
    })
    for _, v := range inventory {
        if v > bound {
            orders -= v - bound 
            res = (res + (v + bound + 1) * (v - bound)/2) % mod
        }
    }
    res = (res + orders * bound) % mod
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/11/05/jj.gif" />
    // Input: inventory = [2,5], orders = 4
    // Output: 14
    // Explanation: Sell the 1st color 1 time (2) and the 2nd color 3 times (5 + 4 + 3).
    // The maximum total value is 2 + 5 + 4 + 3 = 14.
    fmt.Println(maxProfit([]int{2,5}, 4)) // 14
    // Example 2:
    // Input: inventory = [3,5], orders = 6
    // Output: 19
    // Explanation: Sell the 1st color 2 times (3 + 2) and the 2nd color 4 times (5 + 4 + 3 + 2).
    // The maximum total value is 3 + 2 + 5 + 4 + 3 + 2 = 19.
    fmt.Println(maxProfit([]int{3,5}, 6)) // 19

    fmt.Println(maxProfit([]int{2,8,4,10,6}, 20)) // 110
    fmt.Println(maxProfit([]int{1000000000}, 1000000000)) // 21

    fmt.Println(maxProfit1([]int{2,5}, 4)) // 14
    fmt.Println(maxProfit1([]int{3,5}, 6)) // 19
    fmt.Println(maxProfit1([]int{2,8,4,10,6}, 20)) // 110
    fmt.Println(maxProfit1([]int{1000000000}, 1000000000)) // 21

    fmt.Println(maxProfit2([]int{2,5}, 4)) // 14
    fmt.Println(maxProfit2([]int{3,5}, 6)) // 19
    fmt.Println(maxProfit2([]int{2,8,4,10,6}, 20)) // 110
    fmt.Println(maxProfit2([]int{1000000000}, 1000000000)) // 21
}