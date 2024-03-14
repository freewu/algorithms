package main

// 2312. Selling Pieces of Wood
// You are given two integers m and n that represent the height and width of a rectangular piece of wood. 
// You are also given a 2D integer array prices, where prices[i] = [hi, wi, pricei] indicates you can sell a rectangular piece of wood of height hi and width wi for pricei dollars.
// To cut a piece of wood, you must make a vertical or horizontal cut across the entire height or width of the piece to split it into two smaller pieces. 
// After cutting a piece of wood into some number of smaller pieces, you can sell pieces according to prices. 
// You may sell multiple pieces of the same shape, and you do not have to sell all the shapes. 
// The grain of the wood makes a difference, so you cannot rotate a piece to swap its height and width.
// Return the maximum money you can earn after cutting an m x n piece of wood.
// Note that you can cut the piece of wood as many times as you want.
 
// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/04/27/ex1.png"/>
// Input: m = 3, n = 5, prices = [[1,4,2],[2,2,7],[2,1,3]]
// Output: 19
// Explanation: The diagram above shows a possible scenario. It consists of:
// - 2 pieces of wood shaped 2 x 2, selling for a price of 2 * 7 = 14.
// - 1 piece of wood shaped 2 x 1, selling for a price of 1 * 3 = 3.
// - 1 piece of wood shaped 1 x 4, selling for a price of 1 * 2 = 2.
// This obtains a total of 14 + 3 + 2 = 19 money earned.
// It can be shown that 19 is the maximum amount of money that can be earned.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/04/27/ex2new.png"/>
// Input: m = 4, n = 6, prices = [[3,2,10],[1,4,2],[4,1,3]]
// Output: 32
// Explanation: The diagram above shows a possible scenario. It consists of:
// - 3 pieces of wood shaped 3 x 2, selling for a price of 3 * 10 = 30.
// - 1 piece of wood shaped 1 x 4, selling for a price of 1 * 2 = 2.
// This obtains a total of 30 + 2 = 32 money earned.
// It can be shown that 32 is the maximum amount of money that can be earned.
// Notice that we cannot rotate the 1 x 4 piece of wood to obtain a 4 x 1 piece of wood.
 
// Constraints:
//     1 <= m, n <= 200
//     1 <= prices.length <= 2 * 10^4
//     prices[i].length == 3
//     1 <= hi <= m
//     1 <= wi <= n
//     1 <= pricei <= 10^6
//     All the shapes of wood (hi, wi) are pairwise distinct.

import "fmt"

func dp(h, w int, memo, p [][]int64) int64 {
    if h == 0 || w == 0 {
        return 0
    }
    if memo[h][w] >= 0 {
        return memo[h][w]
    }
    memo[h][w] = 0
    r := int64(0)
    if p[h][w] != -1 {
        r = p[h][w]
    }
    for i := 1; i < h; i++ {
        x := dp(i, w, memo, p) + dp(h - i, w, memo, p)
        if x > r {
            r = x
        }        
    }
    for j := 1; j < w; j++ {
        x := dp(h, j, memo, p) + dp(h, w - j, memo, p)
        if x > r {
            r = x
        }
    }
    memo[h][w] = r
    return r    
}

func sellingWood(m int, n int, prices [][]int) int64 {
    memo := make([][]int64, m + 1)
    p := make([][]int64, m + 1)
    for i := 0; i <= m; i++ {
        memo[i] = make([]int64, n + 1)
        p[i] = make([]int64, n + 1)
        for j := 0; j <= n; j++ {
            memo[i][j] = -1
            p[i][j] = -1
        }
    }
    for _, x := range prices {
        p[x[0]][x[1]] = int64(x[2])
    }    
    return dp(m, n, memo, p)
}

func sellingWood1(m, n int, prices [][]int) int64 {
    max := func (a, b int) int { if b > a { return b }; return a }
	pr := make([][]int, m+1)
	for i := range pr {
		pr[i] = make([]int, n+1)
	}
	for _, price := range prices {
		pr[price[0]][price[1]] = price[2]
	}
	f := make([][]int, m+1)
	for i := 1; i <= m; i++ {
		f[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			f[i][j] = pr[i][j]
			for k := 1; k < j; k++ { // 垂直切割
				f[i][j] = max(f[i][j], f[i][k]+f[i][j-k])
			}
			for k := 1; k < i; k++ { // 水平切割
				f[i][j] = max(f[i][j], f[k][j]+f[i-k][j])
			}
		}
	}
	return int64(f[m][n])
}

func main() {
    // Input: m = 3, n = 5, prices = [[1,4,2],[2,2,7],[2,1,3]]
    // Output: 19
    // Explanation: The diagram above shows a possible scenario. It consists of:
    // - 2 pieces of wood shaped 2 x 2, selling for a price of 2 * 7 = 14.
    // - 1 piece of wood shaped 2 x 1, selling for a price of 1 * 3 = 3.
    // - 1 piece of wood shaped 1 x 4, selling for a price of 1 * 2 = 2.
    // This obtains a total of 14 + 3 + 2 = 19 money earned.
    fmt.Println(sellingWood(3,5,[][]int{[]int{1,4,2},[]int{2,2,7},[]int{2,1,3}})) // 19

    // Input: m = 4, n = 6, prices = [[3,2,10],[1,4,2],[4,1,3]]
    // Output: 32
    // Explanation: The diagram above shows a possible scenario. It consists of:
    // - 3 pieces of wood shaped 3 x 2, selling for a price of 3 * 10 = 30.
    // - 1 piece of wood shaped 1 x 4, selling for a price of 1 * 2 = 2.
    // This obtains a total of 30 + 2 = 32 money earned.
    fmt.Println(sellingWood(4,6,[][]int{[]int{3,2,10},[]int{1,4,2},[]int{4,1,3}})) // 32

    fmt.Println(sellingWood1(3,5,[][]int{[]int{1,4,2},[]int{2,2,7},[]int{2,1,3}})) // 19
    fmt.Println(sellingWood1(4,6,[][]int{[]int{3,2,10},[]int{1,4,2},[]int{4,1,3}})) // 32
}