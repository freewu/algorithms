package main

/**
Clarification
Are a and b both 32-bit integers?

Yes.
Can I use bit operation?

Sure you can.
Example
Given a=1 and b=2 return 3.

Challenge
Of course you can just return a + b to get accepted. But Can you challenge not do it like that?(You should not use + or any arithmetic operators.)
**/

import "fmt"

/**
 * @param a: An integer
 * @param b: An integer
 * @return: The sum of a and b 
 */
func aplusb (a int, b int) int {
	// write your code here
	return a + b
}

// 使用位运算
func aplusb1 (a int, b int) int {
	// write your code here
	if 0 == a {
		return b
	}
	if 0 == b {
		return a
	}
	// a 	00000001
	// b 	00000011
	// ^ 	00000010
	var t = a ^ b // 取反 

	// a	00000001
	// b	00000011
	// & 	00000011
	// >> 1	00000110
	var t1 = (a & b) >> 1 // 处理进位 
	return aplusb1(t,t1) // 递归到 0
}

func main() {
    fmt.Println(aplusb1(1,2))
}