package main

// 415. Add Strings
// Given two non-negative integers, num1 and num2 represented as string, return the sum of num1 and num2 as a string.
// You must solve the problem without using any built-in library for handling large integers (such as BigInteger). You must also not convert the inputs to integers directly.

// Example 1:
// Input: num1 = "11", num2 = "123"
// Output: "134"

// Example 2:
// Input: num1 = "456", num2 = "77"
// Output: "533"

// Example 3:
// Input: num1 = "0", num2 = "0"
// Output: "0"

// Constraints:
//     1 <= num1.length, num2.length <= 10^4
//     num1 and num2 consist of only digits.
//     num1 and num2 don't have any leading zeros except for the zero itself.
import "fmt"

func addStrings(num1 string, num2 string) string {
    max := func (x, y int) int { if x > y { return x; }; return y; }
    i, j := len(num1) - 1, len(num2) - 1
    longest := max(len(num1), len(num2))
    res, carry := make([]rune, longest + 1), rune(0)
    for longest >= 0 {
        n1, n2 := rune(0), rune(0)
        if i >= 0 {
            n1 = rune(num1[i] - '0')
            i--
        }
        if j >= 0 {
            n2 = rune(num2[j] - '0')
            j--
        }
        sum := carry + n1 + n2
        res[longest] = sum % 10 + '0'
        if sum > 9 { // 判断是否需要进位
            carry = 1
        } else {
            carry = 0
        }
        longest--
        if longest == 0 && carry > 0 {
            res[longest] = '1'
            return string(res)
        }
    }
    return string(res[1:])
}

func addStrings1(num1 string, num2 string) string {
    res, inx1, inx2, carry := []byte{}, len(num1) - 1, len(num2) - 1, 0
    for inx1 >= 0 || inx2 >= 0 || carry > 0 {
        x, y := 0, 0
        if inx1 >= 0 {
            x = int(num1[inx1] - '0')
        }
        if inx2 >= 0 {
            y = int(num2[inx2] - '0')
        }
        z := x + y + carry
        res = append(res, byte(z % 10 + int('0')))
        carry = z / 10
        inx1--
        inx2--
    }
    l, r := 0, len(res) - 1
    for l < r {
        res[l], res[r] = res[r], res[l]
        l++
        r--
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: num1 = "11", num2 = "123"
    // Output: "134"
    fmt.Println(addStrings("11","123")) // "134"
    // Example 2:
    // Input: num1 = "456", num2 = "77"
    // Output: "533"
    fmt.Println(addStrings("456","77")) // "533"
    // Example 3:
    // Input: num1 = "0", num2 = "0"
    // Output: "0"
    fmt.Println(addStrings("0","0")) // "0"

    fmt.Println(addStrings1("11","123")) // "134"
    fmt.Println(addStrings1("456","77")) // "533"
    fmt.Println(addStrings1("0","0")) // "0"
}