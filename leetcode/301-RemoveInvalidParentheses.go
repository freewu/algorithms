package main

// 301. Remove Invalid Parentheses
// Given a string s that contains parentheses and letters, 
// remove the minimum number of invalid parentheses to make the input string valid.

// Return a list of unique strings that are valid with the minimum number of removals. 
// You may return the answer in any order.

// Example 1:
// Input: s = "()())()"
// Output: ["(())()","()()()"]

// Example 2:
// Input: s = "(a)())()"
// Output: ["(a())()","(a)()()"]

// Example 3:
// Input: s = ")("
// Output: [""]
 
// Constraints:
//     1 <= s.length <= 25
//     s consists of lowercase English letters and parentheses '(' and ')'.
//     There will be at most 20 parentheses in s.

import "fmt"

// bfs
func removeInvalidParentheses(s string) []string {
    res, seen, queue := []string{}, map[string]struct{}{}, []string{s}
    seen[s] = struct{}{}
    isValid := func (s string) bool {
        cnt := 0
        for i := range s { 
            if s[i] == '(' { // （ 则 ++
                cnt++
            } else if s[i] == ')' {
                if cnt > 0 { 
                    cnt--
                } else { // 没有配对的 ( 的直接返回 false
                    return false
                }
            }
        }
        return cnt == 0
    }
    for len(queue) > 0 && len(res) == 0 {
        for _, elem := range queue { // traverse by level
            queue = queue[1:]
            if isValid(elem) {
                res = append(res, elem)
                continue // because on current level there is minimum number of removals
            }
            if len(res) > 0 {
                continue // because on current level there is minimum number of removals
            }
            // check all substrings without one parentheses
            for i := range elem {
                if s[i] != '(' && s[i] != ')' {
                    continue
                }
                newElem := elem[:i] + elem[i+1:]
                if _, ok := seen[newElem]; ok {
                    continue
                }
                seen[newElem] = struct{}{}
                queue = append(queue, newElem)
            }
        }
    }
    return res
}

// 双指针 + dfs
func removeInvalidParentheses1(s string) []string{
    lremove, rremove, res := 0, 0, []string{}
    for _, ch := range s {
        if ch == '(' {
            lremove++
        } else if ch == ')' {
            if lremove == 0 {
                rremove++
            } else {
                lremove--
            }
        }
    }
    isValid := func (str string) bool {
        cnt := 0
        for _, ch := range str {
            if ch == '(' {
                cnt++
            } else if ch == ')' {
                cnt--
                if cnt < 0 {
                    return false
                }
            }
        }
        return cnt == 0
    }
    var dfs func(res *[]string, str string, start, left, right int)
    dfs = func(res *[]string, str string, start, left, right int) {
        if left == 0 && right == 0 {
            if isValid(str) {
                *res = append(*res, str)
            }
            return
        }
        for i := start; i < len(str); i++ {
            if i != start && str[i] == str[i-1] {
                continue
            }
            if left + right > len(str)-i { // 如果剩余的字符无法满足去掉的数量要求，直接返回
                return
            }
            if left > 0 && str[i] == '(' { // 尝试去掉一个左括号
                dfs(res, str[:i]+str[i+1:], i, left - 1, right)
            }
            if right > 0 && str[i] == ')' { // 尝试去掉一个右括号
                dfs(res, str[:i]+str[i+1:], i, left, right - 1)
            }
        }
    }
    dfs(&res, s, 0, lremove, rremove)
    return res
}

func main() {
    fmt.Println(removeInvalidParentheses("()())()")) // ["(())()","()()()"]
    fmt.Println(removeInvalidParentheses("(a)())()")) //  ["(a())()","(a)()()"]
    fmt.Println(removeInvalidParentheses(")(")) //  [""]

    fmt.Println(removeInvalidParentheses1("()())()")) // ["(())()","()()()"]
    fmt.Println(removeInvalidParentheses1("(a)())()")) //  ["(a())()","(a)()()"]
    fmt.Println(removeInvalidParentheses1(")(")) //  [""]
}