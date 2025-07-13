package main

// 3612. Process String with Special Operations I
// You are given a string s consisting of lowercase English letters and the special characters: *, #, and %.

// Build a new string result by processing s according to the following rules from left to right:
//     1. If the letter is a lowercase English letter append it to result.
//     2. A '*' removes the last character from result, if it exists.
//     3. A '#' duplicates the current result and appends it to itself.
//     4. A '%' reverses the current result.

// Return the final string result after processing all characters in s.

// Example 1:
// Input: s = "a#b%*"
// Output: "ba"
// Explanation:
// i	s[i]	Operation	Current result
// 0	'a'	Append 'a'	"a"
// 1	'#'	Duplicate result	"aa"
// 2	'b'	Append 'b'	"aab"
// 3	'%'	Reverse result	"baa"
// 4	'*'	Remove the last character	"ba"
// Thus, the final result is "ba".

// Example 2:
// Input: s = "z*#"
// Output: ""
// Explanation:
// i	s[i]	Operation	Current result
// 0	'z'	Append 'z'	"z"
// 1	'*'	Remove the last character	""
// 2	'#'	Duplicate the string	""
// Thus, the final result is "".

// Constraints:
//     1 <= s.length <= 20
//     s consists of only lowercase English letters and special characters *, #, and %.

import "fmt"

func processStr(s string) string {
    res := []rune{}
    for _, c := range s {
        switch c {
        case '*': //  删除 res 中的最后一个字符
            if len(res) > 0 { 
                res = res[:len(res) - 1]
            }
        case '#': // 复制 当前的 res 并 追加 到其自身后面
            res = append(res, res...)
        case '%': // 反转 当前的 res
            for i, j := 0, len(res) - 1; i < j; i, j = i+1, j-1 {
                res[i], res[j] = res[j], res[i]
            }
        default:
            res = append(res, c)
        }
    }
    return string(res)
}

func processStr1(s string) string {
    res := []byte{}
    for _, c := range s {
        if c == '*' {
            if len(res) > 0 {
                res = res[:len(res) - 1]
            }
        } else if c == '#' {
            res = append(res,res...)
        } else if  c== '%' {
            for i,j  := 0,len(res) - 1; i < j; {
                res[i],res[j] = res[j],res[i]
                i++
                j--
            }
        } else {
            res = append(res,byte(c))
        }
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: s = "a#b%*"
    // Output: "ba"
    // Explanation:
    // i	s[i]	Operation	Current result
    // 0	'a'	Append 'a'	"a"
    // 1	'#'	Duplicate result	"aa"
    // 2	'b'	Append 'b'	"aab"
    // 3	'%'	Reverse result	"baa"
    // 4	'*'	Remove the last character	"ba"
    // Thus, the final result is "ba".
    fmt.Println(processStr("a#b%*")) // "ba"
    // Example 2:
    // Input: s = "z*#"
    // Output: ""
    // Explanation:
    // i	s[i]	Operation	Current result
    // 0	'z'	Append 'z'	"z"
    // 1	'*'	Remove the last character	""
    // 2	'#'	Duplicate the string	""
    // Thus, the final result is "".
    fmt.Println(processStr("z*#")) // ""

    fmt.Println(processStr("blue%frog#")) // "eulbfrogeulbfrog"
    fmt.Println(processStr("leet%code#")) // "teelcodeteelcode"

    fmt.Println(processStr1("a#b%*")) // "ba"
    fmt.Println(processStr1("z*#")) // ""
    fmt.Println(processStr1("blue%frog#")) // "eulbfrogeulbfrog"
    fmt.Println(processStr1("leet%code#")) // "teelcodeteelcode"
}