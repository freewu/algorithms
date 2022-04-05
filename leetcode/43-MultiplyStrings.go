package main

import "fmt"

/**
43. Multiply Strings
Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.
Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.

Constraints:

	1 <= num1.length, num2.length <= 200
	num1 and num2 consist of digits only.
	Both num1 and num2 do not contain any leading zero, except the number 0 itself.


Example 1:

	Input: num1 = "2", num2 = "3"
	Output: "6"

Example 2:

	Input: num1 = "123", num2 = "456"
	Output: "56088"

 */

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	b1, b2, tmp := []byte(num1), []byte(num2), make([]int, len(num1)+len(num2))
	for i := 0; i < len(b1); i++ {
		for j := 0; j < len(b2); j++ {
			tmp[i+j+1] += int(b1[i]-'0') * int(b2[j]-'0')
		}
	}
	for i := len(tmp) - 1; i > 0; i-- {
		tmp[i-1] += tmp[i] / 10
		tmp[i] = tmp[i] % 10
	}
	if tmp[0] == 0 {
		tmp = tmp[1:]
	}
	res := make([]byte, len(tmp))
	for i := 0; i < len(tmp); i++ {
		res[i] = '0' + byte(tmp[i])
	}
	return string(res)
}

func main() {
	fmt.Printf(",multiply(\"2\",\"3\") = %v\n",multiply("2","3")) // "6"
	fmt.Printf(",multiply(\"123\",\"456\") = %v\n",multiply("123","456")) // "56088"
}
