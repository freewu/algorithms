package main

// 3180. Maximum Total Reward Using Operations I
// You are given an integer array rewardValues of length n, representing the values of rewards.

// Initially, your total reward x is 0, and all indices are unmarked. 
// You are allowed to perform the following operation any number of times:
//     1. Choose an unmarked index i from the range [0, n - 1].
//     2. If rewardValues[i] is greater than your current total reward x, 
//        then add rewardValues[i] to x (i.e., x = x + rewardValues[i]), and mark the index i.

// Return an integer denoting the maximum total reward you can collect by performing the operations optimally. 

// Example 1:
// Input: rewardValues = [1,1,3,3]
// Output: 4
// Explanation:
// During the operations, we can choose to mark the indices 0 and 2 in order, and the total reward will be 4, which is the maximum.

// Example 2:
// Input: rewardValues = [1,6,4,3,2]
// Output: 11
// Explanation:
// Mark the indices 0, 2, and 1 in order. The total reward will then be 11, which is the maximum.

// Constraints:
//     1 <= rewardValues.length <= 2000
//     1 <= rewardValues[i] <= 2000

import "fmt"
import "math/big"
import "sort"

func maxTotalReward1(rewardValues []int) int {
    sort.Ints(rewardValues)
    f0, f1 := big.NewInt(1), big.NewInt(0)
    for _, x := range rewardValues {
        mask, one := big.NewInt(0), big.NewInt(1)
        mask.Sub(mask.Lsh(one, uint(x)), one)
        f1.Lsh(f1.And(f0, mask), uint(x))
        f0.Or(f0, f1)
    }
    return f0.BitLen() - 1
}

func maxTotalReward(rewardValues []int) int {
    count, freq := 0, make([]int, 2001)
    for _, v := range rewardValues {
        if freq[v] == 0 { count++ }
        freq[v]++
    } 
    index, arr := 0, make([]int, count)
    for i := 1; i < len(freq); i++ {
        if freq[i] > 0 {
            arr[index] = i
            index++
        }
    }
    n := count // sorting completed
    mx := arr[n - 1] * 2 // for this testcase the max value of reward could be max value x2
    dp := make([]bool, mx + 1)
    dp[0] = true
    for i := 0; i < len(arr); i++ {
        for j := 0; j < arr[i]; j++ {
            dp[j + arr[i]] = dp[j + arr[i]] || dp[j]
        }
    }
    res := 0
    for i := 0; i < len(dp); i++ {
        if dp[i]  {
            res = i
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: rewardValues = [1,1,3,3]
    // Output: 4
    // Explanation:
    // During the operations, we can choose to mark the indices 0 and 2 in order, and the total reward will be 4, which is the maximum.
    fmt.Println(maxTotalReward([]int{1,1,3,3})) // 4
    // Example 2:
    // Input: rewardValues = [1,6,4,3,2]
    // Output: 11
    // Explanation:
    // Mark the indices 0, 2, and 1 in order. The total reward will then be 11, which is the maximum.
    fmt.Println(maxTotalReward([]int{1,6,4,3,2})) // 11

    fmt.Println(maxTotalReward1([]int{1,1,3,3})) // 4
    fmt.Println(maxTotalReward1([]int{1,6,4,3,2})) // 11
}