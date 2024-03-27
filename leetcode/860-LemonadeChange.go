package main

// 860. Lemonade Change
// At a lemonade stand, each lemonade costs $5. 
// Customers are standing in a queue to buy from you and order one at a time (in the order specified by bills). 
// Each customer will only buy one lemonade and pay with either a $5, $10, or $20 bill. 
// You must provide the correct change to each customer so that the net transaction is that the customer pays $5.

// Note that you do not have any change in hand at first.

// Given an integer array bills where bills[i] is the bill the ith customer pays, 
// return true if you can provide every customer with the correct change, or false otherwise.

// Example 1:
// Input: bills = [5,5,5,10,20]
// Output: true
// Explanation: 
// From the first 3 customers, we collect three $5 bills in order.
// From the fourth customer, we collect a $10 bill and give back a $5.
// From the fifth customer, we give a $10 bill and a $5 bill.
// Since all customers got correct change, we output true.

// Example 2:
// Input: bills = [5,5,10,10,20]
// Output: false
// Explanation: 
// From the first two customers in order, we collect two $5 bills.
// For the next two customers in order, we collect a $10 bill and give back a $5 bill.
// For the last customer, we can not give the change of $15 back because we only have two $10 bills.
// Since not every customer received the correct change, the answer is false.
 
// Constraints:
//     1 <= bills.length <= 10^5
//     bills[i] is either 5, 10, or 20.

import "fmt"

func lemonadeChange(bills []int) bool {
    cashes := []int{0,0} // 依次表示 5, 10 面额数量
    for _,bill := range bills {
        switch bill{
            case 5: // 5块的直接收
                cashes[0]++
            case 10:
                if cashes[0] > 0 { // 需要找 5 元
                    cashes[0]--
                    cashes[1]++
                } else {
                    return false // 不够找
                }
            case 20:
                change := 15
                if cashes[1] > 0 { // 优先使用 10 元找零
                    cashes[1]--
                    change -= 10
                }
                for cashes[0] > 0 && change > 0 {
                    cashes[0]--
                    change -= 5
                }
                if change > 0 {
                    return false
                }
        }
    }
    return true
}

func lemonadeChange1(bills []int) bool {
    five, ten := 0, 0
    for _, bill := range bills {
        if bill == 5 {
            five++
        } else if bill == 10 {
            if five == 0 {
                 return false
            }
            five--
            ten++
        } else { // 处理 20 的
            if five > 0 && ten > 0 { // 10,5 都有 找一张 10(ten--) 一张 5(five--)
                five--
                ten--
            } else if five >= 3 { // 只有 5 元 且 大于等于 3张  3 * 5
                five -= 3
            } else {
                return false
            }
        }
    }
    return true
}

func main() {
    // From the first 3 customers, we collect three $5 bills in order.
    // From the fourth customer, we collect a $10 bill and give back a $5.
    // From the fifth customer, we give a $10 bill and a $5 bill.
    // Since all customers got correct change, we output true.
    fmt.Println(lemonadeChange([]int{5,5,5,10,20})) // true

    // Explanation: 
    // From the first two customers in order, we collect two $5 bills.
    // For the next two customers in order, we collect a $10 bill and give back a $5 bill.
    // For the last customer, we can not give the change of $15 back because we only have two $10 bills.
    // Since not every customer received the correct change, the answer is false.
    fmt.Println(lemonadeChange([]int{5,5,10,10,20})) // false

    fmt.Println(lemonadeChange1([]int{5,5,5,10,20})) // true
    fmt.Println(lemonadeChange1([]int{5,5,10,10,20})) // false
}