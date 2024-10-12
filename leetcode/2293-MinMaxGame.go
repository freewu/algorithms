package main

// 2293. Min Max Game
// You are given a 0-indexed integer array nums whose length is a power of 2.

// Apply the following algorithm on nums:
//     1. Let n be the length of nums. If n == 1, end the process. Otherwise, create a new 0-indexed integer array newNums of length n / 2.
//     2. For every even index i where 0 <= i < n / 2, assign the value of newNums[i] as min(nums[2 * i], nums[2 * i + 1]).
//     3. For every odd index i where 0 <= i < n / 2, assign the value of newNums[i] as max(nums[2 * i], nums[2 * i + 1]).
//     4. Replace the array nums with newNums.
//     5. Repeat the entire process starting from step 1.

// Return the last number that remains in nums after applying the algorithm.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/04/13/example1drawio-1.png" />
// Input: nums = [1,3,5,2,4,8,2,2]
// Output: 1
// Explanation: The following arrays are the results of applying the algorithm repeatedly.
// First: nums = [1,5,4,2]
// Second: nums = [1,4]
// Third: nums = [1]
// 1 is the last remaining number, so we return 1.

// Example 2:
// Input: nums = [3]
// Output: 3
// Explanation: 3 is already the last remaining number, so we return 3.

// Constraints:
//     1 <= nums.length <= 1024
//     1 <= nums[i] <= 10^9
//     nums.length is a power of 2.

import "fmt"

// recursive
func minMaxGame(nums []int) int {
    if len(nums) == 1 {  return nums[0] }
    n := len(nums) / 2
    arr := make([]int, n)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        index := i * 2
        if i % 2 == 0 {
            arr[i] = min(nums[index], nums[index + 1])
        } else {
            arr[i] = max(nums[index], nums[index + 1])
        }
    }
    return minMaxGame(arr)
}

func minMaxGame1(nums []int) int {
    ex, ll := len(nums), len(nums)
    for ll > 1 {
        index, newNums := 1, []int{}
        for i := 1; i < ex; i += 2 {
            if index % 2 == 0 {
                if nums[i] > nums[i-1] {
                    newNums = append(newNums, nums[i])
                } else {
                    newNums = append(newNums, nums[i-1])
                }
            } else {
                if nums[i] > nums[i-1] {
                    newNums = append(newNums, nums[i-1])
                } else {
                    newNums = append(newNums, nums[i])
                }
            }
            index = index + 1
        }
        // should use copy(to, from) instead of a loop (S1001)
        ex = ll - len(newNums)
        // 重新规划nums
        copy(nums, newNums)
        ll = ll - ex
    }
    return nums[0]
}

func minMaxGame2(nums []int) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for len(nums) != 1 {
        j := 0
        for i := 0; i < len(nums) >> 1; i += 1 {
            if i & 1 == 0 {
                nums[j] = min(nums[i * 2], nums[2 * i + 1])
            } else {
                nums[j] = max(nums[i * 2], nums[2 * i + 1])
            }
            j += 1
        }
        nums = nums[:j]
    }
    return nums[0]
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/04/13/example1drawio-1.png" />
    // Input: nums = [1,3,5,2,4,8,2,2]
    // Output: 1
    // Explanation: The following arrays are the results of applying the algorithm repeatedly.
    // First: nums = [1,5,4,2]
    // Second: nums = [1,4]
    // Third: nums = [1]
    // 1 is the last remaining number, so we return 1.
    fmt.Println(minMaxGame([]int{1,3,5,2,4,8,2,2})) // 1
    // Example 2:
    // Input: nums = [3]
    // Output: 3
    // Explanation: 3 is already the last remaining number, so we return 3.
    fmt.Println(minMaxGame([]int{3})) // 3

    fmt.Println(minMaxGame1([]int{1,3,5,2,4,8,2,2})) // 1
    fmt.Println(minMaxGame1([]int{3})) // 3

    fmt.Println(minMaxGame2([]int{1,3,5,2,4,8,2,2})) // 1
    fmt.Println(minMaxGame2([]int{3})) // 3
}