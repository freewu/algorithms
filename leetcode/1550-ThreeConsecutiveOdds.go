package main

// 1550. Three Consecutive Odds
// Given an integer array arr, return true if there are three consecutive odd numbers in the array. 
// Otherwise, return false.
 
// Example 1:
// Input: arr = [2,6,4,1]
// Output: false
// Explanation: There are no three consecutive odds.

// Example 2:
// Input: arr = [1,2,34,3,4,5,7,23,12]
// Output: true
// Explanation: [5,7,23] are three consecutive odds.
 
// Constraints:
//     1 <= arr.length <= 1000
//     1 <= arr[i] <= 1000

import "fmt"

func threeConsecutiveOdds(arr []int) bool {
    count := 0
    for _, v := range arr {
        if v % 2 == 1 {
            count++
            if count == 3 { // 有连续3个奇数
                return true
            }
        } else { // 遇偶数归零
            count = 0
        }
    }
    return false
}

func threeConsecutiveOdds1(arr []int) bool {
    for i :=0; i <len(arr) - 2; i++ {
        if arr[i] % 2 != 0 && arr[i+1] % 2 != 0 && arr[i+2] % 2 != 0 {
            return true
        }
    }
    return false
}

func main() {
    // Example 1:
    // Input: arr = [2,6,4,1]
    // Output: false
    // Explanation: There are no three consecutive odds.
    fmt.Println(threeConsecutiveOdds([]int{2,6,4,1})) // false
    // Example 2:
    // Input: arr = [1,2,34,3,4,5,7,23,12]
    // Output: true
    // Explanation: [5,7,23] are three consecutive odds.
    fmt.Println(threeConsecutiveOdds([]int{1,2,34,3,4,5,7,23,12})) // true

    fmt.Println(threeConsecutiveOdds([]int{1,2,3,4,5,6,7,8,9})) // false
    fmt.Println(threeConsecutiveOdds([]int{9,8,7,6,5,4,3,2,1})) // false

    fmt.Println(threeConsecutiveOdds1([]int{2,6,4,1})) // false
    fmt.Println(threeConsecutiveOdds1([]int{1,2,34,3,4,5,7,23,12})) // true
    fmt.Println(threeConsecutiveOdds1([]int{1,2,3,4,5,6,7,8,9})) // false
    fmt.Println(threeConsecutiveOdds1([]int{9,8,7,6,5,4,3,2,1})) // false
}