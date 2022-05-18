/**
2449 · Sleep Sort
Description
In the Solution class of this question there is a sleep_sort method (in Java/C++ it is the sleepSort method) 
that passes in an array nums and a function print_number for printing numbers
in a subthread (in Java it is just an array nums, you can use the class method printNumber under class Main to print numbers, 
and the printNumber function in C++). Please use sleep sort to sort the elements of the array and call the print_number function in a subthread to print the result.

Sleep sorting is mainly implemented according to the CPU's scheduling algorithm. In this algorithm, 
we create different threads for each element of the input array, and then each thread sleeps for a time proportional to the value of the corresponding array element. 
Thus, the thread with the least sleep time is woken up first and prints the number, followed by the second least element, 
and so on. The largest element wakes up after a long time, and then that element is printed last. So the output is an ordered one.

1≤∣nums∣≤10
0 <= nums[i] <= 10

Example
We will print the sequence by calling the sleep_sort (which in Java is just an array nums , you can use the class method printNumber under class Main to print the numbers, in C++ it is the printNumber function) method.

When nums=[0.17, 0.02, 0.1], your code should output.

0.02
0.1
0.17

When nums=[0.9, 0.6, 0.3], your code should output.

0.3
0.6
0.9
 */
// 类实现Runnable接口实现多线程
class SleepSortThread implements Runnable {
	double x;
	public SleepSortThread(double num) {
		super();
		x = num;
	}
	// 重写线程run函数
	public void run() {
		try {
            // 离散化线程休眠时间
			Thread.sleep((int)(x * 1000));
			Main.printNumber(x);
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
}

public class Solution {
    public void sleepSort(double[] nums) throws Exception {
        // write your code
		Thread[] threads = new Thread[nums.length];
		for(int i = 0; i < nums.length; i++) {
            // 将实现Runnable接口的类对象赋值到线程的构造函数               
			threads[i] = new Thread(new SleepSortThread(nums[i]));
			//开启线程
			threads[i].start();
		}
    }
}