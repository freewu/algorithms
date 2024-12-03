package main

// 2847. Smallest Number With Given Digit Product
// Given a positive integer n, return a string representing the smallest positive integer 
// such that the product of its digits is equal to n, or "-1" if no such number exists.

// Example 1:
// Input: n = 105
// Output: "357"
// Explanation: 
// 3 * 5 * 7 = 105. 
// It can be shown that 357 is the smallest number with a product of digits equal to 105. 
// So the answer would be "357".

// Example 2:
// Input: n = 7
// Output: "7"
// Explanation: 
// Since 7 has only one digit, its product of digits would be 7. 
// We will show that 7 is the smallest number with a product of digits equal to 7. 
// Since the product of numbers 1 to 6 is 1 to 6 respectively, so "7" would be the answer.

// Example 3:
// Input: n = 44
// Output: "-1"
// Explanation: 
// It can be shown that there is no number such that its product of digits is equal to 44. 
// So the answer would be "-1".

// Constraints:
//     1 <= n <= 10^18

import "fmt"

func smallestNumber(n int64) string {
    arr, count := []byte{}, [10]int{}
    for i := 9; i > 1; i-- {
        for n % int64(i) == 0 {
            count[i]++
            n /= int64(i)
        }
    }
    if n != 1 { return "-1" }
    for i := 2; i < 10; i++ {
        for j := 0; j < count[i]; j++ {
            arr = append(arr, byte(i) + '0')
        }
    }
    res := string(arr)
    if len(res) > 0 { return res }
    return "1"
}

func smallestNumber1(n int64) string {
    if n == 1 { return "1" }
    res := []byte{}
    for n != 1 {
        found := false
        for i := 9 ; i >= 2; i-- {
            if n % int64(i) == 0 {
                res = append(res, byte('0' + i))
                n /= int64(i)
                found = true
                break
            }
        }
        if !found {  return "-1" }
    }
    reverse := func(arr []byte) {
        for i, j := 0, len(arr) - 1; i < j; i, j = i + 1, j - 1 {
            arr[i], arr[j] = arr[j], arr[i]
        }
    }
    reverse(res)
    return string(res)
}

func main() {
    // Example 1:
    // Input: n = 105
    // Output: "357"
    // Explanation: 
    // 3 * 5 * 7 = 105. 
    // It can be shown that 357 is the smallest number with a product of digits equal to 105. 
    // So the answer would be "357".
    fmt.Println(smallestNumber(105)) // "357"
    // Example 2:
    // Input: n = 7
    // Output: "7"
    // Explanation: 
    // Since 7 has only one digit, its product of digits would be 7. 
    // We will show that 7 is the smallest number with a product of digits equal to 7. 
    // Since the product of numbers 1 to 6 is 1 to 6 respectively, so "7" would be the answer.
    fmt.Println(smallestNumber(7)) // "7"
    // Example 3:
    // Input: n = 44
    // Output: "-1"
    // Explanation: 
    // It can be shown that there is no number such that its product of digits is equal to 44. 
    // So the answer would be "-1".
    fmt.Println(smallestNumber(44)) // "-1"

    fmt.Println(smallestNumber(1)) // 1
    fmt.Println(smallestNumber(1e18)) // 555555555555555555888888

    fmt.Println(smallestNumber1(105)) // "357"
    fmt.Println(smallestNumber1(7)) // "7"
    fmt.Println(smallestNumber1(44)) // "-1"
    fmt.Println(smallestNumber1(1)) // 1
    fmt.Println(smallestNumber1(1e18)) // 555555555555555555888888
}