package main

import "fmt"

// 1690. Stone Game VII
// Alice and Bob take turns playing a game, with Alice starting first.
// There are n stones arranged in a row. On each player's turn, 
// they can remove either the leftmost stone or the rightmost stone from the row and receive points equal to the sum of the remaining stones' values in the row.
// The winner is the one with the higher score when there are no stones left to remove.

// Bob found that he will always lose this game (poor Bob, he always loses), 
// so he decided to minimize the score's difference. Alice's goal is to maximize the difference in the score.

// Given an array of integers stones where stones[i] represents the value of the ith stone from the left, 
// return the difference in Alice and Bob's score if they both play optimally.

// Example 1:
// Input: stones = [5,3,1,4,2]
// Output: 6
// Explanation: 
// - Alice removes 2 and gets 5 + 3 + 1 + 4 = 13 points. Alice = 13, Bob = 0, stones = [5,3,1,4].
// - Bob removes 5 and gets 3 + 1 + 4 = 8 points. Alice = 13, Bob = 8, stones = [3,1,4].
// - Alice removes 3 and gets 1 + 4 = 5 points. Alice = 18, Bob = 8, stones = [1,4].
// - Bob removes 1 and gets 4 points. Alice = 18, Bob = 12, stones = [4].
// - Alice removes 4 and gets 0 points. Alice = 18, Bob = 12, stones = [].
// The score difference is 18 - 12 = 6.

// Example 2:
// Input: stones = [7,90,5,1,100,10,10,2]
// Output: 122
 
// Constraints:
// 		n == stones.length
// 		2 <= n <= 1000
// 		1 <= stones[i] <= 1000

func stoneGameVII(stones []int) int {
	l := len(stones)
	sum := make([]int, l + 1)
	for i := 1; i <= l; i = i + 1 {
		sum[i] = sum[i-1] + stones[i-1] // 前缀和
	}

	dp := make([][]int, l + 1)
	for i := 0; i <= l; i = i + 1 {
		dp[i] = make([]int, l + 1)
	}

	for i := 1; i+1 <= l; i = i + 1{   
		// 初始化情况，区间只剩2颗石头的情况下，先手拿的最大分差
		dp[i][i+1] = max(stones[i-1], stones[i]);
	}
	for n := 2; n <= l; n = n + 1 {   // 区间长度
		for i := 1; i + n <= l; i = i + 1 {   // 左端点
			j := i + n // 右端点
			dp[i][j] = max( dp[i][j], 
							// 上一个人的分差取负，就是我的分差 + 我这次获得的分数（前缀和获取）
							max(-dp[i+1][j]+sum[j]-sum[i],-dp[i][j-1]+sum[j-1]-sum[i-1]));		
			fmt.Println(dp)			
		}
	}
	return dp[1][l]
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

// best solution
func stoneGameVII1(stones []int) int {
	fmt.Println("stones: ",stones)
	n := len(stones)
	s := make([]int, n+1)
	for i, x := range stones {
		s[i+1] = s[i] + x
	}
	fmt.Println("s:", s)
	f := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[j] = max(s[j+1]-s[i+1]-f[j], s[j]-s[i]-f[j-1])
		}
	}
	return f[n-1]
}

func main() {
	fmt.Println(stoneGameVII([]int{5,3,1,4,2})) // 6
	fmt.Println(stoneGameVII([]int{7,90,5,1,100,10,10,2})) // 122

	fmt.Println(stoneGameVII1([]int{5,3,1,4,2})) // 6
	fmt.Println(stoneGameVII1([]int{7,90,5,1,100,10,10,2})) // 122
}