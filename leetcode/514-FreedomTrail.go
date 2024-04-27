package main

// 514. Freedom Trail
// In the video game Fallout 4, the quest "Road to Freedom" requires players to reach a metal dial called the "Freedom Trail Ring" and use the dial to spell a specific keyword to open the door.
// Given a string ring that represents the code engraved on the outer ring and another string key that represents the keyword that needs to be spelled, return the minimum number of steps to spell all the characters in the keyword.
// Initially, the first character of the ring is aligned at the "12:00" direction. You should spell all the characters in key one by one by rotating ring clockwise or anticlockwise to make each character of the string key aligned at the "12:00" direction and then by pressing the center button.

// At the stage of rotating the ring to spell the key character key[i]:
//     You can rotate the ring clockwise or anticlockwise by one place, which counts as one step. The final purpose of the rotation is to align one of ring's characters at the "12:00" direction, where this character must equal key[i].
//     If the character key[i] has been aligned at the "12:00" direction, press the center button to spell, which also counts as one step. After the pressing, you could begin to spell the next character in the key (next stage). Otherwise, you have finished all the spelling.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2018/10/22/ring.jpg">
// Input: ring = "godding", key = "gd"
// Output: 4
// Explanation:
// For the first key character 'g', since it is already in place, we just need 1 step to spell this character. 
// For the second key character 'd', we need to rotate the ring "godding" anticlockwise by two steps to make it become "ddinggo".
// Also, we need 1 more step for spelling.
// So the final output is 4.

// Example 2:
// Input: ring = "godding", key = "godding"
// Output: 13
 
// Constraints:
//     1 <= ring.length, key.length <= 100
//     ring and key consist of only lower case English letters.
//     It is guaranteed that key could always be spelled by rotating ring.

// func findRotateSteps(ring string, key string) int {
// 	n := len(ring)
// 	m := len(key)
// 	dp := make([][]int, m + 1)
// 	for i := 0; i < m + 1; i++ {
// 		dp[i] = make([]int, len(n)) 
// 	}
// 	// int n = ring.length();
// 	// int m = key.length();
// 	// int[][] dp = new int[m+1][n];

// 	for(int i = 1;i<m+1;i++){
// 		for(int j = 0;j<n;j++){
// 			dp[i][j]=Integer.MAX_VALUE;
// 			for(int k=0;k<n;k++){
// 				if(ring.charAt(k)==key.charAt(m-i)){
// 					int diff = Math.abs(j-k);
// 					int step = Math.min(diff,n-diff);
// 					dp[i][j] = Math.min(dp[i][j],step+dp[i-1][k]);
// 				}
// 			}
// 		}
// 	}

// 	return dp[m][0]+m;
// }

import "fmt"

func findRotateSteps(ring string, key string) int {
    n, m, pos, inf := len(ring), len(key), [26][]int{}, 1 << 32 - 1
    for i, c := range ring {
        pos[c-'a'] = append(pos[c-'a'], i)
    }

    min := func (a ...int) int {
        res := a[0]
        for _, v := range a[1:] {
            if v < res {
                res = v
            }
        }
        return res
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
        for j := range dp[i] {
            dp[i][j] = inf
        }
    }
    for _, p := range pos[key[0]-'a'] {
        dp[0][p] = min(p, n-p) + 1
    }
    for i := 1; i < m; i++ {
        for _, j := range pos[key[i]-'a'] {
            for _, k := range pos[key[i-1]-'a'] {
                dp[i][j] = min(dp[i][j], dp[i-1][k] + min(abs(j-k), n - abs(j-k)) + 1)
            }
        }
    }
    return min(dp[m-1]...)
}

func main() {
    // Input: ring = "godding", key = "gd"
    // Output: 4
    // Explanation:
    // For the first key character 'g', since it is already in place, we just need 1 step to spell this character. 
    // For the second key character 'd', we need to rotate the ring "godding" anticlockwise by two steps to make it become "ddinggo".
    // Also, we need 1 more step for spelling.
    // So the final output is 4.
    fmt.Println(findRotateSteps("godding","gd")) // 4
    // Example 2:
    // Input: ring = "godding", key = "godding"
    // Output: 13
    fmt.Println(findRotateSteps("godding","godding")) // 13
}