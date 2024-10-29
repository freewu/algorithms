package main

// 3216. Lexicographically Smallest String After a Swap
// Given a string s containing only digits, 
// return the lexicographically smallest string that can be obtained after swapping adjacent digits in s with the same parity at most once.

// Digits have the same parity if both are odd or both are even. 
// For example, 5 and 9, as well as 2 and 4, have the same parity, while 6 and 9 do not.

// Example 1:
// Input: s = "45320"
// Output: "43520"
// Explanation:
// s[1] == '5' and s[2] == '3' both have the same parity, and swapping them results in the lexicographically smallest string.

// Example 2:
// Input: s = "001"
// Output: "001"
// Explanation:
// There is no need to perform a swap because s is already the lexicographically smallest.

// Constraints:
//     2 <= s.length <= 100
//     s consists only of digits.

import "fmt"
import "reflect"

func getSmallestString(s string) string {
    arr := []byte(s)
    swap := reflect.Swapper(arr)
    for i := 0; i < len(arr) - 1; i++ {
        x, y := arr[i] - '0', arr[i+1] - '0'
        if x % 2 == y % 2 && x > y {
            swap(i, i + 1)
            break
        }
    }
    return string(arr)
}

func getSmallestString1(s string) string {
    arr, n := []byte(s), len(s)
    for i := 0; i < n - 1; i++ {
        if arr[i] % 2 == arr[i+1] % 2 && arr[i] > arr[i+1] { // 同奇偶性, 且 arr[i] > arr[i+1] 执行交换处理
            arr[i], arr[i+1] = arr[i+1], arr[i]
            break
        }
    }
    return string(arr)
}

func main() {
    // Example 1:
    // Input: s = "45320"
    // Output: "43520"
    // Explanation:
    // s[1] == '5' and s[2] == '3' both have the same parity, and swapping them results in the lexicographically smallest string.
    fmt.Println(getSmallestString("45320")) // "43520"
    // Example 2:
    // Input: s = "001"
    // Output: "001"
    // Explanation:
    // There is no need to perform a swap because s is already the lexicographically smallest.
    fmt.Println(getSmallestString("001")) // "001"

    fmt.Println(getSmallestString("8697")) // "6879"

    fmt.Println(getSmallestString1("45320")) // "43520"
    fmt.Println(getSmallestString1("001")) // "001"
    fmt.Println(getSmallestString1("8697")) // "6879"
}