package main

// 1744. Can You Eat Your Favorite Candy on Your Favorite Day?
// You are given a (0-indexed) array of positive integers candiesCount where candiesCount[i] represents the number of candies of the ith type you have. 
// You are also given a 2D array queries where queries[i] = [favoriteTypei, favoriteDayi, dailyCapi].

// You play a game with the following rules:
//     You start eating candies on day 0.
//     You cannot eat any candy of type i unless you have eaten all candies of type i - 1.
//     You must eat at least one candy per day until you have eaten all the candies.

// Construct a boolean array answer such that answer.length == queries.length 
// and answer[i] is true if you can eat a candy of type favoriteTypei on day favoriteDayi without eating more than dailyCapi candies on any day, and false otherwise. 
// Note that you can eat different types of candy on the same day, provided that you follow rule 2.

// Return the constructed array answer.

// Example 1:
// Input: candiesCount = [7,4,5,3,8], queries = [[0,2,2],[4,2,4],[2,13,1000000000]]
// Output: [true,false,true]
// Explanation:
// 1- If you eat 2 candies (type 0) on day 0 and 2 candies (type 0) on day 1, you will eat a candy of type 0 on day 2.
// 2- You can eat at most 4 candies each day.
//    If you eat 4 candies every day, you will eat 4 candies (type 0) on day 0 and 4 candies (type 0 and type 1) on day 1.
//    On day 2, you can only eat 4 candies (type 1 and type 2), so you cannot eat a candy of type 4 on day 2.
// 3- If you eat 1 candy each day, you will eat a candy of type 2 on day 13.

// Example 2:
// Input: candiesCount = [5,2,6,4,1], queries = [[3,1,2],[4,10,3],[3,10,100],[4,100,30],[1,3,1]]
// Output: [false,true,true,false,false]

// Constraints:
//     1 <= candiesCount.length <= 10^5
//     1 <= candiesCount[i] <= 10^5
//     1 <= queries.length <= 10^5
//     queries[i].length == 3
//     0 <= favoriteTypei < candiesCount.length
//     0 <= favoriteDayi <= 10^9
//     1 <= dailyCapi <= 10^9

import "fmt"

// Prefix Sum
func canEat(candiesCount []int, queries [][]int) []bool {
    prefixSum, index := make([]int, len(candiesCount)), 1
    res := make([]bool, len(queries))
    for i, query := range queries {
        favType, favDay := query[0], query[1]
        if favType >= index { // process only up to the favorite type
            for i := index; i <= favType; i++ {
                prefixSum[i] = prefixSum[i-1] + candiesCount[i - 1]
            }
            index = favType + 1
        }
        // the number of candies required to eat
        mn, mx := (prefixSum[favType] + 1), (prefixSum[favType] + candiesCount[favType])
        // the number of candies you can eat based on the cap
        minEat := favDay + 1
        maxEat := minEat * query[2]
        if mn <= maxEat && mx >= minEat { // these ranges must overlap
            res[i] = true
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: candiesCount = [7,4,5,3,8], queries = [[0,2,2],[4,2,4],[2,13,1000000000]]
    // Output: [true,false,true]
    // Explanation:
    // 1- If you eat 2 candies (type 0) on day 0 and 2 candies (type 0) on day 1, you will eat a candy of type 0 on day 2.
    // 2- You can eat at most 4 candies each day.
    //    If you eat 4 candies every day, you will eat 4 candies (type 0) on day 0 and 4 candies (type 0 and type 1) on day 1.
    //    On day 2, you can only eat 4 candies (type 1 and type 2), so you cannot eat a candy of type 4 on day 2.
    // 3- If you eat 1 candy each day, you will eat a candy of type 2 on day 13.
    fmt.Println(canEat([]int{7,4,5,3,8}, [][]int{{0,2,2},{4,2,4},{2,13,1000000000}})) // [true,false,true]
    // Example 2:
    // Input: candiesCount = [5,2,6,4,1], queries = [[3,1,2],[4,10,3],[3,10,100],[4,100,30],[1,3,1]]
    // Output: [false,true,true,false,false]
    fmt.Println(canEat([]int{5,2,6,4,1}, [][]int{{3,1,2},{4,10,3},{3,10,100},{4,100,30},{1,3,1}})) // [false,true,true,false,false]
}