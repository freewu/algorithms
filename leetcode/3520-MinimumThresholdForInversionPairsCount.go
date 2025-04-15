package main

// 3520. Minimum Threshold for Inversion Pairs Count
// You are given an array of integers nums and an integer k.

// An inversion pair with a threshold x is defined as a pair of indices (i, j) such that:
//     1. i < j
//     2. nums[i] > nums[j]
//     3. The difference between the two numbers is at most x (i.e. nums[i] - nums[j] <= x).

// Your task is to determine the minimum integer min_threshold such that there are at least k inversion pairs with threshold min_threshold.

// If no such integer exists, return -1.

// Example 1:
// Input: nums = [1,2,3,4,3,2,1], k = 7
// Output: 2
// Explanation:
// For threshold x = 2, the pairs are:
// (3, 4) where nums[3] == 4 and nums[4] == 3.
// (2, 5) where nums[2] == 3 and nums[5] == 2.
// (3, 5) where nums[3] == 4 and nums[5] == 2.
// (4, 5) where nums[4] == 3 and nums[5] == 2.
// (1, 6) where nums[1] == 2 and nums[6] == 1.
// (2, 6) where nums[2] == 3 and nums[6] == 1.
// (4, 6) where nums[4] == 3 and nums[6] == 1.
// (5, 6) where nums[5] == 2 and nums[6] == 1.
// There are less than k inversion pairs if we choose any integer less than 2 as threshold.

// Example 2:
// Input: nums = [10,9,9,9,1], k = 4
// Output: 8
// Explanation:
// For threshold x = 8, the pairs are:
// (0, 1) where nums[0] == 10 and nums[1] == 9.
// (0, 2) where nums[0] == 10 and nums[2] == 9.
// (0, 3) where nums[0] == 10 and nums[3] == 9.
// (1, 4) where nums[1] == 9 and nums[4] == 1.
// (2, 4) where nums[2] == 9 and nums[4] == 1.
// (3, 4) where nums[3] == 9 and nums[4] == 1.
// There are less than k inversion pairs if we choose any integer less than 8 as threshold.

// Constraints:
//     1 <= nums.length <= 10^4
//     1 <= nums[i] <= 10^9
//     1 <= k <= 10^9

import "fmt"
import "sort"
import "slices"

func minThreshold(nums []int, k int) int {
    // Step 1: Discretize the numbers
    unique := make(map[int]bool)
    for _, num := range nums {
        unique[num] = true
    }
    arr := make([]int, 0, len(unique))
    for num := range unique {
        arr = append(arr, num)
    }
    sort.Ints(arr)
    n := len(arr)

    // Binary search to find the minimal threshold
    res, left, right := -1, 0, slices.Max(nums)
    for left <= right {
        mid := left + (right-left)/2
        if check(nums, arr, mid, k, n) {
            res, right = mid, mid - 1
        } else {
            left = mid + 1
        }
    }
    return res
}

func check(nums, arr []int, x, k, n int) bool {
    count := 0
    tree := make([]int, 4*n)

    for _, val := range nums {
        tmp := x + val
        right := sort.SearchInts(arr, tmp+1) - 1
        left := sort.SearchInts(arr, val)
        if left < len(arr) && arr[left] == val {
            left = left + 1
        }
        if left <= right && right >= 0 {
            count += query(tree, 0, 0, n-1, left, right)
        }
        update(tree, 0, 0, n-1, sort.SearchInts(arr, val), 1)
        if count >= k {
            return true
        }
    }
    return count >= k
}

func update(tree []int, root, start, end, index, val int) {
    if start == end {
        tree[root] += val
        return
    }
    mid := (start + end) / 2
    if index <= mid {
        update(tree, 2*root+1, start, mid, index, val)
    } else {
        update(tree, 2*root+2, mid+1, end, index, val)
    }
    tree[root] = tree[2*root+1] + tree[2*root+2]
}

func query(tree []int, root, start, end, left, right int) int {
    if left > end || right < start {
        return 0
    }
    if left <= start && end <= right {
        return tree[root]
    }
    mid := (start + end) / 2
    return query(tree, 2*root+1, start, mid, left, right) + query(tree, 2*root+2, mid+1, end, left, right)
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4,3,2,1], k = 7
    // Output: 2
    // Explanation:
    // For threshold x = 2, the pairs are:
    // (3, 4) where nums[3] == 4 and nums[4] == 3.
    // (2, 5) where nums[2] == 3 and nums[5] == 2.
    // (3, 5) where nums[3] == 4 and nums[5] == 2.
    // (4, 5) where nums[4] == 3 and nums[5] == 2.
    // (1, 6) where nums[1] == 2 and nums[6] == 1.
    // (2, 6) where nums[2] == 3 and nums[6] == 1.
    // (4, 6) where nums[4] == 3 and nums[6] == 1.
    // (5, 6) where nums[5] == 2 and nums[6] == 1.
    // There are less than k inversion pairs if we choose any integer less than 2 as threshold.
    fmt.Println(minThreshold([]int{1,2,3,4,3,2,1}, 7)) // 2
    // Example 2:
    // Input: nums = [10,9,9,9,1], k = 4
    // Output: 8
    // Explanation:
    // For threshold x = 8, the pairs are:
    // (0, 1) where nums[0] == 10 and nums[1] == 9.
    // (0, 2) where nums[0] == 10 and nums[2] == 9.
    // (0, 3) where nums[0] == 10 and nums[3] == 9.
    // (1, 4) where nums[1] == 9 and nums[4] == 1.
    // (2, 4) where nums[2] == 9 and nums[4] == 1.
    // (3, 4) where nums[3] == 9 and nums[4] == 1.
    // There are less than k inversion pairs if we choose any integer less than 8 as threshold.
    fmt.Println(minThreshold([]int{10,9,9,9,1}, 4)) // 8

    fmt.Println(minThreshold([]int{1,2,3,4,5,6,7,8,9}, 4)) // -1
    fmt.Println(minThreshold([]int{9,8,7,6,5,4,3,2,1}, 4)) // 1
}