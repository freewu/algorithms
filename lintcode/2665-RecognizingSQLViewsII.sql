-- 2665 · Recognizing SQL Views (II)
-- # Description
-- Now we need to see all the table and view names in the current database and its type, write the SQL statement to do so

-- Example
-- The result should be returned after executing your SQL statement:

-- Tables_in_judge	Table_type
-- 'courses'	'BASE TABLE'
-- 'teachers'	'BASE TABLE'
-- 'v_courses'	'VIEW'
-- 'v_courses_teachers'	'VIEW'
-- 'v_teachers'	'VIEW'

-- SHOW TABLES 语法
-- SHOW [FULL] TABLES [FROM db_name] [LIKE 'pattern']
-- SHOW TABLES 列举了给定数据库中的非 TEMPORARY 表。您也可以使用 mysqlshow db_name 命令得到此清单。
-- 本命令也列举数据库中的其它视图。支持 FULL 修改符，这样 SHOW FULL TABLES 就可以显示第二个输出列。对于一个表，第二列的值为 BASE TABLE；对于一个视图，第二列的值为 VIEW。
-- 如果您对于一个表没有权限，则该表不会在来自 SHOW TABLES 或的 mysql show db_name 输出中显示。

SHOW FULL TABLES;