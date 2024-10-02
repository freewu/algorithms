package main

// 1359. Count All Valid Pickup and Delivery Options
// Given n orders, each order consists of a pickup and a delivery service.

// Count all valid pickup/delivery possible sequences such that delivery(i) is always after of pickup(i). 

// Since the answer may be too large, return it modulo 10^9 + 7.

// Example 1:
// Input: n = 1
// Output: 1
// Explanation: Unique order (P1, D1), Delivery 1 always is after of Pickup 1.

// Example 2:
// Input: n = 2
// Output: 6
// Explanation: All possible orders: 
// (P1,P2,D1,D2), (P1,P2,D2,D1), (P1,D1,P2,D2), (P2,P1,D1,D2), (P2,P1,D2,D1) and (P2,D2,P1,D1).
// This is an invalid order (P1,D2,P2,D1) because Pickup 2 is after of Delivery 2.

// Example 3:
// Input: n = 3
// Output: 90

// Constraints:
//     1 <= n <= 500

import "fmt"

func countOrders(n int) int {
    // 计算所有有效的 取货/交付 可能的顺序, delivery(i)总是在pickup(i)之后
    // 由于答案很大，返回答案对10^9+7取余的结果
    // 1           1
    // 2           6 (3+2+1) = 6
    // 3           6*(5+4+3+2+1) = 15*6=90
    res := 1
    for i := 2; i <= n; i++ {
        res = (i * (2 * i - 1) * res) % 1_000_000_007
    }
    return res
}

func countOrders1(n int) int {
    res := 1
    for i := 1; i <= 2 * n; i++ {
        res = res * i // number of ways to arrange pickups
        // We only need to divide the result by 2 n-times.
        // To prevent decimal results we divide after multiplying an even number.
        if i % 2 == 0 {
            res = res / 2
        }
        res = res % 1_000_000_007
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 1
    // Output: 1
    // Explanation: Unique order (P1, D1), Delivery 1 always is after of Pickup 1.
    fmt.Println(countOrders(1)) // 1
    // Example 2:
    // Input: n = 2
    // Output: 6
    // Explanation: All possible orders: 
    // (P1,P2,D1,D2), (P1,P2,D2,D1), (P1,D1,P2,D2), (P2,P1,D1,D2), (P2,P1,D2,D1) and (P2,D2,P1,D1).
    // This is an invalid order (P1,D2,P2,D1) because Pickup 2 is after of Delivery 2.
    fmt.Println(countOrders(2)) // 6
    // Example 3:
    // Input: n = 3
    // Output: 90
    fmt.Println(countOrders(3)) // 90

    fmt.Println(countOrders(10)) // 850728840
    fmt.Println(countOrders(128)) // 761756041
    fmt.Println(countOrders(499)) // 496638171
    fmt.Println(countOrders(500)) // 764678010

    fmt.Println(countOrders1(1)) // 1
    fmt.Println(countOrders1(2)) // 6
    fmt.Println(countOrders1(3)) // 90
    fmt.Println(countOrders1(10)) // 850728840
    fmt.Println(countOrders1(128)) // 761756041
    fmt.Println(countOrders1(499)) // 496638171
    fmt.Println(countOrders1(500)) // 764678010
}