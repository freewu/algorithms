package main

// 698. Partition to K Equal Sum Subsets
// Given an integer array nums and an integer k, 
// return true if it is possible to divide this array into k non-empty subsets whose sums are all equal.

// Example 1:
// Input: nums = [4,3,2,3,5,2,1], k = 4
// Output: true
// Explanation: It is possible to divide it into 4 subsets (5), (1, 4), (2,3), (2,3) with equal sums.

// Example 2:
// Input: nums = [1,2,3,4], k = 3
// Output: false

// Constraints:
//     1 <= k <= nums.length <= 16
//     1 <= nums[i] <= 10^4
//     The frequency of each element is in the range [1, 4].

import "fmt"
import "sort"

func canPartitionKSubsets(nums []int, k int) bool {
    n, sum := len(nums), 0
    for i := 0; i < n; i++ {
        sum += nums[i]
    }
    visited := make([]bool, n)
    if k <= 0 || sum % k!= 0 {
        return false
    }
    var checkPartition func(nums []int,visited []bool,index,n,k,currSum,target int) bool
    checkPartition = func(nums []int,visited []bool,index,n,k,currSum,target int) bool {
        if k == 1 {
            return true
        }
        if currSum == target {
            return checkPartition(nums, visited, 0, n, k - 1, 0, target) 
        }
        for i := index; i < n; i++ {
            if visited[i] || currSum+nums[i] > target {
                continue
            }
            visited[i] = true
            if checkPartition(nums, visited, i + 1, n, k, currSum+nums[i], target) {
                return true
            }
            visited[i] = false
        }
        return false
    }
    return checkPartition(nums,visited,0,n,k,0,sum/k)
}

func canPartitionKSubsets1(nums []int, k int) bool {
    sum, mx, l := 0, 0, len(nums) - 1
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, num := range nums {
        sum += num
        mx = max(mx, num)
    }
    if (sum % k != 0) || (mx > sum / k) {
        return false
    }
    sort.Slice(nums, func (i, j int) bool  {
        return nums[i] <= nums[j]
    })
    var find func(idx int, mid int, kv int, vis []bool) bool
    find = func(idx int, mid int, kv int, vis []bool) bool {
        if kv == k {
            return true
        }
        if mid == sum / k {
            return find(l, 0, kv + 1, vis)
        }
        for i := idx ; i >= 0; i-- {
            if vis[i] || mid + nums[i] > sum / k {
                continue
            }
            vis[i] = true
            if find(i-1, mid+nums[i], kv, vis) {
                return true
            }
            vis[i] = false
            if mid == 0 {
                return false
            }
        }
        return false
    }
    vis := make([]bool, l + 1)
    return find(l, 0, 0, vis)
}

func main() {
    // Example 1:
    // Input: nums = [4,3,2,3,5,2,1], k = 4
    // Output: true
    // Explanation: It is possible to divide it into 4 subsets (5), (1, 4), (2,3), (2,3) with equal sums.
    fmt.Println(canPartitionKSubsets([]int{4,3,2,3,5,2,1}, 4)) // true
    // Example 2:
    // Input: nums = [1,2,3,4], k = 3
    // Output: false
    fmt.Println(canPartitionKSubsets([]int{1,2,3,4}, 3)) // false

    fmt.Println(canPartitionKSubsets1([]int{4,3,2,3,5,2,1}, 4)) // true
    fmt.Println(canPartitionKSubsets1([]int{1,2,3,4}, 3)) // false
}