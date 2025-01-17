package main

// 2512. Reward Top K Students
// You are given two string arrays positive_feedback and negative_feedback, 
// containing the words denoting positive and negative feedback, respectively. 
// Note that no word is both positive and negative.

// Initially every student has 0 points. 
// Each positive word in a feedback report increases the points of a student by 3, 
// whereas each negative word decreases the points by 1.

// You are given n feedback reports, 
// represented by a 0-indexed string array report and a 0-indexed integer array student_id, 
// where student_id[i] represents the ID of the student who has received the feedback report report[i]. 
// The ID of each student is unique.

// Given an integer k, return the top k students after ranking them in non-increasing order by their points. 
// In case more than one student has the same points, the one with the lower ID ranks higher.

// Example 1:
// Input: positive_feedback = ["smart","brilliant","studious"], negative_feedback = ["not"], report = ["this student is studious","the student is smart"], student_id = [1,2], k = 2
// Output: [1,2]
// Explanation: 
// Both the students have 1 positive feedback and 3 points but since student 1 has a lower ID he ranks higher.

// Example 2:
// Input: positive_feedback = ["smart","brilliant","studious"], negative_feedback = ["not"], report = ["this student is not studious","the student is smart"], student_id = [1,2], k = 2
// Output: [2,1]
// Explanation: 
// - The student with ID 1 has 1 positive feedback and 1 negative feedback, so he has 3-1=2 points. 
// - The student with ID 2 has 1 positive feedback, so he has 3 points. 
// Since student 2 has more points, [2,1] is returned.

// Constraints:
//     1 <= positive_feedback.length, negative_feedback.length <= 10^4
//     1 <= positive_feedback[i].length, negative_feedback[j].length <= 100
//     Both positive_feedback[i] and negative_feedback[j] consists of lowercase English letters.
//     No word is present in both positive_feedback and negative_feedback.
//     n == report.length == student_id.length
//     1 <= n <= 10^4
//     report[i] consists of lowercase English letters and spaces ' '.
//     There is a single space between consecutive words of report[i].
//     1 <= report[i].length <= 100
//     1 <= student_id[i] <= 10^9
//     All the values of student_id[i] are unique.
//     1 <= k <= n

import "fmt"
import "strings"
import "sort"

func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {
    if k > len(student_id) { return student_id }
    positiveSet, negativeSet := make(map[string]bool),  make(map[string]bool)
    for _, v := range positive_feedback {
        positiveSet[v] = true
    }
    for _, v := range negative_feedback {
        negativeSet[v] = true
    }
    arr := make([][]int, len(report)) // first dimension is report/student idx, second is a slice, where 0th element is id of student and 1st element is score 
    for i := range arr {
        arr[i] = make([]int, 2)
    }
    for i, v := range report {
        score := 0
        for _, v := range strings.Split(v, " ") {
            if positiveSet[v] { score += 3 }
            if negativeSet[v] { score-- }
        }
        arr[i][0],  arr[i][1] = student_id[i], score
    }
    sort.Slice(arr, func(i, j int) bool { // sort two dimensional slice
        if arr[i][1] == arr[j][1] { return arr[i][0] < arr[j][0]  } // based on index
        return arr[i][1] > arr[j][1] // based on score
    })
    res := make([]int, 0, k)
    for i := 0; i < k; i++ { // append k student_ids
        res = append(res, arr[i][0])
    }
    return res
}

func topStudents1(positiveFeedback, negativeFeedback, report []string, studentId []int, k int) []int {
    words := make(map[string]int)
    for _, w := range positiveFeedback { words[w] = 3  }
    for _, w := range negativeFeedback { words[w] = -1 }
    type Student struct { score, id int }
    arr := make([]Student, len(report))
    for i, v := range report {
        score := 0
        for _, w := range strings.Split(v, " ") {
            score += words[w]
        }
        arr[i] = Student{score, studentId[i] }
    }
    sort.Slice(arr, func(i, j int) bool {
        a, b := arr[i], arr[j]
        return a.score > b.score || a.score == b.score && a.id < b.id
    })
    res := make([]int, k)
    for i, v := range arr[:k] {
        res[i] = v.id
    }
    return res
}

func main() {
    // Example 1:
    // Input: positive_feedback = ["smart","brilliant","studious"], negative_feedback = ["not"], report = ["this student is studious","the student is smart"], student_id = [1,2], k = 2
    // Output: [1,2]
    // Explanation: 
    // Both the students have 1 positive feedback and 3 points but since student 1 has a lower ID he ranks higher.
    fmt.Println(topStudents([]string{"smart","brilliant","studious"}, []string{"not"}, []string{"this student is studious","the student is smart"}, []int{1,2}, 2)) // [1,2]
    // Example 2:
    // Input: positive_feedback = ["smart","brilliant","studious"], negative_feedback = ["not"], report = ["this student is not studious","the student is smart"], student_id = [1,2], k = 2
    // Output: [2,1]
    // Explanation: 
    // - The student with ID 1 has 1 positive feedback and 1 negative feedback, so he has 3-1=2 points. 
    // - The student with ID 2 has 1 positive feedback, so he has 3 points. 
    // Since student 2 has more points, [2,1] is returned.
    fmt.Println(topStudents([]string{"smart","brilliant","studious"}, []string{"not"}, []string{"this student is not studious","the student is smart"}, []int{1,2}, 2)) // [2,1]

    fmt.Println(topStudents1([]string{"smart","brilliant","studious"}, []string{"not"}, []string{"this student is studious","the student is smart"}, []int{1,2}, 2)) // [1,2]
    fmt.Println(topStudents1([]string{"smart","brilliant","studious"}, []string{"not"}, []string{"this student is not studious","the student is smart"}, []int{1,2}, 2)) // [2,1]
}