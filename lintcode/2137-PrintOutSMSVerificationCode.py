# 2137 · Print out SMS verification code
# # Description
# print is one of the most basic methods in Python, and it's probably the first method we'll encounter when learning a language

# When outputting a string, we can sometimes use the format method to output the variable along with it

# var = 
# f"var is {var}"
# "var is {}".format(var)
# "var is {v}".format(v=var)
# All three can output the same string

# Please write python code that prints "Hello, {name}! Your validation code is {SMS_verification_code}, 
# please keep it in secret." The content in {} needs to be entered.
# You need to refine the code at the specified location in main.py 
# to output a message containing the name and the SMS validation code.

# Do not change the code before the comment
# You only need to write the paragraph which print the SMS verification code
# Example
# We'll do this by importing the sys module and passing in name = zhangsan1 and SMS_verification_code = pZqccl by calling the methods in it.
# After executing main.py,your code should output
#   Hello, zhangsan1! Your validation code is pZqccl, please keep it in secret. 

import sys

name = str(sys.argv[1])
SMS_verification_code = str(sys.argv[2])

# print "Hello, name! Your validation code is SMS_verification_code
# please keep it in secret."
print("Hello, {name}! Your validation code is {code}, please keep it in secret. ".format(name = name,code = SMS_verification_code ))