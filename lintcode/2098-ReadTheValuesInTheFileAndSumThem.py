# 2098 · Read the values in the file and sum them
# Description
# Please write the relevant Python code in main.py, 
# read the contents of the file from the given path input_filepath, 
# calculate and sum them, and finally print the result of the calculation.

# Note that the data type when outputting is integer

# Example
# The evaluation opportunity executes your code by executing the command python solution.py {path}, 
# and passing the data in the file as a command line parameter. You can learn how the code runs in solution.py.

# Example 1
# When the input data is:
#   1 2
# The output data is:
#   3

# Example 2
# When the input data is:
#   3 4
# The output data is:
#   7

import sys

input_filepath = sys.argv[1]

# write your code here to read file content
# and print it to console

# 以只读权限打开文件
with open(input_filepath, 'r') as f:
    # 读取文件信息
    data = f.read()

# 使用 split() 以空格为分割，分别将切开的两个字符串赋值给 a, b
a, b = data.split(' ')

# 使用 int 将参数转为整型，求和后使用 print 打印
print(int(a) + int(b))
