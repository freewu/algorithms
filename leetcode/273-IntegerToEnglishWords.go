package main

// 273. Integer to English Words
// Convert a non-negative integer num to its English words representation.

// Example 1:
// Input: num = 123
// Output: "One Hundred Twenty Three"

// Example 2:
// Input: num = 12345
// Output: "Twelve Thousand Three Hundred Forty Five"

// Example 3:
// Input: num = 1234567
// Output: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
 
// Constraints:
//     0 <= num <= 2^31 - 1

import "fmt"
import "strings"

func numberToWords(num int) string {
    // numbers are read from left (larger) to right (smaller)
    INTS := []int{int(1e9),int(1e6),1000,100,90,80,70,60,50,40,30,20,19,18,17,16,15,14,13,12,11,10,9,8,7,6,5,4,3,2,1,0}
    WORDS := []string{"Billion","Million","Thousand","Hundred","Ninety", "Eighty", "Seventy", "Sixty", "Fifty", "Forty", "Thirty", "Twenty", "Nineteen", "Eighteen", "Seventeen", "Sixteen", "Fifteen", "Fourteen", "Thirteen", "Twelve", "Eleven", "Ten", "Nine", "Eight", "Seven", "Six", "Five", "Four", "Three", "Two", "One", "Zero"}

    var dfs func (n int, i int, b *strings.Builder) 
    dfs = func (n int, i int, b *strings.Builder) {
        if n <= 20 { // because [0,20] can be translated one by one
            b.WriteString(WORDS[31-n])
            return
        }
        for n > 0 {
            curr := n/INTS[i] // Use integer division we check if a enumerated number is included in the given number
            if curr > 0 {
                if INTS[i] >= 100 { // Billion,Million,Thousand or Hundred are included
                    dfs(curr, i + 1, b)
                    b.WriteByte(' ')
                }
                b.WriteString(WORDS[i]) // Billion,Million,Thousand or Hundred are not included
                n %= INTS[i] 
                if n > 0 {
                    b.WriteByte(' ')
                }
            }
            i++
        }
    }
    var b strings.Builder
    dfs(num, 0, &b)
    return b.String()
}

func numberToWords1(num int) string {
    var (
        singles   = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
        teens     = []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
        tens      = []string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
        thousands = []string{"", "Thousand", "Million", "Billion"}
    )
    if num == 0 {
        return "Zero"
    }
    sb := &strings.Builder{}
    var dfs func(int)
    dfs = func(num int) {
        switch {
        case num == 0:
        case num < 10:
            sb.WriteString(singles[num])
            sb.WriteByte(' ')
        case num < 20:
            sb.WriteString(teens[num-10])
            sb.WriteByte(' ')
        case num < 100:
            sb.WriteString(tens[num/10])
            sb.WriteByte(' ')
            dfs(num % 10)
        default:
            sb.WriteString(singles[num/100])
            sb.WriteString(" Hundred ")
            dfs(num % 100)
        }
    }
    for i, unit := 3, int(1e9); i >= 0; i-- {
        if curNum := num / unit; curNum > 0 {
            num -= curNum * unit
            dfs(curNum)
            sb.WriteString(thousands[i])
            sb.WriteByte(' ')
        }
        unit /= 1000
    }
    return strings.TrimSpace(sb.String())
}


func numberToWords2(num int) string {
    if num == 0 {
        return "Zero"
    }
    elements := []string{}
    if num >= 1000000000 {
        elements = append(append(elements, numToString(num / 1000000000)...), "Billion")
    }
    if num >= 1000000 {
        millions := num % 1000000000 / 1000000
        if millions > 0 {
            elements = append(append(elements, numToString(millions)...), "Million")
        }
    }
    if num >= 1000 {
        thousands := num % 1000000 / 1000
        if thousands > 0 {
            elements = append(append(elements, numToString(thousands)...), "Thousand")
        }
    }
    elements = append(elements, numToString(num % 1000)...)
    return strings.Join(elements, " ")
}

func numToString(num int) []string {
    switch num {
    case 0:
        return nil
    case 1:
        return []string{"One"}
    case 2:
        return []string{"Two"}
    case 3:
        return []string{"Three"}
    case 4:
        return []string{"Four"}
    case 5:
        return []string{"Five"}
    case 6:
        return []string{"Six"}
    case 7:
        return []string{"Seven"}
    case 8:
        return []string{"Eight"}
    case 9:
        return []string{"Nine"}
    case 10:
        return []string{"Ten"}
    case 11:
        return []string{"Eleven"}
    case 12:
        return []string{"Twelve"}
    case 13:
        return []string{"Thirteen"}
    case 14:
        return []string{"Fourteen"}
    case 15:
        return []string{"Fifteen"}
    case 16:
        return []string{"Sixteen"}
    case 17:
        return []string{"Seventeen"}
    case 18:
        return []string{"Eighteen"}
    case 19:
        return []string{"Nineteen"}
    case 20:
        return []string{"Twenty"}
    case 30:
        return []string{"Thirty"}
    case 40:
        return []string{"Forty"}
    case 50:
        return []string{"Fifty"}
    case 60:
        return []string{"Sixty"}
    case 70:
        return []string{"Seventy"}
    case 80:
        return []string{"Eighty"}
    case 90:
        return []string{"Ninety"}
    }
    elements := []string{}
    hundreds := num / 100
    tens := num % 100 / 10
    ones := num % 10
    if hundreds > 0 {
        elements = append(append(elements, numToString(hundreds)...), "Hundred")
    }
    if tens >= 2 {
        elements = append(append(elements, numToString(tens * 10)...), numToString(ones)...)
    } else {
        elements = append(elements, numToString(num % 100)...)
    }
    return elements
}

func main() {
    // Example 1:
    // Input: num = 123
    // Output: "One Hundred Twenty Three"
    fmt.Println(numberToWords(123)) // "One Hundred Twenty Three"
    // Example 2:
    // Input: num = 12345
    // Output: "Twelve Thousand Three Hundred Forty Five"
    fmt.Println(numberToWords(12345)) // Twelve Thousand Three Hundred Forty Five
    // Example 3:
    // Input: num = 1234567
    // Output: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
    fmt.Println(numberToWords(1234567)) // One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven

    fmt.Println(numberToWords1(123)) // "One Hundred Twenty Three"
    fmt.Println(numberToWords1(12345)) // Twelve Thousand Three Hundred Forty Five
    fmt.Println(numberToWords1(1234567)) // One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven
    
    fmt.Println(numberToWords2(123)) // "One Hundred Twenty Three"
    fmt.Println(numberToWords2(12345)) // Twelve Thousand Three Hundred Forty Five
    fmt.Println(numberToWords2(1234567)) // One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven
}