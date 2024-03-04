package main

// 948. Bag of Tokens
// You start with an initial power of power, an initial score of 0, 
// and a bag of tokens given as an integer array tokens, where each tokens[i] donates the value of tokeni.

// Your goal is to maximize the total score by strategically playing these tokens. 
// n one move, you can play an unplayed token in one of the two ways (but not both for the same token):
//         Face-up: If your current power is at least tokens[i], you may play tokeni, losing tokens[i] power and gaining 1 score.
//         Face-down: If your current score is at least 1, you may play tokeni, gaining tokens[i] power and losing 1 score.

// Return the maximum possible score you can achieve after playing any number of tokens.

// Example 1:
// Input: tokens = [100], power = 50
// Output: 0
// Explanation: 
//     Since your score is 0 initially, you cannot play the token face-down. 
//     You also cannot play it face-up since your power (50) is less than tokens[0] (100).

// Example 2:
// Input: tokens = [200,100], power = 150
// Output: 1
// Explanation: 
//     Play token1 (100) face-up, reducing your power to 50 and increasing your score to 1.
//     There is no need to play token0, since you cannot play it face-up to add to your score. 
//     The maximum score achievable is 1.

// Example 3:
// Input: tokens = [100,200,300,400], power = 200
// Output: 2
// Explanation: 
//     Play the tokens in this order to get a score of 2:
//         Play token0 (100) face-up, reducing power to 100 and increasing score to 1.
//         Play token3 (400) face-down, increasing power to 500 and reducing score to 0.
//         Play token1 (200) face-up, reducing power to 300 and increasing score to 1.
//         Play token2 (300) face-up, reducing power to 0 and increasing score to 2.
//         The maximum score achievable is 2.

// Constraints:
//         0 <= tokens.length <= 1000
//         0 <= tokens[i], power < 10^4

import "fmt"
import "sort"

// 贪心
// 思路:
//      如果让我们来玩令牌放置这个游戏，在让令牌正面朝上的时候，肯定要去找能量最小的令牌。
//      同样的，在让令牌反面朝上的时候，肯定要去找能量最大的令牌
func bagOfTokensScore(tokens []int, P int) int {
    // 升序排列
	sort.Ints(tokens)
	maxScore,score := 0,0
    // 双指针法
	l, r := 0, len(tokens)-1
	for l <= r {
        // 如果能量 > 左指针对应值（即当前能获得的1分需要的最少代价），能量换分值，记录最大值；
		if P >= tokens[l] {
            // 用能量换得分
			score++
			P -= tokens[l]
			l++
			if score > maxScore {
				maxScore = score
			}
		} else if score > 0 { // 如果当前得分>0, 换取能量，可能谋取更多分值
            // 用得分换能量
			score--
			P += tokens[r]
			r--
		} else {
			break
		}
	}
	return maxScore
}

func bagOfTokensScore1(tokens []int, power int) int {
    sort.Ints(tokens)
    n := len(tokens)
    result := 0
    ans := 0
    i:=0
    j:=n-1
    for i <= j {
        if power >= tokens[i] {
            power -= tokens[i]
            i++
            ans++
            result = max(result, ans)
        } else {
            if ans > 0 {
                ans--
                power += tokens[j]
                j--
            } else {
                break
            }
        }
    }
    return result
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    // Since your score is 0 initially, you cannot play the token face-down. 
    // You also cannot play it face-up since your power (50) is less than tokens[0] (100).
    fmt.Println(bagOfTokensScore([]int{100}, 50)) // 0
    // Play token1 (100) face-up, reducing your power to 50 and increasing your score to 1.
    // There is no need to play token0, since you cannot play it face-up to add to your score. 
    // The maximum score achievable is 1.
    fmt.Println(bagOfTokensScore([]int{200,100}, 150)) // 1

//     Play the tokens in this order to get a score of 2:
//         Play token0 (100) face-up, reducing power to 100 and increasing score to 1.
//         Play token3 (400) face-down, increasing power to 500 and reducing score to 0.
//         Play token1 (200) face-up, reducing power to 300 and increasing score to 1.
//         Play token2 (300) face-up, reducing power to 0 and increasing score to 2.
//         The maximum score achievable is 2.
    fmt.Println(bagOfTokensScore([]int{100,200,300,400}, 200)) // 2


    fmt.Println(bagOfTokensScore1([]int{100}, 50)) // 0
    fmt.Println(bagOfTokensScore1([]int{200,100}, 150)) // 1
    fmt.Println(bagOfTokensScore1([]int{100,200,300,400}, 200)) // 2
}