package main

// 93. Restore IP Addresses
// A valid IP address consists of exactly four integers separated by single dots.
// Each integer is between 0 and 255 (inclusive) and cannot have leading zeros.

// For example, "0.1.2.201" and "192.168.1.1" are valid IP addresses,
// but "0.011.255.245", "192.168.1.312" and "192.168@1.1" are invalid IP addresses.
// Given a string s containing only digits, return all possible valid IP addresses that can be formed by inserting dots into s.
// You are not allowed to reorder or remove any digits in s. You may return the valid IP addresses in any order.

// Example 1:
// Input: s = "25525511135"
// Output: ["255.255.11.135","255.255.111.35"]

// Example 2:
// Input: s = "0000"
// Output: ["0.0.0.0"]

// Example 3:
// Input: s = "101023"
// Output: ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]

// Constraints:
//     1 <= s.length <= 20
//     s consists of digits only.

// 解题思路:
//     DFS
//     IP规则 0 <= x <=255

import "fmt"
import "strconv"

func restoreIpAddresses(s string) []string {
    res, ip := []string{}, []int{}
    if s == "" {
        return res
    }
    getString := func (ip []int) string {
        res := strconv.Itoa(ip[0])
        for i := 1; i < len(ip); i++ {
            res += "." + strconv.Itoa(ip[i])
        }
        return res
    }
    var dfs func(s string, index int, ip []int)
    dfs = func(s string, index int, ip []int) {
        if index == len(s) {
            if len(ip) == 4 {
                res = append(res, getString(ip))
            }
            return
        }
        if index == 0 {
            num, _ := strconv.Atoi(string(s[0]))
            ip = append(ip, num)
            dfs(s, index+1, ip)
        } else {
            num, _ := strconv.Atoi(string(s[index]))
            next := ip[len(ip)-1]*10 + num
            if next <= 255 && ip[len(ip)-1] != 0 {
                ip[len(ip)-1] = next
                dfs(s, index+1, ip)
                ip[len(ip)-1] /= 10
            }
            if len(ip) < 4 {
                ip = append(ip, num)
                dfs(s, index+1, ip)
                ip = ip[:len(ip)-1]
            }
        }
    }
    dfs(s, 0, ip)
    return res
}

// best solution
func restoreIpAddresses1(s string) []string {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    valid := func (chars []rune) bool{
        if len(chars) > 1 && chars[0] == '0'{
            return false
        }
        s:= string(chars)
        num,err := strconv.Atoi(s)
        if err != nil{
            return false
        }
        if num <= 255 {
            return true
        }
        return false
    }
    var helper func (chars []rune, left int) []string 
    helper = func (chars []rune, left int) []string {
        res := make([]string, 0)
        if left == 0 || len(chars) == 0{
            return res
        }
        if len(chars) > left*3 || len(chars) < left*1{
            return res
        }
        mx := min(3, len(chars))
        for i := 1; i <= mx; i++{
            if valid(chars[:i]){
                bottoms := helper(chars[i:], left-1)
                if len(bottoms)==0 && left-1 == 0 && i==len(chars){
                    temp := string(chars[:i])
                    res = append(res, temp)
                }
                for _, b := range bottoms {
                    temp := string(chars[:i])
                    temp = temp+"."+b
                    res = append(res, temp)
                }
            }
        }
        return res
    }
    return helper([]rune(s), 4)
}

func main() {
    // Example 1:
    // Input: s = "25525511135"
    // Output: ["255.255.11.135","255.255.111.35"]
    fmt.Printf("restoreIpAddresses(\"25525511135\") = %v\n",restoreIpAddresses("25525511135")) // ["255.255.11.135","255.255.111.35"]
    // Example 2:
    // Input: s = "0000"
    // Output: ["0.0.0.0"]
    fmt.Printf("restoreIpAddresses(\"0000\") = %v\n",restoreIpAddresses("0000")) // ["0.0.0.0"]
    // Example 3:
    // Input: s = "101023"
    // Output: ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
    fmt.Printf("restoreIpAddresses(\"101023\") = %v\n",restoreIpAddresses("101023")) // ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]

    fmt.Printf("restoreIpAddresses1(\"25525511135\") = %v\n",restoreIpAddresses1("25525511135")) // ["255.255.11.135","255.255.111.35"]
    fmt.Printf("restoreIpAddresses1(\"0000\") = %v\n",restoreIpAddresses1("0000")) // ["0.0.0.0"]
    fmt.Printf("restoreIpAddresses1(\"101023\") = %v\n",restoreIpAddresses1("101023")) // ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
}
