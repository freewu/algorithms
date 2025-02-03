package main

// 2836. Maximize Value of Function in a Ball Passing Game
// You are given an integer array receiver of length n and an integer k. 
// n players are playing a ball-passing game.

// You choose the starting player, i. 
// The game proceeds as follows: player i passes the ball to player receiver[i], who then passes it to receiver[receiver[i]], and so on, for k passes in total. 
// The game's score is the sum of the indices of the players who touched the ball, including repetitions, i.e. i + receiver[i] + receiver[receiver[i]] + ... + receiver(k)[i].

// Return the maximum possible score.

// Notes:
//     receiver may contain duplicates.
//     receiver[i] may be equal to i.

// Example 1:
// Input: receiver = [2,0,1], k = 4
// Output: 6
// Explanation:
// Starting with player i = 2 the initial score is 2:
// Pass | Sender Index | Receiver Index | Score
// 1    |      2       |       1        |  3
// 2    |      1       |       0        |  3
// 3    |      0       |       2        |  5
// 4    |      2       |       1        |  6

// Example 2:
// Input: receiver = [1,1,1,2,3], k = 3
// Output: 10
// Explanation:
// Starting with player i = 4 the initial score is 4:
// Pass | Sender Index | Receiver Index | Score
//   1  |      4       |       3        |  7
//   2  |      3       |       2        |  9
//   3  |      2       |       1        |  10

// Constraints:
//     1 <= receiver.length == n <= 10^5
//     0 <= receiver[i] <= n - 1
//     1 <= k <= 10^10

import "fmt"
import "math/bits"

func getMaxFunctionValue(receiver []int, k int64) int64 {
    res, n := int64(-1), len(receiver)
    scores, exits := make([][]int64, 50), make([][]int, 50)
    for i := 0; i < 50; i++ {
        scores[i], exits[i] = make([]int64, n), make([]int, n)
    }
    for i := 0; i < n; i++ {
        scores[0][i], exits[0][i] = int64(receiver[i]), receiver[i]
    }
    for step := 1; step < 50; step++ {
        for i := 0; i < n; i++ {
            exit := exits[step - 1][i]
            exits[step][i] = exits[step - 1][exit]
            scores[step][i] = scores[step - 1][i] + scores[step - 1][exit]
        }
    }
    max := func (x, y int64) int64 { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        score, exit := int64(i), i
        for step := 0; step <= 50; step++ {
            if (1 << step) & k == 0 { continue }
            score += scores[step][exit]
            exit = exits[step][exit]
        }
        res = max(res, score)
    }
    return res
}

func getMaxFunctionValue1(receiver []int, k int64) int64 {
    res, n := int64(-1), len(receiver)
    k++
    degrees := make([]int, n)
    for _, t := range receiver {
        degrees[t]++
    }
    topos, visiteds := make([]int, 0), make([]bool, n)
    for i := range degrees {
        if degrees[i] == 0 {
            visiteds[i] = true
            topos = append(topos, i)
        }
    }
    adjanList := make([][]int, n)
    for i := 0; i < len(topos); i++ {
        re := receiver[topos[i]]
        adjanList[re] = append(adjanList[re], topos[i])
        degrees[re]--
        if degrees[re] == 0 {
            visiteds[re] = true
            topos = append(topos, re)
        }
    }
    rounds := make([][]int, 0)
    for i := 0; i < n; i++ {
        if visiteds[i] {
            continue
        }
        tmp := make([]int, 0)
        for j := i; !visiteds[j]; j = receiver[j] {
            tmp = append(tmp, j)
            visiteds[j] = true
        }
        rounds = append(rounds, tmp)
    }
    countRound := func(prefix []int64, i int, k int64) int64 {
        l := len(prefix)
        tmp := k / int64(l) * prefix[l-1]
        remainders := int(k % int64(l))
        t := prefix[(i-1+remainders+l)%l] - prefix[(i-1+l) % l]
        if t < 0 {
            t += prefix[l-1]
        }
        return tmp + t
    }
    max := func (x, y int64) int64 { if x > y { return x; }; return y; }
    var dfs func(current int, tail []int64, prefix []int64, i int, k int64) int64
    dfs = func(current int, tail []int64, prefix []int64, i int, k int64) int64 {
        res, last := int64(0), len(tail)
        tail = append(tail, 0)
        for _, next := range adjanList[current] {
            tail[last] = tail[last-1] + int64(next)
            if int64(last) >= k {
                res = max(res, tail[last]-tail[last-int(k)])
            } else {
                res = max(res, tail[last] + countRound(prefix, i, k - int64(last)))
            }
            res = max(res, dfs(next, tail, prefix, i, k))
        }
        return res
    }
    caches := make([]int64, 1, n)
    for _, r := range rounds {
        prefix := make([]int64, len(r))
        prefix[0] = int64(r[0])
        for i := 1; i < len(r); i++ {
            prefix[i] = prefix[i-1] + int64(r[i])
        }
        for i := 0; i < len(r); i++ {
            res = max(res, countRound(prefix, i, k))
            res = max(res, dfs(r[i], caches, prefix, i, k))
        }
    }
    return res
}

func getMaxFunctionValue2(receiver []int, k int64) int64 {
    type Pair struct{ pa, sum int }
    res, n, m := 0, len(receiver), bits.Len(uint(k))
    f := make([][]Pair, n)
    for i := 0; i < n; i++  {
        f[i] = make([]Pair, m)
        f[i][0] = Pair{ receiver[i], receiver[i]}
    }
    for j := 1; j < m; j++ {
        for i := 0; i < n; i++  {
            p := f[i][j-1].pa
            f[i][j] = Pair{ f[p][j-1].pa, f[i][j-1].sum + f[p][j-1].sum }
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        count, pa, j, t := i, i, 0, k
        for t > 0 {
            if t & 1 == 1 {
                count += f[pa][j].sum
                pa = f[pa][j].pa
            }
            j++
            t >>= 1
        }
        res = max(res, count)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: receiver = [2,0,1], k = 4
    // Output: 6
    // Explanation:
    // Starting with player i = 2 the initial score is 2:
    // Pass | Sender Index | Receiver Index | Score
    // 1    |      2       |       1        |  3
    // 2    |      1       |       0        |  3
    // 3    |      0       |       2        |  5
    // 4    |      2       |       1        |  6
    fmt.Println(getMaxFunctionValue([]int{2,0,1}, 4)) // 6
    // Example 2:
    // Input: receiver = [1,1,1,2,3], k = 3
    // Output: 10
    // Explanation:
    // Starting with player i = 4 the initial score is 4:
    // Pass | Sender Index | Receiver Index | Score
    //   1  |      4       |       3        |  7
    //   2  |      3       |       2        |  9
    //   3  |      2       |       1        |  10
    fmt.Println(getMaxFunctionValue([]int{1,1,1,2,3}, 3)) // 10

    fmt.Println(getMaxFunctionValue1([]int{2,0,1}, 4)) // 6
    fmt.Println(getMaxFunctionValue1([]int{1,1,1,2,3}, 3)) // 10

    fmt.Println(getMaxFunctionValue2([]int{2,0,1}, 4)) // 6
    fmt.Println(getMaxFunctionValue2([]int{1,1,1,2,3}, 3)) // 10
}