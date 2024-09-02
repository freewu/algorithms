package main

// 972. Equal Rational Numbers
// Given two strings s and t, each of which represents a non-negative rational number, 
// return true if and only if they represent the same number. 
// The strings may use parentheses to denote the repeating part of the rational number.

// A rational number can be represented using up to three parts: 
//     <IntegerPart>, <NonRepeatingPart>, and a <RepeatingPart>. 
// The number will be represented in one of the following three ways:
//     <IntegerPart>
//         For example, 12, 0, and 123.
//     <IntegerPart><.><NonRepeatingPart>
//         For example, 0.5, 1., 2.12, and 123.0001.
//     <IntegerPart><.><NonRepeatingPart><(><RepeatingPart><)>
//         For example, 0.1(6), 1.(9), 123.00(1212).

// The repeating portion of a decimal expansion is conventionally denoted within a pair of round brackets. 
// For example:
//     1/6 = 0.16666666... = 0.1(6) = 0.1666(6) = 0.166(66).
 
// Example 1:
// Input: s = "0.(52)", t = "0.5(25)"
// Output: true
// Explanation: Because "0.(52)" represents 0.52525252..., and "0.5(25)" represents 0.52525252525..... , the strings represent the same number.

// Example 2:
// Input: s = "0.1666(6)", t = "0.166(66)"
// Output: true

// Example 3:
// Input: s = "0.9(9)", t = "1."
// Output: true
// Explanation: "0.9(9)" represents 0.999999999... repeated forever, which equals 1.  [See this link for an explanation.]
// "1." represents the number 1, which is formed correctly: (IntegerPart) = "1" and (NonRepeatingPart) = "".

// Constraints:
//     Each part consists only of digits.
//     The <IntegerPart> does not have leading zeros (except for the zero itself).
//     1 <= <IntegerPart>.length <= 4
//     0 <= <NonRepeatingPart>.length <= 4
//     1 <= <RepeatingPart>.length <= 4

import "fmt"
import "strings"
import "unicode"

func isRationalEqual(s string, t string) bool {
    iter1, iter2 := getIter(s), getIter(t)
    for i := 0; i < 24; i++ {
        a,b := iter1.next(), iter2.next()
        if a == b { continue }
        if b > a {
            a, b = b, a
            iter1, iter2 = iter2, iter1
        }
        if a - b != 1 {
            return false
        }
        for i := 0; i < 24; i++ {
            if iter1.next() != byte('0') || iter2.next() != byte('9') {
                return false
            }
        }
        return true
    }
    return true
}

type iterator struct {
    nonRepeating string
    repeating string
    i int
}

func (iter *iterator) next() byte {
    iter.i++
    if iter.i < len(iter.nonRepeating) {
        if iter.nonRepeating[iter.i] == byte('.') {
            iter.i++
        }
    }
    if iter.i < len(iter.nonRepeating) {
        return iter.nonRepeating[iter.i]
    }
    if len(iter.repeating) == 0 {
        return byte('0')
    }
    index := (iter.i - len(iter.nonRepeating)) % len(iter.repeating)
    return iter.repeating[index]
}

func getIter(s string) *iterator {
    repeatSplit := strings.Split(s, "(")
    if len(repeatSplit) == 1 {
        return &iterator {
            s,"",-1,
        }
    }
    repeat := repeatSplit[1][:len(repeatSplit[1])-1]
    return &iterator {
        repeatSplit[0],
        repeat,
        -1,
    }
}

func isRationalEqual1(s string, t string) bool {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    gcd := func (x, y int) int { x,y = abs(x), abs(y); for y != 0 { x, y = y, x % y; }; return x; }
    decimalToFraction := func(num string) []int {
        fraction, length, index := make([]int, 2), len(num), 0
        integerPart, nonRepeatingPart, repeatingPart, nonRepeatingUnit, repeatingUnit := 0, 0, 0, 1, 0
        for index < length && unicode.IsNumber(rune(num[index])) {
            digit := int(num[index] - '0')
            integerPart = integerPart * 10 + digit
            index++
        }
        if index < length && num[index] == '.' {
            index++
        }
        for index < length && unicode.IsNumber(rune(num[index])) {
            digit := int(num[index] - '0')
            nonRepeatingPart = nonRepeatingPart * 10 + digit
            nonRepeatingUnit *= 10
            index++
        }
        if index < length && num[index] == '(' {
            index++
        }
        for index < length && unicode.IsNumber(rune(num[index])) {
            digit := int(num[index] - '0')
            repeatingPart = repeatingPart * 10 + digit
            repeatingUnit = repeatingUnit * 10 + 9
            index++
        }
        if repeatingPart == 0 {
            fraction[0] = nonRepeatingPart
            fraction[1] = nonRepeatingUnit
        } else {
            fraction[0] = nonRepeatingPart * repeatingUnit + repeatingPart
            fraction[1] = repeatingUnit * nonRepeatingUnit
        }
        fraction[0] += integerPart * fraction[1]
        divisor := gcd(fraction[0], fraction[1])
        fraction[0] /= divisor
        fraction[1] /= divisor
        return fraction
    }
    fraction1, fraction2 := decimalToFraction(s), decimalToFraction(t)
    return fraction1[0] == fraction2[0] && fraction1[1] == fraction2[1]
}

func main() {
    // Example 1:
    // Input: s = "0.(52)", t = "0.5(25)"
    // Output: true
    // Explanation: Because "0.(52)" represents 0.52525252..., and "0.5(25)" represents 0.52525252525..... , the strings represent the same number.
    fmt.Println(isRationalEqual("0.(52)", "0.5(25)")) // true
    // Example 2:
    // Input: s = "0.1666(6)", t = "0.166(66)"
    // Output: true
    fmt.Println(isRationalEqual("0.1666(6)", "0.166(66)")) // true
    // Example 3:
    // Input: s = "0.9(9)", t = "1."
    // Output: true
    // Explanation: "0.9(9)" represents 0.999999999... repeated forever, which equals 1.  [See this link for an explanation.]
    // "1." represents the number 1, which is formed correctly: (IntegerPart) = "1" and (NonRepeatingPart) = "".
    fmt.Println(isRationalEqual("0.9(9)", "1.")) // true

    fmt.Println(isRationalEqual1("0.(52)", "0.5(25)")) // true
    fmt.Println(isRationalEqual1("0.1666(6)", "0.166(66)")) // true
    fmt.Println(isRationalEqual1("0.9(9)", "1.")) // true
}