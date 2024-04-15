package main

// 1492. The kth Factor of n
// You are given two positive integers n and k. 
// A factor of an integer n is defined as an integer i where n % i == 0.

// Consider a list of all factors of n sorted in ascending order, 
// return the kth factor in this list or return -1 if n has less than k factors.

// Example 1:
// Input: n = 12, k = 3
// Output: 3
// Explanation: Factors list is [1, 2, 3, 4, 6, 12], the 3rd factor is 3.

// Example 2:
// Input: n = 7, k = 2
// Output: 7
// Explanation: Factors list is [1, 7], the 2nd factor is 7.

// Example 3:
// Input: n = 4, k = 4
// Output: -1
// Explanation: Factors list is [1, 2, 4], there is only 3 factors. We should return -1.

// Constraints:
//     1 <= k <= n <= 1000

// Follow up:
//     Could you solve this problem in less than O(n) complexity?

import "fmt"

func kthFactor(n int, k int) int {
    cnt := 0
    for i := 1; i < n / 2 + 1; i++ {
        if n % i == 0 { // 正整数 i 满足 n % i == 0 ，那么我们就说正整数 i 是整数 n 的因子
            cnt++ // 累加找到的因子数
        } 
        if cnt == k { // 你返回第 k 个因子
            return i
        }
    }
    if cnt == k - 1 { // 是 cnt = k - 1 说明 第 k 个因子就是 n 本身
        return n
    }
    return -1
}

func main() {
    // Explanation: Factors list is [1, 2, 3, 4, 6, 12], the 3rd factor is 3.
    fmt.Println(kthFactor(12, 3)) // 3
    // Explanation: Factors list is [1, 7], the 2nd factor is 7.
    fmt.Println(kthFactor(7, 2)) // 3
    // Explanation: Factors list is [1, 2, 4], there is only 3 factors. We should return -1.
    fmt.Println(kthFactor(4, 24)) // -1
}