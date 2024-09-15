package main

// 1108. Defanging an IP Address
// Given a valid (IPv4) IP address, return a defanged version of that IP address.

// A defanged IP address replaces every period "." with "[.]".

// Example 1:
// Input: address = "1.1.1.1"
// Output: "1[.]1[.]1[.]1"

// Example 2:
// Input: address = "255.100.50.0"
// Output: "255[.]100[.]50[.]0"
 

// Constraints:
//     The given address is a valid IPv4 address.

import "fmt"

func defangIPaddr(address string) string {
    res := []byte{}
    for i := 0; i < len(address); i++ {
        if address[i] == '.' {
            res = append(res, '[')
            res = append(res, '.')
            res = append(res, ']')
        } else {
            res = append(res, address[i])
        }
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: address = "1.1.1.1"
    // Output: "1[.]1[.]1[.]1"
    fmt.Println(defangIPaddr("1.1.1.1")) // "1[.]1[.]1[.]1"
    // Example 2:
    // Input: address = "255.100.50.0"
    // Output: "255[.]100[.]50[.]0"
    fmt.Println(defangIPaddr("255.100.50.0")) // "255[.]100[.]50[.]0"
}