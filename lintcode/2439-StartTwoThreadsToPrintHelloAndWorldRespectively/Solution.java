/**
2439 · Start 2 threads to print "Hello" and "World" respectively
# Description
The question asks to print HelloWorld.
We want you to do this in a multi-threaded way. 
The code you need to write is a class called Solution, 
where the hello method is executed by the first thread and the world method is executed by the second thread.
Note that these two threads are executed concurrently. 
These two methods are passed two functions printHello, printWorld (print_hello, print_world in Python) to output "Hello" and "World" respectively. 
And after both threads have finished executing, your code should be outputting strings like HelloWorld as described above.

For the printHello (print_hello in Python) function, you don't need to pass any arguments; 
calling the function will print "Hello". Similarly, for the printWorld (print_world in Python) function, 
you don't need to pass any arguments, calling the function will print "World".
So your code needs to figure out how to call the two functions we provide to you in the following order.

printHello()
printWorld()
 */
import java.util.concurrent.Semaphore;

// solution 1 Semaphore
public class Solution {
    // write your code here
    // 创建两个信号量对象 printHello 和 printWorld，printHello 的信号量为 1 则允许访问，printWorld 的信号量为 0 则阻塞等待
    Semaphore printHello = new Semaphore(1);
    Semaphore printWorld = new Semaphore(0);
    
    public void printHello() throws InterruptedException {
        // write your code here
        // 从信号量获取一个许可
        printHello.acquire();
        // 打印 printHello 中的内容
        System.out.print("Hello");
        // 使信号量许可数增加
        printWorld.release();
    }

    public void printWorld() throws InterruptedException {
        // write your code here
        // 从信号量获取一个许可
        printWorld.acquire();
        // 打印 printWorld 中的内容
        System.out.print("World");
    }
}

// solution 2 wait notify
public class Solution {
    // write your code here
    private int flag = 0;
    public synchronized void printHello() throws InterruptedException {
        // write your code here
        // 使用 while 而不用if 防止虚假唤醒
        while (flag != 0) { // 只有当 flag = 0 才能打印Hello 其他情况都会等待
            wait();
        }
        flag = 1;
        System.out.print("Hello");
        notify();
    }

    public synchronized void printWorld() throws InterruptedException {
        // write your code here
        while (flag != 1) { // 只有当flag=1 才能打印World 其他情况都会等待
            wait();
        }
        System.out.print("World");
        notify();
    }
}

