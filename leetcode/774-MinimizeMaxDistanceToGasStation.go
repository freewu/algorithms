package main

// 774. Minimize Max Distance to Gas Station
// You are given an integer array stations that represents the positions of the gas stations on the x-axis. 
// You are also given an integer k.

// You should add k new gas stations. 
// You can add the stations anywhere on the x-axis, and not necessarily on an integer position.

// Let penalty() be the maximum distance between adjacent gas stations after adding the k new stations.
// Return the smallest possible value of penalty(). Answers within 10-6 of the actual answer will be accepted.

// Example 1:
// Input: stations = [1,2,3,4,5,6,7,8,9,10], k = 9
// Output: 0.50000

// Example 2:
// Input: stations = [23,24,36,39,46,56,57,65,84,98], k = 1
// Output: 14.00000

// Constraints:
//     10 <= stations.length <= 2000
//     0 <= stations[i] <= 10^8
//     stations is sorted in a strictly increasing order.
//     1 <= k <= 10^6

import "fmt"

// 二分
func minmaxGasDist(stations []int, k int) float64 {
    maxDistance := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < len(stations); i++ {
        maxDistance = max(maxDistance, stations[i]-stations[i-1])
    }
    left, right := 0.0, float64(maxDistance)
    for (right - left) > 1e-6 {
        mid := (left + right) / 2
        cnt := 0
        for i := 1; i < len(stations); i++ {
            cnt += int(float64(stations[i] - stations[i-1]) / mid)
        }
        if cnt <= k { // 要找到 mid 的最小值，满足 penalty(stations, mid) <= float64(k)
            right = mid
        } else {
            left = mid
        }
    }
    return left
}

// 性质P：相邻两个加油站见的最大距离为target时 
// 新增设的加油站数量 <= k
func minmaxGasDist1(stations []int, k int) float64 {
    check := func(target float64 )bool{
        count := 0
        for i := 1;i < len(stations);i++{
            // 使用除法避免减法的精度丢失
           	count += int(float64(stations[i]-stations[i-1]) / target)
        }
        return count <= k
    }
    maxDistance := 0
	for i := 1; i < len(stations); i++ {
		maxDistance = max(maxDistance, stations[i]-stations[i-1])
	}
    // 闭区间
    // 循环不变量
    //     L-1e-6 < target
    //     R+1e-6 >= target 
    left,right := 1e-6, float64(maxDistance) - 1e-6
    // 注意这里区间为空的判断
    // 不再是 left <= right
    // 而是根据精度要求为 right-left > 1e-6
    for right-left >= (1e-6) {
        mid := left + (right-left)/2.0
        if check(mid){
            right = mid-1e-6
        }else{
            left = mid+1e-6
        }
    }
    return right + 1e-6 // 或 left
}

func main() {
    // Example 1:
    // Input: stations = [1,2,3,4,5,6,7,8,9,10], k = 9
    // Output: 0.50000
    fmt.Println(minmaxGasDist([]int{1,2,3,4,5,6,7,8,9,10}, 9)) // 0.50000
    // Example 2:
    // Input: stations = [23,24,36,39,46,56,57,65,84,98], k = 1
    // Output: 14.00000
    fmt.Println(minmaxGasDist([]int{23,24,36,39,46,56,57,65,84,98}, 1)) // 14.00000

    fmt.Println(minmaxGasDist([]int{1,2,3,4,5,6,7,8,9}, 1)) // 0.9999990463256836
    fmt.Println(minmaxGasDist([]int{9,8,7,6,5,4,3,2,1}, 1)) // 0

    fmt.Println(minmaxGasDist1([]int{1,2,3,4,5,6,7,8,9,10}, 9)) // 0.50000
    fmt.Println(minmaxGasDist1([]int{23,24,36,39,46,56,57,65,84,98}, 1)) // 14.00000
    fmt.Println(minmaxGasDist1([]int{1,2,3,4,5,6,7,8,9}, 1)) // 1
    fmt.Println(minmaxGasDist1([]int{9,8,7,6,5,4,3,2,1}, 1)) // 0
}