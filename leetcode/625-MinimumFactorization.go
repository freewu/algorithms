package main

// 625. Minimum Factorization
// Given a positive integer num, return the smallest positive integer x whose multiplication of each digit equals num. 
// If there is no answer or the answer is not fit in 32-bit signed integer, return 0.

// Example 1:
// Input: num = 48
// Output: 68

// Example 2:
// Input: num = 15
// Output: 35

// Constraints:
//     1 <= num <= 2^31 - 1

import "fmt"
import "strconv"

// 暴力枚举 超出时间限制
func smallestFactorization(num int) int {
    inf := 1 << 32 - 1
    for i := 1; i < 999999999; i++ {
        mul, t := 1, i
        for t != 0 {
            mul *= t % 10
            t /= 10
        }
        if mul == num && mul <= inf {
            return i
        }
    }
    return 0
}

// dfs
func smallestFactorization1(num int) int {
    res, inf := 0, 1 << 32 - 1
    if num < 2 {
        return num
    }
    dig := []int{9, 8, 7, 6, 5, 4, 3, 2}
    var dfs func(dig []int, i int,  a int,  mul int, s string) bool
    dfs = func(dig []int, i int,  a int,  mul int, s string) bool {
        if mul > a || i == len(dig) {
            return false;
        }
        if mul == a {
            res, _ = strconv.Atoi(s)
            return true;
        }
        return dfs(dig, i, a, mul * dig[i], strconv.Itoa(dig[i]) + s) || dfs(dig, i + 1, a, mul, s)
    }
    if dfs(dig, 0, num, 1, "") && res <= inf {
        return res
    }
    return 0
}

// 因子分解
func smallestFactorization2(num int) int {
    if num < 2 {
        return num
    }
    res, mul := 0, 1
    for i := 9; i >= 2; i-- {
        for num % i == 0 {
            num /= i
            res = mul * i + res
            mul *= 10
        }
    }
    if num < 2 && res <= 2147483647 {
        return int(res)
    }
    return 0
}

func main() {
    // Example 1:
    // Input: num = 48
    // Output: 68
    fmt.Println(smallestFactorization(48)) // 68
    // Example 2:
    // Input: num = 15
    // Output: 35
    fmt.Println(smallestFactorization(15)) // 35

    fmt.Println(smallestFactorization1(48)) // 68
    fmt.Println(smallestFactorization1(15)) // 35

    fmt.Println(smallestFactorization2(48)) // 68
    fmt.Println(smallestFactorization2(15)) // 35
}