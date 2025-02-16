package main

// 3113. Find the Number of Subarrays Where Boundary Elements Are Maximum
// You are given an array of positive integers nums.

// Return the number of subarrays of nums, 
// where the first and the last elements of the subarray are equal to the largest element in the subarray.

// Example 1:
// Input: nums = [1,4,3,3,2]
// Output: 6
// Explanation:
// There are 6 subarrays which have the first and the last elements equal to the largest element of the subarray:
// subarray [1,4,3,3,2], with its largest element 1. The first element is 1 and the last element is also 1.
// subarray [1,4,3,3,2], with its largest element 4. The first element is 4 and the last element is also 4.
// subarray [1,4,3,3,2], with its largest element 3. The first element is 3 and the last element is also 3.
// subarray [1,4,3,3,2], with its largest element 3. The first element is 3 and the last element is also 3.
// subarray [1,4,3,3,2], with its largest element 2. The first element is 2 and the last element is also 2.
// subarray [1,4,3,3,2], with its largest element 3. The first element is 3 and the last element is also 3.
// Hence, we return 6.

// Example 2:
// Input: nums = [3,3,3]
// Output: 6
// Explanation:
// There are 6 subarrays which have the first and the last elements equal to the largest element of the subarray:
// subarray [3,3,3], with its largest element 3. The first element is 3 and the last element is also 3.
// subarray [3,3,3], with its largest element 3. The first element is 3 and the last element is also 3.
// subarray [3,3,3], with its largest element 3. The first element is 3 and the last element is also 3.
// subarray [3,3,3], with its largest element 3. The first element is 3 and the last element is also 3.
// subarray [3,3,3], with its largest element 3. The first element is 3 and the last element is also 3.
// subarray [3,3,3], with its largest element 3. The first element is 3 and the last element is also 3.
// Hence, we return 6.

// Example 3:
// Input: nums = [1]
// Output: 1
// Explanation:
// There is a single subarray of nums which is [1], with its largest element 1. The first element is 1 and the last element is also 1.
// Hence, we return 1.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"

func numberOfSubarrays(nums []int) int64 {
    res := 0
    stack := [][]int{}
    for _, v := range nums {
        for len(stack) > 0 && stack[len(stack) - 1][0] < v {
            stack = stack[:len(stack) - 1]
        }
        if len(stack) == 0 || stack[len(stack) - 1][0] != v {
            stack = append(stack, []int{ v, 1 })
        } else {
            stack[len(stack) - 1][1]++
        }
        res += stack[len(stack) - 1][1]
    }
    return int64(res)
}

func numberOfSubarrays1(nums []int) int64 {
    // 单调递减栈
    res := len(nums)
    type Pair struct { val, count int }
    // 加入哨兵
    stack := []Pair{{ 1 << 31, 0 }}
    for _, v := range nums {
        // 如果当前v大于栈顶 那么栈顶元素在后面不可能构成
        for v > stack[len(stack) - 1].val {
            stack = stack[:len(stack) - 1]
        }
        if v == stack[len(stack) - 1].val {
            res += stack[len(stack) - 1].count
            stack[len(stack) - 1].count++
        } else {
            stack = append(stack, Pair{ v, 1 })
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [1,4,3,3,2]
    // Output: 6
    // Explanation:
    // There are 6 subarrays which have the first and the last elements equal to the largest element of the subarray:
    // subarray [1,4,3,3,2], with its largest element 1. The first element is 1 and the last element is also 1.
    // subarray [1,4,3,3,2], with its largest element 4. The first element is 4 and the last element is also 4.
    // subarray [1,4,3,3,2], with its largest element 3. The first element is 3 and the last element is also 3.
    // subarray [1,4,3,3,2], with its largest element 3. The first element is 3 and the last element is also 3.
    // subarray [1,4,3,3,2], with its largest element 2. The first element is 2 and the last element is also 2.
    // subarray [1,4,3,3,2], with its largest element 3. The first element is 3 and the last element is also 3.
    // Hence, we return 6.
    fmt.Println(numberOfSubarrays([]int{1,4,3,3,2})) // 6
    // Example 2:
    // Input: nums = [3,3,3]
    // Output: 6
    // Explanation:
    // There are 6 subarrays which have the first and the last elements equal to the largest element of the subarray:
    // subarray [3,3,3], with its largest element 3. The first element is 3 and the last element is also 3.
    // subarray [3,3,3], with its largest element 3. The first element is 3 and the last element is also 3.
    // subarray [3,3,3], with its largest element 3. The first element is 3 and the last element is also 3.
    // subarray [3,3,3], with its largest element 3. The first element is 3 and the last element is also 3.
    // subarray [3,3,3], with its largest element 3. The first element is 3 and the last element is also 3.
    // subarray [3,3,3], with its largest element 3. The first element is 3 and the last element is also 3.
    // Hence, we return 6.
    fmt.Println(numberOfSubarrays([]int{3,3,3})) // 6
    // Example 3:
    // Input: nums = [1]
    // Output: 1
    // Explanation:
    // There is a single subarray of nums which is [1], with its largest element 1. The first element is 1 and the last element is also 1.
    // Hence, we return 1.
    fmt.Println(numberOfSubarrays([]int{1})) // 1

    fmt.Println(numberOfSubarrays([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(numberOfSubarrays([]int{9,8,7,6,5,4,3,2,1})) // 9

    fmt.Println(numberOfSubarrays1([]int{1,4,3,3,2})) // 6
    fmt.Println(numberOfSubarrays1([]int{3,3,3})) // 6
    fmt.Println(numberOfSubarrays1([]int{1})) // 1
    fmt.Println(numberOfSubarrays1([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(numberOfSubarrays1([]int{9,8,7,6,5,4,3,2,1})) // 9
}