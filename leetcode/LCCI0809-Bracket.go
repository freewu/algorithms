package main

// 面试题 08.09. Bracket LCCI
// Implement an algorithm to print all valid (e.g., properly opened and closed) combinations of n pairs of parentheses.

// Note: The result set should not contain duplicated subsets.

// For example, given n = 3, the result should be:
// [
//   "((()))",
//   "(()())",
//   "(())()",
//   "()(())",
//   "()()()"
// ]

import "fmt"
import "bytes"

func generateParenthesis(n int) []string {
    res, path := []string{}, []int{}
    var dfs func(i, balance int)
    dfs = func(i, balance int) { // balance = 左括号个数 - 右括号个数
        if len(path) == n {
            s := bytes.Repeat([]byte{')'}, n * 2)
            for _, j := range path {
                s[j] = '('
            }
            res = append(res, string(s))
            return
        }
        // 可以填 0 到 balance 个右括号
        for close := 0; close <= balance; close++ { // 填 close 个右括号
            path = append(path, i+close) // 填 1 个左括号
            dfs(i + close + 1, balance - close + 1)
            path = path[:len(path) - 1]
        }
    }
    dfs(0, 0)
    return res
}

func generateParenthesis1(n int) []string {
    var backTrack func(left, right int, cur string) []string
    backTrack = func(left, right int, cur string) []string {
        if left == 0 && right == 0 {
            return []string{ cur }
        }
        res := []string{}
        if left > 0 {
            res = append(res, backTrack(left - 1, right, cur + "(")...)
        }
        if right > left {
            res = append(res, backTrack(left, right - 1, cur + ")")...)
        }
        return res
    }
    return backTrack(n, n, "")
}



func main() {
    // For example, given n = 3, the result should be:
    // [
    //   "((()))",
    //   "(()())",
    //   "(())()",
    //   "()(())",
    //   "()()()"
    // ]
    fmt.Println(generateParenthesis(3)) // [((())) (()()) (())() ()(()) ()()()]

    fmt.Println(generateParenthesis(1)) // [()]
    fmt.Println(generateParenthesis(2)) // [(()) ()()]
    fmt.Println(generateParenthesis(4)) // [(((()))) ((()())) ((())()) ((()))() (()(())) (()()()) (()())() (())(()) (())()() ()((())) ()(()()) ()(())() ()()(()) ()()()()]
    fmt.Println(generateParenthesis(5)) // [((((())))) (((()()))) (((())())) (((()))()) (((())))() ((()(()))) ((()()())) ((()())()) ((()()))() ((())(())) ((())()()) ((())())() ((()))(()) ((()))()() (()((()))) (()(()())) (()(())()) (()(()))() (()()(())) (()()()()) (()()())() (()())(()) (()())()() (())((())) (())(()()) (())(())() (())()(()) (())()()() ()(((()))) ()((()())) ()((())()) ()((()))() ()(()(())) ()(()()()) ()(()())() ()(())(()) ()(())()() ()()((())) ()()(()()) ()()(())() ()()()(()) ()()()()()]

    fmt.Println(generateParenthesis1(3)) // [((())) (()()) (())() ()(()) ()()()]
    fmt.Println(generateParenthesis1(1)) // [()]
    fmt.Println(generateParenthesis1(2)) // [(()) ()()]
    fmt.Println(generateParenthesis1(4)) // [(((()))) ((()())) ((())()) ((()))() (()(())) (()()()) (()())() (())(()) (())()() ()((())) ()(()()) ()(())() ()()(()) ()()()()]
    fmt.Println(generateParenthesis1(5)) // [((((())))) (((()()))) (((())())) (((()))()) (((())))() ((()(()))) ((()()())) ((()())()) ((()()))() ((())(())) ((())()()) ((())())() ((()))(()) ((()))()() (()((()))) (()(()())) (()(())()) (()(()))() (()()(())) (()()()()) (()()())() (()())(()) (()())()() (())((())) (())(()()) (())(())() (())()(()) (())()()() ()(((()))) ()((()())) ()((())()) ()((()))() ()(()(())) ()(()()()) ()(()())() ()(())(()) ()(())()() ()()((())) ()()(()()) ()()(())() ()()()(()) ()()()()()]
}