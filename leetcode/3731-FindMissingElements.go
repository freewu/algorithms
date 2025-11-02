package main

// 3731. Find Missing Elements 
// You are given an integer array nums consisting of unique integers.

// Originally, nums contained every integer within a certain range. However, some integers might have gone missing from the array.

// The smallest and largest integers of the original range are still present in nums.

// Return a sorted list of all the missing integers in this range. If no integers are missing, return an empty list.

// Example 1:
// Input: nums = [1,4,2,5]
// Output: [3]
// Explanation:
// The smallest integer is 1 and the largest is 5, so the full range should be [1,2,3,4,5]. Among these, only 3 is missing.

// Example 2:
// Input: nums = [7,8,6,9]
// Output: []
// Explanation:
// The smallest integer is 6 and the largest is 9, so the full range is [6,7,8,9]. All integers are already present, so no integer is missing.

// Example 3:
// Input: nums = [5,1]
// Output: [2,3,4]
// Explanation:
// The smallest integer is 1 and the largest is 5, so the full range should be [1,2,3,4,5]. The missing integers are 2, 3, and 4.

// Constraints:
//     2 <= nums.length <= 100
//     1 <= nums[i] <= 100

import "fmt"
import "sort"

func findMissingElements(nums []int) []int {
    mn, mx := 1 << 31, -(1 << 31)
    res, mp := []int{}, map[int]bool{}
    for _, v := range nums {
        mn = min(mn, v)
        mx = max(mx, v)
        mp[v] = true
    }
    for i := mn + 1; i < mx; i++ {
        if !mp[i] {
            res = append(res, i)
        }
    }
    return res
}

func findMissingElements1(nums []int) []int {
    sort.Ints(nums)
    res, count, n := []int{}, nums[0] + 1, len(nums)
    for i := 1; i < n && count < nums[n-1]; count++ {
        if nums[i] > count {
            res = append(res, count)
        }else if nums[i] == count {
            i++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,4,2,5]
    // Output: [3]
    // Explanation:
    // The smallest integer is 1 and the largest is 5, so the full range should be [1,2,3,4,5]. Among these, only 3 is missing.
    fmt.Println(findMissingElements([]int{1,4,2,5})) // [3]
    // Example 2:
    // Input: nums = [7,8,6,9]
    // Output: []
    // Explanation:
    // The smallest integer is 6 and the largest is 9, so the full range is [6,7,8,9]. All integers are already present, so no integer is missing.
    fmt.Println(findMissingElements([]int{7,8,6,9})) // []
    // Example 3:
    // Input: nums = [5,1]
    // Output: [2,3,4]
    // Explanation:
    // The smallest integer is 1 and the largest is 5, so the full range should be [1,2,3,4,5]. The missing integers are 2, 3, and 4.
    fmt.Println(findMissingElements([]int{5,1})) // [2,3,4]

    fmt.Println(findMissingElements([]int{1,2,3,4,5,6,7,8,9})) // []
    fmt.Println(findMissingElements([]int{9,8,7,6,5,4,3,2,1})) // []

    fmt.Println(findMissingElements([]int{1,4,2,5})) // [3]
    fmt.Println(findMissingElements([]int{7,8,6,9})) // []
    fmt.Println(findMissingElements([]int{5,1})) // [2,3,4]
    fmt.Println(findMissingElements([]int{1,2,3,4,5,6,7,8,9})) // []
    fmt.Println(findMissingElements([]int{9,8,7,6,5,4,3,2,1})) // []
}