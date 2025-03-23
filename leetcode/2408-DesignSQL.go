package main

// 2408. Design SQL
// // You are given n tables represented with two arrays names and columns, 
// // where names[i] is the name of the ith table and columns[i] is the number of columns of the ith table.

// // You should be able to perform the following operations:
// //     1. Insert a row in a specific table. 
// //        Each row you insert has an id. The id is assigned using an auto-increment method where the id of the first inserted row is 1, 
// //        and the id of each other row inserted into the same table is the id of the last inserted row (even if it was deleted) plus one.
// //     2. Delete a row from a specific table. Note that deleting a row does not affect the id of the next inserted row.
// //     3. Select a specific cell from any table and return its value.

// // Implement the SQL class:
// //     SQL(String[] names, int[] columns) 
// //         Creates the n tables.
// //     void insertRow(String name, String[] row) 
// //         Adds a row to the table name. 
// //         It is guaranteed that the table will exist, and the size of the array row is equal to the number of columns in the table.
// //     void deleteRow(String name, int rowId) 
// //         Removes the row rowId from the table name. 
// //         It is guaranteed that the table and row will exist.
// //     String selectCell(String name, int rowId, int columnId) 
// //         Returns the value of the cell in the row rowId and the column columnId from the table name.

// // Example 1:
// // Input
// // ["SQL", "insertRow", "selectCell", "insertRow", "deleteRow", "selectCell"]
// // [[["one", "two", "three"], [2, 3, 1]], ["two", ["first", "second", "third"]], ["two", 1, 3], ["two", ["fourth", "fifth", "sixth"]], ["two", 1], ["two", 2, 2]]
// // Output
// // [null, null, "third", null, null, "fifth"]
// // Explanation
// // SQL sql = new SQL(["one", "two", "three"], [2, 3, 1]); // creates three tables.
// // sql.insertRow("two", ["first", "second", "third"]); // adds a row to the table "two". Its id is 1.
// // sql.selectCell("two", 1, 3); // return "third", finds the value of the third column in the row with id 1 of the table "two".
// // sql.insertRow("two", ["fourth", "fifth", "sixth"]); // adds another row to the table "two". Its id is 2.
// // sql.deleteRow("two", 1); // deletes the first row of the table "two". Note that the second row will still have the id 2.
// // sql.selectCell("two", 2, 2); // return "fifth", finds the value of the second column in the row with id 2 of the table "two".

// // Constraints:
// //     n == names.length == columns.length
// //     1 <= n <= 10^4
// //     1 <= names[i].length, row[i].length, name.length <= 20
// //     names[i], row[i], and name consist of lowercase English letters.
// //     1 <= columns[i] <= 100
// //     All the strings of names are distinct.
// //     name exists in the array names.
// //     row.length equals the number of columns in the chosen table.
// //     rowId and columnId will be valid.
// //     At most 250 calls will be made to insertRow and deleteRow.
// //     At most 104 calls will be made to selectCell.

// import "fmt"

// type SQL struct {
//     tables map[string][][]string
// }

// func Constructor(names []string, columns []int) SQL {
//     return SQL{ tables: map[string][][]string{} }
// }

// func (this *SQL) InsertRow(name string, row []string) {
//     this.tables[name] = append(this.tables[name], row)
// }

// func (this *SQL) DeleteRow(name string, rowId int) {
// }

// func (this *SQL) SelectCell(name string, rowId int, columnId int) string {
//     return this.tables[name][rowId-1][columnId-1]
// }

// /**
//  * Your SQL object will be instantiated and called as such:
//  * obj := Constructor(names, columns);
//  * obj.InsertRow(name,row);
//  * obj.DeleteRow(name,rowId);
//  * param_3 := obj.SelectCell(name,rowId,columnId);
//  */

// func main() {
//     // SQL sql = new SQL(["one", "two", "three"], [2, 3, 1]); // creates three tables.
//     obj := Constructor([]string{"one", "two", "three"}, []int{2, 3, 1}) 
//     fmt.Println(obj)
//     // sql.insertRow("two", ["first", "second", "third"]); // adds a row to the table "two". Its id is 1.
//     obj.InsertRow("two", []string{"first", "second", "third"})
//     fmt.Println(obj)
//     // sql.selectCell("two", 1, 3); // return "third", finds the value of the third column in the row with id 1 of the table "two".
//     fmt.Println(obj.SelectCell("two", 1, 3)) // "third"
//     // sql.insertRow("two", ["fourth", "fifth", "sixth"]); // adds another row to the table "two". Its id is 2.
//     obj.InsertRow("two", []string{"fourth", "fifth", "sixth"})
//     fmt.Println(obj)
//     // sql.deleteRow("two", 1); // deletes the first row of the table "two". Note that the second row will still have the id 2.
//     obj.DeleteRow("two", 1)
//     fmt.Println(obj)
//     // sql.selectCell("two", 2, 2); // return "fifth", finds the value of the second column in the row with id 2 of the table "two".
//     fmt.Println(obj.SelectCell("two", 2, 2)) // "fifth"
// }


