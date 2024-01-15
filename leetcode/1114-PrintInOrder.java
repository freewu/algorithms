package leetcode;

// 1114. Print in Order
// 并发问题
// 并发问题来自并发计算的场景，该场景下，程序在多线程（或多进程）中 同时 执行。
// 同时进行并不是完全指进程或线程在不同的物理 CPU 上独立运行，更多情况下，是在一个物理 CPU 上交替执行多个线程或进程。
// 并发既可在线程中，也可在进程中。

// 并发主要为多任务情况设计。但如果应用不当，可能会引发一些漏洞。按照情况不同，可以分为三种：
// 
//      竞态条件：由于多进程之间的竞争执行，导致程序未按照期望的顺序输出。
//      死锁：并发程序等待一些必要资源，导致没有程序可以执行。
//      资源不足：进程被永久剥夺了运行所需的资源。

//

class Foo {

    // 首先初始化共享变量 firstJobDone 和 secondJobDone，初始值表示所有方法未执行
    private AtomicInteger firstJobDone = new AtomicInteger(0);
    private AtomicInteger secondJobDone = new AtomicInteger(0);


    public Foo() {
        
    }

    public void first(Runnable printFirst) throws InterruptedException {
        // printFirst.run() outputs "first". Do not change or remove this line.
        printFirst.run();
        // 方法 first() 没有依赖关系，可以直接执行。在方法最后更新变量 firstJobDone 表示该方法执行完成。
        firstJobDone.incrementAndGet();
    }

    public void second(Runnable printSecond) throws InterruptedException {
        // 检查 firstJobDone 的状态。如果未更新则进入等待状态，否则执行方法 second()。
        while (firstJobDone.get() != 1) {
            // waiting for the first job to be done.
        }
        // printSecond.run() outputs "second". Do not change or remove this line.
        printSecond.run();
        // mark the second as done, by increasing its count.

        // 在方法末尾，更新变量 secondJobDone 表示方法 second() 执行完成
        secondJobDone.incrementAndGet();
    }

    public void third(Runnable printThird) throws InterruptedException {
        // 检查 secondJobDone 的状态。与方法 second() 类似，执行 third() 之前，需要先等待 secondJobDone 的状态
        while (secondJobDone.get() != 1) {
            // waiting for the second job to be done.
        }
        // printThird.run() outputs "third". Do not change or remove this line.
        printThird.run();
    }
}

// CountDownLatch
class Foo1 {
    private CountDownLatch count1;
    private CountDownLatch count2;
    public Foo1() {
        count1 = new CountDownLatch(1);
        count2 = new CountDownLatch(1);
    }

    public void first(Runnable printFirst) throws InterruptedException {
        // printFirst.run() outputs "first". Do not change or remove this line.
        printFirst.run();
        count1.countDown();
    }

    public void second(Runnable printSecond) throws InterruptedException {
        count1.await();
        // printSecond.run() outputs "second". Do not change or remove this line.
        printSecond.run();
        count2.countDown();
    }

    public void third(Runnable printThird) throws InterruptedException {
        count2.await();
        // printThird.run() outputs "third". Do not change or remove this line.
        printThird.run();
    }
}

// use Semaphore
class Foo2 {
    private Semaphore two = new Semaphore(0);
    private Semaphore three = new Semaphore(0);

    public Foo2() {

    }

    public void first(Runnable printFirst) throws InterruptedException {

        // printFirst.run() outputs "first". Do not change or remove this line.
        printFirst.run();
        two.release();
    }

    public void second(Runnable printSecond) throws InterruptedException {
        two.acquire();
        // printSecond.run() outputs "second". Do not change or remove this line.
        printSecond.run();
        three.release();
    }

    public void third(Runnable printThird) throws InterruptedException {
        three.acquire();
        // printThird.run() outputs "third". Do not change or remove this line.
        printThird.run();
    }
}

// 同步原语 synchronized + volatile
public class Foo3 {

    private volatile int flag = 1;
    private final Object object = new Object();

    public Foo3() {

    }

    public void first(Runnable printFirst) throws InterruptedException {
        synchronized (object) {
            while (flag != 1) object.wait();
            printFirst.run();
            flag = 2;
            object.notifyAll();
        }
    }

    public void second(Runnable printSecond) throws InterruptedException {
        synchronized (object) {
            while (flag != 2) object.wait();
            printSecond.run();
            flag = 3;
            object.notifyAll();
        }
    }

    public void third(Runnable printThird) throws InterruptedException {
        synchronized (object) {
            while (flag != 3) object.wait();
        }
        printThird.run();
    }
}