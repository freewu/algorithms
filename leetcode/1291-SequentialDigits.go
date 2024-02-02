package main

// 1291. Sequential Digits
// An integer has sequential digits if and only if each digit in the number is one more than the previous digit.
// Return a sorted list of all the integers in the range [low, high] inclusive that have sequential digits.

// Example 1:
// Input: low = 100, high = 300
// Output: [123,234]

// Example 2:
// Input: low = 1000, high = 13000
// Output: [1234,2345,3456,4567,5678,6789,12345]
 
// Constraints:
// 		10 <= low <= high <= 10^9

// 范围最大为 9 位数，所以直接枚举即可。采用的方法是滑动窗口；
// 定义最长串  s := "123456789";
// 然后以最小长度到最大长长度在最长串上截取
import "strconv"
import "fmt"

func sequentialDigits(low int, high int) []int {
    ans := make([]int, 0)
	s := "123456789";
	min_len, max_len := len(strconv.Itoa(low)), len(strconv.Itoa(high))
	for min_len <= max_len {
		temp := 0;
		for i := 0; i <= 9 - min_len; i = i + 1 {
			temp, _ = strconv.Atoi(s[i:i + min_len]);
			if temp < low {
				continue
			} else if temp > high {
				break
			} else {
				ans = append(ans, temp);
			}
		}
		min_len = min_len + 1
	}
	return ans;
}

func main() {
	fmt.Println(sequentialDigits(100,300)) // [123,234]
	fmt.Println(sequentialDigits(1000,13000)) // [1234,2345,3456,4567,5678,6789,12345]
}