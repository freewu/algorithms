package main

// 1287. Element Appearing More Than 25% In Sorted Array
// Given an integer array sorted in non-decreasing order, 
// there is exactly one integer in the array that occurs more than 25% of the time, return that integer.

// Example 1:
// Input: arr = [1,2,2,6,6,6,6,7,10]
// Output: 6

// Example 2:
// Input: arr = [1,1]
// Output: 1

// Constraints:
//     1 <= arr.length <= 10^4
//     0 <= arr[i] <= 10^5

import "fmt"

func findSpecialInteger(arr []int) int {
    target, mp := len(arr) / 4, make(map[int]int)
    for _, v := range arr {
        mp[v]++
        if mp[v] > target {
            return v
        }
    }
    return -1
}

func findSpecialInteger1(arr []int) int {
    count, total, num := 1, len(arr), arr[0]
    for i := 1; i < total; i++ {
        if arr[i] == num {
            count++
        } else { // 不一样了
            if float32(count) / float32(total) > 0.25 {
                break
            }
            num, count = arr[i], 1
        }
    }
    return num
}

func findSpecialInteger2(arr []int) int {
    maxCount, currentCount, res := 1, 1, arr[0]
    for i := 1; i < len(arr); i++ {
        if arr[i] == arr[i-1] {
            currentCount++
        } else {
            if currentCount > maxCount {
                maxCount, res = currentCount, arr[i - 1]
            }
            currentCount = 1
        }
    }
    if currentCount > maxCount {
        res = arr[len(arr)-1]
    }
    return res
}

func main() {
    // Example 1:
    // Input: arr = [1,2,2,6,6,6,6,7,10]
    // Output: 6
    fmt.Println(findSpecialInteger([]int{1,2,2,6,6,6,6,7,10})) // 6
    // Example 2:
    // Input: arr = [1,1]
    // Output: 1
    fmt.Println(findSpecialInteger([]int{1,1})) // 1

    fmt.Println(findSpecialInteger1([]int{1,2,2,6,6,6,6,7,10})) // 6
    fmt.Println(findSpecialInteger1([]int{1,1})) // 1

    fmt.Println(findSpecialInteger2([]int{1,2,2,6,6,6,6,7,10})) // 6
    fmt.Println(findSpecialInteger2([]int{1,1})) // 1
}