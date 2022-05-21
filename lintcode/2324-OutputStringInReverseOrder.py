# 2324 · Output String In Reverse Order (Python version)
# Description
# Please obtain a string s from the standard input stream (console), 
# and use the print statement to output the string in reverse order.

# 1 ≤ len(s) ≤ 10^9
 

# Example
# The evaluation opportunity executes your code by executing the command python main.py, 
# and enters s as standard input from the console.

# Example 1
# When s = 12345, the printed result of the program execution is:
#   54321

# Example 2
# When s = abcd, the printed result of the program execution is:
#   dcba

# Example 3
# When s = 1000000, the printed result of the program execution is:
#   0000001

# write your code here
# read data from console
s = list(input())

# output the answer to the console according to the requirements of the question
s.reverse()
print(''.join(s))