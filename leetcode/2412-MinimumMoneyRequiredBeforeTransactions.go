package main

// 2412. Minimum Money Required Before Transactions
// You are given a 0-indexed 2D integer array transactions, where transactions[i] = [costi, cashbacki].

// The array describes transactions, where each transaction must be completed exactly once in some order. 
// At any given moment, you have a certain amount of money. 
// In order to complete transaction i, money >= costi must hold true. 
// After performing a transaction, money becomes money - costi + cashbacki.

// Return the minimum amount of money required before any transaction so that all of the transactions can be completed regardless of the order of the transactions.

// Example 1:
// Input: transactions = [[2,1],[5,0],[4,2]]
// Output: 10
// Explanation:
// Starting with money = 10, the transactions can be performed in any order.
// It can be shown that starting with money < 10 will fail to complete all transactions in some order.

// Example 2:
// Input: transactions = [[3,0],[0,3]]
// Output: 3
// Explanation:
// - If transactions are in the order [[3,0],[0,3]], the minimum money required to complete the transactions is 3.
// - If transactions are in the order [[0,3],[3,0]], the minimum money required to complete the transactions is 0.
// Thus, starting with money = 3, the transactions can be performed in any order.

// Constraints:
//     1 <= transactions.length <= 10^5
//     transactions[i].length == 2
//     0 <= costi, cashbacki <= 10^9

import "fmt"

func minimumMoney(transactions [][]int) int64 {
    base, extra := 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, t := range transactions {
        if t[0] > t[1] {
            base += t[0] - t[1]
            extra = max(extra, t[1])
        } else {
            extra = max(extra, t[0])
        }
    }
    return int64(base + extra)
}

func minimumMoney1(transactions [][]int) int64 {
    sum, mx := 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, t := range transactions {
        sum += max(t[0] - t[1], 0)
        mx = max(mx, min(t[0], t[1]))
    }
    return int64(sum + mx)
}

func main() {
    // Example 1:
    // Input: transactions = [[2,1],[5,0],[4,2]]
    // Output: 10
    // Explanation:
    // Starting with money = 10, the transactions can be performed in any order.
    // It can be shown that starting with money < 10 will fail to complete all transactions in some order.
    fmt.Println(minimumMoney([][]int{{2,1},{5,0},{4,2}})) // 10
    // Example 2:
    // Input: transactions = [[3,0],[0,3]]
    // Output: 3
    // Explanation:
    // - If transactions are in the order [[3,0],[0,3]], the minimum money required to complete the transactions is 3.
    // - If transactions are in the order [[0,3],[3,0]], the minimum money required to complete the transactions is 0.
    // Thus, starting with money = 3, the transactions can be performed in any order.
    fmt.Println(minimumMoney([][]int{{3,0},{0,3}})) // 3

    fmt.Println(minimumMoney1([][]int{{2,1},{5,0},{4,2}})) // 10
    fmt.Println(minimumMoney1([][]int{{3,0},{0,3}})) // 3
}