/**
# Semaphore
Semaphore（信号量）是用来控制同时访问特定资源的线程数量，通过协调各个线程以保证合理地使用公共资源。
Semaphore 通过使用计数器来控制对共享资源的访问。 如果计数器大于 0，则允许访问。 如果为 0，则拒绝访问。 
计数器所计数的是允许访问共享资源的许可。 因此，要访问资源，必须从信号量中授予线程许可。
主要方法：

void acquire()：从信号量获取一个许可，如果无可用许可前将一直阻塞等待，
void acquire(int permits)：获取指定数目的许可，如果无可用许可前也将会一直阻塞等待
boolean tryAcquire()：从信号量尝试获取一个许可，如果无可用许可，直接返回 false，不会阻塞
boolean tryAcquire(int permits)： 尝试获取指定数目的许可，如果无可用许可直接返回 false
boolean tryAcquire(int permits, long timeout, TimeUnit unit)：在指定的时间内尝试从信号量中获取许可，如果在指定的时间内获取成功，返回 true，否则返回 false
void release()：释放一个许可，别忘了在 finally 中使用，注意：多次调用该方法，会使信号量的许可数增加，达到动态扩展的效果，如：初始 permits 为 1，调用了两次 release，最大许可会改变为 2
int availablePermits()：获取当前信号量可用的许可
Semaphore构造函数：
public Semaphore(int permits) {
     // permits 初始许可数，也就是最大访问线程数
     sync = new NonfairSync(permits);
 }

public Semaphore(int permits, boolean fair) {
    // fair 当设置为 false 时，创建的信号量为非公平锁；当设置为 true 时，信号量是公平锁  
    sync = fair ? new FairSync(permits) : new NonfairSync(permits);
}
示例：
Semaphore 登录限流

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Semaphore;

public class Main {
    public static void main(String[] args) {

        // 允许最大的登录数
        int slots = 10;
        ExecutorService executorService = Executors.newFixedThreadPool(slots);
        LoginQueueUsingSemaphore loginQueue = new LoginQueueUsingSemaphore(slots);
        // 线程池模拟登录
        for (int i = 1; i <= slots; i++) {
            final int num = i;
            executorService.execute(() -> {
                if (loginQueue.tryLogin()) {
                    System.out.println("用户:" + num + "登录成功！");
                } else {
                    System.out.println("用户:" + num + "登录失败！");
                }
            });
        }
        executorService.shutdown();


        System.out.println("当前可用许可证数：" + loginQueue.availableSlots());

        // 此时已经登录了 10 个用户，再次登录的时候会返回 false
        if (loginQueue.tryLogin()) {
            System.out.println("登录成功！");
        } else {
            System.out.println("系统登录用户已满，登录失败！");
        }
        // 有用户退出登录
        loginQueue.logout();

        // 再次登录
        if (loginQueue.tryLogin()) {
            System.out.println("登录成功！");
        } else {
            System.out.println("系统登录用户已满，登录失败！");
        }

    }
}

class LoginQueueUsingSemaphore {

    private Semaphore semaphore;


    public LoginQueueUsingSemaphore(int slotLimit) {
        semaphore = new Semaphore(slotLimit);
    }

    boolean tryLogin() {
        //获取一个凭证
        return semaphore.tryAcquire();
    }

    void logout() {
        semaphore.release();
    }

    int availableSlots() {
        return semaphore.availablePermits();
    }
}
以上实例编译运行结果如下：

用户:1登录成功！
用户:3登录成功！
用户:2登录成功！
用户:5登录成功！
用户:4登录成功！
用户:6登录成功！
用户:7登录成功！
用户:8登录成功！
用户:9登录成功！
用户:10登录成功！
当前可用许可证数：0
系统登录用户已满，登录失败！
登录成功！ 
示例：
获取许可和释放许可

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Semaphore;


public class SemaPhoreTest {
    private Semaphore semaphore = new Semaphore(3);

    public static void main(String[] args) {
        SemaPhoreTest semaPhoreTest = new SemaPhoreTest();

        // 同步队列线程
        ExecutorService executorService = Executors.newCachedThreadPool();
        executorService.submit(semaPhoreTest.new TaskThread(1));
        executorService.submit(semaPhoreTest.new TaskThread(2));
        executorService.submit(semaPhoreTest.new TaskThread(3));
        executorService.submit(semaPhoreTest.new TaskThread(4));
        executorService.submit(semaPhoreTest.new TaskThread(5));
        executorService.submit(semaPhoreTest.new TaskThread(6));
        executorService.submit(semaPhoreTest.new TaskThread(7));
        executorService.shutdown();
    }

    class TaskThread implements Runnable {
        private int id;

        public TaskThread(int id) {
            this.id = id;
        }

        @Override
        public void run() {
            try {
                // 获取许可
                semaphore.acquire();
                System.out.println("Thread " + id + " is working");
                // 线程休眠
                Thread.sleep(2000);
                // 释放许可
                semaphore.release();
                System.out.println("Thread " + id + " is over");
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
}
以上实例编译运行结果如下：

1
2
3
4
5
6
7
8
9
10
11
12
13
14
Thread 1 is working
Thread 3 is working
Thread 2 is working
Thread 4 is working
Thread 1 is over
Thread 3 is over
Thread 2 is over
Thread 5 is working
Thread 6 is working
Thread 7 is working
Thread 6 is over
Thread 5 is over
Thread 4 is over
Thread 7 is over 
 */