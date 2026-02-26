package main

// 1404. Number of Steps to Reduce a Number in Binary Representation to One
// Given the binary representation of an integer as a string s, 
// return the number of steps to reduce it to 1 under the following rules:
//     If the current number is even, you have to divide it by 2.
//     If the current number is odd, you have to add 1 to it.

// It is guaranteed that you can always reach one for all test cases.

// Example 1:
// Input: s = "1101"
// Output: 6
// Explanation: "1101" corressponds to number 13 in their decimal representation.
// Step 1) 13 is odd, add 1 and obtain 14. 
// Step 2) 14 is even, divide by 2 and obtain 7.
// Step 3) 7 is odd, add 1 and obtain 8.
// Step 4) 8 is even, divide by 2 and obtain 4.  
// Step 5) 4 is even, divide by 2 and obtain 2. 
// Step 6) 2 is even, divide by 2 and obtain 1.  

// Example 2:
// Input: s = "10"
// Output: 1
// Explanation: "10" corressponds to number 2 in their decimal representation.
// Step 1) 2 is even, divide by 2 and obtain 1.  

// Example 3:
// Input: s = "1"
// Output: 0

// Constraints:
//     1 <= s.length <= 500
//     s consists of characters '0' or '1'
//     s[0] == '1'

import "fmt"

// 模拟
func numSteps(s string) int {
    res, arr := 0, []byte(s)
    for l := len(arr); l != 1; l = len(arr) {
        if arr[l-1] == '0' { // 如果当前数字为偶数，则将其除以 2  >> 1 
            arr = arr[:l-1]
        } else {
            for i := l - 1; i >= 0; i-- {
                if arr[i] == '0' {
                    arr[i] = '1'
                    break
                } else {
                    arr[i] = '0'
                }
            }
            if arr[0] == '0' {
                arr = append([]byte{'1'}, arr...)
            }
        }
        res++
    }
    return res
}

func main() {
// Example 1:
    // Input: s = "1101"
    // Output: 6
    // Explanation: "1101" corressponds to number 13 in their decimal representation.
    // Step 1) 13 is odd, add 1 and obtain 14. 
    // Step 2) 14 is even, divide by 2 and obtain 7.
    // Step 3) 7 is odd, add 1 and obtain 8.
    // Step 4) 8 is even, divide by 2 and obtain 4.  
    // Step 5) 4 is even, divide by 2 and obtain 2. 
    // Step 6) 2 is even, divide by 2 and obtain 1.  
    fmt.Println(numSteps("1101")) // 6
    // Example 2:
    // Input: s = "10"
    // Output: 1
    // Explanation: "10" corressponds to number 2 in their decimal representation.
    // Step 1) 2 is even, divide by 2 and obtain 1.  
    fmt.Println(numSteps("10")) // 1
    // Example 3:
    // Input: s = "1"
    // Output: 0
    fmt.Println(numSteps("1")) // 0

    fmt.Println(numSteps("1111111111")) // 11
    fmt.Println(numSteps("0000000000")) // 9
    fmt.Println(numSteps("0000011111")) // 17
    fmt.Println(numSteps("1111100000")) // 11
    fmt.Println(numSteps("1010101010")) // 15
    fmt.Println(numSteps("0101010101")) // 17
}