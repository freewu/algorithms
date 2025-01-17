package main

// 2513. Minimize the Maximum of Two Arrays
// We have two arrays arr1 and arr2 which are initially empty. 
// You need to add positive integers to them such that they satisfy all the following conditions:
//     1. arr1 contains uniqueCnt1 distinct positive integers, each of which is not divisible by divisor1.
//     2. arr2 contains uniqueCnt2 distinct positive integers, each of which is not divisible by divisor2.
//     3. No integer is present in both arr1 and arr2.

// Given divisor1, divisor2, uniqueCnt1, and uniqueCnt2, 
// return the minimum possible maximum integer that can be present in either array.

// Example 1:
// Input: divisor1 = 2, divisor2 = 7, uniqueCnt1 = 1, uniqueCnt2 = 3
// Output: 4
// Explanation: 
// We can distribute the first 4 natural numbers into arr1 and arr2.
// arr1 = [1] and arr2 = [2,3,4].
// We can see that both arrays satisfy all the conditions.
// Since the maximum value is 4, we return it.

// Example 2:
// Input: divisor1 = 3, divisor2 = 5, uniqueCnt1 = 2, uniqueCnt2 = 1
// Output: 3
// Explanation: 
// Here arr1 = [1,2], and arr2 = [3] satisfy all conditions.
// Since the maximum value is 3, we return it.

// Example 3:
// Input: divisor1 = 2, divisor2 = 4, uniqueCnt1 = 8, uniqueCnt2 = 2
// Output: 15
// Explanation: 
// Here, the final possible arrays can be arr1 = [1,3,5,7,9,11,13,15], and arr2 = [2,6].
// It can be shown that it is not possible to obtain a lower maximum satisfying all conditions. 

// Constraints:
//     2 <= divisor1, divisor2 <= 10^5
//     1 <= uniqueCnt1, uniqueCnt2 < 10^9
//     2 <= uniqueCnt1 + uniqueCnt2 <= 10^9

import "fmt"

func minimizeSet(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {
    gcd := func(x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    lcm := func(x int, y int) int { return (x * y) / gcd(x, y) }
    left, right, divisor := 1, 10_000_000_000, lcm(divisor1, divisor2)
    for left < right {
        mid := (left + right) >> 1
        count1, count2 := mid / divisor1 * (divisor1 - 1) + mid % divisor1, mid / divisor2 * (divisor2 - 1) + mid % divisor2
        count := mid / divisor * (divisor - 1) + mid % divisor
        if count1 >= uniqueCnt1 && count2 >= uniqueCnt2 && count >= uniqueCnt1 + uniqueCnt2 {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}

func minimizeSet1(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {
    gcd := func(x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    lcm := func(x int, y int) int { return (x * y) / gcd(x, y) }
    g := lcm(divisor1, divisor2)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    check := func(k int) bool {
        sum := k - k / divisor1 - k / divisor2 + k / g // for 1 and 2
        a := k / divisor1 - k / g // only for 2
        b := k / divisor2 - k / g // only for 1
        return sum >= max(uniqueCnt1 - b, 0) + max(uniqueCnt2 - a, 0)
    }
    left, right := 0, int(2e9)
    for left + 1 < right {
        mid := left + ((right - left) >> 1)
        if check(mid) {
            right = mid
        } else {
            left = mid
        }
    }
    return right
}

func main() {
    // Example 1:
    // Input: divisor1 = 2, divisor2 = 7, uniqueCnt1 = 1, uniqueCnt2 = 3
    // Output: 4
    // Explanation: 
    // We can distribute the first 4 natural numbers into arr1 and arr2.
    // arr1 = [1] and arr2 = [2,3,4].
    // We can see that both arrays satisfy all the conditions.
    // Since the maximum value is 4, we return it.
    fmt.Println(minimizeSet(2,7,1,3)) // 4
    // Example 2:
    // Input: divisor1 = 3, divisor2 = 5, uniqueCnt1 = 2, uniqueCnt2 = 1
    // Output: 3
    // Explanation: 
    // Here arr1 = [1,2], and arr2 = [3] satisfy all conditions.
    // Since the maximum value is 3, we return it.
    fmt.Println(minimizeSet(3,5,2,1)) // 3
    // Example 3:
    // Input: divisor1 = 2, divisor2 = 4, uniqueCnt1 = 8, uniqueCnt2 = 2
    // Output: 15
    // Explanation: 
    // Here, the final possible arrays can be arr1 = [1,3,5,7,9,11,13,15], and arr2 = [2,6].
    // It can be shown that it is not possible to obtain a lower maximum satisfying all conditions.
    fmt.Println(minimizeSet(2,4,8,1)) // 15

    fmt.Println(minimizeSet(2,2,1,1)) // 3
    fmt.Println(minimizeSet(2,2,500_000_000,500_000_000)) // 1999999999
    fmt.Println(minimizeSet(2,2,1,999_999_999)) // 1999999999
    fmt.Println(minimizeSet(2,2,999_999_999,1)) // 1999999999
    fmt.Println(minimizeSet(2,100_000,500_000_000,500_000_000)) // 1000010000
    fmt.Println(minimizeSet(100_000,2,500_000_000,500_000_000)) // 1000010000
    fmt.Println(minimizeSet(100_000,100_000,1,1)) // 2
    fmt.Println(minimizeSet(100_000,100_000,500_000_000,500_000_000)) // 1000010000

    fmt.Println(minimizeSet1(2,7,1,3)) // 4
    fmt.Println(minimizeSet1(3,5,2,1)) // 3
    fmt.Println(minimizeSet1(2,4,8,1)) // 15
    fmt.Println(minimizeSet1(2,2,1,1)) // 3
    fmt.Println(minimizeSet1(2,2,500_000_000,500_000_000)) // 1999999999
    fmt.Println(minimizeSet1(2,2,1,999_999_999)) // 1999999999
    fmt.Println(minimizeSet1(2,2,999_999_999,1)) // 1999999999
    fmt.Println(minimizeSet1(2,100_000,500_000_000,500_000_000)) // 1000010000
    fmt.Println(minimizeSet1(100_000,2,500_000_000,500_000_000)) // 1000010000
    fmt.Println(minimizeSet1(100_000,100_000,1,1)) // 2
    fmt.Println(minimizeSet1(100_000,100_000,500_000_000,500_000_000)) // 1000010000
}