package main

// 3709. Design Exam Scores Tracker
// Alice frequently takes exams and wants to track her scores and calculate the total scores over specific time periods.

// Implement the ExamTracker class:
//     ExamTracker(): Initializes the ExamTracker object.
//     void record(int time, int score): Alice takes a new exam at time time and achieves the score score.
//     long long totalScore(int startTime, int endTime): Returns an integer that represents the total score of all exams taken by Alice between startTime and endTime (inclusive). If there are no recorded exams taken by Alice within the specified time interval, return 0.
//     It is guaranteed that the function calls are made in chronological order. That is,

// Calls to record() will be made with strictly increasing time.
// Alice will never ask for total scores that require information from the future. That is, if the latest record() is called with time = t, then totalScore() will always be called with startTime <= endTime <= t.
 
// Example 1:
// Input:
// ["ExamTracker", "record", "totalScore", "record", "totalScore", "totalScore", "totalScore", "totalScore"]
// [[], [1, 98], [1, 1], [5, 99], [1, 3], [1, 5], [3, 4], [2, 5]]

// Output:
// [null, null, 98, null, 98, 197, 0, 99]
// Explanation
// ExamTracker examTracker = new ExamTracker();
// examTracker.record(1, 98); // Alice takes a new exam at time 1, scoring 98.
// examTracker.totalScore(1, 1); // Between time 1 and time 1, Alice took 1 exam at time 1, scoring 98. The total score is 98.
// examTracker.record(5, 99); // Alice takes a new exam at time 5, scoring 99.
// examTracker.totalScore(1, 3); // Between time 1 and time 3, Alice took 1 exam at time 1, scoring 98. The total score is 98.
// examTracker.totalScore(1, 5); // Between time 1 and time 5, Alice took 2 exams at time 1 and 5, scoring 98 and 99. The total score is 98 + 99 = 197.
// examTracker.totalScore(3, 4); // Alice did not take any exam between time 3 and time 4. Therefore, the answer is 0.
// examTracker.totalScore(2, 5); // Between time 2 and time 5, Alice took 1 exam at time 5, scoring 99. The total score is 99.
 
// Constraints:
//     1 <= time <= 10^9
//     1 <= score <= 10^9
//     1 <= startTime <= endTime <= t, where t is the value of time from the most recent call of record().
//     Calls of record() will be made with strictly increasing time.
//     After ExamTracker(), the first function call will always be record().
//     At most 105 calls will be made in total to record() and totalScore().

import "fmt"
import "sort"

type ExamTracker struct {
    times  []int
    prefixSum []int64
}

func Constructor() ExamTracker {
    return ExamTracker{[]int{}, []int64{0}}
}

func (e *ExamTracker) Record(time, score int) {
    e.times = append(e.times, time)
	e.prefixSum = append(e.prefixSum, e.prefixSum[len(e.prefixSum)-1]+int64(score))
}

func (e *ExamTracker) TotalScore(startTime, endTime int) int64 {
    left := sort.SearchInts(e.times, startTime)
    right := sort.SearchInts(e.times, endTime+1) // 也可以在 e.times[left:] 中二分
    return e.prefixSum[right] - e.prefixSum[left]
}

/**
 * Your ExamTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Record(time,score);
 * param_2 := obj.TotalScore(startTime,endTime);
 */

func main() {
    // ExamTracker examTracker = new ExamTracker();
    examTracker := Constructor()
    // examTracker.record(1, 98); // Alice takes a new exam at time 1, scoring 98.
    examTracker.Record(1, 98)
    // examTracker.totalScore(1, 1); // Between time 1 and time 1, Alice took 1 exam at time 1, scoring 98. The total score is 98.
    fmt.Println(examTracker.TotalScore(1, 1)) 
    // examTracker.record(5, 99); // Alice takes a new exam at time 5, scoring 99.
    examTracker.Record(5, 99)
    // examTracker.totalScore(1, 3); // Between time 1 and time 3, Alice took 1 exam at time 1, scoring 98. The total score is 98.
    fmt.Println(examTracker.TotalScore(1, 3))
    // examTracker.totalScore(1, 5); // Between time 1 and time 5, Alice took 2 exams at time 1 and 5, scoring 98 and 99. The total score is 98 + 99 = 197. 
    fmt.Println(examTracker.TotalScore(1, 5))// examTracker.totalScore(3, 4); // Alice did not take any exam between time 3 and time 4. Therefore, the answer is 0.
    // examTracker.totalScore(3, 4); // Alice did not take any exam between time 3 and time 4. Therefore, the answer is 0.
    fmt.Println(examTracker.TotalScore(3, 4))
    // examTracker.totalScore(2, 5); // Between time 2 and time 5, Alice took 1 exam at time 5, scoring 99. The total score is 99.
    fmt.Println(examTracker.TotalScore(2, 5))
}