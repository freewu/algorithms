package main

// 3629. Minimum Jumps to Reach End via Prime Teleportation
// You are given an integer array nums of length n.

// You start at index 0, and your goal is to reach index n - 1.

// From any index i, you may perform one of the following operations:
//     1. Adjacent Step: Jump to index i + 1 or i - 1, if the index is within bounds.
//     2. Prime Teleportation: If nums[i] is a prime number p, you may instantly jump to any index j != i such that nums[j] % p == 0.

// Return the minimum number of jumps required to reach index n - 1.

// A prime number is a natural number greater than 1 with only two factors, 1 and itself.

// Example 1:
// Input: nums = [1,2,4,6]
// Output: 2
// Explanation:
// One optimal sequence of jumps is:
// Start at index i = 0. Take an adjacent step to index 1.
// At index i = 1, nums[1] = 2 is a prime number. Therefore, we teleport to index i = 3 as nums[3] = 6 is divisible by 2.
// Thus, the answer is 2.

// Example 2:
// Input: nums = [2,3,4,7,9]
// Output: 2
// Explanation:
// One optimal sequence of jumps is:
// Start at index i = 0. Take an adjacent step to index i = 1.
// At index i = 1, nums[1] = 3 is a prime number. Therefore, we teleport to index i = 4 since nums[4] = 9 is divisible by 3.
// Thus, the answer is 2.

// Example 3:
// Input: nums = [4,6,5,8]
// Output: 3
// Explanation:
// Since no teleportation is possible, we move through 0 → 1 → 2 → 3. Thus, the answer is 3.
 
// Constraints:
//     1 <= n == nums.length <= 10^5
//     1 <= nums[i] <= 10^6

import "fmt"

// const MX = 1_000_001
// var factors [MX][]int

// func init() {
//     for i := 2; i < MX; i++ {
//         if len(factors[i]) == 0 {
//             for j := i; j < MX; j += i {
//                 factors[j] = append(factors[j], i)
//             }
//         }
//     }
// }

// func minJumps(nums []int) int {
//     res, n := 0, len(nums)
//     edges := make(map[int][]int)
//     for i, a := range nums {
//         if len(factors[a]) == 1 {
//             edges[a] = append(edges[a], i)
//         }
//     }
//     seen := make([]bool, n)
//     seen[n-1] = true
//     queue := []int{ n - 1 }
//     for {
//         newqueue := []int{}
//         for _, i := range queue {
//             if i == 0 {
//                 return res
//             }
//             if i > 0 && !seen[i-1] {
//                 seen[i-1] = true
//                 newqueue = append(newqueue, i-1)
//             }
//             if i < n-1 && !seen[i+1] {
//                 seen[i+1] = true
//                 newqueue = append(newqueue, i+1)
//             }
//             for _, p := range factors[nums[i]] {
//                 if list, ok := edges[p]; ok {
//                     for _, j := range list {
//                         if !seen[j] {
//                             seen[j] = true
//                             newqueue = append(newqueue, j)
//                         }
//                     }
//                     delete(edges, p)
//                 }
//             }
//         }
//         queue = newqueue
//         res++
//     }
// }

const MAXN = 100_000
const MAXV = 1_000_000
const SQRTMAXV = 1_000

type Bitset [(MAXV + 64) / 64]uint64

func (b *Bitset) Set(x int) {
    b[x/64] |= 1 << (x % 64)
}

func (b *Bitset) Get(x int) bool {
    return b[x/64]>>(x%64)&1 != 0
}

var sieve Bitset

func init() {
    sieve.Set(0)
    sieve.Set(1)
    for i := 4; i <= MAXV; i += 2 {
        sieve.Set(i)
    }
    for i := 3; i <= SQRTMAXV; i += 2 {
        if !sieve.Get(i) {
            for j := i * i; j <= MAXV; j += 2 * i {
                sieve.Set(j)
            }
        }
    }
}

