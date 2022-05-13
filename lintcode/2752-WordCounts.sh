# 2752 · Word counts
# Description
# The wc command is used to calculate the word count, so complete the following 3 steps in order.

# Step 1 use the cat command to view the contents of the file file
cat file

# Step 2 use the wc command to see the number of lines, words, and bytes in the file file
wc file 

# Step 3 use the wc command to see the number of lines and bytes in the file file
wc -lc file

# Linux wc命令用于计算字数。
# 利用wc指令我们可以计算文件的Byte数、字数、或是列数，若不指定文件名称、或是所给予的文件名为"-"，则wc指令会从标准输入设备读取数据。

# 语法
# wc [-clw][--help][--version][文件...]
# 参数：
# -c或--bytes或--chars 只显示Bytes数。
# -l或--lines 显示行数。
# -w或--words 只显示字数。
# --help 在线帮助。
# --version 显示版本信息。
# 实例
# 在默认的情况下，wc将计算指定文件的行数、字数，以及字节数。使用的命令为：

# wc testfile 
# 先查看testfile文件的内容，可以看到：

# $ cat testfile  
# Linux networks are becoming more and more common, but scurity is often an overlooked  
# issue. Unfortunately, in today’s environment all networks are potential hacker targets,  
# fro0m tp-secret military research networks to small home LANs.  
# Linux Network Securty focuses on securing Linux in a networked environment, where the  
# security of the entire network needs to be considered rather than just isolated machines.  
# It uses a mix of theory and practicl techniques to teach administrators how to install and  
# use security applications, as well as how the applcations work and why they are necesary. 
# 使用 wc统计，结果如下：
# $ wc testfile           # testfile文件的统计信息  
# 3 92 598 testfile       # testfile文件的行数为3、单词数92、字节数598 
# 其中，3 个数字分别表示testfile文件的行数、单词数，以及该文件的字节数。

# 如果想同时统计多个文件的信息，例如同时统计testfile、testfile_1、testfile_2，可使用如下命令：

# wc testfile testfile_1 testfile_2   #统计三个文件的信息 
# 输出结果如下：

# $ wc testfile testfile_1 testfile_2  #统计三个文件的信息  
# 3 92 598 testfile                    #第一个文件行数为3、单词数92、字节数598  
# 9 18 78 testfile_1                   #第二个文件的行数为9、单词数18、字节数78  
# 3 6 32 testfile_2                    #第三个文件的行数为3、单词数6、字节数32  
# 15 116 708 总用量                    #三个文件总共的行数为15、单词数116、字节数708 