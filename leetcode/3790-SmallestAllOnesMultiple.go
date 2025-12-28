package main

// 3790. Smallest All-Ones Multiple
// You are given a positive integer k.

// Find the smallest integer n divisible by k that consists of only the digit 1 in its decimal representation (e.g., 1, 11, 111, ...).

// Return an integer denoting the number of digits in the decimal representation of n. If no such n exists, return -1.

// Example 1:
// Input: k = 3
// Output: 3
// Explanation:
// n = 111 because 111 is divisible by 3, but 1 and 11 are not. The length of n = 111 is 3.

// Example 2:
// Input: k = 7
// Output: 6
// Explanation:
// n = 111111. The length of n = 111111 is 6.

// Example 3:
// Input: k = 2
// Output: -1
// Explanation:
// There does not exist a valid n that is a multiple of 2.

// Constraints:
//     2 <= k <= 10^5

import "fmt"
import "math"

func minAllOneMultiple(k int) int {
    var pow func(a, b int) int
    pow = func(a, b int) int {
        if b == 0 {return 1}
        tmp := pow((a * a) % k, b / 2) % k
        if b & 1 == 1 { return (tmp * (a % k)) % k }
        return tmp
    }
    visited := make(map[int]bool)
    temp := 1
    for i := 1; i < math.MaxInt; i++ {
        if temp % k == 0 { return i }
        if visited[temp] { return -1 }
        visited[temp] = true    
        temp = (temp + pow(10 % k, i)) % k
    }
    return -1
}

func minAllOneMultiple1(k int) int {
    if k % 2 == 0 || k % 5 == 0 { 
        return -1 
    }
    cur, res := 0, 1
    for {
        cur = (cur * 10 + 1) % k
        if cur == 0 { 
            return res 
        }
        res++
    }
    return -1
}

func main() {
    // Example 1:
    // Input: k = 3
    // Output: 3
    // Explanation:
    // n = 111 because 111 is divisible by 3, but 1 and 11 are not. The length of n = 111 is 3.
    fmt.Println(minAllOneMultiple(3)) // 3
    // Example 2:
    // Input: k = 7
    // Output: 6
    // Explanation:
    // n = 111111. The length of n = 111111 is 6.
    fmt.Println(minAllOneMultiple(7)) // 6
    // Example 3:
    // Input: k = 2
    // Output: -1
    // Explanation:
    // There does not exist a valid n that is a multiple of 2.  
    fmt.Println(minAllOneMultiple(2)) // -1

    fmt.Println(minAllOneMultiple(8)) // -1
    fmt.Println(minAllOneMultiple(64)) // -1
    fmt.Println(minAllOneMultiple(99)) // 18
    fmt.Println(minAllOneMultiple(100)) // -1
    fmt.Println(minAllOneMultiple(999)) // 27
    fmt.Println(minAllOneMultiple(1024)) // -1
    fmt.Println(minAllOneMultiple(99_999)) // 45
    fmt.Println(minAllOneMultiple(100_000)) // -1

    fmt.Println(minAllOneMultiple1(3)) // 3
    fmt.Println(minAllOneMultiple1(7)) // 6 
    fmt.Println(minAllOneMultiple1(2)) // -1
    fmt.Println(minAllOneMultiple1(8)) // -1
    fmt.Println(minAllOneMultiple1(64)) // -1
    fmt.Println(minAllOneMultiple1(99)) // 18
    fmt.Println(minAllOneMultiple1(100)) // -1
    fmt.Println(minAllOneMultiple1(999)) // 27
    fmt.Println(minAllOneMultiple1(1024)) // -1
    fmt.Println(minAllOneMultiple1(99_999)) // 45
    fmt.Println(minAllOneMultiple1(100_000)) // -1
}