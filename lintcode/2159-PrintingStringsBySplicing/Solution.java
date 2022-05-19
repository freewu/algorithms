/**
2159 · Printing strings by splicing
Description
Write Java statements that splice the incoming strings.
The problem provides the Solution class, 
which has a splice method that passes a parameter name of type String and a parameter phone of type String, 
which is to return (name spliced with a : and then ), and the return value is of type String.

You need to write your code under // write your code here

Note the format of the splice
The input string cannot be empty
Example
The evaluator will compile your code into an executable file Main and execute it, 
the code will read in each test data in turn from the input folder for evaluation.
You can find out more about this in Main.java. In each test data, we include 2 parameters, 
the first parameter is name and the second parameter is phone.

Your code should output different results for the different values of name and phone.

Example 1

When name takes the value "Jack" and phone takes the value "15570729587", your code should output

    Jack: 15570729587

Example 2

When name takes the value "Rose " and phone takes the value "15270729587", your code should output.

    Rose: 15270729587
 */

public class Solution {
    public String splice(String name, String phone) {
        // write your code here
        return name + ": " + phone;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.splice("Jack", "15570729587"));
        System.out.println(solution.splice("Rose", "15270729587"));
    }
}