# 2623 · Find eligible files in a directory (i)
# Description
# Use the find command as required to find files that match the requirements in the current folder,
# completing the following three steps.

# Step 1 query all files and directories in the current directory and its subdirectories and list
find 

# Step 2 search for all files in the current directory and its subdirectories and list them
find -type f

# Step 3 search for and list all the directories in the current directory and its subdirectories
find -type d


# Linux find 命令用来在指定目录下查找文件。任何位于参数之前的字符串都将被视为欲查找的目录名。
# 如果使用该命令时，不设置任何参数，则 find 命令将在当前目录下查找子目录与文件。并且将查找到的子目录和文件全部进行显示。

# 语法
# find  path -option  [-print][-exec -ok command] {} \;
# 参数说明 :

# find 根据下列规则判断 path 和 expression，在命令列上第一个 - ( ) , ! 之前的部份为 path，之后的是 expression。如果 path 是空字串则使用目前路径，如果 expression 是空字串则使用 -print 为预设 expression。

# expression 中可使用的选项有二三十个之多，在此只介绍最常用的部份。
# -mount, -xdev : 只检查和指定目录在同一个文件系统下的文件，避免列出其它文件系统中的文件
# -amin n : 在过去 n 分钟内被读取过
# -anewer file : 比文件 file 更晚被读取过的文件
# -atime n : 在过去 n 天内被读取过的文件
# -cmin n : 在过去 n 分钟内被修改过
# -cnewer file :比文件 file 更新的文件
# -ctime n : 在过去 n 天内创建的文件
# -mtime n : 在过去 n 天内修改过的文件
# -empty : 空的文件-gid n or -group name : gid 是 n 或是 group 名称是 name
# -ipath p, -path p : 路径名称符合 p 的文件，ipath 会忽略大小写
# -name name, -iname name : 文件名称符合 name 的文件。iname 会忽略大小写
# -size n : 文件大小 是 n 单位，b 代表 512 位元组的区块，c 表示字元数，k 表示 kilo bytes，w 是二个位元组。
# -type c : 文件类型是 c 的文件。
#       d: 目录
#       c: 字型装置文件
#       b: 区块装置文件
#       p: 具名贮列
#       f: 一般文件
#       l: 符号连结
#       s: socket
# -pid n : process id 是 n 的文件
# 你可以使用 ( ) 将运算式分隔，并使用下列运算。
#   exp1 -and exp2
#   ! expr
#   -not expr
#   exp1 -or exp2
#   exp1, exp2



# 将当前目录及其子目录下所有文件后缀为 .c 的文件列出来:

    find . -name "*.c"

# 将当前目录及其子目录中的所有文件列出：

    find . -type f

# 将当前目录及其子目录下所有最近 20 天内更新过的文件列出:

    find . -ctime  20

# 查找 /var/log 目录中更改时间在 7 日以前的普通文件，并在删除之前询问它们：

    find /var/log -type f -mtime +7 -ok rm {} \;

# 查找当前目录中文件属主具有读、写权限，并且文件所属组的用户和其他用户具有读权限的文件：

    find . -type f -perm 644 -exec ls -l {} \;

# 查找系统中所有文件长度为 0 的普通文件，并列出它们的完整路径：

    find / -type f -size 0 -exec ls -l {} \;

# 查找系统中所有文件长度不为 0 的普通文件，并列出它们的完整路径：

    find / -type f ! -size 0 -exec ls -l {} \;