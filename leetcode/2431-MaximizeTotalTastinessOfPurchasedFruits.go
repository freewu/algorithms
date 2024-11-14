package main

// 2431. Maximize Total Tastiness of Purchased Fruits
// You are given two non-negative integer arrays price and tastiness, both arrays have the same length n. 
// You are also given two non-negative integers maxAmount and maxCoupons.

// For every integer i in range [0, n - 1]:
//     1. price[i] describes the price of ith fruit.
//     2. tastiness[i] describes the tastiness of ith fruit.

// You want to purchase some fruits such that total tastiness is maximized and the total price does not exceed maxAmount.

// Additionally, you can use a coupon to purchase fruit for half of its price (rounded down to the closest integer). 
// You can use at most maxCoupons of such coupons.

// Return the maximum total tastiness that can be purchased.

// Note that:
//     1. You can purchase each fruit at most once.
//     2. You can use coupons on some fruit at most once.

// Example 1:
// Input: price = [10,20,20], tastiness = [5,8,8], maxAmount = 20, maxCoupons = 1
// Output: 13
// Explanation: It is possible to make total tastiness 13 in following way:
// - Buy first fruit without coupon, so that total price = 0 + 10 and total tastiness = 0 + 5.
// - Buy second fruit with coupon, so that total price = 10 + 10 and total tastiness = 5 + 8.
// - Do not buy third fruit, so that total price = 20 and total tastiness = 13.
// It can be proven that 13 is the maximum total tastiness that can be obtained.

// Example 2:
// Input: price = [10,15,7], tastiness = [5,8,20], maxAmount = 10, maxCoupons = 2
// Output: 28
// Explanation: It is possible to make total tastiness 20 in following way:
// - Do not buy first fruit, so that total price = 0 and total tastiness = 0.
// - Buy second fruit with coupon, so that total price = 0 + 7 and total tastiness = 0 + 8.
// - Buy third fruit with coupon, so that total price = 7 + 3 and total tastiness = 8 + 20.
// It can be proven that 28 is the maximum total tastiness that can be obtained.

// Constraints:
//     n == price.length == tastiness.length
//     1 <= n <= 100
//     0 <= price[i], tastiness[i], maxAmount <= 1000
//     0 <= maxCoupons <= 5

import "fmt"

func maxTastiness(price []int, tastiness []int, maxAmount int, maxCoupons int) int {
    n := len(price)
    dp := make([][]int, maxAmount + 1)
    for i := range dp {
        dp[i] = make([]int, maxCoupons + 1)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ { // 遍历 tastiness 和 price
        t, p := tastiness[i], price[i]
        for j := maxAmount; j >= p / 2; j-- {// 覆盖所有情况，终点为最小能到的剩余预算
            for k := maxCoupons; k >= 0; k-- { // 覆盖所有情况，没有优惠券了也要算进去
                if j >= p && k > 0 {// 用不用优惠券都能买
                    dp[j][k] = max(dp[j][k], max(dp[j - p][k], dp[j - p / 2][k - 1]) + t) 
                } else if k > 0 { // 只能用优惠券买了
                    dp[j][k] = max(dp[j][k], dp[j - p / 2][k - 1] + t)
                } else if j >= p { // 没有优惠券了，要够钱才能买
                    dp[j][k] = max(dp[j][k], dp[j - p][k] + t)
                }
            }
        }
    }
    // 当买不起的时候，在买不起的范围内，预算多一点少一点，能买到的水果都是一样的。
    // 所以当少几块预算(maxAmount - ε)就能取到最大值的时候，再给几块预算到maxAmount还是买不起，他们都是最大值，直接取maxAmount即为答案。
    return dp[maxAmount][maxCoupons]
}

func main() {
    // Example 1:
    // Input: price = [10,20,20], tastiness = [5,8,8], maxAmount = 20, maxCoupons = 1
    // Output: 13
    // Explanation: It is possible to make total tastiness 13 in following way:
    // - Buy first fruit without coupon, so that total price = 0 + 10 and total tastiness = 0 + 5.
    // - Buy second fruit with coupon, so that total price = 10 + 10 and total tastiness = 5 + 8.
    // - Do not buy third fruit, so that total price = 20 and total tastiness = 13.
    // It can be proven that 13 is the maximum total tastiness that can be obtained.
    fmt.Println(maxTastiness([]int{10,20,20}, []int{5,8,8}, 20, 1)) // 13
    // Example 2:
    // Input: price = [10,15,7], tastiness = [5,8,20], maxAmount = 10, maxCoupons = 2
    // Output: 28
    // Explanation: It is possible to make total tastiness 20 in following way:
    // - Do not buy first fruit, so that total price = 0 and total tastiness = 0.
    // - Buy second fruit with coupon, so that total price = 0 + 7 and total tastiness = 0 + 8.
    // - Buy third fruit with coupon, so that total price = 7 + 3 and total tastiness = 8 + 20.
    // It can be proven that 28 is the maximum total tastiness that can be obtained.
    fmt.Println(maxTastiness([]int{10,15,7}, []int{5,8,20}, 10, 2)) // 28
}