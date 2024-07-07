package main

// 722. Remove Comments
// Given a C++ program, remove comments from it. 
// The program source is an array of strings source where source[i] is the ith line of the source code. 
// This represents the result of splitting the original source code string by the newline character '\n'.

// In C++, there are two types of comments, line comments, and block comments.
//     The string "//" denotes a line comment, which represents that it and the rest of the characters to the right of it in the same line should be ignored.
//     The string "/*" denotes a block comment, which represents that all characters until the next (non-overlapping) occurrence of "*/" should be ignored. (Here, occurrences happen in reading order: line by line from left to right.) To be clear, the string "/*/" does not yet end the block comment, as the ending would be overlapping the beginning.

// The first effective comment takes precedence over others.
//     For example, if the string "//" occurs in a block comment, it is ignored.
//     Similarly, if the string "/*" occurs in a line or block comment, it is also ignored.

// If a certain line of code is empty after removing comments, you must not output that line: each string in the answer list will be non-empty.

// There will be no control characters, single quote, or double quote characters.

//     For example, source = "string s = "/* Not a comment. */";" will not be a test case.
    
// Also, nothing else such as defines or macros will interfere with the comments.
// It is guaranteed that every open block comment will eventually be closed, so "/*" outside of a line or block comment always starts a new comment.

// Finally, implicit newline characters can be deleted by block comments. 
// Please see the examples below for details.
// After removing the comments from the source code, return the source code in the same format.

// Example 1:
// Input: source = ["/*Test program */", "int main()", "{ ", "  // variable declaration ", "int a, b, c;", "/* This is a test", "   multiline  ", "   comment for ", "   testing */", "a = b + c;", "}"]
// Output: ["int main()","{ ","  ","int a, b, c;","a = b + c;","}"]
// Explanation: The line by line code is visualized as below:
// /*Test program */
// int main()
// { 
//   // variable declaration 
// int a, b, c;
// /* This is a test
//    multiline  
//    comment for 
//    testing */
// a = b + c;
// }
// The string /* denotes a block comment, including line 1 and lines 6-9. The string // denotes line 4 as comments.
// The line by line output code is visualized as below:
// int main()
// { 
// int a, b, c;
// a = b + c;
// }

// Example 2:
// Input: source = ["a/*comment", "line", "more_comment*/b"]
// Output: ["ab"]
// Explanation: The original source string is "a/*comment\nline\nmore_comment*/b", where we have bolded the newline characters.  After deletion, the implicit newline characters are deleted, leaving the string "ab", which when delimited by newline characters becomes ["ab"].

// Constraints:
//     1 <= source.length <= 100
//     0 <= source[i].length <= 80
//     source[i] consists of printable ASCII characters.
//     Every open block comment is eventually closed.
//     There are no single-quote or double-quote in the input.

import "fmt"

func removeComments(source []string) []string {
    res, cur, block := []string{}, "", false
    for _, line := range source {
        n := len(line)
        for i := 0; i < n; i++ {
            char, isLast, next := line[i], false, byte(' ')
            if i == n - 1 {
                isLast = true
            } else {
                next = line[i + 1]
            }
            if block {
                if char == '*' && !isLast && next == '/' {
                    block = false
                    i++
                }
            } else {
                if char == '/' && !isLast && (next == '/' || next == '*') {
                    if next == '*' {
                        block = true
                        i++
                    } else {
                        break
                    }
                } else {
                    cur += string(char)
                }
            }
        }
        if !block && cur != "" {
            res = append(res, cur)
            cur = ""
        }
    }
    return res
}

func removeComments1(source []string) []string {
    res, new_line, in_block := []string{}, []byte{}, false
    for _, line := range source {
        for i := 0; i < len(line); i++ {
            if in_block { // "/*" 里内容都是注释直到找到  "*/"
                if i + 1 < len(line) && line[i] == '*' && line[i + 1] == '/' { // */ 注释块结束
                    in_block = false
                    i++
                }
            } else {
                if i + 1 < len(line) && line[i] == '/' && line[i + 1] == '*' { // /* 注释块开始
                    in_block = true
                    i++
                } else if i + 1 < len(line) && line[i] == '/' && line[i + 1] == '/' {  // 单行注释开始
                    break
                } else {
                    new_line = append(new_line, line[i])
                }
            }
        }
        if  !in_block && len(new_line) > 0 {
            res = append(res, string(new_line))
            new_line = []byte{}
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: source = ["/*Test program */", "int main()", "{ ", "  // variable declaration ", "int a, b, c;", "/* This is a test", "   multiline  ", "   comment for ", "   testing */", "a = b + c;", "}"]
    // Output: ["int main()","{ ","  ","int a, b, c;","a = b + c;","}"]
    // Explanation: The line by line code is visualized as below:
    // /*Test program */
    // int main()
    // { 
    //   // variable declaration 
    // int a, b, c;
    // /* This is a test
    //    multiline  
    //    comment for 
    //    testing */
    // a = b + c;
    // }
    // The string /* denotes a block comment, including line 1 and lines 6-9. The string // denotes line 4 as comments.
    // The line by line output code is visualized as below:
    // int main()
    // { 
    // int a, b, c;
    // a = b + c;
    // }
    source1 := []string{
        "/*Test program */", 
        "int main()", 
        "{ ", 
        "  // variable declaration ", 
        "int a, b, c;", 
        "/* This is a test", 
        "   multiline  ", 
        "   comment for ", 
        "   testing */", 
        "a = b + c;", 
        "}",
    }
    fmt.Println(removeComments(source1)) // ["int main()","{ ","  ","int a, b, c;","a = b + c;","}"]
    // Example 2:
    // Input: source = ["a/*comment", "line", "more_comment*/b"]
    // Output: ["ab"]
    // Explanation: The original source string is "a/*comment\nline\nmore_comment*/b", where we have bolded the newline characters.  After deletion, the implicit newline characters are deleted, leaving the string "ab", which when delimited by newline characters becomes ["ab"].
    source2 := []string{
        "a/*comment", 
        "line", 
        "more_comment*/b",
    }
    fmt.Println(removeComments(source2)) // ["ab"]

    fmt.Println(removeComments1(source1)) // ["int main()","{ ","  ","int a, b, c;","a = b + c;","}"]
    fmt.Println(removeComments1(source2)) // ["ab"]
}