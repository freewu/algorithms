# 2757 · Setting up the shell
# Description
# The set command is commonly used to set shell options, complete the following steps in order.

# Step 1 uncomment the wildcard character
set -f # 取消使用通配符。

# Step 2 Cancel the set parameter you just used
set +f # 取消某个set曾启动的参数。

# Step 3 View all the file names starting with f in the current folder
ls | grep f

# Linux set命令用于设置shell。
# set指令能设置所使用shell的执行方式，可依照不同的需求来做设置。
# 语法
#   set [+-abCdefhHklmnpPtuvx]
# 参数说明：

# -a 　标示已修改的变量，以供输出至环境变量。
# -b 　使被中止的后台程序立刻回报执行状态。
# -C 　转向所产生的文件无法覆盖已存在的文件。
# -d 　Shell预设会用杂凑表记忆使用过的指令，以加速指令的执行。使用-d参数可取消。
# -e 　若指令传回值不等于0，则立即退出shell。
# -f　 取消使用通配符。
# -h 　自动记录函数的所在位置。
# -H   Shell 　可利用"!"加<指令编号>的方式来执行history中记录的指令。
# -k 　指令所给的参数都会被视为此指令的环境变量。
# -l 　记录for循环的变量名称。
# -m 　使用监视模式。
# -n 　只读取指令，而不实际执行。
# -p 　启动优先顺序模式。
# -P 　启动-P参数后，执行指令时，会以实际的文件或目录来取代符号连接。
# -t 　执行完随后的指令，即退出shell。
# -u 　当执行时使用到未定义过的变量，则显示错误信息。
# -v 　显示shell所读取的输入值。
# -x 　执行指令后，会先显示该指令及所下的参数。
# +<参数> 　取消某个set曾启动的参数。

# 显示环境变量
# # set
# BASH=/bin/bash
# BASH_ARGC=()
# BASH_ARGV=()
# BASH_LINENO=()
# BASH_SOURCE=()
# BASH_VERSINFO=([0]="3" [1]="00" [2]="15" [3]="1" [4]="release" [5]="i386-redhat-linux-gnu")
# BASH_VERSION='3.00.15(1)-release'
# COLORS=/etc/DIR_COLORS.xterm
# COLUMNS=99
# DIRSTACK=()
# EUID=0
# GROUPS=()
# G_BROKEN_FILENAMES=1
# HISTFILE=/root/.bash_history
# HISTFILESIZE=1000
# HISTSIZE=1000
# HOME=/root
# HOSTNAME=hnlinux
# HOSTTYPE=i386
# IFS=$' '
# INPUTRC=/etc/inputrc
# KDEDIR=/usr
# LANG=zh_CN.GB2312
# LESSOPEN='|/usr/bin/lesspipe.sh %s'
# LINES=34
# L
# MAIL=/var/spool/mail/root
# MAILCHECK=60
# OLDPWD=/home/uptech
# OPTERR=1
# OPTIND=1
# OSTYPE=linux-gnu
# PATH=/usr/kerberos/sbin:/usr/kerberos/bin:/usr/local/sbin:/usr/local/bin:/sbin:/bin:/usr/sbin:/usr/bin:/usr/X11R6/bin:/root/bin:/opt/crosstools/gcc-3.4.6-glibc-2.3.6/bin
# PIPESTATUS=([0]="2")
# PPID=26005
# PROMPT_COMMAND='echo -ne "