-- 1929 · Analysis of Online Class II
-- # Description
-- The online_class_situations table shows the behavioral activities of some students in online classes.
-- Each row of data records the number of courses (may be 0) that a student has listened to after logging in to the course with the same device on the same day before quitting the online course.
-- Write a SQL statement to query the date each student The id of the device used to log in to the platform for the first time.

-- Table definition: online_class_situations

-- columns_name	type	explaination
-- student_id	int	student's id
-- device_id	int	device's id
-- date	date	Class date of the course
-- course_number	int	course number
-- The primary key of the table is (student_id, date) combined primary key
-- Tip:

-- The first time you log in to the platform to listen to a lesson, it means you have logged in to the platform and have a lesson record, that is, course_number> 0
-- Example
-- Example 1:

-- Table content: online_class_situations

-- student_id	device_id	date	course_number
-- 1	2	2020-03-01	5
-- 1	2	2020-04-02	6
-- 2	3	2020-05-25	1
-- 3	1	2020-03-02	0
-- 3	4	2020-12-03	5
-- After running your SQL statement, the table should return:

-- student_id	device_id
-- 1	2
-- 2	3
-- 3	4
-- Example 2:

-- Table content: online_class_situations

-- student_id	device_id	date	course_number
-- 1	2	2020-03-01	5
-- 2	2	2020-04-02	6
-- 3	2	2020-05-25	1
-- After running your SQL statement, the table should return:

-- student_id	device_id
-- 1	2
-- 2	2
-- 3	2
SELECT
	r.student_id AS student_id,
	r.device_id AS device_id
FROM
	online_class_situations AS r,
	(
		SELECT 
			student_id,
			MIN(`date`) AS `date`
		FROM 
			online_class_situations
		WHERE 
			course_number > 0
		GROUP BY student_id
	) AS p
WHERE 
	r.student_id = p.student_id AND
	r.date = p.date