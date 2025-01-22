package main

// 2570. Merge Two 2D Arrays by Summing Values
// You are given two 2D integer arrays nums1 and nums2.
//     nums1[i] = [idi, vali] indicate that the number with the id idi has a value equal to vali.
//     nums2[i] = [idi, vali] indicate that the number with the id idi has a value equal to vali.

// Each array contains unique ids and is sorted in ascending order by id.

// Merge the two arrays into one array that is sorted in ascending order by id, respecting the following conditions:
//     Only ids that appear in at least one of the two arrays should be included in the resulting array.
//     Each id should be included only once and its value should be the sum of the values of this id in the two arrays. If the id does not exist in one of the two arrays then its value in that array is considered to be 0.

// Return the resulting array. The returned array must be sorted in ascending order by id.

// Example 1:
// Input: nums1 = [[1,2],[2,3],[4,5]], nums2 = [[1,4],[3,2],[4,1]]
// Output: [[1,6],[2,3],[3,2],[4,6]]
// Explanation: The resulting array contains the following:
// - id = 1, the value of this id is 2 + 4 = 6.
// - id = 2, the value of this id is 3.
// - id = 3, the value of this id is 2.
// - id = 4, the value of this id is 5 + 1 = 6.

// Example 2:
// Input: nums1 = [[2,4],[3,6],[5,5]], nums2 = [[1,3],[4,3]]
// Output: [[1,3],[2,4],[3,6],[4,3],[5,5]]
// Explanation: There are no common ids, so we just include each id with its value in the resulting list.

// Constraints:
//     1 <= nums1.length, nums2.length <= 200
//     nums1[i].length == nums2[j].length == 2
//     1 <= idi, vali <= 1000
//     Both arrays contain unique ids.
//     Both arrays are in strictly ascending order by id.

import "fmt"

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
    res, i, j := [][]int{}, 0, 0
    for i < len(nums1) && j < len(nums2) {
        if nums1[i][0] < nums2[j][0] {
            res = append(res, nums1[i])
            i++
        } else if nums1[i][0] > nums2[j][0] {
            res = append(res, nums2[j])
            j++
        } else {
            res = append(res, []int{ nums1[i][0], nums1[i][1] + nums2[j][1]})
            i++
            j++
        }
    }
    for i < len(nums1) { // 处理 nums1 剩余
        res = append(res, nums1[i])
        i++
    }
    for j < len(nums2) { // 处理 nums2 剩余
        res = append(res, nums2[j])
        j++
    }
    return res
}

func mergeArrays1(nums1 [][]int, nums2 [][]int) [][]int {
    res, i, j := [][]int{}, 0, 0
    for i < len(nums1) && j < len(nums2) {
        if nums1[i][0] < nums2[j][0] {
            res = append(res, nums1[i])
            i++
        } else if nums1[i][0] > nums2[j][0] {
            res = append(res, nums2[j])
            j++
        } else {
            res = append(res, []int{ nums1[i][0], nums1[i][1] + nums2[j][1]})
            i++
            j++
        }
    }
    if i >= len(nums1) {
        res = append(res, nums2[j:]...)
    } else {
        res = append(res, nums1[i:]...)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums1 = [[1,2],[2,3],[4,5]], nums2 = [[1,4],[3,2],[4,1]]
    // Output: [[1,6],[2,3],[3,2],[4,6]]
    // Explanation: The resulting array contains the following:
    // - id = 1, the value of this id is 2 + 4 = 6.
    // - id = 2, the value of this id is 3.
    // - id = 3, the value of this id is 2.
    // - id = 4, the value of this id is 5 + 1 = 6.
    fmt.Println(mergeArrays([][]int{{1,2},{2,3},{4,5}}, [][]int{{1,4},{3,2},{4,1}})) // [[1,6],[2,3],[3,2],[4,6]]
    // Example 2:
    // Input: nums1 = [[2,4],[3,6],[5,5]], nums2 = [[1,3],[4,3]]
    // Output: [[1,3],[2,4],[3,6],[4,3],[5,5]]
    // Explanation: There are no common ids, so we just include each id with its value in the resulting list.
    fmt.Println(mergeArrays([][]int{{2,4},{3,6},{5,5}}, [][]int{{1,3},{4,3}})) // [[1,3],[2,4],[3,6],[4,3],[5,5]]

    fmt.Println(mergeArrays1([][]int{{1,2},{2,3},{4,5}}, [][]int{{1,4},{3,2},{4,1}})) // [[1,6],[2,3],[3,2],[4,6]]
    fmt.Println(mergeArrays1([][]int{{2,4},{3,6},{5,5}}, [][]int{{1,3},{4,3}})) // [[1,3],[2,4],[3,6],[4,3],[5,5]]
}