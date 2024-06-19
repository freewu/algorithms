package main

// 640. Solve the Equation
// Solve a given equation and return the value of 'x' in the form of a string "x=#value". 
// The equation contains only '+', '-' operation, the variable 'x' and its coefficient. 
// You should return "No solution" if there is no solution for the equation, or "Infinite solutions" if there are infinite solutions for the equation.

// If there is exactly one solution for the equation, we ensure that the value of 'x' is an integer.

// Example 1:
// Input: equation = "x+5-3+x=6+x-2"
// Output: "x=2"

// Example 2:
// Input: equation = "x=x"
// Output: "Infinite solutions"

// Example 3:
// Input: equation = "2x=x"
// Output: "x=0"

// Constraints:
//     3 <= equation.length <= 1000
//     equation has exactly one '='.
//     equation consists of integers with an absolute value in the range [0, 100] without any leading zeros, and the variable 'x'.

import "fmt"
import "strings"
import "strconv"

func solveEquation(equation string) string {
    s := strings.Split(equation, "=")
    if len(s) != 2 {
        return "No solution"
    }
    parse := func(equation string) (int, int) {
        equation += "+" // to simplify the code
        sign, current, coefficient, number := 1, "", 0, 0
        for _, v := range equation {
            switch v {
            case '+':
                n, _ := strconv.Atoi(current)
                number += sign * n
                sign = 1
                current = ""
            case '-':
                n, _ := strconv.Atoi(current)
                number += sign * n
                sign = -1
                current = ""
            case 'x':
                if current == "" { // if current is empty, set it "1"
                    current = "1"
                }
                n, _ := strconv.Atoi(current)
                coefficient += sign * n
                current = "0"
            default:
                current += string(v)
            }
        }
        return coefficient, number
    }
    left, right := s[0], s[1]
    lx, ln := parse(left)
    rx, rn := parse(right)
    x := lx - rx
    n := rn - ln
    if x == 0 {
        if n == 0 {
            return "Infinite solutions"
        }
        return "No solution"
    }
    return "x=" + strconv.Itoa(n / x)
}

func main() {
    // Example 1:
    // Input: equation = "x+5-3+x=6+x-2"
    // Output: "x=2"
    fmt.Println(solveEquation("x+5-3+x=6+x-2")) // x=2
    // Example 2:
    // Input: equation = "x=x"
    // Output: "Infinite solutions"
    fmt.Println(solveEquation("x=x")) // "Infinite solutions"
    // Example 3:
    // Input: equation = "2x=x"
    // Output: "x=0"
    fmt.Println(solveEquation("2x=x")) // x=0
}