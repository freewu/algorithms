package main

// 2600. K Items With the Maximum Sum
// There is a bag that consists of items, each item has a number 1, 0, or -1 written on it.

// You are given four non-negative integers numOnes, numZeros, numNegOnes, and k.

// The bag initially contains:
//     1. numOnes items with 1s written on them.
//     2. numZeroes items with 0s written on them.
//     3. numNegOnes items with -1s written on them.

// We want to pick exactly k items among the available items. 
// Return the maximum possible sum of numbers written on the items.

// Example 1:
// Input: numOnes = 3, numZeros = 2, numNegOnes = 0, k = 2
// Output: 2
// Explanation: We have a bag of items with numbers written on them {1, 1, 1, 0, 0}. We take 2 items with 1 written on them and get a sum in a total of 2.
// It can be proven that 2 is the maximum possible sum.

// Example 2:
// Input: numOnes = 3, numZeros = 2, numNegOnes = 0, k = 4
// Output: 3
// Explanation: We have a bag of items with numbers written on them {1, 1, 1, 0, 0}. We take 3 items with 1 written on them, and 1 item with 0 written on it, and get a sum in a total of 3.
// It can be proven that 3 is the maximum possible sum.

// Constraints:
//     0 <= numOnes, numZeros, numNegOnes <= 50
//     0 <= k <= numOnes + numZeros + numNegOnes

import "fmt"

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    maxSum := 0
    // ones & zeros:
    adding := min(numOnes, k) 
    maxSum += adding
    k -= adding + numZeros // we can do this, as zeros do not affect maxSum
    if k <= 0 {
        return maxSum
    }
    // negatives:
    return maxSum - k // we can do this with given constraint: 0 <= k <= numOnes + numZeros + numNegOnes
}

func kItemsWithMaximumSum1(numOnes int, numZeros int, numNegOnes int, k int) int {
    res := 0
    for i := 0; i < k; i++ {
        if numOnes > 0 {
            res++
            numOnes--
            continue
        }
        if numZeros > 0 {
            numZeros--
            continue
        }
        if numNegOnes > 0 {
            res--
            numNegOnes--
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: numOnes = 3, numZeros = 2, numNegOnes = 0, k = 2
    // Output: 2
    // Explanation: We have a bag of items with numbers written on them {1, 1, 1, 0, 0}. We take 2 items with 1 written on them and get a sum in a total of 2.
    // It can be proven that 2 is the maximum possible sum.
    fmt.Println(kItemsWithMaximumSum(3, 3, 0, 2)) // 2
    // Example 2:
    // Input: numOnes = 3, numZeros = 2, numNegOnes = 0, k = 4
    // Output: 3
    // Explanation: We have a bag of items with numbers written on them {1, 1, 1, 0, 0}. We take 3 items with 1 written on them, and 1 item with 0 written on it, and get a sum in a total of 3.
    // It can be proven that 3 is the maximum possible sum.
    fmt.Println(kItemsWithMaximumSum(3, 2, 0, 4)) // 3

    fmt.Println(kItemsWithMaximumSum(0, 0, 0, 0)) // 0
    fmt.Println(kItemsWithMaximumSum(50, 50, 50, 0)) // 0
    fmt.Println(kItemsWithMaximumSum(50, 50, 50, 150)) // 0
    fmt.Println(kItemsWithMaximumSum(0, 50, 50, 100)) // -50
    fmt.Println(kItemsWithMaximumSum(50, 0, 50, 100)) // 0
    fmt.Println(kItemsWithMaximumSum(50, 50, 0, 100)) // 50

    fmt.Println(kItemsWithMaximumSum1(3, 3, 0, 2)) // 2
    fmt.Println(kItemsWithMaximumSum1(3, 2, 0, 4)) // 3
    fmt.Println(kItemsWithMaximumSum1(0, 0, 0, 0)) // 0
    fmt.Println(kItemsWithMaximumSum1(50, 50, 50, 0)) // 0
    fmt.Println(kItemsWithMaximumSum1(50, 50, 50, 150)) // 0
    fmt.Println(kItemsWithMaximumSum1(0, 50, 50, 100)) // -50
    fmt.Println(kItemsWithMaximumSum1(50, 0, 50, 100)) // 0
    fmt.Println(kItemsWithMaximumSum1(50, 50, 0, 100)) // 50
}