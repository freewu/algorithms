/**
2848 · Refining the right class
# Description
You need to refine the relevant code in the Solution class based on the code in the Main.java file. 
Add the show() method to the Solution class to print the information.

Example
No input is required for this question. The evaluator compiles the entire project's code into an executable Main program and executes your code in such a way that your code prints the result to the standard output stream (console) when it is finished running.

The result of the run should be

Name: Lemon, age: 20
 
 */

class Solution {
	// write your code here
    public String name;
    public int age;

    public void show() {
        System.out.println("Name: " + name + ", age: " + age);
    }

    public static void main(String[] args) {
		Solution s1 = new Solution();
		s1.name = "Lemon";
		s1.age = 20;
		s1.show();
	}
}
