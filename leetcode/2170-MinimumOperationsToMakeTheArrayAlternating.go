package main

// 2170. Minimum Operations to Make the Array Alternating
// You are given a 0-indexed array nums consisting of n positive integers.

// The array nums is called alternating if:
//     nums[i - 2] == nums[i], where 2 <= i <= n - 1.
//     nums[i - 1] != nums[i], where 1 <= i <= n - 1.

// In one operation, you can choose an index i and change nums[i] into any positive integer.

// Return the minimum number of operations required to make the array alternating.

// Example 1:
// Input: nums = [3,1,3,2,4,3]
// Output: 3
// Explanation:
// One way to make the array alternating is by converting it to [3,1,3,1,3,1].
// The number of operations required in this case is 3.
// It can be proven that it is not possible to make the array alternating in less than 3 operations. 

// Example 2:
// Input: nums = [1,2,2,2,2]
// Output: 2
// Explanation:
// One way to make the array alternating is by converting it to [1,2,1,2,1].
// The number of operations required in this case is 2.
// Note that the array cannot be converted to [2,2,2,2,2] because in this case nums[0] == nums[1] which violates the conditions of an alternating array.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"

func minimumOperations(nums []int) int {
    n := len(nums)
    get := func(i int) [][]int {
        freq := make(map[int]int)
        for ; i < n; i += 2 {
            freq[nums[i]]++
        }
        a, n1, b, n2 := 0, 0, 0, 0
        for k, v := range freq {
            if v > n1 {
                b, n2, a, n1 = a, n1, k, v
            } else if v > n2 {
                b, n2 = k, v
            }
        }
        return [][]int{ { a, n1 }, { b, n2} }
    }
    res := 1 << 31
    for _, e1 := range get(0) {
        for _, e2 := range get(1) {
            if e1[0] != e2[0] && res > (n - (e1[1]+e2[1])) {
                res = n - (e1[1] + e2[1])
            }
        }
    }
    return res
}

func minimumOperations1(nums []int) int {
    n := len(nums)
    oddCount, evenCount := make(map[int]int, n), make(map[int]int, n) // Count the top 2 most prevalent numbers on even and odd positions
    for i, v := range nums {
        if i % 2 == 0 {
            evenCount[v]++
        } else {
            oddCount[v]++
        }
    }
    var maxOddCount, maxOddNums, maxEvenCounts, maxEvenNums [2]int
    for num, count := range oddCount {
        if count > maxOddCount[0] {
            maxOddNums[0], maxOddNums[1], maxOddCount[1] = num, maxOddNums[0], maxOddCount[0]
            maxOddCount[0] = count
        } else if count > maxOddCount[1] {
            maxOddNums[1], maxOddCount[1] = num, count
        }
    }
    for num, count := range evenCount {
        if count > maxEvenCounts[0] {
            maxEvenNums[0], maxEvenNums[1], maxEvenCounts[1] = num, maxEvenNums[0], maxEvenCounts[0]
            maxEvenCounts[0] = count
        } else if count > maxEvenCounts[1] {
            maxEvenNums[1], maxEvenCounts[1] = num, count
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    if maxEvenNums[0] == maxOddNums[0] {
        return min(n - maxEvenCounts[0] - maxOddCount[1], n - maxEvenCounts[1] - maxOddCount[0] )
    }
    return n - maxEvenCounts[0] - maxOddCount[0]
}

func main() {
    // Example 1:
    // Input: nums = [3,1,3,2,4,3]
    // Output: 3
    // Explanation:
    // One way to make the array alternating is by converting it to [3,1,3,1,3,1].
    // The number of operations required in this case is 3.
    // It can be proven that it is not possible to make the array alternating in less than 3 operations. 
    fmt.Println(minimumOperations([]int{3,1,3,2,4,3})) // 3
    // Example 2:
    // Input: nums = [1,2,2,2,2]
    // Output: 2
    // Explanation:
    // One way to make the array alternating is by converting it to [1,2,1,2,1].
    // The number of operations required in this case is 2.
    // Note that the array cannot be converted to [2,2,2,2,2] because in this case nums[0] == nums[1] which violates the conditions of an alternating array.
    fmt.Println(minimumOperations([]int{1,2,2,2,2})) // 2

    fmt.Println(minimumOperations1([]int{3,1,3,2,4,3})) // 3
    fmt.Println(minimumOperations1([]int{1,2,2,2,2})) // 2
}