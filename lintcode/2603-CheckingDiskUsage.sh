# 2603 · Checking disk usage
# Description
# To check the disk space usage of the file system using the df command, complete the following 3 steps in order.

# Step 1 check the disk space occupation of the file system
df

# Step 2 check the disk space usage of the file system, including all file systems with 0 Blocks
df -a

# Step 3 check the disk space usage of the file system and display the file system form
df -T

# 语法：
#       df [-ahHiklmPT][--block-size=<区块大小>][-t <文件系统类型>][-x <文件系统类型>][--help][--no-sync][--sync][--version][文件或设备]
#  显示磁盘的文件系统与使用情形。

# # 参数说明:

# -a --all   包含全部的文件系统
# [root@bluefrog-laptop ~]# df -a
# Filesystem     1K-blocks    Used Available Use% Mounted on
# /dev/xvda1      41282880 1808452  37377380   5% /
# proc                   0       0         0    - /proc
# sysfs                  0       0         0    - /sys
# devpts                 0       0         0    - /dev/pts
# tmpfs             509164       0    509164   0% /dev/shm
# none                   0       0         0    - /proc/sys/fs/binfmt_misc
# none                   0       0         0    - /proc/xen

# --block-size=<区块大小>   以指定的区块大小来显示区块数目。
# [root@bluefrog-laptop ~]# df --block-size=10240
# Filesystem     10K-blocks   Used Available Use% Mounted on
# /dev/xvda1        4128288 180846   3737738   5% /
# tmpfs               50917      0     50917   0% /dev/shm

# -h --human-readable   以可读性较高的方式来显示信息。
# [root@bluefrog-laptop ~]# df -h
# Filesystem      Size  Used Avail Use% Mounted on
# /dev/xvda1       40G  1.8G   36G   5% /
# tmpfs           498M     0  498M   0% /dev/shm

#  -H --si   与 -h 参数相同，但在计算时是以1000 Bytes为换算单位而非1024 Bytes。
# [root@bluefrog-laptop ~]# df -H
# Filesystem      Size  Used Avail Use% Mounted on
# /dev/xvda1       43G  1.9G   39G   5% /
# tmpfs           522M     0  522M   0% /dev/shm

# -i --inodes   显示inode的信息。
# [root@bluefrog-laptop ~]# df -i
# Filesystem      Inodes IUsed   IFree IUse% Mounted on
# /dev/xvda1     2621440 48349 2573091    2% /
# tmpfs           127291     1  127290    1% /dev/shm

#  -k --kilobytes   指定区块大小为1024字节。
# [root@bluefrog-laptop ~]# df -k
# Filesystem     1K-blocks    Used Available Use% Mounted on
# /dev/xvda1      41282880 1808456  37377376   5% /
# tmpfs             509164       0    509164   0% /dev/shm

# -l --local   仅显示本地端的文件系统。
# [root@bluefrog-laptop ~]# df -l
# Filesystem     1K-blocks    Used Available Use% Mounted on
# /dev/xvda1      41282880 1808456  37377376   5% /
# tmpfs             509164       0    509164   0% /dev/shm

# -m --megabytes   指定区块大小为1048576字节。
# [root@bluefrog-laptop ~]# df -m
# Filesystem     1M-blocks  Used Available Use% Mounted on
# /dev/xvda1         40316  1767     36502   5% /
# tmpfs                498     0       498   0% /dev/shm

#  --no-sync   在取得磁盘使用信息前，不要执行sync指令，此为预设值。
# [root@bluefrog-laptop ~]# df --no-sync
# Filesystem     1K-blocks    Used Available Use% Mounted on
# /dev/xvda1      41282880 1808460  37377372   5% /
# tmpfs             509164       0    509164   0% /dev/shm

# -P --portability   使用POSIX的输出格式。
# [root@bluefrog-laptop ~]# df -P
# Filesystem     1024-blocks    Used Available Capacity Mounted on
# /dev/xvda1        41282880 1808460  37377372       5% /
# tmpfs               509164       0    509164       0% /dev/shm

# --sync   在取得磁盘使用信息前，先执行sync指令。
# [root@bluefrog-laptop ~]# df --sync
# Filesystem     1K-blocks    Used Available Use% Mounted on
# /dev/xvda1      41282880 1808460  37377372   5% /
# tmpfs             509164       0    509164   0% /dev/shm

# -t<文件系统类型> --type=<文件系统类型>   仅显示指定文件系统类型的磁盘信息。
# [root@bluefrog-laptop ~]# df -t tmpfs
# Filesystem     1K-blocks  Used Available Use% Mounted on
# tmpfs             509164     0    509164   0% /dev/shm

#  -T --print-type   显示文件系统的类型。
# [root@bluefrog-laptop ~]# df -T
# Filesystem     Type  1K-blocks    Used Available Use% Mounted on
# /dev/xvda1     ext3   41282880 1808468  37377364   5% /
# tmpfs          tmpfs    509164       0    509164   0% /dev/shm

# -x<文件系统类型> --exclude-type=<文件系统类型>   不要显示指定文件系统类型的磁盘信息。
# [root@bluefrog-laptop ~]# df -x tmpfs
# Filesystem     1K-blocks    Used Available Use% Mounted on
# /dev/xvda1      41282880 1808468  37377364   5% /

#  --direct      show statistics for a file instead of mount point
# [root@bluefrog-laptop ~]# df --direct
# Filesystem     1K-blocks    Used Available Use% File
# /dev/xvda1      41282880 1808468  37377364   5% /
# tmpfs             509164       0    509164   0% /dev/shm

# --total       produce a grand total
# [root@bluefrog-laptop ~]# df --total
# Filesystem     1K-blocks    Used Available Use% Mounted on
# /dev/xvda1      41282880 1808468  37377364   5% /
# tmpfs             509164       0    509164   0% /dev/shm
# total           41792044 1808468  37886528   5%

# --help   显示帮助
# --version   显示版本信息