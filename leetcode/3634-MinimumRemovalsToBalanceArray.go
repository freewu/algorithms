package main

// 3634. Minimum Removals to Balance Array
// You are given an integer array nums and an integer k.

// An array is considered balanced if the value of its maximum element is at most k times the minimum element.

// You may remove any number of elements from nums​​​​​​​ without making it empty.

// Return the minimum number of elements to remove so that the remaining array is balanced.

// Note: An array of size 1 is considered balanced as its maximum and minimum are equal, and the condition always holds true.

// Example 1:
// Input: nums = [2,1,5], k = 2
// Output: 1
// Explanation:
// Remove nums[2] = 5 to get nums = [2, 1].
// Now max = 2, min = 1 and max <= min * k as 2 <= 1 * 2. Thus, the answer is 1.

// Example 2:
// Input: nums = [1,6,2,9], k = 3
// Output: 2
// Explanation:
// Remove nums[0] = 1 and nums[3] = 9 to get nums = [6, 2].
// Now max = 6, min = 2 and max <= min * k as 6 <= 2 * 3. Thus, the answer is 2.

// Example 3:
// Input: nums = [4,6], k = 2
// Output: 0
// Explanation:
// Since nums is already balanced as 6 <= 4 * 2, no elements need to be removed.
 
// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     1 <= k <= 10^5

import "fmt"
import "sort"

func minRemoval(nums []int, k int) int {
    sort.Ints(nums)
    mx, left := 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, v := range nums {
        for nums[left] * k < v {
            left++
        }
        mx = max(mx, i - left + 1)
    }
    return len(nums) - mx
}

func minRemoval1(nums []int, k int) int {
    sort.Ints(nums)
    save := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for left, i := 0, 0; i < len(nums); i++ {
        for nums[left] * k < nums[i] {
            left++
        }
        save = max(save, i - left + 1)
    }
    return len(nums) - save
}

func minRemoval2(nums []int, k int) int {
    sort.Ints(nums)
    p ,count := 0 ,0
    for i := 0;i < len(nums);i ++ {
        for float64(nums[i]) > float64(nums[p]) * float64(k) {
            p ++
        }
        if i - p + 1 > count {
            count = i - p + 1
        }
    }
    return len(nums) - count
}

func main() {
    // Example 1:
    // Input: nums = [2,1,5], k = 2
    // Output: 1
    // Explanation:
    // Remove nums[2] = 5 to get nums = [2, 1].
    // Now max = 2, min = 1 and max <= min * k as 2 <= 1 * 2. Thus, the answer is 1.
    fmt.Println(minRemoval([]int{2,1,5}, 2)) // 1
    // Example 2:
    // Input: nums = [1,6,2,9], k = 3
    // Output: 2
    // Explanation:
    // Remove nums[0] = 1 and nums[3] = 9 to get nums = [6, 2].
    // Now max = 6, min = 2 and max <= min * k as 6 <= 2 * 3. Thus, the answer is 2.
    fmt.Println(minRemoval([]int{1,6,2,9}, 3)) // 2
    // Example 3:
    // Input: nums = [4,6], k = 2
    // Output: 0
    // Explanation:
    // Since nums is already balanced as 6 <= 4 * 2, no elements need to be removed. 
    fmt.Println(minRemoval([]int{4,6}, 2)) // 0

    fmt.Println(minRemoval([]int{1,2,3,4,5,6,7,8,9}, 2)) // 4
    fmt.Println(minRemoval([]int{9,8,7,6,5,4,3,2,1}, 2)) // 4

    fmt.Println(minRemoval1([]int{2,1,5}, 2)) // 1
    fmt.Println(minRemoval1([]int{1,6,2,9}, 3)) // 2
    fmt.Println(minRemoval1([]int{4,6}, 2)) // 0
    fmt.Println(minRemoval1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 4
    fmt.Println(minRemoval1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 4

    fmt.Println(minRemoval2([]int{2,1,5}, 2)) // 1
    fmt.Println(minRemoval2([]int{1,6,2,9}, 3)) // 2
    fmt.Println(minRemoval2([]int{4,6}, 2)) // 0
    fmt.Println(minRemoval2([]int{1,2,3,4,5,6,7,8,9}, 2)) // 4
    fmt.Println(minRemoval2([]int{9,8,7,6,5,4,3,2,1}, 2)) // 4
}