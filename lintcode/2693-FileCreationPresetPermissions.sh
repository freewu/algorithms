# 2693 · File creation preset permissions
# Description
# To combine the umask and ls commands, query the permission mask you preset when you created the file, 
# and create a new file to verify it, complete the following 3 steps in order.

# Step 1 use the umask command to query the permission mask of the file you created
umask

# Step 2 use the touch command to create a new file new_file 
touch new_file

# Step 3 use the ls command to query the new_file file for permission information
ls -l new_file

#  umask 是什么
# 当我们登录系统之后创建一个文件是会有一个默认权限的，那么这个权限是怎么来的呢？这就是umask干的事情。
# umask用于设置用户创建文件或者目录的默认权限，umask设置的是权限的“补码”，而我们常用chmod设置的是文件权限码。
# 一般在/etc/profile,HOME/.bashprofile或者HOME/.profile中设置umask值。

# 2 umask是用来做什么的
# 默认情况下的umask值是022(可以用umask命令查看），此时你建立的文件默认权限是644(6-0,6-2,6-2)，
# 建立的目录的默认 权限是755(7-0,7-2,7-2)，可以用ls -l验证一下,　现在应该知道umask的用途了吧，它是为了控制默认权限的。


# [root@bogon test]# id
# uid=0(root) gid=0(root) groups=0(root) context=unconfined_u:unconfined_r:unconfined_t:s0-s0:c0.c1023
# [root@bogon test]# umask
# 0022
# [root@bogon test]# touch a.txt
# [root@bogon test]# ls -l
# total 0
# -rw-r--r--. 1 root root 0 Jul 3 00:40 a.txt
# [root@bogon test]# mkdir b
# [root@bogon test]# ls -l
# total 0
# -rw-r--r--. 1 root root 0 Jul 3 00:40 a.txt
# drwxr-xr-x. 2 root root 6 Jul 3 00:41 b
# 从上面可以看到， root 的umask是022(第一个0 代表特殊权限位，这里先不考虑)， 创建的文件默认权限是644，创建的目录是755。

# 3 基本权限讲解
# 讲解umask的使用之前， 需要先讲解下文件的基本权限

# linux文件权限
#  	r	w	x
# 文件	 可以查看文件内容	 可以修改文件	 可以把文件启动为一个运行的程序
# 目录	 可以ls查看目录中的文件名	 可以在目录中创建或者删除文件（只有w权限没法创建，需要x配合）	 可以使用cd 进入这个目录ls-l显示目录内文件的元数据的信息
# 4 umask计算权限
# 对于文件和目录来说， 最大的权限其实都是777，但是执行权限对于文件来说，很可怕，而对目录来说执行权限是个基本权限。所以默认目录的最大权限是777，而文件的默认最大权限就是666。

# 对于root用户的umask=022这个来说，777权限二进制码就是（111）（111）（111），022权限二进制码为（000）（010）（010）。

# 所有权限二进制的1:代表有这个权限
# umask二进制1：代表要去掉这个权限，不管你原来有没有权限，你最终一定没有这个权限。
# umask二进制的0：代表我不关心对应位的权限，你原来有权限就有权限，没有就没有， 我不影响你。
# umask为002的文件默认权限计算方法

#  	所有者 r	所有者 w	所有者 x	所在组 r	所在组 w	所在组 x	其他 r	其他 w	其他 x 
# 所有权限777	1	1	1	1	1	1	1	1	1
# umask掩码002	0	0	0	0	1	0	0	1	0
# 计算后的值	1	1	1	1	0	1	1	0	1
# umask为002的目录默认权限计算方法
#  	所有者 r	所有者 w	所有者 x	所在组 r	所在组 w	所在组 x	其他 r	其他 w	其他 x 
# 所有权限666	1	1	0	1	1	0	1	1	0
# umask掩码002	0	0	0	0	1	0	0	1	0
# 计算后的值	1	1	0	1	0	0	1	0	0
# umask为023的目录默认权限计算方法
#  	所有者 r	所有者 w	所有者 x	所在组 r	所在组 w	所在组 x	其他 r	其他 w	其他 x 
# 所有权限777	1	1	1	1	1	1	1	1	1
# umask掩码023	0	0	0	0	1	0	0	1	1
# 计算后的值	1	1	1	1	0	1	1	0	0
# umask为023的文件默认权限计算方法
#  	所有者 r	所有者 w	所有者 x	所在组 r	所在组 w	所在组 x	其他 r	其他 w	其他 x 
# 所有权限666	1	1	0	1	1	0	1	1	0
# umask掩码023	0	0	0	0	1	0	0	1	1
# 计算后的值	1	1	0	1	0	0	1	0	0
# 上面就是一个umask的正常计算过程，但是这样实在是太麻烦了。我们使用如下的简单的方法快速计算。

# 对于目录，直接使用777-umask即可，就得到了最终结果。
# 对于文件，先使用666-umask。
# 如果对应位上为偶数：最终权限就是这个偶数值。
# 如果上面的对应为上有奇数，就对应位+1。
# 上面的这个方法计算是非常方便的， 为何得到奇数要+1呢。

# 文件的最大权限是666，都是偶数，你得到奇数，说明你的umask有奇数啊，读为4，写为2，都是偶数，说明你有执行权限的。

# 就按照上面的umask=023为例，在计算其他用户权限的时候6-3=3 ，6是读写，3是写和执行，其实应该是读写权限减去读权限的得到写权限的，相当于我们多减去了一个执行权限。所以结果加1。

# 5 umask的修改
# umask 的修改分2中， 临时修改的和永久修改的

# 临时修改：

# [root@bogon test]# umask 023
# [root@bogon test]# umask
# 0023
# [root@bogon test]#
# 永久修改：

# 可以编辑以下文件 添加umask=022。

# 交互式登陆的配置生效：

# /etc/profile < /etc/profile.d/*.sh < ~/.bash_profile < ~/.bashrc </etc/bashrc 【/etc/bashrc的配置最有效 可以覆盖前面的配置】

# 非交互登陆的配置生效：

# ~/.bashrc < /etc/bashrc  < /etc/profile.d/*.sh

# 6 常用umask

# [root@bogon test]# umask 002
# [root@bogon test]# umask
# 0002
# [root@bogon test]# umask 022
# [root@bogon test]# umask
# 0022