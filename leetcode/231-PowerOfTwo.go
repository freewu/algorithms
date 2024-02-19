package main

// 231. Power of Two
// Given an integer n, return true if it is a power of two. Otherwise, return false.
// An integer n is a power of two, if there exists an integer x such that n == 2x.

// Example 1:
// Input: n = 1
// Output: true
// Explanation: 2^0 = 1

// Example 2:
// Input: n = 16
// Output: true
// Explanation: 2^4 = 16

// Example 3:
// Input: n = 3
// Output: false

// Example 4:
// Input: n = 4
// Output: true
// Explanation: 2^2 = 4

// Example 5:
// Input: n = 5
// Output: false

// Constraints:
//         -2^31 <= n <= 2^31 - 1

// Follow up: Could you solve it without loops/recursion?

import "fmt"

// 位运算
func isPowerOfTwo(num int) bool {
	return (num > 0 && ((num & (num - 1)) == 0))
}

// 数论
func isPowerOfTwo1(num int) bool {
    // 2^30 = 1073741824
    // 2^31 = 21474836482
	return num > 0 && (1073741824%num == 0)
}

// 枚举 map
func isPowerOfTwo2(num int) bool {
	allPowerOfTwoMap := map[int]int{1: 1, 2: 2, 4: 4, 8: 8, 16: 16, 32: 32, 64: 64, 128: 128, 256: 256, 512: 512, 1024: 1024, 2048: 2048, 4096: 4096, 8192: 8192, 16384: 16384, 32768: 32768, 65536: 65536, 131072: 131072, 262144: 262144, 524288: 524288, 1048576: 1048576, 2097152: 2097152, 4194304: 4194304, 8388608: 8388608, 16777216: 16777216, 33554432: 33554432, 67108864: 67108864, 134217728: 134217728, 268435456: 268435456, 536870912: 536870912, 1073741824: 1073741824}
	_, ok := allPowerOfTwoMap[num]
	return ok
}

// 循环 % 2 
func isPowerOfTwo3(num int) bool {
	for num >= 2 {
        // 每次除一半 判断是否能被 2 整除
		if num % 2 == 0 {
			num = num / 2
		} else {
			return false
		}
	}
	return num == 1
}

func main() {
    fmt.Println(isPowerOfTwo(1)) // true 2^0
    fmt.Println(isPowerOfTwo(16)) // true 2^4
    fmt.Println(isPowerOfTwo(3)) // false
    fmt.Println(isPowerOfTwo(4)) // true 2^2
    fmt.Println(isPowerOfTwo(5)) // false

    fmt.Println(isPowerOfTwo1(1)) // true 2^0
    fmt.Println(isPowerOfTwo1(16)) // true 2^4
    fmt.Println(isPowerOfTwo1(3)) // false
    fmt.Println(isPowerOfTwo1(4)) // true 2^2
    fmt.Println(isPowerOfTwo1(5)) // false

    fmt.Println(isPowerOfTwo2(1)) // true 2^0
    fmt.Println(isPowerOfTwo2(16)) // true 2^4
    fmt.Println(isPowerOfTwo2(3)) // false
    fmt.Println(isPowerOfTwo2(4)) // true 2^2
    fmt.Println(isPowerOfTwo2(5)) // false

    fmt.Println(isPowerOfTwo3(1)) // true 2^0
    fmt.Println(isPowerOfTwo3(16)) // true 2^4
    fmt.Println(isPowerOfTwo3(3)) // false
    fmt.Println(isPowerOfTwo3(4)) // true 2^2
    fmt.Println(isPowerOfTwo3(5)) // false
}