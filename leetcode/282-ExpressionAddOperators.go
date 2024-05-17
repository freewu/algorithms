package main

// 282. Expression Add Operators
// Given a string num that contains only digits and an integer target, 
// return all possibilities to insert the binary operators '+', '-', and/or '*' 
// between the digits of num so that the resultant expression evaluates to the target value.

// Note that operands in the returned expressions should not contain leading zeros.

// Example 1:
// Input: num = "123", target = 6
// Output: ["1*2*3","1+2+3"]
// Explanation: Both "1*2*3" and "1+2+3" evaluate to 6.

// Example 2:
// Input: num = "232", target = 8
// Output: ["2*3+2","2+3*2"]
// Explanation: Both "2*3+2" and "2+3*2" evaluate to 8.

// Example 3:
// Input: num = "3456237490", target = 9191
// Output: []
// Explanation: There are no expressions that can be created from "3456237490" to evaluate to 9191.
 
// Constraints:
//     1 <= num.length <= 10
//     num consists of only digits.
//     -2^31 <= target <= 2^31 - 1

import "fmt"
import "strconv"

func addOperators(num string, target int) []string {
    res := []string{}
    var backtrack func(string, string, int, int)
    backtrack = func(num, prefix string, val, last int) {
        if len(num) == 0 {
            if val == target {
                res = append(res, prefix)
            }
            return
        }
        for i := 1; i <= len(num); i++ {
            str := num[:i]
            digit, _ := strconv.Atoi(str)
            if num[0] == '0'&& i != 1 { // prevent numbers starting with 0 except the number 0 itself
                continue
            }
            if prefix == "" { // for the first call
                backtrack(num[i:], str, digit, digit)
            } else { // + | - | *  都尝试
                backtrack(num[i:], prefix + "+" + str, val + digit, digit)
                backtrack(num[i:], prefix + "-" + str, val - digit, -digit)
                backtrack(num[i:], prefix + "*" + str, val - last + last * digit, last * digit)
            }
        }
    }
    backtrack(num, "", 0, 0)
    return res
}

func addOperators1(num string, target int) []string {
    res, path := make([]string, 0), make([]byte, 0, 2*len(num))
    var backtrack func (num string, idx int, sum, pre, target int, path []byte)
    backtrack = func (num string, idx int, sum, pre, target int, path []byte) {
        if idx >= len(num) {
            if sum == target {
                res = append(res, string(path))
            }
            return
        }
        pidx := len(path)
        if idx != 0 {
            path = append(path, 0)
        }
        cur := 0
        for j := idx; j < len(num); j++ {
            if j > idx && num[idx] == '0' {
                break
            }
            path = append(path, num[j])
            cur = cur * 10 + int(num[j] - '0')
            if idx == 0 {
                backtrack(num, j+1, cur, cur, target, path)
            } else {
                path[pidx] = '+'
                backtrack(num, j+1, sum+cur, cur, target, path)
                path[pidx] = '-'
                backtrack(num, j+1, sum-cur, -cur, target, path)
                path[pidx] = '*'
                backtrack(num, j+1, sum-pre+pre*cur, pre*cur, target, path)
            }
        }
    }
    backtrack(num, 0, 0, 0, target, path)
    return res
}

func main() {
    // Example 1:
    // Input: num = "123", target = 6
    // Output: ["1*2*3","1+2+3"]
    // Explanation: Both "1*2*3" and "1+2+3" evaluate to 6.
    fmt.Println(addOperators("123", 6)) // ["1*2*3","1+2+3"]
    // Example 2:
    // Input: num = "232", target = 8
    // Output: ["2*3+2","2+3*2"]
    // Explanation: Both "2*3+2" and "2+3*2" evaluate to 8.
    fmt.Println(addOperators("232", 8)) // ["2*3+2","2+3*2"]
    // Example 3:
    // Input: num = "3456237490", target = 9191
    // Output: []
    // Explanation: There are no expressions that can be created from "3456237490" to evaluate to 9191.
    fmt.Println(addOperators("3456237490", 9191)) // []

    fmt.Println(addOperators1("123", 6)) // ["1*2*3","1+2+3"]
    fmt.Println(addOperators1("232", 8)) // ["2*3+2","2+3*2"]
    fmt.Println(addOperators1("3456237490", 9191)) // []
}