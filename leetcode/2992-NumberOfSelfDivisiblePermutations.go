package main

// 2992. Number of Self-Divisible Permutations
// Given an integer n, return the number of permutations of the 1-indexed array nums = [1, 2, ..., n], 
// such that it's self-divisible.

// A 1-indexed array a of length n is self-divisible if for every 1 <= i <= n, gcd(a[i], i) == 1.

// A permutation of an array is a rearrangement of the elements of that array, 
// for example here are all of the permutations of the array [1, 2, 3]:
//     [1, 2, 3]
//     [1, 3, 2]
//     [2, 1, 3]
//     [2, 3, 1]
//     [3, 1, 2]
//     [3, 2, 1]

// Example 1:
// Input: n = 1
// Output: 1
// Explanation: The array [1] has only 1 permutation which is self-divisible.

// Example 2:
// Input: n = 2
// Output: 1
// Explanation: The array [1,2] has 2 permutations and only one of them is self-divisible:
// nums = [1,2]: This is not self-divisible since gcd(nums[2], 2) != 1.
// nums = [2,1]: This is self-divisible since gcd(nums[1], 1) == 1 and gcd(nums[2], 2) == 1.

// Example 3:
// Input: n = 3
// Output: 3
// Explanation: The array [1,2,3] has 3 self-divisble permutations: [1,3,2], [3,1,2], [2,3,1].
// It can be shown that the other 3 permutations are not self-divisible. Hence the answer is 3.

// Constraints:
//     1 <= n <= 12

import "fmt"
import "math/bits"

func selfDivisiblePermutationCount(n int) int {
    var gcd func(a, b int) int
    gcd = func(a, b int) int {
        mx, mn := a, b
        if mx < mn {
            mx, mn = b, a
        }
        k := mx % mn
        if k == 0 { return mn }
        return gcd(k, mn)
    }
    var dfs func(visited []bool, index, count int) int
    dfs = func(visited []bool, index, count int) int {
        if index >= len(visited) {
            return count + 1
        }
        for i := 0; i < len(visited); i++ {
            if !visited[i] && gcd(i + 1, index + 1) == 1 {
                visited[i] = true
                count = dfs(visited, index + 1, count)
                visited[i] = false
            }
        }
        return count
    }
    return dfs(make([]bool, n), 0, 0)
}

func selfDivisiblePermutationCount1(n int) int {
    memo := make([]int, 1 << n)
    for i := range memo {
        memo[i] = -1
    }
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    var dfs func(int) int
    dfs = func(mask int) int {
        index := bits.OnesCount(uint(mask))
        if index == n { return 1 }
        if memo[mask] != -1 { return memo[mask] }
        memo[mask] = 0
        for j := 0; j < n; j++ {
            if 1 << j & mask == 0 && gcd(j+1, index+1) == 1 {
                memo[mask] += dfs(1 << j | mask)
            }
        }
        return memo[mask]
    }
    return dfs(0)
}

// table trick
func selfDivisiblePermutationCount2(n int) int {
    table := []int{0,1,1,3,4,28,16,256,324,3600,3600,129744,63504}
    return table[n]
}

func main() {
    // Example 1:
    // Input: n = 1
    // Output: 1
    // Explanation: The array [1] has only 1 permutation which is self-divisible.
    fmt.Println(selfDivisiblePermutationCount(1)) // 1
    // Example 2:
    // Input: n = 2
    // Output: 1
    // Explanation: The array [1,2] has 2 permutations and only one of them is self-divisible:
    // nums = [1,2]: This is not self-divisible since gcd(nums[2], 2) != 1.
    // nums = [2,1]: This is self-divisible since gcd(nums[1], 1) == 1 and gcd(nums[2], 2) == 1.
    fmt.Println(selfDivisiblePermutationCount(2)) // 1
    // Example 3:
    // Input: n = 3
    // Output: 3
    // Explanation: The array [1,2,3] has 3 self-divisble permutations: [1,3,2], [3,1,2], [2,3,1].
    // It can be shown that the other 3 permutations are not self-divisible. Hence the answer is 3.
    fmt.Println(selfDivisiblePermutationCount(3)) // 3

    fmt.Println(selfDivisiblePermutationCount(4)) // 4
    fmt.Println(selfDivisiblePermutationCount(5)) // 28
    fmt.Println(selfDivisiblePermutationCount(6)) // 16
    fmt.Println(selfDivisiblePermutationCount(7)) // 256
    fmt.Println(selfDivisiblePermutationCount(8)) // 324
    fmt.Println(selfDivisiblePermutationCount(9)) // 3600
    fmt.Println(selfDivisiblePermutationCount(10)) // 3600
    fmt.Println(selfDivisiblePermutationCount(11)) // 129744
    fmt.Println(selfDivisiblePermutationCount(12)) // 63504

    fmt.Println(selfDivisiblePermutationCount1(1)) // 1
    fmt.Println(selfDivisiblePermutationCount1(2)) // 1
    fmt.Println(selfDivisiblePermutationCount1(3)) // 3
    fmt.Println(selfDivisiblePermutationCount1(4)) // 4
    fmt.Println(selfDivisiblePermutationCount1(5)) // 28
    fmt.Println(selfDivisiblePermutationCount1(6)) // 16
    fmt.Println(selfDivisiblePermutationCount1(7)) // 256
    fmt.Println(selfDivisiblePermutationCount1(8)) // 324
    fmt.Println(selfDivisiblePermutationCount1(9)) // 3600
    fmt.Println(selfDivisiblePermutationCount1(10)) // 3600
    fmt.Println(selfDivisiblePermutationCount1(11)) // 129744
    fmt.Println(selfDivisiblePermutationCount1(12)) // 63504

    fmt.Println(selfDivisiblePermutationCount2(1)) // 1
    fmt.Println(selfDivisiblePermutationCount2(2)) // 1
    fmt.Println(selfDivisiblePermutationCount2(3)) // 3
    fmt.Println(selfDivisiblePermutationCount2(4)) // 4
    fmt.Println(selfDivisiblePermutationCount2(5)) // 28
    fmt.Println(selfDivisiblePermutationCount2(6)) // 16
    fmt.Println(selfDivisiblePermutationCount2(7)) // 256
    fmt.Println(selfDivisiblePermutationCount2(8)) // 324
    fmt.Println(selfDivisiblePermutationCount2(9)) // 3600
    fmt.Println(selfDivisiblePermutationCount2(10)) // 3600
    fmt.Println(selfDivisiblePermutationCount2(11)) // 129744
    fmt.Println(selfDivisiblePermutationCount2(12)) // 63504
}