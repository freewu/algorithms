package main

// 911. Online Election
// You are given two integer arrays persons and times. 
// In an election, the ith vote was cast for persons[i] at time times[i].

// For each query at a time t, find the person that was leading the election at time t. 
// Votes cast at time t will count towards our query. 
// In the case of a tie, the most recent vote (among tied candidates) wins.

// Implement the TopVotedCandidate class:
//     TopVotedCandidate(int[] persons, int[] times) 
//         Initializes the object with the persons and times arrays.
//     int q(int t) 
//         Returns the number of the person that was leading the election at time t according to the mentioned rules.

// Example 1:
// Input
// ["TopVotedCandidate", "q", "q", "q", "q", "q", "q"]
// [[[0, 1, 1, 0, 0, 1, 0], [0, 5, 10, 15, 20, 25, 30]], [3], [12], [25], [15], [24], [8]]
// Output
// [null, 0, 1, 1, 0, 0, 1]
// Explanation
// TopVotedCandidate topVotedCandidate = new TopVotedCandidate([0, 1, 1, 0, 0, 1, 0], [0, 5, 10, 15, 20, 25, 30]);
// topVotedCandidate.q(3); // return 0, At time 3, the votes are [0], and 0 is leading.
// topVotedCandidate.q(12); // return 1, At time 12, the votes are [0,1,1], and 1 is leading.
// topVotedCandidate.q(25); // return 1, At time 25, the votes are [0,1,1,0,0,1], and 1 is leading (as ties go to the most recent vote.)
// topVotedCandidate.q(15); // return 0
// topVotedCandidate.q(24); // return 0
// topVotedCandidate.q(8); // return 1

// Constraints:
//     1 <= persons.length <= 5000
//     times.length == persons.length
//     0 <= persons[i] < persons.length
//     0 <= times[i] <= 10^9
//     times is sorted in a strictly increasing order.
//     times[0] <= t <= 10^9
//     At most 10^4 calls will be made to q.

import "fmt"
import "sort"

type TopVotedCandidate struct {
    list [][2]int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
    list, dict, winner := [][2]int{}, make(map[int]int), -1
    for i, v := range persons {
        if winner == -1 || v == winner {
            winner = v
            dict[winner]++
        } else {
            dict[v]++
            if dict[v] >= dict[winner] {
                winner = v
            }
        }
        list = append(list, [2]int{times[i], winner})
    }
    return TopVotedCandidate{list: list}
}

func (this *TopVotedCandidate) Q(t int) int {
    index := sort.Search(len(this.list), func(i int) bool {
        return (this.list)[i][0] > t
    })
    return (this.list)[index - 1][1]
}


type TopVotedCandidate1 struct {
    arr, times []int
}

func Constructor1(persons []int, times []int) TopVotedCandidate1 {
    arr, cnt, cur := make([]int, len(times)), make([]int, len(times)), -1
    for i := range times {
        cnt[persons[i]]++
        if cur < 0 || cnt[persons[i]] >= cnt[cur] {
            cur = persons[i]
        }
        arr[i] = cur
    }
    return TopVotedCandidate1{arr, times}
}

func (this *TopVotedCandidate1) Q(t int) int {
    i := sort.Search(len(this.times), func(i int) bool {
        return this.times[i] > t
    })
    return this.arr[i-1]
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */

func main() {
    // TopVotedCandidate topVotedCandidate = new TopVotedCandidate([0, 1, 1, 0, 0, 1, 0], [0, 5, 10, 15, 20, 25, 30]);
    obj := Constructor([]int{0, 1, 1, 0, 0, 1, 0},[]int{0, 5, 10, 15, 20, 25, 30})
    fmt.Println(obj)
    // topVotedCandidate.q(3); // return 0, At time 3, the votes are [0], and 0 is leading.
    fmt.Println(obj.Q(3)) // 0
    // topVotedCandidate.q(12); // return 1, At time 12, the votes are [0,1,1], and 1 is leading.
    fmt.Println(obj.Q(12)) // 1
    // topVotedCandidate.q(25); // return 1, At time 25, the votes are [0,1,1,0,0,1], and 1 is leading (as ties go to the most recent vote.)
    fmt.Println(obj.Q(25)) // 1
    // topVotedCandidate.q(15); // return 0
    fmt.Println(obj.Q(15)) // 0
    // topVotedCandidate.q(24); // return 0
    fmt.Println(obj.Q(24)) // 0
    // topVotedCandidate.q(8); // return 1
    fmt.Println(obj.Q(8)) // 1


    obj1 := Constructor([]int{0, 1, 1, 0, 0, 1, 0},[]int{0, 5, 10, 15, 20, 25, 30})
    fmt.Println(obj1)
    // topVotedCandidate.q(3); // return 0, At time 3, the votes are [0], and 0 is leading.
    fmt.Println(obj1.Q(3)) // 0
    // topVotedCandidate.q(12); // return 1, At time 12, the votes are [0,1,1], and 1 is leading.
    fmt.Println(obj1.Q(12)) // 1
    // topVotedCandidate.q(25); // return 1, At time 25, the votes are [0,1,1,0,0,1], and 1 is leading (as ties go to the most recent vote.)
    fmt.Println(obj1.Q(25)) // 1
    // topVotedCandidate.q(15); // return 0
    fmt.Println(obj1.Q(15)) // 0
    // topVotedCandidate.q(24); // return 0
    fmt.Println(obj1.Q(24)) // 0
    // topVotedCandidate.q(8); // return 1
    fmt.Println(obj1.Q(8)) // 1
}