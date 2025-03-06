package main

// 面试题 16.08. English Int LCCI
// Given any integer, print an English phrase that describes the integer 
// (e.g., "One Thousand Two Hundred Thirty Four").

// Example 1:
// Input: 123
// Output: "One Hundred Twenty Three"

// Example 2:
// Input: 12345
// Output: "Twelve Thousand Three Hundred Forty Five"

// Example 3:
// Input: 1234567
// Output: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"

// Example 4:
// Input: 1234567891
// Output: "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One"

import "fmt"
import "strings"

func numberToWords(num int) string {
    if num == 0 { return "Zero" }
    var (
        singles   = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
        teens     = []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
        tens      = []string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
        thousands = []string{"", "Thousand", "Million", "Billion"}
    )
    sb := &strings.Builder{}
    var recursion func(int)
    recursion = func(num int) {
        switch {
        case num == 0:
        case num < 10: // 1 - 9
            sb.WriteString(singles[num])
            sb.WriteByte(' ')
        case num < 20: // 10 - 19
            sb.WriteString(teens[num - 10])
            sb.WriteByte(' ')
        case num < 100: // 20 - 90
            sb.WriteString(tens[num / 10])
            sb.WriteByte(' ')
            recursion(num % 10)
        default:
            sb.WriteString(singles[num / 100])
            sb.WriteString(" Hundred ")
            recursion(num % 100)
        }
    }
    for i, unit := 3, int(1e9); i >= 0; i-- {
        if cur:= num / unit; cur > 0 {
            num -= cur * unit
            recursion(cur)
            sb.WriteString(thousands[i]) // 逢 3 
            sb.WriteByte(' ')
        }
        unit /= 1000
    }
    return strings.TrimSpace(sb.String())
}

func main() {
    // Example 1:
    // Input: 123
    // Output: "One Hundred Twenty Three"
    fmt.Println(numberToWords(123)) // "One Hundred Twenty Three"
    // Example 2:
    // Input: 12345
    // Output: "Twelve Thousand Three Hundred Forty Five"
    fmt.Println(numberToWords(12345)) // "Twelve Thousand Three Hundred Forty Five"
    // Example 3:
    // Input: 1234567
    // Output: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
    fmt.Println(numberToWords(1234567)) // "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
    // Example 4:
    // Input: 1234567891
    // Output: "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One"
    fmt.Println(numberToWords(1234567891)) // "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One"

    fmt.Println(numberToWords(123456789)) // One Hundred Twenty Three Million Four Hundred Fifty Six Thousand Seven Hundred Eighty Nine
    fmt.Println(numberToWords(987654321)) // Nine Hundred Eighty Seven Million Six Hundred Fifty Four Thousand Three Hundred Twenty One
}