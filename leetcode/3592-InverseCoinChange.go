package main

// 3592. Inverse Coin Change
// You are given a 1-indexed integer array numWays, where numWays[i] represents the number of ways to select a total amount i using an infinite supply of some fixed coin denominations. 
// Each denomination is a positive integer with value at most numWays.length.

// However, the exact coin denominations have been lost. 
// Your task is to recover the set of denominations that could have resulted in the given numWays array.

// Return a sorted array containing unique integers which represents this set of denominations.

// If no such set exists, return an empty array.

// Example 1:
// Input: numWays = [0,1,0,2,0,3,0,4,0,5]
// Output: [2,4,6]
// Explanation:
// Amount	Number of ways	Explanation
// 1	0	There is no way to select coins with total value 1.
// 2	1	The only way is [2].
// 3	0	There is no way to select coins with total value 3.
// 4	2	The ways are [2, 2] and [4].
// 5	0	There is no way to select coins with total value 5.
// 6	3	The ways are [2, 2, 2], [2, 4], and [6].
// 7	0	There is no way to select coins with total value 7.
// 8	4	The ways are [2, 2, 2, 2], [2, 2, 4], [2, 6], and [4, 4].
// 9	0	There is no way to select coins with total value 9.
// 10	5	The ways are [2, 2, 2, 2, 2], [2, 2, 2, 4], [2, 4, 4], [2, 2, 6], and [4, 6].

// Example 2:
// Input: numWays = [1,2,2,3,4]
// Output: [1,2,5]
// Explanation:
// Amount	Number of ways	Explanation
// 1	1	The only way is [1].
// 2	2	The ways are [1, 1] and [2].
// 3	2	The ways are [1, 1, 1] and [1, 2].
// 4	3	The ways are [1, 1, 1, 1], [1, 1, 2], and [2, 2].
// 5	4	The ways are [1, 1, 1, 1, 1], [1, 1, 1, 2], [1, 2, 2], and [5].

// Example 3:
// Input: numWays = [1,2,3,4,15]
// Output: []
// Explanation:
// No set of denomination satisfies this array.

// Constraints:
//     1 <= numWays.length <= 100
//     0 <= numWays[i] <= 2 * 10^8

import "fmt"

func findCoins(numWays []int) []int {
    n := len(numWays)
    numWays = append([]int{1}, numWays...)
    res, myWays := make([]int, 0), make([]int, n + 1)
    myWays[0] = 1
    for i := 1; i <= n; i++ {
        // If `myWays[x] == numWays[x]`, move on.
        if myWays[i] == numWays[i] { continue }
        // If `myWays[x] + 1 == numWays[x]` → add that value as a coin in our basket and update `myWays`, so `myWays[x...n]` accounts for ways with the new coin.
        if numWays[i]-myWays[i] == 1 {
            res = append(res, i)
            for j := i; j <= n; j++ {
                myWays[j] += myWays[j-i]
            }
        } else { // If `myWays[x] + 1 < numWays[x]` → no solution. (*see below for why)
            return []int{}
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: numWays = [0,1,0,2,0,3,0,4,0,5]
    // Output: [2,4,6]
    // Explanation:
    // Amount	Number of ways	Explanation
    // 1	0	There is no way to select coins with total value 1.
    // 2	1	The only way is [2].
    // 3	0	There is no way to select coins with total value 3.
    // 4	2	The ways are [2, 2] and [4].
    // 5	0	There is no way to select coins with total value 5.
    // 6	3	The ways are [2, 2, 2], [2, 4], and [6].
    // 7	0	There is no way to select coins with total value 7.
    // 8	4	The ways are [2, 2, 2, 2], [2, 2, 4], [2, 6], and [4, 4].
    // 9	0	There is no way to select coins with total value 9.
    // 10	5	The ways are [2, 2, 2, 2, 2], [2, 2, 2, 4], [2, 4, 4], [2, 2, 6], and [4, 6].
    fmt.Println(findCoins([]int{0,1,0,2,0,3,0,4,0,5})) // [2,4,6]
    // Example 2:
    // Input: numWays = [1,2,2,3,4]
    // Output: [1,2,5]
    // Explanation:
    // Amount	Number of ways	Explanation
    // 1	1	The only way is [1].
    // 2	2	The ways are [1, 1] and [2].
    // 3	2	The ways are [1, 1, 1] and [1, 2].
    // 4	3	The ways are [1, 1, 1, 1], [1, 1, 2], and [2, 2].
    // 5	4	The ways are [1, 1, 1, 1, 1], [1, 1, 1, 2], [1, 2, 2], and [5].
    fmt.Println(findCoins([]int{1,2,2,3,4})) // [1,2,5]
    // Example 3:
    // Input: numWays = [1,2,3,4,15]
    // Output: []
    // Explanation:
    // No set of denomination satisfies this array.
    fmt.Println(findCoins([]int{1,2,3,4,15})) // []

    fmt.Println(findCoins([]int{1,2,3,4,5,6,7,8,9})) // []
    fmt.Println(findCoins([]int{9,8,7,6,5,4,3,2,11})) // []
}