package main

// 166. Fraction to Recurring Decimal
// Given two integers representing the numerator and denominator of a fraction, return the fraction in string format.
// If the fractional part is repeating, enclose the repeating part in parentheses.
// If multiple answers are possible, return any of them.
// It is guaranteed that the length of the answer string is less than 10^4 for all the given inputs.

// Example 1:
// Input: numerator = 1, denominator = 2
// Output: "0.5"

// Example 2:
// Input: numerator = 2, denominator = 1
// Output: "2"

// Example 3:
// Input: numerator = 4, denominator = 333
// Output: "0.(012)"
 
// Constraints:
//     -2^31 <= numerator, denominator <= 2^31 - 1
//     denominator != 0

import "fmt"
import "strings"
import "strconv"

func fractionToDecimal(numerator int, denominator int) string {
    res := ""
    if numerator * denominator < 0 { // 判断正负符号
        numerator = numerator * -1
        res = "-"
    }
    res = fmt.Sprintf("%s%d", res, numerator / denominator)
    numerator = numerator % denominator
    if numerator != 0 { // 判读是否有小数
        res += "."
    }
    mapNumberIndex := map[string]int{}
    numberFound := []int{}
    for numerator != 0 {
        k := numerator * 10 / denominator
        numerator = (numerator * 10) % denominator
        signature := fmt.Sprintf("%d%d", k, numerator)
        if repeatStarter, existed := mapNumberIndex[signature]; existed {
            for i, n := range numberFound {
                if i == repeatStarter {
                    res = fmt.Sprintf("%s(%d", res, n)
                    continue
                }
                res = fmt.Sprintf("%s%d", res, n)
            }
            return res + ")"
        }
        numberFound = append(numberFound, k)
        mapNumberIndex[signature] = len(numberFound) - 1
    }
    for _, n := range numberFound {
        res = fmt.Sprintf("%s%d", res, n)
    }
    return res
}

func fractionToDecimal1(numerator int, denominator int) string {
    const zero = '0'
    if numerator == 0 {
        return "0"
    }
    var sb strings.Builder
    negative := (numerator < 0) != (denominator < 0)
    if negative {
        sb.WriteByte('-')
    }
    if numerator < 0 {
        numerator = -numerator
    }
    if denominator < 0 {
        denominator = -denominator
    }
    whole := numerator / denominator
    sb.WriteString(strconv.Itoa(whole))
    numerator = numerator % denominator
    if numerator == 0 {
        return sb.String()
    }
    sb.WriteByte('.')
    seen := map[int]int{}
    for idx := sb.Len(); seen[numerator] == 0; idx++ {
        seen[numerator] = idx
        numerator = numerator % denominator * 10
        if numerator == 0 {
            return sb.String()
        }
        digit := numerator / denominator
        sb.WriteByte(byte(digit) + zero)
    }
    repeatPos := seen[numerator]
    out := sb.String()
    repeatingDigits := out[repeatPos:]
    decimalDigits := out[:repeatPos]
    for c := true; c; {
        decimalDigits, c = strings.CutSuffix(decimalDigits, repeatingDigits)
    }
    sb.Reset()
    sb.WriteByte('(')
    sb.WriteString(repeatingDigits)
    sb.WriteByte(')')
    return string(append([]byte(decimalDigits), sb.String()...))
}

func main() {
    // Example 1:
    // Input: numerator = 1, denominator = 2
    // Output: "0.5"
    fmt.Println(fractionToDecimal(1, 2)) // "0.5"
    // Example 2:
    // Input: numerator = 2, denominator = 1
    // Output: "2"
    fmt.Println(fractionToDecimal(2, 1)) // "2"
    // Example 3:
    // Input: numerator = 4, denominator = 333
    // Output: "0.(012)"
    fmt.Println(fractionToDecimal(4, 333)) // "0.(012)"

    fmt.Println(fractionToDecimal(1 << 31, -1 << 31)) // -1
    fmt.Println(fractionToDecimal(1 << 31, 1 << 31)) // 1
    fmt.Println(fractionToDecimal(-1 << 31, -1 << 31)) // 1
    fmt.Println(fractionToDecimal(-1 << 31, 1 << 31)) // -1

    fmt.Println(fractionToDecimal1(1, 2)) // "0.5"
    fmt.Println(fractionToDecimal1(2, 1)) // "2"
    fmt.Println(fractionToDecimal1(4, 333)) // "0.(012)"
    fmt.Println(fractionToDecimal1(1 << 31, -1 << 31)) // -1
    fmt.Println(fractionToDecimal1(1 << 31, 1 << 31)) // 1
    fmt.Println(fractionToDecimal1(-1 << 31, -1 << 31)) // 1
    fmt.Println(fractionToDecimal(-1 << 31, 1 << 31)) // -1
}