package main

// 3348. Smallest Divisible Digit Product II
// You are given a string num which represents a positive integer, and an integer t.

// A number is called zero-free if none of its digits are 0.

// Return a string representing the smallest zero-free number greater than 
// or equal to num such that the product of its digits is divisible by t. 
// If no such number exists, return "-1".

// Example 1:
// Input: num = "1234", t = 256
// Output: "1488"
// Explanation:
// The smallest zero-free number that is greater than 1234 and has the product of its digits divisible by 256 is 1488, with the product of its digits equal to 256.

// Example 2:
// Input: num = "12355", t = 50
// Output: "12355"
// Explanation:
// 12355 is already zero-free and has the product of its digits divisible by 50, with the product of its digits equal to 150.

// Example 3:
// Input: num = "11111", t = 26
// Output: "-1"
// Explanation:
// No number greater than 11111 has the product of its digits divisible by 26.

// Constraints:
//     2 <= num.length <= 2 * 10^5
//     num consists only of digits in the range ['0', '9'].
//     num does not contain leading zeros.
//     1 <= t <= 10^14

import "fmt"
import "slices"

// 从右到左枚举
func smallestNumber(num string, t int64) string {
    tmp := int(t)
    for i := 9; i > 1; i-- {
        for tmp % i == 0 { tmp /= i }
    }
    if tmp > 1 { return "-1" } // t 包含大于 7 的质因子
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    n := len(num)
    leftT := make([]int, n + 1)
    leftT[0] = int(t)
    i0 := n - 1
    for i, c := range num {
        if c == '0' {
            i0 = i
            break
        }
        leftT[i+1] = leftT[i] / gcd(leftT[i], int(c-'0'))
    }
    if leftT[n] == 1 { return num } // num 的数位之积是 t 的倍数
    s := []byte(num) // 假设答案和 num 一样长
    for i := i0; i >= 0; i-- {
        for s[i]++; s[i] <= '9'; s[i]++ {
            tt := leftT[i] / gcd(leftT[i], int(s[i]-'0'))
            k := 9
            for j := n - 1; j > i; j-- {
                for tt % k > 0 { k-- }
                tt /= k
                s[j] = '0' + byte(k)
            }
            if tt == 1 {
                return string(s)
            }
        }
    }
    res := []byte{} // 答案一定比 num 长
    for i := int64(9); i > 1; i-- {
        for t%i == 0 {
            res = append(res, '0' + byte(i))
            t /= i
        }
    }
    for len(res) <= n {
        res = append(res, '1')
    }
    slices.Reverse(res)
    return string(res)
}

func main() {
    // Example 1:
    // Input: num = "1234", t = 256
    // Output: "1488"
    // Explanation:
    // The smallest zero-free number that is greater than 1234 and has the product of its digits divisible by 256 is 1488, with the product of its digits equal to 256.
    fmt.Println(smallestNumber("1234", 256)) // "1488"
    // Example 2:
    // Input: num = "12355", t = 50
    // Output: "12355"
    // Explanation:
    // 12355 is already zero-free and has the product of its digits divisible by 50, with the product of its digits equal to 150.
    fmt.Println(smallestNumber("12355", 50)) // "12355"
    // Example 3:
    // Input: num = "11111", t = 26
    // Output: "-1"
    // Explanation:
    // No number greater than 11111 has the product of its digits divisible by 26.
    fmt.Println(smallestNumber("11111", 26)) // "-1"
}