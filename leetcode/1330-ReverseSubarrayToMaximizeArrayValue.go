package main

// 1330. Reverse Subarray To Maximize Array Value
// You are given an integer array nums. 
// The value of this array is defined as the sum of |nums[i] - nums[i + 1]| for all 0 <= i < nums.length - 1.

// You are allowed to select any subarray of the given array and reverse it. 
// You can perform this operation only once.

// Find maximum possible value of the final array.

// Example 1:
// Input: nums = [2,3,1,5,4]
// Output: 10
// Explanation: By reversing the subarray [3,1,5] the array becomes [2,5,1,3,4] whose value is 10.

// Example 2:
// Input: nums = [2,4,9,24,2,1,10]
// Output: 68

// Constraints:
//     2 <= nums.length <= 3 * 10^4
//     -10^5 <= nums[i] <= 10^5
//     The answer is guaranteed to fit in a 32-bit integer.

import "fmt"

func maxValueAfterReverse(nums []int) int {
    res, n, mx, mn := 0, len(nums), -1 << 31, 1 << 31
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    //it is only possible to increase when in a, b ... c, d
    //the distance between min(c,d) and max(a, b) is positive
    //then will increase by 2*(min(c, d) - max(a, b)
    for i := 0; i < n - 1; i++ {
        l, r := i, i + 1
        mx = max(mx, min(nums[l], nums[r]))
        mn = min(mn, max(nums[l], nums[r]))
    }
    change := max(0, (mx - mn)*2)
    //another situation is reverse from nums[0] or reverse from nums[-1]
    //if reverse from nums[0] to nums[i]
    //change = -abs(nums[i]-nums[i+1])+abs(nums[i+1]-nums[0])
    //if reverse from nums[i] to nums[-1]
    //change = -abs(nums[i]-nums[i-1])+abs(nums[-1]-nums[i-1])
    head, tail := -1 << 31, -1 << 31
    for i := 0; i < n - 1; i++ {
        head = max(head, -abs(nums[i] - nums[i+1]) + abs(nums[i+1] - nums[0]))
        tail = max(tail, -abs(nums[i+1] - nums[i]) + abs(nums[i] - nums[n -1]))
    }
    change = max(max(tail, head), change)
    for i := 0; i < n - 1; i++ {
        res += abs(nums[i] - nums[i+1])
    }
    return res + change
}

func maxValueAfterReverse1(nums []int) int {
    base, d, n, mx, mn := 0, 0, len(nums), -1 << 31, 1 << 31
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < n; i++ {
        a, b := nums[i-1], nums[i]
        base += abs(a - b)
        mx = max(mx, min(a, b))
        mn = min(mn, max(a, b))
        d = max(d, max(abs(nums[0]-b) - abs(a-b), // i = 0
            abs(nums[n-1]-a) - abs(a-b))) // j = n-1
    }
    return base + max(d, 2 * (mx - mn))
}

func main() {
    // Example 1:
    // Input: nums = [2,3,1,5,4]
    // Output: 10
    // Explanation: By reversing the subarray [3,1,5] the array becomes [2,5,1,3,4] whose value is 10.
    fmt.Println(maxValueAfterReverse([]int{2,3,1,5,4})) // 10
    // Example 2:
    // Input: nums = [2,4,9,24,2,1,10]
    // Output: 68
    // Explanation: By reversing the subarray [3,1,5] the array becomes [2,5,1,3,4] whose value is 10.
    fmt.Println(maxValueAfterReverse([]int{2,4,9,24,2,1,10})) // 68

    fmt.Println(maxValueAfterReverse1([]int{2,3,1,5,4})) // 10
    fmt.Println(maxValueAfterReverse1([]int{2,4,9,24,2,1,10})) // 68
}