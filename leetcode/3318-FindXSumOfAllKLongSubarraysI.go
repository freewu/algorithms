package main

// 3318. Find X-Sum of All K-Long Subarrays I
// You are given an array nums of n integers and two integers k and x.

// The x-sum of an array is calculated by the following procedure:
//     1. Count the occurrences of all elements in the array.
//     2. Keep only the occurrences of the top x most frequent elements. 
//        If two elements have the same number of occurrences, the element with the bigger value is considered more frequent.
//     3. Calculate the sum of the resulting array.

// Note that if an array has less than x distinct elements, its x-sum is the sum of the array.

// Return an integer array answer of length n - k + 1 where answer[i] is the x-sum of the subarray nums[i..i + k - 1].

// Example 1:
// Input: nums = [1,1,2,2,3,4,2,3], k = 6, x = 2
// Output: [6,10,12]
// Explanation:
// For subarray [1, 1, 2, 2, 3, 4], only elements 1 and 2 will be kept in the resulting array. Hence, answer[0] = 1 + 1 + 2 + 2.
// For subarray [1, 2, 2, 3, 4, 2], only elements 2 and 4 will be kept in the resulting array. Hence, answer[1] = 2 + 2 + 2 + 4. Note that 4 is kept in the array since it is bigger than 3 and 1 which occur the same number of times.
// For subarray [2, 2, 3, 4, 2, 3], only elements 2 and 3 are kept in the resulting array. Hence, answer[2] = 2 + 2 + 2 + 3 + 3.

// Example 2:
// Input: nums = [3,8,7,8,7,5], k = 2, x = 2
// Output: [11,15,15,15,12]
// Explanation:
// Since k == x, answer[i] is equal to the sum of the subarray nums[i..i + k - 1].

// Constraints:
//     1 <= n == nums.length <= 50
//     1 <= nums[i] <= 50
//     1 <= x <= k <= nums.length

import "fmt"
import "slices"
import "sort"

func findXSum(nums []int, k int, x int) []int {
    type Pair struct {
        First, Second int
    }
    n := len(nums)
    res := make([]int, n - k + 1)
    for i := 0; i < n - k + 1; i++ {
        mp := make(map[int]int)
        for j := i; j < i + k; j++ {
            mp[nums[j]]++
        }
        arr := []Pair{}
        for k, v := range mp {
            arr = append(arr, Pair{ v, k })
        }
        slices.SortFunc(arr, func(a, b Pair) int {
            if a.First == b.First {
                return b.Second - a.Second
            }
            return b.First - a.First
        })
        sum, uniq := 0, 0 
        for _, v := range arr {
            if uniq >= x { break }
            sum += (v.First * v.Second)
            uniq++
        }
        res[i] = sum 
    }
    return res
}

func findXSum1(nums []int, k int, x int) []int {
    n := len(nums)
    res := make([]int, n - k + 1)
    type Element struct { // Create a slice to sort the elements by frequency and value
        Value, Count int
    }
    for i := 0; i <= n - k; i++ {
        mp := make(map[int]int)
        for j := i; j < i + k; j++ { // Count occurrences in the current subarray nums[i..i+k-1]
            mp[nums[j]]++
        }
        elements := make([]Element, 0, len(mp))
        for k, v := range mp {
            elements = append(elements, Element{ k, v })
        }
        sort.Slice(elements, func(a, b int) bool { // Sort elements first by count (descending) and then by value (descending)
            if elements[a].Count == elements[b].Count {
                return elements[a].Value > elements[b].Value
            }
            return elements[a].Count > elements[b].Count
        })
        sum := 0 
        for j := 0; j < x && j < len(elements); j++ { // Calculate the x-sum for the current subarray
            sum += ( elements[j].Value * elements[j].Count )
        }
        res[i] = sum
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,1,2,2,3,4,2,3], k = 6, x = 2
    // Output: [6,10,12]
    // Explanation:
    // For subarray [1, 1, 2, 2, 3, 4], only elements 1 and 2 will be kept in the resulting array. Hence, answer[0] = 1 + 1 + 2 + 2.
    // For subarray [1, 2, 2, 3, 4, 2], only elements 2 and 4 will be kept in the resulting array. Hence, answer[1] = 2 + 2 + 2 + 4. Note that 4 is kept in the array since it is bigger than 3 and 1 which occur the same number of times.
    // For subarray [2, 2, 3, 4, 2, 3], only elements 2 and 3 are kept in the resulting array. Hence, answer[2] = 2 + 2 + 2 + 3 + 3.
    fmt.Println(findXSum([]int{1,1,2,2,3,4,2,3}, 6, 2)) // [6,10,12]
    // Example 2:
    // Input: nums = [3,8,7,8,7,5], k = 2, x = 2
    // Output: [11,15,15,15,12]
    // Explanation:
    // Since k == x, answer[i] is equal to the sum of the subarray nums[i..i + k - 1].
    fmt.Println(findXSum([]int{3,8,7,8,7,5}, 2, 2)) // [11,15,15,15,12]

    fmt.Println(findXSum([]int{1,2,3,4,5,6,7,8,9}, 2, 2)) // [3 5 7 9 11 13 15 17]
    fmt.Println(findXSum([]int{9,8,7,6,5,4,3,2,1}, 2, 2)) // [17 15 13 11 9 7 5 3]

    fmt.Println(findXSum1([]int{1,1,2,2,3,4,2,3}, 6, 2)) // [6,10,12]
    fmt.Println(findXSum1([]int{3,8,7,8,7,5}, 2, 2)) // [11,15,15,15,12]
    fmt.Println(findXSum1([]int{1,2,3,4,5,6,7,8,9}, 2, 2)) // [3 5 7 9 11 13 15 17]
    fmt.Println(findXSum1([]int{9,8,7,6,5,4,3,2,1}, 2, 2)) // [17 15 13 11 9 7 5 3]
}