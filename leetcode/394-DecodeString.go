package main

// 394. Decode String
// Given an encoded string, return its decoded string.
// The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. 
// Note that k is guaranteed to be a positive integer.
// You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc.
// Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. 
// For example, there will not be input like 3a or 2[4].

// The test cases are generated so that the length of the output will never exceed 10^5.

// Example 1:
// Input: s = "3[a]2[bc]"
// Output: "aaabcbc"

// Example 2:
// Input: s = "3[a2[c]]"
// Output: "accaccacc"

// Example 3:
// Input: s = "2[abc]3[cd]ef"
// Output: "abcabccdcdcdef"
 
// Constraints:
//     1 <= s.length <= 30
//     s consists of lowercase English letters, digits, and square brackets '[]'.
//     s is guaranteed to be a valid input.
//     All the integers in s are in the range [1, 300].

import "fmt"
import "strings"
import "unicode"
import "strconv"

func decodeString(s string) string {
    res := strings.Builder{}
    for i := 0; i < len(s); i++ {
        if s[i] >= '0' && s[i] <= '9' {
            count := 0
            for i < len(s) && s[i] != '[' {
                count = count * 10 + int(s[i] - '0')
                i++
            }
            startIndex := i + 1
            for j := 1; j > 0; i++ {
                if s[i + 1] == '[' { j++ }
                if s[i + 1] == ']' { j-- }
            }
            for k := 0; k < count; k++ {
                res.WriteString(decodeString(s[startIndex:i]))
            }
        } else {
            res.WriteByte(s[i])
        }
    }
    return res.String()
}

// stack
type Stack []string

func (s *Stack) IsEmpty() bool {
    return len(*s) <= 0
}

func(s *Stack) Push(item string){
    *s = append(*s , item)
}

func (s *Stack) Pop() (string){
    if s.IsEmpty(){
        return ""
    }
    item := (*s)[len(*s)-1]
    (*s) = (*s)[:len(*s)-1]
    
    return item
}

func (s *Stack) ToString() string{
    result := ""
    for !s.IsEmpty(){
        result+=s.Pop()
    }
    return result
}

func (s *Stack) Peek() string{
    item := (*s)[len(*s)-1]
    return item
}

func decodeString1(s string) string {
    var stack Stack
    res := ""
    isDigit := func (s string) bool {
        for _, c := range s {
            if !unicode.IsDigit(c) {
                return false
            }
        }
        return true
    }
    for _,item := range s {
        charAt := string(item)  
        if charAt != "]"{
            stack.Push(charAt)
        } else {
            item := ""
            for stack.Peek() != "["{
                item = stack.Pop() + item
            }
            //remove "["
            stack.Pop()
            digit := ""
            for !stack.IsEmpty() && isDigit(stack.Peek()){
                digit = stack.Pop() + digit
            }
            k,_ :=  strconv.ParseInt(digit, 10,32)
            subStr := ""
            for i := 0 ; i< int(k); i++{
                subStr += item
            }
            stack.Push(subStr)
        }
        
    }
    for !stack.IsEmpty(){
        res = stack.Pop() + res
    }
    return res
}


// 外层的解码需要等待内层解码的结果。先扫描的字符还用不上，但不能忘了它们。
// 我们准备由内到外，层层解决[ ]，需要保持对字符的记忆，于是用栈。
func decodeString2(s string) string {
    numStack, strStack := make([]int, 0), make([]string, 0)
    num, res := 0, ""
    for _, ch := range s {
        if ch >= '0' && ch <= '9' { // 叠加十进制数字
            num = num*10 + int(ch-'0')
        } else if ch == '[' { // 数字和字符串分别入栈、置空
            numStack = append(numStack, num)
            num = 0
            strStack = append(strStack, res)
            res = ""
        } else if ch == ']' { // 数字和字符串分别出栈、拼接
            count := numStack[len(numStack)-1]
            numStack = numStack[:len(numStack)-1]
            str := strStack[len(strStack)-1]
            strStack = strStack[:len(strStack)-1]
            res = str + strings.Repeat(res, count) // ！！！当前res*count，再接在str后面
        } else { // ['a','z'] || ['A','Z']
            res += string(ch) // 叠加字符
        }
    }
    return res
}

func main() {
    fmt.Println(decodeString("3[a]2[bc]")) // aaabcbc
    fmt.Println(decodeString("3[a2[c]]")) // accaccacc
    fmt.Println(decodeString("2[abc]3[cd]ef")) // abcabccdcdcdef

    fmt.Println(decodeString1("3[a]2[bc]")) // aaabcbc
    fmt.Println(decodeString1("3[a2[c]]")) // accaccacc
    fmt.Println(decodeString1("2[abc]3[cd]ef")) // abcabccdcdcdef

    fmt.Println(decodeString2("3[a]2[bc]")) // aaabcbc
    fmt.Println(decodeString2("3[a2[c]]")) // accaccacc
    fmt.Println(decodeString2("2[abc]3[cd]ef")) // abcabccdcdcdef
}