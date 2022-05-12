-- 2760 · Understanding SQL Procedures (I)
-- Description
-- We have provided a SQL procedure 'getTeachers', 
-- please write SQL statement to call this procedure to see its content.

-- Example
-- Input

-- teachers：
-- id	name	email	age	country
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 3	'Western Venom'	'western.venom@163.com'	18	'USA'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'

-- Return：
-- id	name	email	age	country
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 3	'Western Venom'	'western.venom@163.com'	18	'USA'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'

-- Write your SQL here --
CALL getTeachers();

-- # 创建存储过程（CREATE PROCEDURE）
-- MySQL存储过程是一些 SQL 语句的集合，比如有的时候我们可能需要一大串的 SQL 语句，或者说在编写 SQL 语句的过程中还需要设置一些变量的值，这个时候我们就完全有必要编写一个存储过程。
-- 基本语法
-- 可以使用 CREATE PROCEDURE 语句创建存储过程。
-- 语法格式如下：
-- CREATE PROCEDURE <过程名> ( [过程参数[,…] ] ) <过程体>
-- [过程参数[,…] ] 格式
-- [ IN | OUT | INOUT ] <参数名> <类型>

-- 语法说明如下：
-- 1) 过程名
--      存储过程的名称，默认在当前数据库中创建。
--      若需要在特定数据库中创建存储过程，则要在名称前面加上数据库的名称，即 db_name.sp_name。
--      需要注意的是，名称应当尽量避免选取与 MySQL 内置函数相同的名称，否则会发生错误。
-- 2) 过程参数
--      存储过程的参数列表。其中，<参数名>为参数名，<类型>为参数的类型（可以是任何有效的 MySQL 数据类型）。
--      当有多个参数时，参数列表中彼此间用逗号分隔。存储过程可以没有参数（此时存储过程的名称后仍需加上一对括号），也可以有 1 个或多个参数。
--      MySQL 存储过程支持三种类型的参数，即输入参数、输出参数和输入/输出参数，分别用 IN、OUT 和 INOUT 三个关键字标识。
--      其中，输入参数可以传递给一个存储过程，输出参数用于存储过程需要返回一个操作结果的情形，而输入/输出参数既可以充当输入参数也可以充当输出参数。
--      需要注意的是，参数的取名不要与数据表的列名相同，否则尽管不会返回出错信息，但是存储过程的 SQL 语句会将参数名看作列名，从而引发不可预知的结果。
-- 3) 过程体
--      存储过程的主体部分，也称为存储过程体，包含在过程调用的时候必须执行的 SQL 语句。
--      这个部分以关键字 BEGIN 开始，以关键字 END 结束。若存储过程体中只有一条 SQL 语句，则可以省略 BEGIN-END 标志。
--
-- 在存储过程的创建中，经常会用到一个十分重要的 MySQL 命令，即 DELIMITER 命令，
-- 特别是对于通过命令行的方式来操作 MySQL 数据库的使用者，更是要学会使用该命令。
-- 在 MySQL 中，服务器处理 SQL 语句默认是以分号作为语句结束标志的。
-- 然而，在创建存储过程时，存储过程体可能包含有多条 SQL 语句，这些 SQL 语句如果仍以分号作为语句结束符，
-- 那么 MySQL 服务器在处理时会以遇到的第一条 SQL 语句结尾处的分号作为整个程序的结束符，而不再去处理存储过程体中后面的 SQL 语句，这样显然不行。
-- 为解决这个问题，通常可使用 DELIMITER 命令将结束命令修改为其他字符。
-- 语法格式如下：

-- DELIMITER $$

-- 语法说明如下：
-- $$ 是用户定义的结束符，通常这个符号可以是一些特殊的符号，如两个 ? 或两个 ￥ 等。
-- 当使用 DELIMITER 命令时，应该避免使用反斜杠“\”字符，因为它是 MySQL 的转义字符。
-- 在 MySQL 命令行客户端输入如下SQL语句。

-- mysql > DELIMITER ??
-- 成功执行这条 SQL 语句后，任何命令、语句或程序的结束标志就换为两个问号“??”了。
-- 若希望换回默认的分号“;”作为结束标志，则在 MySQL 命令行客户端输入下列语句即可：
-- mysql > DELIMITER ;
-- 注意：DELIMITER 和分号“;”之间一定要有一个空格。
-- 在创建存储过程时，必须具有 CREATE ROUTINE 权限

-- # 查看数据库中存在哪些存储过程
--      SHOW PROCEDURE STATUS

-- # 查看某个存储过程的具体信息
--      SHOW CREATE PROCEDURE <存储过程名>。

-- # 创建不带参数的存储过程
-- mysql> DELIMITER ??
-- mysql> CREATE PROCEDURE ShowScore()
-- mysql> BEGIN
-- mysql>       SELECT * FROM tb_scores;
-- mysql> END ?? --  存储过程结束
-- mysql> DELIMITER ; 把 结束符定义重设成 ;
-- mysql> -- 调用存储过程
-- mysql> CALL ShowScore();

-- # 创建带参数的存储过程
-- mysql> DELIMITER ??
-- mysql> CREATE PROCEDURE GetScoreByStudent(IN pname VARCHAR(30))
-- mysql> BEGIN
-- mysql>       SELECT * FROM tb_scores WHERE name=pname;
-- mysql> END ?? --  存储过程结束
-- mysql> DELIMITER ; 把 结束符定义重设成 ;
-- mysql> -- 调用存储过程
-- mysql> CALL GetScoreByStudent('bluefrog');

-- # MySQL修改存储过程（ALTER PROCEDURE）
--  在实际开发过程中，业务需求修改的情况时有发生，这样，不可避免的需要修改 MySQL 中存储过程的特征 。
--  可以使用 ALTER PROCEDURE 语句修改存储过程的某些特征。
--  语法格式如下：
--      ALTER PROCEDURE <过程名> [ <特征> … ]
--
--  这个语法用于修改存储过程的某些特征，如要修改存储过程的内容，可以先删除该存储过程，再重新创建。
--  修改存储过程的内容可以通过删除原存储过程，再以相同的命名创建新的存储过程。

-- # 删除存储过程（DROP PROCEDURE）
-- 存储过程被创建后，保存在数据库服务器上，直至被删除。可以使用 DROP PROCEDURE 语句删除数据库中已创建的存储过程。
-- 语法格式如下：
--      DROP { PROCEDURE | FUNCTION } [ IF EXISTS ] <过程名>
-- 语法说明如下：
--  1) 过程名
--            指定要删除的存储过程的名称。
--  2) IF EXISTS
--          指定这个关键字，用于防止因删除不存在的存储过程而引发的错误。
-- 
--  注意：存储过程名称后面没有参数列表，也没有括号，在删除之前，必须确认该存储过程没有任何依赖关系，否则会导致其他与之关联的存储过程无法运行。
--  DROP PROCEDURE GetScores