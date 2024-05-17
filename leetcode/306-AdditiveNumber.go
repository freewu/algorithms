package main

// 306. Additive Number
// An additive number is a string whose digits can form an additive sequence.
// A valid additive sequence should contain at least three numbers.
// Except for the first two numbers, each subsequent number in the sequence must be the sum of the preceding two.
// Given a string containing only digits, return true if it is an additive number or false otherwise.
// Note: Numbers in the additive sequence cannot have leading zeros, so sequence 1, 2, 03 or 1, 02, 3 is invalid.

// Example 1:
// Input: "112358"
// Output: true
// Explanation: 
// The digits can form an additive sequence: 1, 1, 2, 3, 5, 8. 
// 1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8

// Example 2:
// Input: "199100199"
// Output: true
// Explanation: 
// The additive sequence is: 1, 99, 100, 199. 
// 1 + 99 = 100, 99 + 100 = 199
 
// Constraints:
//     1 <= num.length <= 35
//     num consists only of digits.
    
// Follow up: How would you handle overflow for very large input integers?

import "fmt"
import "strconv"

func isAdditiveNumber(num string) bool {
    n := len(num)
    var isValid func (num string, i, j, k int) bool
    isValid = func (num string, i, j, k int) bool {
        if k == len(num) {
            return true
        }
        num1, _ := strconv.Atoi(num[i:j])
        num2, _ := strconv.Atoi(num[j:k])
        sum := strconv.Itoa(num1 + num2)
        if k + len(sum) > len(num) || num[k:k + len(sum)] != sum {
            return false
        }
        return isValid(num, j, k, k + len(sum))
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i <= n / 2; i++ {
        if num[0] == '0' && i > 1 {
            break
        }
        for j := i + 1; max(i, j - i) <= n - j; j++ {
            if num[i] == '0' && j-i > 1 {
                break
            }
            if isValid(num, 0, i, j) {
                return true
            }
        }
    }
    return false
}

func main() {
    // Example 1:
    // Input: "112358"
    // Output: true
    // Explanation: 
    // The digits can form an additive sequence: 1, 1, 2, 3, 5, 8. 
    // 1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8
    fmt.Println(isAdditiveNumber("112358")) // true
    // Example 2:
    // Input: "199100199"
    // Output: true
    // Explanation: 
    // The additive sequence is: 1, 99, 100, 199. 
    // 1 + 99 = 100, 99 + 100 = 199
    fmt.Println(isAdditiveNumber("199100199")) // true
}