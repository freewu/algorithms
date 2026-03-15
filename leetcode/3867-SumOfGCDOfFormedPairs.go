package main

// 3867. Sum of GCD of Formed Pairs
// You are given an integer array nums of length n.

// Construct an array prefixGcd where for each index i:
//     1. Let mxi = max(nums[0], nums[1], ..., nums[i]).
//     2. prefixGcd[i] = gcd(nums[i], mxi).

// After constructing prefixGcd:
//     1. Sort prefixGcd in non-decreasing order.
//     2. Form pairs by taking the smallest unpaired element and the largest unpaired element.
//     3. Repeat this process until no more pairs can be formed.
//     4. For each formed pair, compute the gcd of the two elements.
//     5. If n is odd, the middle element in the prefixGcd array remains unpaired and should be ignored.

// Return an integer denoting the sum of the GCD values of all formed pairs.

// The term gcd(a, b) denotes the greatest common divisor of a and b.
 
// Example 1:
// Input: nums = [2,6,4]
// Output: 2
// Explanation:
// Construct prefixGcd:
// i  | nums[i] | mxi | prefixGcd[i]
// 0  |    2    | 2   | 2
// 1  |    6    | 6   | 6
// 2  |    4    | 6   | 2
// prefixGcd = [2, 6, 2]. After sorting, it forms [2, 2, 6].
// Pair the smallest and largest elements: gcd(2, 6) = 2. The remaining middle element 2 is ignored. Thus, the sum is 2.

// Example 2:
// Input: nums = [3,6,2,8]
// Output: 5
// Explanation:
// Construct prefixGcd:
// i  | nums[i] | mxi | prefixGcd[i]
// 0  |    3    | 3   | 3
// 1  |    6    | 6   | 6
// 2  |    2    | 6   | 2
// 3  |    8    | 8   | 8
// prefixGcd = [3, 6, 2, 8]. After sorting, it forms [2, 3, 6, 8].
// Form pairs: gcd(2, 8) = 2 and gcd(3, 6) = 3. Thus, the sum is 2 + 3 = 5.

// Constraints:
//     1 <= n == nums.length <= 10^5
//     1 <= nums[i] <= 10​​​​​​​^9

import "fmt"
import "sort"

func gcdSum(nums []int) int64 {
    res, mx, n := int64(0), 0, len(nums)
    prefix := make([]int, n)
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    for i, v := range nums {
        mx = max(mx, v)
        prefix[i] = gcd(v, mx)
    }
    sort.Ints(prefix)
    for i := range n / 2 {
        res += int64(gcd(prefix[i], prefix[n-1-i]))
    }
    return res
}

func gcdSum1(nums []int) int64 {
    res, mx, n := 0, 0, len(nums)
    prefix := make([]int, n)
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    for i, num := range nums {
        mx = max(mx, num)
        prefix[i] = gcd(mx, num)
    }
    sort.Ints(prefix)
    l, r := 0, n-1
    for l < r {
        gd := gcd(prefix[l], prefix[r])
        l++
        r--
        res += gd
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [2,6,4]
    // Output: 2
    // Explanation:
    // Construct prefixGcd:
    // i  | nums[i] | mxi | prefixGcd[i]
    // 0  |    2    | 2   | 2
    // 1  |    6    | 6   | 6
    // 2  |    4    | 6   | 2
    // prefixGcd = [2, 6, 2]. After sorting, it forms [2, 2, 6].
    // Pair the smallest and largest elements: gcd(2, 6) = 2. The remaining middle element 2 is ignored. Thus, the sum is 2.
    fmt.Println(gcdSum([]int{2,6,4})) // 2
    // Example 2:
    // Input: nums = [3,6,2,8]
    // Output: 5
    // Explanation:
    // Construct prefixGcd:
    // i  | nums[i] | mxi | prefixGcd[i]
    // 0  |    3    | 3   | 3
    // 1  |    6    | 6   | 6
    // 2  |    2    | 6   | 2
    // 3  |    8    | 8   | 8
    // prefixGcd = [3, 6, 2, 8]. After sorting, it forms [2, 3, 6, 8].
    // Form pairs: gcd(2, 8) = 2 and gcd(3, 6) = 3. Thus, the sum is 2 + 3 = 5.
    fmt.Println(gcdSum([]int{3,6,2,8})) // 5

    fmt.Println(gcdSum([]int{1,2,3,4,5,6,7,8,9})) // 6
    fmt.Println(gcdSum([]int{9,8,7,6,5,4,3,2,1})) // 4

    fmt.Println(gcdSum1([]int{2,6,4})) // 2
    fmt.Println(gcdSum1([]int{3,6,2,8})) // 5
    fmt.Println(gcdSum1([]int{1,2,3,4,5,6,7,8,9})) // 6
    fmt.Println(gcdSum1([]int{9,8,7,6,5,4,3,2,1})) // 4
}