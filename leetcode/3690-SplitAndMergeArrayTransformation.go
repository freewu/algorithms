package main

// 3690. Split and Merge Array Transformation
// You are given two integer arrays nums1 and nums2, each of length n. 
// You may perform the following split-and-merge operation on nums1 any number of times:
//     1. Choose a subarray nums1[L..R].
//     2. Remove that subarray, leaving the prefix nums1[0..L-1] (empty if L = 0) and the suffix nums1[R+1..n-1] (empty if R = n - 1).
//     3. Re-insert the removed subarray (in its original order) at any position in the remaining array 
//        (i.e., between any two elements, at the very start, or at the very end).

// Return the minimum number of split-and-merge operations needed to transform nums1 into nums2.

// Example 1:
// Input: nums1 = [3,1,2], nums2 = [1,2,3]
// Output: 1
// Explanation:
// Split out the subarray [3] (L = 0, R = 0); the remaining array is [1,2].
// Insert [3] at the end; the array becomes [1,2,3].

// Example 2:
// Input: nums1 = [1,1,2,3,4,5], nums2 = [5,4,3,2,1,1]
// Output: 3
// Explanation:
// Remove [1,1,2] at indices 0 - 2; remaining is [3,4,5]; insert [1,1,2] at position 2, resulting in [3,4,1,1,2,5].
// Remove [4,1,1] at indices 1 - 3; remaining is [3,2,5]; insert [4,1,1] at position 3, resulting in [3,2,5,4,1,1].
// Remove [3,2] at indices 0 - 1; remaining is [5,4,1,1]; insert [3,2] at position 2, resulting in [5,4,3,2,1,1].
 
// Constraints:
//     2 <= n == nums1.length == nums2.length <= 6
//     -10^5 <= nums1[i], nums2[i] <= 10^5
//     nums2 is a permutation of nums1.

import "fmt"
import "slices"

func minSplitMerge(nums1, nums2 []int) int {
    res, n := 0, len(nums1)
    t := [6]int{}
    for j, v := range nums1 {
        t[j] = v
    }
    visited, queue := map[[6]int]bool{t: true}, [][]int{ nums1 }
    for ; ; res++ {
        temp := queue
        queue = nil
        for _, a := range temp {
            if slices.Equal(a, nums2) { return res }
            for l := 0; l < n; l++ {
                for r := l + 1; r <= n; r++ {
                    b := slices.Clone(a)
                    sub := slices.Clone(b[l:r])
                    b = append(b[:l], b[r:]...) // 从 b 中移除 sub
                    for i := 0; i <= len(b); i++ {
                        c := slices.Insert(slices.Clone(b), i, sub...)
                        t := [6]int{}
                        for j, v := range c {
                            t[j] = v
                        }
                        if !visited[t] {
                            visited[t] = true
                            queue = append(queue, c)
                        }
                    }
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums1 = [3,1,2], nums2 = [1,2,3]
    // Output: 1
    // Explanation:
    // Split out the subarray [3] (L = 0, R = 0); the remaining array is [1,2].
    // Insert [3] at the end; the array becomes [1,2,3].
    fmt.Println(minSplitMerge([]int{3,1,2}, []int{1,2,3})) // 1
    // Example 2:
    // Input: nums1 = [1,1,2,3,4,5], nums2 = [5,4,3,2,1,1]
    // Output: 3
    // Explanation:
    // Remove [1,1,2] at indices 0 - 2; remaining is [3,4,5]; insert [1,1,2] at position 2, resulting in [3,4,1,1,2,5].
    // Remove [4,1,1] at indices 1 - 3; remaining is [3,2,5]; insert [4,1,1] at position 3, resulting in [3,2,5,4,1,1].
    // Remove [3,2] at indices 0 - 1; remaining is [5,4,1,1]; insert [3,2] at position 2, resulting in [5,4,3,2,1,1].
    fmt.Println(minSplitMerge([]int{1,1,2,3,4,5}, []int{5,4,3,2,1,1})) // 3

    fmt.Println(minSplitMerge([]int{1,2,3,4,5,6}, []int{1,2,3,4,5,6})) // 0
    fmt.Println(minSplitMerge([]int{1,2,3,4,5,6}, []int{6,5,4,3,2,1})) // 4
    fmt.Println(minSplitMerge([]int{6,5,4,3,2,1}, []int{1,2,3,4,5,6})) // 4
    fmt.Println(minSplitMerge([]int{6,5,4,3,2,1}, []int{6,5,4,3,2,1})) // 0
}