-- 2576 · Update data processing of recruitment information statistics table
-- # Description
-- The students table stores information about all students, including student id and student name name
-- The companies table stores all company information, including company id and company name name.
-- The recording table stores all resume submissions, including student id (student_id) and company id (company_id)
-- ** Please write SQL statements to process the updated data in the recording table **

-- If student_id in the update data does not exist in the student table or company_id does not exist in the companies table, the update is not allowed.
-- When student_id does not exist, the value is not updated
-- Do not update the value when company_id does not exist
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
-- 1	'2021-01-14'	2	1
-- 2	'2021-03-21'	1	3
-- 3	'2021-04-13'	2	4
-- 4	'2021-02-22'	3	5
-- 5	'2021-02-16'	4	2
-- 6	'2021-01-26'	1	2
-- After running your SQL statement, the table should return.

-- id	delivery_date	company_id	student_id
-- 1	'2021-01-14'	4	3
-- 2	'2021-03-21'	1	3
-- 3	'2021-04-13'	2	4
-- 4	'2021-02-22'	3	5
-- 5	'2021-02-16'	3	2
-- 6	'2021-01-26'	3	2

-- Write your SQL here --
DROP TRIGGER IF EXISTS `before_recording_update`;
CREATE TRIGGER `before_recording_update`
BEFORE UPDATE ON `recording`
FOR EACH ROW
BEGIN
	DECLARE s int;
    	DECLARE c int;
	-- 是否存在该学生 0 / 1
	SET s = (SELECT COUNT(*) FROM `students` WHERE id = new.student_id);
	-- 是否存在该公司 0 / 1
	SET c = (SELECT COUNT(*) FROM `companies` WHERE id = new.company_id);
	-- 当 student_id 不存在时，则不更新该值
	IF s = 0 THEN
		SET new.student_id = old.student_id;
	END IF;
	-- 当 company_id 不存在时，则不更新该值
	IF c = 0 THEN
		SET new.company_id = old.company_id;
	END IF;
END