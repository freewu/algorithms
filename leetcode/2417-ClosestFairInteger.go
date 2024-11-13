package main

// 2417. Closest Fair Integer
// You are given a positive integer n.

// We call an integer k fair if the number of even digits in k is equal to the number of odd digits in it.

// Return the smallest fair integer that is greater than or equal to n.

// Example 1:
// Input: n = 2
// Output: 10
// Explanation: The smallest fair integer that is greater than or equal to 2 is 10.
// 10 is fair because it has an equal number of even and odd digits (one odd digit and one even digit).

// Example 2:
// Input: n = 403
// Output: 1001
// Explanation: The smallest fair integer that is greater than or equal to 403 is 1001.
// 1001 is fair because it has an equal number of even and odd digits (two odd digits and two even digits).

// Constraints:
//     1 <= n <= 10^9

import "fmt"
import "strconv"
import "math"

func closestFair(n int) int {
    cache := []int{0, 0, 10, 0, 1001, 0, 100011, 0, 10000111, 0, 1000001111}  // 缓存 [2, 4, 6, 8, 10]位数中最小的公平整数
    res, l := n, len(strconv.Itoa(n)) 
    if l % 2 == 1 { return cache[l + 1] }
    n1 := int(math.Pow10((l)))
    isFairIntegerChecker := func(n int) bool {
        count0, count1 := 0, 0 
        for n > 0 {
            r := n % 10 
            if r % 2 == 0 {
                count0++
            } else {
                count1++
            }
            n = n / 10 
        } 
        return count0 == count1
    }
    for isFairIntegerChecker(res) != true {
        res += 1 
        if res == n1 {
            res = cache[l + 2] 
            break 
        }
    }
    return res 
}

func main() {
    // Example 1:
    // Input: n = 2
    // Output: 10
    // Explanation: The smallest fair integer that is greater than or equal to 2 is 10.
    // 10 is fair because it has an equal number of even and odd digits (one odd digit and one even digit).
    fmt.Println(closestFair(2)) // 10
    // Example 2:
    // Input: n = 403
    // Output: 1001
    // Explanation: The smallest fair integer that is greater than or equal to 403 is 1001.
    // 1001 is fair because it has an equal number of even and odd digits (two odd digits and two even digits).
    fmt.Println(closestFair(403)) // 1001

    fmt.Println(closestFair(1)) // 10
    fmt.Println(closestFair(16)) // 16
    fmt.Println(closestFair(1024)) // 1025
    fmt.Println(closestFair(999_999_999)) // 1000001111
    fmt.Println(closestFair(1_000_000_000)) // 1000001111
}