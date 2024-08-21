package main

// 3133. Minimum Array End
// You are given two integers n and x. 
// You have to construct an array of positive integers nums of size n where for every 0 <= i < n - 1, 
// nums[i + 1] is greater than nums[i], 
// and the result of the bitwise AND operation between all elements of nums is x.

// Return the minimum possible value of nums[n - 1].

// Example 1:
// Input: n = 3, x = 4
// Output: 6
// Explanation:
// nums can be [4,5,6] and its last element is 6.

// Example 2:
// Input: n = 2, x = 7
// Output: 15
// Explanation:
// nums can be [7,15] and its last element is 15.

// Constraints:
//     1 <= n, x <= 10^8

import "fmt"

func minEnd(n int, x int) int64 {
    res := x
    for n > 1 {
        res = (res + 1) | x
        n--
    }
    return int64(res)
}

func minEnd1(n, x int) int64 {
    n-- // 先把 n 减一，这样下面讨论的 n 就是原来的 n-1
    i, j := 0, 0
    for n >> j > 0 {
        // x 的第 i 个比特值是 0，即「空位」
        if x >> i &1  == 0 {
            // 空位填入 n 的第 j 个比特值
            x |= n >> j & 1 << i
            j++
        }
        i++
    }
    return int64(x)
}

func main() {
    // Example 1:
    // Input: n = 3, x = 4
    // Output: 6
    // Explanation:
    // nums can be [4,5,6] and its last element is 6.
    fmt.Println(minEnd(3, 4)) // 6
    // Example 2:
    // Input: n = 2, x = 7
    // Output: 15
    // Explanation:
    // nums can be [7,15] and its last element is 15.
    fmt.Println(minEnd(2, 7)) // 15

    fmt.Println(minEnd1(3, 4)) // 6
    fmt.Println(minEnd1(2, 7)) // 15
}