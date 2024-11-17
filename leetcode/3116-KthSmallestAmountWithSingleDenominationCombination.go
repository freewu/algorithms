package main

// 3116. Kth Smallest Amount With Single Denomination Combination
// You are given an integer array coins representing coins of different denominations and an integer k.

// You have an infinite number of coins of each denomination. 
// However, you are not allowed to combine coins of different denominations.

// Return the kth smallest amount that can be made using these coins.

// Example 1:
// Input: coins = [3,6,9], k = 3
// Output:  9
// Explanation: The given coins can make the following amounts:
// Coin 3 produces multiples of 3: 3, 6, 9, 12, 15, etc.
// Coin 6 produces multiples of 6: 6, 12, 18, 24, etc.
// Coin 9 produces multiples of 9: 9, 18, 27, 36, etc.
// All of the coins combined produce: 3, 6, 9, 12, 15, etc.

// Example 2:
// Input: coins = [5,2], k = 7
// Output: 12 
// Explanation: The given coins can make the following amounts:
// Coin 5 produces multiples of 5: 5, 10, 15, 20, etc.
// Coin 2 produces multiples of 2: 2, 4, 6, 8, 10, 12, etc.
// All of the coins combined produce: 2, 4, 5, 6, 8, 10, 12, 14, 15, etc.

// Constraints:
//     1 <= coins.length <= 15
//     1 <= coins[i] <= 25
//     1 <= k <= 2 * 10^9
//     coins contains pairwise distinct integers.

import "fmt"
import "sort"
import "math/bits"

func findKthSmallest(coins []int, k int) int64 {
    sort.Ints(coins)
    arr := coins[:0]
    for _, x := range coins {
        for _, y := range arr {
            if x % y == 0 { break }
        }
        arr = append(arr, x)
    }
    gcd := func(x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    lcm := func(x, y int) int { return x / gcd(x, y) * y }
    subsetLcm := make([]int, 1 << len(arr))
    subsetLcm[0] = 1
    for i, x := range arr {
        bit := 1 << i
        for mask, l := range subsetLcm[:bit] {
            subsetLcm[bit|mask] = lcm(l, x)
        }
    }
    for i := range subsetLcm {
        if bits.OnesCount(uint(i)) % 2 == 0 {
            subsetLcm[i] *= -1
        }
    }
    res := sort.Search(arr[0] * k, func(m int) bool {
        count := 0
        for _, v := range subsetLcm[1:] {
            count += m / v
        }
        return count >= k
    })
    return int64(res)
}

func main() {
    // Example 1:
    // Input: coins = [3,6,9], k = 3
    // Output:  9
    // Explanation: The given coins can make the following amounts:
    // Coin 3 produces multiples of 3: 3, 6, 9, 12, 15, etc.
    // Coin 6 produces multiples of 6: 6, 12, 18, 24, etc.
    // Coin 9 produces multiples of 9: 9, 18, 27, 36, etc.
    // All of the coins combined produce: 3, 6, 9, 12, 15, etc.
    fmt.Println(findKthSmallest([]int{3,6,9}, 3)) // 9
    // Example 2:
    // Input: coins = [5,2], k = 7
    // Output: 12 
    // Explanation: The given coins can make the following amounts:
    // Coin 5 produces multiples of 5: 5, 10, 15, 20, etc.
    // Coin 2 produces multiples of 2: 2, 4, 6, 8, 10, 12, etc.
    // All of the coins combined produce: 2, 4, 5, 6, 8, 10, 12, 14, 15, etc.
    fmt.Println(findKthSmallest([]int{5,2}, 7)) // 12
}