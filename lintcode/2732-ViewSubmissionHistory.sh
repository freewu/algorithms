# 2732 · View submission history
# The git log command can be used to view the history of commits in a number of different formats, 
# so complete the following steps in order.

# Step 1 go to the /nobody/my-repo directory and use the git log command to view the history of commits
cd /nobody/my-repo
git log

# Step 2 view the history of the last two commits
git log -2

# Step 3 show the differences introduced by each commit
git log -p

# Step 4 list under each commit all the files that have been modified, how many files have been modified, and which lines have been removed or added to the modified files
git log --stat

# Step 5 query the commit history for commits with the string test in their description
git log --grep test

# --oneline
#   使用 --oneline 参数，只显示提交的 SHA1 值和提交信息，SHA1 还是缩短显示前几位，一般为前七位，
#   前七位冲突的概率基本已经很小了，所以Git使用前七位代表整个40位复杂的SHA1值

# --graph
#   绘制一个 ASCII 图像来展示提交历史的分支结构，类似于一个树形结构

# --stat
#   使用--stat参数主要可以在git log 的基础上输出文件增删改的统计数据。
#       $ git log --stat
#       commit fa71c098e2912b69a1c82348d403b3260f2dc64e (HEAD -> temp_temp)
#       Author: zz203203zhang <zz203203zhang@gmail.com>
#       Date:   Wed Aug 12 17:19:05 2020 +0800
#           add txt file and dir           # commit信息
#       txt/a.txt | 1 +                   # 文件修改状态，添加或删除了多少行
#       1 file changed, 1 insertion(+)    # 统计变更文件数量

# -p
#   -p参数与--stat类似，不过-p参数更为详细，可以看到每个文件更为详细的修改内容，团队协作做项目中很有用途。
#   控制输出每个commit具体修改的内容，输出的形式以diff的形式给出。

# -<number>
#   将提交的显示次数限制
#   git log -3 // 显示近3条日志

# --author
#   --author用来过滤commit,限定输出给定的用户，这个有利于查找团队某个人的提交历史
#   git log --author="bluefrog"

# --after --before
#   限定指定日期范围的log，就是按照日期查找
#   git log --after '08-12-2020'

# --decoreate
#   该参数用来控制log输出时，显示对应commit所属的branch和tag信息
