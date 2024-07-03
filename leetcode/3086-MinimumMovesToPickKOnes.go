package main

// 3086. Minimum Moves to Pick K Ones
// You are given a binary array nums of length n, a positive integer k and a non-negative integer maxChanges.

// Alice plays a game, where the goal is for Alice to pick up k ones from nums using the minimum number of moves. 
// When the game starts, Alice picks up any index aliceIndex in the range [0, n - 1] and stands there. 
// If nums[aliceIndex] == 1 , Alice picks up the one and nums[aliceIndex] becomes 0(this does not count as a move). 
// After this, Alice can make any number of moves (including zero) where in each move Alice must perform exactly one of the following actions:

//     1. Select any index j != aliceIndex such that nums[j] == 0 and set nums[j] = 1. 
//        This action can be performed at most maxChanges times.
//     2. Select any two adjacent indices x and y (|x - y| == 1) such that nums[x] == 1, nums[y] == 0, then swap their values (set nums[y] = 1 and nums[x] = 0). 
//        If y == aliceIndex, Alice picks up the one after this move and nums[y] becomes 0.

// Return the minimum number of moves required by Alice to pick exactly k ones.

// Example 1:
// Input: nums = [1,1,0,0,0,1,1,0,0,1], k = 3, maxChanges = 1
// Output: 3
// Explanation: Alice can pick up 3 ones in 3 moves, if Alice performs the following actions in each move when standing at aliceIndex == 1:
// At the start of the game Alice picks up the one and nums[1] becomes 0. nums becomes [1,1,1,0,0,1,1,0,0,1].
// Select j == 2 and perform an action of the first type. nums becomes [1,0,1,0,0,1,1,0,0,1]
// Select x == 2 and y == 1, and perform an action of the second type. nums becomes [1,1,0,0,0,1,1,0,0,1]. As y == aliceIndex, Alice picks up the one and nums becomes [1,0,0,0,0,1,1,0,0,1].
// Select x == 0 and y == 1, and perform an action of the second type. nums becomes [0,1,0,0,0,1,1,0,0,1]. As y == aliceIndex, Alice picks up the one and nums becomes [0,0,0,0,0,1,1,0,0,1].
// Note that it may be possible for Alice to pick up 3 ones using some other sequence of 3 moves.

// Example 2:
// Input: nums = [0,0,0,0], k = 2, maxChanges = 3
// Output: 4
// Explanation: Alice can pick up 2 ones in 4 moves, if Alice performs the following actions in each move when standing at aliceIndex == 0:
// Select j == 1 and perform an action of the first type. nums becomes [0,1,0,0].
// Select x == 1 and y == 0, and perform an action of the second type. nums becomes [1,0,0,0]. As y == aliceIndex, Alice picks up the one and nums becomes [0,0,0,0].
// Select j == 1 again and perform an action of the first type. nums becomes [0,1,0,0].
// Select x == 1 and y == 0 again, and perform an action of the second type. nums becomes [1,0,0,0]. As y == aliceIndex, Alice picks up the one and nums becomes [0,0,0,0].

// Constraints:
//     2 <= n <= 10^5
//     0 <= nums[i] <= 1
//     1 <= k <= 10^5
//     0 <= maxChanges <= 10^5
//     maxChanges + sum(nums) >= k

import "fmt"

