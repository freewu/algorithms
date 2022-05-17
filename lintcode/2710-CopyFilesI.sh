# 2710 · Copy Files (i)
# Description
# In this topic, you can practice the comm command to compare two files that have been ordered by completing the following 4 steps.

# Step 1 use the cat command to view the contents of the file file1 1. use the cat command to view the contents of the file file2
cat file1

# Step 2 use the cat command to view the contents of file file2 1. use the comm command to view the contents of file file2
cat file2

# Step 3 use the comm command to compare the differences between files file1 and file2 by column
comm file1 file2

# Step 4 use the comm command to compare the differences between files file2 and file1 by column
comm file2 file1

# Linux comm 命令用于比较两个已排过序的文件。
# 这项指令会一列列地比较两个已排序文件的差异，并将其结果显示出来，
# 如果没有指定任何参数，则会把结果分成 3 列显示：
#       第 1 列仅是在第 1 个文件中出现过的列，
#       第 2 列是仅在第 2 个文件中出现过的列，
#       第 3 列则是在第 1 与第 2 个文件里都出现过的列。
# 若给予的文件名称为 - ，则 comm 指令会从标准输入设备读取数据。

# 语法
# comm [-123][--help][--version][第1个文件][第2个文件]
# 参数：
#   -1 不显示只在第 1 个文件里出现过的列。
#   -2 不显示只在第 2 个文件里出现过的列。
#   -3 不显示只在第 1 和第 2 个文件里出现过的列。
#   --help 在线帮助。
#   --version 显示版本信息。

# 实例
# aaa.txt 与 bbb.txt 的文件内容如下：
# [root@localhost text]# cat aaa.txt 
# aaa 
# bbb 
# ccc 
# ddd 
# eee 
# 111 
# 222
# [root@localhost text]# cat bbb.txt
# bbb 
# ccc 
# aaa 
# hhh 
# ttt 
# jjj
# 执行 comm 命令输出结果如下：

# [root@localhost text]# comm aaa.txt bbb.txt 
# aaa
#                 bbb
#                 ccc
#         aaa
# ddd
# eee
# 111
# 222
#         hhh
#         ttt
#         jjj
# 第一列  第二列  第三列