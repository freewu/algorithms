package main

// 2449. Minimum Number of Operations to Make Arrays Similar
// You are given two positive integer arrays nums and target, of the same length.

// In one operation, you can choose any two distinct indices i and j where 0 <= i, j < nums.length and:
//     1. set nums[i] = nums[i] + 2 and
//     2. set nums[j] = nums[j] - 2.

// Two arrays are considered to be similar if the frequency of each element is the same.

// Return the minimum number of operations required to make nums similar to target. 
// The test cases are generated such that nums can always be similar to target.

// Example 1:
// Input: nums = [8,12,6], target = [2,14,10]
// Output: 2
// Explanation: It is possible to make nums similar to target in two operations:
// - Choose i = 0 and j = 2, nums = [10,12,4].
// - Choose i = 1 and j = 2, nums = [10,14,2].
// It can be shown that 2 is the minimum number of operations needed.

// Example 2:
// Input: nums = [1,2,5], target = [4,1,3]
// Output: 1
// Explanation: We can make nums similar to target in one operation:
// - Choose i = 1 and j = 2, nums = [1,4,3].

// Example 3:
// Input: nums = [1,1,1,1,1], target = [1,1,1,1,1]
// Output: 0
// Explanation: The array nums is already similiar to target.

// Constraints:
//     n == nums.length == target.length
//     1 <= n <= 10^5
//     1 <= nums[i], target[i] <= 10^6
//     It is possible to make nums similar to target.

import "fmt"
import "sort"

func makeSimilar(nums []int, target []int) int64 {
    even, odd, teven, todd := []int{}, []int{}, []int{}, []int{}
    for i := 0; i < len(nums); i++ {
        if nums[i] % 2 == 0 {
            even = append(even, nums[i])
        } else {
            odd = append(odd, nums[i])
        }
        if target[i] % 2 == 0 {
            teven = append(teven, target[i])
        } else {
            todd = append(todd, target[i])
        }
    }
    sort.Ints(even)
    sort.Ints(odd)
    sort.Ints(teven)
    sort.Ints(todd)
    res := 0
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 0; i < len(even); i++ {
        res += abs(even[i] - teven[i]) / 2
    }
    for i := 0; i < len(odd); i++ {
        res += abs(odd[i] - todd[i]) / 2
    }
    return int64(res / 2)
}

func makeSimilar1(nums []int, target []int) int64 {
    helper := func(arr []int) {
        for i, v := range arr {
            if v % 2 > 0 {
                arr[i] = -v
            }
        }
        sort.Ints(arr)
    }
    helper(nums)
    helper(target)
    res := 0
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i, v := range nums {
        res += abs(v - target[i])
    }
    return int64(res / 4)
}

func main() {
    // Example 1:
    // Input: nums = [8,12,6], target = [2,14,10]
    // Output: 2
    // Explanation: It is possible to make nums similar to target in two operations:
    // - Choose i = 0 and j = 2, nums = [10,12,4].
    // - Choose i = 1 and j = 2, nums = [10,14,2].
    // It can be shown that 2 is the minimum number of operations needed.
    fmt.Println(makeSimilar([]int{8,12,6}, []int{2,14,10})) // 2
    // Example 2:
    // Input: nums = [1,2,5], target = [4,1,3]
    // Output: 1
    // Explanation: We can make nums similar to target in one operation:
    // - Choose i = 1 and j = 2, nums = [1,4,3].
    fmt.Println(makeSimilar([]int{1,2,5}, []int{4,1,3})) // 1
    // Example 3:
    // Input: nums = [1,1,1,1,1], target = [1,1,1,1,1]
    // Output: 0
    // Explanation: The array nums is already similiar to target.
    fmt.Println(makeSimilar([]int{1,1,1,1,1}, []int{1,1,1,1,1})) // 0

    fmt.Println(makeSimilar([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 0
    fmt.Println(makeSimilar([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(makeSimilar([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(makeSimilar([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(makeSimilar1([]int{8,12,6}, []int{2,14,10})) // 2
    fmt.Println(makeSimilar1([]int{1,2,5}, []int{4,1,3})) // 1
    fmt.Println(makeSimilar1([]int{1,1,1,1,1}, []int{1,1,1,1,1})) // 0
    fmt.Println(makeSimilar1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 0
    fmt.Println(makeSimilar1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(makeSimilar1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(makeSimilar1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 0
}