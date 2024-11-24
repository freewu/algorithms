package main

// 1996. The Number of Weak Characters in the Game
// You are playing a game that contains multiple characters, and each of the characters has two main properties: attack and defense. 
// You are given a 2D integer array properties where properties[i] = [attacki, defensei] represents the properties of the ith character in the game.

// A character is said to be weak if any other character has both attack and defense levels strictly greater than this character's attack and defense levels. 
// More formally, a character i is said to be weak if there exists another character j where attackj > attacki and defensej > defensei.

// Return the number of weak characters.

// Example 1:
// Input: properties = [[5,5],[6,3],[3,6]]
// Output: 0
// Explanation: No character has strictly greater attack and defense than the other.

// Example 2:
// Input: properties = [[2,2],[3,3]]
// Output: 1
// Explanation: The first character is weak because the second character has a strictly greater attack and defense.

// Example 3:
// Input: properties = [[1,5],[10,4],[4,3]]
// Output: 1
// Explanation: The third character is weak because the second character has a strictly greater attack and defense.

// Constraints:
//     2 <= properties.length <= 10^5
//     properties[i].length == 2
//     1 <= attacki, defensei <= 10^5

import "fmt"
import "sort"

func numberOfWeakCharacters(properties [][]int) int {
    sort.Slice(properties, func(i, j int) bool {
        if properties[i][0] != properties[j][0] { return properties[i][0] > properties[j][0] } // 攻击越强越靠前
        return properties[i][1] < properties[j][1] // 防御越弱越靠前
    })	
    res, mx := 0, -1
    for _, v := range properties {
        if v[1] < mx { // 如果认为角色 i 弱于 存在的另一个角色 j ，那么 attackj > attacki 且 defensej > defensei
            res++
        } else {
            mx = v[1]
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: properties = [[5,5],[6,3],[3,6]]
    // Output: 0
    // Explanation: No character has strictly greater attack and defense than the other.
    fmt.Println(numberOfWeakCharacters([][]int{{5,5},{6,3},{3,6}})) // 0
    // Example 2:
    // Input: properties = [[2,2],[3,3]]
    // Output: 1
    // Explanation: The first character is weak because the second character has a strictly greater attack and defense.
    fmt.Println(numberOfWeakCharacters([][]int{{2,2},{3,3}})) // 1
    // Example 3:
    // Input: properties = [[1,5],[10,4],[4,3]]
    // Output: 1
    // Explanation: The third character is weak because the second character has a strictly greater attack and defense.
    fmt.Println(numberOfWeakCharacters([][]int{{1,5},{10,4},{4,3}})) // 1
}