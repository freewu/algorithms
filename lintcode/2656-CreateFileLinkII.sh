# 2656 · Create file link (II)
# Description
# Use the Linux command line to create a link between the temp.txt file and the tmp folder with ln. 
# You need to complete the exercise through the following four steps:

# Step 1 Use ls -l to get the detailed information of the files in the current directory.
ls -l 

# Step 2 Create a hard link of temp.txt in the same directory and name it h-temp.txt.
ln temp.txt h-temp.txt.

# Step 3 Use ls -l to get the detailed information of the files in the current directory.
ls -l 

# Step 4 Create a soft link of the tmp folder under the create-link-here folder and name it s-tmp
ln -s /nobody/tmp /nobody/create-link-here/s-tmp

#  Linux允许同一个文件可以有好几个不同的名字，而它们共享一个数据。一个文件发生改变，其他的文件也都会发生改变。这就是Hard Link。
#  用命令行创建一个链接默认即为硬链接。命令如下：
#       ln  {source}  {link}
#  下面是一个例子：
#       ln foo bar
# 即为foo这个文件创建一个硬链接，新名字叫做bar。使用命令查看它们的编号一样：
#       ls -i foo
#       ls -i bar