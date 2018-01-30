package main

/*
Given a roman numeral, convert it to an integer.

Input is guaranteed to be within the range from 1 to 3999.

('I', 1);
('V', 5);
('X', 10);
('L', 50);
('C', 100);
('D', 500);
('M', 1000);
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

func main() {
	fmt.Println(romanToInt("VII"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("VIII"))
	fmt.Println(romanToInt("V"))
	fmt.Println(romanToInt("XV"))
	fmt.Println(romanToInt("DCXXI")) // 621
}