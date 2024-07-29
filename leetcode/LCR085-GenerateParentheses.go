package main

// LCR 085. 括号生成
// 正整数 n 代表生成括号的对数，请设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

// 示例 1：
// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]

// 示例 2：
// 输入：n = 1
// 输出：["()"]

// 提示：
//     1 <= n <= 8

import "fmt"

// dfs
func generateParenthesis(n int) []string {
    res := []string{}
    if n == 0 {
        return res
    }
    var dfs func(l, r int, str string) 
    dfs = func(l, r int, str string) {
        if l == 0 && r == 0 {
            res = append(res, str)
            return
        }
        if l > 0 {
            dfs(l-1, r, str + "(")
        }
        if r > 0 && l < r {
            dfs(l, r - 1, str + ")")
        }
    }
    dfs(n, n, "")
    return res
}

// best solution
func generateParenthesis1(n int) []string {
    res := []string{}
    var dfs func (n int, o int, c int, cur string)
    dfs = func (n int, o int, c int, cur string) {
        if n == o && n == c {
            res = append(res, cur)
            return
        }
        if o < n {
            dfs(n, o+1, c, cur + "(")
        }
        if c < o {
            dfs(n, o, c+1, cur + ")")
        }
    }
    dfs(n, 0, 0, "")
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