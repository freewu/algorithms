package main

// 846. Hand of Straights
// Alice has some number of cards and she wants to rearrange the cards into groups so that each group is of size groupSize, 
// and consists of groupSize consecutive cards.

// Given an integer array hand where hand[i] is the value written on the ith card and an integer groupSize, 
// return true if she can rearrange the cards, or false otherwise.

// Example 1:
// Input: hand = [1,2,3,6,2,3,4,7,8], groupSize = 3
// Output: true
// Explanation: Alice's hand can be rearranged as [1,2,3],[2,3,4],[6,7,8]

// Example 2:
// Input: hand = [1,2,3,4,5], groupSize = 4
// Output: false
// Explanation: Alice's hand can not be rearranged into groups of 4.

// Constraints:
//     1 <= hand.length <= 10^4
//     0 <= hand[i] <= 10^9
//     1 <= groupSize <= hand.length

import "fmt"
import "sort"
import "slices"

func isNStraightHand(hand []int, groupSize int) bool {
    if len(hand) % groupSize != 0 { // 不够分组
        return false
    }
    sort.Ints(hand)
    totalSplits, mapped := 0, make(map[int]int)
    for _, digit := range hand { // 统计牌数量
        mapped[digit]++
    }
    for _, digit := range hand {
        if mapped[digit] > 0 {
            totalSplits++
            for index := 0; index < groupSize; index++ { // 分组
                if mapped[digit] <= 0 { // 不够分组了
                    return false
                } else {
                    mapped[digit]--
                }
                digit++ // 1,2,3 | 2,3,4
            }
        }
    }
    return totalSplits == len(hand) / groupSize
}

func isNStraightHand1(hand []int, groupSize int) bool {
    if groupSize == 1 {
        return true
    }
    if len(hand) % groupSize != 0 { // 不够分组
        return false
    }
    slices.Sort(hand)
    list := make([][]int, 0)
    for _, num := range hand {
        flag := true
        for i := 0; i < len(list); i++ {
            v := list[i][len(list[i])-1]
            if v + 1 == num {
                list[i] = append(list[i], num)
                if len(list[i]) == groupSize {
                    list = slices.Delete(list, i, i+1)
                }
                flag = false
                break
            } else if v + 1 < num {
                return false
            }
        }
        if flag {
            list = append(list, []int{num})
        }
    }
    return len(list) == 0
}

func main() {
    // Example 1:
    // Input: hand = [1,2,3,6,2,3,4,7,8], groupSize = 3
    // Output: true
    // Explanation: Alice's hand can be rearranged as [1,2,3],[2,3,4],[6,7,8]
    fmt.Println(isNStraightHand([]int{1,2,3,6,2,3,4,7,8}, 3)) // true
    // Example 2:
    // Input: hand = [1,2,3,4,5], groupSize = 4
    // Output: false
    // Explanation: Alice's hand can not be rearranged into groups of 4.
    fmt.Println(isNStraightHand([]int{1,2,3,4,5}, 4)) // false

    fmt.Println(isNStraightHand1([]int{1,2,3,6,2,3,4,7,8}, 3)) // true
    fmt.Println(isNStraightHand1([]int{1,2,3,4,5}, 4)) // false
}