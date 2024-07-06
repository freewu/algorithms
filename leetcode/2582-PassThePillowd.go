package main

// 2582. Pass the Pillow
// There are n people standing in a line labeled from 1 to n. 
// The first person in the line is holding a pillow initially. 
// Every second, the person holding the pillow passes it to the next person standing in the line. 
// Once the pillow reaches the end of the line, the direction changes, and people continue passing the pillow in the opposite direction.
//     For example, once the pillow reaches the nth person they pass it to the n - 1th person, 
//     then to the n - 2th person and so on.

// Given the two positive integers n and time, return the index of the person holding the pillow after time seconds.

// Example 1:
// Input: n = 4, time = 5
// Output: 2
// Explanation: People pass the pillow in the following way: 1 -> 2 -> 3 -> 4 -> 3 -> 2.
// After five seconds, the 2nd person is holding the pillow.

// Example 2:
// Input: n = 3, time = 2
// Output: 3
// Explanation: People pass the pillow in the following way: 1 -> 2 -> 3.
// After two seconds, the 3rd person is holding the pillow.

// Constraints:
//     2 <= n <= 1000
//     1 <= time <= 1000

import "fmt"
// 1 -> 2 -> 3 
// 1 <- 2 <- 3
// 模拟
func passThePillow(n int, time int) int {
    res, flag := 1, true
    for time > 0 {
        if flag { // -> 传递
            res++
        } else { // <- 传递
            res--
        }
        if res == 1 || res == n { // 换传递方向
            flag = !flag
        }
        time--
    }
    return res
}

// math
func passThePillow1(n int, time int) (a int) {
    if n > time {
        return time + 1
    }
    res := time % ((n-1) * 2) + 1
    if res > n {
        return n - time % (n-1)
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 4, time = 5
    // Output: 2
    // Explanation: People pass the pillow in the following way: 1 -> 2 -> 3 -> 4 -> 3 -> 2.
    // After five seconds, the 2nd person is holding the pillow.
    fmt.Println(passThePillow(4, 5)) // 2
    // Example 2:
    // Input: n = 3, time = 2
    // Output: 3
    // Explanation: People pass the pillow in the following way: 1 -> 2 -> 3.
    // After two seconds, the 3rd person is holding the pillow.
    fmt.Println(passThePillow(3, 2)) // 3

    fmt.Println(passThePillow(1000, 1000)) // 999

    fmt.Println(passThePillow1(4, 5)) // 2
    fmt.Println(passThePillow1(3, 2)) // 3
    fmt.Println(passThePillow1(1000, 1000)) // 999
}