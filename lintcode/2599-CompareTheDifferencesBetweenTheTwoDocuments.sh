# 2599 · Compare the differences between the two documents
# Description
# Use the diff command to compare the differences between the specified two files, completing the following 3 steps.

# Step 1 compare the difference between file1 and file2
diff file1 file2

# Step 2 compare the difference between file1 and file4, ignoring case
diff -i file1 file4

# Step 3 compare the difference between file1 and file3, ignoring changes in the amount of whitespaceers
diff -b file1 file3


# diff 命令用于比较文件的差异。
#   diff 以逐行的方式，比较文本文件的异同处。如果指定要比较目录，则 diff 会比较目录中相同文件名的文件，但不会比较其中子目录。
# # 语法
# diff [-abBcdefHilnNpPqrstTuvwy][-<行数>][-C <行数>][-D <巨集名称>][-I <字符或字符串>][-S <文件>][-W <宽度>][-x <文件或目录>][-X <文件>][--help][--left-column][--suppress-common-line][文件或目录1][文件或目录2]
# # 参数：
# -<行数> 　
#       指定要显示多少行的文本。此参数必须与-c或-u参数一并使用。
# -a --text 　
#       diff预设只会逐行比较文本文件。
# -b --ignore-space-change 　
#       不检查空格字符的不同。
# -B --ignore-blank-lines 　
#       不检查空白行。
# -c 　
#       显示全部内文，并标出不同之处。
# -C<行数> --context<行数> 　
#       与执行"-c-<行数>"指令相同。
# -d -minimal 　
#       使用不同的演算法，以较小的单位来做比较。
# -D<巨集名称> ifdef<巨集名称> 　
#       此参数的输出格式可用于前置处理器巨集。
# -e --ed 　
#       此参数的输出格式可用于ed的script文件。
# -f -forward-ed 　
#       输出的格式类似ed的script文件，但按照原来文件的顺序来显示不同处。
# -H --speed-large-files 　
#       比较大文件时，可加快速度。
# -I<字符或字符串> --ignore-matching-lines<字符或字符串> 　
#       若两个文件在某几行有所不同，而这几行同时都包含了选项中指定的字符或字符串，则不显示这两个文件的差异。
# -i --ignore-case 　
#       不检查大小写的不同。
# -l --paginate 　
#       将结果交由pr程序来分页。
# -n --rcs 　
#       将比较结果以RCS的格式来显示。
# -N --new-file 　
#       在比较目录时，若文件A仅出现在某个目录中，预设会显示：
#       Only in目录：文件A若使用-N参数，则diff会将文件A与一个空白的文件比较。
# -p 　
#       若比较的文件为C语言的程序码文件时，显示差异所在的函数名称。
# -P --unidirectional-new-file
#    　 与-N类似，但只有当第二个目录包含了一个第一个目录所没有的文件时，才会将这个文件与空白的文件做比较。
# -q --brief 　
#       仅显示有无差异，不显示详细的信息。
# -r --recursive 　
#       比较子目录中的文件。
# -s --report-identical-files 　
#       若没有发现任何差异，仍然显示信息。
# -S<文件> --starting-file<文件> 　
#       在比较目录时，从指定的文件开始比较。
# -t --expand-tabs 　
#       在输出时，将tab字符展开。
# -T --initial-tab 
#       在每行前面加上tab字符以便对齐。
# -u,-U<列数> --unified=<列数> 　
#       以合并的方式来显示文件内容的不同。
# -v --version 　
#       显示版本信息。
# -w --ignore-all-space 　
#       忽略全部的空格字符。
# -W<宽度> --width<宽度>
#       在使用-y参数时，指定栏宽。
# -x<文件名或目录> --exclude<文件名或目录> 
#       不比较选项中所指定的文件或目录。
# -X<文件> --exclude-from<文件> 　
#       可以将文件或目录类型存成文本文件，然后在=<文件>中指定此文本文件。
# -y --side-by-side 　
#       以并列的方式显示文件的异同之处。
# --help 　
#       显示帮助。
# --left-column 　
#       在使用-y参数时，若两个文件某一行内容相同，则仅在左侧的栏位显示该行内容。
# --suppress-common-lines 　
#       在使用-y参数时，仅显示不同之处。