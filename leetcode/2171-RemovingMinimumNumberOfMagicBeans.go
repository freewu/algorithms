package main

// 2171. Removing Minimum Number of Magic Beans
// You are given an array of positive integers beans, 
// where each integer represents the number of magic beans found in a particular magic bag.
// Remove any number of beans (possibly none) from each bag such that the number of beans in each remaining non-empty bag (still containing at least one bean) is equal. 
// Once a bean has been removed from a bag, you are not allowed to return it to any of the bags.

// Return the minimum number of magic beans that you have to remove.

// Example 1:
// Input: beans = [4,1,6,5]
// Output: 4
// Explanation: 
// - We remove 1 bean from the bag with only 1 bean.
//   This results in the remaining bags: [4,0,6,5]
// - Then we remove 2 beans from the bag with 6 beans.
//   This results in the remaining bags: [4,0,4,5]
// - Then we remove 1 bean from the bag with 5 beans.
//   This results in the remaining bags: [4,0,4,4]
// We removed a total of 1 + 2 + 1 = 4 beans to make the remaining non-empty bags have an equal number of beans.
// There are no other solutions that remove 4 beans or fewer.

// Example 2:
// Input: beans = [2,10,3,2]
// Output: 7
// Explanation:
// - We remove 2 beans from one of the bags with 2 beans.
//   This results in the remaining bags: [0,10,3,2]
// - Then we remove 2 beans from the other bag with 2 beans.
//   This results in the remaining bags: [0,10,3,0]
// - Then we remove 3 beans from the bag with 3 beans. 
//   This results in the remaining bags: [0,10,0,0]
// We removed a total of 2 + 2 + 3 = 7 beans to make the remaining non-empty bags have an equal number of beans.
// There are no other solutions that removes 7 beans or fewer.
 
// Constraints:
//     1 <= beans.length <= 10^5
//     1 <= beans[i] <= 10^5

import "fmt"
import "sort"

func minimumRemoval(beans []int) int64 {
    sort.Ints(beans)
    res, total, n := 0, 0, len(beans)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, bean := range beans {
        res = max(res, bean * n)
        n--
        total += bean
    }
    return int64(total - res)
}

func minimumRemoval1(beans []int) int64 {
    sort.Ints(beans)
    n, total := len(beans), 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for _, bean := range beans {
        total += bean
    }
    res := total
    for i := 0; i < n; i++ {
        res = min(res, total - beans[i] * (n - i))
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: beans = [4,1,6,5]
    // Output: 4
    // Explanation: 
    // - We remove 1 bean from the bag with only 1 bean.
    //   This results in the remaining bags: [4,0,6,5]
    // - Then we remove 2 beans from the bag with 6 beans.
    //   This results in the remaining bags: [4,0,4,5]
    // - Then we remove 1 bean from the bag with 5 beans.
    //   This results in the remaining bags: [4,0,4,4]
    // We removed a total of 1 + 2 + 1 = 4 beans to make the remaining non-empty bags have an equal number of beans.
    // There are no other solutions that remove 4 beans or fewer.
    fmt.Println(minimumRemoval([]int{4,1,6,5})) // 4
    // Example 2:
    // Input: beans = [2,10,3,2]
    // Output: 7
    // Explanation:
    // - We remove 2 beans from one of the bags with 2 beans.
    //   This results in the remaining bags: [0,10,3,2]
    // - Then we remove 2 beans from the other bag with 2 beans.
    //   This results in the remaining bags: [0,10,3,0]
    // - Then we remove 3 beans from the bag with 3 beans. 
    //   This results in the remaining bags: [0,10,0,0]
    // We removed a total of 2 + 2 + 3 = 7 beans to make the remaining non-empty bags have an equal number of beans.
    // There are no other solutions that removes 7 beans or fewer.
    fmt.Println(minimumRemoval([]int{2,10,3,2})) // 7

    fmt.Println(minimumRemoval1([]int{4,1,6,5})) // 4
    fmt.Println(minimumRemoval1([]int{2,10,3,2})) // 7
}