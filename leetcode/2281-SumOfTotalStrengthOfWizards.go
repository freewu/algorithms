package main

// 2281. Sum of Total Strength of Wizards
// As the ruler of a kingdom, you have an army of wizards at your command.

// You are given a 0-indexed integer array strength, where strength[i] denotes the strength of the ith wizard. 
// For a contiguous group of wizards (i.e. the wizards' strengths form a subarray of strength), the total strength is defined as the product of the following two values:
//     The strength of the weakest wizard in the group.
//     The total of all the individual strengths of the wizards in the group.

// Return the sum of the total strengths of all contiguous groups of wizards.
// Since the answer may be very large, return it modulo 10^9 + 7.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: strength = [1,3,1,2]
// Output: 44
// Explanation: The following are all the contiguous groups of wizards:
// - [1] from [1,3,1,2] has a total strength of min([1]) * sum([1]) = 1 * 1 = 1
// - [3] from [1,3,1,2] has a total strength of min([3]) * sum([3]) = 3 * 3 = 9
// - [1] from [1,3,1,2] has a total strength of min([1]) * sum([1]) = 1 * 1 = 1
// - [2] from [1,3,1,2] has a total strength of min([2]) * sum([2]) = 2 * 2 = 4
// - [1,3] from [1,3,1,2] has a total strength of min([1,3]) * sum([1,3]) = 1 * 4 = 4
// - [3,1] from [1,3,1,2] has a total strength of min([3,1]) * sum([3,1]) = 1 * 4 = 4
// - [1,2] from [1,3,1,2] has a total strength of min([1,2]) * sum([1,2]) = 1 * 3 = 3
// - [1,3,1] from [1,3,1,2] has a total strength of min([1,3,1]) * sum([1,3,1]) = 1 * 5 = 5
// - [3,1,2] from [1,3,1,2] has a total strength of min([3,1,2]) * sum([3,1,2]) = 1 * 6 = 6
// - [1,3,1,2] from [1,3,1,2] has a total strength of min([1,3,1,2]) * sum([1,3,1,2]) = 1 * 7 = 7
// The sum of all the total strengths is 1 + 9 + 1 + 4 + 4 + 4 + 3 + 5 + 6 + 7 = 44.

// Example 2:
// Input: strength = [5,4,6]
// Output: 213
// Explanation: The following are all the contiguous groups of wizards: 
// - [5] from [5,4,6] has a total strength of min([5]) * sum([5]) = 5 * 5 = 25
// - [4] from [5,4,6] has a total strength of min([4]) * sum([4]) = 4 * 4 = 16
// - [6] from [5,4,6] has a total strength of min([6]) * sum([6]) = 6 * 6 = 36
// - [5,4] from [5,4,6] has a total strength of min([5,4]) * sum([5,4]) = 4 * 9 = 36
// - [4,6] from [5,4,6] has a total strength of min([4,6]) * sum([4,6]) = 4 * 10 = 40
// - [5,4,6] from [5,4,6] has a total strength of min([5,4,6]) * sum([5,4,6]) = 4 * 15 = 60
// The sum of all the total strengths is 25 + 16 + 36 + 36 + 40 + 60 = 213.

// Constraints:
//     1 <= strength.length <= 10^5
//     1 <= strength[i] <= 10^9

import "fmt"

// Prefix Sum + Monotonic Stack
func totalStrength(strength []int) int {
    res, mod, n := 0, 1_000_000_007, len(strength)
    prePresum := make([]int, n + 2)
    for i := 0; i < n; i++ { // Calculate presum
        prePresum[i+2] = (prePresum[i+1] + strength[i]) % mod
    }
    for i := 1; i <= n; i++ { // Calculate presum of presum
        prePresum[i+1] = (prePresum[i+1] + prePresum[i])%mod
    }
    right, left := make([]int, n), make([]int, n)
    for i := range right {
        right[i] = n
    }
    for i := range left {
        left[i] = -1
    }
    stack := []int{}
    for i := 0; i < n; i++ { // Get the first index of non-larger value to right of strength[i]
        for len(stack) > 0 && strength[stack[len(stack)-1]] >= strength[i] {
            it := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            right[it] = i
        }
        stack = append(stack, i)
    }
    stack = []int{} // clear stack
    for i := n-1; i >= 0; i-- { // Get the first index of smaller value to left of strength[i]
        for len(stack) > 0 && strength[stack[len(stack)-1]] > strength[i] {
            it := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            left[it] = i
        }
        stack = append(stack, i)
    }
    for i := 0; i < n; i++ { // For each element in strength, we get the value of R_term - L_term.
        leftBound,rightBound := left[i], right[i] // Get the left index and the right index.
        leftCount, rightCount := i - leftBound, rightBound - i // Get the leftCount and rightCount

        negPresum := (prePresum[i+1] - prePresum[i-leftCount+1]) % mod
        posPresum := (prePresum[i + rightCount+1] - prePresum[i+1]) % mod
        
        // The total strength of all subarrays that have strength[i] as the minimum.
        res += (posPresum * leftCount - negPresum * rightCount) % mod * strength[i] % mod
        res %= mod
    }
    return (res + mod) % mod
}

