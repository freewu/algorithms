package main

// 2818. Apply Operations to Maximize Score
// You are given an array nums of n positive integers and an integer k.

// Initially, you start with a score of 1. 
// You have to maximize your score by applying the following operation at most k times:
//     1. Choose any non-empty subarray nums[l, ..., r] that you haven't chosen previously.
//     2. Choose an element x of nums[l, ..., r] with the highest prime score. 
//        If multiple such elements exist, choose the one with the smallest index.
//     3. Multiply your score by x.

// Here, nums[l, ..., r] denotes the subarray of nums starting at index l and ending at the index r, both ends being inclusive.

// The prime score of an integer x is equal to the number of distinct prime factors of x. 
// For example, the prime score of 300 is 3 since 300 = 2 * 2 * 3 * 5 * 5.

// Return the maximum possible score after applying at most k operations.

// Since the answer may be large, return it modulo 109 + 7.

// Example 1:
// Input: nums = [8,3,9,3,8], k = 2
// Output: 81
// Explanation: To get a score of 81, we can apply the following operations:
// - Choose subarray nums[2, ..., 2]. nums[2] is the only element in this subarray. Hence, we multiply the score by nums[2]. The score becomes 1 * 9 = 9.
// - Choose subarray nums[2, ..., 3]. Both nums[2] and nums[3] have a prime score of 1, but nums[2] has the smaller index. Hence, we multiply the score by nums[2]. The score becomes 9 * 9 = 81.
// It can be proven that 81 is the highest score one can obtain.

// Example 2:
// Input: nums = [19,12,14,6,10,18], k = 3
// Output: 4788
// Explanation: To get a score of 4788, we can apply the following operations: 
// - Choose subarray nums[0, ..., 0]. nums[0] is the only element in this subarray. Hence, we multiply the score by nums[0]. The score becomes 1 * 19 = 19.
// - Choose subarray nums[5, ..., 5]. nums[5] is the only element in this subarray. Hence, we multiply the score by nums[5]. The score becomes 19 * 18 = 342.
// - Choose subarray nums[2, ..., 3]. Both nums[2] and nums[3] have a prime score of 2, but nums[2] has the smaller index. Hence, we multipy the score by nums[2]. The score becomes 342 * 14 = 4788.
// It can be proven that 4788 is the highest score one can obtain.

// Constraints:
//     1 <= nums.length == n <= 10^5
//     1 <= nums[i] <= 10^5
//     1 <= k <= min(n * (n + 1) / 2, 10^9)

import "fmt"
import "sort"

func maximumScore(nums []int, k int) int {
    res, n, mod := 1, len(nums), 1_000_000_007
    pow := func(a, n int) int {
        res := 1
        for ; n > 0; n >>= 1 {
            if n&1 == 1 {
                res = res * a % mod
            }
            a = a * a % mod
        }
        return res
    }
    primeFactors := func(n int) int {
        i, mp := 2, make(map[int]bool)
        for i <= n / i {
            for n % i == 0 {
                mp[i] = true
                n /= i
            }
            i++
        }
        if n > 1 {
            mp[n] = true
        }
        return len(mp)
    }
    arr, left, right := make([][3]int, n), make([]int, n), make([]int, n)
    for i, v := range nums {
        left[i], right[i], arr[i] = -1, n, [3]int{i, primeFactors(v), v}
    }
    stack := []int{}
    for _, e := range arr {
        i, f := e[0], e[1]
        for len(stack) > 0 && arr[stack[len(stack) - 1]][1] < f {
            stack = stack[:len(stack) - 1]
        }
        if len(stack) > 0 {
            left[i] = stack[len(stack) - 1]
        }
        stack = append(stack, i)
    }
    stack = []int{}
    for i := n - 1; i >= 0; i-- {
        f := arr[i][1]
        for len(stack) > 0 && arr[stack[len(stack) - 1]][1] <= f {
            stack = stack[:len(stack) - 1]
        }
        if len(stack) > 0 {
            right[i] = stack[len(stack) - 1]
        }
        stack = append(stack, i)
    }
    sort.Slice(arr, func(i, j int) bool { 
        return arr[i][2] > arr[j][2] 
    })
    for _, e := range arr {
        i, x := e[0], e[2]
        l, r := left[i], right[i]
        count := (i - l) * (r - i)
        if count <= k {
            res = res * pow(x, count) % mod
            k -= count
        } else {
            res = res * pow(x, k) % mod
            break
        }
    }
    return res
}