// You are given two string arrays, names and columns, both of size n.
// The ith table is represented by the name names[i] and contains columns[i] number of columns.

// You need to implement a class that supports the following operations:
//     1. Insert a row in a specific table with an id assigned using an auto-increment method, where the id of the first inserted row is 1, and the id of each new row inserted into the same table is one greater than the id of the last inserted row, even if the last row was removed.
//     2. Remove a row from a specific table. Removing a row does not affect the id of the next inserted row.
//     3. Select a specific cell from any table and return its value.
//     4. Export all rows from any table in csv format.

// Implement the SQL class:
//     SQL(String[] names, int[] columns)
//         Creates the n tables.
//     bool ins(String name, String[] row)
//         Inserts row into the table name and returns true.
//         If row.length does not match the expected number of columns, or name is not a valid table, returns false without any insertion.
//     void rmv(String name, int rowId)
//         Removes the row rowId from the table name.
//         If name is not a valid table or there is no row with id rowId, no removal is performed.
//     String sel(String name, int rowId, int columnId)
//         Returns the value of the cell at the specified rowId and columnId in the table name.
//         If name is not a valid table, or the cell (rowId, columnId) is invalid, returns "<null>".
//     String[] exp(String name)
//         Returns the rows present in the table name.
//         If name is not a valid table, returns an empty array. 
//         Each row is represented as a string, with each cell value (including the row's id) separated by a ",".

// Example 1:
// Input:
// ["SQL","ins","sel","ins","exp","rmv","sel","exp"]
// [[["one","two","three"],[2,3,1]],["two",["first","second","third"]],["two",1,3],["two",["fourth","fifth","sixth"]],["two"],["two",1],["two",2,2],["two"]]
// Output:
// [null,true,"third",true,["1,first,second,third","2,fourth,fifth,sixth"],null,"fifth",["2,fourth,fifth,sixth"]]
// Explanation:
// // Creates three tables.
// SQL sql = new SQL(["one", "two", "three"], [2, 3, 1]);
// // Adds a row to the table "two" with id 1. Returns True.
// sql.ins("two", ["first", "second", "third"]);
// // Returns the value "third" from the third column
// // in the row with id 1 of the table "two".
// sql.sel("two", 1, 3);
// // Adds another row to the table "two" with id 2. Returns True.
// sql.ins("two", ["fourth", "fifth", "sixth"]);
// // Exports the rows of the table "two".
// // Currently, the table has 2 rows with ids 1 and 2.
// sql.exp("two");
// // Removes the first row of the table "two". Note that the second row
// // will still have the id 2.
// sql.rmv("two", 1);
// // Returns the value "fifth" from the second column
// // in the row with id 2 of the table "two".
// sql.sel("two", 2, 2);
// // Exports the rows of the table "two".
// // Currently, the table has 1 row with id 2.
// sql.exp("two");

// Example 2:
// Input:
// ["SQL","ins","sel","rmv","sel","ins","ins"]
// [[["one","two","three"],[2,3,1]],["two",["first","second","third"]],["two",1,3],["two",1],["two",1,2],["two",["fourth","fifth"]],["two",["fourth","fifth","sixth"]]]
// Output:
// [null,true,"third",null,"<null>",false,true]
// Explanation:
// // Creates three tables.
// SQL sQL = new SQL(["one", "two", "three"], [2, 3, 1]); 
// // Adds a row to the table "two" with id 1. Returns True. 
// sQL.ins("two", ["first", "second", "third"]); 
// // Returns the value "third" from the third column 
// // in the row with id 1 of the table "two".
// sQL.sel("two", 1, 3); 
// // Removes the first row of the table "two".
// sQL.rmv("two", 1); 
// // Returns "<null>" as the cell with id 1 
// // has been removed from table "two".
// sQL.sel("two", 1, 2); 
// // Returns False as number of columns are not correct.
// sQL.ins("two", ["fourth", "fifth"]); 
// // Adds a row to the table "two" with id 2. Returns True.
// sQL.ins("two", ["fourth", "fifth", "sixth"]); 

// Constraints:
//     n == names.length == columns.length
//     1 <= n <= 104
//     1 <= names[i].length, row[i].length, name.length <= 10
//     names[i], row[i], and name consist only of lowercase English letters.
//     1 <= columns[i] <= 10
//     1 <= row.length <= 10
//     All names[i] are distinct.
//     At most 2000 calls will be made to ins and rmv.
//     At most 104 calls will be made to sel.
//     At most 500 calls will be made to exp.
// Follow-up: Which approach would you choose if the table might become sparse due to many deletions, and why? Consider the impact on memory usage and performance.

