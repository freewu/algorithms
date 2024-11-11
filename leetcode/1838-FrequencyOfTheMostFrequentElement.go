package main

// 1838. Frequency of the Most Frequent Element
// The frequency of an element is the number of times it occurs in an array.

// You are given an integer array nums and an integer k. 
// In one operation, you can choose an index of nums and increment the element at that index by 1.

// Return the maximum possible frequency of an element after performing at most k operations.

// Example 1:
// Input: nums = [1,2,4], k = 5
// Output: 3
// Explanation: Increment the first element three times and the second element two times to make nums = [4,4,4].
// 4 has a frequency of 3.

// Example 2:
// Input: nums = [1,4,8,13], k = 5
// Output: 2
// Explanation: There are multiple optimal solutions:
// - Increment the first element three times to make nums = [4,4,8,13]. 4 has a frequency of 2.
// - Increment the second element four times to make nums = [1,8,8,13]. 8 has a frequency of 2.
// - Increment the third element five times to make nums = [1,4,13,13]. 13 has a frequency of 2.

// Example 3:
// Input: nums = [3,9,6], k = 2
// Output: 1

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5
//     1 <= k <= 10^5

import "fmt"
import "sort"
//import "slices" 

func maxFrequency(nums []int, k int) int {
    sort.Ints(nums)
    res, sum, n := 1, 0, len(nums)
    pivot, head, freq := n - 1, n - 2, 1
    for {
        if head < 0 { return res }
        if k >= sum {
            sum = sum + nums[pivot] - nums[head]
            head--
            freq++
        }
        if k >= sum {
            if freq > res { res = freq }
        } else if k < sum {
            freq--
            sum = sum - (freq * (nums[pivot] - nums[pivot - 1]))
            pivot--
        }
    }
    return res 
}

func maxFrequency1(nums []int, k int) int {
    n := len(nums) 
    if n == 1 { return 1 } // corner-case 
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    countingSort := func(arr []int) { // 计数排序
        mn, mx := arr[0], arr[0] 
        for _, v := range arr {
            mn, mx = min(mn, v), max(mx, v) 
        }
        count := make([]int, mx + 1) 
        for _, v := range arr {
            count[v]++
        }
        index := 0 
        for i := mn; i <= mx; i++ {
            for count[i] > 0 {
                arr[index] = i
                index++
                count[i]--
            }
        }
    }
    countingSort(nums) // slices.Sort(nums) 
    res, sum, start := 1, nums[0] , 0 // window窗口的左端点=  
    for i := 1; i < n; i++ {
        v := nums[i]; 
        sum += v
        f := (i - start + 1) 
        if sum + k >= v * f { // 补足k个位置后，最高频率能够达到f个 v
            res = max(res, f) 
        } else {
            for start < i {
                sum -= nums[start]; 
                start++
                f--
                if sum + k >= v * f { 
                    res = max(res, f) 
                    break 
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,4], k = 5
    // Output: 3
    // Explanation: Increment the first element three times and the second element two times to make nums = [4,4,4].
    // 4 has a frequency of 3.
    fmt.Println(maxFrequency([]int{1,2,4}, 5)) // 3
    // Example 2:
    // Input: nums = [1,4,8,13], k = 5
    // Output: 2
    // Explanation: There are multiple optimal solutions:
    // - Increment the first element three times to make nums = [4,4,8,13]. 4 has a frequency of 2.
    // - Increment the second element four times to make nums = [1,8,8,13]. 8 has a frequency of 2.
    // - Increment the third element five times to make nums = [1,4,13,13]. 13 has a frequency of 2.
    fmt.Println(maxFrequency([]int{1,4,8,13}, 5)) // 2
    // Example 3:
    // Input: nums = [3,9,6], k = 2
    // Output: 1
    fmt.Println(maxFrequency([]int{3,9,6}, 2)) // 1

    fmt.Println(maxFrequency1([]int{1,2,4}, 5)) // 3
    fmt.Println(maxFrequency1([]int{1,4,8,13}, 5)) // 2
    fmt.Println(maxFrequency1([]int{3,9,6}, 2)) // 1
}