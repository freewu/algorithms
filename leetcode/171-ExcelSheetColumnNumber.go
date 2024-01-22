package main

import "fmt"

// 171. Excel Sheet Column Number
// Given a string columnTitle that represents the column title as appears in an Excel sheet, 
// return its corresponding column number.
// For example:
// 		A -> 1
// 		B -> 2
// 		C -> 3
// 		...
// 		Z -> 26
// 		AA -> 27
// 		AB -> 28 
// 		...
 
// Example 1:
// Input: columnTitle = "A"
// Output: 1

// Example 2:
// Input: columnTitle = "AB"
// Output: 28

// Example 3:
// Input: columnTitle = "ZY"
// Output: 701
 
// Constraints:

// 		1 <= columnTitle.length <= 7
// 		columnTitle consists only of uppercase English letters.
// 		columnTitle is in the range ["A", "FXSHRXW"].

func titleToNumber(s string) int {
	val, res := 0, 0
	for i := 0; i < len(s); i++ {
		// 按照 26 进制还原成十进制
		val = int(s[i] - 'A' + 1)
		res = res*26 + val
	}
	return res
}

func main() {
	fmt.Println(titleToNumber("A")) // 1
	fmt.Println(titleToNumber("AB")) // 28
	fmt.Println(titleToNumber("ZY")) // 701
	fmt.Println(titleToNumber("AMJ")) // 1024
	fmt.Println(titleToNumber("BZT")) // 2048
	fmt.Println(titleToNumber("FXSHRXW")) // 2147483647
}