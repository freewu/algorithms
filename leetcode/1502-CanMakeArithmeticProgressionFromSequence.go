package main

// 1502. Can Make Arithmetic Progression From Sequence
// A sequence of numbers is called an arithmetic progression if the difference between any two consecutive elements is the same.
// Given an array of numbers arr, return true if the array can be rearranged to form an arithmetic progression. Otherwise, return false.

// Example 1:
// Input: arr = [3,5,1]
// Output: true
// Explanation: We can reorder the elements as [1,3,5] or [5,3,1] with differences 2 and -2 respectively, between each consecutive elements.

// Example 2:
// Input: arr = [1,2,4]
// Output: false
// Explanation: There is no way to reorder the elements to obtain an arithmetic progression.
 
// Constraints:
//     2 <= arr.length <= 1000
//     -10^6 <= arr[i] <= 10^6

import "fmt"
import "sort"

func canMakeArithmeticProgression(arr []int) bool {
    if len(arr) < 3 {
        return true
    }
    sort.Ints(arr)
    for i := 1; i < len(arr) - 1; i++ {
        if (arr[i] - arr[i - 1]) != (arr[i+1] - arr[i]) {
            return false
        }
    }
    return true
}

func canMakeArithmeticProgression1(arr []int) bool {
    sort.Ints(arr)
    // 计算好 for 里减少 一次计算
    diff := arr[1] - arr[0]
    for i := 2; i < len(arr); i++ {
        if arr[i] - arr[i-1] != diff {
            return false
        }
    }
    return true
}

func main( ) {
    fmt.Println(canMakeArithmeticProgression([]int{3,5,1})) // true
    fmt.Println(canMakeArithmeticProgression([]int{1,2,4})) // false
    fmt.Println(canMakeArithmeticProgression([]int{1,2,3,4,5,6})) // true
    fmt.Println(canMakeArithmeticProgression([]int{1,100})) // true
    fmt.Println(canMakeArithmeticProgression([]int{-13,-17,-8,-10,-20,2,3,-19,2,-18,-5,7,-12,18,-17,12,-1})) // false

    fmt.Println(canMakeArithmeticProgression1([]int{3,5,1})) // true
    fmt.Println(canMakeArithmeticProgression1([]int{1,2,4})) // false
    fmt.Println(canMakeArithmeticProgression1([]int{1,2,3,4,5,6})) // true
    fmt.Println(canMakeArithmeticProgression1([]int{1,100})) // true
    fmt.Println(canMakeArithmeticProgression1([]int{-13,-17,-8,-10,-20,2,3,-19,2,-18,-5,7,-12,18,-17,12,-1})) // false
}