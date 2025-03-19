package main

// 3191. Minimum Operations to Make Binary Array Elements Equal to One I
// You are given a binary array nums.

// You can do the following operation on the array any number of times (possibly zero):
//     Choose any 3 consecutive elements from the array and flip all of them.

// Flipping an element means changing its value from 0 to 1, and from 1 to 0.

// Return the minimum number of operations required to make all elements in nums equal to 1. 
// If it is impossible, return -1.

// Example 1:
// Input: nums = [0,1,1,1,0,0]
// Output: 3
// Explanation:
// We can do the following operations:
//     Choose the elements at indices 0, 1 and 2. The resulting array is nums = [1,0,0,1,0,0].
//     Choose the elements at indices 1, 2 and 3. The resulting array is nums = [1,1,1,0,0,0].
//     Choose the elements at indices 3, 4 and 5. The resulting array is nums = [1,1,1,1,1,1].

// Example 2:
// Input: nums = [0,1,1,1]
// Output: -1
// Explanation:
// It is impossible to make all elements equal to 1.

// Constraints:
//     3 <= nums.length <= 10^5
//     0 <= nums[i] <= 1

import "fmt"

func minOperations(nums []int) int {
    res, n := 0, len(nums)
    for i := 0; i < n - 2; i++ {
        if nums[i] == 0 { // 每次翻转3个
            res++
            nums[i+1] ^= 1
            nums[i+2] ^= 1
        }
    }
    if nums[n-1] == 0 || nums[n-2] == 0 { return -1 } // 最后一翻转还有 0 的存在说明无法完成操作
    return res
}

func minOperations1(nums []int) int {
    res, n := 0, len(nums)
    for i, x := range nums[:n-2] {
        if x == 0 { // 必须操作
            nums[i+1] ^= 1
            nums[i+2] ^= 1
            res++
        }
    }
    if nums[n-2] == 0 || nums[n-1] == 0 { return -1 }
    return res
}

func minOperations2(nums []int) int {
    count, n := 0, len(nums) - 2
    for i := 0; i < n; i++ {
        if nums[i] != 0 { continue }
        count++
        nums[i] = 1
        nums[i + 1] ^= 1
        nums[i + 2] ^= 1
    }
    if nums[n] + nums[n + 1] != 2 { return -1 }
    return count
}

func minOperations3(nums []int) int {
    res, n := 0, len(nums)
    for i := 0; i < n - 2; i++ {
        if nums[i] == 0 {
            res++
            nums[i] = 1
            nums[i+1] ^= 1
            nums[i+2] ^= 1
        }
    }
    for i := n-2; i < n; i++ {
        if nums[i] != 1 {
            return -1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [0,1,1,1,0,0]
    // Output: 3
    // Explanation:
    // We can do the following operations:
    //     Choose the elements at indices 0, 1 and 2. The resulting array is nums = [1,0,0,1,0,0].
    //     Choose the elements at indices 1, 2 and 3. The resulting array is nums = [1,1,1,0,0,0].
    //     Choose the elements at indices 3, 4 and 5. The resulting array is nums = [1,1,1,1,1,1].
    fmt.Println(minOperations([]int{0,1,1,1,0,0})) // 3
    // Example 2:
    // Input: nums = [0,1,1,1]
    // Output: -1
    // Explanation:
    // It is impossible to make all elements equal to 1.
    fmt.Println(minOperations([]int{0,1,1,1})) // -1
    fmt.Println(minOperations([]int{0,1,0,1,0,1,0,1,0,1,0,1})) // 6
    fmt.Println(minOperations([]int{1,0,1,0,1,0,1,0,1,0,1,0})) // 6
    fmt.Println(minOperations([]int{1,1,1,1,1,0,0,0,0,0})) // -1
    fmt.Println(minOperations([]int{0,0,0,0,0,1,1,1,1,1})) // -1

    fmt.Println(minOperations1([]int{0,1,1,1,0,0})) // 3
    fmt.Println(minOperations1([]int{0,1,1,1})) // -1
    fmt.Println(minOperations1([]int{0,1,0,1,0,1,0,1,0,1,0,1})) // 6
    fmt.Println(minOperations1([]int{1,0,1,0,1,0,1,0,1,0,1,0})) // 6
    fmt.Println(minOperations1([]int{1,1,1,1,1,0,0,0,0,0})) // -1
    fmt.Println(minOperations1([]int{0,0,0,0,0,1,1,1,1,1})) // -1

    fmt.Println(minOperations2([]int{0,1,1,1,0,0})) // 3
    fmt.Println(minOperations2([]int{0,1,1,1})) // -1
    fmt.Println(minOperations2([]int{0,1,0,1,0,1,0,1,0,1,0,1})) // 6
    fmt.Println(minOperations2([]int{1,0,1,0,1,0,1,0,1,0,1,0})) // 6
    fmt.Println(minOperations2([]int{1,1,1,1,1,0,0,0,0,0})) // -1
    fmt.Println(minOperations2([]int{0,0,0,0,0,1,1,1,1,1})) // -1

    fmt.Println(minOperations3([]int{0,1,1,1,0,0})) // 3
    fmt.Println(minOperations3([]int{0,1,1,1})) // -1
    fmt.Println(minOperations3([]int{0,1,0,1,0,1,0,1,0,1,0,1})) // 6
    fmt.Println(minOperations3([]int{1,0,1,0,1,0,1,0,1,0,1,0})) // 6
    fmt.Println(minOperations3([]int{1,1,1,1,1,0,0,0,0,0})) // -1
    fmt.Println(minOperations3([]int{0,0,0,0,0,1,1,1,1,1})) // -1
}