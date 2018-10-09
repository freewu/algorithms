package main

/**
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output:  321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
*/

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var t = []int{}

	for {
		//fmt.Println(x)
		//fmt.Println(x % 10)
		var s = x % 10
		x = x / 10;

		t = append(t,s)

		if 0 == x {
			//fmt.Println(t)
			break
		}
	}
	var l = len(t)
	// fmt.Println(l)
	var s = 0
	for i := 0; i < l; i++ {
		//fmt.Println(math.Pow10(l - 1))
		//fmt.Println(i)
		s += t[i] * int( math.Pow10(l - i - 1))
	}

	if s > math.MaxInt32 || s < math.MinInt32 {
		return 0
	}
    return s
}

func reverse1(x int) int {
	var s int64 = 0
	for {
		s = s * 10 + int64(x % 10)
		x /= 10;
		if x == 0 {
			break
		}
	}
	if s > math.MaxInt32 || s < math.MinInt32 {
		return 0
	}
	return int(s)
}

/*
while (x > 0) {
	res = res * 10 + x % 10;
	x /= 10;
}
if (res > INT_MAX) return 0;
if (isPositive) return res;
*/

func main() {
	fmt.Println(reverse1(123))
	fmt.Println(reverse1(120))
	fmt.Println(reverse1(-1234))

	fmt.Println(reverse(123))
	fmt.Println(reverse(120))

	fmt.Println(reverse(-123))
	fmt.Println(reverse(-1234))
}
