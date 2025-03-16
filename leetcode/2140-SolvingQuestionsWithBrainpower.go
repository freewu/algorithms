package main

// 2140. Solving Questions With Brainpower
// You are given a 0-indexed 2D integer array questions where questions[i] = [pointsi, brainpoweri].
// The array describes the questions of an exam, where you have to process the questions in order (i.e., starting from question 0) and make a decision whether to solve or skip each question. 
// Solving question i will earn you pointsi points but you will be unable to solve each of the next brainpoweri questions. 
// If you skip question i, you get to make the decision on the next question.

// For example, given questions = [[3, 2], [4, 3], [4, 4], [2, 5]]:
//     If question 0 is solved, you will earn 3 points but you will be unable to solve questions 1 and 2.
//     If instead, question 0 is skipped and question 1 is solved, you will earn 4 points but you will be unable to solve questions 2 and 3.

// Return the maximum points you can earn for the exam.

// Example 1:
// Input: questions = [[3,2],[4,3],[4,4],[2,5]]
// Output: 5
// Explanation: The maximum points can be earned by solving questions 0 and 3.
// - Solve question 0: Earn 3 points, will be unable to solve the next 2 questions
// - Unable to solve questions 1 and 2
// - Solve question 3: Earn 2 points
// Total points earned: 3 + 2 = 5. There is no other way to earn 5 or more points.

// Example 2:
// Input: questions = [[1,1],[2,2],[3,3],[4,4],[5,5]]
// Output: 7
// Explanation: The maximum points can be earned by solving questions 1 and 4.
// - Skip question 0
// - Solve question 1: Earn 2 points, will be unable to solve the next 2 questions
// - Unable to solve questions 2 and 3
// - Solve question 4: Earn 5 points
// Total points earned: 2 + 5 = 7. There is no other way to earn 7 or more points.

// Constraints:
//     1 <= questions.length <= 10^5
//     questions[i].length == 2
//     1 <= pointsi, brainpoweri <= 10^5

import "fmt"

// dp
func mostPoints(questions [][]int) int64 {
    dp := make([]int, len(questions))
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func (questions [][]int, index, temp int, dp []int) int
    dfs = func (questions [][]int, index, temp int, dp []int) int {
        if int(index) >= len(questions) {
            return temp
        }
        if dp[index] != 0 {
            return dp[index] + temp
        }
        skip := dfs(questions, index + 1, temp, dp) // 需要发弃的分数
        get := dfs(questions, index + questions[index][1] + 1, temp + questions[index][0], dp) // 需要获得的分数
        res := max(skip, get)
        dp[index] = res
        return res
    }
    return int64(dfs(questions,0,0, dp))
}

func mostPoints1(questions [][]int) int64 {
    l := len(questions)
    dp := make([]int, l)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    dp[l-1] = questions[l-1][0]
    for i := l-2; i >= 0; i-- {
        next := i + questions[i][1] + 1
        if next < l {
            dp[i] = max(questions[i][0]+dp[next], dp[i+1])
        } else {
            dp[i] = max(questions[i][0], dp[i+1])
        }
    }
    return int64(dp[0])
}

func mostPoints2(questions [][]int) int64 {
    res, n := 0, len(questions)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := n - 1; i >= 0; i-- {
         if i + questions[i][1] + 1 < n {
              questions[i][0] = questions[i][0] + questions[i + questions[i][1] + 1][0]
         }
         res = max(res,questions[i][0])
         questions[i][0] = res
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: questions = [[3,2],[4,3],[4,4],[2,5]]
    // Output: 5
    // Explanation: The maximum points can be earned by solving questions 0 and 3.
    // - Solve question 0: Earn 3 points, will be unable to solve the next 2 questions
    // - Unable to solve questions 1 and 2
    // - Solve question 3: Earn 2 points
    // Total points earned: 3 + 2 = 5. There is no other way to earn 5 or more points.
    fmt.Println(mostPoints([][]int{{3,2},{4,3},{4,4},{2,5}})) // 5
    // Example 2:
    // Input: questions = [[1,1],[2,2],[3,3],[4,4],[5,5]]
    // Output: 7
    // Explanation: The maximum points can be earned by solving questions 1 and 4.
    // - Skip question 0
    // - Solve question 1: Earn 2 points, will be unable to solve the next 2 questions
    // - Unable to solve questions 2 and 3
    // - Solve question 4: Earn 5 points
    // Total points earned: 2 + 5 = 7. There is no other way to earn 7 or more points.
    fmt.Println(mostPoints([][]int{{1,1},{2,2},{3,3},{4,4},{5,5}})) // 7

    fmt.Println(mostPoints1([][]int{{3,2},{4,3},{4,4},{2,5}})) // 5
    fmt.Println(mostPoints1([][]int{{1,1},{2,2},{3,3},{4,4},{5,5}})) // 7

    fmt.Println(mostPoints2([][]int{{3,2},{4,3},{4,4},{2,5}})) // 5
    fmt.Println(mostPoints2([][]int{{1,1},{2,2},{3,3},{4,4},{5,5}})) // 7
}