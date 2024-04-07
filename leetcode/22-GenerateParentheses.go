package main

// 22. Generate Parentheses
// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

// Example 1:
// Input: n = 3
// Output: ["((()))","(()())","(())()","()(())","()()()"]

// Example 2:
// Input: n = 1
// Output: ["()"]

// Constraints:
//     1 <= n <= 8

import "fmt"

// dfs
func generateParenthesis(n int) []string {
    res := []string{}
    if n == 0 {
        return res
    }
    var dfs func(l, r int, str string, res *[]string) 
    dfs = func(l, r int, str string, res *[]string) {
        if l == 0 && r == 0 {
            *res = append(*res, str)
            return
        }
        if l > 0 {
            dfs(l-1, r, str + "(", res)
        }
        if r > 0 && l < r {
            dfs(l, r - 1, str + ")", res)
        }
    }
    dfs(n, n, "", &res)
    return res
}

// best solution
func generateParenthesis1(n int) []string {
	res := []string{}
    var dfs func (n int, o int, c int, cur string, res *[]string)
    dfs = func (n int, o int, c int, cur string, res *[]string) {
        if n == o && n == c {
            *res = append(*res, cur)
            return
        }
        if o < n {
            dfs(n, o+1, c, cur + "(", res)
        }
        if c < o {
            dfs(n, o, c+1, cur + ")", res)
        }
    }
    dfs(n, 0, 0, "", &res)
    return res
}

func main() {
    fmt.Printf("generateParenthesis1(1) = %v\n",generateParenthesis1(1)) // [()]
    fmt.Printf("generateParenthesis1(2) = %v\n",generateParenthesis1(2)) // [(()) ()()]
    fmt.Printf("generateParenthesis1(3) = %v\n",generateParenthesis1(3)) // [((())) (()()) (())() ()(()) ()()()]
    // fmt.Printf("generateParenthesis1(4) = %v\n",generateParenthesis1(4))
    // fmt.Printf("generateParenthesis1(5) = %v\n",generateParenthesis1(5))
    // fmt.Printf("generateParenthesis1(6) = %v\n",generateParenthesis1(6))
    // fmt.Printf("generateParenthesis1(7) = %v\n",generateParenthesis1(7))
    // fmt.Printf("generateParenthesis1(8) = %v\n",generateParenthesis1(8))

    fmt.Printf("generateParenthesis(1) = %v\n",generateParenthesis(1)) // [()]
    fmt.Printf("generateParenthesis(2) = %v\n",generateParenthesis(2)) // [(()) ()()]
    fmt.Printf("generateParenthesis(3) = %v\n",generateParenthesis(3)) // [((())) (()()) (())() ()(()) ()()()]
    // fmt.Printf("generateParenthesis(4) = %v\n",generateParenthesis(4))
    // fmt.Printf("generateParenthesis(5) = %v\n",generateParenthesis(5))
    // fmt.Printf("generateParenthesis(6) = %v\n",generateParenthesis(6))
    // fmt.Printf("generateParenthesis(7) = %v\n",generateParenthesis(7))
    // fmt.Printf("generateParenthesis(8) = %v\n",generateParenthesis(8))
}
