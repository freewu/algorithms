package main

// 4055. Count Shadow Pairs II
// You are given an integer array nums of length n.

// A pair of indices (i, j) is called a shadow pair if all of the following conditions are satisfied:
//     1. 0 <= i < j < n
//     2. nums[i] < nums[j]
//     3. There does not exist an index k such that i < k < j and nums[i] < nums[k] < nums[j].

// Return the total number of shadow pairs.

// Example 1:
// Input: nums = [3,1,4,2,5]
// Output: 5
// Explanation:
// (i, j) | nums[i] | nums[j] | Shadow Pair
// (0, 2) | 3       | 4       | nums[1] = 1 is not strictly between 3 and 4
// (1, 2) | 1       | 4       | No index k exists such that 1 < k < 2
// (1, 3) | 1       | 2       | nums[2] = 4 is not strictly between 1 and 2
// (2, 4) | 4       | 5       | nums[3] = 2 is not strictly between 4 and 5
// (3, 4) | 5       | 5       | No index k exists such that 3 < k < 4
// Thus, the answer is 5.

// Example 2:
// Input: nums = [6,7,8,9]
// Output: 3
// Explanation:
// (i, j) | nums[i] | nums[j] | Shadow Pair
// (0, 1) | 6       | 7       | No index k exists such that 0 < k < 1
// (1, 2) | 7       | 8       | No index k exists such that 1 < k < 2
// (2, 3) | 8       | 9       | No index k exists such that 2 < k < 3
// Thus, the answer is 3.

// Constraints:
//     3 <= n == nums.length <= 5 * 10^4
//     1 <= nums[i] <= 10^9

import "fmt"
import "sort"
import "slices"

func shadowPairs(nums []int) int {
    var solve func (arr []int, low, high int) int 
    solve = func (arr []int, low, high int) int {
        res, n := 0, len(arr)
        if n <= 1 || low == high {
            return res
        }
        lowStack, highStack, b, c := []int{}, []int{}, []int{}, []int{}
        mid := (low + high) / 2
        for i, v := range arr {
            if v <= mid { // x 在下部，作为 nums[i]
                for len(lowStack) > 0 && arr[lowStack[len(lowStack) - 1]] < v {
                    lowStack = lowStack[:len(lowStack)-1] // 因为 x 的出现，栈顶不能作为 nums[i]
                }
                lowStack = append(lowStack, i)
                b = append(b, v)
            } else { // v 在上部，作为 nums[j]
                // 找到 v 左侧第一个小于 v 的最近元素，作为 nums[k]
                for len(highStack) > 0 && arr[highStack[len(highStack)-1]] >= v {
                    highStack = highStack[:len(highStack)-1]
                }
                res += len(lowStack)
                if len(highStack) > 0 {
                    // lowStack 中 < highStack[len(highStack)-1] 的下标不能作为 nums[i]
                    res -= sort.SearchInts(lowStack, highStack[len(highStack)-1])
                }
                highStack = append(highStack, i)
                c = append(c, v)
            }
        }
        return res + solve(b, low, mid) + solve(c, mid+1, high)
    }
    sorted := slices.Clone(nums)
    slices.Sort(sorted)
    sorted = slices.Compact(sorted)
    for i, v := range nums {
        nums[i] = sort.SearchInts(sorted, v)
    }
    return solve(nums, 0, len(sorted)-1)
}

func main() {
    // Example 1:
    // Input: nums = [3,1,4,2,5]
    // Output: 5
    // Explanation:
    // (i, j) | nums[i] | nums[j] | Shadow Pair
    // (0, 2) | 3       | 4       | nums[1] = 1 is not strictly between 3 and 4
    // (1, 2) | 1       | 4       | No index k exists such that 1 < k < 2
    // (1, 3) | 1       | 2       | nums[2] = 4 is not strictly between 1 and 2
    // (2, 4) | 4       | 5       | nums[3] = 2 is not strictly between 4 and 5
    // (3, 4) | 5       | 5       | No index k exists such that 3 < k < 4
    // Thus, the answer is 5.
    fmt.Println(shadowPairs([]int{3,1,4,2,5})) // 5
    // Example 2:
    // Input: nums = [6,7,8,9]
    // Output: 3
    // Explanation:
    // (i, j) | nums[i] | nums[j] | Shadow Pair
    // (0, 1) | 6       | 7       | No index k exists such that 0 < k < 1
    // (1, 2) | 7       | 8       | No index k exists such that 1 < k < 2
    // (2, 3) | 8       | 9       | No index k exists such that 2 < k < 3
    // Thus, the answer is 3.
    fmt.Println(shadowPairs([]int{6,7,8,9})) // 3

    fmt.Println(shadowPairs([]int{1,2,3,4,5,6,7,8,9})) // 8
    fmt.Println(shadowPairs([]int{9,8,7,6,5,4,3,2,1})) // 0
}