package main

// 2540. Minimum Common Value
// Given two integer arrays nums1 and nums2, sorted in non-decreasing order, return the minimum integer common to both arrays. If there is no common integer amongst nums1 and nums2, return -1.
// Note that an integer is said to be common to nums1 and nums2 if both arrays have at least one occurrence of that integer.

// Example 1:
// Input: nums1 = [1,2,3], nums2 = [2,4]
// Output: 2
// Explanation: The smallest element common to both arrays is 2, so we return 2.

// Example 2:
// Input: nums1 = [1,2,3,6], nums2 = [2,3,4,5]
// Output: 2
// Explanation: There are two common elements in the array 2 and 3 out of which 2 is the smallest, so 2 is returned.
 
// Constraints:
//         1 <= nums1.length, nums2.length <= 10^5
//         1 <= nums1[i], nums2[j] <= 10^9
//         Both nums1 and nums2 are sorted in non-decreasing order.

import "fmt"

func getCommon(nums1 []int, nums2 []int) int {
    i, j,l1,l2 := 0, 0, len(nums1) - 1, len(nums2) - 1
    for {
        // 已经有一个结束了
        if i > l1 || j > l2 {
            break
        }
        // 小一边的要追上才能比对
        if nums1[i] > nums2[j] { 
            j++
        } else if nums1[i] < nums2[j]{
            i++
        } else { // 第一个发现相同的直接返回就是最小的共同数了
            return nums1[i]
        }
    }
    return -1
}

func main() {
    fmt.Println(getCommon(
        []int{1,2,3},
        []int{2,4},
    )) // 2
    fmt.Println(getCommon(
        []int{1,2,3,6},
        []int{2,3,4,5},
    )) // 2
    fmt.Println(getCommon(
        []int{1,3},
        []int{3,4},
    )) // 3
}