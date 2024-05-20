package main

// 468. Validate IP Address
// Given a string queryIP, 
//     return "IPv4" if IP is a valid IPv4 address, 
//     "IPv6" if IP is a valid IPv6 address or "Neither" 
//     if IP is not a correct IP of any type.

// A valid IPv4 address is an IP in the form "x1.x2.x3.x4" where 0 <= xi <= 255 and xi cannot contain leading zeros.
//     For example, "192.168.1.1" and "192.168.1.0" are valid IPv4 addresses 
//     while "192.168.01.1", "192.168.1.00", and "192.168@1.1" are invalid IPv4 addresses.

// A valid IPv6 address is an IP in the form "x1:x2:x3:x4:x5:x6:x7:x8" where:
//     1 <= xi.length <= 4
//     xi is a hexadecimal string which may contain digits, lowercase English letter ('a' to 'f') and upper-case English letters ('A' to 'F').
//     Leading zeros are allowed in xi.

//     For example, "2001:0db8:85a3:0000:0000:8a2e:0370:7334" and "2001:db8:85a3:0:0:8A2E:0370:7334" are valid IPv6 addresses, 
//     while "2001:0db8:85a3::8A2E:037j:7334" and "02001:0db8:85a3:0000:0000:8a2e:0370:7334" are invalid IPv6 addresses.

// Example 1:
// Input: queryIP = "172.16.254.1"
// Output: "IPv4"
// Explanation: This is a valid IPv4 address, return "IPv4".

// Example 2:
// Input: queryIP = "2001:0db8:85a3:0:0:8A2E:0370:7334"
// Output: "IPv6"
// Explanation: This is a valid IPv6 address, return "IPv6".

// Example 3:
// Input: queryIP = "256.256.256.256"
// Output: "Neither"
// Explanation: This is neither a IPv4 address nor a IPv6 address.

// Constraints:
//     queryIP consists only of English letters, digits and the characters '.' and ':'.

import "fmt"
import "strconv"

func validIPAddress(queryIP string) string {
    if isIPv4(queryIP) { return "IPv4" }
    if isIPv6(queryIP) { return "IPv6" }
    return "Neither"
}

func isIPv4(s string) bool {
    i, j, groups := 0, 0, 0
    for j < len(s) + 1 {
        if j < len(s) {
            if int(s[j]) < 48 || int(s[j]) > 57 { // ASCII for 0-9
                if int(s[j]) != 46 {  return false } // ASCII for '.'
            }
        }
        if j == len(s) || s[j] == '.' {
            if j - i == 0 { return false }
            if j - i > 1 && s[i] == '0' { return false } // no leading zeroes
            if j - i > 3 { return false } // max 3 digits
            v, _ := strconv.Atoi(s[i:j])
            if v > 255 { return false }
            i = j + 1
            groups++
        }
        j++
    }
    if groups != 4 { return false }
    return true
}

func isIPv6(s string) bool {
    i, j, groups := 0, 0, 0
    for j < len(s) + 1 {
        if j < len(s) {
            if !(int(s[j]) >= 48 && int(s[j]) <= 57) && !(int(s[j]) >= 65 && int(s[j]) <= 70) &&
                !(int(s[j]) >= 97 && int(s[j]) <= 102) && int(s[j]) != 58 { // not 0-9 or a-f or A-F or :
                return false
            }
        }
        if j == len(s) || s[j] == ':' {
            if j - i == 0 { return false }
            if j - i > 4 { return false } // max 4 digits
            i = j + 1
            groups++
        }
        j++
    }
    if groups != 8 { return false }
    return true
}

func main() {
    // Example 1:
    // Input: queryIP = "172.16.254.1"
    // Output: "IPv4"
    // Explanation: This is a valid IPv4 address, return "IPv4".
    fmt.Println(validIPAddress("172.16.254.1")) // IPv4
    // Example 2:
    // Input: queryIP = "2001:0db8:85a3:0:0:8A2E:0370:7334"
    // Output: "IPv6"
    // Explanation: This is a valid IPv6 address, return "IPv6".
    fmt.Println(validIPAddress("2001:0db8:85a3:0:0:8A2E:0370:7334")) // IPv6
    // Example 3:
    // Input: queryIP = "256.256.256.256"
    // Output: "Neither"
    // Explanation: This is neither a IPv4 address nor a IPv6 address.
    fmt.Println(validIPAddress("256.256.256.256")) // Neither
}