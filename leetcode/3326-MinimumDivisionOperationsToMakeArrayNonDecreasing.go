package main

// 3326. Minimum Division Operations to Make Array Non Decreasing
// You are given an integer array nums.

// Any positive divisor of a natural number x that is strictly less than x is called a proper divisor of x. 
// For example, 2 is a proper divisor of 4, while 6 is not a proper divisor of 6.

// You are allowed to perform an operation any number of times on nums, 
// where in each operation you select any one element from nums and divide it by its greatest proper divisor.

// Return the minimum number of operations required to make the array non-decreasing.

// If it is not possible to make the array non-decreasing using any number of operations, return -1.

// Example 1:
// Input: nums = [25,7]
// Output: 1
// Explanation:
// Using a single operation, 25 gets divided by 5 and nums becomes [5, 7].

// Example 2:
// Input: nums = [7,7,6]
// Output: -1

// Example 3:
// Input: nums = [1,1,1,1]
// Output: 0

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^6

import "fmt"

const mx = 1_000_001
var facts = [mx]int{}

func init() {
    for i := 2; i < mx; i++ {
        if facts[i] == 0 {
            for j := i; j < mx; j += i {
                if facts[j] == 0 {
                    facts[j] = i
                }
            }
        }
    }
}

func minOperations1(nums []int) int {
    res, n := 0, len(nums)
    for i := n - 2; i >= 0; i-- {
        if nums[i] > nums[i + 1] {
            nums[i] = facts[nums[i]]
            if nums[i] > nums[i + 1] {
                return -1
            }
            res++
        }
    }
    return res
}

// // 解答错误 715 / 720
// func minOperations(nums []int) int {
//     res, n, limit := 0, len(nums), 1000
//     prime, sieve := make([]int, 0), make([]bool, limit + 1)
//     for i := 2; i <= limit; i++ {
//         if sieve[i] { continue }
//         prime = append(prime, i)
//         for j := i; j <= limit; j += i {
//             sieve[j] = true
//         }
//     }
//     for i := n - 2; i >= 0; i-- {
//         if nums[i + 1] >= nums[i] { continue }
//         for j := 0; j < len(prime); j++ {
//             if nums[i] == prime[j] { continue }
//             if nums[i] % prime[j] == 0 && nums[i + 1] >= nums[i] / prime[j] {
//                 nums[i] /= prime[j]
//                 res++
//                 break
//             }
//             if nums[i] % prime[len(prime) - 1 - j] == 0 && nums[i + 1] >= prime[len(prime) - 1 - j] {
//                 nums[i] = prime[len(prime) - 1 - j]
//                 res++
//                 break
//             }
//         }
//     }
//     for i := 1; i < n; i++ {
//         if nums[i] < nums[i-1] { return -1 }
//     }
//     return res
// }

func minOperations(nums []int) int {
    res ,n := 0, len(nums)
    mp := make(map[int]int)
    leastPrime := func(num int) int {
        if num % 2 == 0 { return 2 }
        for prime := 3; prime * prime <= num; prime += 2 {
            if num % prime == 0 { return prime }
        }
        return num
    }
    for i := n - 2; i >= 0; i-- {
        if nums[i] <= nums[i + 1] { continue }
        v, ok := mp[nums[i]]
        if !ok {
            v = leastPrime(nums[i])
            mp[nums[i]] = v
        }
        nums[i] = v
        if nums[i] > nums[i + 1] { return -1 }
        res++
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [25,7]
    // Output: 1
    // Explanation:
    // Using a single operation, 25 gets divided by 5 and nums becomes [5, 7].
    fmt.Println(minOperations([]int{25,7})) // 1
    // Example 2:
    // Input: nums = [7,7,6]
    // Output: -1
    fmt.Println(minOperations([]int{7,7,6})) // -1
    // Example 3:
    // Input: nums = [1,1,1,1]
    // Output: 0
    fmt.Println(minOperations([]int{1,1,1,1})) // 0

    fmt.Println(minOperations([]int{5,51,25})) // -1
    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1})) // -1

    fmt.Println(minOperations1([]int{25,7})) // 1
    fmt.Println(minOperations1([]int{7,7,6})) // -1
    fmt.Println(minOperations1([]int{1,1,1,1})) // 0
    fmt.Println(minOperations1([]int{5,51,25})) // -1
    fmt.Println(minOperations1([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minOperations1([]int{9,8,7,6,5,4,3,2,1})) // -1
}