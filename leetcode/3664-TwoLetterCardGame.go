package main

// 3664. Two-Letter Card Game
// You are given a deck of cards represented by a string array cards, and each card displays two lowercase letters.

// You are also given a letter x. You play a game with the following rules:
//     1. Start with 0 points.
//     2. On each turn, you must find two compatible cards from the deck that both contain the letter x in any position.
//     3. Remove the pair of cards and earn 1 point.
//     4. The game ends when you can no longer find a pair of compatible cards.

// Return the maximum number of points you can gain with optimal play.

// Two cards are compatible if the strings differ in exactly 1 position.

// Example 1:
// Input: cards = ["aa","ab","ba","ac"], x = "a"
// Output: 2
// Explanation:
// On the first turn, select and remove cards "ab" and "ac", which are compatible because they differ at only index 1.
// On the second turn, select and remove cards "aa" and "ba", which are compatible because they differ at only index 0.
// Because there are no more compatible pairs, the total score is 2.

// Example 2:
// Input: cards = ["aa","ab","ba"], x = "a"
// Output: 1
// Explanation:
// On the first turn, select and remove cards "aa" and "ba".
// Because there are no more compatible pairs, the total score is 1.

// Example 3:
// Input: cards = ["aa","ab","ba","ac"], x = "b"
// Output: 0
// Explanation:
// The only cards that contain the character 'b' are "ab" and "ba". However, they differ in both indices, so they are not compatible. Thus, the output is 0.

// Constraints:
//     2 <= cards.length <= 10^5
//     cards[i].length == 2
//     Each cards[i] is composed of only lowercase English letters between 'a' and 'j'.
//     x is a lowercase English letter between 'a' and 'j'.

import "fmt"

func score(cards []string, x byte) int {
    count1, count2 := [10]int{}, [10]int{}
    for _, s := range cards {
        if s[0] == x {
            count1[s[1]-'a']++
        }
        if s[1] == x {
            count2[s[0]-'a']++
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    calc := func(cnt []int, x byte) (int, int) {  // 计算这一组的得分（配对个数），以及剩余元素个数
        sum, mx := 0, 0
        for i, c := range cnt {
            if i != int(x-'a') {
                sum += c
                mx = max(mx, c)
            }
        }
        pairs := min(sum / 2, sum - mx)
        return pairs, sum - pairs * 2
    }
    pairs1, left1 := calc(count1[:], x)
    pairs2, left2 := calc(count2[:], x)
    res := pairs1 + pairs2 // 不考虑 xx 时的得分
    v := count1[x-'a']
    // 把 xx 和剩下的 x? 和 ?x 配对
    // 每有 1 个 xx，得分就能增加一，但这不能超过剩下的 x? 和 ?x 的个数 left1+left2
    if v > 0 {
        mn := min(v, left1 + left2)
        res += mn
        v -= mn
    }
    // 如果还有 xx，就撤销之前的配对，比如 (ax,bx) 改成 (ax,xx) 和 (bx,xx)
    // 每有 2 个 xx，得分就能增加一，但这不能超过之前的配对个数 pairs1+pairs2
    // 由于这种方案平均每个 xx 只能增加 0.5 分，不如上面的，所以先考虑把 xx 和剩下的 x? 和 ?x 配对，再考虑撤销之前的配对
    if v > 0 {
        res += min(v/2, pairs1 + pairs2)
    }
    return res
}

func main() {
    // Example 1:
    // Input: cards = ["aa","ab","ba","ac"], x = "a"
    // Output: 2
    // Explanation:
    // On the first turn, select and remove cards "ab" and "ac", which are compatible because they differ at only index 1.
    // On the second turn, select and remove cards "aa" and "ba", which are compatible because they differ at only index 0.
    // Because there are no more compatible pairs, the total score is 2.
    fmt.Println(score([]string{"aa","ab","ba","ac"}, 'a')) // 2
    // Example 2:
    // Input: cards = ["aa","ab","ba"], x = "a"
    // Output: 1
    // Explanation:
    // On the first turn, select and remove cards "aa" and "ba".
    // Because there are no more compatible pairs, the total score is 1.
    fmt.Println(score([]string{"aa","ab","ba"}, 'a')) // 1 
    // Example 3:
    // Input: cards = ["aa","ab","ba","ac"], x = "b"
    // Output: 0
    // Explanation:
    // The only cards that contain the character 'b' are "ab" and "ba". However, they differ in both indices, so they are not compatible. Thus, the output is 0.
    fmt.Println(score([]string{"aa","ab","ba","ac"}, 'b')) // 0
}