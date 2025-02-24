package main

// 面试题 05.02. Binary Number to String LCCI
// Given a real number between 0 and 1 (e.g., 0.72) that is passed in as a double, print the binary representation. 
// If the number cannot be represented accurately in binary with at most 32 characters, print "ERROR".

// Example1:
// Input: 0.625
// Output: "0.101"

// Example2:
// Input: 0.1
// Output: "ERROR"
// Note: 0.1 cannot be represented accurately in binary.

// Note:
//     This two charaters "0." should be counted into 32 characters.
//     The number of decimal places for num is at most 6 digits

import "fmt"
import "strings"

func printBin(num float64) string {
    var sb strings.Builder
    sb.WriteString("0.")
    var r float64
    for num > 0 && num < 1 {
        r = num * 2
        if r >= 1 {
            sb.WriteString("1")
            num = r - 1
        } else {
            sb.WriteString("0")
            num = r
        }
        if sb.Len() > 32 {
            return "ERROR"
        }
    }
    return sb.String()
}

func printBin1(num float64) string {
    res, r := []byte{'0', '.'}, float64(0)
    for num > 0 && num < 1 {
        r = num * 2
        if r >= 1 {
            res, num = append(res, '1'), r - 1
        } else {
            res, num = append(res, '0'), r
        }
        if len(res) > 32 {
            return "ERROR"
        }
    }
    return string(res)
}

func main() {
    // Example1:
    // Input: 0.625
    // Output: "0.101"
    fmt.Println(printBin(0.625)) // "0.101"
    // Example2:
    // Input: 0.1
    // Output: "ERROR"
    // Note: 0.1 cannot be represented accurately in binary.
    fmt.Println(printBin(0.1)) // "ERROR"

    fmt.Println(printBin1(0.625)) // "0.101"
    fmt.Println(printBin1(0.1)) // "ERROR"
}