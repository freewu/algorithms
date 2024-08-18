package main

// 842. Split Array into Fibonacci Sequence
// You are given a string of digits num, such as "123456579". 
// We can split it into a Fibonacci-like sequence [123, 456, 579].

// Formally, a Fibonacci-like sequence is a list f of non-negative integers such that:
//     0 <= f[i] < 231, (that is, each integer fits in a 32-bit signed integer type),
//     f.length >= 3, and
//     f[i] + f[i + 1] == f[i + 2] for all 0 <= i < f.length - 2.

// Note that when splitting the string into pieces, each piece must not have extra leading zeroes, 
// except if the piece is the number 0 itself.

// Return any Fibonacci-like sequence split from num, or return [] if it cannot be done.

// Example 1:
// Input: num = "1101111"
// Output: [11,0,11,11]
// Explanation: The output [110, 1, 111] would also be accepted.

// Example 2:
// Input: num = "112358130"
// Output: []
// Explanation: The task is impossible.

// Example 3:
// Input: num = "0123"
// Output: []
// Explanation: Leading zeroes are not allowed, so "01", "2", "3" is not valid.

// Constraints:
//     1 <= num.length <= 200
//     num contains only digits.

import "fmt"
import "strconv"

func splitIntoFibonacci(num string) []int {
    res, n  := []int{}, len(num)
    //"a" denotes the position of the first digit of F[1]
    for a := 1; a < n - 1; a ++{ 
        if num[0] == '0' && a > 1 { // avoid cases such as "01"
            continue
        }
        n1, _ := strconv.Atoi(num[0 : a]) // n1 = F[0]
        if n1 > 2147483647 {
            continue
        }
        //"b" denotes the position of the first digit of F[2]
        for b := a + 1; b < n; b ++{
            if num[a] == '0' && b > a + 1 { // avoid cases such as "01"
                continue
            }    
            n2, _ := strconv.Atoi(num[a : b]) // n2 = F[1]
            s := strconv.Itoa(n1 + n2) // s = string(F[2])
            m1 := n1; m2 := n2
            h := b; t := h + len(s) - 1 // "h" and "t" denote the head and tail position of the sum
            fibonacci := true
            for h < n {
                if t >= n || s != num[h : t + 1] || m1 + m2 > 2147483647 {
                    fibonacci = false
                    break
                }
                m1, m2 = m2, m1 + m2
                s = strconv.Itoa(m1 + m2)
                h = t + 1; t = h + len(s) - 1
            }
            if fibonacci { // to save in the result array and return result
                v1, _ := strconv.Atoi(num[0 : a])
                res = append(res, v1)
                v2, _ := strconv.Atoi(num[a : b])
                res = append(res, v2)
                h := b // head and tail positions for the sum
                for t := h + len(strconv.Itoa(v1 + v2)) - 1; t < n; {
                    res = append(res, v1 + v2)
                    h = t + 1; t = h + len(strconv.Itoa(v1 + v2)) - 1
                    v1, v2 = v2, v1 + v2
                }
                return res
            }
        }
    }    
    return nil
}

func splitIntoFibonacci1(num string) []int {
    res, n, inf := []int{}, len(num), 1 << 31
    var backtrack func(index, sum, prev int) bool
    backtrack = func(index, sum, prev int) bool {
        if index == n {
            return len(res) >= 3
        }
        cur := 0
        for i := index; i < n; i++ {
            // 每个块的数字一定不要以零开头，除非这个块是数字 0 本身
            if i > index && num[index] == '0' {
                break
            }
            cur = cur * 10 + int(num[i]-'0')
            if cur > inf { // 拆出的整数要符合 32 位有符号整数类型
                break
            }
            // res[i] + res[i+1] = res[i+2]
            if len(res) >= 2 {
                if cur < sum { continue }
                if cur > sum { break }
            }
            // cur 符合要求，加入序列 res
            res = append(res, cur)
            if backtrack(i+1, prev + cur, cur) {
                return true
            }
            res = res[:len(res)-1]
        }
        return false
    }
    backtrack(0, 0, 0)
    return res
}

func main() {
    // Example 1:
    // Input: num = "1101111"
    // Output: [11,0,11,11]
    // Explanation: The output [110, 1, 111] would also be accepted.
    fmt.Println(splitIntoFibonacci("1101111")) // [11,0,11,11] | [110, 1, 111]
    // Example 2:
    // Input: num = "112358130"
    // Output: []
    // Explanation: The task is impossible.
    fmt.Println(splitIntoFibonacci("112358130")) // []
    // Example 3:
    // Input: num = "0123"
    // Output: []
    // Explanation: Leading zeroes are not allowed, so "01", "2", "3" is not valid.
    fmt.Println(splitIntoFibonacci("0123")) // []

    fmt.Println(splitIntoFibonacci1("1101111")) // [11,0,11,11] | [110, 1, 111]
    fmt.Println(splitIntoFibonacci1("112358130")) // []
    fmt.Println(splitIntoFibonacci1("0123")) // []
}