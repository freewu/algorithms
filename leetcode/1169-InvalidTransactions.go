package main

// 1169. Invalid Transactions
// A transaction is possibly invalid if:
//     the amount exceeds $1000, or;
//     if it occurs within (and including) 60 minutes of another transaction with the same name in a different city.

// You are given an array of strings transaction 
// where transactions[i] consists of comma-separated values representing the name, time (in minutes), amount, and city of the transaction.

// Return a list of transactions that are possibly invalid. You may return the answer in any order.

// Example 1:
// Input: transactions = ["alice,20,800,mtv","alice,50,100,beijing"]
// Output: ["alice,20,800,mtv","alice,50,100,beijing"]
// Explanation: The first transaction is invalid because the second transaction occurs within a difference of 60 minutes, have the same name and is in a different city. Similarly the second one is invalid too.

// Example 2:
// Input: transactions = ["alice,20,800,mtv","alice,50,1200,mtv"]
// Output: ["alice,50,1200,mtv"]

// Example 3:
// Input: transactions = ["alice,20,800,mtv","bob,50,1200,mtv"]
// Output: ["bob,50,1200,mtv"]

// Constraints:
//     transactions.length <= 1000
//     Each transactions[i] takes the form "{name},{time},{amount},{city}"
//     Each {name} and {city} consist of lowercase English letters, and have lengths between 1 and 10.
//     Each {time} consist of digits, and represent an integer between 0 and 1000.
//     Each {amount} consist of digits, and represent an integer between 0 and 2000.

import "fmt"
import "strings"
import "strconv"

func invalidTransactions(transactions []string) []string {
    type TransInfo struct {
        Time int
        City string
        Raw string
    }
    transMap, transCountMap, resMap := make(map[string][]TransInfo), make(map[string]int), make(map[string]bool)
    for _, transaction := range transactions {
        data := strings.Split(transaction, ",")
        name, city := data[0], data[3]
        time, _ := strconv.Atoi(data[1])
        amt, _ := strconv.Atoi(data[2])
        if amt > 1000 {
            resMap[transaction] = true
        } 
        if dups, ok := transMap[name]; ok {
            for _, dup := range dups {
                timeDiff := time - dup.Time
                if timeDiff <= 60 && timeDiff >= -60 && city != dup.City {
                    resMap[transaction], resMap[dup.Raw] = true, true
                }
            }
        }
        transMap[name] = append(transMap[name], TransInfo{ time, city, transaction })
        transCountMap[transaction] += 1
    }
    res := []string{}
    for v := range resMap {
        for i := 0; i < transCountMap[v]; i++ {
            res = append(res, v)
        }
    }
    return res
}

func invalidTransactions1(transactions []string) []string {
    res := []string{}
    type Transaction struct {
        name, city string
        time, amount int
    }
    transactionList, transactionByName := make([]Transaction, len(transactions)), make(map[string][]Transaction)
    for i, transaction := range transactions {
        parts := strings.Split(transaction, ",")
        time, _ := strconv.Atoi(parts[1])
        amount, _ := strconv.Atoi(parts[2])
        transactionList[i] = Transaction{
            name: parts[0],
            time: time,
            amount: amount,
            city: parts[3],
        }
        transactionByName[parts[0]] = append(transactionByName[parts[0]], transactionList[i])
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for _, trans := range transactionList {
        if trans.amount > 1000 {
            res = append(res, trans.name + "," + strconv.Itoa(trans.time) + "," + strconv.Itoa(trans.amount) + "," + trans.city)
            continue
        }
        for _, other := range transactionByName[trans.name] {
            if other.city != trans.city && abs(other.time - trans.time) <= 60 {
                res = append(res, trans.name + "," + strconv.Itoa(trans.time) + "," + strconv.Itoa(trans.amount) + "," + trans.city)
                break
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: transactions = ["alice,20,800,mtv","alice,50,100,beijing"]
    // Output: ["alice,20,800,mtv","alice,50,100,beijing"]
    // Explanation: The first transaction is invalid because the second transaction occurs within a difference of 60 minutes, have the same name and is in a different city. Similarly the second one is invalid too.
    fmt.Println(invalidTransactions([]string{"alice,20,800,mtv","alice,50,100,beijing"})) // ["alice,20,800,mtv","alice,50,100,beijing"]
    // Example 2:
    // Input: transactions = ["alice,20,800,mtv","alice,50,1200,mtv"]
    // Output: ["alice,50,1200,mtv"]
    fmt.Println(invalidTransactions([]string{"alice,20,800,mtv","alice,50,1200,mtv"})) // ["alice,50,1200,mtv"]
    // Example 3:
    // Input: transactions = ["alice,20,800,mtv","bob,50,1200,mtv"]
    // Output: ["bob,50,1200,mtv"]
    fmt.Println(invalidTransactions([]string{"alice,20,800,mtv","bob,50,1200,mtv"})) // ["bob,50,1200,mtv"]

    fmt.Println(invalidTransactions1([]string{"alice,20,800,mtv","alice,50,100,beijing"})) // ["alice,20,800,mtv","alice,50,100,beijing"]
    fmt.Println(invalidTransactions1([]string{"alice,20,800,mtv","alice,50,1200,mtv"})) // ["alice,50,1200,mtv"]
    fmt.Println(invalidTransactions1([]string{"alice,20,800,mtv","bob,50,1200,mtv"})) // ["bob,50,1200,mtv"]
}