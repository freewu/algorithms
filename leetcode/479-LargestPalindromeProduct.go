package main

// 479. Largest Palindrome Product
// Given an integer n, return the largest palindromic integer that can be represented as the product of two n-digits integers.
// Since the answer can be very large, return it modulo 1337.

// Example 1:
// Input: n = 2
// Output: 987
// Explanation: 99 x 91 = 9009, 9009 % 1337 = 987

// Example 2:
// Input: n = 1
// Output: 9
 
// Constraints:
//     1 <= n <= 8

import "fmt"
import "math"

// func largestPalindrome(n int) int {
//     // just to forget about 1-digit case
//     if n == 1 {
//         return 9
//     }
//     // min number with n digits (for ex. for n = 4, min_num = 1000) & // max number with n digits (for ex. 9999)
//     max_pal, min_num, max_num := 0, 10 ** (n - 1), 10 ** n - 1

    
//     // step is equal to 2, because we have to get a number, the 1st digit of which is 9, so we have to   
//     // iterate only over odd numbers
//     for i in range(max_num, min_num - 1, -2): 
        
//         # since we are looking for the maximum palindrome number, it makes no sense to iterate over the 
//         # product less than the max_pal obtained from the last iteration
//         if i * i < max_pal:
//             break
            
//         for j in range(max_num, i - 1, -2):
//             product = i * j
            
//             # since a palindrome with an even number of digits must be mod 11 == 0 and we have no reason to 
//             # check the product which less or equal than max_pal
//             if product % 11 != 0 and product >= max_pal:
//                 continue
                
//             # check if product is a palindrome then update the max_pal
//             if str(product) == str(product)[::-1]:
//                 max_pal = product

//     return max_pal % 1337
// }

// 打表
func largestPalindrome1(n int) int {
    mp := []int{9,987,123,597,677,1218,877,475}
    return mp[n-1]
}

func largestPalindrome(n int) int {
    if n == 1 {
        return 9
    }
    upper := int(math.Pow10(n) - 1)
    for left := upper; ; left-- { // 10^n - 1 枚举左半部分
        num := left
        for x := left; x > 0; x /= 10 {  // 翻转左半部分到其自身末尾，构造回文数
            num = num * 10 + x % 10       
        }
        for y := upper; y*y >= num; y-- {
            if num % y == 0 { // y 是 p 的因子
                return num % 1337
            }
        }
    }
}

func main() {
    // Example 1:
    // Input: n = 2
    // Output: 987
    // Explanation: 99 x 91 = 9009, 9009 % 1337 = 987
    fmt.Println(largestPalindrome(2)) // 987
    // Example 2:
    // Input: n = 1
    // Output: 9
    fmt.Println(largestPalindrome(1)) // 9

    fmt.Println(largestPalindrome(3)) // 123
    fmt.Println(largestPalindrome(4)) // 597
    fmt.Println(largestPalindrome(5)) // 677
    fmt.Println(largestPalindrome(6)) // 1218
    fmt.Println(largestPalindrome(7)) // 877
    fmt.Println(largestPalindrome(8)) // 475


    fmt.Println(largestPalindrome1(2)) // 987
    fmt.Println(largestPalindrome1(1)) // 9
    fmt.Println(largestPalindrome1(3)) // 123
    fmt.Println(largestPalindrome1(4)) // 597
    fmt.Println(largestPalindrome1(5)) // 677
    fmt.Println(largestPalindrome1(6)) // 1218
    fmt.Println(largestPalindrome1(7)) // 877
    fmt.Println(largestPalindrome1(8)) // 475
}