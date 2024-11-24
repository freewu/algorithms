package main

// 1998. GCD Sort of an Array
// You are given an integer array nums, and you can perform the following operation any number of times on nums:
//     Swap the positions of two elements nums[i] and nums[j] if gcd(nums[i], nums[j]) > 1 
//     where gcd(nums[i], nums[j]) is the greatest common divisor of nums[i] and nums[j].

// Return true if it is possible to sort nums in non-decreasing order using the above swap method, or false otherwise.

// Example 1:
// Input: nums = [7,21,3]
// Output: true
// Explanation: We can sort [7,21,3] by performing the following operations:
// - Swap 7 and 21 because gcd(7,21) = 7. nums = [21,7,3]
// - Swap 21 and 3 because gcd(21,3) = 3. nums = [3,7,21]

// Example 2:
// Input: nums = [5,2,6,2]
// Output: false
// Explanation: It is impossible to sort the array because 5 cannot be swapped with any other element.

// Example 3:
// Input: nums = [10,5,9,3,15]
// Output: true
// We can sort [10,5,9,3,15] by performing the following operations:
// - Swap 10 and 15 because gcd(10,15) = 5. nums = [15,5,9,3,10]
// - Swap 15 and 3 because gcd(15,3) = 3. nums = [3,5,9,15,10]
// - Swap 10 and 15 because gcd(10,15) = 5. nums = [3,5,9,10,15]

// Constraints:
//     1 <= nums.length <= 3 * 10^4
//     2 <= nums[i] <= 10^5

import "fmt"
import "sort"

func gcdSort(nums []int) bool {
    n := 100001
    visited, seen, saved := make([]bool, n), make(map[int]bool), make([]int, len(nums))
    for _, v := range nums {
        seen[v] = true
    }
    uf := make([]int, n)
    for i := range uf {
        uf[i] = i
    }
    var find func(v int) int
    find = func(v int) int {
        if uf[v] != v {
            uf[v] = find(uf[v])
        }
        return uf[v]
    }
    for i := 2; i < n; i++ {
        if visited[i] { continue }
        for j := i; j < n; j += i {
            visited[j] = true
            if seen[j] {
                uf[find(j)] = find(i)
            }
        }
    }
    copy(saved, nums)
    sort.Ints(nums)
    for i := range nums {
        if find(nums[i]) != find(saved[i]) {
            return false
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: nums = [7,21,3]
    // Output: true
    // Explanation: We can sort [7,21,3] by performing the following operations:
    // - Swap 7 and 21 because gcd(7,21) = 7. nums = [21,7,3]
    // - Swap 21 and 3 because gcd(21,3) = 3. nums = [3,7,21]
    fmt.Println(gcdSort([]int{7,21,3})) // true
    // Example 2:
    // Input: nums = [5,2,6,2]
    // Output: false
    // Explanation: It is impossible to sort the array because 5 cannot be swapped with any other element.
    fmt.Println(gcdSort([]int{5,2,6,2})) // false
    // Example 3:
    // Input: nums = [10,5,9,3,15]
    // Output: true
    // We can sort [10,5,9,3,15] by performing the following operations:
    // - Swap 10 and 15 because gcd(10,15) = 5. nums = [15,5,9,3,10]
    // - Swap 15 and 3 because gcd(15,3) = 3. nums = [3,5,9,15,10]
    // - Swap 10 and 15 because gcd(10,15) = 5. nums = [3,5,9,10,15]
    fmt.Println(gcdSort([]int{10,5,9,3,15})) // true
}