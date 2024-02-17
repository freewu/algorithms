package main

// 70. Climbing Stairs
// You are climbing a stair case. It takes n steps to reach to the top.
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
// Note: Given n will be a positive integer.

// Example 1:
// Input: 2
// Output:  2
// Explanation:  
// 		There are two ways to climb to the top.
// 		1. 1 step + 1 step
// 		2. 2 steps

// Example 2:
// Input: 3
// Output:  3
// Explanation:  There are three ways to climb to the top.
// 		1. 1 step + 1 step + 1 step
// 		2. 1 step + 2 steps
// 		3. 2 steps + 1 step

// Example 3:
// Input: 5
// Output: 8
// Explanation:
// 		1 + 1 + 1 + 1 + 1
// 		1 + 1 + 1 + 2
// 		1 + 1 + 2 + 1
// 		1 + 2 + 1 + 1
// 		1 + 2 + 2
// 		2 + 1 + 1 + 1
// 		2 + 1 + 2
// 		2 + 2 + 1

// 1     1   (1)
// 2     2   (1 + 1 / 2) 
// 3     3   3 = 2 + 1   (1 + 1 + 1 / 2 + 1 / 1 + 2)
// 4     5   5 = 3 + 2   (1 + 1 + 1 + 1 / 1 + 1 + 2 / 1 + 2 + 1 / 2 + 1 + 1 / 2 + 2)
// 5     8   8 = 5 + 3
// 			1 + 1 + 1 + 2
// 			1 + 1 + 2 + 1
// 			1 + 2 + 1 + 1
// 			1 + 2 + 2
// 			2 + 1 + 1 + 1
// 			2 + 1 + 2
// 			2 + 2 + 1
// 6     13  13 = 8 + 5

import "fmt"
import "time"

// 递归
func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	return climbStairs(n - 1) + climbStairs(n - 2)
}

// 利用缓存 备忘录
func climbStairs1(n int) int  {
	if n == 1 || n == 2 {
		return n
	} else {
		arr := make([]int, n + 1)
		arr[1] = 1
		arr[2] = 2
		return dfs(n,arr)
	}
}

func dfs(n int,arr []int) int {
	if arr[n] != 0 {
		return arr[n]
	} else {
		arr[n] = dfs(n -1, arr) + dfs(n - 2, arr)
		return arr[n]
	}
}


// 动态规划法 (利用数组来存储)
func climbStairs2(n int) int  {
	if n == 1 || n == 2 {
		return n
	}
	arr := make([]int, n + 1)
	arr[1] = 1
	arr[2] = 2
	for i := 3 ; i <= n; i++ {
		arr[i] = arr[i - 1] + arr[i - 2]
	}
	return arr[n]
}


func main() {
	var start,end int64
	start = time.Now().UnixNano()
	fmt.Println(climbStairs(2)) // 2
	fmt.Println(climbStairs(3)) // 3
	fmt.Println(climbStairs(5)) // 8
	end = time.Now().UnixNano()
	fmt.Printf("climbStairs used: %d ns \n",end - start)

	start = time.Now().UnixNano()
	fmt.Println(climbStairs1(2)) // 2
	fmt.Println(climbStairs1(3)) // 3
	fmt.Println(climbStairs1(5)) // 8
	end = time.Now().UnixNano()
	fmt.Printf("climbStairs1 used: %d ns \n",end - start)

	start = time.Now().UnixNano()
	fmt.Println(climbStairs2(2)) // 2
	fmt.Println(climbStairs2(3)) // 3
	fmt.Println(climbStairs2(5)) // 8
	end = time.Now().UnixNano()
	fmt.Printf("climbStairs2 used: %d ns \n",end - start)
}
