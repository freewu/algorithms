package main

import "fmt"
import "math"

// 168. Excel Sheet Column Title
// Given an integer columnNumber, return its corresponding column title as it appears in an Excel sheet.

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
// Input: columnNumber = 1
// Output: "A"

// Example 2:
// Input: columnNumber = 28
// Output: "AB"

// Example 3:
// Input: columnNumber = 701
// Output: "ZY"

// Constraints:
// 		1 <= columnNumber <= 2^31 - 1


// 类似短除法的计算过程 26 进制的字母编码 按照短除法先除，然后余数逆序输出
func convertToTitle(n int) string {
	result := []byte{}
	for n > 0 {
		result = append(result, 'A' + byte((n-1)%26))
		n = (n - 1) / 26 
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}

func powInt(x, y int) int {
    return int(math.Pow(float64(x), float64(y)))
}

func convertToTitle1(columnNumber int) string {
    res := ""
    for columnNumber > 0 {
        columnNumber--
		// 取 26  余数
		// fmt.Println(columnNumber % 26)
        res = string('A' + int32(columnNumber%26)) + res
        columnNumber /= 26  // 每次循环除 26
    }
    return res

}

func main() {
	fmt.Println(convertToTitle(1)) // A
	fmt.Println(convertToTitle(28)) // AB
	fmt.Println(convertToTitle(701)) // ZY
	fmt.Println(convertToTitle(1024)) // AMJ
	fmt.Println(convertToTitle(2048)) // BZT
	fmt.Println(convertToTitle(powInt(2,31) - 1)) // FXSHRXW

	fmt.Println(convertToTitle1(1)) // A
	fmt.Println(convertToTitle1(28)) // AB
	fmt.Println(convertToTitle1(701)) // ZY
	fmt.Println(convertToTitle1(1024)) // AMJ
	fmt.Println(convertToTitle1(2048)) // BZT
	fmt.Println(convertToTitle1(powInt(2,31) - 1)) // FXSHRXW
}