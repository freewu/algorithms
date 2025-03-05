package main

// 面试题 16.21. Sum Swap LCCI
// Given two arrays of integers, find a pair of values (one value from each array) that you can swap to give the two arrays the same sum.

// Return an array, where the first element is the element in the first array that will be swapped, 
// and the second element is another one in the second array. 
// If there are more than one answers, return any one of them. 
// If there is no answer, return an empty array.

// Example1:
// Input: array1 = [4, 1, 2, 1, 1, 2], array2 = [3, 6, 3, 3]
// Output: [1, 3]

// Example2:
// Input: array1 = [1, 2, 3], array2 = [4, 5, 6]
// Output: []

// Note:
//     1 <= array1.length, array2.length <= 100000

import "fmt"

func findSwapValues(array1 []int, array2 []int) []int {
    sum1, sum2 := 0, 0
    for _, v := range array1 {
        sum1 += v
    }
    for _, v := range array2 {
        sum2 += v
    }
    diff := sum2 - sum1
    if diff & 1 == 1 {
        return nil
    }
    diff >>= 1
    mp := map[int]bool{}
    for _,v := range array2 {
        mp[v] = true
    }
    for _,v := range array1 {
        if mp[diff + v] {
            return []int{ v, diff + v}
        }
    }
    return nil
}

func findSwapValues1(array1 []int, array2 []int) []int {
    mp := make(map[int]bool, len(array1))
    for _, v := range array1 {
        mp[v] = true
    }
    sum := func(arr []int) (res int) { for _, v := range arr {; res += v; }; return ; }
    sum1, sum2 := sum(array1), sum(array2)
    if (sum1 - sum2) & 1 != 0 {
        return nil
    }
    diff := (sum1 - sum2) / 2
    for _, v := range array2 {
        if mp[v + diff] {
            return []int{v + diff, v}
        }
    }
    return nil
}

func main() {
    // Example1:
    // Input: array1 = [4, 1, 2, 1, 1, 2], array2 = [3, 6, 3, 3]
    // Output: [1, 3]
    fmt.Println(findSwapValues([]int{4, 1, 2, 1, 1, 2}, []int{3, 6, 3, 3})) // [1, 3]
    // Example2:
    // Input: array1 = [1, 2, 3], array2 = [4, 5, 6]
    // Output: []
    fmt.Println(findSwapValues([]int{1, 2, 3}, []int{4, 5, 6})) // []

    fmt.Println(findSwapValues([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // [1 1]
    fmt.Println(findSwapValues([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // [1 1]
    fmt.Println(findSwapValues([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // [9 9]
    fmt.Println(findSwapValues([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // [9 9]

    fmt.Println(findSwapValues1([]int{4, 1, 2, 1, 1, 2}, []int{3, 6, 3, 3})) // [1, 3]
    fmt.Println(findSwapValues1([]int{1, 2, 3}, []int{4, 5, 6})) // []
    fmt.Println(findSwapValues1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // [1 1]
    fmt.Println(findSwapValues1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // [1 1]
    fmt.Println(findSwapValues1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // [9 9]
    fmt.Println(findSwapValues1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // [9 9]
}