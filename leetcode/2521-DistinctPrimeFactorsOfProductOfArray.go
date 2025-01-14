package main

// 2521. Distinct Prime Factors of Product of Array
// Given an array of positive integers nums, 
// return the number of distinct prime factors in the product of the elements of nums.

// Note that:
//     A number greater than 1 is called prime if it is divisible by only 1 and itself.
//     An integer val1 is a factor of another integer val2 if val2 / val1 is an integer.

// Example 1:
// Input: nums = [2,4,3,7,10,6]
// Output: 4
// Explanation:
// The product of all the elements in nums is: 2 * 4 * 3 * 7 * 10 * 6 = 10080 = 25 * 32 * 5 * 7.
// There are 4 distinct prime factors so we return 4.

// Example 2:
// Input: nums = [2,4,8,16]
// Output: 1
// Explanation:
// The product of all the elements in nums is: 2 * 4 * 8 * 16 = 1024 = 210.
// There is 1 distinct prime factor so we return 1.

// Constraints:
//     1 <= nums.length <= 10^4
//     2 <= nums[i] <= 1000

import "fmt"

func distinctPrimeFactors(nums []int) int {
    mp := make(map[int]int)
    for _, v := range nums {
        for i := 2; v > 1; i++ {
            for v % i == 0 {
                mp[i]++
                v = v / i
            }
        }
    }
    return len(mp)
}

func distinctPrimeFactors1(nums []int) int {
    mp := make(map[int]bool)
    for _, v := range nums {
        for i := 2; i <= v / i; i++ {
            if v % i == 0 {
                mp[i] = true
                for v % i == 0 {
                    v /= i
                }
            }
        }
        if v > 1 { mp[v] = true }
    }
    return len(mp)
}

func main() {
    // Example 1:
    // Input: nums = [2,4,3,7,10,6]
    // Output: 4
    // Explanation:
    // The product of all the elements in nums is: 2 * 4 * 3 * 7 * 10 * 6 = 10080 = 25 * 32 * 5 * 7.
    // There are 4 distinct prime factors so we return 4.
    fmt.Println(distinctPrimeFactors([]int{2,4,3,7,10,6})) // 4
    // Example 2:
    // Input: nums = [2,4,8,16]
    // Output: 1
    // Explanation:
    // The product of all the elements in nums is: 2 * 4 * 8 * 16 = 1024 = 210.
    // There is 1 distinct prime factor so we return 1.
    fmt.Println(distinctPrimeFactors([]int{2,4,8,16})) // 1

    fmt.Println(distinctPrimeFactors([]int{1,2,3,4,5,6,7,8,9})) // 4
    fmt.Println(distinctPrimeFactors([]int{9,8,7,6,5,4,3,2,1})) // 4

    fmt.Println(distinctPrimeFactors1([]int{2,4,3,7,10,6})) // 4
    fmt.Println(distinctPrimeFactors1([]int{2,4,8,16})) // 1
    fmt.Println(distinctPrimeFactors1([]int{1,2,3,4,5,6,7,8,9})) // 4
    fmt.Println(distinctPrimeFactors1([]int{9,8,7,6,5,4,3,2,1})) // 4
}