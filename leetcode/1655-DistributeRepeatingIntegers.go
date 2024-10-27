package main

// 1655. Distribute Repeating Integers
// You are given an array of n integers, nums, where there are at most 50 unique values in the array. 
// You are also given an array of m customer order quantities, quantity, where quantity[i] is the amount of integers the ith customer ordered. 
// Determine if it is possible to distribute nums such that:
//     The ith customer gets exactly quantity[i] integers,
//     The integers the ith customer gets are all equal, and
//     Every customer is satisfied.

// Return true if it is possible to distribute nums according to the above conditions.

// Example 1:
// Input: nums = [1,2,3,4], quantity = [2]
// Output: false
// Explanation: The 0th customer cannot be given two different integers.

// Example 2:
// Input: nums = [1,2,3,3], quantity = [2]
// Output: true
// Explanation: The 0th customer is given [3,3]. The integers [1,2] are not used.

// Example 3:
// Input: nums = [1,1,2,2], quantity = [2,2]
// Output: true
// Explanation: The 0th customer is given [1,1], and the 1st customer is given [2,2].

// Constraints:
//     n == nums.length
//     1 <= n <= 10^5
//     1 <= nums[i] <= 1000
//     m == quantity.length
//     1 <= m <= 10
//     1 <= quantity[i] <= 10^5
//     There are at most 50 unique values in nums.

import "fmt"
import "sort"

// // 超出时间限制 106 / 109 
// func canDistribute(nums []int, quantity []int) bool {
//     mp := make(map[int]int)
//     for _, v := range nums {
//         mp[v]++
//     }
//     var dfs func(mp map[int]int, quantity []int) bool
//     dfs = func(mp map[int]int, quantity []int) bool {
//         if len(quantity) == 0 { return true }
//         visited := make(map[int]bool)
//         for i := range mp {
//             if visited[mp[i]] { continue } // Just to make sure that same frequency is not visited again
//             visited[mp[i]] = true
//             if mp[i] >= quantity[0] {
//                 mp[i] -= quantity[0]
//                 if dfs(mp, quantity[1:]) { return true }
//                 mp[i] += quantity[0]
//             }
//         }
//         return false
//     }
//     return dfs(mp, quantity)
// }

func canDistribute(nums []int, quantity []int) bool {
    n, sum := len(nums), 0
    for _, v := range quantity { 
        sum += v
    }
    if sum > n { return false }
    mp := make(map[int]int)
    for _, v := range nums { 
        mp[v]++
    }
    sort.Ints(quantity)
    var backtrack func(mp map[int]int, quantity []int, index int) bool 
    backtrack = func(mp map[int]int, quantity []int, index int) bool {
        if index == -1 { return true }
        flag, curr := false, quantity[index]
        for k, v := range mp {
            if v >= curr {
                mp[k] = v - curr
                flag = backtrack(mp, quantity, index - 1)
                if flag { break }
                mp[k] = v
            }
        }
        return flag
    }
    return backtrack(mp, quantity, len(quantity) - 1)
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4], quantity = [2]
    // Output: false
    // Explanation: The 0th customer cannot be given two different integers.
    fmt.Println(canDistribute([]int{1,2,3,4}, []int{2})) // false
    // Example 2:
    // Input: nums = [1,2,3,3], quantity = [2]
    // Output: true
    // Explanation: The 0th customer is given [3,3]. The integers [1,2] are not used.
    fmt.Println(canDistribute([]int{1,2,3,3}, []int{2})) // true
    // Example 3:
    // Input: nums = [1,1,2,2], quantity = [2,2]
    // Output: true
    // Explanation: The 0th customer is given [1,1], and the 1st customer is given [2,2].
    fmt.Println(canDistribute([]int{1,1,2,2}, []int{2,2})) // true
}