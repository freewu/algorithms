package main

// 592. Fraction Addition and Subtraction
// Given a string expression representing an expression of fraction addition and subtraction, 
// return the calculation result in string format.

// The final result should be an irreducible fraction. 
// If your final result is an integer, change it to the format of a fraction that has a denominator 1. 
// So in this case, 2 should be converted to 2/1.

// Example 1:
// Input: expression = "-1/2+1/2"
// Output: "0/1"

// Example 2:
// Input: expression = "-1/2+1/2+1/3"
// Output: "1/3"

// Example 3:
// Input: expression = "1/3-1/2"
// Output: "-1/6"

// Constraints
//     The input string only contains '0' to '9', '/', '+' and '-'. So does the output.
//     Each fraction (input and output) has the format ±numerator/denominator. If the first input fraction or the output is positive, then '+' will be omitted.
//     The input only contains valid irreducible fractions, where the numerator and denominator of each fraction will always be in the range [1, 10]. If the denominator is 1, it means this fraction is actually an integer in a fraction format defined above.
//     The number of given fractions will be in the range [1, 10].
//     The numerator and denominator of the final result are guaranteed to be valid and in the range of 32-bit int.

import "fmt"
import "strconv"

// stack
func fractionAddition(expression string) string {
    convertToInt := func (st []rune) int {
        n, _ := strconv.Atoi(string(st))
        return n
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    calc := func (nums []int) []int {
        a, b, c, d := nums[0], nums[1], nums[2], nums[3]
        n, d := a * d + c * b, b * d 
        if n % d == 0 {
            return []int{n/d, 1}
        }
        m := abs(gcd(n, d))
        return []int{n/m, d/m}
    }
    nums, stack := make([]int, 0, 4), make([]rune, 0, len(expression))
    for i, ch := range expression {
        switch {
        case ch == '-' || ch == '+':
            if len(stack) > 0 {
                num := convertToInt(stack)
                nums = append(nums, num)
                stack = stack[len(stack):]
            }
            stack = append(stack, ch)
        case ch == '/':
            num := convertToInt(stack)
            nums = append(nums, num)
            stack = stack[len(stack):]
        default:
            stack = append(stack, ch)
        }
        if i == len(expression) - 1 {
            nums = append(nums, convertToInt(stack))
        }
        if len(nums) == 4 {
            nums = calc(nums)
        }
    }
    return strconv.Itoa(nums[0]) + "/" + strconv.Itoa(nums[1]) 
}

func main() {
    // Example 1:
    // Input: expression = "-1/2+1/2"
    // Output: "0/1"
    fmt.Println(fractionAddition("-1/2+1/2")) // "0/1"
    // Example 2:
    // Input: expression = "-1/2+1/2+1/3"
    // Output: "1/3"
    fmt.Println(fractionAddition("-1/2+1/2+1/3")) // 1/3
    // Example 3:
    // Input: expression = "1/3-1/2"
    // Output: "-1/6"
    fmt.Println(fractionAddition("1/3-1/2")) // "-1/6"
}