package main

// 201. Bitwise AND of Numbers Range
// Given two integers left and right that represent the range [left, right], 
// return the bitwise AND of all numbers in this range, inclusive.

// Example 1:
// Input: left = 5, right = 7
// Output: 4

// Example 2:
// Input: left = 0, right = 0
// Output: 0

// Example 3:
// Input: left = 1, right = 2147483647
// Output: 0
 
// Constraints:
// 		0 <= left <= right <= 2^31 - 1

import "fmt"

func rangeBitwiseAnd2(left int, right int) int {
    start := left
	res := left
	for start < right {
		res = res & start
		start += 1
		//fmt.Println("start: ",start)
	}
	return res
}

// 解法一
func rangeBitwiseAnd1(m int, n int) int {
	if m == 0 {
		return 0
	}
	moved := 0
	for m != n {
		m >>= 1
		n >>= 1
		moved++
	}
	return m << uint32(moved)
}

// Brian Kernighan's algorithm
func rangeBitwiseAnd(m int, n int) int {
	for n > m {
		n &= (n - 1) // 清除最低位的 1
	}
	return n
}

func main() {
	fmt.Println(rangeBitwiseAnd(5,7)) // 4
	fmt.Println(rangeBitwiseAnd(0,0)) // 0
	fmt.Println(rangeBitwiseAnd(1,2147483647)) // 0

	fmt.Println(rangeBitwiseAnd1(5,7)) // 4
	fmt.Println(rangeBitwiseAnd1(0,0)) // 0
	fmt.Println(rangeBitwiseAnd1(1,2147483647)) // 0

	fmt.Println(rangeBitwiseAnd2(5,7)) // 4
	fmt.Println(rangeBitwiseAnd2(0,0)) // 0
	fmt.Println(rangeBitwiseAnd2(1,2147483647)) // 0
}