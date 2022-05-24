# 2273 · Sequential Output (Python version)
# Description
# Please obtain a positive integer n from the standard input stream (console). 
# It is required to output all integers in the range [1,n] through the print statement, and each number must be output in a new line.

# 100 1≤ n ≤100

# Example
# The evaluation opportunity executes your code by executing the command python main.py, and enters n as standard input from the console.

# Example 1
# When n = 3, the printed result of the program execution is:
# 1
# 2
# 3

# Example 2
# When n = 5, the printed result of program execution is:
# 1
# 2
# 3
# 4
# 5

# write your code here
# read data from console
n = int(input())

# output the answer to the console according to the requirements of the question
# i = 1
# while (n >= i):
#     print(i)
#     i += 1

# 使用循环变量在 for 循环中进行输出
for i in range(1, n + 1):
    print(i)