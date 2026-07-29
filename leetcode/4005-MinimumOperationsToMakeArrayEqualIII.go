package main

// 4005. Minimum Operations to Make Array Equal III
// You are given an integer array nums.

// In one operation, you may choose any element nums[i] and perform one of the following:
//     1. Multiply nums[i] by an integer k, where k >= 2.
//     2. Divide nums[i] by an integer k, where 2 <= k < nums[i], provided that nums[i] is divisible by k.

// Return the minimum number of operations required to make all elements of nums equal.

// Example 1:
// Input: nums = [6,12,8]
// Output: 3
// Explanation:
// We can perform following operates to make all numbers to 6:
// Divide nums[1] = 12 by 2 to get 6.
// Divide nums[2] = 8 by 4 to get 2.
// Multiply nums[2] = 2 by 3 to get 6.

// Example 2:
// Input: nums = [5,15,20]
// Output: 2
// Explanation:
// We can perform following operates to make all numbers to 5:
// Divide nums[1] = 15 by 3 to get 5.
// Divide nums[2] = 20 by 4 to get 5.

// Example 3:
// Input: nums = [7,7,7]
// Output: 0
// Explanation:
// All elements are already equal, so no operations are needed.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10​​​​​​^​9

// Tips:
//     1. For any target value greater than 1, an element can be changed to the target in at most two operations.
//     2. An element requires one operation if it divides the target or the target divides it. 
//     Otherwise, it requires two operations.
//     3. Choosing a common multiple of all elements gives an answer of at most nums.length. 
//     Therefore, it is enough to compare this answer with targets that already appear in nums.
//     4. For each distinct target x, count how many array elements divide x or are divisible by x. 
//     Use frequencies and divisor enumeration to compute these counts efficiently.
//     5. Elements equal to x require zero operations, other comparable elements require one operation, and all remaining elements require two operations.

import "fmt"

// 生成 [2, limit] 内的所有素数
func getPrimes(limit int) []int {
    sieve := make([]bool, limit+1)
    for i := 2; i*i <= limit; i++ {
        if !sieve[i] {
            for j := i * i; j <= limit; j += i {
                sieve[j] = true
            }
        }
    }
    primes := []int{}
    for i := 2; i <= limit; i++ {
        if !sieve[i] {
            primes = append(primes, i)
        }
    }
    return primes
}

// 快速幂取模
func powMod(a, exp, mod int) int {
    res := 1
    a %= mod
    for exp > 0 {
        if exp&1 == 1 {
            res = (res * a) % mod
        }
        a = (a * a) % mod
        exp >>= 1
    }
    return res
}

// Miller-Rabin 素性测试，适用于 32 位整数
func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    if n == 2 || n == 3 || n == 5 || n == 7 {
        return true
    }
    if n % 2 == 0 {
        return false
    }
    d := n - 1
    s := 0
    for d%2 == 0 {
        d /= 2
        s++
    }
    for _, a := range []int{2, 7, 61} {
        if a >= n {
            continue
        }
        x := powMod(a, d, n)
        if x == 1 || x == n-1 {
            continue
        }
        composite := true
        for r := 0; r < s-1; r++ {
            x = (x * x) % n
            if x == n-1 {
                composite = false
                break
            }
        }
        if composite {
            return false
        }
    }
    return true
}

// 用试除法分解合数，返回质因数和指数
func factorize(n int, primes []int) map[int]int {
    factors := make(map[int]int)
    temp := n
    for _, p := range primes {
        if p*p > temp {
            break
        }
        if temp%p == 0 {
            cnt := 0
            for temp%p == 0 {
                temp /= p
                cnt++
            }
            factors[p] = cnt
        }
    }
    if temp > 1 {
        factors[temp]++
    }
    return factors
}

// 根据质因数分解生成所有因子（包含 1 和自身）
func getDivisors(factors map[int]int) []int {
    divs := []int{1}
    for p, exp := range factors {
        curLen := len(divs)
        pow := 1
        for e := 1; e <= exp; e++ {
            pow *= p
            for i := 0; i < curLen; i++ {
                divs = append(divs, divs[i]*pow)
            }
        }
    }
    return divs
}

func minOperations(nums []int) int64 {
    freq := make(map[int]int)
    ones := 0
    for _, v := range nums {
        if v == 1 {
            ones++
        } else {
            freq[v]++
        }
    }
    count := len(nums) - ones
    if count == 0 {
        return 0
    }
    // S[x] = 2*cnt[x] + covered[x] ，初始化时 covered 为 0
    S := make(map[int]int, len(freq))
    for v, c := range freq {
        S[v] = 2 * c
    }
    primes := getPrimes(31623) // sqrt(10^9) ≈ 31623
    // 枚举所有整除对
    for v, c := range freq {
        if isPrime(v) { // 质数没有真因子
            continue
        }
        factors := factorize(v, primes)
        divisors := getDivisors(factors)
        for _, d := range divisors {
            if d <= 1 || d >= v {
                continue
            }
            if cntD, ok := freq[d]; ok {
                // d 是 v 的因子 => d 与 v 有整除关系
                S[d] += c
                S[v] += cntD
            }
        }
    }
    // 基线：选取一个非常大的公倍数，此时 S = count
    mx := count
    for _, s := range S {
        if s > mx {
            mx = s
        }
    }
    return int64(ones + 2 * count - mx)
}

func main() {
    // Example 1:
    // Input: nums = [6,12,8]
    // Output: 3
    // Explanation:
    // We can perform following operates to make all numbers to 6:
    // Divide nums[1] = 12 by 2 to get 6.
    // Divide nums[2] = 8 by 4 to get 2.
    // Multiply nums[2] = 2 by 3 to get 6.
    fmt.Println(minOperations([]int{6,12,8})) // 3
    // Example 2:
    // Input: nums = [5,15,20]
    // Output: 2
    // Explanation:
    // We can perform following operates to make all numbers to 5:
    // Divide nums[1] = 15 by 3 to get 5.
    // Divide nums[2] = 20 by 4 to get 5.
    fmt.Println(minOperations([]int{5,15,20})) // 2
    // Example 3:
    // Input: nums = [7,7,7]
    // Output: 0
    // Explanation:
    // All elements are already equal, so no operations are needed.
    fmt.Println(minOperations([]int{7,7,7})) // 0
    // Example 4:
    // Input: nums = [1,1,1,3,2]
    // Output: 5
    fmt.Println(minOperations([]int{1,1,1,3,2})) // 5
    // Example 5:
    // Input: nums = [1,2,5,3,1]
    // Output: 5
    fmt.Println(minOperations([]int{1,2,5,3,1})) // 5

    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1})) // 9
}