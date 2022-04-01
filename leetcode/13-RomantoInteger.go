package main

/*
13. Roman to Integer
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

	Symbol       Value
	I             1
	V             5
	X             10
	L             50
	C             100
	D             500
	M             1000
For example, two is written as II in Roman numeral, just two one’s added together.
Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.
Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII.
Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

	I can be placed before V (5) and X (10) to make 4 and 9.
	X can be placed before L (50) and C (100) to make 40 and 90.
	C can be placed before D (500) and M (1000) to make 400 and 900.

Given a roman numeral, convert it to an integer. Input is guaranteed to be within the range from 1 to 3999.

Example 1:

	Input: "III"
	Output: 3

Example 2:

	Input: "IV"
	Output: 4

Example 3:

	Input: "IX"
	Output: 9

Example 4:

	Input: "LVIII"
	Output: 58
	Explanation: L = 50, V= 5, III = 3.

Example 5:

	Input: "MCMXCIV"
	Output: 1994
	Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
*/

import (
	"fmt"
)

var m = map[string]int{"I":1,"V":5,"X":10,"L":50,"C":100,"D":500,"M":1000}

func romanToInt(s string) int {
	
	//fmt.Println(m)
	//fmt.Println(len(s))
	//fmt.Println(m[string(s[0])])
	var n = m[string(s[0])]

	for i := 1; i < len(s); i++ {
		//fmt.Println(string(s[i]))
		//fmt.Println(string(s[i-1]))

		if m[string(s[i])] > m[string(s[i - 1])] {
			// 1 + 5 - 2 * 1 IV
			n = n + m[string(s[i])] - 2 * m[string(s[i - 1])]
		} else {
			n += m[string(s[i])]
		}
		//fmt.Println(n)
	}
    return n
}

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt1(s string) int {
	if s == "" {
		return 0
	}
	num, lastint, total := 0, 0, 0
	for i := 0; i < len(s); i++ {
		char := s[len(s)-(i+1) : len(s)-i]
		num = roman[char]
		if num < lastint {
			total = total - num
		} else {
			total = total + num
		}
		lastint = num
	}
	return total
}

// best solution
func romanToIntBest(s string)  int{
	k := []rune(s)
	num := 0

	for i:=0; i<len(k);i++{
		if i < len(k)-1 && roman_nums(k[i]) < roman_nums(k[i+1]){
			num+= roman_nums(k[i+1])- roman_nums(k[i])
			i= i+1
		}else{
			num+=roman_nums(k[i])
		}
	}
	return num
}

func roman_nums(numbers rune) int{
	switch numbers{
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}

func main() {
	fmt.Printf("romanToInt(\"VII\") = %v\n",romanToInt("VII"))
	fmt.Printf("romanToInt(\"IV\") = %v\n",romanToInt("IV"))
	fmt.Printf("romanToInt(\"IX\") = %v\n",romanToInt("IX"))
	fmt.Printf("romanToInt(\"VIII\") = %v\n",romanToInt("VIII"))
	fmt.Printf("romanToInt(\"V\") = %v\n",romanToInt("V"))
	fmt.Printf("romanToInt(\"XV\") = %v\n",romanToInt("XV"))
	fmt.Printf("romanToInt(\"DCXXI\") = %v\n",romanToInt("DCXXI")) // 621

	fmt.Printf("romanToInt1(\"VII\") = %v\n",romanToInt1("VII"))
	fmt.Printf("romanToInt1(\"IV\") = %v\n",romanToInt1("IV"))
	fmt.Printf("romanToInt1(\"IX\") = %v\n",romanToInt1("IX"))
	fmt.Printf("romanToInt1(\"VIII\") = %v\n",romanToInt1("VIII"))
	fmt.Printf("romanToInt1(\"V\") = %v\n",romanToInt1("V"))
	fmt.Printf("romanToInt1(\"XV\") = %v\n",romanToInt1("XV"))
	fmt.Printf("romanToInt1(\"DCXXI\") = %v\n",romanToInt1("DCXXI")) // 621

	fmt.Printf("romanToIntBest(\"VII\") = %v\n",romanToIntBest("VII"))
	fmt.Printf("romanToIntBest(\"IV\") = %v\n",romanToIntBest("IV"))
	fmt.Printf("romanToIntBest(\"IX\") = %v\n",romanToIntBest("IX"))
	fmt.Printf("romanToIntBest(\"VIII\") = %v\n",romanToIntBest("VIII"))
	fmt.Printf("romanToIntBest(\"V\") = %v\n",romanToIntBest("V"))
	fmt.Printf("romanToIntBest(\"XV\") = %v\n",romanToIntBest("XV"))
	fmt.Printf("romanToIntBest(\"DCXXI\") = %v\n",romanToIntBest("DCXXI")) // 621
}