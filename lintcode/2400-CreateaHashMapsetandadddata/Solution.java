/**
 * 2400 · Create a HashMap set and add data
# Description
Write code that creates a HashMap collection, adds a piece of data to the collection, and returns the collection.
A HashMap is a hash table that stores the contents of a key-value pair (key-value) map.
There is a createHashMap method in the Solution class of this question, 
which has an int type parameter number and a String type parameter str, 
number represents the key to add data to the collection, and str represents the value to add data to the collection. 
represents the value to add to the collection. This method creates a HashMap collection, 
adds a data to the collection and returns the collection. The return value is of type HashMap.
Write your code in the createHashMap method of class name Solution under // write your code here.here`.

Note the generalization of HashMap set

Example
The reviewer will compile the entire project's code into an executable Main program and execute your code in this manner Main <number> <str>

Sample 1

    If the input data is 1 bejing then return the result and output.
    {1=beijing}

Sample 2

    If the input data is 2 shanghai then the result is returned and the output is
    {2=shanghai}
 */
import java.util.HashMap;

public class Solution {
    /**
     * @param number: Represents the key to pass data into the set
     * @param str: Represents the value of the data passed into the set
     * @return: Represents the return of the created set
     */
    public HashMap<Integer, String> createHashMap(int number, String str) {
        // writer your code here
        HashMap<Integer,String> map = new HashMap<>();
        map.put(number,str);
        return map;
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.createHashMap(1, "beijing"));
        System.out.println(s.createHashMap(2, "shanghai"));
    }
}