/**
204 · Singleton

# Description
Singleton is a most widely used design pattern. 
If a class has and only has one instance at every moment, we call this design as singleton. 
For example, for class Mouse (not a animal mouse), we should design it in singleton.
You job is to implement a getInstance method for given class, 
return the same instance of this class every time you call this method.


Example

In Java:

	A a = A.getInstance();
	A b = A.getInstance();

a should equal to b.

# Challenge
    If we call getInstance concurrently, can you make sure your code could run correctly?
 */
class Solution {
    private static Solution instance = null;
    /**
     * @return: The same instance of this class every time
     */
    public static Solution getInstance() {
        // write your code here
        if(instance == null) instance = new Solution();
        return instance;
    }

    public static void main(String[] args) {
        Solution a = Solution.getInstance();
	    Solution b = Solution.getInstance();

        System.out.println(a);
        System.out.println(b);
        System.out.println(a.equals(b)); // true
    }
};

// V1:双重检查+synchronized锁懒汉版
public class Solution{
    
    // 1. 私有单例对象，禁止通过 类名.属性 访问
    // 2. 将其声明为静态成员，使得在静态方法中得以访问
    // 3. 使用volatile关键字，消除指令重排序的影响
    private static volatile Solution instance;
    
    // 1. 私有构造函数
    private Soltion(){
        
    }
    
    // 静态方法，返回单例对象。双重检查+synchroinzed锁，保证多线程下instance对象唯一
    public static Solution getInstance() {
        if(instance == null) {
            synchronized( Solution.class ) {
                if(instance == null) {
                    // 多线程并发访问，只会有一个线程初始化成功
                    instance = new Soltion();
                }
            }
        }
        return instance;    
    }
}

// V2: 静态内部类版
public class Solution{
    
    static class InnerClass{
        private static Solution instance = new Sotluion();
    }
    
    public static Solution getInstance(){
        return Solution.instance;
    }
    
}

// V3: 枚举类版
public enum Solution{
    INSTANCE;
    
    public static Solutin getInstance() {
        return Solution.INSTANCE;
    }
}