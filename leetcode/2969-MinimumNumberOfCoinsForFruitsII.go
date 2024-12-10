package main

// 2969. Minimum Number of Coins for Fruits II
// You are at a fruit market with different types of exotic fruits on display.

// You are given a 1-indexed array prices, where prices[i] denotes the number of coins needed to purchase the ith fruit.

// The fruit market has the following offer:
//     If you purchase the ith fruit at prices[i] coins, you can get the next i fruits for free.

// Note that even if you can take fruit j for free, you can still purchase it for prices[j] coins to receive a new offer.

// Return the minimum number of coins needed to acquire all the fruits.

// Example 1:
// Input: prices = [3,1,2]
// Output: 4
// Explanation: You can acquire the fruits as follows:
// - Purchase the 1st fruit with 3 coins, and you are allowed to take the 2nd fruit for free.
// - Purchase the 2nd fruit with 1 coin, and you are allowed to take the 3rd fruit for free.
// - Take the 3rd fruit for free.
// Note that even though you were allowed to take the 2nd fruit for free, you purchased it because it is more optimal.
// It can be proven that 4 is the minimum number of coins needed to acquire all the fruits.

// Example 2:
// Input: prices = [1,10,1,1]
// Output: 2
// Explanation: You can acquire the fruits as follows:
// - Purchase the 1st fruit with 1 coin, and you are allowed to take the 2nd fruit for free.
// - Take the 2nd fruit for free.
// - Purchase the 3rd fruit for 1 coin, and you are allowed to take the 4th fruit for free.
// - Take the 4th fruit for free.
// It can be proven that 2 is the minimum number of coins needed to acquire all the fruits.

// Constraints:
//     1 <= prices.length <= 10^5
//     1 <= prices[i] <= 10^5

import "fmt"

// template
type Deque struct{ l, r []int }

func (q Deque) Empty() bool {
    return len(q.l) == 0 && len(q.r) == 0
}

func (q Deque) Size() int {
    return len(q.l) + len(q.r)
}

func (q *Deque) PushFront(v int) {
    q.l = append(q.l, v)
}

func (q *Deque) PushBack(v int) {
    q.r = append(q.r, v)
}

func (q *Deque) PopFront() (v int) {
    if len(q.l) > 0 {
        q.l, v = q.l[:len(q.l)-1], q.l[len(q.l)-1]
    } else {
        v, q.r = q.r[0], q.r[1:]
    }
    return
}

func (q *Deque) PopBack() (v int) {
    if len(q.r) > 0 {
        q.r, v = q.r[:len(q.r)-1], q.r[len(q.r)-1]
    } else {
        v, q.l = q.l[0], q.l[1:]
    }
    return
}

func (q Deque) Front() int {
    if len(q.l) > 0 {
        return q.l[len(q.l)-1]
    }
    return q.r[0]
}

func (q Deque) Back() int {
    if len(q.r) > 0 {
        return q.r[len(q.r)-1]
    }
    return q.l[0]
}

func (q Deque) Get(i int) int {
    if i < len(q.l) {
        return q.l[len(q.l)-1-i]
    }
    return q.r[i-len(q.l)]
}

func minimumCoins(prices []int) int {
    n := len(prices)
    q := Deque{}
    for i := n; i > 0; i-- {
        for q.Size() > 0 && q.Front() > i*2+1 {
            q.PopFront()
        }
        if i <= (n-1)/2 {
            prices[i-1] += prices[q.Front()-1]
        }
        for q.Size() > 0 && prices[q.Back()-1] >= prices[i-1] {
            q.PopBack()
        }
        q.PushBack(i)
    }
    return prices[0]
}

func main() {
    // Example 1:
    // Input: prices = [3,1,2]
    // Output: 4
    // Explanation: You can acquire the fruits as follows:
    // - Purchase the 1st fruit with 3 coins, and you are allowed to take the 2nd fruit for free.
    // - Purchase the 2nd fruit with 1 coin, and you are allowed to take the 3rd fruit for free.
    // - Take the 3rd fruit for free.
    // Note that even though you were allowed to take the 2nd fruit for free, you purchased it because it is more optimal.
    // It can be proven that 4 is the minimum number of coins needed to acquire all the fruits.
    fmt.Println(minimumCoins([]int{3,1,2})) // 4
    // Example 2:
    // Input: prices = [1,10,1,1]
    // Output: 2
    // Explanation: You can acquire the fruits as follows:
    // - Purchase the 1st fruit with 1 coin, and you are allowed to take the 2nd fruit for free.
    // - Take the 2nd fruit for free.
    // - Purchase the 3rd fruit for 1 coin, and you are allowed to take the 4th fruit for free.
    // - Take the 4th fruit for free.
    // It can be proven that 2 is the minimum number of coins needed to acquire all the fruits.
    fmt.Println(minimumCoins([]int{1,10,1,1})) // 2
}