const mx = 100_000
const mod = 1_000_000_007
const inf = int(1e18)

var primes [mx + 1]int

func init() {
    primes[0] = inf
    for i := 2; i <= mx; i++ {
        if primes[i] == 0 {
            for j := i; j <= mx; j += i {
                primes[j]++
            }
        }
    }
}

func maximumScore1(nums []int, k int) int {
    nums = append(nums, 0)
    stack, count := []int{ -1 }, make(map[int]int)
    for i, v := range nums {
        for len(stack) > 1 && primes[v] > primes[nums[stack[len(stack) - 1]]] {
            p := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            count[nums[p]] += (p - stack[len(stack) - 1]) * (i - p)
        }
        stack = append(stack, i)
    }
    keys := []int{}
    for v := range count {
        keys = append(keys, v)
    }
    sort.Slice(keys, func(i, j int) bool { 
        return keys[i] > keys[j] 
    })
    pow := func(x, n int) int {
        res := 1
        for ; n > 0; n >>= 1 {
            if n & 1 == 1 {
                res = res * x % mod
            }
            x = x * x % mod
        }
        return res
    }
    res := 1
    for i := 0; k > 0; i++ {
        t := min(count[keys[i]], k)
        k -= t
        res = res * pow(keys[i], t) % mod
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [8,3,9,3,8], k = 2
    // Output: 81
    // Explanation: To get a score of 81, we can apply the following operations:
    // - Choose subarray nums[2, ..., 2]. nums[2] is the only element in this subarray. Hence, we multiply the score by nums[2]. The score becomes 1 * 9 = 9.
    // - Choose subarray nums[2, ..., 3]. Both nums[2] and nums[3] have a prime score of 1, but nums[2] has the smaller index. Hence, we multiply the score by nums[2]. The score becomes 9 * 9 = 81.
    // It can be proven that 81 is the highest score one can obtain.
    fmt.Println(maximumScore([]int{8,3,9,3,8}, 2)) // 81
    // Example 2:
    // Input: nums = [19,12,14,6,10,18], k = 3
    // Output: 4788
    // Explanation: To get a score of 4788, we can apply the following operations: 
    // - Choose subarray nums[0, ..., 0]. nums[0] is the only element in this subarray. Hence, we multiply the score by nums[0]. The score becomes 1 * 19 = 19.
    // - Choose subarray nums[5, ..., 5]. nums[5] is the only element in this subarray. Hence, we multiply the score by nums[5]. The score becomes 19 * 18 = 342.
    // - Choose subarray nums[2, ..., 3]. Both nums[2] and nums[3] have a prime score of 2, but nums[2] has the smaller index. Hence, we multipy the score by nums[2]. The score becomes 342 * 14 = 4788.
    // It can be proven that 4788 is the highest score one can obtain.
    fmt.Println(maximumScore([]int{19,12,14,6,10,18}, 3)) // 4788

    fmt.Println(maximumScore([]int{1,2,3,4,5,6,7,8,9}, 2)) // 72
    fmt.Println(maximumScore([]int{9,8,7,6,5,4,3,2,1}, 2)) // 81

    fmt.Println(maximumScore1([]int{8,3,9,3,8}, 2)) // 81
    fmt.Println(maximumScore1([]int{19,12,14,6,10,18}, 3)) // 4788
    fmt.Println(maximumScore1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 72
    fmt.Println(maximumScore1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 81
}