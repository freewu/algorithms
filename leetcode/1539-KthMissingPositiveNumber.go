package main

// 1539. Kth Missing Positive Number
// Given an array arr of positive integers sorted in a strictly increasing order, and an integer k.
// Return the kth positive integer that is missing from this array.

// Example 1:
// Input: arr = [2,3,4,7,11], k = 5
// Output: 9
// Explanation: The missing positive integers are [1,5,6,8,9,10,12,13,...]. The 5th missing positive integer is 9.

// Example 2:
// Input: arr = [1,2,3,4], k = 2
// Output: 6
// Explanation: The missing positive integers are [5,6,7,...]. The 2nd missing positive integer is 6.

// Constraints:
//     1 <= arr.length <= 1000
//     1 <= arr[i] <= 1000
//     1 <= k <= 1000
//     arr[i] < arr[j] for 1 <= i < j <= arr.length

// Follow up:
//     Could you solve this problem in less than O(n) complexity?

import "fmt"
import "sort"

func findKthPositive(arr []int, k int) int {
    nums, missing := []int{}, []int{}
    m := make(map[int]bool)
    // 生成一个正常的数组 长为 arr[len(arr) - 1] 最大值 + k
    for i := 1; i <= arr[len(arr)-1] + k; i++ {
        nums = append(nums, i)
    }
    for _, v := range arr {
        m[v] = true
    }
    for _, v := range nums {
        // 不存在 arr 中 就添加
        if m[v] == false {
            missing = append(missing, v)
        }
    }
    // 返回缺失数组的第 k 位
    return missing[k -1]
}

func findKthPositive1(arr []int, k int) int {
    arr = append(arr, 2024)
    index := sort.Search(len(arr), func(i int) bool {
        return arr[i] - i - 1 >= k
    })
    return index + k
}

func main() {
    // Explanation: The missing positive integers are [1,5,6,8,9,10,12,13,...]. The 5th missing positive integer is 9.
    fmt.Println(findKthPositive([]int{2,3,4,7,11}, 5)) // 9
    // Explanation: The missing positive integers are [5,6,7,...]. The 2nd missing positive integer is 6.
    fmt.Println(findKthPositive([]int{1,2,3,4}, 2)) // 6

    fmt.Println(findKthPositive1([]int{2,3,4,7,11}, 5)) // 9
    fmt.Println(findKthPositive1([]int{1,2,3,4}, 2)) // 6
}