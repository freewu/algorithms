package main

// 458. Poor Pigs
// There are buckets buckets of liquid, where exactly one of the buckets is poisonous. 
// To figure out which one is poisonous, you feed some number of (poor) pigs the liquid to see whether they will die or not.
// Unfortunately, you only have minutesToTest minutes to determine which bucket is poisonous.

// You can feed the pigs according to these steps:
//     1 Choose some live pigs to feed.
//     2 For each pig, choose which buckets to feed it. 
//       The pig will consume all the chosen buckets simultaneously and will take no time. 
//       Each pig can feed from any number of buckets, and each bucket can be fed from by any number of pigs.
//     3 Wait for minutesToDie minutes. You may not feed any other pigs during this time.
//     4 After minutesToDie minutes have passed, any pigs that have been fed the poisonous bucket will die, and all others will survive.
//     5 Repeat this process until you run out of time.

// Given buckets, minutesToDie, and minutesToTest, return the minimum number of pigs needed to figure out which bucket is poisonous within the allotted time.

// Example 1:
// Input: buckets = 4, minutesToDie = 15, minutesToTest = 15
// Output: 2
// Explanation: We can determine the poisonous bucket as follows:
// At time 0, feed the first pig buckets 1 and 2, and feed the second pig buckets 2 and 3.
// At time 15, there are 4 possible outcomes:
// - If only the first pig dies, then bucket 1 must be poisonous.
// - If only the second pig dies, then bucket 3 must be poisonous.
// - If both pigs die, then bucket 2 must be poisonous.
// - If neither pig dies, then bucket 4 must be poisonous.

// Example 2:
// Input: buckets = 4, minutesToDie = 15, minutesToTest = 30
// Output: 2
// Explanation: We can determine the poisonous bucket as follows:
// At time 0, feed the first pig bucket 1, and feed the second pig bucket 2.
// At time 15, there are 2 possible outcomes:
// - If either pig dies, then the poisonous bucket is the one it was fed.
// - If neither pig dies, then feed the first pig bucket 3, and feed the second pig bucket 4.
// At time 30, one of the two pigs must die, and the poisonous bucket is the one it was fed.

// Constraints:
//     1 <= buckets <= 1000
//     1 <= minutesToDie <= minutesToTest <= 100

import "fmt"
import "math"

// 25 / 26 个通过的测试用例
// 使用数学方法,以 minutesToDie=15, minutesToTest=60, 1 只小猪为例,可以测试 5 只桶
//     0-15 小猪吃第一个桶中的液体,如果死去，则第一个桶有毒,否则继续测试
//     15-30 小猪吃第二个桶中的液体，如果死去，则第二个桶有毒,否则继续测试
//     30-45 小猪吃第三个桶中的液体，如果死去，则第三个桶有毒,否则继续测试
//     45-60 小猪吃第四个桶中的液体，如果死去，则第四个桶有毒
// 如果最后小猪没有死去，则第五个桶有毒
// 所以一只小猪在 minutesToDie 和 minutesToTest 时间一定的情况下可以最多判断 base = minutesToTest / minutesToDie + 1 个桶
// 假设小猪的数量是 num,那么 pow(base, num) >= buckets,根据对数运算规则，两边分别取对数得到: num >= Log10(buckets) / Log10(base)
func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
    base := minutesToTest / minutesToDie + 1
    return int(math.Ceil(math.Log10(float64(buckets)) / math.Log10(float64(base))))
}

func poorPigs1(buckets int, minutesToDie int, minutesToTest int) int {
    testRound := minutesToTest / minutesToDie
    testGroup := testRound + 1
    res := int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(testGroup))))
    return res
}

// 依次递增猪的个数求猪能判断的桶数，也就是 1只猪能判断的桶数 的 次方。
// 然后比传进来的桶数要大就Ok
func poorPigs2(buckets int, minutesToDie int, minutesToTest int) int {
    if buckets == 1 {
        return 0
    }
    in, res := int(minutesToTest / minutesToDie) + 1, 1
    for {
        t := math.Pow(float64(in), float64(res))
        if t >= float64(buckets) {
            return res
        }
        res++
    }
    return 0
}

func poorPigs3(buckets int, minutesToDie int, minutesToTest int) int {
    res, base := 0, minutesToTest / minutesToDie + 1
    for p := 1; p < buckets; p *= base {
        res++
    }
    return res
}

func main() {
    // Example 1:
    // Input: buckets = 4, minutesToDie = 15, minutesToTest = 15
    // Output: 2
    // Explanation: We can determine the poisonous bucket as follows:
    // At time 0, feed the first pig buckets 1 and 2, and feed the second pig buckets 2 and 3.
    // At time 15, there are 4 possible outcomes:
    // - If only the first pig dies, then bucket 1 must be poisonous.
    // - If only the second pig dies, then bucket 3 must be poisonous.
    // - If both pigs die, then bucket 2 must be poisonous.
    // - If neither pig dies, then bucket 4 must be poisonous.
    fmt.Println(poorPigs(4, 15, 15)) // 2
    // Example 2:
    // Input: buckets = 4, minutesToDie = 15, minutesToTest = 30
    // Output: 2
    // Explanation: We can determine the poisonous bucket as follows:
    // At time 0, feed the first pig bucket 1, and feed the second pig bucket 2.
    // At time 15, there are 2 possible outcomes:
    // - If either pig dies, then the poisonous bucket is the one it was fed.
    // - If neither pig dies, then feed the first pig bucket 3, and feed the second pig bucket 4.
    // At time 30, one of the two pigs must die, and the poisonous bucket is the one it was fed.
    fmt.Println(poorPigs(4, 15, 30)) // 2
    fmt.Println(poorPigs(1000, 15, 60)) // 5
    fmt.Println(poorPigs(125, 1, 4)) // 4

    fmt.Println(poorPigs1(4, 15, 15)) // 2
    fmt.Println(poorPigs1(4, 15, 30)) // 2
    fmt.Println(poorPigs1(1000, 15, 60)) // 5
    fmt.Println(poorPigs1(125, 1, 4)) // 4

    fmt.Println(poorPigs2(4, 15, 15)) // 2
    fmt.Println(poorPigs2(4, 15, 30)) // 2
    fmt.Println(poorPigs2(1000, 15, 60)) // 5
    fmt.Println(poorPigs2(125, 1, 4)) // 3

    fmt.Println(poorPigs3(4, 15, 15)) // 2
    fmt.Println(poorPigs3(4, 15, 30)) // 2
    fmt.Println(poorPigs3(1000, 15, 60)) // 5
    fmt.Println(poorPigs3(125, 1, 4)) // 3
}