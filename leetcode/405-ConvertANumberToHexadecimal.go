package main


// 405. Convert a Number to Hexadecimal
// Given an integer num, return a string representing its hexadecimal representation. 
// For negative integers, two’s complement method is used.

// All the letters in the answer string should be lowercase characters, 
// and there should not be any leading zeros in the answer except for the zero itself.

// Note: You are not allowed to use any built-in library method to directly solve this problem.

// Example 1:
// Input: num = 26
// Output: "1a"

// Example 2:
// Input: num = -1
// Output: "ffffffff"
 
// Constraints:
//     -2^31 <= num <= 2^31 - 1

import "fmt"

// with sys lib
func toHex(num int) string {
    return fmt.Sprintf("%x", uint32(num))
}

func toHex1(num int) string {
    if num == 0 {
        return "0"
    }
    res, n, ref := "", uint32(num), "0123456789abcdef"  
    for n != 0 {
        res = string(ref[n & 0xF]) + res
        n >>= 4
    }
    return res
}

func toHex2(num int) string {
    if num == 0 {
        return "0"
    }
    if num < 0 { // 处理负数的情况
        num += 1 << 32
    }
    mp := map[int]string{
        0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9",
        10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f",
    }
    res, bitArr := "", []string{}
    for num > 0 {
        bitArr = append(bitArr, mp[num%16])
        num /= 16
    }
    for i := len(bitArr) - 1; i >= 0; i-- {
        res += bitArr[i]
    }
    return res
}

func toHex3(num int) string {
    hexes := "0123456789abcdef"
    if num == 0 {
        return "0"
    }
    if num < 0 { // 处理负数
        num += 1 << 32
    }
    res := make([]byte, 0, 8)
    for num > 0 {
        res = append(res, hexes[num % 16])
        num /= 16
    }
    for i, n := 0, len(res); i < n/2; i++ { // 翻转
        res[i], res[n-i-1] = res[n-i-1], res[i]
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: num = 26
    // Output: "1a"
    fmt.Println(toHex(26)) // la
    // Example 2:
    // Input: num = -1
    // Output: "ffffffff"
    fmt.Println(toHex(-1)) // ffffffff

    fmt.Println(toHex1(26)) // la
    fmt.Println(toHex1(-1)) // ffffffff

    fmt.Println(toHex2(26)) // la
    fmt.Println(toHex2(-1)) // ffffffff

    fmt.Println(toHex3(26)) // la
    fmt.Println(toHex3(-1)) // ffffffff
}