package main

// 3765. Complete Prime Number
// You are given an integer num.

// A number num is called a Complete Prime Number if every prefix and every suffix of num is prime.

// Return true if num is a Complete Prime Number, otherwise return false.

// Note:
//     1. A prefix of a number is formed by the first k digits of the number.
//     2. A suffix of a number is formed by the last k digits of the number.
//     3. A prime number is a natural number greater than 1 with only two factors, 1 and itself.
//     4. Single-digit numbers are considered Complete Prime Numbers only if they are prime.

// Example 1:
// Input: num = 23
// Output: true
// Explanation:
// ​​​​​​​Prefixes of num = 23 are 2 and 23, both are prime.
// Suffixes of num = 23 are 3 and 23, both are prime.
// All prefixes and suffixes are prime, so 23 is a Complete Prime Number and the answer is true.

// Example 2:
// Input: num = 39
// Output: false
// Explanation:
// Prefixes of num = 39 are 3 and 39. 3 is prime, but 39 is not prime.
// Suffixes of num = 39 are 9 and 39. Both 9 and 39 are not prime.
// At least one prefix or suffix is not prime, so 39 is not a Complete Prime Number and the answer is false.

// Example 3:
// Input: num = 7
// Output: true
// Explanation:
// 7 is prime, so all its prefixes and suffixes are prime and the answer is true.

// Constraints:
//     1 <= num <= 10^9

import "fmt"
import "strconv"
import "math"

func completePrime(num int) bool {
    divisor, suffix, prefix := 1, num, num
    for num >= 10 {
        num /= 10
        divisor *= 10
    }
    // Time complexity: O(sqrt(num))
    isPrime := func(num int) bool {
        if num <= 1 { return false }
        if num == 2 { return true } // Check 2 separately
        if num % 2 == 0 { return false }
        for i := 3; i * i <= num; i += 2 { // Check for odd divisors up to sqrt(num)
            if num % i == 0 {
                return false
            }
        }
        return true
    }
    for prefix > 0 {
        if !isPrime(prefix) {
            return false
        }
        prefix %= divisor
        divisor /= 10
    }
    for suffix > 0 {
        if !isPrime(suffix) {
            return false
        }
        suffix /= 10
    }
    return true
}

func completePrime1(num int) bool {
    str := strconv.Itoa(num)
    n := len(str)
    isPrime := func(n int) bool {
        if n < 2 { return false } // 0和1不是质数
        if n == 2 || n == 3 { return true } // 2和3是质数
        if n % 2 == 0 || n % 3 == 0 {  return false } // 排除偶数（除2外）和能被3整除的数
        if n % 6 != 1 && n % 6 != 5 { return false } // 所有大于5的质数必位于6x±1两侧（如5,7,11,13）
        for i := (5); i <= int(math.Sqrt(float64(n))) ; i += 6 { // 只需检查到 √n 范围内的因子, 步长为6，检查i和i+2
            if n % i == 0 || n % (i + 2) == 0 {
                return false
            }
        }
        return true
    }
    for i := 0; i < n; i++ { // prefix
        prefix, _ := strconv.Atoi(str[i:])
        if !isPrime(prefix) {
            return false
        }
    }
    for i := n; i > 0; i-- { // subfix
        subfix, _ := strconv.Atoi(str[:i])
        // fmt.Printf("subfix[:%d]=%d\n", i, subfix)
        if !isPrime(subfix) {
            return false
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: num = 23
    // Output: true
    // Explanation:
    // ​​​​​​​Prefixes of num = 23 are 2 and 23, both are prime.
    // Suffixes of num = 23 are 3 and 23, both are prime.
    // All prefixes and suffixes are prime, so 23 is a Complete Prime Number and the answer is true.
    fmt.Println(completePrime(23)) // true
    // Example 2:
    // Input: num = 39
    // Output: false
    // Explanation:
    // Prefixes of num = 39 are 3 and 39. 3 is prime, but 39 is not prime.
    // Suffixes of num = 39 are 9 and 39. Both 9 and 39 are not prime.
    // At least one prefix or suffix is not prime, so 39 is not a Complete Prime Number and the answer is false.
    fmt.Println(completePrime(39)) // false
    // Example 3:
    // Input: num = 7
    // Output: true
    // Explanation:
    // 7 is prime, so all its prefixes and suffixes are prime and the answer is true.
    fmt.Println(completePrime(7)) // true

    fmt.Println(completePrime(1)) // false
    fmt.Println(completePrime(13)) // false
    fmt.Println(completePrime(1024)) // false
    fmt.Println(completePrime(1_000_000_000)) // false
    fmt.Println(completePrime(1_000_000_007)) // false

    fmt.Println(completePrime1(23)) // true
    fmt.Println(completePrime1(39)) // false
    fmt.Println(completePrime1(7)) // true
    fmt.Println(completePrime1(1)) // false
    fmt.Println(completePrime1(13)) // false
    fmt.Println(completePrime1(1024)) // false
    fmt.Println(completePrime1(1_000_000_000)) // false
    fmt.Println(completePrime1(1_000_000_007)) // false
}
