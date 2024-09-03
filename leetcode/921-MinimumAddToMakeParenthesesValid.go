package main

// 921. Minimum Add to Make Parentheses Valid
// A parentheses string is valid if and only if:
//     It is the empty string,
//     It can be written as AB (A concatenated with B), where A and B are valid strings, or
//     It can be written as (A), where A is a valid string.

// You are given a parentheses string s. 
// In one move, you can insert a parenthesis at any position of the string.

//     For example, if s = "()))", you can insert an opening parenthesis to be "(()))" or a closing parenthesis to be "())))".

// Return the minimum number of moves required to make s valid.
 
// Example 1:
// Input: s = "())"
// Output: 1

// Example 2:
// Input: s = "((("
// Output: 3

// Constraints:
//     1 <= s.length <= 1000
//     s[i] is either '(' or ')'.

import "fmt"

// func minAddToMakeValid(s string) int {
//     stack1, stack2 := []byte{}, []byte{}
//     for i := 0; i < len(s); i++ {
//         if s[i] == '(' {
//             stack1 = append(stack1, '(')
//             if len(stack2) > 0 {
//                 stack2 = stack2[:len(stack2)]
//             }
//         } else { // ')'
//             stack2 = append(stack1, ')')
//             if len(stack1) > 0 {
//                 stack1 = stack1[:len(stack1)]
//             }
//         }
//     }
//     abs := func(x int) int { if x < 0 { return -x; }; return x; }
//     return abs(len(stack1) - len(stack2))
// }

func minAddToMakeValid(s string) int {
    steps, stack := 0,  0
    for _, v := range s {
        if v == '(' {
            stack++
        } else if v == ')' {
            if stack == 0 {
                steps++ // "(" need to be added
            } else {
                stack--
            }
        }
    }
    // at this point if stack has some item left then closed bracket is required
    steps += stack
    return steps
}

func minAddToMakeValid1(s string) int {
    if len(s) == 0 {
        return 0
    }
    mp, stack := map[byte]byte{  ')': '(', }, []byte{ s[0] }
    for i := 1; i < len(s); i++ {
        if len(stack) > 0 && stack[len(stack)-1] == mp[s[i]] {
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, s[i])
        }
    }
    return len(stack)
}

func main() {
    // Example 1:
    // Input: s = "())"
    // Output: 1
    fmt.Println(minAddToMakeValid("())")) // 1
    // Example 2:
    // Input: s = "((("
    // Output: 3
    fmt.Println(minAddToMakeValid("(((")) // 3

    fmt.Println(minAddToMakeValid1("())")) // 1
    fmt.Println(minAddToMakeValid1("(((")) // 3
}