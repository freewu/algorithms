-- 2570 · New data processing of recruitment information statistics table
-- # Description
-- The students table stores information about all students, including student id and student name name
-- The companies table stores all company information, including company id and company name name.
-- The recording table stores all resume submissions, including student id (student_id) and company id (company_id)
-- Write SQL statements to process the new data in the resume submission table, 
-- setting student_id to 0 when it does not exist in the student table, and setting company_id to 0 when it does not exist in the companies table.

-- Table Definition 1: students (students table)

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	student name
-- Table Definition 2: companies (Company Table)

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	company name
-- address	varchar	company address
-- Table Definition 3: recording (record table)

-- column name	type	comment
-- id	int unsigned	primary key
-- delivery_date	date	date of delivery
-- company_id	int	company id
-- student_id	int	student id

-- Example
-- Table Contents 1: students

-- id	name
-- 1	Da Ming
-- 2	Amy
-- 3	Mike
-- 4	Park
-- 5	George
-- Table 2: Companies

-- id	name	address
-- 1	Alibaba	Hang Zhou
-- 2	NetEase	Guang Zhou
-- 3	Baidu	Bei Jing
-- 4	Tencent	Shen Zhen
-- Table 3: recording

-- id	delivery_date	company_id	student_id
-- After running your SQL statement, it should return.

-- id	delivery_date	company_id	student_id
-- 1	'2021-01-14'	2	1
-- 2	'2021-03-21'	1	0
-- 3	'2021-04-13'	0	4
-- 4	'2021-02-26'	0	0

-- Write your SQL here --
DROP TRIGGER IF EXISTS `before_recording_insert`;
CREATE TRIGGER `before_recording_insert`
BEFORE INSERT ON `recording`
FOR EACH ROW
BEGIN
	DECLARE s int;
	DECLARE c int;

	-- 判断 学生 id 是否存在
	SET s = (SELECT COUNT(*) FROM `students` WHERE id = new.student_id);
	-- 学生 id 不存在
	IF s = 0 THEN
		-- 当 student_id 在 students 表中不存在时，置为 0
		SET new.student_id = 0;
	END IF;

	-- 判断 公司 id 是否存在
	SET c = (SELECT COUNT(*) FROM `companies` WHERE id = new.company_id);
	-- 公司 id 不存在
	IF c = 0 THEN
		-- 当 company_id 在 companies 表中不存在时，置为 0
		SET new.company_id = 0;
	END IF;
END