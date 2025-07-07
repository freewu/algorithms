package main

// 3605. Minimum Stability Factor of Array
// You are given an integer array nums and an integer maxC.

// A subarray is called stable if the highest common factor (HCF) of all its elements is greater than or equal to 2.

// The stability factor of an array is defined as the length of its longest stable subarray.

// You may modify at most maxC elements of the array to any integer.

// Return the minimum possible stability factor of the array after at most maxC modifications. 
// If no stable subarray remains, return 0.

// Note:
//     1. The highest common factor (HCF) of an array is the largest integer that evenly divides all the array elements.
//     2. A subarray of length 1 is stable if its only element is greater than or equal to 2, since HCF([x]) = x.

// Example 1:
// Input: nums = [3,5,10], maxC = 1
// Output: 1
// Explanation:
// The stable subarray [5, 10] has HCF = 5, which has a stability factor of 2.
// Since maxC = 1, one optimal strategy is to change nums[1] to 7, resulting in nums = [3, 7, 10].
// Now, no subarray of length greater than 1 has HCF >= 2. Thus, the minimum possible stability factor is 1.

// Example 2:
// Input: nums = [2,6,8], maxC = 2
// Output: 1
// Explanation:
// The subarray [2, 6, 8] has HCF = 2, which has a stability factor of 3.
// Since maxC = 2, one optimal strategy is to change nums[1] to 3 and nums[2] to 5, resulting in nums = [2, 3, 5].
// Now, no subarray of length greater than 1 has HCF >= 2. Thus, the minimum possible stability factor is 1.

// Example 3:
// Input: nums = [2,4,9,6], maxC = 1
// Output: 2
// Explanation:
// The stable subarrays are:
// [2, 4] with HCF = 2 and stability factor of 2.
// [9, 6] with HCF = 3 and stability factor of 2.
// Since maxC = 1, the stability factor of 2 cannot be reduced due to two separate stable subarrays. Thus, the minimum possible stability factor is 2.

// Constraints:
//     1 <= n == nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     0 <= maxC <= n

import "fmt"

func minStable(nums []int, maxC int) int {
    n := len(nums)
    l, r := 0, n + 1
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    countBroken := func(mid int) int {
        res := 0
        for i := 0; i + mid - 1 < n; {
            g := nums[i]
            j := i + 1
            for j < i + mid && g > 1 {
                g = gcd(g, nums[j])
                j++
            }
            if g > 1 {
                res++
                i += mid
            } else {
                i++
            }
        }
        return res
    }
    for l + 1 < r {
        mid := l + (r - l) / 2
        if countBroken(mid) <= maxC {
            r = mid
        } else {
            l = mid
        }
    }
    return l
}

func minStable1(nums []int, maxC int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    countNon1 := 0
    for _, v := range nums {
        if v != 1 {
            countNon1++
        }
    }
    log2Table := make([]int, n+1)
    if n >= 1 {
        log2Table[1] = 0
    }
    for i := 2; i <= n; i++ {
        log2Table[i] = log2Table[i>>1] + 1
    }
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    maxLog := log2Table[n] + 1
    st := make([][]int, maxLog)
    st[0] = make([]int, n)
    copy(st[0], nums)
    for j := 1; j < maxLog; j++ {
        step := 1 << (j-1)
        totalLen := 1 << j
        if totalLen > n { break }
        st[j] = make([]int, n - totalLen + 1)
        for i := 0; i <= n-totalLen; i++ {
            st[j][i] = gcd(st[j-1][i], st[j-1][i+step])
        }
    }
    queryFunc := func(l, r int) int {
        if l > r { return 0 }
        n := r - l + 1
        j := log2Table[n]
        rightStart := r - (1 << j) + 1
        return gcd(st[j][l], st[j][rightStart])
    }
    check := func(x int) bool {
        if x == 0 { return countNon1 <= maxC }
        L := x + 1
        if L > n { return true }
        breaks, lastBreak := 0, -1
        for i := 0; i <= n-L; i++ {
            if lastBreak != -1 && lastBreak >= i && lastBreak <= i+L-1 { continue }
            g := queryFunc(i, i+L-1)
            if g >= 2 {
                if breaks >= maxC { return false }
                breaks++
                lastBreak = i + L - 1
            }
        }
        return true
    }
    res, low, high := n, 0, n
    for low <= high {
        mid := (low + high) >> 1
        if check(mid) {
            res = mid
            high = mid - 1
        } else {
            low = mid + 1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,5,10], maxC = 1
    // Output: 1
    // Explanation:
    // The stable subarray [5, 10] has HCF = 5, which has a stability factor of 2.
    // Since maxC = 1, one optimal strategy is to change nums[1] to 7, resulting in nums = [3, 7, 10].
    // Now, no subarray of length greater than 1 has HCF >= 2. Thus, the minimum possible stability factor is 1.
    fmt.Println(minStable([]int{3,5,10}, 1)) // 1
    // Example 2:
    // Input: nums = [2,6,8], maxC = 2
    // Output: 1
    // Explanation:
    // The subarray [2, 6, 8] has HCF = 2, which has a stability factor of 3.
    // Since maxC = 2, one optimal strategy is to change nums[1] to 3 and nums[2] to 5, resulting in nums = [2, 3, 5].
    // Now, no subarray of length greater than 1 has HCF >= 2. Thus, the minimum possible stability factor is 1.
    fmt.Println(minStable([]int{2,6,8}, 2)) // 1
    // Example 3:
    // Input: nums = [2,4,9,6], maxC = 1
    // Output: 2
    // Explanation:
    // The stable subarrays are:
    // [2, 4] with HCF = 2 and stability factor of 2.
    // [9, 6] with HCF = 3 and stability factor of 2.
    // Since maxC = 1, the stability factor of 2 cannot be reduced due to two separate stable subarrays. Thus, the minimum possible stability factor is 2.
    fmt.Println(minStable([]int{2,4,9,6}, 1)) // 2

    fmt.Println(minStable([]int{1,2,3,4,5,6,7,8,9}, 1)) // 1
    fmt.Println(minStable([]int{9,8,7,6,5,4,3,2,1}, 1)) // 1

    fmt.Println(minStable1([]int{3,5,10}, 1)) // 1
    fmt.Println(minStable1([]int{2,6,8}, 2)) // 1
    fmt.Println(minStable1([]int{2,4,9,6}, 1)) // 2
    fmt.Println(minStable1([]int{1,2,3,4,5,6,7,8,9}, 1)) // 1
    fmt.Println(minStable1([]int{9,8,7,6,5,4,3,2,1}, 1)) // 1
}