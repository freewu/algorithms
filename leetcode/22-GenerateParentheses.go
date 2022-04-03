package main

import "fmt"

/**
22. Generate Parentheses
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

Constraints:

	1 <= n <= 8

Example 1:

	Input: n = 3
	Output: ["((()))","(()())","(())()","()(())","()()()"]

Example 2:

	Input: n = 1
	Output: ["()"]

 */

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	var res []string
	findGenerateParenthesis(n, n, "", &res)
	return res
}

func findGenerateParenthesis(lindex, rindex int, str string, res *[]string) {
	if lindex == 0 && rindex == 0 {
		*res = append(*res, str)
		return
	}
	if lindex > 0 {
		findGenerateParenthesis(lindex-1, rindex, str+"(", res)
	}
	if rindex > 0 && lindex < rindex {
		findGenerateParenthesis(lindex, rindex-1, str+")", res)
	}
}

// best solution
func generateParenthesisBest(n int) []string {
	var result []string
	helper(n, 0, 0, "", &result)
	return result
}

func helper(n int, o int, c int, cur string, result *[]string) {
	if n == o && n == c {
		*result = append(*result, cur)
		return
	}
	if o < n {
		helper(n, o+1, c, cur + "(", result)
	}
	if c < o {
		helper(n, o, c+1, cur + ")", result)
	}
}

func main() {
	fmt.Printf("generateParenthesisBest(1) = %v\n",generateParenthesisBest(1))
	fmt.Printf("generateParenthesisBest(2) = %v\n",generateParenthesisBest(2))
	fmt.Printf("generateParenthesisBest(3) = %v\n",generateParenthesisBest(3))
	fmt.Printf("generateParenthesisBest(4) = %v\n",generateParenthesisBest(4))
	fmt.Printf("generateParenthesisBest(5) = %v\n",generateParenthesisBest(5))
	fmt.Printf("generateParenthesisBest(6) = %v\n",generateParenthesisBest(6))
	fmt.Printf("generateParenthesisBest(7) = %v\n",generateParenthesisBest(7))
	fmt.Printf("generateParenthesisBest(8) = %v\n",generateParenthesisBest(8))

	fmt.Printf("generateParenthesis(1) = %v\n",generateParenthesis(1))
	fmt.Printf("generateParenthesis(2) = %v\n",generateParenthesis(2))
	fmt.Printf("generateParenthesis(3) = %v\n",generateParenthesis(3))
	fmt.Printf("generateParenthesis(4) = %v\n",generateParenthesis(4))
	fmt.Printf("generateParenthesis(5) = %v\n",generateParenthesis(5))
	fmt.Printf("generateParenthesis(6) = %v\n",generateParenthesis(6))
	fmt.Printf("generateParenthesis(7) = %v\n",generateParenthesis(7))
	fmt.Printf("generateParenthesis(8) = %v\n",generateParenthesis(8))
}
