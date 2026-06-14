package main

// 3960. Frequency Balance Subarray
// You are given an integer array ​​​​​​​nums.

// Define a frequency balance subarray as follows:
//     1. If the subarray contains only one element, it is frequency balanced.
//     2. If the subarray contains at least two elements, 
//         then every element with maximum frequency must occur exactly twice as many times as every other distinct value in that subarray.

// Return an integer denoting the length of the longest frequency balance subarray.

// A subarray is a contiguous non-empty​​​​​​​ sequence of elements within an array.

// The frequency of an element x is the number of times it occurs in the array.
 
// Example 1:
// Input: nums = [1,2,2,1,2,3,3,3]
// Output: 5
// Explanation:
// The longest frequency balance subarray is [2, 1, 2, 3, 3].
// The elements that appear most frequently are 2 and 3, both appearing twice.
// The remaining element 1 appears once, meeting the requirements.

// Example 2:
// Input: nums = [5,5,5,5]
// Output: 4
// Explanation:
// The longest frequency balance subarray is [5, 5, 5, 5].
// The element that appears most frequently is 5.
// There are no other elements meeting the requirements.

// Example 3:
// Input: nums = [1,2,3,4]
// Output: 1
// Explanation:
// Since all elements appear only once, the length of the longest frequency balance subarray is 1.

// Constraints:
//     1 <= nums.length <= 10​​​​​​​^3
//     1 <= nums[i] <= 10​​​​​​​^9

import "fmt"

func getLength(nums []int) int {
    res := 0
    for i := 0; i < len(nums); i++ {
        mx, count, freq := 0, make(map[int]int), make(map[int]int)
        for j := i; j < len(nums); j++ {
            if freq[count[nums[j]]] > 0 {
                freq[count[nums[j]]]--
            }
            if freq[count[nums[j]]] == 0 {delete(freq, count[nums[j]])}
            count[nums[j]]++
            freq[count[nums[j]]]++
            mx = max(mx, count[nums[j]])
            if len(freq) >= 3 { 
                continue
            }
            if len(count) == 1 {
                res = max(res, j-i+1) 
                continue
            }
            low, high := 1 << 61, -1 << 61
            for k := range freq {
                low, high = min(low, k), max(high, k)
            }
            if low * 2 == high {
                res = max(res, j - i + 1) 
            }
        }
    }
    return res
}

func getLength1(nums []int) int {
    res, n := 1, len(nums)
    if n == 1{
        return 1
    }
    mp := make(map[int]int)
    id := 0
    for _, v := range nums{
        if _, ok := mp[v]; !ok {
            mp[v] = id
            id++
        }
    }
    arr := make([]int, n)
    for i, v := range nums {
        arr[i] = mp[v]
    }
    freq, c := make([]int, id), make([]int, n + 1)
    for i  := range nums {
        if n - i <= res {
            break
        }
        f, df := 0, 0
        for j := i; j < n; j++ {
            x := arr[j]
            w := freq[x]
            if w > 0{
                c[w]--
                if c[w] == 0 {
                    df--
                }
            }
            w++
            freq[x] = w
            if c[w] == 0 {
                df++
            }
            c[w]++
            f = max(f, w)
            if j - i + 1 > res {
                if (df == 1 && c[f] == 1) || (df == 2 && f&1 == 0 && c[f >> 1] > 0) {
                    res = j - i + 1
                }
            }
        }
        for j := i; j < n; j++ {
            freq[arr[j]] = 0
        }
        clear(c[:f + 1])
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,2,1,2,3,3,3]
    // Output: 5
    // Explanation:
    // The longest frequency balance subarray is [2, 1, 2, 3, 3].
    // The elements that appear most frequently are 2 and 3, both appearing twice.
    // The remaining element 1 appears once, meeting the requirements.
    fmt.Println(getLength([]int{1,2,2,1,2,3,3,3})) // 5
    // Example 2:
    // Input: nums = [5,5,5,5]
    // Output: 4
    // Explanation:
    // The longest frequency balance subarray is [5, 5, 5, 5].
    // The element that appears most frequently is 5.
    // There are no other elements meeting the requirements.
    fmt.Println(getLength([]int{5,5,5,5})) // 4
    // Example 3:
    // Input: nums = [1,2,3,4]
    // Output: 1
    // Explanation:
    // Since all elements appear only once, the length of the longest frequency balance subarray is 1.
    fmt.Println(getLength([]int{1,2,3,4})) // 1

    fmt.Println(getLength([]int{1,2,3,4,5,6,7,8,9})) // 1
    fmt.Println(getLength([]int{9,8,7,6,5,4,3,2,1})) // 1

    fmt.Println(getLength1([]int{1,2,2,1,2,3,3,3})) // 5
    fmt.Println(getLength1([]int{5,5,5,5})) // 4
    fmt.Println(getLength1([]int{1,2,3,4})) // 1
    fmt.Println(getLength1([]int{1,2,3,4,5,6,7,8,9})) // 1
    fmt.Println(getLength1([]int{9,8,7,6,5,4,3,2,1})) // 1
}