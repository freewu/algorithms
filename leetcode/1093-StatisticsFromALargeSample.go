package main

// 1093. Statistics from a Large Sample
// You are given a large sample of integers in the range [0, 255]. 
// Since the sample is so large, it is represented by an array count where count[k] is the number of times that k appears in the sample.

// Calculate the following statistics:
//     minimum: The minimum element in the sample.
//     maximum: The maximum element in the sample.
//     mean: The average of the sample, calculated as the total sum of all elements divided by the total number of elements.
//     median:
//         If the sample has an odd number of elements, then the median is the middle element once the sample is sorted.
//         If the sample has an even number of elements, then the median is the average of the two middle elements once the sample is sorted.
//     mode: The number that appears the most in the sample. It is guaranteed to be unique.

// Return the statistics of the sample as an array of floating-point numbers [minimum, maximum, mean, median, mode]. 
// Answers within 10^-5 of the actual answer will be accepted.

// Example 1:
// Input: count = [0,1,3,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
// Output: [1.00000,3.00000,2.37500,2.50000,3.00000]
// Explanation: The sample represented by count is [1,2,2,2,3,3,3,3].
// The minimum and maximum are 1 and 3 respectively.
// The mean is (1+2+2+2+3+3+3+3) / 8 = 19 / 8 = 2.375.
// Since the size of the sample is even, the median is the average of the two middle elements 2 and 3, which is 2.5.
// The mode is 3 as it appears the most in the sample.

// Example 2:
// Input: count = [0,4,3,2,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
// Output: [1.00000,4.00000,2.18182,2.00000,1.00000]
// Explanation: The sample represented by count is [1,1,1,1,2,2,2,3,3,4,4].
// The minimum and maximum are 1 and 4 respectively.
// The mean is (1+1+1+1+2+2+2+3+3+4+4) / 11 = 24 / 11 = 2.18181818... (for display purposes, the output shows the rounded number 2.18182).
// Since the size of the sample is odd, the median is the middle element 2.
// The mode is 1 as it appears the most in the sample.

// Constraints:
//     count.length == 256
//     0 <= count[i] <= 10^9
//     1 <= sum(count) <= 10^9
//     The mode of the sample that count represents is unique.

import "fmt"

// 双指针
func sampleStats(count []int) []float64 {
    mx, mn  := 0, 265 // 最大值, 最小值
    counts, total := 0, 0 // 样本个数, 样本总值
    median, mode := 0.0, 0 // 中位数, 保众数
    i, j := 0, len(count) - 1
    for i < j {
        for count[i] == 0 { i++ } // 跳过所有空值
        for count[j] == 0 { j-- }// 跳过所有空值
        if i >= j { // 上一步可能 i >= j
            break
        }
        if mn == 265 { // 有序数组，左指针碰到的第一个非0值就是最小值
            mn = i
        }
        if mx == 0 { // 有序数组，右指针碰到的第一个非0值就是最大值
            mx = j
        }
        // 计算总值和个数
        counts += count[j]
        total += count[j] * j
        counts += count[i]
        total += count[i] * i
        // 取众数
        if count[i] > count[mode] {
            mode = i
        }
        if count[j] > count[mode] {
            mode = j
        }

        // 移动指针
        i++
        j--
    }

    // 补充 i == j 的值
    if i == j {
        counts += count[i]
        total += count[i] * i
    }
    // 取中位数
    // 中间个数
    mid := float64(counts) / 2
    for k := 0; k < len(count); k++ {
        if count[k] == 0 {
            continue
        }
        // 说明有偶数个
        if mid-float64(count[k]) == 0 {
            median = float64(k)
        }
        if mid-float64(count[k]) < 0 {
            if median > 0 {
                // 补充偶数个时的另一个值
                median += float64(k)
            } else {
                // 奇数个
                median = 2 * float64(k)
            }
            break
        }
        mid -= float64(count[k])
    }
    return []float64{float64(mn), float64(mx), float64(total) / float64(counts) , median / 2, float64(mode)}
}

func sampleStats1(count []int) []float64 {
    samCount, samNum, mostVal, mostNum, c := 0, 0, 0, 0, 0
    vCount := make([][2]int, 0, len(count))
    mx, mn := 0, -1
    for i, v := range count {
        if v == 0 { continue }
        if mn == -1 {
            mn = i
        }
        mx = i
        samNum += v
        samCount += v * i
        if v > mostNum {
            mostVal = i
            mostNum = v
        }
        vCount = append(vCount, [2]int{i, v})
    }
    odd := (samNum % 2) == 0
    mid := samNum / 2
    var median float64
    for i, v := range vCount {
        if c + v[1] > mid {
            median = float64(v[0])
            break
        } else if c+v[1] == mid {
            if !odd {
                median = float64(vCount[i+1][0])
                break
            } else {
                median = float64(vCount[i+1][0]+v[0]) / 2
                break
            }
        }
        c += v[1]
    }
    return []float64{float64(mn), float64(mx), float64(samCount) / float64(samNum), median, float64(mostVal)}
}

func main() {
    // Example 1:
    // Input: count = [0,1,3,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
    // Output: [1.00000,3.00000,2.37500,2.50000,3.00000]
    // Explanation: The sample represented by count is [1,2,2,2,3,3,3,3].
    // The minimum and maximum are 1 and 3 respectively.
    // The mean is (1+2+2+2+3+3+3+3) / 8 = 19 / 8 = 2.375.
    // Since the size of the sample is even, the median is the average of the two middle elements 2 and 3, which is 2.5.
    // The mode is 3 as it appears the most in the sample.
    sample1 := []int{0,1,3,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
    fmt.Println(sampleStats(sample1)) // [1.00000,3.00000,2.37500,2.50000,3.00000]
    // Example 2:
    // Input: count = [0,4,3,2,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
    // Output: [1.00000,4.00000,2.18182,2.00000,1.00000]
    // Explanation: The sample represented by count is [1,1,1,1,2,2,2,3,3,4,4].
    // The minimum and maximum are 1 and 4 respectively.
    // The mean is (1+1+1+1+2+2+2+3+3+4+4) / 11 = 24 / 11 = 2.18181818... (for display purposes, the output shows the rounded number 2.18182).
    // Since the size of the sample is odd, the median is the middle element 2.
    // The mode is 1 as it appears the most in the sample.
    sample2 := []int{0,4,3,2,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
    fmt.Println(sampleStats(sample2)) // [1.00000,4.00000,2.18182,2.00000,1.00000]

    fmt.Println(sampleStats1(sample1)) // [1.00000,3.00000,2.37500,2.50000,3.00000]
    fmt.Println(sampleStats1(sample2)) // [1.00000,4.00000,2.18182,2.00000,1.00000]
}