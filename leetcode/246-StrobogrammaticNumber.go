package main

// 246. Strobogrammatic Number
// Given a string num which represents an integer, return true if num is a strobogrammatic number.
// A strobogrammatic number is a number that looks the same when rotated 180 degrees (looked at upside down).

// Example 1:
// Input: num = "69"
// Output: true

// Example 2:
// Input: num = "88"
// Output: true

// Example 3:
// Input: num = "962"
// Output: false

// Constraints:
//     1 <= num.length <= 50
//     num consists of only digits.
//     num does not contain any leading zeros except for zero itself.

import "fmt"

// 双指针
func isStrobogrammatic(num string) bool {
    // 构建一个数字翻转后的数组，如果数字翻转后是非数字，则用字母代替
    // 6 => 9  9 => 6
    l := []byte{'0', '1', 'a', 'b', 'c', 'd', '9', 'e', '8', '6'}
    i, j := 0, len(num) - 1
    for  i <= j {
        if l[num[i]-'0'] != num[j] {
            return false
        }
        i++
        j--
    }
    return true
}

func main() {
    // Example 1:
    // Input: num = "69"
    // Output: true
    fmt.Println(isStrobogrammatic("69")) // true
    // Example 2:
    // Input: num = "88"
    // Output: true
    fmt.Println(isStrobogrammatic("88")) // true
    // Example 3:
    // Input: num = "962"
    // Output: false
    fmt.Println(isStrobogrammatic("962")) // false
}