import "fmt"
import "strings"
import "strconv"

type SQL struct {
    tables map[string]*Table
}

type Table struct {
    id         int
    size       int
    mapIDToRow map[int][]string
}

func Constructor(names []string, columns []int) SQL {
    sql := SQL{ map[string]*Table{} }
    for i := 0; i < len(names); i++ {
        table := &Table{
            id:         1,
            size:       columns[i],
            mapIDToRow: make(map[int][]string),
        }
        sql.tables[names[i]] = table
    }
    return sql
}

func (this *SQL) Ins(name string, row []string) bool {
    if this.tables[name] == nil { return false } // 表不存在
    if len(row) != this.tables[name].size { return false }
    id := this.tables[name].id
    this.tables[name].mapIDToRow[id] = row
    this.tables[name].id++
    return true
}

func (this *SQL) Rmv(name string, rowId int) {
    if this.tables[name] == nil { return }
    if _, ok := this.tables[name].mapIDToRow[rowId]; !ok { return }
    delete(this.tables[name].mapIDToRow, rowId)
}

func (this *SQL) Sel(name string, rowId int, columnId int) string {
    if this.tables[name] == nil { return "<null>" }
    _, ok := this.tables[name].mapIDToRow[rowId]
    if !ok || columnId > this.tables[name].size { return "<null>" }
    return this.tables[name].mapIDToRow[rowId][columnId - 1]
}

func (this *SQL) Exp(name string) []string {
    res := []string{}
    if this.tables[name] == nil { return res }
    for id, row := range this.tables[name].mapIDToRow {
        currentRow := strings.Builder{}
        rowContent := strings.Join(row, ",")
        currentRow.WriteString(strconv.Itoa(id))
        currentRow.WriteString(",")
        currentRow.WriteString(rowContent)
        res = append(res, currentRow.String())
    }
    return res
}

func main() {
    // Explanation:
    // // Creates three tables.
    // SQL sql = new SQL(["one", "two", "three"], [2, 3, 1]);
    obj1 := Constructor([]string{"one", "two", "three"}, []int{2, 3, 1})
    fmt.Println(obj1)
    // // Adds a row to the table "two" with id 1. Returns True.
    // sql.ins("two", ["first", "second", "third"]);
    obj1.Ins("two",[]string{"first", "second", "third"})
    fmt.Println(obj1)
    // // Returns the value "third" from the third column
    // // in the row with id 1 of the table "two".
    // sql.sel("two", 1, 3);
    fmt.Println(obj1.Sel("two", 1, 3))  // third
    // // Adds another row to the table "two" with id 2. Returns True.
    // sql.ins("two", ["fourth", "fifth", "sixth"]);
    obj1.Ins("two",[]string{"fourth", "fifth", "sixth"})
    fmt.Println(obj1)
    // // Exports the rows of the table "two".
    // // Currently, the table has 2 rows with ids 1 and 2.
    // sql.exp("two");
    fmt.Println(obj1.Exp("two"))
    // // Removes the first row of the table "two". Note that the second row
    // // will still have the id 2.
    // sql.rmv("two", 1);
    obj1.Rmv("two", 1)
    fmt.Println(obj1)
    // // Returns the value "fifth" from the second column
    // // in the row with id 2 of the table "two".
    // sql.sel("two", 2, 2);
    fmt.Println(obj1.Sel("two", 2, 2))  // fifth
    // // Exports the rows of the table "two".
    // // Currently, the table has 1 row with id 2.
    // sql.exp("two");
    fmt.Println(obj1.Exp("two"))

    // Example 2:
    // Input:
    // ["SQL","ins","sel","rmv","sel","ins","ins"]
    // [[["one","two","three"],[2,3,1]],["two",["first","second","third"]],["two",1,3],["two",1],["two",1,2],["two",["fourth","fifth"]],["two",["fourth","fifth","sixth"]]]
    // Output:
    // [null,true,"third",null,"<null>",false,true]
    // Explanation:
    // // Creates three tables.
    // SQL sQL = new SQL(["one", "two", "three"], [2, 3, 1]); 
    // // Adds a row to the table "two" with id 1. Returns True. 
    // sQL.ins("two", ["first", "second", "third"]); 
    // // Returns the value "third" from the third column 
    // // in the row with id 1 of the table "two".
    // sQL.sel("two", 1, 3); 
    // // Removes the first row of the table "two".
    // sQL.rmv("two", 1); 
    // // Returns "<null>" as the cell with id 1 
    // // has been removed from table "two".
    // sQL.sel("two", 1, 2); 
    // // Returns False as number of columns are not correct.
    // sQL.ins("two", ["fourth", "fifth"]); 
    // // Adds a row to the table "two" with id 2. Returns True.
    // sQL.ins("two", ["fourth", "fifth", "sixth"]); 

}