package main

// 2233. Maximum Product After K Increments
// You are given an array of non-negative integers nums and an integer k. 
// In one operation, you may choose any element from nums and increment it by 1.

// Return the maximum product of nums after at most k operations. 
// Since the answer may be very large, return it modulo 10^9 + 7. 
// Note that you should maximize the product before taking the modulo. 

// Example 1:
// Input: nums = [0,4], k = 5
// Output: 20
// Explanation: Increment the first number 5 times.
// Now nums = [5, 4], with a product of 5 * 4 = 20.
// It can be shown that 20 is maximum product possible, so we return 20.
// Note that there may be other ways to increment nums to have the maximum product.

// Example 2:
// Input: nums = [6,3,3,2], k = 2
// Output: 216
// Explanation: Increment the second number 1 time and increment the fourth number 1 time.
// Now nums = [6, 4, 3, 3], with a product of 6 * 4 * 3 * 3 = 216.
// It can be shown that 216 is maximum product possible, so we return 216.
// Note that there may be other ways to increment nums to have the maximum product.

// Constraints:
//     1 <= nums.length, k <= 10^5
//     0 <= nums[i] <= 10^6

import "fmt"
import "container/heap"
import "sort"

type PriorityQueue []int

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i] < pq[j] }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x any) { *pq = append(*pq, x.(int)) }
func (pq *PriorityQueue) Pop() any {
    n := len(*pq)
    v := (*pq)[n - 1]
    *pq = (*pq)[:n - 1]
    return v
}

func maximumProduct(nums []int, k int) int {
    pq := PriorityQueue(nums)
    heap.Init(&pq)
    for i := 0; i < k; i++ {
        heap.Push(&pq, heap.Pop(&pq).(int) + 1)
    }
    res := pq[0]
    for i := 1; i < pq.Len(); i++ {
        res = (res * pq[i]) % 1_000_000_007
    }
    return res
}

func maximumProduct1(nums []int, k int) int {
    mod := 1_000_000_007
    count := make(map[int]int) // 计数器，用于统计每个数字出现的次数
    for _, v := range nums {
        count[v]++
    }
    uniqueNums := []int{}
    for i := range count {
        uniqueNums = append(uniqueNums, i)
    }
    sort.Ints(uniqueNums) // 将所有数字排序
    for k > 0 && len(uniqueNums) > 0 {
        if len(uniqueNums) > 1 {
            a, b := uniqueNums[0], uniqueNums[1]
            ca := count[a]
            if ca * (b - a) <= k { // 可以将a全都加成b
                k -= ca * (b - a)
                delete(count, a)
                count[b] += ca
                uniqueNums = uniqueNums[1:]
            } else {
                solidAdd, scatterAdd := k / ca, k % ca
                if solidAdd > 0 {
                    oa := a
                    a += solidAdd
                    count[a] = ca
                    delete(count, oa)
                }
                if scatterAdd > 0 {
                    count[a] -= scatterAdd
                    count[a+1] += scatterAdd
                }
                k = 0
            }
        } else { // 如果只有一个元素，直接分散增加
            a := uniqueNums[0]
            ca := count[a]
            solidAdd, scatterAdd := k / ca, k % ca
            if solidAdd > 0 {
                oa := a
                a += solidAdd
                count[a] = ca
                delete(count, oa)
            }
            if scatterAdd > 0 {
                count[a] -= scatterAdd
                count[a+1] = scatterAdd
            }
            k = 0
        }
    }
    // 快速幂计算 (a^b) % mod
    qmi := func(a, b int) int {
        res := 1
        for b > 0 {
            if b&1 == 1 {
                res = res * a % mod
            }
            a = a * a % mod
            b >>= 1
        }
        return res
    }
    // 计算最终答案
    res := 1
    for k, v := range count {
        res *= qmi(k, v) % mod
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [0,4], k = 5
    // Output: 20
    // Explanation: Increment the first number 5 times.
    // Now nums = [5, 4], with a product of 5 * 4 = 20.
    // It can be shown that 20 is maximum product possible, so we return 20.
    // Note that there may be other ways to increment nums to have the maximum product.
    fmt.Println(maximumProduct([]int{0,4}, 5)) // 20
    // Example 2:
    // Input: nums = [6,3,3,2], k = 2
    // Output: 216
    // Explanation: Increment the second number 1 time and increment the fourth number 1 time.
    // Now nums = [6, 4, 3, 3], with a product of 6 * 4 * 3 * 3 = 216.
    // It can be shown that 216 is maximum product possible, so we return 216.
    // Note that there may be other ways to increment nums to have the maximum product.
    fmt.Println(maximumProduct([]int{6,3,3,2}, 2)) // 216
    fmt.Println(maximumProduct([]int{24,5,64,53,26,38}, 54)) // 180820950
   

    // fmt.Println(maximumProduct1([]int{0,4}, 5)) // 20
    // fmt.Println(maximumProduct1([]int{6,3,3,2}, 2)) // 216
}