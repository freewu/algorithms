package main

import "fmt"
import "slices"
import "cmp"

// 1686. Stone Game VI
// Alice and Bob take turns playing a game, with Alice starting first.
// There are n stones in a pile. On each player's turn,
// they can remove a stone from the pile and receive points based on the stone's value. 
// Alice and Bob may value the stones differently.
// You are given two integer arrays of length n, aliceValues and bobValues.
// Each aliceValues[i] and bobValues[i] represents how Alice and Bob, respectively, value the ith stone.
// The winner is the person with the most points after all the stones are chosen. 
// If both players have the same amount of points, the game results in a draw. 
// Both players will play optimally. Both players know the other's values.

// Determine the result of the game, and:

// 		If Alice wins, return 1.
// 		If Bob wins, return -1.
// 		If the game results in a draw, return 0.
 

// Example 1:
// Input: aliceValues = [1,3], bobValues = [2,1]
// Output: 1
// Explanation:
// If Alice takes stone 1 (0-indexed) first, Alice will receive 3 points.
// Bob can only choose stone 0, and will only receive 2 points.
// Alice wins.

// Example 2:
// Input: aliceValues = [1,2], bobValues = [3,1]
// Output: 0
// Explanation:
// If Alice takes stone 0, and Bob takes stone 1, they will both have 1 point.
// Draw.

// Example 3:
// Input: aliceValues = [2,4,3], bobValues = [1,6,7]
// Output: -1
// Explanation:
// Regardless of how Alice plays, Bob will be able to have more points than Alice.
// For example, if Alice takes stone 1, Bob can take stone 2, and Alice takes stone 0, Alice will have 6 points to Bob's 7.
// Bob wins.
 
// Constraints:
// 		n == aliceValues.length == bobValues.length
// 		1 <= n <= 10^5
// 		1 <= aliceValues[i], bobValues[i] <= 100

func stoneGameVI(aliceValues []int, bobValues []int) int {
	vals := make([][2]int, len(aliceValues))
	for i, a := range aliceValues {
		vals[i] = [2]int{a + bobValues[i], i}
	}
	slices.SortFunc(vals, func(a, b [2]int) int { return b[0] - a[0] })
	a, b := 0, 0
	for k, v := range vals {
		i := v[1]
		if k%2 == 0 {
			a += aliceValues[i]
		} else {
			b += bobValues[i]
		}
	}
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}


// best solution
func stoneGameVI1(a, b []int) int {
    type pair struct{ x, y int }
    pairs := make([]pair, len(a))
    for i, x := range a {
        pairs[i] = pair{x, b[i]}
    }
    slices.SortFunc(pairs, func(p, q pair) int { return q.x + q.y - p.x - p.y })
    diff := 0
    for i, p := range pairs {
        if i%2 == 0 {
            diff += p.x
        } else {
            diff -= p.y
        }
    }
    return cmp.Compare(diff, 0)
}

func main() {
	fmt.Println(stoneGameVI([]int{1,3},[]int{2,1})) // 1
	fmt.Println(stoneGameVI([]int{1,2},[]int{3,1})) // 0	
	fmt.Println(stoneGameVI([]int{2,4,3},[]int{1,6,7})) // -1	

	fmt.Println(stoneGameVI1([]int{1,3},[]int{2,1})) // 1
	fmt.Println(stoneGameVI1([]int{1,2},[]int{3,1})) // 0	
	fmt.Println(stoneGameVI1([]int{2,4,3},[]int{1,6,7})) // -1	
}