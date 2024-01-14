package main

import "fmt"

/**
6. ZigZag Conversion
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

	P   A   H   N
	A P L S I I G
	Y   I   R

And then read line by line: "PAHNAPLSIIGYIR"
Write the code that will take a string and make this conversion given a number of rows:
string convert(string s, int numRows);

Constraints:

	1 <= s.length <= 1000
	s consists of English letters (lower-case and upper-case), ',' and '.'.
	1 <= numRows <= 1000

Example 1:

	Input: s = "PAYPALISHIRING", numRows = 3
	Output: "PAHNAPLSIIGYIR"

Example 2:

	Input: s = "PAYPALISHIRING", numRows = 4
	Output: "PINALSIGYAHRPI"
	Explanation:
		P     I    N
		A   L S  I G
		Y A   H R
		P     I

Example 3:

	Input: s = "A", numRows = 1
	Output: "A"
*/

func convert(s string, numRows int) string {
	matrix, down, up := make([][]byte, numRows, numRows), 0, numRows-2
	for i := 0; i != len(s); {
		fmt.Printf("matrix: %v\n", matrix)
		if down != numRows {
			matrix[down] = append(matrix[down], byte(s[i]))
			down++
			i++
		} else if up > 0 {
			matrix[up] = append(matrix[up], byte(s[i]))
			up--
			i++
		} else {
			up = numRows - 2
			down = 0
		}
	}
	solution := make([]byte, 0, len(s))
	for _, row := range matrix {
		for _, item := range row {
			solution = append(solution, item)
		}
	}
	return string(solution)
}

func main() {
	fmt.Printf("convert(\"PAYPALISHIRING\",3) = %v\n", convert("PAYPALISHIRING", 3))
	fmt.Printf("convert(\"PAYPALISHIRING\",4) = %v\n", convert("PAYPALISHIRING", 4))
	fmt.Printf("convert(\"PAYPALISHIRING\",5) = %v\n", convert("PAYPALISHIRING", 5))
	fmt.Printf("convert(\"A\",1) = %v\n", convert("A", 1))
}
