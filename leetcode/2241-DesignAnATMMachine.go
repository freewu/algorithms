package main

// 2241. Design an ATM Machine
// There is an ATM machine that stores banknotes of 5 denominations: 20, 50, 100, 200, and 500 dollars. 
// Initially the ATM is empty. The user can use the machine to deposit or withdraw any amount of money.

// When withdrawing, the machine prioritizes using banknotes of larger values.
//     1. For example, if you want to withdraw $300 
//        and there are 2 $50 banknotes, 1 $100 banknote, and 1 $200 banknote, 
//        then the machine will use the $100 and $200 banknotes.
//     2. However, if you try to withdraw $600 and there are 3 $200 banknotes and 1 $500 banknote, 
//        then the withdraw request will be rejected because the machine will first try to use the $500 banknote 
//        and then be unable to use banknotes to complete the remaining $100. 
//        Note that the machine is not allowed to use the $200 banknotes instead of the $500 banknote.

// Implement the ATM class:
//     ATM() 
//         Initializes the ATM object.
//     void deposit(int[] banknotesCount) 
//         Deposits new banknotes in the order $20, $50, $100, $200, and $500.
//     int[] withdraw(int amount) 
//         Returns an array of length 5 of the number of banknotes 
//         that will be handed to the user in the order $20, $50, $100, $200, and $500, 
//         and update the number of banknotes in the ATM after withdrawing. 
//         Returns [-1] if it is not possible (do not withdraw any banknotes in this case).

// Example 1:
// Input
// ["ATM", "deposit", "withdraw", "deposit", "withdraw", "withdraw"]
// [[], [[0,0,1,2,1]], [600], [[0,1,0,1,1]], [600], [550]]
// Output
// [null, null, [0,0,1,0,1], null, [-1], [0,1,0,0,1]]
// Explanation
// ATM atm = new ATM();
// atm.deposit([0,0,1,2,1]); // Deposits 1 $100 banknote, 2 $200 banknotes,
//                           // and 1 $500 banknote.
// atm.withdraw(600);        // Returns [0,0,1,0,1]. The machine uses 1 $100 banknote
//                           // and 1 $500 banknote. The banknotes left over in the
//                           // machine are [0,0,0,2,0].
// atm.deposit([0,1,0,1,1]); // Deposits 1 $50, $200, and $500 banknote.
//                           // The banknotes in the machine are now [0,1,0,3,1].
// atm.withdraw(600);        // Returns [-1]. The machine will try to use a $500 banknote
//                           // and then be unable to complete the remaining $100,
//                           // so the withdraw request will be rejected.
//                           // Since the request is rejected, the number of banknotes
//                           // in the machine is not modified.
// atm.withdraw(550);        // Returns [0,1,0,0,1]. The machine uses 1 $50 banknote
//                           // and 1 $500 banknote.

// Constraints:
// banknotesCount.length == 5
//     0 <= banknotesCount[i] <= 10^9
//     1 <= amount <= 10^9
//     At most 5000 calls in total will be made to withdraw and deposit.
//     At least one call will be made to each function withdraw and deposit.
//     Sum of banknotesCount[i] in all deposits doesn't exceed 10^9

import "fmt"

type ATM struct {
    Banknotes   [5]int // 面值
    Money       [5]int // 数量
    Amount      int // 账户金额
}

func Constructor() ATM {
    return ATM{
        Banknotes: [5]int{20,50,100,200,500},
        Money: [5]int{},
        Amount: 0,
    }
}

func (this *ATM) Deposit(banknotesCount []int)  {
    for i,v := range banknotesCount {
        this.Money[i] += v 
        this.Amount += banknotesCount[i] * this.Banknotes[i]
    }
}

func (this *ATM) Withdraw(amount int) []int {
    if amount > this.Amount { return []int{-1} }
    res, temp := make([]int,5), this.Money
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 4; i >= 0; i-- {
        if(temp[i] > 0 && amount >= this.Banknotes[i]) {
            count := (amount / this.Banknotes[i])
            res[i] = min(count,temp[i])
            temp[i] -= res[i]
            amount -= res[i] * this.Banknotes[i]
        }
    }    
    if amount != 0 { return []int{-1} }
    this.Money = temp
    this.Amount -= amount
    return res
}

/**
* Your ATM object will be instantiated and called as such:
* obj := Constructor();
* obj.Deposit(banknotesCount);
* param_2 := obj.Withdraw(amount);
*/

func main() {
    // ATM atm = new ATM();
    obj := Constructor()
    fmt.Println(obj)
    // atm.deposit([0,0,1,2,1]); // Deposits 1 $100 banknote, 2 $200 banknotes,
    //                           // and 1 $500 banknote.
    obj.Deposit([]int{0,0,1,2,1})
    fmt.Println(obj)
    // atm.withdraw(600);        // Returns [0,0,1,0,1]. The machine uses 1 $100 banknote
    //                           // and 1 $500 banknote. The banknotes left over in the
    //                           // machine are [0,0,0,2,0].
    fmt.Println(obj.Withdraw(600)) // [0,0,1,0,1]
    // atm.deposit([0,1,0,1,1]); // Deposits 1 $50, $200, and $500 banknote.
    //                           // The banknotes in the machine are now [0,1,0,3,1].
    obj.Deposit([]int{0,1,0,1,1})
    fmt.Println(obj)
    // atm.withdraw(600);        // Returns [-1]. The machine will try to use a $500 banknote
    //                           // and then be unable to complete the remaining $100,
    //                           // so the withdraw request will be rejected.
    //                           // Since the request is rejected, the number of banknotes
    //                           // in the machine is not modified.
    fmt.Println(obj.Withdraw(600)) // [-1]
    // atm.withdraw(550);        // Returns [0,1,0,0,1]. The machine uses 1 $50 banknote
    //                           // and 1 $500 banknote.
    fmt.Println(obj.Withdraw(550)) // [0,1,0,0,1]
}