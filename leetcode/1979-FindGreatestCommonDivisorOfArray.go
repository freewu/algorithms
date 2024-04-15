package main

// 1979. Find Greatest Common Divisor of Array
// Given an integer array nums, return the greatest common divisor of the smallest number and largest number in nums.
// The greatest common divisor of two numbers is the largest positive integer that evenly divides both numbers.

// Example 1:
// Input: nums = [2,5,6,9,10]
// Output: 2
// Explanation:
// The smallest number in nums is 2.
// The largest number in nums is 10.
// The greatest common divisor of 2 and 10 is 2.

// Example 2:
// Input: nums = [7,5,6,8,3]
// Output: 1
// Explanation:
// The smallest number in nums is 3.
// The largest number in nums is 8.
// The greatest common divisor of 3 and 8 is 1.

// Example 3:
// Input: nums = [3,3]
// Output: 3
// Explanation:
// The smallest number in nums is 3.
// The largest number in nums is 3.
// The greatest common divisor of 3 and 3 is 3.
 
// Constraints:
//     2 <= nums.length <= 1000
//     1 <= nums[i] <= 1000

import "fmt"
import "sort"

func findGCD(nums []int) int {
    sort.Ints(nums)
    // 朴素的欧几里德原理有 gcd(a,b)=gcd(b,a mod b); 辗转相除法 递归
    var gcd func(a int, b int) int 
    gcd = func(a int, b int) int { 
        if a == 0 { 
            return b
        }
        return gcd(b % a, a)
    }
    return gcd(nums[0], nums[len(nums) - 1])
}

func findGCD1(nums []int) int {
    sort.Ints(nums)
    // 朴素的欧几里德原理有 gcd(a,b) = gcd(b,a mod b); 辗转相除法 迭代
    gcd := func(a int, b int) int { 
        for b != 0 { // 为什么用b判断呢？因为b就是用来存a%b的，即上面算法步骤里的r的
            t := b
            b = a % b
            a = t
        }
        return a
    }
    return gcd(nums[0], nums[len(nums) - 1])
}

// # 辗转相减法
// 算法步骤：
//     若a > b，则a = a - b
//     若b > a，则b = b - a
//     若a == b，则a(或b)即为最大公约数
//     若a != b，则回到1

// 求32,12的最大公约数：
//     32 - 12 = 20 (20 > 12)
//     20 - 12 = 8 (8 < 12)
//     12 - 8 = 4 (4 < 8)
//     8 - 4 = 4 (4 == 4)
//     所以最大公约数是4.
func findGCD2(nums []int) int {
    sort.Ints(nums)
    // 辗转相减法 递归实现
    var gcd func(a int, b int) int 
    gcd = func(a int, b int) int { 
        if a == b {
            return a
        }
        if a > b {
            return gcd(a - b, b)
        } 
        return gcd(a, b - a)
    }
    return gcd(nums[0], nums[len(nums) - 1])
}

func findGCD3(nums []int) int {
    sort.Ints(nums)
    // 辗转相减法 迭代实现
    gcd := func(a int, b int) int { 
        for a != b { // 如果a,b不相等，则用大的数减去小的数，直到相等为止
            if a > b {
                a = a - b
            } else {
                b = b - a
            }
        }
        return a
    }
    return gcd(nums[0], nums[len(nums) - 1])
}


func findGCD4(nums []int) int {
    sort.Ints(nums)
    // 穷举法
    gcd := func(a int, b int) int { 
        min := func (x, y int) int { if x < y { return x; }; return y; }
        mn := min(a, b)  // 先找出 a, b 中小的那个
        for i := 2; i <= mn ; i++ {
            if (a % i == 0) && (b % i == 0) { // 如果能同时被 a 和 b 除尽，则更新 res
                return i
            }
        }
        return 1
    }
    return gcd(nums[0], nums[len(nums) - 1])
}

func findGCD5(nums []int) int {
    mn, mx := 1 << 32 - 1, -1 << 32 - 1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // 取到最大&最小值
    for _, v := range nums {
        mn = min(v, mn)
        mx = max(v, mx)
    }
    // 朴素的欧几里德原理有 gcd(a,b)=gcd(b,a mod b); 辗转相除法 递归
    var gcd func(a int, b int) int 
    gcd = func(a int, b int) int { 
        if a == 0 { 
            return b
        }
        return gcd(b % a, a)
    }
    return gcd(mn, mx)
}

func main() {
    // Explanation:
    // The smallest number in nums is 2.
    // The largest number in nums is 10.
    // The greatest common divisor of 2 and 10 is 2.
    fmt.Println(findGCD([]int{2,5,6,9,10})) // 2
    // Explanation:
    // The smallest number in nums is 3.
    // The largest number in nums is 8.
    // The greatest common divisor of 3 and 8 is 1.
    fmt.Println(findGCD([]int{7,5,6,8,3})) // 1
    // The smallest number in nums is 3.
    // The largest number in nums is 3.
    // The greatest common divisor of 3 and 3 is 3.
    fmt.Println(findGCD([]int{3,3})) // 3

    fmt.Println(findGCD1([]int{2,5,6,9,10})) // 2
    fmt.Println(findGCD1([]int{7,5,6,8,3})) // 1
    fmt.Println(findGCD1([]int{3,3})) // 3

    fmt.Println(findGCD2([]int{2,5,6,9,10})) // 2
    fmt.Println(findGCD2([]int{7,5,6,8,3})) // 1
    fmt.Println(findGCD2([]int{3,3})) // 3

    fmt.Println(findGCD3([]int{2,5,6,9,10})) // 2
    fmt.Println(findGCD3([]int{7,5,6,8,3})) // 1
    fmt.Println(findGCD3([]int{3,3})) // 3

    fmt.Println(findGCD4([]int{2,5,6,9,10})) // 2
    fmt.Println(findGCD4([]int{7,5,6,8,3})) // 1
    fmt.Println(findGCD4([]int{3,3})) // 3

    fmt.Println(findGCD5([]int{2,5,6,9,10})) // 2
    fmt.Println(findGCD5([]int{7,5,6,8,3})) // 1
    fmt.Println(findGCD5([]int{3,3})) // 3
}