package main

// 719. Find K-th Smallest Pair Distance
// The distance of a pair of integers a and b is defined as the absolute difference between a and b.

// Given an integer array nums and an integer k, 
// return the kth smallest distance among all the pairs nums[i] and nums[j] where 0 <= i < j < nums.length.

// Example 1:
// Input: nums = [1,3,1], k = 1
// Output: 0
// Explanation: Here are all the pairs:
// (1,3) -> 2
// (1,1) -> 0
// (3,1) -> 2
// Then the 1st smallest distance pair is (1,1), and its distance is 0.

// Example 2:
// Input: nums = [1,1,1], k = 2
// Output: 0

// Example 3:
// Input: nums = [1,6,1], k = 3
// Output: 5

// Constraints:
//     n == nums.length
//     2 <= n <= 10^4
//     0 <= nums[i] <= 10^6
//     1 <= k <= n * (n - 1) / 2

import "fmt"
import "sort"

func smallestDistancePair(nums []int, k int) int {
    sort.Ints(nums)
    return sort.Search(nums[len(nums) - 1] - nums[0] + 1, func (i int) bool {
        res := 0
        for p := 0; p <= len(nums) - 2; p++ {
            if nums[p] + i < nums[p + 1] {
                continue
            }
            index := sort.Search(len(nums), func(j int) bool {
                return nums[j] - nums[p] > i
            }) 
            res += index - p - 1
        }
        return res >= k
    })
}

func smallestDistancePair1(nums []int, k int) int {
    sort.Ints(nums)
    findPairs := func(nums []int, m int) int {
        count, i, j := 0, 0, 1
        for i < len(nums) {
            for j < len(nums) && nums[j]-nums[i] <= m {
                count += j - i
                j++
            }
            i++
        }
        return count
    }
    low, high := nums[1] - nums[0], nums[len(nums)-1] - nums[0]
    for i := 2; i < len(nums); i++ {
        diff := nums[i] - nums[i-1]
        if diff < low {
            low = diff
        }
    }
    for low < high {
        mid := (low + high) >> 1
        if findPairs(nums, mid) < k {
            low = mid + 1
        } else {
            high = mid
        }
    }
    return low
}

func main() {
    // Example 1:
    // Input: nums = [1,3,1], k = 1
    // Output: 0
    // Explanation: Here are all the pairs:
    // (1,3) -> 2
    // (1,1) -> 0
    // (3,1) -> 2
    // Then the 1st smallest distance pair is (1,1), and its distance is 0.
    fmt.Println(smallestDistancePair([]int{1,3,1}, 1)) // 0
    // Example 2:
    // Input: nums = [1,1,1], k = 2
    // Output: 0
    fmt.Println(smallestDistancePair([]int{1,1,1}, 2)) // 0
    // Example 3:
    // Input: nums = [1,6,1], k = 3
    // Output: 5
    fmt.Println(smallestDistancePair([]int{1,6,1}, 3)) // 5

    fmt.Println(smallestDistancePair1([]int{1,3,1}, 1)) // 0
    fmt.Println(smallestDistancePair1([]int{1,1,1}, 2)) // 0
    fmt.Println(smallestDistancePair1([]int{1,6,1}, 3)) // 5
}