func minimumMoves(nums []int, k int, maxChanges int) int64 {
    pos := make([]int, 0)
    for i, num := range nums {
        if num != 0 {
            pos = append(pos, i)
        }
    }
    if len(pos) == 0 {
        return int64(2 * k)
    }
    if k == 1 {
        return int64(0)
    }
    m := len(pos)
    suffixSum, prefixSum := make([]int, m), make([]int, m)
    prefixSum[0] = pos[0]
    for i := 1; i < m; i++ {
        prefixSum[i] = prefixSum[i - 1] + pos[i]
    }
    for i := m - 2; i >= 0; i-- {
        suffixSum[i] = suffixSum[i + 1] + (pos[m - 1] - pos[i])
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res := 1 << 32 - 1
    maxChanges = min(maxChanges, k - 1)
    for change := max(0, maxChanges - 2); change <= maxChanges; change++ {
        required, changeCost := k - change, change * 2
        l, r := 0, required - 1
        mid := (l + r) / 2
        for r < m {
            tmp := changeCost
            tmp += suffixSum[l] - suffixSum[mid] - (mid - l) * (pos[m - 1] - pos[mid])
            tmp += prefixSum[r] - prefixSum[mid] - (r - mid) * (pos[mid])
            res = min(res, tmp)
            l++
            r++
            mid++
        }
    }
    return int64(res)
}

type prefixArray struct {
    rawArray  []int
    prefixSum []int
}

func (this *prefixArray) subSum(l, r int) int {
    sum := this.prefixSum[r]
    if l-1 >= 0 {
        sum -= this.prefixSum[l-1]
    }
    return sum
}

func newPrefixArray(array []int) prefixArray {
    prefixSum := make([]int, len(array))
    for i := range array {
        prefixSum[i] = array[i]
        if i-1 >= 0 {
            prefixSum[i] += prefixSum[i-1]
        }
    }
    return prefixArray{rawArray: array, prefixSum: prefixSum}
}

func minimumMoves1(nums []int, k int, maxChanges int) int64 {
    res := 1 << 32 - 1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    calcSumK := func(nums []int, k int) (int, bool) {
        if k == 0 {
            return 0, true
        }
        ones := []int{}
        for i := range nums {
            if nums[i] == 1 {
                ones = append(ones, i)
            }
        }
        if len(ones) < k {
            return 0, false
        }
        res, prefix := 1 << 32 - 1, newPrefixArray(ones)
        for r := k - 1; r < len(ones); r++ {
            mid := (k + 1) / 2
            //fmt.Println(ones, r, half, k)
            sum := prefix.subSum(r - mid + 1, r) - prefix.subSum(r - k + 1, r - k + mid)
            res = min(res, sum)
        }
        return res, true
    }
    for _, a := range []int{0, 1, 2, 3} {
        b := k - a
        if b < 0 || b > maxChanges {
            continue
        }
        t := b * 2
        v, ok := calcSumK(nums, a)
        if !ok {
            continue
        }
        res = min(res, t+v)
    }
    if k >= maxChanges {
        t := maxChanges * 2
        v, ok := calcSumK(nums, k-maxChanges)
        if ok {
            res = min(res, t+v)
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [1,1,0,0,0,1,1,0,0,1], k = 3, maxChanges = 1
    // Output: 3
    // Explanation: Alice can pick up 3 ones in 3 moves, if Alice performs the following actions in each move when standing at aliceIndex == 1:
    // At the start of the game Alice picks up the one and nums[1] becomes 0. nums becomes [1,1,1,0,0,1,1,0,0,1].
    // Select j == 2 and perform an action of the first type. nums becomes [1,0,1,0,0,1,1,0,0,1]
    // Select x == 2 and y == 1, and perform an action of the second type. nums becomes [1,1,0,0,0,1,1,0,0,1]. As y == aliceIndex, Alice picks up the one and nums becomes [1,0,0,0,0,1,1,0,0,1].
    // Select x == 0 and y == 1, and perform an action of the second type. nums becomes [0,1,0,0,0,1,1,0,0,1]. As y == aliceIndex, Alice picks up the one and nums becomes [0,0,0,0,0,1,1,0,0,1].
    // Note that it may be possible for Alice to pick up 3 ones using some other sequence of 3 moves.
    fmt.Println(minimumMoves([]int{1,1,0,0,0,1,1,0,0,1}, 3, 1)) // 3
    // Example 2:
    // Input: nums = [0,0,0,0], k = 2, maxChanges = 3
    // Output: 4
    // Explanation: Alice can pick up 2 ones in 4 moves, if Alice performs the following actions in each move when standing at aliceIndex == 0:
    // Select j == 1 and perform an action of the first type. nums becomes [0,1,0,0].
    // Select x == 1 and y == 0, and perform an action of the second type. nums becomes [1,0,0,0]. As y == aliceIndex, Alice picks up the one and nums becomes [0,0,0,0].
    // Select j == 1 again and perform an action of the first type. nums becomes [0,1,0,0].
    // Select x == 1 and y == 0 again, and perform an action of the second type. nums becomes [1,0,0,0]. As y == aliceIndex, Alice picks up the one and nums becomes [0,0,0,0].
    fmt.Println(minimumMoves([]int{0,0,0,0}, 2, 3)) // 4

    fmt.Println(minimumMoves1([]int{1,1,0,0,0,1,1,0,0,1}, 3, 1)) // 3
    fmt.Println(minimumMoves1([]int{0,0,0,0}, 2, 3)) // 4
}