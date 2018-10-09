package main

import (
	"fmt"
)

/*
Implement int sqrt(int x).

Compute and return the square root of x.

x is guaranteed to be a non-negative integer.

Example 1:

Input: 4
Output: 2
Example 2:

Input: 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since we want to return an integer, the decimal part part will be truncated.

1 1
2 1
3 1
4 2
5 2
6 2
7 2
8 2
9 3
*/

// 二分搜索
func mySqrt1(x int) int {
	if x == 0 {
		return 0
	}
	var l = 0
	var r = (x / 2) + 1

	for {
		if l > r {
			break
		}
		var m = (l + r) / 2 // 取均值
		var s = m * m       // 均值的平方
		if s == x {         // 1 / 4 / 9 的情况
			return m
		} else if s < x { //
			l = m + 1
		} else {
			r = m - 1
		}

		fmt.Println("m: ", m, " r: ", r, " l: ", l)
	}
	return r
}

// 牛顿法
// http://www.cnblogs.com/AnnieKim/archive/2013/04/18/3028607.html
// xi+1=xi - (xi2 - n) / (2xi) = xi - xi / 2 + n / (2xi) = xi / 2 + n / 2xi = (xi + n/xi) / 2
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	var r = 1
	var p = 0
	for {
		if r == p {
			break
		}
		p = r
		r = (r + (x / r)) / 2
		fmt.Println("p:", p, " r:", r)
	}
	return r
}

func main() {
	fmt.Println(mySqrt(10)) // 3
	// fmt.Println(mySqrt(2))  // 1
	// fmt.Println(mySqrt(3))  // 1
	// fmt.Println(mySqrt(4))  // 2
	// fmt.Println(mySqrt(8))  // 2
	// fmt.Println(mySqrt(9))  // 3
	// fmt.Println(mySqrt(15)) // 3
	// fmt.Println(mySqrt(16)) // 4
	// fmt.Println(mySqrt(24)) // 4
	// fmt.Println(mySqrt(25)) // 5
}
