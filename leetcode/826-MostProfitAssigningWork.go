package main

// 826. Most Profit Assigning Work
// You have n jobs and m workers. 
// You are given three arrays: difficulty, profit, and worker where:
//     difficulty[i] and profit[i] are the difficulty and the profit of the ith job, and
//     worker[j] is the ability of jth worker (i.e., the jth worker can only complete a job with difficulty at most worker[j]).

// Every worker can be assigned at most one job, but one job can be completed multiple times.
//     For example, if three workers attempt the same job that pays $1, then the total profit will be $3. 
//     If a worker cannot complete any job, their profit is $0.

// Return the maximum profit we can achieve after assigning the workers to the jobs.

// Example 1:
// Input: difficulty = [2,4,6,8,10], profit = [10,20,30,40,50], worker = [4,5,6,7]
// Output: 100
// Explanation: Workers are assigned jobs of difficulty [4,4,6,6] and they get a profit of [20,20,30,30] separately.

// Example 2:
// Input: difficulty = [85,47,57], profit = [24,66,99], worker = [40,25,25]
// Output: 0
 
// Constraints:
//     n == difficulty.length
//     n == profit.length
//     m == worker.length
//     1 <= n, m <= 10^4
//     1 <= difficulty[i], profit[i], worker[i] <= 10^5

import "fmt"
import "sort"

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
    type JobPair struct {
        Difficulty int
        Profit     int
    }
    DifficultyAndProfits := []JobPair{}
    for i := 0; i < len(difficulty); i++ { // 将工作难度和收益关联起来
        DifficultyAndProfits = append(DifficultyAndProfits, JobPair{ Difficulty: difficulty[i], Profit: profit[i]})
    }
    sort.Slice(DifficultyAndProfits, func(i, j int) bool {
        return DifficultyAndProfits[i].Difficulty < DifficultyAndProfits[j].Difficulty
    })
    sort.Ints(worker)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // tricky, we dont reset them to 0 for each worker!
    // currentMaxProfit will hold the best profit till current j, and we dont need to iterate again and again.
    // We can just move forward and check if there is any greater profit value available if worker capability is greater than the difficulty of next value
    maxProfit, j, currentMaxProfit := 0, 0, 0
    for i := 0; i < len(worker); i++ {
        for j < len(DifficultyAndProfits) && worker[i] >= DifficultyAndProfits[j].Difficulty {
            currentMaxProfit = max(currentMaxProfit, DifficultyAndProfits[j].Profit)
            j++
        }
        maxProfit += currentMaxProfit
    }
    return maxProfit
}

func maxProfitAssignment1(difficulty []int, profit []int, worker []int) int {
    jobs := make([][2]int, len(difficulty))
    for i := range difficulty { // 将工作的难度&收益关联起来 index 
        jobs[i] = [2]int{difficulty[i], profit[i]}
    }
    sort.Slice(jobs, func(i, j int) bool { return jobs[i][0] < jobs[j][0] }) // 按难度从容易到难
    sort.Ints(worker) // 工人能力从小到大

    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, i, best := 0, 0, 0
    for _, w := range worker {
        for i < len(jobs) && w >= jobs[i][0] {
            best = max(best, jobs[i][1])
            i++
        }
        res += best
    }
    return res
}

func main() {
    // Example 1:
    // Input: difficulty = [2,4,6,8,10], profit = [10,20,30,40,50], worker = [4,5,6,7]
    // Output: 100
    // Explanation: Workers are assigned jobs of difficulty [4,4,6,6] and they get a profit of [20,20,30,30] separately.
    fmt.Println(maxProfitAssignment([]int{2,4,6,8,10},[]int{10,20,30,40,50},[]int{4,5,6,7})) // 100
    // Example 2:
    // Input: difficulty = [85,47,57], profit = [24,66,99], worker = [40,25,25]
    // Output: 0
    fmt.Println(maxProfitAssignment([]int{85,47,57},[]int{24,66,99},[]int{40,25,25})) // 0

    fmt.Println(maxProfitAssignment1([]int{2,4,6,8,10},[]int{10,20,30,40,50},[]int{4,5,6,7})) // 100
    fmt.Println(maxProfitAssignment1([]int{85,47,57},[]int{24,66,99},[]int{40,25,25})) // 0
}