func totalStrength1(strength []int) int {
    res, n, mod := 0, len(strength), 1_000_000_007
    left, right, stack := make([]int, n), make([]int, n), []int{-1}
    for i := range right {
        right[i] = n
    }
    for i, v := range strength {
        for len(stack)>1 && strength[stack[len(stack)-1]] >= v {
            right[stack[len(stack)-1]] = i
            stack = stack[:len(stack)-1]
        }
        left[i] = stack[len(stack)-1]
        stack = append(stack, i)
    }
    s, ss := 0, make([]int, n+2)
    for i, v := range strength {
        s += v 
        ss[i+2] = (ss[i+1] + s) % mod
    }
    for i, v := range strength {
        l, r := left[i]+1, right[i]-1
        total := ((i-l+1) * (ss[r+2]-ss[i+1]) - (r-i+1)*(ss[i+1]-ss[l])) % mod
        res = (res + v * total) % mod
    }
    return (res + mod) % mod
}

func main() {
    // Example 1:
    // Input: strength = [1,3,1,2]
    // Output: 44
    // Explanation: The following are all the contiguous groups of wizards:
    // - [1] from [1,3,1,2] has a total strength of min([1]) * sum([1]) = 1 * 1 = 1
    // - [3] from [1,3,1,2] has a total strength of min([3]) * sum([3]) = 3 * 3 = 9
    // - [1] from [1,3,1,2] has a total strength of min([1]) * sum([1]) = 1 * 1 = 1
    // - [2] from [1,3,1,2] has a total strength of min([2]) * sum([2]) = 2 * 2 = 4
    // - [1,3] from [1,3,1,2] has a total strength of min([1,3]) * sum([1,3]) = 1 * 4 = 4
    // - [3,1] from [1,3,1,2] has a total strength of min([3,1]) * sum([3,1]) = 1 * 4 = 4
    // - [1,2] from [1,3,1,2] has a total strength of min([1,2]) * sum([1,2]) = 1 * 3 = 3
    // - [1,3,1] from [1,3,1,2] has a total strength of min([1,3,1]) * sum([1,3,1]) = 1 * 5 = 5
    // - [3,1,2] from [1,3,1,2] has a total strength of min([3,1,2]) * sum([3,1,2]) = 1 * 6 = 6
    // - [1,3,1,2] from [1,3,1,2] has a total strength of min([1,3,1,2]) * sum([1,3,1,2]) = 1 * 7 = 7
    // The sum of all the total strengths is 1 + 9 + 1 + 4 + 4 + 4 + 3 + 5 + 6 + 7 = 44.
    fmt.Println(totalStrength([]int{1,3,1,2})) // 44
    // Example 2:
    // Input: strength = [5,4,6]
    // Output: 213
    // Explanation: The following are all the contiguous groups of wizards: 
    // - [5] from [5,4,6] has a total strength of min([5]) * sum([5]) = 5 * 5 = 25
    // - [4] from [5,4,6] has a total strength of min([4]) * sum([4]) = 4 * 4 = 16
    // - [6] from [5,4,6] has a total strength of min([6]) * sum([6]) = 6 * 6 = 36
    // - [5,4] from [5,4,6] has a total strength of min([5,4]) * sum([5,4]) = 4 * 9 = 36
    // - [4,6] from [5,4,6] has a total strength of min([4,6]) * sum([4,6]) = 4 * 10 = 40
    // - [5,4,6] from [5,4,6] has a total strength of min([5,4,6]) * sum([5,4,6]) = 4 * 15 = 60
    // The sum of all the total strengths is 25 + 16 + 36 + 36 + 40 + 60 = 213.
    fmt.Println(totalStrength([]int{5,4,6})) // 213

    fmt.Println(totalStrength1([]int{1,3,1,2})) // 44
    fmt.Println(totalStrength1([]int{5,4,6})) // 213
}