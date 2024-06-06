package main

// 1060. Missing Element in Sorted Array
// Given an integer array nums which is sorted in ascending order and all of its elements are unique and given also an integer k, return the kth missing number starting from the leftmost number of the array.

// Example 1:
// Input: nums = [4,7,9,10], k = 1
// Output: 5
// Explanation: The first missing number is 5.

// Example 2:
// Input: nums = [4,7,9,10], k = 3
// Output: 8
// Explanation: The missing numbers are [5,6,8,...], hence the third missing number is 8.

// Example 3:
// Input: nums = [1,2,4], k = 3
// Output: 6
// Explanation: The missing numbers are [3,5,6,7,...], hence the third missing number is 6.

// Constraints:
// 1 <= nums.length <= 5 * 10^4
// 1 <= nums[i] <= 10^7
// nums is sorted in ascending order, and all the elements are unique.
// 1 <= k <= 10^8
 
// Follow up: Can you find a logarithmic time complexity (i.e., O(log(n))) solution?

import "fmt"

func missingElement(nums []int, k int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]+1
    }
    helper := func(i int) int { // 指定索引缺了多少个数字
      return nums[i] - nums[0] - i;
    }
    if v := helper(n - 1); k > v { // 缺的元素在最大值后面时
        return nums[n - 1] + k - v
    }
    l, r := 0, len(nums)-1
    for l < r{
        mid := (l + r) >> 1
        if helper(mid) < k { // 缺的个数小于 k 时
            l = mid + 1
        }else{
            r = mid
        }
    }
    return nums[l - 1] + k - helper(l - 1)
}

func main() {
    // Example 1:
    // Input: nums = [4,7,9,10], k = 1
    // Output: 5
    // Explanation: The first missing number is 5.
    fmt.Println(missingElement([]int{4,7,9,10}, 1)) // 5
    // Example 2:
    // Input: nums = [4,7,9,10], k = 3
    // Output: 8
    // Explanation: The missing numbers are [5,6,8,...], hence the third missing number is 8.
    fmt.Println(missingElement([]int{4,7,9,10}, 3)) // 8
    // Example 3:
    // Input: nums = [1,2,4], k = 3
    // Output: 6
    // Explanation: The missing numbers are [3,5,6,7,...], hence the third missing number is 6.
    fmt.Println(missingElement([]int{1,2,4}, 3)) // 6
}