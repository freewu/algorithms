package main

// 1213. Intersection of Three Sorted Arrays
// Given three integer arrays arr1, arr2 and arr3 sorted in strictly increasing order, 
// return a sorted array of only the integers that appeared in all three arrays.

// Example 1:
// Input: arr1 = [1,2,3,4,5], arr2 = [1,2,5,7,9], arr3 = [1,3,4,5,8]
// Output: [1,5]
// Explanation: Only 1 and 5 appeared in the three arrays.

// Example 2:
// Input: arr1 = [197,418,523,876,1356], arr2 = [501,880,1593,1710,1870], arr3 = [521,682,1337,1395,1764]
// Output: []

// Constraints:
//     1 <= arr1.length, arr2.length, arr3.length <= 1000
//     1 <= arr1[i], arr2[i], arr3[i] <= 2000

import "fmt"

// 三指针
func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
    res, l1,l2,l3,n1,n2,n3 := []int{}, 0,0,0,len(arr1),len(arr2),len(arr3)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for l1 < n1 && l2 < n2 && l3 < n3 {
        if arr1[l1] == arr2[l2] && arr2[l2] == arr3[l3] { // 都相等
            res = append(res, arr1[l1])
            l1++
            l2++
            l3++
            continue
        }
        // 谁最小谁进位
        mn := min(arr1[l1], min(arr2[l2],arr3[l3]))
        if arr1[l1] == mn { l1++ }
        if arr2[l2] == mn { l2++ }
        if arr3[l3] == mn { l3++ }
    }   
    return res
}

func main() {
    // Example 1:
    // Input: arr1 = [1,2,3,4,5], arr2 = [1,2,5,7,9], arr3 = [1,3,4,5,8]
    // Output: [1,5]
    // Explanation: Only 1 and 5 appeared in the three arrays.
    fmt.Println(arraysIntersection([]int{1,2,3,4,5}, []int{1,2,5,7,9}, []int{1,3,4,5,8})) // [1,5]
    // Example 2:
    // Input: arr1 = [197,418,523,876,1356], arr2 = [501,880,1593,1710,1870], arr3 = [521,682,1337,1395,1764]
    // Output: []
    fmt.Println(arraysIntersection([]int{197,418,523,876,1356}, []int{501,880,1593,1710,1870}, []int{521,682,1337,1395,1764})) // []
}