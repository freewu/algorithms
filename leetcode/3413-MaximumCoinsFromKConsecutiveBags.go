package main

// 3413. Maximum Coins From K Consecutive Bags
// There are an infinite amount of bags on a number line, one bag for each coordinate. 
// Some of these bags contain coins.

// You are given a 2D array coins, where coins[i] = [li, ri, ci] denotes that every bag from li to ri contains ci coins.

// The segments that coins contain are non-overlapping.

// You are also given an integer k.

// Return the maximum amount of coins you can obtain by collecting k consecutive bags.

// Example 1:
// Input: coins = [[8,10,1],[1,3,2],[5,6,4]], k = 4
// Output: 10
// Explanation:
// Selecting bags at positions [3, 4, 5, 6] gives the maximum number of coins: 2 + 0 + 4 + 4 = 10.

// Example 2:
// Input: coins = [[1,10,3]], k = 2
// Output: 6
// Explanation:
// Selecting bags at positions [1, 2] gives the maximum number of coins: 3 + 3 = 6.

// Constraints:
//     1 <= coins.length <= 10^5
//     1 <= k <= 10^9
//     coins[i] == [li, ri, ci]
//     1 <= li <= ri <= 10^9
//     1 <= ci <= 1000
//     The given segments are non-overlapping.

import "fmt"
import "sort"
import "slices"

func maximumCoins(coins [][]int, k int) int64 {
    sort.Slice(coins, func(i, j int) bool {
        return coins[i][0] < coins[j][0]
    })
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, cur, n := 0, 0, len(coins)
    // Start at coins[i][0]
    for i, j := 0, 0; i < n; i++ {
        for j < n && coins[j][1] <= (coins[i][0] + k - 1) {
            cur += (coins[j][1] - coins[j][0] + 1) * coins[j][2]
            j++
        }
        if j < n {
            part := max(0, coins[i][0] + k - 1 - coins[j][0] + 1) * coins[j][2]
            res = max(res, cur + part)
        }
        cur -= (coins[i][1] - coins[i][0] + 1) * coins[i][2]
    }
    // End at coins[i][1]
    cur = 0
    for i, j := 0, 0; i < n; i++ {
        cur += (coins[i][1] - coins[i][0] + 1) * coins[i][2]
        for coins[j][1] < coins[i][1] - k + 1 {
            cur -= (coins[j][1] - coins[j][0] + 1) * coins[j][2]
            j++
        }
        part := max(0, coins[i][1] - k - coins[j][0] + 1) * coins[j][2]
        res = max(res, cur - part)
    }
    return int64(res)
}

func maximumCoins1(coins [][]int, k int) int64 {
    slices.SortFunc(coins, func(a, b []int) int { 
        return a[0] - b[0] 
    })
    max := func (x, y int) int { if x > y { return x; }; return y; }
    maximumWhiteTiles := func(tiles [][]int, carpetLen int) int {
        res, cover, left := 0, 0, 0
        for _, tile := range tiles {
            tl, tr, c := tile[0], tile[1], tile[2]
            cover += (tr - tl + 1) * c
            for tiles[left][1] + carpetLen - 1 < tr {
                cover -= (tiles[left][1] - tiles[left][0] + 1) * tiles[left][2]
                left++
            }
            uncover := max((tr - carpetLen + 1 - tiles[left][0]) * tiles[left][2], 0)
            res = max(res, cover - uncover)
        }
        return res
    }
    res := maximumWhiteTiles(coins, k)
    slices.Reverse(coins)
    for _, t := range coins {
        t[0], t[1] = -t[1], -t[0]
    }
    return int64(max(res, maximumWhiteTiles(coins, k)))
}

func main() {
    // Example 1:
    // Input: coins = [[8,10,1],[1,3,2],[5,6,4]], k = 4
    // Output: 10
    // Explanation:
    // Selecting bags at positions [3, 4, 5, 6] gives the maximum number of coins: 2 + 0 + 4 + 4 = 10.
    fmt.Println(maximumCoins([][]int{{8,10,1},{1,3,2},{5,6,4}}, 4)) // 10
    // Example 2:
    // Input: coins = [[1,10,3]], k = 2
    // Output: 6
    // Explanation:
    // Selecting bags at positions [1, 2] gives the maximum number of coins: 3 + 3 = 6.
    fmt.Println(maximumCoins([][]int{{1,10,3}}, 2)) // 6

    fmt.Println(maximumCoins1([][]int{{8,10,1},{1,3,2},{5,6,4}}, 4)) // 10
    fmt.Println(maximumCoins1([][]int{{1,10,3}}, 2)) // 6
}