func minJumps(nums []int) int {
    var adj [MAXV + 2]uint32
    var nxt, dist, queue [MAXN + 2]uint32
    var gen uint32
    gen += 1 << 20
    maxv, n := 0, uint32(len(nums))
    for i := n - 1; i != ^uint32(0); i-- {
        maxv = max(maxv, nums[i])
        nxt[i] = adj[nums[i]]
        adj[nums[i]] = i | gen
    }
    var visited Bitset
    head, tail := 0, 1
    queue[0] = 0
    dist[0] = gen
    for head < tail && queue[head] < n-1 {
        u := queue[head]
        head++
        d := dist[u]
        enqueue := func(v uint32) {
            if dist[v] < gen {
                dist[v] = d + 1
                queue[tail] = v
                tail++
            }
        }
        if u+1 < n {
            enqueue(u + 1)
        }
        if u > 0 {
            enqueue(u - 1)
        }
        p := nums[u]
        if sieve.Get(p) || visited.Get(p) {
            continue
        }
        visited.Set(p)
        for i := p; i <= maxv; i += p {
            cur := adj[i]
            for cur >= gen {
                v := cur &^ gen
                enqueue(v)
                cur = nxt[v]
            }
            adj[i] = 0
        }
    }
    return int(dist[queue[head]] - gen)
}

func minJumps1(nums []int) int {
    n := len(nums)
    if n <= 1 { return 0 }
    mx := nums[0]
    for _, v := range nums {
        if v > mx {
            mx = v
        }
    }
    spf := make([]int, mx + 1)
    for i := 0; i <= mx; i++ {
        spf[i] = i
    }
    for i := 2; i*i <= mx; i++ {
        if spf[i] == i {
            for j := i * i; j <= mx; j += i {
                if spf[j] == j {
                    spf[j] = i
                }
            }
        }
    }
    factorToIndices := make(map[int][]int, n * 2)
    for i := 0; i < n; i++ {
        prev, x := 0, nums[i]
        for x > 1 {
            p := spf[x]
            if p != prev {
                factorToIndices[p] = append(factorToIndices[p], i)
                prev = p
            }
            x /= p
        }
    }
    dist := make([]int, n)
    for i := range dist {
        dist[i] = 1 << 31
    }
    dist[0] = 0
    q := []int{0}
    for len(q) > 0 {
        u := q[0]
        q = q[1:]
        d := dist[u]
        if u == n-1 {
            return d
        }
        if u+1 < n && dist[u+1] == 1 << 31 {
            dist[u+1] = d + 1
            q = append(q, u+1)
        }
        if u-1 >= 0 && dist[u-1] == 1 << 31 {
            dist[u-1] = d + 1
            q = append(q, u-1)
        }
        val := nums[u]
        if val > 1 && spf[val] == val {
            if indices, ok := factorToIndices[val]; ok {
                for _, v := range indices {
                    if dist[v] == 1 << 31 {
                        dist[v] = d + 1
                        q = append(q, v)
                    }
                }
                delete(factorToIndices, val)
            }
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: nums = [1,2,4,6]
    // Output: 2
    // Explanation:
    // One optimal sequence of jumps is:
    // Start at index i = 0. Take an adjacent step to index 1.
    // At index i = 1, nums[1] = 2 is a prime number. Therefore, we teleport to index i = 3 as nums[3] = 6 is divisible by 2.
    // Thus, the answer is 2.
    fmt.Println(minJumps([]int{1,2,4,6})) // 2
    // Example 2:
    // Input: nums = [2,3,4,7,9]
    // Output: 2
    // Explanation:
    // One optimal sequence of jumps is:
    // Start at index i = 0. Take an adjacent step to index i = 1.
    // At index i = 1, nums[1] = 3 is a prime number. Therefore, we teleport to index i = 4 since nums[4] = 9 is divisible by 3.
    // Thus, the answer is 2.
    fmt.Println(minJumps([]int{2,3,4,7,9})) // 2
    // Example 3:
    // Input: nums = [4,6,5,8]
    // Output: 3
    // Explanation:
    // Since no teleportation is possible, we move through 0 → 1 → 2 → 3. Thus, the answer is 3.
    fmt.Println(minJumps([]int{4,6,5,8})) // 3

    fmt.Println(minJumps([]int{1,2,3,4,5,6,7,8,9})) // 3
    fmt.Println(minJumps([]int{9,8,7,6,5,4,3,2,1})) // 8

    fmt.Println(minJumps1([]int{1,2,4,6})) // 2
    fmt.Println(minJumps1([]int{2,3,4,7,9})) // 2
    fmt.Println(minJumps1([]int{4,6,5,8})) // 3
    fmt.Println(minJumps1([]int{1,2,3,4,5,6,7,8,9})) // 3
    fmt.Println(minJumps1([]int{9,8,7,6,5,4,3,2,1})) // 8
}