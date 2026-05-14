package main

// 3930. Power Update After K-th Largest Insertion II
// You are given an integer array nums and an integer p.

// You are also given a 2D integer array queries, where each queries[i] = [vali, ki].

// For each query:
//     1. Insert vali into nums.
//     2. Let x be the kith largest element in the current nums.
//     3. Update p to px % (10^9 + 7).

// Return an array ans where the ans[i] represents the value of p after processing the ith query.

// Example 1:
// Input: nums = [2], p = 4, queries = [[3,1],[1,2]]
// Output: [64,4096]
// Explanation:
// i | vali | Current nums | ki | kith largest | p  | New p = pk % (10^9 + 7)
// 0 | 3    | [2, 3]       | 1	 | 3            | 4	 | 4^3 % (10^9 + 7) = 64
// 1 | 1    | [2, 3, 1]    | 2	 | 2            | 64 | 64^2 % (10^9 + 7) = 4096
// Thus, ans = [64, 4096].

// Example 2:
// Input: nums = [7,5], p = 6, queries = [[4,3],[7,2]]
// Output: [1296,220296870]
// Explanation:
// i | vali | Current nums | ki | kith largest | p    | New p = pk % (10^9 + 7)
// 0 | 4    | [7, 5, 4]    | 1	 | 4            | 6    | 6^4 % (10^9 + 7) = 1296
// 1 | 7    | [7, 5, 4, 7]	| 2	 | 7            | 1296 | 1296^7 % (10^9 + 7) = 220296870
// Thus, ans = [1296, 220296870]

// Constraints:
//     1 <= nums.length <= 2 * 10^4
//     1 <= nums[i] <= 10^9
//     ​​​​​​​1 <= p <= 10^9
//     1 <= queries.length <= 2 * 10^4
//     ​​​​​​​1 <= vali <= 10^9
//     1 <= ki <= n + i + 1​​​​​​​

import "fmt"

func powerUpdate(nums []int, p int, queries [][]int) []int {
    const MOD = 1_000_000_007
    // Pair { 数值, 唯一索引 }
    type Pair struct {
        val int
        index int
    }
    pow := func(a, b int) int { // 快速幂取模
        res := 1
        a %= MOD
        for b > 0 {
            if b&1 == 1 {
                res = (res * a) % MOD
            }
            a = (a * a) % MOD
            b >>= 1
        }
        return res
    }
    findInsertPos := func(arr []Pair, p Pair) int { // 二分查找：找到降序排列中，新元素应该插入的位置
        l, r := 0, len(arr)
        for l < r {
            mid := (l + r) / 2
            // 降序规则：先比数值，数值相同比索引（保证唯一）
            if arr[mid].val < p.val || (arr[mid].val == p.val && arr[mid].index < p.index) {
                r = mid
            } else {
                l = mid + 1
            }
        }
        return l
    }
    res := make([]int, 0, len(queries))
    orderedList := []Pair{} // 初始化降序有序切片
    curr, n := p, len(nums)
    // 插入初始nums元素（和C++逻辑一致：数值+原始索引）
    for i := 0; i < n; i++ {
        pos := findInsertPos(orderedList, Pair{val: nums[i], index: i})
        // 切片插入：在pos位置插入元素
        orderedList = append(orderedList[:pos], append([]Pair{{val: nums[i], index: i}}, orderedList[pos:]...)...)
    }
    // 处理每个查询
    for i, q := range queries {
        val, k := q[0], q[1]
        // 插入新元素：索引 = n+i（和C++完全一致，保证唯一）
        newPair := Pair{val: val, index: n + i}
        pos := findInsertPos(orderedList, newPair)
        orderedList = append(orderedList[:pos], append([]Pair{newPair}, orderedList[pos:]...)...)
        // 取第k-1个元素（find_by_order(k-1)）
        x := orderedList[k-1].val
        // 更新幂次
        curr = pow(curr, x)
        res = append(res, curr)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2], p = 4, queries = [[3,1],[1,2]]
    // Output: [64,4096]
    // Explanation:
    // i | vali | Current nums | ki | kith largest | p  | New p = pk % (10^9 + 7)
    // 0 | 3    | [2, 3]       | 1	 | 3            | 4	 | 4^3 % (10^9 + 7) = 64
    // 1 | 1    | [2, 3, 1]    | 2	 | 2            | 64 | 64^2 % (10^9 + 7) = 4096
    // Thus, ans = [64, 4096].
    fmt.Println(powerUpdate([]int{2}, 4, [][]int{{3,1},{1,2}})) // [64,4096]
    // Example 2:
    // Input: nums = [7,5], p = 6, queries = [[4,3],[7,2]]
    // Output: [1296,220296870]
    // Explanation:
    // i | vali | Current nums | ki | kith largest | p    | New p = pk % (10^9 + 7)
    // 0 | 4    | [7, 5, 4]    | 1	 | 4            | 6    | 6^4 % (10^9 + 7) = 1296
    // 1 | 7    | [7, 5, 4, 7]	| 2	 | 7            | 1296 | 1296^7 % (10^9 + 7) = 220296870
    // Thus, ans = [1296, 220296870]
    fmt.Println(powerUpdate([]int{7,5}, 6, [][]int{{4,3},{7,2}})) // [1296,220296870]

    fmt.Println(powerUpdate([]int{1,2,3,4,5,6,7,8,9}, 6, [][]int{{4,3},{7,2}})) // [279936 592081930]
    fmt.Println(powerUpdate([]int{9,8,7,6,5,4,3,2,1}, 6, [][]int{{4,3},{7,2}})) // [279936 592081930]
}