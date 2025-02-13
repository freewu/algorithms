package main

// 3370. Smallest Number With All Set Bits
// You are given a positive number n.

// Return the smallest number x greater than or equal to n, 
// such that the binary representation of x contains only set bits

// Example 1:
// Input: n = 5
// Output: 7
// Explanation:
// The binary representation of 7 is "111".

// Example 2:
// Input: n = 10
// Output: 15
// Explanation:
// The binary representation of 15 is "1111".

// Example 3:
// Input: n = 3
// Output: 3
// Explanation:
// The binary representation of 3 is "11".

// Constraints:
//     1 <= n <= 1000

import "fmt"
import "math"
import "math/bits"

func smallestNumber(n int) int {
    if n == 1 { return n }
    log := func(base, x float64) float64 { return math.Log(x) / math.Log(base); }
    power := math.Floor(log(2, float64(n)))
    return int(math.Pow(2, power + 1)) - 1
}

func smallestNumber1(n int) int {
    v := 2
    for {
        if v - 1 >= n {
            return v - 1
        }
        v *= 2
    }
    return -1
}

func smallestNumber2(n int) int {
    return 1 << bits.Len(uint(n)) - 1
}

func main() {
    // Example 1:
    // Input: n = 5
    // Output: 7
    // Explanation:
    // The binary representation of 7 is "111".
    fmt.Println(smallestNumber(5)) // 7
    // Example 2:
    // Input: n = 10
    // Output: 15
    // Explanation:
    // The binary representation of 15 is "1111".
    fmt.Println(smallestNumber(10)) // 15
    // Example 3:
    // Input: n = 3
    // Output: 3
    // Explanation:
    // The binary representation of 3 is "11".
    fmt.Println(smallestNumber(3)) // 3

    fmt.Println(smallestNumber(1)) // 1
    fmt.Println(smallestNumber(2)) // 3
    fmt.Println(smallestNumber(8)) // 15
    fmt.Println(smallestNumber(64)) // 127
    fmt.Println(smallestNumber(99)) // 127
    fmt.Println(smallestNumber(100)) // 127
    fmt.Println(smallestNumber(128)) // 255
    fmt.Println(smallestNumber(256)) // 511
    fmt.Println(smallestNumber(999)) // 1023
    fmt.Println(smallestNumber(1000)) // 1023

    fmt.Println(smallestNumber1(5)) // 7
    fmt.Println(smallestNumber1(10)) // 15
    fmt.Println(smallestNumber1(3)) // 3
    fmt.Println(smallestNumber1(1)) // 1
    fmt.Println(smallestNumber1(2)) // 3
    fmt.Println(smallestNumber1(8)) // 15
    fmt.Println(smallestNumber1(64)) // 127
    fmt.Println(smallestNumber1(99)) // 127
    fmt.Println(smallestNumber1(100)) // 127
    fmt.Println(smallestNumber1(128)) // 255
    fmt.Println(smallestNumber1(256)) // 511
    fmt.Println(smallestNumber1(999)) // 1023
    fmt.Println(smallestNumber1(1000)) // 1023

    fmt.Println(smallestNumber2(5)) // 7
    fmt.Println(smallestNumber2(10)) // 15
    fmt.Println(smallestNumber2(3)) // 3
    fmt.Println(smallestNumber2(1)) // 1
    fmt.Println(smallestNumber2(2)) // 3
    fmt.Println(smallestNumber2(8)) // 15
    fmt.Println(smallestNumber2(64)) // 127
    fmt.Println(smallestNumber2(99)) // 127
    fmt.Println(smallestNumber2(100)) // 127
    fmt.Println(smallestNumber2(128)) // 255
    fmt.Println(smallestNumber2(256)) // 511
    fmt.Println(smallestNumber2(999)) // 1023
    fmt.Println(smallestNumber2(1000)) // 1023
}