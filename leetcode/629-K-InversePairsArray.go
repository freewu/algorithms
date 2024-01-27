package main

import "fmt"

// 629. K Inverse Pairs Array
// For an integer array nums：
// 		an inverse pair is a pair of integers [i, j] 
// 		where 0 <= i < j < nums.length and nums[i] > nums[j].
// Given two integers n and k, 
// return the number of different arrays consist of numbers from 1 to n such that there are exactly k inverse pairs. 
// Since the answer can be huge, return it modulo 109 + 7.

// Example 1:
// Input: n = 3, k = 0
// Output: 1
// Explanation: Only the array [1,2,3] which consists of numbers from 1 to 3 has exactly 0 inverse pairs.

// Example 2:
// Input: n = 3, k = 1
// Output: 2
// Explanation: The array [1,3,2] and [2,1,3] have exactly 1 inverse pair.

// Constraints:

// 		1 <= n <= 1000
// 		0 <= k <= 1000

// dp
func kInversePairs(n int, k int) int {
	M := 1000000007;
	dp := make([][]int, n + 1)
	// 先构造一个 (n + 1) * (k + 1) 的数组
	for i := 0; i < n + 1; i++ {
		dp[i] = make([]int, k + 1)
	}
	dp[0][0] = 1
	for  i := 0; i <= n; i = i + 1 {
		for  j := 0; j < i; j = j + 1 {
			for  m := 0; m <= k; m = m + 1  {
				if m - j >= 0 && m - j <= k {
					dp[i][m] = (dp[i][m] + dp[i - 1][m - j]) % M
				}
			}
		}
	}
	return dp[n][k]
}

// best solution
const mod = 1e9 + 7
func kInversePairs1(n, k int) int {
	dp := make([]int, k+1)

	dp[0] = 1
	queue := make([]int, k+1)

	for i := 2; i <= n; i++ {
		queue[0] = 1
		sum := 1

		for j := 1; j <= k; j++ {
			sum = (sum + dp[j]) % mod
			if j >= i {
				sum = (sum + mod - queue[j%i]) % mod
			}

			queue[j%i] = dp[j]
			dp[j] = sum

			fmt.Println(dp)
		}
	}

	return dp[k]
}

func main() {
	fmt.Println(kInversePairs(3,0)) // 1
	fmt.Println(kInversePairs(3,1)) // 2

	fmt.Println(kInversePairs1(3,0)) // 1
	fmt.Println(kInversePairs1(3,1)) // 2
}