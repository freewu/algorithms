-- 2720 · Creating a view that ensures consistency (I)
-- # Description
-- Now you are asked to create an updatable view v_teachers that only allows viewing and inserting information 
-- about teachers who are less than 30 years old, write SQL statement to achieve it.

-- Table Definition : teachers (Teachers table)
-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	Lecturer's name
-- email	varchar	Instructor's email
-- age	int	lecturer's age
-- country	varchar	tutor's nationality

-- View Definition: v_teachers(View)
-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	Lecturer's name
-- email	varchar	Instructor's email
-- age	int	lecturer's age
-- country	varchar	Tutor's nationality


-- Example
-- Input：

-- teachers：
-- id	name	email	age	country
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 3	'Western Venom'	'western.venom@163.com'	28	'USA'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'

-- Return：
-- id	name	email	age	country
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 3	'Western Venom'	'western.venom@163.com'	28	'USA'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'

-- Write your SQL here --
CREATE OR REPLACE VIEW 
	`v_teachers`
AS 
	SELECT 
		* 
	FROM
		`teachers`
	WHERE
		age < 30
WITH CHECK OPTION;


-- # 创建视图
-- CREATE [OR REPLACE] [ALGORITHM = {UNDEFINED | MERGE | TEMPTABLE}]
--     VIEW view_name [(column_list)]
--     AS select_statement
--    [WITH [CASCADED | LOCAL] CHECK OPTION]
--
-- 1）OR REPLACE：表示替换已有视图
-- 2）ALGORITHM：表示视图选择算法，默认算法是 UNDEFINED(未定义的)： MySQL自动选择要使用的算法 ；merge合并；temptable临时表
--              MERGE：将视图的语句与视图定义合并起来，使得视图定义的某一部分取代语句的对应部分
--              TEMPTABLE：将视图的结果存入临时表，然后使用临时表执行语句
-- 3）select_statement：表示select语句
-- 4）[WITH [CASCADED | LOCAL] CHECK OPTION]：表示视图在更新时保证在视图的权限范围之内
-- 　　cascade 是默认值，表示更新视图的时候，要满足视图和表的相关条件
-- 　　local表示更新视图的时候，要满足该视图定义的一个条件即可
--     推荐使用WHIT [CASCADED|LOCAL] CHECK OPTION选项，可以保证数据的安全性 