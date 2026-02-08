package main

// 3833. Count Dominant Indices
// You are given an integer array nums of length n.

// An element at index i is called dominant if: nums[i] > average(nums[i + 1], nums[i + 2], ..., nums[n - 1])

// Your task is to count the number of indices i that are dominant.

// The average of a set of numbers is the value obtained by adding all the numbers together and dividing the sum by the total number of numbers.

// Note: The rightmost element of any array is not dominant.

// Example 1:
// Input: nums = [5,4,3]
// Output: 2
// Explanation:
// At index i = 0, the value 5 is dominant as 5 > average(4, 3) = 3.5.
// At index i = 1, the value 4 is dominant over the subarray [3].
// Index i = 2 is not dominant as there are no elements to its right. Thus, the answer is 2.

// Example 2:
// Input: nums = [4,1,2]
// Output: 1
// Explanation:
// At index i = 0, the value 4 is dominant over the subarray [1, 2].
// At index i = 1, the value 1 is not dominant.
// Index i = 2 is not dominant as there are no elements to its right. Thus, the answer is 1.

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 100​​​​​​​

import "fmt"

// 双指针
func dominantIndices(nums []int) int {
    res, k, i, j := 0, 0, 1, len(nums)
    for i < j {
        sum := 0
        for x := i; x < j; x++ {
            sum += nums[x]
        }
        if float64(nums[k]) > float64(sum) / float64(j-i) {
            res++
        }
        k++
        i++
    }
    return res
}

func dominantIndices1(nums []int) int {
    res, sum, n := 0, 0, len(nums)
    for i := n - 2; i >= 0; i-- {
        sum += nums[i+1]
        if nums[i] * (n - 1 - i) > sum {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [5,4,3]
    // Output: 2
    // Explanation:
    // At index i = 0, the value 5 is dominant as 5 > average(4, 3) = 3.5.
    // At index i = 1, the value 4 is dominant over the subarray [3].
    // Index i = 2 is not dominant as there are no elements to its right. Thus, the answer is 2.
    fmt.Println(dominantIndices([]int{5,4,3})) // 2
    // Example 2:
    // Input: nums = [4,1,2]
    // Output: 1
    // Explanation:
    // At index i = 0, the value 4 is dominant over the subarray [1, 2].
    // At index i = 1, the value 1 is not dominant.
    // Index i = 2 is not dominant as there are no elements to its right. Thus, the answer is 1.
    fmt.Println(dominantIndices([]int{4,1,2})) // 1

    fmt.Println(dominantIndices([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(dominantIndices([]int{9,8,7,6,5,4,3,2,1})) // 8

    fmt.Println(dominantIndices1([]int{5,4,3})) // 2
    fmt.Println(dominantIndices1([]int{4,1,2})) // 1
    fmt.Println(dominantIndices1([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(dominantIndices1([]int{9,8,7,6,5,4,3,2,1})) // 8
}