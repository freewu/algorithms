package main

// 1658. Minimum Operations to Reduce X to Zero
// You are given an integer array nums and an integer x. 
// In one operation, you can either remove the leftmost or the rightmost element from the array nums and subtract its value from x. 
// Note that this modifies the array for future operations.

// Return the minimum number of operations to reduce x to exactly 0 if it is possible, otherwise, return -1.

// Example 1:
// Input: nums = [1,1,4,2,3], x = 5
// Output: 2
// Explanation: The optimal solution is to remove the last two elements to reduce x to zero.

// Example 2:
// Input: nums = [5,6,7,8,9], x = 4
// Output: -1

// Example 3:
// Input: nums = [3,2,20,1,1,3], x = 10
// Output: 5
// Explanation: The optimal solution is to remove the last three elements and the first two elements (5 operations in total) to reduce x to zero.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^4
//     1 <= x <= 10^9

import "fmt"

// Sliding Window 
func minOperations(nums []int, x int) int {
    i, j, count, res, flag := 0, len(nums) - 1, 0, len(nums) - 1, false
    for i <= j && nums[i] <= x { // find the left bound
        x -= nums[i]
        i++
        count++
    }
    if i > j { // means the sum of all elems is <= x
        if x == 0 { return count } 
        return -1
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    if x == 0 { // means x can be minused to 0, set temp min Opt nums
        flag, res = true, min(res, count)
    }
    for i >= 0 { // find the right bound
        if nums[j] > x && i > 0 { // x can not be minused, slide the left bound
            i--
            x += nums[i]
            count--
        }
        if nums[j] > x { // left bound can not be slided, because is left most
            if i == 0 { break }
            continue
        }
        x -= nums[j] // slide the right bound
        j--
        count++
        if x == 0 { // record the opt num
            flag, res = true, min(res, count)
        }
    }
    if flag { return res }
    return -1
}

func minOperations1(nums []int, x int) int {
    sum, n := -x, len(nums)
    for _, v := range nums {
        sum += v
    }
    if sum < 0 { return -1 }
    if sum == 0 { return n }
    res, left := 0, 0
    for i := 0; i < n; i++ {
        sum -= nums[i]
        for sum < 0 {
            sum += nums[left]
            left++
        }
        if sum == 0 && res < i - left + 1 {
            res = i - left + 1
        }
    }
    if res == 0 { return -1 }
    return n - res
}

func main() {
    // Example 1:
    // Input: nums = [1,1,4,2,3], x = 5
    // Output: 2
    // Explanation: The optimal solution is to remove the last two elements to reduce x to zero.
    fmt.Println(minOperations([]int{1,1,4,2,3}, 5)) // 2
    // Example 2:
    // Input: nums = [5,6,7,8,9], x = 4
    // Output: -1
    fmt.Println(minOperations([]int{5,6,7,8,9}, 4)) // -1
    // Example 3:
    // Input: nums = [3,2,20,1,1,3], x = 10
    // Output: 5
    // Explanation: The optimal solution is to remove the last three elements and the first two elements (5 operations in total) to reduce x to zero.
    fmt.Println(minOperations([]int{3,2,20,1,1,3}, 10)) // 5

    fmt.Println(minOperations1([]int{1,1,4,2,3}, 5)) // 2
    fmt.Println(minOperations1([]int{5,6,7,8,9}, 4)) // -1
    fmt.Println(minOperations1([]int{3,2,20,1,1,3}, 10)) // 5
}