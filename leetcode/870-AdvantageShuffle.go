package main

// 870. Advantage Shuffle
// You are given two integer arrays nums1 and nums2 both of the same length. 
// The advantage of nums1 with respect to nums2 is the number of indices i for which nums1[i] > nums2[i].

// Return any permutation of nums1 that maximizes its advantage with respect to nums2.

// Example 1:
// Input: nums1 = [2,7,11,15], nums2 = [1,10,4,11]
// Output: [2,11,7,15]

// Example 2:
// Input: nums1 = [12,24,8,32], nums2 = [13,25,32,11]
// Output: [24,32,8,12]

// Constraints:
//     1 <= nums1.length <= 10^5
//     nums2.length == nums1.length
//     0 <= nums1[i], nums2[i] <= 10^9

import "fmt"
import "sort"
import "slices"

func advantageCount(nums1, nums2 []int) []int {
    sort.Ints(nums1)
    arr := make([][2]int, len(nums2))
    for i, v := range nums2 {
        arr[i] = [2]int{v, i}
    }
    sort.Slice(arr, func(i, j int) bool {
        return arr[i][0] < arr[j][0]
    })
    res := make([]int, len(nums1))
    left, right := 0, len(nums1) - 1
    for i := len(arr) - 1; i >= 0; i-- {
        v, index := arr[i][0], arr[i][1]
        if nums1[right] > v {
            res[index] = nums1[right]
            right--
        } else {
            res[index] = nums1[left]
            left++
        }
    }
    return res
}

func advantageCount1(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    n := len(nums1)
    ids := make([]int, n)
    for i := range ids {
        ids[i] = i
    }
    //ids索引根据nums的值进行排序，ids[0]就是nums2的最小值索引
    slices.SortFunc(ids, func(i, j int) int { 
        return nums2[i] - nums2[j] 
    })
    res, left, right := make([]int, n), 0, n-1
    for _, x := range nums1 {
        if x > nums2[ids[left]] { // 用下等马比下等马
            res[ids[left]] = x
            left++
        } else { // 用下等马比上等马
            res[ids[right]] = x
            right--
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums1 = [2,7,11,15], nums2 = [1,10,4,11]
    // Output: [2,11,7,15]
    fmt.Println(advantageCount([]int{2,7,11,15}, []int{1,10,4,11})) // [2,11,7,15]
    // Example 2:
    // Input: nums1 = [12,24,8,32], nums2 = [13,25,32,11]
    // Output: [24,32,8,12]
    fmt.Println(advantageCount([]int{12,24,8,32}, []int{13,25,32,11})) // [24,32,8,12]

    fmt.Println(advantageCount1([]int{2,7,11,15}, []int{1,10,4,11})) // [2,11,7,15]
    fmt.Println(advantageCount1([]int{12,24,8,32}, []int{13,25,32,11})) // [24,32,8,12]
}