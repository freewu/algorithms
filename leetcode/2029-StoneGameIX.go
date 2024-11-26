package main

// 2029. Stone Game IX
// Alice and Bob continue their games with stones. 
// There is a row of n stones, and each stone has an associated value. 
// You are given an integer array stones, where stones[i] is the value of the ith stone.

// Alice and Bob take turns, with Alice starting first. 
// On each turn, the player may remove any stone from stones. 
// The player who removes a stone loses if the sum of the values of all removed stones is divisible by 3. 
// Bob will win automatically if there are no remaining stones (even if it is Alice's turn).

// Assuming both players play optimally, return true if Alice wins and false if Bob wins.

// Example 1:
// Input: stones = [2,1]
// Output: true
// Explanation: The game will be played as follows:
// - Turn 1: Alice can remove either stone.
// - Turn 2: Bob removes the remaining stone. 
// The sum of the removed stones is 1 + 2 = 3 and is divisible by 3. Therefore, Bob loses and Alice wins the game.

// Example 2:
// Input: stones = [2]
// Output: false
// Explanation: Alice will remove the only stone, and the sum of the values on the removed stones is 2. 
// Since all the stones are removed and the sum of values is not divisible by 3, Bob wins the game.

// Example 3:
// Input: stones = [5,1,2,4,3]
// Output: false
// Explanation: Bob will always win. One possible way for Bob to win is shown below:
// - Turn 1: Alice can remove the second stone with value 1. Sum of removed stones = 1.
// - Turn 2: Bob removes the fifth stone with value 3. Sum of removed stones = 1 + 3 = 4.
// - Turn 3: Alices removes the fourth stone with value 4. Sum of removed stones = 1 + 3 + 4 = 8.
// - Turn 4: Bob removes the third stone with value 2. Sum of removed stones = 1 + 3 + 4 + 2 = 10.
// - Turn 5: Alice removes the first stone with value 5. Sum of removed stones = 1 + 3 + 4 + 2 + 5 = 15.
// Alice loses the game because the sum of the removed stones (15) is divisible by 3. Bob wins the game.
 
// Constraints:
//     1 <= stones.length <= 10^5
//     1 <= stones[i] <= 10^4

import "fmt"

func stoneGameIX(stones []int) bool {
    if len(stones) == 1 { return false }
    one, two, three := 0, 0, 0
    for _, stone := range stones {
        if stone % 3 == 1 {
            one++
        } else if stone % 3 == 2 {
            two++
        } else {
            three++
        }
    }
    if three % 2 == 0 { 
        return one != 0 && two != 0
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    return (one >= 3 && two == 0) || (two >= 3 && one == 0) || (abs(one - two) >= 3)
}

func stoneGameIX1(stones []int) bool {
    count1 := [3]int{}
    for _, v := range stones {
        count1[v % 3]++
    }
    count2 := [3]int{ count1[0], count1[2], count1[1] }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    check := func(count [3]int) bool {
        if count[1] == 0 { return false }
        count[1]--
        res := 1 + min(count[1], count[2]) * 2 + count[0]
        if count[1] > count[2] {
            count[1]--
            res++
        }
        return res % 2 == 1 && count[1] != count[2]
    }
    return check(count1) || check(count2)
}

func main() {
    // Example 1:
    // Input: stones = [2,1]
    // Output: true
    // Explanation: The game will be played as follows:
    // - Turn 1: Alice can remove either stone.
    // - Turn 2: Bob removes the remaining stone. 
    // The sum of the removed stones is 1 + 2 = 3 and is divisible by 3. Therefore, Bob loses and Alice wins the game.
    fmt.Println(stoneGameIX([]int{2,1})) // true
    // Example 2:
    // Input: stones = [2]
    // Output: false
    // Explanation: Alice will remove the only stone, and the sum of the values on the removed stones is 2. 
    // Since all the stones are removed and the sum of values is not divisible by 3, Bob wins the game.
    fmt.Println(stoneGameIX([]int{2})) // false
    // Example 3:
    // Input: stones = [5,1,2,4,3]
    // Output: false
    // Explanation: Bob will always win. One possible way for Bob to win is shown below:
    // - Turn 1: Alice can remove the second stone with value 1. Sum of removed stones = 1.
    // - Turn 2: Bob removes the fifth stone with value 3. Sum of removed stones = 1 + 3 = 4.
    // - Turn 3: Alices removes the fourth stone with value 4. Sum of removed stones = 1 + 3 + 4 = 8.
    // - Turn 4: Bob removes the third stone with value 2. Sum of removed stones = 1 + 3 + 4 + 2 = 10.
    // - Turn 5: Alice removes the first stone with value 5. Sum of removed stones = 1 + 3 + 4 + 2 + 5 = 15.
    // Alice loses the game because the sum of the removed stones (15) is divisible by 3. Bob wins the game.
    fmt.Println(stoneGameIX([]int{5,1,2,4,3})) // false

    fmt.Println(stoneGameIX1([]int{2,1})) // true
    fmt.Println(stoneGameIX1([]int{2})) // false
    fmt.Println(stoneGameIX1([]int{5,1,2,4,3})) // false
}