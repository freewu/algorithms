package main

// 1323. Maximum 69 Number
// You are given a positive integer num consisting only of digits 6 and 9.

// Return the maximum number you can get by changing at most one digit (6 becomes 9, and 9 becomes 6).

// Example 1:
// Input: num = 9669
// Output: 9969
// Explanation: 
// Changing the first digit results in 6669.
// Changing the second digit results in 9969.
// Changing the third digit results in 9699.
// Changing the fourth digit results in 9666.
// The maximum number is 9969.

// Example 2:
// Input: num = 9996
// Output: 9999
// Explanation: Changing the last digit 6 to 9 results in the maximum number.

// Example 3:
// Input: num = 9999
// Output: 9999
// Explanation: It is better not to apply any change.

// Constraints:
//     1 <= num <= 10^4
//     num consists of only 6 and 9 digits.

import "fmt"
import "strconv" 

func maximum69Number(num int) int {
    flag, arr := true, []byte(strconv.Itoa(num))
    for i := 0; i < len(arr); i++ {
        if arr[i] == '6' { // 改变第一个遇到的 6 => 9
            arr[i] = '9'
            flag = false
            break
        }
    }
    if flag { return num } // 没有改变
    res, _ := strconv.Atoi(string(arr))
    return res
}

func maximum69Number1(num int) int {
    snum := []byte(strconv.Itoa(num))
    res := num
    pairs := map[byte]byte{ '6': '9', '9': '6', }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, v := range snum {
        snum := []byte(strconv.Itoa(num))
        value, _ := pairs[v]
        snum[i] = value
        inum, _ := strconv.Atoi(string(snum))
        res = max(res, inum)
    }
    return res
}

func main() {
    // Example 1:
    // Input: num = 9669
    // Output: 9969
    // Explanation: 
    // Changing the first digit results in 6669.
    // Changing the second digit results in 9969.
    // Changing the third digit results in 9699.
    // Changing the fourth digit results in 9666.
    // The maximum number is 9969.
    fmt.Println(maximum69Number(9669)) // 9969
    // Example 2:
    // Input: num = 9996
    // Output: 9999
    // Explanation: Changing the last digit 6 to 9 results in the maximum number.
    fmt.Println(maximum69Number(9996)) // 9999
    // Example 3:
    // Input: num = 9999
    // Output: 9999
    // Explanation: It is better not to apply any change.
    fmt.Println(maximum69Number(9999)) // 9999

    fmt.Println(maximum69Number1(9669)) // 9969
    fmt.Println(maximum69Number1(9996)) // 9999
    fmt.Println(maximum69Number1(9999)) // 9999
}