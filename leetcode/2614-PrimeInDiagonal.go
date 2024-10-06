package main

// 2614. Prime In Diagonal
// You are given a 0-indexed two-dimensional integer array nums.

// Return the largest prime number that lies on at least one of the diagonals of nums. 
// In case, no prime is present on any of the diagonals, return 0.

// Note that:
//     1. An integer is prime if it is greater than 1 and has no positive integer divisors other than 1 and itself.
//     2. An integer val is on one of the diagonals of nums if there exists an integer i for which nums[i][i] = val 
//        or an i for which nums[i][nums.length - i - 1] = val.

// In the above diagram, one diagonal is [1,5,9] and another diagonal is [3,5,7].

// Example 1:
// Input: nums = [[1,2,3],[5,6,7],[9,10,11]]
// Output: 11
// Explanation: The numbers 1, 3, 6, 9, and 11 are the only numbers present on at least one of the diagonals. Since 11 is the largest prime, we return 11.

// Example 2:
// Input: nums = [[1,2,3],[5,17,7],[9,11,10]]
// Output: 17
// Explanation: The numbers 1, 3, 9, 10, and 17 are all present on at least one of the diagonals. 17 is the largest prime, so we return 17.

// Constraints:
//     1 <= nums.length <= 300
//     nums.length == numsi.length
//     1 <= nums[i][j] <= 4*10^6

import "fmt"

func diagonalPrime(nums [][]int) int {
    res, n := 0, len(nums)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    isPrime := func(n int) bool {
        if n <= 1 { return false }
        if n == 2 || n == 3 { return true }
        if n % 2 == 0 || n % 3 == 0 { return false }
        for i := 5; i*i <= n; i += 6 {
            if n % i == 0 || n % (i+2) == 0 { return false }
        }
        return true
    }
    for i := 0; i < n; i++ {
        if isPrime(nums[i][i]) { res = max(res, nums[i][i]) }
        if isPrime(nums[i][n-i-1]) { res = max(res, nums[i][n-i-1]) }
    }
    return res 
}

func main() {
    // Example 1:
    // Input: nums = [[1,2,3],[5,6,7],[9,10,11]]
    // Output: 11
    // Explanation: The numbers 1, 3, 6, 9, and 11 are the only numbers present on at least one of the diagonals. Since 11 is the largest prime, we return 11.
    fmt.Println(diagonalPrime([][]int{{1,2,3},{5,6,7},{9,10,11}})) // 11
    // Example 2:
    // Input: nums = [[1,2,3],[5,17,7],[9,11,10]]
    // Output: 17
    // Explanation: The numbers 1, 3, 9, 10, and 17 are all present on at least one of the diagonals. 17 is the largest prime, so we return 17.
    fmt.Println(diagonalPrime([][]int{{1,2,3},{5,17,7},{9,11,10}})) // 17
}