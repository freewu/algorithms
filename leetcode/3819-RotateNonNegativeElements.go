package main

// 3819. Rotate Non Negative Elements
// You are given an integer array nums and an integer k.

// Rotate only the non-negative elements of the array to the left by k positions, in a cyclic manner.

// All negative elements must stay in their original positions and must not move.

// After rotation, place the non-negative elements back into the array in the new order, 
// filling only the positions that originally contained non-negative values and skipping all negative positions.

// Return the resulting array.

// Example 1:
// Input: nums = [1,-2,3,-4], k = 3
// Output: [3,-2,1,-4]
// Explanation:​​​​​​​
// The non-negative elements, in order, are [1, 3].
// Left rotation with k = 3 results in:
// [1, 3] -> [3, 1] -> [1, 3] -> [3, 1]
// Placing them back into the non-negative indices results in [3, -2, 1, -4].

// Example 2:
// Input: nums = [-3,-2,7], k = 1
// Output: [-3,-2,7]
// Explanation:
// The non-negative elements, in order, are [7].
// Left rotation with k = 1 results in [7].
// Placing them back into the non-negative indices results in [-3, -2, 7].

// Example 3:
// Input: nums = [5,4,-9,6], k = 2
// Output: [6,5,-9,4]
// Explanation:
// The non-negative elements, in order, are [5, 4, 6].
// Left rotation with k = 2 results in [6, 5, 4].
// Placing them back into the non-negative indices results in [6, 5, -9, 4].

// Constraints:
//     1 <= nums.length <= 10^5
//     -10^5 <= nums[i] <= 10^5
//     0 <= k <= 10^5

import "fmt"
import "slices"

func rotateElements(nums []int, k int) []int {
    arr := []int{}
    for _, v := range nums { // 取出非负数
        if v >= 0 {
            arr = append(arr, v)
        }
    }
    n := len(arr)
    if n == 0 { return nums } // 没有非负数，无需操作
    rotateLeft := func (a []int, k int) { // 轮转数组
        slices.Reverse(a[:k])
        slices.Reverse(a[k:])
        slices.Reverse(a)
    }
    rotateLeft(arr, k % n) // 向左轮替 k 个位置（原地操作）
    // 双指针，把 arr 填入 nums，跳过负数
    j := 0
    for i, v := range nums {
        if v >= 0 {
            nums[i] = arr[j]
            j++
        }
    }
    return nums
}

// 双指针
func rotateElements1(nums []int, k int) []int {
    arr := []int{}
    for _, v := range nums { // 取出非负数
        if v >= 0 {
            arr = append(arr, v)
        }
    }
    j := k // 双指针，把 arr 填入 nums，跳过负数
    for i, v := range nums {
        if v >= 0 {
            nums[i] = arr[j % len(arr)]
            j++
        }
    }
    return nums
}

func rotateElements2(nums []int, k int) []int {
    count, n, arr := 0, len(nums),make([]int,0)
    for i := 0; i < n; i++ { // 取出非负数
        if nums[i] >= 0 {
            count++
            arr = append(arr,nums[i])
        }
    }
    if count == 0 { return nums }
    k = k % count
    arr1 := make([]int,count)
    copy(arr1[:count-k], arr[k:])
    copy(arr1[count-k:], arr[:k])
    j := 0
    for i := 0; i < n; i++ {
        if nums[i] >= 0 {
            nums[i] = arr1[j]
            j++
        }
    }
    return nums
}

func main() {
    // Example 1:
    // Input: nums = [1,-2,3,-4], k = 3
    // Output: [3,-2,1,-4]
    // Explanation:​​​​​​​
    // The non-negative elements, in order, are [1, 3].
    // Left rotation with k = 3 results in:
    // [1, 3] -> [3, 1] -> [1, 3] -> [3, 1]
    // Placing them back into the non-negative indices results in [3, -2, 1, -4].
    fmt.Println(rotateElements([]int{1,-2,3,-4}, 3)) // [3, -2, 1, -4]
    // Example 2:
    // Input: nums = [-3,-2,7], k = 1
    // Output: [-3,-2,7]
    // Explanation:
    // The non-negative elements, in order, are [7].
    // Left rotation with k = 1 results in [7].
    // Placing them back into the non-negative indices results in [-3, -2, 7].
    fmt.Println(rotateElements([]int{-3,-2,7}, 1)) // [-3, -2, 7]
    // Example 3:
    // Input: nums = [5,4,-9,6], k = 2
    // Output: [6,5,-9,4]
    // Explanation:
    // The non-negative elements, in order, are [5, 4, 6].
    // Left rotation with k = 2 results in [6, 5, 4].
    // Placing them back into the non-negative indices results in [6, 5, -9, 4].
    fmt.Println(rotateElements([]int{5,4,-9,6}, 2)) // [6, 5, -9, 4]

    fmt.Println(rotateElements([]int{1,2,3,4,5,6,7,8,9}, 2)) // [3 4 5 6 7 8 9 1 2]
    fmt.Println(rotateElements([]int{9,8,7,6,5,4,3,2,1}, 2)) // [7 6 5 4 3 2 1 9 8]

    fmt.Println(rotateElements1([]int{1,-2,3,-4}, 3)) // [3, -2, 1, -4]
    fmt.Println(rotateElements1([]int{-3,-2,7}, 1)) // [-3, -2, 7]
    fmt.Println(rotateElements1([]int{5,4,-9,6}, 2)) // [6, 5, -9, 4]
    fmt.Println(rotateElements1([]int{1,2,3,4,5,6,7,8,9}, 2)) // [3 4 5 6 7 8 9 1 2]
    fmt.Println(rotateElements1([]int{9,8,7,6,5,4,3,2,1}, 2)) // [7 6 5 4 3 2 1 9 8]

    fmt.Println(rotateElements2([]int{1,-2,3,-4}, 3)) // [3, -2, 1, -4]
    fmt.Println(rotateElements2([]int{-3,-2,7}, 1)) // [-3, -2, 7]
    fmt.Println(rotateElements2([]int{5,4,-9,6}, 2)) // [6, 5, -9, 4]   
    fmt.Println(rotateElements2([]int{1,2,3,4,5,6,7,8,9}, 2)) // [3 4 5 6 7 8 9 1 2]
    fmt.Println(rotateElements2([]int{9,8,7,6,5,4,3,2,1}, 2)) // [7 6 5 4 3 2 1 9 8]
}