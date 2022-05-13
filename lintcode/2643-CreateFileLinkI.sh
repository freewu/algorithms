# 2643 · Create file link (I)
# Description
# When we need to use the same file in different directories, 
# we don’t need to put a file that must be the same in each directory that we need. 
# We just need to put the file in a fixed directory, and then put it in other directories. 
# You can use the ln command to link it under the directory under the directory,
# so you don't need to repeatedly occupy the disk space.
# Use the Linux command to create a link to the file temp.txt with ln. 
# You need to complete the exercise through the following four steps:

# Step 1 Use ls -l to get the detailed information of the files in the current directory.
ls -l

# Step 2 Create a soft link of temp.txt in the same directory and name it s-temp.txt.
ln -s temp.txt s-temp.txt

# Step 3 Use ls -l to get the detailed information of the files in the current directory.
ls -l

# Step 4 Create a soft link of temp2.txt in the same directory and name it s-temp.txt, because it already exists and needs to be overwritten.tten.
ln -sf temp2.txt s-temp.txt

# ln  是linux系统中一个非常重要命令，英文全称是“link”，即链接的意思，它的功能是为某一个文件在另外一个位置建立一个同步的链接。 一种是hard link，又称为硬链接；另一种是symbolic link，又称为符号链接。
# 通俗一点理解，可以把硬链接当成源文件的副本，他和源文件一样的大小，但是事实上却不占任何空间。符号链接可以理解为类似windows一样的快捷方式。

# 符号链接 ：

# 1. 符号链接以路径的形式存在，类似于Windows操作系统中的快捷方式。
# 2. 符号链接可以跨文件系统 ，硬链接不可以。
# 3. 符号链接可以对一个不存在的文件名进行链接，硬链接不可以。
# 4. 符号链接可以对目录进行链接，硬链接不可以。

# 硬链接：

# 1.硬链接以文件副本的形式存在，但不占用实际空间。
# 2. 硬链接不允许给目录创建硬链接。
# 3.硬链接只有在同一个文件系统中才能创建。

# 语法格式： ln [参数] [源文件或目录] [目标文件或目录]

# 常用参数：
# -b	为每个已存在的目标文件创建备份文件
# -d	此选项允许“root”用户建立目录的硬链接
# -f	强制创建链接，即使目标文件已经存在
# -n    把指向目录的符号链接视为一个普通文件
# -i	交互模式，若目标文件已经存在，则提示用户确认进行覆盖
# -s	对源文件建立符号链接，而非硬链接
# -v	详细信息模式，输出指令的详细执行过程
