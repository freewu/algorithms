package main 

// 2009. Minimum Number of Operations to Make Array Continuous
// You are given an integer array nums. In one operation, you can replace any element in nums with any integer.
// nums is considered continuous if both of the following conditions are fulfilled:
//     All elements in nums are unique.
//     The difference between the maximum element and the minimum element in nums equals nums.length - 1.
//     For example, nums = [4, 2, 5, 3] is continuous, but nums = [1, 2, 3, 5, 6] is not continuous.

// Return the minimum number of operations to make nums continuous.

// Example 1:
// Input: nums = [4,2,5,3]
// Output: 0
// Explanation: nums is already continuous.

// Example 2:
// Input: nums = [1,2,3,5,6]
// Output: 1
// Explanation: One possible solution is to change the last element to 4.
// The resulting array is [1,2,3,5,4], which is continuous.

// Example 3:
// Input: nums = [1,10,100,1000]
// Output: 3
// Explanation: One possible solution is to:
// - Change the second element to 2.
// - Change the third element to 3.
// - Change the fourth element to 4.
// The resulting array is [1,2,3,4], which is continuous.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"
import "sort"
import "slices"

func minOperations(nums []int) int {
    l := len(nums)
    m := make(map[int]struct{}, l)
    for _, v := range nums {
        m[v] = struct{}{}
    }
    arr := make([]int, 0, len(m))
    for k := range m {
        arr = append(arr, k)
    }
    sort.Ints(arr)

    max := 0
    for i, v := range arr {
        n := sort.SearchInts(arr[i:], v+l)
        if n > max {
            max = n
        }
    }
    return l - max
}

func minOperations1(nums []int) int {
    n := len(nums)
    sort.Ints(nums)
    a := slices.Compact(nums)
    res, left := 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, x := range a {
        for a[left] < x-n+1 {
            left++
        }
        res = max(res, i - left + 1)
    }
    return n - res
}

func main() {
    // Explanation: nums is already continuous.
    fmt.Println(minOperations([]int{4,2,5,3})) // 0
    // Explanation: One possible solution is to change the last element to 4.
    // The resulting array is [1,2,3,5,4], which is continuous.
    fmt.Println(minOperations([]int{1,2,3,5,6})) // 1
    // Explanation: One possible solution is to:
    // - Change the second element to 2.
    // - Change the third element to 3.
    // - Change the fourth element to 4.
    // The resulting array is [1,2,3,4], which is continuous.
    fmt.Println(minOperations([]int{1,10,100,1000})) // 3

    fmt.Println(minOperations1([]int{4,2,5,3})) // 0
    fmt.Println(minOperations1([]int{1,2,3,5,6})) // 1
    fmt.Println(minOperations1([]int{1,10,100,1000})) // 3
}