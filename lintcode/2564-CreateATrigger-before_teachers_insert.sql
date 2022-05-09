-- 2564 · Create a Trigger "before_teachers_insert"
-- # Description
-- We want to create a trigger for the teachers table and name it before_teachers_insert, 
-- which should be executed before adding a new data, the trigger execution should be SET new.country = 'CN', 
-- please write SQL statement to achieve it

-- Table definition : teachers (Teachers table)
-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	Lecturer's name
-- email	varchar	Instructor's email
-- age	int	lecturer's age
-- country	varchar	Tutor's nationality


-- Example
-- Input data :

-- id	name	email	age	country
-- 1	'Linghu Chong'	None	18	'CN'
-- Return results :

-- id	name	email	age	country
-- 1	'Linghu Chong'	None	18	'CN'
-- 2	'Eastern heretic'	'eastern.heretic@gmail.com'	33	'CN'
-- 3	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 4	'Western Venom'	'western.venom@163.com'	28	'CN'
-- 5	'Southern Emperor'	'southern.emperor@qq.com'	31	'CN'

-- MySQL触发器的创建语法：
-- CREATE　[DEFINER = { 'user' | CURRENT_USER }]　
-- TRIGGER trigger_name
-- trigger_time trigger_event
-- ON table_name
-- FOR EACH ROW
-- [trigger_order]
-- trigger_body

-- DEFINER=	可选参数，指定创建者，默认为当前登录用户（CURRENT_USER）；
-- 该触发器将以此参数指定的用户执行，所以需要考虑权限问题；	DEFINER='root@%'
-- DEFINER=CURRENT_USER
-- trigger_name	触发器名称，最好由表名+触发事件关键词+触发时间关键词组成；	 
-- trigger_time	触发时间，在某个事件之前还是之后；	BEFORE、AFTER
-- trigger_event	触发事件，如插入时触发、删除时触发；
-- 　　INSERT：插入操作触发器，INSERT、LOAD DATA、REPLACE时触发；
-- 　　UPDATE：更新操作触发器，UPDATE操作时触发；
-- 　　DELETE：删除操作触发器，DELETE
--    REPLACE操作时触发；	INSERT、UPDATE、DELETE
-- table_name 	触发操作时间的表名；	 
-- trigger_order	可选参数，如果定义了多个具有相同触发事件和触法时间的触发器时（
-- 如：BEFORE UPDATE），默认触发顺序与触发器的创建顺序一致，可以
-- 使用此参数来改变它们触发顺序。mysql 5.7.2起开始支持此参数。
-- 　　FOLLOWS：当前创建触发器在现有触发器之后激活；
-- 　　PRECEDES：当前创建触发器在现有触发器之前激活；	FOLLOWS、PRECEDES
-- trigger_body	触发执行的SQL语句内容，一般以begin开头，end结尾	begin .. end

-- MySQL触发器创建进阶：
-- 1、MySQL触发器中使用变量：
-- 　　MySQL触发器中变量变量前面加'@'，无需定义，可以直接使用：
-- -- 变量直接赋值
-- set @num=999;
 
-- -- 使用select语句查询出来的数据方式赋值，需要加括号：
-- set @name =(select name from table);
-- 2、MySQL触发器中使用if语做条件判断：
-- -- 简单的if语句：
-- set sex = if (new.sex=1, '男', '女');
 
-- -- 多条件if语句：
-- if old.type=1 then
--     update table ...;
-- elseif old.type=2 then
--     update table ...;
-- end if;
-- 　
-- mysql 查看触发器：
-- 1 可以使用“show triggers;”查看触发器
--      show triggers;
-- 2 由于MySQL创建的触发器保存在“information_schema库中的triggers表中，所以还可以通过查询此表查看触发器：
-- 通过information_schema.triggers表查看触发器：
--      select * from information_schema.triggers;
 
-- -- mysql 查看当前数据库的触发器
-- show triggers;
 
-- -- mysql 查看指定数据库"aiezu"的触发器
-- show triggers from aiezu;

-- MySQL触发器的创建语法：
-- CREATE　[DEFINER = { 'user' | CURRENT_USER }]　
-- TRIGGER trigger_name
-- trigger_time trigger_event
-- ON table_name
-- FOR EACH ROW
-- [trigger_order]
-- trigger_body

-- CREATE TRIGGER <trigger_name>
-- <trigger_time> <trigger_event> ON <table_name>
-- FOR EACH ROW
-- BEGIN
--   ...
-- END;
--
-- trigger_name 最好由表名+触发事件关键词+触发时间关键词组成；
-- trigger_time  触发时间，在某个事件之前还是之后；	BEFORE、AFTER
-- trigger_event 触发事件，如插入时触发、删除时触发；
-- 　　INSERT：插入操作触发器，INSERT、LOAD DATA、REPLACE时触发；
-- 　　UPDATE：更新操作触发器，UPDATE操作时触发；
-- 　　DELETE：删除操作触发器，DELETE
--    REPLACE操作时触发；	INSERT、UPDATE、DELETE
-- table_name 触发操作时间的表名；
-- trigger_order	可选参数，如果定义了多个具有相同触发事件和触法时间的触发器时（
-- 如：BEFORE UPDATE），默认触发顺序与触发器的创建顺序一致，可以
-- 使用此参数来改变它们触发顺序。mysql 5.7.2起开始支持此参数。
-- 　　FOLLOWS：当前创建触发器在现有触发器之后激活；
-- 　　PRECEDES：当前创建触发器在现有触发器之前激活；	FOLLOWS、PRECEDES
-- trigger_body	触发执行的SQL语句内容，一般以begin开头，end结尾	begin .. end

-- Write your SQL here --
DROP TRIGGER IF EXISTS `before_teachers_insert`;
CREATE TRIGGER `before_teachers_insert`
BEFORE INSERT ON `teachers`
FOR EACH ROW
BEGIN
	SET new.country = 'CN';
END
