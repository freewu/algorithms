package main

// 7. Reverse Integer
// Given a 32-bit signed integer, reverse digits of an integer.

// Example 1:
// Input: 123
// Output:  321

// Example 2:
// Input: -123
// Output: -321

// Example 3:
// Input: 120
// Output: 21

// Constraints:
//     -2^31 <= x <= 2^31 - 1


import "fmt"
import "math"

func reverse(x int) int {
    arr := []int{}
    for {
        rem := x % 10
        x = x / 10;
        arr = append(arr, rem)
        if 0 == x {
            break
        }
    }
    res, n := 0, len(arr)
    for i := 0; i < n; i++ {
        //fmt.Println(math.Pow10(l - 1))
        //fmt.Println(i)
        res += arr[i] * int( math.Pow10(n - i - 1))
    }
    if res > math.MaxInt32 || res < math.MinInt32 {
        return 0
    }
    return res
}

func reverse1(x int) int {
    res := 0
    for {
        res = res * 10 + (x % 10)
        x /= 10;
        if x == 0 {
            break
        }
    }
    if res > math.MaxInt32 || res < math.MinInt32 {
        return 0
    }
    return res
}

func main() {
    // Example 1:
    // Input: 123
    // Output:  321
    fmt.Println(reverse(123)) // 321
    // Example 2:
    // Input: -123
    // Output: -321
    fmt.Println(reverse(-123)) // -321
    // Example 3:
    // Input: 120
    // Output: 21
    fmt.Println(reverse(120)) // 21

    fmt.Println(reverse(-1234)) // -4321
    fmt.Println(reverse(0)) // 0
    fmt.Println(reverse(1)) // 1
    fmt.Println(reverse((-1 << 30) + 1)) // 0
    fmt.Println(reverse((1 << 30) - 1)) // 0

    fmt.Println(reverse1(123)) // 321
    fmt.Println(reverse1(120)) // 21
    fmt.Println(reverse1(-123)) // -321
    fmt.Println(reverse1(-1234)) // -4321
    fmt.Println(reverse1(0)) // 0
    fmt.Println(reverse1(1)) // 1
    fmt.Println(reverse1((-1 << 30) + 1)) // 0
    fmt.Println(reverse1((1 << 30) - 1)) // 0
}
