package main

// 1058. Minimize Rounding Error to Meet Target
// Given an array of prices [p1,p2...,pn] and a target, 
// round each price pi to Roundi(pi) so that the rounded array [Round1(p1),Round2(p2)...,Roundn(pn)] sums to the given target. 
// Each operation Roundi(pi) could be either Floor(pi) or Ceil(pi).

// Return the string "-1" if the rounded array is impossible to sum to target. 
// Otherwise, return the smallest rounding error, which is defined as Σ |Roundi(pi) - (pi)| for i from 1 to n, as a string with three places after the decimal.

// Example 1:
// Input: prices = ["0.700","2.800","4.900"], target = 8
// Output: "1.000"
// Explanation:
// Use Floor, Ceil and Ceil operations to get (0.7 - 0) + (3 - 2.8) + (5 - 4.9) = 0.7 + 0.2 + 0.1 = 1.0 .

// Example 2:
// Input: prices = ["1.500","2.500","3.500"], target = 10
// Output: "-1"
// Explanation: It is impossible to meet the target.

// Example 3:
// Input: prices = ["1.500","2.500","3.500"], target = 9
// Output: "1.500"

// Constraints:
//     1 <= prices.length <= 500
//     Each string prices[i] represents a real number in the range [0.0, 1000.0] and has exactly 3 decimal places.
//     0 <= target <= 10^6

import "fmt"
import "sort"
import "strconv"

func minimizeError(prices []string, target int) string {
    toNums := func(price string) (int, int) {
        // 1.500 -> 1, 500
        a := 0
        for i := range price {
            if price[i] == '.' {
                continue
            }
            a = 10 * a + int(price[i] - '0')
        }
        return a / 1000, a % 1000
    }
    toString := func (b int) string {
        // 1000 -> 1.000
        a := b / 1000
        b = b % 1000
        res := strconv.Itoa(a) + "."
        if b < 10 {
            res += "00"
        } else if b < 100 {
            res += "0"
        }
        return res + strconv.Itoa(b)
    }
    // 全向上 ceil 操作遍历，得到最大和，过程中也计算小数为零的个数
    // 如果每次不向上而是向下，差值为1，则可得到最小和 = 最大和 - （prices长度 - 小数为零个数)
    // 目标值，应该为最大和与最小和之间
    // 因差值为1，向下 floor 操作次数 = 最大和 - target
    mx, zeroNum := 0, 0
    roundArr, roundSum := make([]int, len(prices)), 0 // 存储 ceil 舍入误差
    for i := 0; i < len(prices); i++ {
        a, b := toNums(prices[i])
        mx += a
        if b > 0 { 
            mx++ // 向上取整后加一
            roundArr[i] = 1000 - b
            roundSum += 1000 - b
        } else {
            zeroNum++
        }
    }
    num := mx - target  // 向下 floor 操作次数
    if num < 0 || num > len(prices) - zeroNum {
        return "-1"
    }
    // 为返回最小的舍入误差，将舍入误差排序，替换为向下操作
    sort.Ints(roundArr)
    for i := len(prices) - 1; i > len(prices)-1-num; i-- {
        roundSum -= roundArr[i]
        roundSum += 1000 - roundArr[i]
    }
    return toString(roundSum)
}

func main() {
    // Example 1:
    // Input: prices = ["0.700","2.800","4.900"], target = 8
    // Output: "1.000"
    // Explanation:
    // Use Floor, Ceil and Ceil operations to get (0.7 - 0) + (3 - 2.8) + (5 - 4.9) = 0.7 + 0.2 + 0.1 = 1.0 .
    fmt.Println(minimizeError([]string{"0.700","2.800","4.900"}, 8)) // "1.000"
    // Example 2:
    // Input: prices = ["1.500","2.500","3.500"], target = 10
    // Output: "-1"
    // Explanation: It is impossible to meet the target.
    fmt.Println(minimizeError([]string{"1.500","2.500","3.500"}, 10)) // "-1"
    // Example 3:
    // Input: prices = ["1.500","2.500","3.500"], target = 9
    // Output: "1.500"
    fmt.Println(minimizeError([]string{"1.500","2.500","3.500"}, 9)) // "1.500"
}