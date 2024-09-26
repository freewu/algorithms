package main

// 1678. Goal Parser Interpretation
// You own a Goal Parser that can interpret a string command. 
// The command consists of an alphabet of "G", "()" and/or "(al)" in some order. 
// The Goal Parser will interpret "G" as the string "G", "()" as the string "o", and "(al)" as the string "al". 
// The interpreted strings are then concatenated in the original order.

// Given the string command, return the Goal Parser's interpretation of command.

// Example 1:
// Input: command = "G()(al)"
// Output: "Goal"
// Explanation: The Goal Parser interprets the command as follows:
// G -> G
// () -> o
// (al) -> al
// The final concatenated result is "Goal".

// Example 2:
// Input: command = "G()()()()(al)"
// Output: "Gooooal"

// Example 3:
// Input: command = "(al)G(al)()()G"
// Output: "alGalooG"

// Constraints:
//     1 <= command.length <= 100
//     command consists of "G", "()", and/or "(al)" in some order.

import "fmt"

func interpret(command string) string {
    res, n := []byte{}, len(command)
    for i := 0; i < n; i++ {
        if command[i] == '(' && i < n {
            if command[i + 1] == ')' { // () => o
                res = append(res, 'o')
            }
        } else if command[i] != '(' && command[i] != ')' {
            res = append(res, command[i])
        }
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: command = "G()(al)"
    // Output: "Goal"
    // Explanation: The Goal Parser interprets the command as follows:
    // G -> G
    // () -> o
    // (al) -> al
    // The final concatenated result is "Goal".
    fmt.Println(interpret("G()(al)")) // "Goal"
    // Example 2:
    // Input: command = "G()()()()(al)"
    // Output: "Gooooal"
    fmt.Println(interpret("G()()()()(al)")) // "Gooooal"
    // Example 3:
    // Input: command = "(al)G(al)()()G"
    // Output: "alGalooG"
    fmt.Println(interpret("(al)G(al)()()G")) // "alGalooG"
}