package main

// 3533. Concatenated Divisibility
// You are given an array of positive integers nums and a positive integer k.

// A permutation of nums is said to form a divisible concatenation if, 
// when you concatenate the decimal representations of the numbers in the order specified by the permutation, 
// the resulting number is divisible by k.

// Return the lexicographically smallest permutation (when considered as a list of integers) that forms a divisible concatenation. 
// If no such permutation exists, return an empty list.

// Example 1:
// Input: nums = [3,12,45], k = 5
// Output: [3,12,45]
// Explanation:
// Permutation	Concatenated Value	Divisible by 5
// [3, 12, 45]	31245	Yes
// [3, 45, 12]	34512	No
// [12, 3, 45]	12345	Yes
// [12, 45, 3]	12453	No
// [45, 3, 12]	45312	No
// [45, 12, 3]	45123	No
// The lexicographically smallest permutation that forms a divisible concatenation is [3,12,45].

// Example 2:
// Input: nums = [10,5], k = 10
// Output: [5,10]
// Explanation:
// Permutation	Concatenated Value	Divisible by 10
// [5, 10]	510	Yes
// [10, 5]	105	No
// The lexicographically smallest permutation that forms a divisible concatenation is [5,10].

// Example 3:
// Input: nums = [1,2,3], k = 5
// Output: []
// Explanation:
// Since no permutation of nums forms a valid divisible concatenation, return an empty list.

// Constraints:
//     1 <= nums.length <= 13
//     1 <= nums[i] <= 10^5
//     1 <= k <= 100

import "fmt"
import "sort"
import "math"
import "math/bits"
import "slices"
import "strconv"

func concatenatedDivisibility(nums []int, k int) []int {
    sort.Ints(nums)
    n := len(nums)
    dp, seen := make([]int, n), make(map[[3]int]bool)
    var dfs func(index, check, rem int) []int 
    dfs = func(index, check, rem int) []int {
        if index == n {
            if rem == 0 { return dp }
            return nil
        }
        if seen[[3]int{index, check, rem}] { return nil }
        for i := 0; i < len(nums); i++ {
            checki := 1 << i
            if checki & check == 0 {
                dp[index] = nums[i]
                next := ((rem * int(math.Pow10(len(fmt.Sprintf("%d", nums[i])))) % k) + nums[i] % k) % k
                res := dfs(index + 1, check | checki, next)
                if res != nil {
                    return res
                }
            }
        }
        seen[[3]int{index, check, rem}] = true
        return nil
    }
    return dfs(0, 0, 0)
}

func concatenatedDivisibility1(nums []int, k int) (ans []int) {
    slices.Sort(nums)
    n := len(nums)
    res, b10 := []int{}, make([]int, n)
    for i, v := range nums {
        b10[i] = int(math.Pow10(len(strconv.Itoa(v))))
    }
    visited := make([][]bool, 1 << n)
    for i := range visited {
        visited[i] = make([]bool, k)
    }
    var f func(int, int) bool
    f = func(i, x int) bool {
        if i == 1 << n - 1 { return x == 0 }
        if visited[i][x] { return false }
        visited[i][x] = true
        for s := uint(1<<n-1^i); s > 0; s &= s - 1 {
            p := bits.TrailingZeros(s)
            if f(i | 1 << p, (x * b10[p] + nums[p]) % k) {
                res = append(res, nums[p])
                return true
            }
        }
        return false
    }
    if !f(0, 0) { return nil }
    slices.Reverse(res)
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,12,45], k = 5
    // Output: [3,12,45]
    // Explanation:
    // Permutation	Concatenated Value	Divisible by 5
    // [3, 12, 45]	31245	Yes
    // [3, 45, 12]	34512	No
    // [12, 3, 45]	12345	Yes
    // [12, 45, 3]	12453	No
    // [45, 3, 12]	45312	No
    // [45, 12, 3]	45123	No
    // The lexicographically smallest permutation that forms a divisible concatenation is [3,12,45].
    fmt.Println(concatenatedDivisibility([]int{3,12,45}, 5)) // [3,12,45]
    // Example 2:
    // Input: nums = [10,5], k = 10
    // Output: [5,10]
    // Explanation:
    // Permutation	Concatenated Value	Divisible by 10
    // [5, 10]	510	Yes
    // [10, 5]	105	No
    // The lexicographically smallest permutation that forms a divisible concatenation is [5,10].
    fmt.Println(concatenatedDivisibility([]int{10,5}, 10)) // [5,10]
    // Example 3:
    // Input: nums = [1,2,3], k = 5
    // Output: []
    // Explanation:
    // Since no permutation of nums forms a valid divisible concatenation, return an empty list.
    fmt.Println(concatenatedDivisibility([]int{1,2,3}, 5)) // []

    fmt.Println(concatenatedDivisibility([]int{1,2,3,4,5,6,7,8,9}, 5)) // [1 2 3 4 6 7 8 9 5]
    fmt.Println(concatenatedDivisibility([]int{9,8,7,6,5,4,3,2,1}, 5)) // [1 2 3 4 6 7 8 9 5]

    fmt.Println(concatenatedDivisibility1([]int{3,12,45}, 5)) // [3,12,45]
    fmt.Println(concatenatedDivisibility1([]int{10,5}, 10)) // [5,10]
    fmt.Println(concatenatedDivisibility1([]int{1,2,3}, 5)) // []
    fmt.Println(concatenatedDivisibility1([]int{1,2,3,4,5,6,7,8,9}, 5)) // [1 2 3 4 6 7 8 9 5]
    fmt.Println(concatenatedDivisibility1([]int{9,8,7,6,5,4,3,2,1}, 5)) // [1 2 3 4 6 7 8 9 5]
}