package main

import "fmt"

/**
17. Letter Combinations of a Phone Number
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.
A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

	2 abc
	3 def
	4 ghi
	5 jkl
	6 mno
	7 pqs
	8 tuv
	9 xyz

Constraints:

	0 <= digits.length <= 4
	digits[i] is a digit in the range ['2', '9'].

Example 1:

	Input: digits = "23"
	Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

Example 2:

	Input: digits = ""
	Output: []

Example 3:

	Input: digits = "2"
	Output: ["a","b","c"]

 */

var (
	letterMap = []string{
		" ",    //0
		"",     //1
		"abc",  //2
		"def",  //3
		"ghi",  //4
		"jkl",  //5
		"mno",  //6
		"pqrs", //7
		"tuv",  //8
		"wxyz", //9
	}
	res   = []string{}
	final = 0
)

// 解法一 DFS
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	res = []string{}
	findCombination(&digits, 0, "")
	return res
}

func findCombination(digits *string, index int, s string) {
	if index == len(*digits) {
		res = append(res, s)
		return
	}
	num := (*digits)[index]
	letter := letterMap[num-'0']
	for i := 0; i < len(letter); i++ {
		findCombination(digits, index+1, s+string(letter[i]))
	}
	return
}

// 解法二 非递归
func letterCombinations1(digits string) []string {
	if digits == "" {
		return []string{}
	}
	index := digits[0] - '0'
	letter := letterMap[index]
	tmp := []string{}
	for i := 0; i < len(letter); i++ {
		if len(res) == 0 {
			res = append(res, "")
		}
		for j := 0; j < len(res); j++ {
			tmp = append(tmp, res[j]+string(letter[i]))
		}
	}
	res = tmp
	final++
	letterCombinations(digits[1:])
	final--
	if final == 0 {
		tmp = res
		res = []string{}
	}
	return tmp
}

// 解法三 回溯（参考回溯模板，类似DFS）
var result []string
var dict = map[string][]string{
	"2" : []string{"a","b","c"},
	"3" : []string{"d", "e", "f"},
	"4" : []string{"g", "h", "i"},
	"5" : []string{"j", "k", "l"},
	"6" : []string{"m", "n", "o"},
	"7" : []string{"p", "q", "r", "s"},
	"8" : []string{"t", "u", "v"},
	"9" : []string{"w", "x", "y", "z"},
}

func letterCombinationsBT(digits string) []string {
	result = []string{}
	if digits == "" {
		return result
	}
	letterFunc("", digits)
	return result
}

func letterFunc(res string, digits string) {
	if digits == "" {
		result = append(result, res)
		return
	}
	k := digits[0:1]
	digits = digits[1:]
	for i := 0; i < len(dict[k]); i++ {
		res += dict[k][i]
		letterFunc(res, digits)
		res = res[0 : len(res)-1]
	}
}

func main() {
	fmt.Printf("letterCombinations(\"23\") = %v\n",letterCombinations("23")) //  ["ad","ae","af","bd","be","bf","cd","ce","cf"]
	fmt.Printf("letterCombinations(\"1\") = %v\n",letterCombinations("1")) // []
	fmt.Printf("letterCombinations(\"2\") = %v\n",letterCombinations("2")) //  ["a","b","c"]

	fmt.Printf("letterCombinations1(\"23\") = %v\n",letterCombinations1("23")) //  ["ad","ae","af","bd","be","bf","cd","ce","cf"]
	fmt.Printf("letterCombinations1(\"1\") = %v\n",letterCombinations1("1")) // []
	fmt.Printf("letterCombinations1(\"2\") = %v\n",letterCombinations1("2")) //  ["a","b","c"]

	fmt.Printf("letterCombinationsBT(\"23\") = %v\n",letterCombinationsBT("23")) //  ["ad","ae","af","bd","be","bf","cd","ce","cf"]
	fmt.Printf("letterCombinationsBT(\"1\") = %v\n",letterCombinationsBT("1")) // []
	fmt.Printf("letterCombinationsBT(\"2\") = %v\n",letterCombinationsBT("2")) //  ["a","b","c"]
}
