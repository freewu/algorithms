package main

// 546. Remove Boxes
// You are given several boxes with different colors represented by different positive numbers.

// You may experience several rounds to remove boxes until there is no box left. 
// Each time you can choose some continuous boxes with the same color (i.e., composed of k boxes, k >= 1), 
// remove them and get k * k points.

// Return the maximum points you can get.

// Example 1:
// Input: boxes = [1,3,2,2,2,3,4,3,1]
// Output: 23
// Explanation:
// [1, 3, 2, 2, 2, 3, 4, 3, 1] 
// ----> [1, 3, 3, 4, 3, 1] (3*3=9 points) 
// ----> [1, 3, 3, 3, 1] (1*1=1 points) 
// ----> [1, 1] (3*3=9 points) 
// ----> [] (2*2=4 points)

// Example 2:
// Input: boxes = [1,1,1]
// Output: 9

// Example 3:
// Input: boxes = [1]
// Output: 1
 
// Constraints:
//     1 <= boxes.length <= 100
//     1 <= boxes[i] <= 100

import "fmt"

// 一维数组 dp
func removeBoxes(boxes []int) int {
    n := len(boxes)
    dp := make([]int, n * n * n)
    var dfs func(boxes []int, i, j, k, n int, dp []int) int
    dfs = func (boxes []int, i, j, k, n int, dp []int) int {
        if i > j {
            return 0
        }
        index := k * n * n + j * n + i
        if r := dp[index]; r > 0 {
            return r
        }
        mx := (k+1) * (k+1) + dfs(boxes, i + 1, j, 0, n, dp)
        for l := i + 1; l <= j; l++ {
            if boxes[i] == boxes[l] {
                x := dfs(boxes, i + 1, l - 1, 0, n, dp) + dfs(boxes, l, j, k + 1, n, dp)
                if x > mx {
                    mx = x
                }
            }
        }
        dp[index] = mx
        return mx
    }
    return dfs(boxes, 0, n - 1, 0, n, dp)
}

// 三维数组 dp
func removeBoxes1(boxes []int) int {
    dp, n:= [100][100][100]int{}, len(boxes)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(boxes []int, l int, r int, k int) int
    dfs = func(boxes []int, l int, r int, k int) int {
        if l > r {
            return 0
        }
        if dp[l][r][k] != 0 {
            return dp[l][r][k]
        }
        s := l
        for s <= r && boxes[s] == boxes[l] {
            s++
        }
        s--
        cnt := k + s - l + 1
        res := cnt * cnt + dfs(boxes, s + 1, r, 0)
        //  2 2 2    2 2 2 3 3 2 2 4 4 2 2 5 5
        //  k =3     l   s     m1      m2    r
        for i := s + 2; i <= r; i++ {
            if boxes[i] == boxes[l] {
                res = max(res, dfs(boxes, s + 1, i - 1, 0) + dfs(boxes, i, r, cnt))
            }
        }
        dp[l][r][k] = res
        return res
    }
    return dfs(boxes, 0, n - 1, 0)
}

func main() {
    // Example 1:
    // Input: boxes = [1,3,2,2,2,3,4,3,1]
    // Output: 23
    // Explanation:
    // [1, 3, 2, 2, 2, 3, 4, 3, 1] 
    // ----> [1, 3, 3, 4, 3, 1] (3*3=9 points) 
    // ----> [1, 3, 3, 3, 1] (1*1=1 points) 
    // ----> [1, 1] (3*3=9 points) 
    // ----> [] (2*2=4 points)
    fmt.Println(removeBoxes([]int{1, 3, 2, 2, 2, 3, 4, 3, 1})) // 23
    // Example 2:
    // Input: boxes = [1,1,1]
    // Output: 9
    fmt.Println(removeBoxes([]int{1,1,1})) // 9
    // Example 3:
    // Input: boxes = [1]
    // Output: 1
    fmt.Println(removeBoxes([]int{1})) // 1

    fmt.Println(removeBoxes1([]int{1, 3, 2, 2, 2, 3, 4, 3, 1})) // 23
    fmt.Println(removeBoxes1([]int{1,1,1})) // 9
    fmt.Println(removeBoxes1([]int{1})) // 1
}