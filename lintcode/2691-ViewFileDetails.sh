# 2691 · View file details
# Description
# In Linux, there always seems to be a clever way to accomplish tasks. 
# For any task, there are always multiple command line utilities to perform it in a better way. 
# The Linux stat command is a command line tool used to display detailed information about files. 
# You need to complete the exercise through the following three steps:

# Step 1 Use stat to view the contents of the file temp1.
stat temp1

# Step 2 View the information of the file system where the file temp1 is located.
stat -f temp1

# Step 3 View the information of all current files.
stat *

# Linux stat 命令用于显示 inode 内容。
# stat 以文字的格式来显示 inode 的内容。
# 语法
#       stat [文件或目录]
#
# # 查看 testfile 文件的inode内容内容，可以用以下命令：

#       stat testfile 
# 执行以上命令输出结果：

#   # stat testfile                # 输入命令
#       File: `testfile'
#       Size: 102             Blocks: 8          IO Block: 4096   regular file
#   Device: 807h/2055d      Inode: 1265161     Links: 1
#   Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
#   Access: 2014-08-13 14:07:20.000000000 +0800
#   Modify: 2014-08-13 14:07:07.000000000 +0800
#   Change: 2014-08-13 14:07:07.000000000 +0800

# ## stat命令可以查看的信息包括：
# File：显示文件名
# Size：显示文件大小
# Blocks：文件使用的数据块总数
# IO Block：IO块大小
# regular file：文件类型（常规文件）
# Device：设备编号
# Inode：Inode号
# Links：链接数
# Access：文件的权限
# Gid、Uid：文件所有权的Gid和Uid
# access time：表示我们最后一次访问（仅仅是访问，没有改动）文件的时间
# modify time：表示我们最后一次修改文件的时间
# change time：表示我们最后一次对文件属性改变的时间，包括权限，大小，属性等等
# Birth time : 文件创建时间，crtime，不过据查此属性linux已废弃，目前状态显示结果均为-

# # 查看文件状态信息
# [root@s145]~# stat test.sh
# File: ‘test.sh’
# Size: 72 Blocks: 8 IO Block: 4096 regular file
# Device: fd00h/64768d Inode: 101932951 Links: 1
# Access: (0744/-rwxr–r--) Uid: ( 0/ root) Gid: ( 0/ root)
# Context: unconfined_u:object_r:admin_home_t:s0
# Access: 2022-01-12 15:43:11.851607487 +0800
# Modify: 2020-01-12 15:43:02.520972106 +0800
# Change: 2022-01-12 15:43:02.520650817 +0800
# Birth: -

# # -f 查看文件所在文件系统状态信息
# [root@s145]~# stat -f test.sh
# File: “test.sh”
# ID: fd0000000000 Namelen: 255 Type: xfs
# Block size: 4096 Fundamental block size: 4096
# Blocks: Total: 9703804 Free: 8676628 Available: 8676628
# Inodes: Total: 19417088 Free: 19372526

# # -t 以简洁形式输出文件信息
# # 简洁形式主要是把每个值的列头去掉了，只显示值，三个时间按照纪元至今秒数显示。
# [root@s145]~# stat -t test.sh
# test.sh 72 8 81e4 0 0 fd00 101932951 1 0 0 1641973391 1578814982 1641973382 0 4096 unconfined_u:object_r:admin_home_t:s

# # -L 显示软连接原始文件状态信息


# # 获取文件inode
# [root@s145]~# stat -c %i test.sh
# 101932951

# # 获取文件所属用户或者用户ID
# [root@s145]~# stat -c %u test.sh
# 1000
# [root@s145]~# stat -c %U test.sh
# wuhs

# # 获取文件所属用户组或者用户组ID
# [root@s145]~# stat -c %G test.sh
# wuhs
# [root@s145]~# stat -c %g test.sh
# 1000

# # 获取文件磁盘挂载点
# [root@s145]~# stat -c %m test.sh
# /
# [root@s145]~# stat -c %m /home/wuhs
# /home

# # 获取文件访问时间
# [root@s145]~# stat -c %x test.sh
# 2022-01-12 15:43:11.851607487 +0800

# # 获取文件修改时间
# [root@s145]~# stat -c %y test.sh
# 2020-01-12 15:43:02.520972106 +0800

# # 获取文件更改时间
# [root@s145]~# stat -c %z test.sh
# 2022-01-12 16:02:18.374012017 +0800

# # 获取文件权限
# [root@s145]~# stat -c %A test.sh
# -rwxr–r--
# [root@s145]~# stat -c %a test.sh
# 744

# ## 使用语法
# # 用法：
# #    stat [OPTION]… FILE…

# ## 参数说明
# 参数	参数说明
# -L	支持符号连接；
# -f	显示文件系统状态而非文件状态；
# -c	按照指定格式输出；
# -t	以简洁方式输出信息；
# –help	显示指令的帮助信息；
# –version	显示指令的版本信息。
# ## 文件的有效格式序列
# 格式符	格式符说明
# %a	八进制中的访问权限
# %A	人类可读形式的访问权
# %b	分配的块数（请参阅%B）
# %B	%b报告的每个块的大小（字节）
# %C	SELinux安全上下文字符串
# %d	十进制设备编号
# %D	十六进制的设备编号
# %f	十六进制原始模式
# %F	文件类型
# %g	所有者的组ID
# %G	所有者的组名称
# %h	硬链接数
# %i	inode数
# %m	挂载点
# %n	文件名
# %N	带取消引用（如果是符号链接）的带引号的文件名
# %o	最佳I/O传输大小提示
# %s	总大小，以字节为单位
# %t	主要设备类型（十六进制），用于字符/块设备特殊文件
# %T	次要设备类型（十六进制），用于字符/块设备特殊文件
# %u	所有者的用户ID
# %U	所有者的用户名
# %w	文件出生时间，人类可读；-如果未知
# %W	文件生成时间，自纪元起的秒数；如果未知，则为0
# %x	最后一次访问的时间，人类可读
# %X	上次访问的时间，自纪元起的秒数
# %y	上次修改的时间，人类可读
# %Y	上次修改的时间，自纪元起的秒数
# %z	最后更改的时间，人类可读
# %Z	上次更改的时间，自新纪元起的秒数
# ## 文件系统的有效格式序列格式符	格式符说明
# %a	非超级用户可用的空闲块
# %b	文件系统中的数据块总数
# %c	文件系统中的文件节点总数
# %d	文件系统中的空闲文件节点
# %f	文件系统中的空闲块
# %i	十六进制文件系统ID
# %l	文件名的最大长度
# %n	文件名
# %s	块大小（用于更快的传输）
# %S	基本块大小（用于块计数）
# %t	十六进制文件系统类型
# %T	人类可读形式的文件系统类
