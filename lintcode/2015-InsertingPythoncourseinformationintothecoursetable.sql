-- 2015 · Inserting Python course information into the course table
-- Description
-- Write an SQL statement to insert a new course record into the course table courses,the record is as follows:

-- id	name	student_count	create_time	teacher_id
-- 13	Python	400	2021-05-23	3
-- Table definition: courses

-- column name	type	comment
-- id	int	primary key
-- name	varchar	course name
-- student_count	int	student count
-- created_at	datetime	course start time
-- teacher_id	int	teacher id

-- The inserted record field type should match the table definition field type

-- Example
-- Table content: Courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-06-01	4
-- 2	System Design	1350	2020-07-18	3
-- 3	Django	780	2020-02-29	3
-- 4	Web	340	2020-04-22	4
-- 5	Big Data	700	2020-09-11	1
-- 6	Artificial Intelligence	1660	2018-05-13	3
-- 7	Java P6+	780	2019-01-19	3
-- 8	Data Analysis	500	2019-07-12	1
-- 10	Object Oriented Design	300	2020-08-08	4
-- After running your SQL statement, we will execute the following statement to check the results:

-- SELECT *
-- FROM `courses`;
-- return to:

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	800	2020-06-01	4
-- 2	System Design	1350	2020-07-18	3
-- 3	Django	800	2020-02-29	3
-- 4	Web	340	2020-04-22	4
-- 5	Big Data	700	2020-09-11	1
-- 6	Artificial Intelligence	1660	2018-05-13	3
-- 7	Java P6+	780	2019-01-19	3
-- 8	Data Analysis	500	2019-07-12	1
-- 10	Object Oriented Design	300	2020-08-08	4
-- 13	Python	400	2021-05-23	3

INSERT INTO `courses`(
	`id`,
	`name`,
	`student_count`,
	`created_at`,
	`teacher_id`
) VALUES (
	13,
	'Python',
	400,
	'2021-05-23',
	3
)
