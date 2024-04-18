package main

// 88. Merge Sorted Array
// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, 
// and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

// Merge nums1 and nums2 into a single array sorted in non-decreasing order.

// The final sorted array should not be returned by the function, but instead be stored inside the array nums1. 
// To accommodate this, nums1 has a length of m + n,
// where the first m elements denote the elements that should be merged, 
// and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

// Example 1:
// Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// Output: [1,2,2,3,5,6]
// Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
// The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.

// Example 2:
// Input: nums1 = [1], m = 1, nums2 = [], n = 0
// Output: [1]
// Explanation: The arrays we are merging are [1] and [].
// The result of the merge is [1].

// Example 3:
// Input: nums1 = [0], m = 0, nums2 = [1], n = 1
// Output: [1]
// Explanation: The arrays we are merging are [] and [1].
// The result of the merge is [1].
// Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.

// Constraints:
//     nums1.length == m + n
//     nums2.length == n
//     0 <= m, n <= 200
//     1 <= m + n <= 200
//     -10^9 <= nums1[i], nums2[j] <= 10^9
 
// Follow up: Can you come up with an algorithm that runs in O(m + n) time?

import "fmt"
import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
    for i := m + n; m > 0 && n > 0; i-- {
        if nums1[m-1] <= nums2[n-1] {
            nums1[i-1] = nums2[n-1]
            n--
        } else {
            nums1[i-1] = nums1[m-1]
            m--
        }
    }
    for ; n > 0; n-- {
        nums1[n-1] = nums2[n-1]
    }
}

// best solution
func merge1(nums1 []int, m int, nums2 []int, n int) {
    m1, i1 ,i2 := len(nums1) - 1, m - 1, n - 1
    for m1 >= 0 && i2 >= 0 {
        if (i1 < 0 ) ||  (nums1[i1] < nums2[i2]) {
            nums1[m1] = nums2[i2]
            i2--
            m1--
        } else {
            nums1[m1] = nums1[i1]
            i1--
            m1--
        }
    }
}

func merge2(nums1 []int, m int, nums2 []int, n int)  {
    copy(nums1[m:], nums2)
    sort.Ints(nums1)
}

func main() { 
    // Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
    // The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
    nums11 := []int{1,2,3,0,0,0}
    nums12 := []int{2,5,6}
    fmt.Println("nums11: ", nums11)
    fmt.Println("nums12: ", nums12)
    merge(nums11, 3, nums12, 3)
    fmt.Println("merge after nums11: ", nums11) // [1,2,2,3,5,6]

    // Explanation: The arrays we are merging are [1] and [].
    // The result of the merge is [1].
    nums21 := []int{1}
    nums22 := []int{}
    fmt.Println("nums21: ", nums21)
    fmt.Println("nums22: ", nums22)
    merge(nums21, 1, nums22, 0)
    fmt.Println("merge after nums21: ", nums21) // [1]

    // Explanation: The arrays we are merging are [] and [1].
    // The result of the merge is [1].
    // Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.
    nums31 := []int{0}
    nums32 := []int{1}
    fmt.Println("nums31: ", nums31)
    fmt.Println("nums32: ", nums32)
    merge(nums31, 0, nums32, 1)
    fmt.Println("merge after nums31: ", nums31) // [1]

    nums111 := []int{1,2,3,0,0,0}
    nums112 := []int{2,5,6}
    fmt.Println("nums111: ", nums111)
    fmt.Println("nums112: ", nums112)
    merge(nums111, 3, nums112, 3)
    fmt.Println("merge after nums111: ", nums111) // [1,2,2,3,5,6]

    nums121 := []int{1}
    nums122 := []int{}
    fmt.Println("nums121: ", nums121)
    fmt.Println("nums122: ", nums122)
    merge(nums121, 1, nums122, 0)
    fmt.Println("merge after nums121: ", nums121) // [1]

    nums131 := []int{0}
    nums132 := []int{1}
    fmt.Println("nums131: ", nums131)
    fmt.Println("nums132: ", nums132)
    merge(nums131, 0, nums132, 1)
    fmt.Println("merge after nums131: ", nums131) // [1]

    nums211 := []int{1,2,3,0,0,0}
    nums212 := []int{2,5,6}
    fmt.Println("nums211: ", nums211)
    fmt.Println("nums212: ", nums212)
    merge(nums211, 3, nums212, 3)
    fmt.Println("merge after nums211: ", nums211) // [1,2,2,3,5,6]

    nums221 := []int{1}
    nums222 := []int{}
    fmt.Println("nums221: ", nums221)
    fmt.Println("nums222: ", nums222)
    merge(nums221, 1, nums222, 0)
    fmt.Println("merge after nums221: ", nums221) // [1]

    nums231 := []int{0}
    nums232 := []int{1}
    fmt.Println("nums231: ", nums231)
    fmt.Println("nums232: ", nums232)
    merge(nums231, 0, nums232, 1)
    fmt.Println("merge after nums231: ", nums231) // [1]
}
