package main

// 2553. Separate the Digits in an Array
// Given an array of positive integers nums, 
// return an array answer that consists of the digits of each integer in nums after separating them in the same order they appear in nums.

// To separate the digits of an integer is to get all the digits it has in the same order.
//     For example, for the integer 10921, the separation of its digits is [1,0,9,2,1].

// Example 1:
// Input: nums = [13,25,83,77]
// Output: [1,3,2,5,8,3,7,7]
// Explanation: 
// - The separation of 13 is [1,3].
// - The separation of 25 is [2,5].
// - The separation of 83 is [8,3].
// - The separation of 77 is [7,7].
// answer = [1,3,2,5,8,3,7,7]. Note that answer contains the separations in the same order.

// Example 2:
// Input: nums = [7,1,3,9]
// Output: [7,1,3,9]
// Explanation: The separation of each integer in nums is itself.
// answer = [7,1,3,9].

// Constraints:
//     1 <= nums.length <= 1000
//     1 <= nums[i] <= 10^5

import "fmt"
import "strconv"

func separateDigits(nums []int) []int {
    res, str := []int{}, ""
    for _, v := range nums {
        str = strconv.Itoa(v)
        for j := 0; j < len(str); j++ {
            t, _ := strconv.Atoi(str[j:j+1])
            res = append(res, t)
        }
    }
    return res
}

func separateDigits1(nums []int) []int {
    reverse := func(arr []int) []int {
        n := len(arr)
        for i := 0; i < n / 2; i++ {
            arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
        }
        return arr
    }
    res := []int{}
    nums = reverse(nums)
    for _, v := range nums {
        for v != 0 {
            res = append(res, v % 10)
            v /= 10
        }
    }
    return reverse(res)
}


func main() {
    // Example 1:
    // Input: nums = [13,25,83,77]
    // Output: [1,3,2,5,8,3,7,7]
    // Explanation: 
    // - The separation of 13 is [1,3].
    // - The separation of 25 is [2,5].
    // - The separation of 83 is [8,3].
    // - The separation of 77 is [7,7].
    // answer = [1,3,2,5,8,3,7,7]. Note that answer contains the separations in the same order.
    fmt.Println(separateDigits([]int{13,25,83,77})) // [1,3,2,5,8,3,7,7]
    // Example 2:
    // Input: nums = [7,1,3,9]
    // Output: [7,1,3,9]
    // Explanation: The separation of each integer in nums is itself.
    // answer = [7,1,3,9].
    fmt.Println(separateDigits([]int{7,1,3,9})) // [7,1,3,9]

    fmt.Println(separateDigits([]int{1,2,3,4,5,6,7,8,9})) // [1 2 3 4 5 6 7 8 9]
    fmt.Println(separateDigits([]int{9,8,7,6,5,4,3,2,1})) // [9 8 7 6 5 4 3 2 1]

    fmt.Println(separateDigits1([]int{13,25,83,77})) // [1,3,2,5,8,3,7,7]
    fmt.Println(separateDigits1([]int{7,1,3,9})) // [7,1,3,9]
    fmt.Println(separateDigits1([]int{1,2,3,4,5,6,7,8,9})) // [1 2 3 4 5 6 7 8 9]
    fmt.Println(separateDigits1([]int{9,8,7,6,5,4,3,2,1})) // [9 8 7 6 5 4 3 2 1]
}