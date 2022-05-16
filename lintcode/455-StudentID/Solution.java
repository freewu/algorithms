/**
455 · Student ID

# Description
Implement a class Class with the following attributes and methods:

A public attribute students which is a array of Student instances.
A constructor with a parameter n, which is the total number of students in this class. 
The constructor should create n Student instances and initialized with student id from 0 ~ n-1

Example 1:

    Input: 3
    Output: [0, 1, 2]
    Explanation: For 3 students, your cls.students[0] should equal to 0,
    cls.students[1] should equal to 1 and the cls.students[2] should equal to 2.

Example 2:

    Input: 5
    Output: [0, 1, 2, 3, 4]
 */

public class Solution {
    class Student {
        public int id;
        
        public Student(int id) {
            this.id = id;
        }
    }
    
    public class Class {
        /**
         * Declare a public attribute `students` which is an array of `Student`
         * instances
         */
        // write your code here.
        public Student[] students;
         
        /**
         * Declare a constructor with a parameter n which is the total number of
         * students in the *class*. The constructor should create n Student
         * instances and initialized with student id from 0 ~ n-1
         */
        // write your code here
        public Class(int n) {
            this.students = new Student[n];
            for(int i = 0;i < n; i++) {
                this.students[i] = new Student(i);
            }
        }
    }

    public static void main(String[] args) {
        System.out.println(new Class(3));
        System.out.println(new Class(5));
    }
}
