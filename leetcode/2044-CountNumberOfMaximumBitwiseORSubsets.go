package main

// 2044. Count Number of Maximum Bitwise-OR Subsets
// Given an integer array nums, find the maximum possible bitwise OR of a subset of nums 
// and return the number of different non-empty subsets with the maximum bitwise OR.

// An array a is a subset of an array b if a can be obtained from b by deleting some (possibly zero) elements of b. 
// Two subsets are considered different if the indices of the elements chosen are different.

// The bitwise OR of an array a is equal to a[0] OR a[1] OR ... OR a[a.length - 1] (0-indexed).

// Example 1:
// Input: nums = [3,1]
// Output: 2
// Explanation: The maximum possible bitwise OR of a subset is 3. There are 2 subsets with a bitwise OR of 3:
// - [3]
// - [3,1]

// Example 2:
// Input: nums = [2,2,2]
// Output: 7
// Explanation: All non-empty subsets of [2,2,2] have a bitwise OR of 2. There are 23 - 1 = 7 total subsets.

// Example 3:
// Input: nums = [3,2,1,5]
// Output: 6
// Explanation: The maximum possible bitwise OR of a subset is 7. There are 6 subsets with a bitwise OR of 7:
// - [3,5]
// - [3,1,5]
// - [3,2,5]
// - [3,2,1,5]
// - [2,5]
// - [2,1,5]

// Constraints:
//     1 <= nums.length <= 16
//     1 <= nums[i] <= 10^5

import "fmt"

func countMaxOrSubsets(nums []int) int {
    res, mx, n := 0, 0, len(nums)
    for _, v := range nums { // 找到最大值
        mx |= v
    }
    var dfs func(cur int, index int) // index序号(nums的), cur当前nums[index]的或值
    dfs = func (cur int, index int) {
        if cur == mx {// 或值为最大时增加数量
            res += 1 << (n - index)
            return
        }
        if index == n { return }
        dfs(cur | nums[index], index+1)// 或上nums[下一个index]
        dfs(cur, index + 1) // 不或上nums[index]
    }
    dfs(0, 0) // 遍历选或者不选某元素的结果
    return res
}

func countMaxOrSubsets1(nums []int) int {
    mx, n := 0, len(nums)
    for _, n := range nums {
        mx |= n
    }
    var dfs func(i, cur, mx int) int
    dfs = func(i, cur, mx int) int {
        if cur == mx { return 1 << (n - i) }
        if i == n    {  return 0 }
        return dfs(i + 1, cur, mx) + dfs(i + 1, cur | nums[i], mx)
    }
    return dfs(0, 0, mx)
}

func main() {
    // Example 1:
    // Input: nums = [3,1]
    // Output: 2
    // Explanation: The maximum possible bitwise OR of a subset is 3. There are 2 subsets with a bitwise OR of 3:
    // - [3]
    // - [3,1]
    fmt.Println(countMaxOrSubsets([]int{3,1})) // 2
    // Example 2:
    // Input: nums = [2,2,2]
    // Output: 7
    // Explanation: All non-empty subsets of [2,2,2] have a bitwise OR of 2. There are 2^3 - 1 = 7 total subsets.
    fmt.Println(countMaxOrSubsets([]int{2,2,2})) // 7
    // Example 3:
    // Input: nums = [3,2,1,5]
    // Output: 6
    // Explanation: The maximum possible bitwise OR of a subset is 7. There are 6 subsets with a bitwise OR of 7:
    // - [3,5]
    // - [3,1,5]
    // - [3,2,5]
    // - [3,2,1,5]
    // - [2,5]
    // - [2,1,5]
    fmt.Println(countMaxOrSubsets([]int{3,2,1,5})) // 6

    fmt.Println(countMaxOrSubsets([]int{1,2,3,4,5,6,7,8,9})) // 337
    fmt.Println(countMaxOrSubsets([]int{9,8,7,6,5,4,3,2,1})) // 337

    fmt.Println(countMaxOrSubsets1([]int{3,1})) // 2
    fmt.Println(countMaxOrSubsets1([]int{2,2,2})) // 7
    fmt.Println(countMaxOrSubsets1([]int{3,2,1,5})) // 6
    fmt.Println(countMaxOrSubsets1([]int{1,2,3,4,5,6,7,8,9})) // 337
    fmt.Println(countMaxOrSubsets1([]int{9,8,7,6,5,4,3,2,1})) // 337
}