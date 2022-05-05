-- 2018 · Insert 'Flash Sale' course information into the specified column of the course schedule
-- # Description
-- Write an SQL statement to insert a new course record into the course table courses,the record is as follows:

-- name	student_count	created_at	teacher_id
-- Flash Sale	100	2018-01-01	5
-- Table definition: courses

-- column name	type	comment
-- id	int	primary key
-- name	varchar	course name
-- student_count	int	student count
-- created_at	datetime	course start time
-- teacher_id	int	teacher id


-- The id column of the courses table course is the primary key, self-growing and does not require a value to be set.
-- The inserted record field type should match the table definition field type
-- Example
-- Example 1

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
-- 10	Object Oriented Design	300	2020-08-08	1
-- 12	Dynamic Programming	2000	2018-08-18	1
-- After running your SQL statement, we will execute the following statement:

-- SELECT *
-- FROM `courses`;
-- return to:

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-06-01	4
-- 2	System Design	1350	2020-07-18	3
-- 3	Django	780	2020-02-29	3
-- 4	Web	340	2020-04-22	4
-- 5	Big Data	700	2020-09-11	1
-- 6	Artificial Intelligence	1660	2018-05-13	3
-- 7	Java P6+	780	2019-01-19	3
-- 8	Data Analysis	500	2019-07-12	1
-- 10	Object Oriented Design	300	2020-08-08	1
-- 12	Dynamic Programming	2000	2018-08-18	1
-- 13	Flash Sale	100	2018-01-01	5
-- Example 2

-- Table content: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-06-01	4
-- 2	System Design	1350	2020-07-18	3
-- 3	Django	780	2020-02-29	3
-- 4	Web	340	2020-04-22	4
-- After running your SQL statement, we will execute the following statement:

-- SELECT *
-- FROM `courses`;
-- return to:

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-06-01	4
-- 2	System Design	1350	2020-07-18	3
-- 3	Django	780	2020-02-29	3
-- 4	Web	340	2020-04-22	4
-- 5	Flash Sale	100	2018-01-01	5

INSERT INTO `courses`(
	`name`,
	`student_count`,
	`created_at`,
	`teacher_id`
) VALUES (
	'Flash Sale',
	100,
	'2018-01-01',
	5
)