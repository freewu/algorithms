package main 

// 649. Dota2 Senate
// In the world of Dota2, there are two parties: the Radiant and the Dire.

// The Dota2 senate consists of senators coming from two parties. 
// Now the Senate wants to decide on a change in the Dota2 game. 
// The voting for this change is a round-based procedure. 
// In each round, each senator can exercise one of the two rights:
//     Ban one senator's right: A senator can make another senator lose all his rights in this and all the following rounds.
//     Announce the victory: If this senator found the senators who still have rights to vote are all from the same party, he can announce the victory and decide on the change in the game.

// Given a string senate representing each senator's party belonging. 
// The character 'R' and 'D' represent the Radiant party and the Dire party. 
// Then if there are n senators, the size of the given string will be n.

// The round-based procedure starts from the first senator to the last senator in the given order. 
// This procedure will last until the end of voting. All the senators who have lost their rights will be skipped during the procedure.

// Suppose every senator is smart enough and will play the best strategy for his own party. 
// Predict which party will finally announce the victory and change the Dota2 game. 
// The output should be "Radiant" or "Dire".

// Example 1:
// Input: senate = "RD"
// Output: "Radiant"
// Explanation: 
// The first senator comes from Radiant and he can just ban the next senator's right in round 1. 
// And the second senator can't exercise any rights anymore since his right has been banned. 
// And in round 2, the first senator can just announce the victory since he is the only guy in the senate who can vote.

// Example 2:
// Input: senate = "RDD"
// Output: "Dire"
// Explanation: 
// The first senator comes from Radiant and he can just ban the next senator's right in round 1. 
// And the second senator can't exercise any rights anymore since his right has been banned. 
// And the third senator comes from Dire and he can ban the first senator's right in round 1. 
// And in round 2, the third senator can just announce the victory since he is the only guy in the senate who can vote.

// Constraints:
//     n == senate.length
//     1 <= n <= 10^4
//     senate[i] is either 'R' or 'D'.

import "fmt"

func predictPartyVictory(senate string) string {
    radiant, dire := []int{}, []int{}  // Slice to store the indices of Radiant and Dire senators
    for i, party := range senate { // Initialize the slices
        if party == 'R' {
            radiant = append(radiant, i)
        } else {
            dire = append(dire, i)
        }
    }
    for len(radiant) > 0 && len(dire) > 0 {
        if radiant[0] < dire[0] { // Find the first Radiant senator to ban a Dire senator
            radiant = append(radiant, radiant[0] + len(senate))
        } else {
            dire = append(dire, dire[0] + len(senate))
        }
        radiant = radiant[1:] // Remove the Radiant senator
        dire = dire[1:] // Remove the Dire senator
    }
    if len(radiant) > 0 {
        return "Radiant"
    } else {
        return "Dire"
    }
}

// 消灭的策略是，尽量消灭自己后面的对手，因为前面的对手已经使用过权利了，而后序的对手依然可以使用权利消灭自己的同伴。局部最优：有一次权利机会，就消灭自己后面的对手。全局最优：为自己的阵营赢取最大利益
// 时间复杂度O(n) 空间复杂度O(1)
func predictPartyVictory1(senate string) string {
    R, D := true, true // R = true表示本轮循环结束后，字符串里依然有R。D同理
    flag := 0  // 当flag大于0时，R在D前出现，R可以消灭D。当flag小于0时，D在R前出现，D可以消灭R
    senateByte := []byte(senate)

    for R && D { // 一旦 R 或者 D 为 false，就结束循环，说明本轮结束后只剩下 R 或者 D 了
        R, D = false, false
        for i := 0; i < len(senateByte); i++ {
            if senateByte[i] == 'R' {
                if flag < 0 {
                    senateByte[i] = 0 // 消灭R，R此时为false
                } else {
                    R = true // 如果没被消灭，本轮循环结束有R
                }
                flag++
            }
            if senateByte[i] == 'D' {
                if flag > 0 {
                    senateByte[i] = 0
                } else {
                    D = true
                }
                flag--
            }
        }
    }
    if R { return "Radiant"; }
    return "Dire"
}

func main() {
    // Example 1:
    // Input: senate = "RD"
    // Output: "Radiant"
    // Explanation: 
    // The first senator comes from Radiant and he can just ban the next senator's right in round 1. 
    // And the second senator can't exercise any rights anymore since his right has been banned. 
    // And in round 2, the first senator can just announce the victory since he is the only guy in the senate who can vote.
    fmt.Println(predictPartyVictory("RD")) // "Radiant"
    // Example 2:
    // Input: senate = "RDD"
    // Output: "Dire"
    // Explanation: 
    // The first senator comes from Radiant and he can just ban the next senator's right in round 1. 
    // And the second senator can't exercise any rights anymore since his right has been banned. 
    // And the third senator comes from Dire and he can ban the first senator's right in round 1. 
    // And in round 2, the third senator can just announce the victory since he is the only guy in the senate who can vote.
    fmt.Println(predictPartyVictory("RDD")) // "Dire"

    fmt.Println(predictPartyVictory1("RD")) // "Radiant"
    fmt.Println(predictPartyVictory1("RDD")) // "Dire"
}