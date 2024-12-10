package main

// 2240. Number of Ways to Buy Pens and Pencils
// You are given an integer total indicating the amount of money you have. 
// You are also given two integers cost1 and cost2 indicating the price of a pen and pencil respectively. Y
// ou can spend part or all of your money to buy multiple quantities (or none) of each kind of writing utensil.

// Return the number of distinct ways you can buy some number of pens and pencils.

// Example 1:
// Input: total = 20, cost1 = 10, cost2 = 5
// Output: 9
// Explanation: The price of a pen is 10 and the price of a pencil is 5.
// - If you buy 0 pens, you can buy 0, 1, 2, 3, or 4 pencils.
// - If you buy 1 pen, you can buy 0, 1, or 2 pencils.
// - If you buy 2 pens, you cannot buy any pencils.
// The total number of ways to buy pens and pencils is 5 + 3 + 1 = 9.

// Example 2:
// Input: total = 5, cost1 = 10, cost2 = 10
// Output: 1
// Explanation: The price of both pens and pencils are 10, which cost more than total, so you cannot buy any writing utensils. Therefore, there is only 1 way: buy 0 pens and 0 pencils.

// Constraints:
//     1 <= total, cost1, cost2 <= 10^6

import "fmt"

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
    res := 0
    for i := 0; i <= (total / cost1); i++  {
        // rem := total - (i * cost1)
        // res += (rem / cost2 + 1)
        res += ((total - cost1 * i) / cost2 + 1)
    }
    return int64(res)
}

func waysToBuyPensPencils1(total int, cost1 int, cost2 int) int64 {
    if total < cost1 && total < cost2 { return 1 }
    res := 0
    for i := 0; cost1 * i <= total; i++ {
        left := total - cost1 * i // How many pens you buy
        res += (left / cost2 + 1) // The configuration of pencils you can afford: 0, 1, 2, ...
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: total = 20, cost1 = 10, cost2 = 5
    // Output: 9
    // Explanation: The price of a pen is 10 and the price of a pencil is 5.
    // - If you buy 0 pens, you can buy 0, 1, 2, 3, or 4 pencils.
    // - If you buy 1 pen, you can buy 0, 1, or 2 pencils.
    // - If you buy 2 pens, you cannot buy any pencils.
    // The total number of ways to buy pens and pencils is 5 + 3 + 1 = 9.
    fmt.Println(waysToBuyPensPencils(20, 10, 5)) // 9
    // Example 2:
    // Input: total = 5, cost1 = 10, cost2 = 10
    // Output: 1
    // Explanation: The price of both pens and pencils are 10, which cost more than total, so you cannot buy any writing utensils. Therefore, there is only 1 way: buy 0 pens and 0 pencils.
    fmt.Println(waysToBuyPensPencils(5, 10, 10)) // 1

    fmt.Println(waysToBuyPensPencils(1, 1, 1)) // 3
    fmt.Println(waysToBuyPensPencils(10_000_000, 10_000_000, 10_000_000)) // 3
    fmt.Println(waysToBuyPensPencils(1, 10_000_000, 10_000_000)) // 1
    fmt.Println(waysToBuyPensPencils(10_000_000, 1, 10_000_000)) // 10000002
    fmt.Println(waysToBuyPensPencils(10_000_000, 10_000_000, 1)) // 10000002
    fmt.Println(waysToBuyPensPencils(1, 1, 10_000_000)) // 2
    fmt.Println(waysToBuyPensPencils(10_000_000, 1, 1)) // 50000015000001
    fmt.Println(waysToBuyPensPencils(1, 10_000_000, 1)) // 2

    fmt.Println(waysToBuyPensPencils1(20, 10, 5)) // 9
    fmt.Println(waysToBuyPensPencils1(5, 10, 10)) // 1
    fmt.Println(waysToBuyPensPencils1(1, 1, 1)) // 3
    fmt.Println(waysToBuyPensPencils1(10_000_000, 10_000_000, 10_000_000)) // 3
    fmt.Println(waysToBuyPensPencils1(1, 10_000_000, 10_000_000)) // 1
    fmt.Println(waysToBuyPensPencils1(10_000_000, 1, 10_000_000)) // 10000002
    fmt.Println(waysToBuyPensPencils1(10_000_000, 10_000_000, 1)) // 10000002
    fmt.Println(waysToBuyPensPencils1(1, 1, 10_000_000)) // 2
    fmt.Println(waysToBuyPensPencils1(10_000_000, 1, 1)) // 50000015000001
    fmt.Println(waysToBuyPensPencils1(1, 10_000_000, 1)) // 2
}