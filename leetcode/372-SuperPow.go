package main

// 372. Super Pow
// Your task is to calculate a^b mod 1337 where a is a positive integer 
// and b is an extremely large positive integer given in the form of an array.

// Example 1:
// Input: a = 2, b = [3]
// Output: 8

// Example 2:
// Input: a = 2, b = [1,0]
// Output: 1024

// Example 3:
// Input: a = 1, b = [4,3,3,8,5,2]
// Output: 1

// Constraints:
//     1 <= a <= 2^31 - 1
//     1 <= b.length <= 2000
//     0 <= b[i] <= 9
//     b does not contain leading zeros.

import "fmt"

// 快速幂 
// mod 计算的几个运算性质：
//     模运算性质一：(a + b) % p = (a % p + b % p) % p
//     模运算性质二：(a - b) % p = (a % p - b % p + p) % p
//     模运算性质三：(a * b) % p = (a % p * b % p) % p
//     模运算性质四：a ^ b % p = ((a % p)^b) % p
// Example：
//     12345^678 % 1337 = (12345^670 * 12345^8) % 1337
//                     = ((12345^670 % 1337) * (12345^8 % 1337)) % 1337  ---> 利用性质 三
//                     = (((12345^67)^10 % 1337) * (12345^8 % 1337)) % 1337  ---> 乘方性质
//                     = ((12345^67 % 1337)^10) % 1337 * (12345^8 % 1337)) % 1337  ---> 利用性质 四
//                     = (((12345^67 % 1337)^10) * (12345^8 % 1337)) % 1337  ---> 反向利用性质 三
func superPow1(a int, b []int) int {
    res := 1
    quickpow := func (x, n int) int { // 快速幂计算 x^n
        res := 1
        x %= 1337
        for n > 0 {
            if (n & 1) == 1 {
                res = (res * x) % 1337
            }
            x = (x * x) % 1337
            n >>= 1
        }
        return res
    }
    for i := 0; i < len(b); i++ {
        res = (quickpow(res, 10) * quickpow(a, b[i])) % 1337
    }
    return res
}

// 暴力
// 利用上面的性质，可以得到：a^1234567 % 1337 = (a^1234560 % 1337) * (a^7 % 1337) % k = ((((a^123456) % 1337)^10)% 1337 * (a^7 % 1337))% 1337;
func superPow(a int, b []int) int {
    if len(b) == 0 {
        return 1
    }
    last, l  := b[len(b) - 1], 1
    for i := 1; i <= last; i++ { // 先计算个位的 a^x 结果，对应上面例子中的 (a^7 % 1337)% 1337
        l = l * a % 1337
    }
    temp := superPow(a, b[:len(b) - 1]) // 再计算除去个位以外的 a^y 的结果，对应上面例子中的 (a^123456) % 1337)
    f := 1
    // 对应上面例子中的 (((a^123456) % 1337)^10)% 1337
    for i := 1; i <= 10; i++ {
        f = f * temp % 1337
    }
    return f * l % 1337
}

func superPow2(a int, b []int) int {
    res, mod := 1, 1337
    pow := func (a, b int) int {
        res := 1
        for ; b > 0; b >>= 1 {
            if (b & 1) == 1 {
                res = (res * a) % mod
            }
            a = ( a * a ) % mod
        }
        return res
    }
    for i := len(b)-1; i >= 0; i-- {
        res = res * pow(a,b[i]) % mod
        a = pow(a,10)
    }
    return res
}

func main() {
    // Example 1:
    // Input: a = 2, b = [3]
    // Output: 8
    fmt.Println(superPow(2, []int{3})) // 8 2^3
    // Example 2:
    // Input: a = 2, b = [1,0]
    // Output: 1024
    fmt.Println(superPow(2, []int{1,0})) // 1024 2^10
    // Example 3:
    // Input: a = 1, b = [4,3,3,8,5,2]
    // Output: 1
    fmt.Println(superPow(1, []int{ 4,3,3,8,5,2 })) // 1 1^433852

    fmt.Println(superPow1(2, []int{3})) // 8 2^3
    fmt.Println(superPow1(2, []int{1,0})) // 1024 2^10
    fmt.Println(superPow1(1, []int{ 4,3,3,8,5,2 })) // 1 1^433852

    fmt.Println(superPow2(2, []int{3})) // 8 2^3
    fmt.Println(superPow2(2, []int{1,0})) // 1024 2^10
    fmt.Println(superPow2(1, []int{ 4,3,3,8,5,2 })) // 1 1^433852
}