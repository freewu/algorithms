package main

import "fmt"

/**
12. Integer to Roman
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

	Symbol       Value
	I             1
	V             5
	X             10
	L             50
	C             100
	D             500
	M             1000

For example, 2 is written as II in Roman numeral, just two one’s added together.
12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII.
Instead, the number four is written as IV. Because the one is before the five we subtract it making four.
The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

	I can be placed before V (5) and X (10) to make 4 and 9.
	X can be placed before L (50) and C (100) to make 40 and 90.
	C can be placed before D (500) and M (1000) to make 400 and 900.

Given an integer, convert it to a roman numeral.

Constraints:

	1 <= num <= 3999

Example 1:

	Input: num = 3
	Output: "III"

Example 2:

	Input: num = 4
	Output: "IV"

Example 3:

	Input: num = 9
	Output: "IX"

Example 4:

	Input: num = 58
	Output: "LVIII"
	Explanation: L = 50, V = 5, III = 3.

Example 5:

	Input: num = 1994
	Output: "MCMXCIV"
	Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

解题思路:
	贪心算法
	将 1-3999 范围内的罗马数字从大到小放在数组中，从头选择到尾，即可把整数转成罗马数字。

 */

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res, i := "", 0
	for num != 0 {
		for values[i] > num { // 如果能被完全包含就选择下一级的
			i++
		}
		num -= values[i]
		res += symbols[i]
	}
	return res
}

func main() {
	fmt.Printf("intToRoman(1) =  %v\n",intToRoman(1))
	fmt.Printf("intToRoman(3) =  %v\n",intToRoman(3))
	fmt.Printf("intToRoman(4) =  %v\n",intToRoman(4))
	fmt.Printf("intToRoman(5) =  %v\n",intToRoman(5))
	fmt.Printf("intToRoman(9) =  %v\n",intToRoman(9))
	fmt.Printf("intToRoman(58) =  %v\n",intToRoman(58))
	fmt.Printf("intToRoman(199) =  %v\n",intToRoman(199))
	fmt.Printf("intToRoman(1994) =  %v\n",intToRoman(1994))
	fmt.Printf("intToRoman(3998) =  %v\n",intToRoman(3998))
}
