package main

// 1694. Reformat Phone Number
// You are given a phone number as a string number. 
// number consists of digits, spaces ' ', and/or dashes '-'.

// You would like to reformat the phone number in a certain manner. 
// Firstly, remove all spaces and dashes. 
// Then, group the digits from left to right into blocks of length 3 until there are 4 or fewer digits. 
// The final digits are then grouped as follows:
//     2 digits: A single block of length 2.
//     3 digits: A single block of length 3.
//     4 digits: Two blocks of length 2 each.

// The blocks are then joined by dashes. 
// Notice that the reformatting process should never produce any blocks of length 1 and produce at most two blocks of length 2.

// Return the phone number after formatting.

// Example 1:
// Input: number = "1-23-45 6"
// Output: "123-456"
// Explanation: The digits are "123456".
// Step 1: There are more than 4 digits, so group the next 3 digits. The 1st block is "123".
// Step 2: There are 3 digits remaining, so put them in a single block of length 3. The 2nd block is "456".
// Joining the blocks gives "123-456".

// Example 2:
// Input: number = "123 4-567"
// Output: "123-45-67"
// Explanation: The digits are "1234567".
// Step 1: There are more than 4 digits, so group the next 3 digits. The 1st block is "123".
// Step 2: There are 4 digits left, so split them into two blocks of length 2. The blocks are "45" and "67".
// Joining the blocks gives "123-45-67".

// Example 3:
// Input: number = "123 4-5678"
// Output: "123-456-78"
// Explanation: The digits are "12345678".
// Step 1: The 1st block is "123".
// Step 2: The 2nd block is "456".
// Step 3: There are 2 digits left, so put them in a single block of length 2. The 3rd block is "78".
// Joining the blocks gives "123-456-78".

// Constraints:
//     2 <= number.length <= 100
//     number consists of digits and the characters '-' and ' '.
//     There are at least two digits in number.

import "fmt"
import "regexp"

func reformatNumber(number string) string {
    arr := []byte{}
    for i := range number { // 提取 0 - 9
        if number[i] >= '0' && number[i] <= '9' {
            arr = append(arr, number[i])
        }
    }
    index, n, res := 0, len(arr), []byte{}
    if n <= 3 {
        return string(arr)
    }
    for n > 0 {
        if n <= 4 {
            if n == 4 { // 处理结尾只有4个的情况 需要分成  xx-xx
                res = append(res, arr[index:index + 2]...)
                res = append(res, '-')
                res = append(res, arr[index + 2:]...)
            } else { // 2-3 直接添加到尾部即可
                res = append(res, arr[index:]...)
            }
            break
        } else {
            res = append(res, arr[index:index + 3]...)
            res = append(res, '-')
            index += 3
            n -= 3
        }
    }
    return string(res)
}

func reformatNumber1(number string) string {
    var solve func(number string) string
    solve = func(number string) string {
        n := len(number)
        switch {
            case n <= 3: return number
            case n == 4: return number[0:2] + "-" + number[2:]
            default: return number[0:3] + "-" + solve(number[3:])
        }
    }
    return solve(regexp.MustCompile(`[ -]`).ReplaceAllString(number, ""))
}

func reformatNumber2(number string) string {
    res, num := "", ""
    for _, v := range number {
        if v == ' ' || v == '-' { continue
        }
        num += string(v)
    }
    switch len(num) % 3 {
    case 0:
        for i := 0; i < len(num); i += 3 {
            res += num[i:i+3] + "-"
        }
        res = res[:len(res)-1]
    case 1: // xx-xx  的情况
        i := 0
        for ; i < len(num)-4; i += 3 {
            res += num[i:i+3] + "-"
        }
        res += num[i:i+2] + "-" + num[i+2:]
    case 2:
        i := 0
        for ; i < len(num)-2; i += 3 {
            res += num[i:i+3] + "-"
        }
        res += num[i:]
    }
    return res
}

func main() {
    // Example 1:
    // Input: number = "1-23-45 6"
    // Output: "123-456"
    // Explanation: The digits are "123456".
    // Step 1: There are more than 4 digits, so group the next 3 digits. The 1st block is "123".
    // Step 2: There are 3 digits remaining, so put them in a single block of length 3. The 2nd block is "456".
    // Joining the blocks gives "123-456".
    fmt.Println(reformatNumber("1-23-45 6")) // "123-456"
    // Example 2:
    // Input: number = "123 4-567"
    // Output: "123-45-67"
    // Explanation: The digits are "1234567".
    // Step 1: There are more than 4 digits, so group the next 3 digits. The 1st block is "123".
    // Step 2: There are 4 digits left, so split them into two blocks of length 2. The blocks are "45" and "67".
    // Joining the blocks gives "123-45-67".
    fmt.Println(reformatNumber("123 4-567")) // "123-45-67"
    // Example 3:
    // Input: number = "123 4-5678"
    // Output: "123-456-78"
    // Explanation: The digits are "12345678".
    // Step 1: The 1st block is "123".
    // Step 2: The 2nd block is "456".
    // Step 3: There are 2 digits left, so put them in a single block of length 2. The 3rd block is "78".
    // Joining the blocks gives "123-456-78".
    fmt.Println(reformatNumber("123 4-5678")) // "123-456-78"

    fmt.Println(reformatNumber1("1-23-45 6")) // "123-456"
    fmt.Println(reformatNumber1("123 4-567")) // "123-45-67"
    fmt.Println(reformatNumber1("123 4-5678")) // "123-456-78"

    fmt.Println(reformatNumber2("1-23-45 6")) // "123-456"
    fmt.Println(reformatNumber2("123 4-567")) // "123-45-67"
    fmt.Println(reformatNumber2("123 4-5678")) // "123-456-78"
}