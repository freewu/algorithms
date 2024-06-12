package main

// 2813. Maximum Elegance of a K-Length Subsequence
// You are given a 0-indexed 2D integer array items of length n and an integer k.
// items[i] = [profiti, categoryi], where profiti and categoryi denote the profit and category of the ith item respectively.

// Let's define the elegance of a subsequence of items as total_profit + distinct_categories2, 
// where total_profit is the sum of all profits in the subsequence, 
// and distinct_categories is the number of distinct categories from all the categories in the selected subsequence.

// Your task is to find the maximum elegance from all subsequences of size k in items.
// Return an integer denoting the maximum elegance of a subsequence of items with size exactly k.
// Note: A subsequence of an array is a new array generated from the original array by deleting some elements (possibly none) without changing the remaining elements' relative order.

// Example 1:
// Input: items = [[3,2],[5,1],[10,1]], k = 2
// Output: 17
// Explanation: In this example, we have to select a subsequence of size 2.
// We can select items[0] = [3,2] and items[2] = [10,1].
// The total profit in this subsequence is 3 + 10 = 13, and the subsequence contains 2 distinct categories [2,1].
// Hence, the elegance is 13 + 22 = 17, and we can show that it is the maximum achievable elegance. 

// Example 2:
// Input: items = [[3,1],[3,1],[2,2],[5,3]], k = 3
// Output: 19
// Explanation: In this example, we have to select a subsequence of size 3. 
// We can select items[0] = [3,1], items[2] = [2,2], and items[3] = [5,3]. 
// The total profit in this subsequence is 3 + 2 + 5 = 10, and the subsequence contains 3 distinct categories [1,2,3]. 
// Hence, the elegance is 10 + 32 = 19, and we can show that it is the maximum achievable elegance.

// Example 3:
// Input: items = [[1,1],[2,1],[3,1]], k = 3
// Output: 7
// Explanation: In this example, we have to select a subsequence of size 3. 
// We should select all the items. 
// The total profit will be 1 + 2 + 3 = 6, and the subsequence contains 1 distinct category [1]. 
// Hence, the maximum elegance is 6 + 12 = 7.  

// Constraints:
//     1 <= items.length == n <= 10^5
//     items[i].length == 2
//     items[i][0] == profiti
//     items[i][1] == categoryi
//     1 <= profiti <= 10^9
//     1 <= categoryi <= n
//     1 <= k <= n

import "fmt"
import "sort"

func findMaximumElegance(items [][]int, k int) int64 {
    sort.Slice(items, func(i, j int) bool {
        return items[i][0] > items[j][0]
    })
    cat, sum := make(map[int]int), 0
    for i := 0; i < k; i++ {
        cat[items[i][1]]++
        sum += items[i][0]
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    cats, res := len(cat), sum + len(cat) * len(cat)  // 初始最大值, 找出profit最大的k個元素
    // 找出有沒有一種組合 可以損失 profit但換到更多的 category
    p := k - 1
    for i := k; i < len(items); i++ { // 從k個開始往後找 找到更小的profit
        if cat[items[i][1]] > 0 { // 如果這個category出現過 顯然不是我要的
            continue
        }
        for p >= 0 {
            // k - 1 是選中的組合裡面 profit最小的 從他開始看
            remove := items[p]
            p--
            rProfit, rCat := remove[0], remove[1]
            if cat[rCat] > 1 {
                sum -= rProfit
                sum += items[i][0]
                // remove from cat
                cat[rCat]--
                // add to cat
                cat[items[i][1]]++
                cats++
                res = max(res, sum + cats * cats)
                break
            }
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: items = [[3,2],[5,1],[10,1]], k = 2
    // Output: 17
    // Explanation: In this example, we have to select a subsequence of size 2.
    // We can select items[0] = [3,2] and items[2] = [10,1].
    // The total profit in this subsequence is 3 + 10 = 13, and the subsequence contains 2 distinct categories [2,1].
    // Hence, the elegance is 13 + 22 = 17, and we can show that it is the maximum achievable elegance. 
    fmt.Println(findMaximumElegance([][]int{{3,2},{5,1},{10,1}}, 2)) // 17
    // Example 2:
    // Input: items = [[3,1],[3,1],[2,2],[5,3]], k = 3
    // Output: 19
    // Explanation: In this example, we have to select a subsequence of size 3. 
    // We can select items[0] = [3,1], items[2] = [2,2], and items[3] = [5,3]. 
    // The total profit in this subsequence is 3 + 2 + 5 = 10, and the subsequence contains 3 distinct categories [1,2,3]. 
    // Hence, the elegance is 10 + 32 = 19, and we can show that it is the maximum achievable elegance.
    fmt.Println(findMaximumElegance([][]int{{3,1},{3,1},{2,2},{5,3}}, 3)) // 19
    // Example 3:
    // Input: items = [[1,1],[2,1],[3,1]], k = 3
    // Output: 7
    // Explanation: In this example, we have to select a subsequence of size 3. 
    // We should select all the items. 
    // The total profit will be 1 + 2 + 3 = 6, and the subsequence contains 1 distinct category [1]. 
    // Hence, the maximum elegance is 6 + 12 = 7.  
    fmt.Println(findMaximumElegance([][]int{{1,1},{2,1},{3,1}}, 3)) // 7
}