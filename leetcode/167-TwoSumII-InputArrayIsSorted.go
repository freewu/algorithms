package main

// 167. Two Sum II - Input array is sorted
// Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order,
// find two numbers such that they add up to a specific target number.
// Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.
// Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.
// The tests are generated such that there is exactly one solution. You may not use the same element twice.
// Your solution must use only constant extra space.

// Example 1:
// Input: numbers = [2,7,11,15], target = 9
// Output: [1,2]
// Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].

// Example 2:
// Input: numbers = [2,3,4], target = 6
// Output: [1,3]
// Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].

// Example 3:
// Input: numbers = [-1,0], target = -1
// Output: [1,2]
// Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].

// Constraints:
//     2 <= numbers.length <= 3 * 10^4
//     -1000 <= numbers[i] <= 1000
//     numbers is sorted in non-decreasing order.
//     -1000 <= target <= 1000
//     The tests are generated such that there is exactly one solution.

import "fmt"

// 利用数组有序的特性的解法
func twoSum(numbers []int, target int) []int {
    i, j := 0, len(numbers)-1
    for i < j { // 从外向内缩进
        val := numbers[i] + numbers[j]
        if val == target { // 如果刚好匹配则返回
            return []int{i + 1, j + 1}
        }
        if val < target { // 如果过小，说明开头需要向里走 ->
            i++
        } else { // 过大, 从尾部向头收缩 <-
            j--
        }
    }
    return nil
}

// Two Sum 的解法
func twoSum1(numbers []int, target int) []int {
    m := make(map[int]int)
    for i := 0; i < len(numbers); i++ {
        another := target - numbers[i]
        if index, ok := m[another]; ok {
            return []int{index + 1, i + 1}
        }
        m[numbers[i]] = i
    }
    return nil
}

func main() {
    // Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].
    fmt.Printf("twoSum([]int{ 2,7,11,15 },9) = %v\n",twoSum([]int{ 2,7,11,15 },9)) // [1,2]  2 + 7  = 9
    // Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].
    fmt.Printf("twoSum([]int{ 2,3,4 },6) = %v\n",twoSum([]int{ 2,3,4 },6)) // [1,3]  2 + 4  = 6
    // Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].
    fmt.Printf("twoSum([]int{ -1,0 },-1) = %v\n",twoSum([]int{ -1,0 },-1)) // [1,2]  -1 + 0  = -1

    fmt.Printf("twoSum1([]int{ 2,7,11,15 },9) = %v\n",twoSum1([]int{ 2,7,11,15 },9)) // [1,2]  2 + 7  = 9
    fmt.Printf("twoSum1([]int{ 2,3,4 },6) = %v\n",twoSum1([]int{ 2,3,4 },6)) // [1,3]  2 + 4  = 6
    fmt.Printf("twoSum1([]int{ -1,0 },-1) = %v\n",twoSum1([]int{ -1,0 },-1)) // [1,2]  -1 + 0  = -1
}