package main

// 1024. Video Stitching
// You are given a series of video clips from a sporting event that lasted time seconds. 
// These video clips can be overlapping with each other and have varying lengths.

// Each video clip is described by an array clips where clips[i] = [starti, endi] indicates 
// that the ith clip started at starti and ended at endi.

// We can cut these clips into segments freely.
//     For example, a clip [0, 7] can be cut into segments [0, 1] + [1, 3] + [3, 7].

// Return the minimum number of clips needed so that we can cut the clips into segments 
// that cover the entire sporting event [0, time]. 
// If the task is impossible, return -1.

// Example 1:
// Input: clips = [[0,2],[4,6],[8,10],[1,9],[1,5],[5,9]], time = 10
// Output: 3
// Explanation: We take the clips [0,2], [8,10], [1,9]; a total of 3 clips.
// Then, we can reconstruct the sporting event as follows:
// We cut [1,9] into segments [1,2] + [2,8] + [8,9].
// Now we have segments [0,2] + [2,8] + [8,10] which cover the sporting event [0, 10].

// Example 2:
// Input: clips = [[0,1],[1,2]], time = 5
// Output: -1
// Explanation: We cannot cover [0,5] with only [0,1] and [1,2].

// Example 3:
// Input: clips = [[0,1],[6,8],[0,2],[5,6],[0,4],[0,3],[6,7],[1,3],[4,7],[1,4],[2,5],[2,6],[3,4],[4,5],[5,7],[6,9]], time = 9
// Output: 3
// Explanation: We can take clips [0,4], [4,7], and [6,9].

// Constraints:
//     1 <= clips.length <= 100
//     0 <= starti <= endi <= 100
//     1 <= time <= 100

import "fmt"
import "sort"

func videoStitching(clips [][]int, time int) int {
    sort.Slice(clips, func(i, j int) bool {
        if clips[i][0] == clips[j][0] { return clips[j][1] < clips[i][1] }
        return clips[i][0] < clips[j][0]
    })
    if clips[0][0] > 0 { return -1 }
    cur, end := 0, clips[0][1]
    for i := range clips {
        cur++
        if end >= time {
            return cur
        }
        nextEnd := end
        for j := i + 1; j < len(clips); j++ {
            if clips[j][0] <= end && clips[j][1] > nextEnd {
                nextEnd = clips[j][1]
                i = j
            }
        }
        if nextEnd == end {
            return -1
        }
        i--
        end = nextEnd
    }
    return -1
}

func videoStitching1(clips [][]int, time int) int {
    max := func (x, y int) int { if x > y { return x; }; return y; }
    mx := make([]int, time) // 转换为跳跃游戏的数组形式
    for _, clip := range clips {
        l, r := clip[0], clip[1]
        if l < time  {
            mx[l] = max(mx[l], r-l)
        }
    }
    maxPos, jumpPos, step := 0, 0, 0
    for i := 0; i < time; i++ {
        num := mx[i]
        if i > jumpPos {
            return -1
        }
        if i + num > maxPos {
            maxPos = i + num
        }
        if i == jumpPos {
            step++
            jumpPos = maxPos
        }
    }
    if jumpPos >= time {
        return step
    }
    return -1
}

func main() {
    // Example 1:
    // Input: clips = [[0,2],[4,6],[8,10],[1,9],[1,5],[5,9]], time = 10
    // Output: 3
    // Explanation: We take the clips [0,2], [8,10], [1,9]; a total of 3 clips.
    // Then, we can reconstruct the sporting event as follows:
    // We cut [1,9] into segments [1,2] + [2,8] + [8,9].
    // Now we have segments [0,2] + [2,8] + [8,10] which cover the sporting event [0, 10].
    fmt.Println(videoStitching([][]int{{0,2},{4,6},{8,10},{1,9},{1,5},{5,9}}, 10)) // 3
    // Example 2:
    // Input: clips = [[0,1],[1,2]], time = 5
    // Output: -1
    // Explanation: We cannot cover [0,5] with only [0,1] and [1,2].
    fmt.Println(videoStitching([][]int{{0,1},{1,2}}, 5)) // -1
    // Example 3:
    // Input: clips = [[0,1],[6,8],[0,2],[5,6],[0,4],[0,3],[6,7],[1,3],[4,7],[1,4],[2,5],[2,6],[3,4],[4,5],[5,7],[6,9]], time = 9
    // Output: 3
    // Explanation: We can take clips [0,4], [4,7], and [6,9].
    fmt.Println(videoStitching([][]int{{0,1},{6,8},{0,2},{5,6},{0,4},{0,3},{6,7},{1,3},{4,7},{1,4},{2,5},{2,6},{3,4},{4,5},{5,7},{6,9}}, 9)) // 3

    fmt.Println(videoStitching1([][]int{{0,2},{4,6},{8,10},{1,9},{1,5},{5,9}}, 10)) // 3
    fmt.Println(videoStitching1([][]int{{0,1},{1,2}}, 5)) // -1
    fmt.Println(videoStitching1([][]int{{0,1},{6,8},{0,2},{5,6},{0,4},{0,3},{6,7},{1,3},{4,7},{1,4},{2,5},{2,6},{3,4},{4,5},{5,7},{6,9}}, 9)) // 3
}