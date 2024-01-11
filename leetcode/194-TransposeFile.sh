# 194. Transpose File
# Given a text file file.txt, transpose its content.
# You may assume that each row has the same number of columns, and each field is separated by the ' ' character.

# Example:
# If file.txt has the following content:
#   name age
#   alice 21
#   ryan 30

# Output the following:
#   name alice ryan
#   age 21 30

# use shell
ROWS=()
while IFS=' '; read -a columns; do
    i=0
    for col in "${columns[@]}"; do
        if [ -z "${ROWS[$i]}" ]; then
            ROWS[$i]="$col"
        else
            ROWS[$i]+=" $col"
        fi
        ((i++))
    done
done <file.txt

for row in "${ROWS[@]}"; do
    echo "$row"
done

# use awk
awk '
{
    for (i = 1; i <= NF; i++) {
        if (FNR == 1) {
            t[i] = $i;
        } else {
            t[i] = t[i] " " $i
        }
    }
}
END {
    for (i = 1; t[i] != ""; i++) {
        print t[i]
    }
}
' file.txt

# FNR == 1 判断为第一行时  name age
# t[i] = t[i] " " $i 后面循环到的数据拼接到  t[i]
# for (i = 1; t[i] != ""; i++) 输出

# NF(number of field)：域的个数
# $NF ：最后一个Field(列)
# FS：输入字段分隔符， 默认为空白字符
# OFS：输出字段分隔符， 默认为空白字符
# RS：输入记录分隔符(输入换行符)， 指定输入时的换行符
# ORS：输出记录分隔符（输出换行符），输出时用指定符号代替换行符
# NF：number of Field，当前行的字段的个数(即当前行被分割成了几列)，字段数量
# NR：行号，当前处理的文本行的行号。
# FNR：各文件分别计数的行号
# FILENAME：当前文件名
# ARGC：命令行参数的个数
# ARGV：数组，保存的是命令行所给定的各参数

# best solution
COLUMNS=$(head -n 1 file.txt | wc -w)

for i in $(seq 1 $COLUMNS); do
    cut -d ' ' -f"$i" file.txt | paste -s -d' '  -
done
