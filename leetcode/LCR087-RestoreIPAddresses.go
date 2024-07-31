package main

// LCR 087. 复原 IP 地址
// 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能从 s 获得的 有效 IP 地址 。你可以按任何顺序返回答案。
// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
//     例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，
//     但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。

// 示例 1：
// 输入：s = "25525511135"
// 输出：["255.255.11.135","255.255.111.35"]

// 示例 2：
// 输入：s = "0000"
// 输出：["0.0.0.0"]

// 示例 3：
// 输入：s = "1111"
// 输出：["1.1.1.1"]

// 示例 4：
// 输入：s = "010010"
// 输出：["0.10.0.10","0.100.1.0"]

// 示例 5：
// 输入：s = "10203040"
// 输出：["10.20.30.40","102.0.30.40","10.203.0.40"]

// 提示：
//     0 <= s.length <= 3000
//     s 仅由数字组成

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
    fmt.Printf("restoreIpAddresses(\"25525511135\") = %v\n",restoreIpAddresses("25525511135")) // ["255.255.11.135","255.255.111.35"]
    fmt.Printf("restoreIpAddresses(\"0000\") = %v\n",restoreIpAddresses("0000")) // ["0.0.0.0"]
    fmt.Printf("restoreIpAddresses(\"101023\") = %v\n",restoreIpAddresses("101023")) // ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]

    fmt.Printf("restoreIpAddresses1(\"25525511135\") = %v\n",restoreIpAddresses1("25525511135")) // ["255.255.11.135","255.255.111.35"]
    fmt.Printf("restoreIpAddresses1(\"0000\") = %v\n",restoreIpAddresses1("0000")) // ["0.0.0.0"]
    fmt.Printf("restoreIpAddresses1(\"101023\") = %v\n",restoreIpAddresses1("101023")) // ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
}
