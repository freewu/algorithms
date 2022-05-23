package main

/**
1 · A + B Problem

Description
Write a function that add two numbers a and b, and return the answer as an integer(int).
There is no need to read data from standard input stream. Both parameters are given in function aplusb,
your job is to calculate the sum and return it.

	-2^31 <= a, b <= 2^31
	-2^31 <= a + b <= 2^31


Example 1:

	Input:
		a = 1
		b = 2
	Output:
		3
	Explanation:
		a + b = 1 + 2 = 3

Example 2:

	Input:
		a = -1
		b = 1
	Output:
		0
	Explanation:
		a + b = -1 + 1 = 0

Challenge
	Of course you can just return a + b to get accepted.
	But Can you challenge not do it like that?(You should not use + or any arithmetic operators.)
*/

import "fmt"

/**
 * @param a: An integer
 * @param b: An integer
 * @return: The sum of a and b 
 */
func Aplusb(a int, b int) int {
    // write your code here
	return a + b
}

// 不使用 + 号
func Aplusb1(a int, b int) int {
	// write your code here
	fmt.Printf("a = %#v, b = %#v\n",a,b)
	fmt.Printf("bin: a = %b, b = %b\n",a,b)
	if b == 0 { // 当 b 等于0的时候a的值就是两数的和，返回A的值
        return a;    
    }
	return Aplusb1(a ^ b,(a & b) << 1)
}

func main() {
	fmt.Printf("Aplusb(1,2) = %v\n",Aplusb(1,2)) // 3
	fmt.Printf("Aplusb(1,-1) = %v\n",Aplusb(1,-1)) // 0
	fmt.Printf("Aplusb(2,7) = %v\n",Aplusb(2,7)) // 9

	fmt.Printf("Aplusb1(1,2) = %v\n",Aplusb1(1,2)) // 3
	fmt.Printf("Aplusb1(1,-1) = %v\n",Aplusb1(1,-1)) // 0
	fmt.Printf("Aplusb1(2,7) = %v\n",Aplusb1(2,7)) // 9
}