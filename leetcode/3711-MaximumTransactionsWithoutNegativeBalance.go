package main

// 3711. Maximum Transactions Without Negative Balance
// You are given an integer array transactions, where transactions[i] represents the amount of the ith transaction:
//     1. A positive value means money is received.
//     2. A negative value means money is sent.

// The account starts with a balance of 0, and the balance must never become negative. 
// Transactions must be considered in the given order, but you are allowed to skip some transactions.

// Return an integer denoting the maximum number of transactions that can be performed without the balance ever going negative.

// Example 1:
// Input: transactions = [2,-5,3,-1,-2]
// Output: 4
// Explanation:
// One optimal sequence is [2, 3, -1, -2], balance: 0 → 2 → 5 → 4 → 2.

// Example 2:
// Input: transactions = [-1,-2,-3]
// Output: 0
// Explanation:
// All transactions are negative. Including any would make the balance negative.

// Example 3:
// Input: transactions = [3,-2,3,-2,1,-1]
// Output: 6
// Explanation:
// All transactions can be taken in order, balance: 0 → 3 → 1 → 4 → 2 → 3 → 2.

// Constraints:
//     1 <= transactions.length <= 10^5
//     -10^9 <= transactions[i] <= 10^9

import "fmt"
import "container/heap"

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] } // Max-heap
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func maxTransactions(transactions []int) int {
    res, balance := 0, 0
    hp := &MaxHeap{}
    heap.Init(hp)
    for _, t := range transactions {
        balance += t
        res++ 
        // If the transaction is negative, add it to the max-heap
        if t < 0 {
            heap.Push(hp, t)
        }
        // If balance becomes negative, try to remove the largest negative transaction
        // to maximize the number of transactions we can keep
        for balance < 0 && hp.Len() > 0 {
            // Remove the largest negative transaction (most harmful)
            largestNegative := heap.Pop(hp).(int)
            balance -= largestNegative
            res--
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: transactions = [2,-5,3,-1,-2]
    // Output: 4
    // Explanation:
    // One optimal sequence is [2, 3, -1, -2], balance: 0 → 2 → 5 → 4 → 2.
    fmt.Println(maxTransactions([]int{2,-5,3,-1,-2})) // 4
    // Example 2:
    // Input: transactions = [-1,-2,-3]
    // Output: 0
    // Explanation:
    // All transactions are negative. Including any would make the balance negative.
    fmt.Println(maxTransactions([]int{-1,-2,-3})) // 0
    // Example 3:
    // Input: transactions = [3,-2,3,-2,1,-1]
    // Output: 6
    // Explanation:
    // All transactions can be taken in order, balance: 0 → 3 → 1 → 4 → 2 → 3 → 2.
    fmt.Println(maxTransactions([]int{3,-2,3,-2,1,-1})) // 6

    fmt.Println(maxTransactions([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(maxTransactions([]int{9,8,7,6,5,4,3,2,1})) // 9
}