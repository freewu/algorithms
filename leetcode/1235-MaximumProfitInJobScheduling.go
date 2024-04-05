package main

// 1235. Maximum Profit in Job Scheduling
// We have n jobs, where every job is scheduled to be done from startTime[i] to endTime[i], obtaining a profit of profit[i].

// You're given the startTime, endTime and profit arrays, 
// return the maximum profit you can take such that there are no two jobs in the subset with overlapping time range.

// If you choose a job that ends at time X you will be able to start another job that starts at time X.

// Example 1:
// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/10/19/sample1_1584.png" />
// Input: startTime = [1,2,3,3], endTime = [3,4,5,6], profit = [50,10,40,70]
// Output: 120
// Explanation: The subset chosen is the first and fourth job. 
// Time range [1-3]+[3-6] , we get profit of 120 = 50 + 70.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/10/10/sample22_1584.png" />
// Input: startTime = [1,2,3,4,6], endTime = [3,5,10,6,9], profit = [20,20,100,70,60]
// Output: 150
// Explanation: The subset chosen is the first, fourth and fifth job. 
// Profit obtained 150 = 20 + 70 + 60.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2019/10/10/sample3_1584.png" />
// Input: startTime = [1,1,1], endTime = [2,3,4], profit = [5,6,4]
// Output: 6
 
// Constraints:
//     1 <= startTime.length == endTime.length == profit.length <= 5 * 10^4
//     1 <= startTime[i] < endTime[i] <= 10^9
//     1 <= profit[i] <= 10^4

import "fmt"
import "sort"

// dp
func jobScheduling(startTime []int, endTime []int, profit []int) int {
    dp := make([]int, len(startTime)+1)
    sum := make([][3]int, len(startTime))
    for i := 0; i < len(startTime); i++ {
        sum[i][0], sum[i][1], sum[i][2] = startTime[i], endTime[i], profit[i]
    }
    sort.Slice(sum, func(i, j int)bool {return sum[i][0] < sum[j][0]})
    sort.Ints(startTime)
    
    for i := len(startTime) - 1; i >= 0; i-- {
        dp[i] = sum[i][2]
        index := sort.SearchInts(startTime, sum[i][1])
        if index < len(startTime) {dp[i] += dp[index]}
        if dp[i] < dp[i+1] {dp[i] = dp[i+1]}
    }
    return dp[0]
}

func jobScheduling1(startTime []int, endTime []int, profit []int) int {
    // 以i份工组为最后的工作能获取的最大利润
    // 如果打了第i份工，开始时间为start_i, 那么上一份工的结束时间一定要小于等于 start_i，
    // 利用二分查找找到最后一个endtime小于等于start_i的工作， 加上其最大利润，为当前的最大利润
    // 假设有序的
    sort.Sort(SortHelper{startTime, endTime, profit})
    max := func (x, y int) int { if x > y { return x; }; return y; }
    dp := make([]int, len(startTime)+1)
    dp[0] = 0
    res := 0
    for i := 0; i < len(startTime); i++ {
        index := sort.Search(i, func(index int)bool {
            return endTime[index] >= startTime[i] + 1
        }) - 1

        dp[i+1] = dp[i]
        cur := profit[i]
        if index >= 0 && endTime[index] <= startTime[i] {
           cur += dp[index+1]
        }
        dp[i+1] = max(dp[i+1], cur)
        res = max(res, dp[i+1])
    }
    return res
}

type SortHelper struct {
    start, end, profit []int
}

func (s SortHelper) Len() int {
    return len(s.start)
}

func(s SortHelper) Swap(i, j int) {
    s.start[i], s.start[j] = s.start[j], s.start[i]
    s.end[i], s.end[j] = s.end[j], s.end[i]
    s.profit[i], s.profit[j] = s.profit[j], s.profit[i]
}

func (s SortHelper) Less(i, j int) bool {
    return s.end[i] < s.end[j]
}

type Job struct {
    start  int // 开始时间
    end    int // 结果时间
    profit int // 报酬
}

func jobScheduling2(start []int, end []int, profit []int) int {
    n, jobs := len(start), make([]Job, len(start))
    for i := 0; i < n; i++ {
        jobs[i] = Job{start: start[i], end: end[i], profit: profit[i]}
    }
    sort.Slice(jobs, func(i, j int) bool {
        return jobs[i].start < jobs[j].start
    })
    jump := make([]int, n)
    for i := 0; i < n; i++ {
        jump[i] = sort.Search(n, func(j int) bool {
            return jobs[j].start >= jobs[i].end
        })
    }
    dp := make([]int, n+1)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := n - 1; i >= 0; i-- {
        dp[i] = max(dp[i+1], jobs[i].profit + dp[jump[i]])
    }
    return dp[0]
}

func main() {
    // Explanation: The subset chosen is the first and fourth job. 
    // Time range [1-3]+[3-6] , we get profit of 120 = 50 + 70.
    fmt.Println(jobScheduling([]int{1,2,3,3},[]int{3,4,5,6},[]int{50,10,40,70})) // 120

    // Explanation: The subset chosen is the first, fourth and fifth job. 
    // Profit obtained 150 = 20 + 70 + 60.
    fmt.Println(jobScheduling([]int{1,2,3,4,6},[]int{3,5,10,6,9},[]int{20,20,100,70,60})) // 150

    fmt.Println(jobScheduling([]int{1,1,1},[]int{2,3,4},[]int{5,6,4})) // 6

    fmt.Println(jobScheduling1([]int{1,2,3,3},[]int{3,4,5,6},[]int{50,10,40,70})) // 120
    fmt.Println(jobScheduling1([]int{1,2,3,4,6},[]int{3,5,10,6,9},[]int{20,20,100,70,60})) // 150
    fmt.Println(jobScheduling1([]int{1,1,1},[]int{2,3,4},[]int{5,6,4})) // 6

    fmt.Println(jobScheduling2([]int{1,2,3,3},[]int{3,4,5,6},[]int{50,10,40,70})) // 120
    fmt.Println(jobScheduling2([]int{1,2,3,4,6},[]int{3,5,10,6,9},[]int{20,20,100,70,60})) // 150
    fmt.Println(jobScheduling2([]int{1,1,1},[]int{2,3,4},[]int{5,6,4})) // 6
}