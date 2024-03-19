package main

// 1969. Minimum Non-Zero Product of the Array Elements
// You are given a positive integer p. 
// Consider an array nums (1-indexed) that consists of the integers in the inclusive range [1, 2p - 1] in their binary representations. 
// You are allowed to do the following operation any number of times:
//     Choose two elements x and y from nums.
//     Choose a bit in x and swap it with its corresponding bit in y. Corresponding bit refers to the bit that is in the same position in the other integer.

// For example, if x = 1101 and y = 0011, after swapping the 2nd bit from the right, we have x = 1111 and y = 0001.
// Find the minimum non-zero product of nums after performing the above operation any number of times. 
// Return this product modulo 10^9 + 7.

// Note: The answer should be the minimum product before the modulo operation is done.

// Example 1:
// Input: p = 1
// Output: 1
// Explanation: nums = [1].
// There is only one element, so the product equals that element.

// Example 2:
// Input: p = 2
// Output: 6
// Explanation: nums = [01, 10, 11].
// Any swap would either make the product 0 or stay the same.
// Thus, the array product of 1 * 2 * 3 = 6 is already minimized.

// Example 3:
// Input: p = 3
// Output: 1512
// Explanation: nums = [001, 010, 011, 100, 101, 110, 111]
// - In the first operation we can swap the leftmost bit of the second and fifth elements.
//     - The resulting array is [001, 110, 011, 100, 001, 110, 111].
// - In the second operation we can swap the middle bit of the third and fourth elements.
//     - The resulting array is [001, 110, 001, 110, 001, 110, 111].
// The array product is 1 * 6 * 1 * 6 * 1 * 6 * 7 = 1512, which is the minimum possible product.
 
// Constraints:
//     1 <= p <= 60

import "fmt"

func minNonZeroProduct(p int) int {
    powerOfTwo := make([]int, p + 1)
    powerOfTwo[0] = 1
    baseValue := 1
    // 跑出相关的2^p的 数据
    // p = 1 [1 2] [01,11]
    // p = 2 [1 2 4] [001,010,100]
    // p = 3 [1 2 4 8] [0001,0010,0100,1000]
    for i := 1; i <= p; i++ {
        baseValue *= 2
        powerOfTwo[i] = baseValue
    }
    // fmt.Println(powerOfTwo)
    // products[i]: 2**i products[0] multi result
    products := make([]int, p+1)
    products[0] = (baseValue - 2) % 1000000007
    for i := 1; i <= p; i++ {
        products[i] = products[i-1] * products[i-1] % 1000000007
    }
    factor := (baseValue - 2) / 2
    minProduct := (baseValue - 1) % 1000000007
    for index := p; factor > 0; {
        for powerOfTwo[index] > factor {
            index--
        }
        factor -= powerOfTwo[index]
        minProduct = minProduct * products[index] % 1000000007
    }
    return minProduct
}

func minNonZeroProduct1(p int) int {
    if p == 1 {
        return 1
    }

    fastPow := func(x, n, mod int64) int64 {
        res := int64(1)
        for n != 0 {
            if n & 1 == 1 {
                res = (res * x) % mod
            }
            x = (x * x) % mod
            n >>= 1
        }
        return res
    }

    mod := int64(1e9 + 7)
    x := fastPow(2, int64(p), mod) - 1
    y := int64(1) << uint(p - 1)
    return int(fastPow(x - 1, y - 1, mod) * x % mod)
}



func main() {
    // Explanation: nums = [1].
    // There is only one element, so the product equals that element.
    fmt.Println(minNonZeroProduct(1)) // 1

    // Explanation: nums = [01, 10, 11].
    // Any swap would either make the product 0 or stay the same.
    // Thus, the array product of 1 * 2 * 3 = 6 is already minimized.
    fmt.Println(minNonZeroProduct(2)) // 6

    // Explanation: nums = [001, 010, 011, 100, 101, 110, 111]
    // - In the first operation we can swap the leftmost bit of the second and fifth elements.
    //     - The resulting array is [001, 110, 011, 100, 001, 110, 111].
    // - In the second operation we can swap the middle bit of the third and fourth elements.
    //     - The resulting array is [001, 110, 001, 110, 001, 110, 111].
    // The array product is 1 * 6 * 1 * 6 * 1 * 6 * 7 = 1512, which is the minimum possible product.
    fmt.Println(minNonZeroProduct(3)) // 1512
    fmt.Println(minNonZeroProduct(8)) // 9253531

    fmt.Println(minNonZeroProduct1(1)) // 1
    fmt.Println(minNonZeroProduct1(2)) // 6
    fmt.Println(minNonZeroProduct1(3)) // 1512
    fmt.Println(minNonZeroProduct1(8)) // 9253531
}