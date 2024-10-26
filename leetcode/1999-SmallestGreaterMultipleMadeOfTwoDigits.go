package main

// 1999. Smallest Greater Multiple Made of Two Digits
// Given three integers, k, digit1, and digit2, you want to find the smallest integer that is:
//     Larger than k,
//     A multiple of k, and
//     Comprised of only the digits digit1 and/or digit2.

// Return the smallest such integer. 
// If no such integer exists or the integer exceeds the limit of a signed 32-bit integer (2^31 - 1), return -1.

// Example 1:
// Input: k = 2, digit1 = 0, digit2 = 2
// Output: 20
// Explanation:
// 20 is the first integer larger than 2, a multiple of 2, and comprised of only the digits 0 and/or 2.

// Example 2:
// Input: k = 3, digit1 = 4, digit2 = 2
// Output: 24
// Explanation:
// 24 is the first integer larger than 3, a multiple of 3, and comprised of only the digits 4 and/or 2.

// Example 3:
// Input: k = 2, digit1 = 0, digit2 = 0
// Output: -1
// Explanation:
// No integer meets the requirements so return -1.

// Constraints:
//     1 <= k <= 1000
//     0 <= digit1 <= 9
//     0 <= digit2 <= 9

import "fmt"
import "math"
import "strconv"

func findInteger(k int, digit1 int, digit2 int) int {
    var dfs func(k, digit1, digit2, mx, res int, nowStr string) int
    dfs = func(k, digit1, digit2, mx, res int, nowStr string) int {
        n, _ := strconv.Atoi(nowStr)
        if n > mx || (res != -1 && n >= res) || len(nowStr) > len(strconv.Itoa(mx)) {
            return res
        }
        if n > k && n % k == 0 {
            if res == -1 || n < res {
                res = n
            }
        } else {
            res = dfs(k, digit1, digit2, mx, res, nowStr+strconv.Itoa(digit1))
            res = dfs(k, digit1, digit2, mx, res, nowStr+strconv.Itoa(digit2))
        }
        return res
    }
    return dfs(k, digit1, digit2, int(math.Pow(2, 31)) - 1, -1, "")
}


func main() {
    // Example 1:
    // Input: k = 2, digit1 = 0, digit2 = 2
    // Output: 20
    // Explanation:
    // 20 is the first integer larger than 2, a multiple of 2, and comprised of only the digits 0 and/or 2.
    fmt.Println(findInteger(2,0,2)) // 20
    // Example 2:
    // Input: k = 3, digit1 = 4, digit2 = 2
    // Output: 24
    // Explanation:
    // 24 is the first integer larger than 3, a multiple of 3, and comprised of only the digits 4 and/or 2.
    fmt.Println(findInteger(3,4,2)) // 24
    // Example 3:
    // Input: k = 2, digit1 = 0, digit2 = 0
    // Output: -1
    // Explanation:
    // No integer meets the requirements so return -1.
    fmt.Println(findInteger(2,0,0)) // -1

    fmt.Println(findInteger(1000, 0, 0)) // -1
    fmt.Println(findInteger(1000, 9, 9)) // -1
    fmt.Println(findInteger(1, 0, 0)) // -1
    fmt.Println(findInteger(1, 9, 9)) // 9
}