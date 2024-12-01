package main

// 2122. Recover the Original Array
// Alice had a 0-indexed array arr consisting of n positive integers. 
// She chose an arbitrary positive integer k and created two new 0-indexed integer arrays lower and higher in the following manner:
//     1. lower[i] = arr[i] - k, for every index i where 0 <= i < n
//     2. higher[i] = arr[i] + k, for every index i where 0 <= i < n

// Unfortunately, Alice lost all three arrays. 
// However, she remembers the integers that were present in the arrays lower and higher, but not the array each integer belonged to. 
// Help Alice and recover the original array.

// Given an array nums consisting of 2n integers, where exactly n of the integers were present in lower and the remaining in higher, return the original array arr. 
// In case the answer is not unique, return any valid array.

// Note: The test cases are generated such that there exists at least one valid array arr.

// Example 1:
// Input: nums = [2,10,6,4,8,12]
// Output: [3,7,11]
// Explanation:
// If arr = [3,7,11] and k = 1, we get lower = [2,6,10] and higher = [4,8,12].
// Combining lower and higher gives us [2,6,10,4,8,12], which is a permutation of nums.
// Another valid possibility is that arr = [5,7,9] and k = 3. In that case, lower = [2,4,6] and higher = [8,10,12]. 

// Example 2:
// Input: nums = [1,1,3,3]
// Output: [2,2]
// Explanation:
// If arr = [2,2] and k = 1, we get lower = [1,1] and higher = [3,3].
// Combining lower and higher gives us [1,1,3,3], which is equal to nums.
// Note that arr cannot be [1,3] because in that case, the only possible way to obtain [1,1,3,3] is with k = 0.
// This is invalid since k must be positive.

// Example 3:
// Input: nums = [5,435]
// Output: [220]
// Explanation:
// The only possible combination is arr = [220] and k = 215. Using them, we get lower = [5] and higher = [435].

// Constraints:
//     2 * n == nums.length
//     1 <= n <= 1000
//     1 <= nums[i] <= 10^9
//     The test cases are generated such that there exists at least one valid array arr.

import "fmt"
import "sort"

func recoverArray(nums []int) []int {
    sort.Ints(nums)
    n := len(nums)
    res, count, remaining := make([]int, n / 2), make(map[int]int), make(map[int]int)
    for _, v := range nums {
        count[v]++
    }
    for i := 1; i <= n / 2; i++ {
        if nums[i] == nums[i - 1] { continue }
        for i, c := range count {
            remaining[i] = c
        }
        k := nums[i] - nums[0]
        if k % 2 == 1 { continue }
        k /= 2
        rp := 0
        for j := 0; j < n; j++ {
            if remaining[nums[j]] == 0 { continue }
            candidate := nums[j] + 2 * k
            if remaining[candidate] == 0 {
                k = -1
                break
            }
            res[rp] = nums[j] + k
            rp++
            remaining[nums[j]]--
            remaining[candidate]--
        }
        if k > 0 { break }
    }
    return res
}

func recoverArray1(nums []int) []int {
    sort.Ints(nums)
    for i, n := 1, len(nums); ; i++ {
        if nums[i] == nums[i-1] { continue } // 优化：如果与上一个元素相同，那么我们会得到同样的 k，同样找不到原数组，此时应直接跳过
        d := nums[i] - nums[0] // 此时 d > 0 必然成立
        if d&1 > 0 { continue } // k 必须是整数
        k := d / 2
        visited := make([]bool, n) // 用来标记出现在 higher 中的数（用 nums 的下标）
        visited[i] = true
        res := []int{ (nums[0] + nums[i]) / 2 }
        for low, high := 0, i + 1; high < n; high++ { // 双指针：lo 指向 lower，hi 指向 higher
            for low++; visited[low]; low++ {} // 找 lower：跳过出现在 higher 中的数
            for ; high < n && nums[high] - nums[low] < 2 * k; high++ {} // 找 higher
            if high == n || nums[high] - nums[low] > 2 * k { break } // 不存在满足等式的 higher 值
            visited[high] = true
            res = append(res, (nums[low] + nums[high])/2) // 找到一对满足等式的 (lower, higher)
        }
        if len(res) == n / 2 { 
            return res 
        }
    }
}

func main() {
    // Example 1:
    // Input: nums = [2,10,6,4,8,12]
    // Output: [3,7,11]
    // Explanation:
    // If arr = [3,7,11] and k = 1, we get lower = [2,6,10] and higher = [4,8,12].
    // Combining lower and higher gives us [2,6,10,4,8,12], which is a permutation of nums.
    // Another valid possibility is that arr = [5,7,9] and k = 3. In that case, lower = [2,4,6] and higher = [8,10,12]. 
    fmt.Println(recoverArray([]int{2,10,6,4,8,12})) // [3,7,11]
    // Example 2:
    // Input: nums = [1,1,3,3]
    // Output: [2,2]
    // Explanation:
    // If arr = [2,2] and k = 1, we get lower = [1,1] and higher = [3,3].
    // Combining lower and higher gives us [1,1,3,3], which is equal to nums.
    // Note that arr cannot be [1,3] because in that case, the only possible way to obtain [1,1,3,3] is with k = 0.
    // This is invalid since k must be positive.
    fmt.Println(recoverArray([]int{1,1,3,3})) // [2,2]
    // Example 3:
    // Input: nums = [5,435]
    // Output: [220]
    // Explanation:
    // The only possible combination is arr = [220] and k = 215. Using them, we get lower = [5] and higher = [435].
    fmt.Println(recoverArray([]int{5,435})) // [220]

    fmt.Println(recoverArray1([]int{2,10,6,4,8,12})) // [3,7,11]
    fmt.Println(recoverArray1([]int{1,1,3,3})) // [2,2]
    fmt.Println(recoverArray1([]int{5,435})) // [220]
}