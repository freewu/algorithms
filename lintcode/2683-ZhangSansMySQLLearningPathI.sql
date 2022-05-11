-- 2683 · Zhang San's MySQL Learning Path (I)
-- # Description
-- Zhang San has recently become interested in MySQL database and decided to learn it well,
-- but... The road of learning is always full of hardships. 
-- Zhang San encountered a problem today, and he couldn't come up with a solution even though he thought about it left and right,
-- and thought hard about it. The situation is like this, the existing teachers table records the information of teachers who have left, 
-- the value of id is the order of teachers leaving, and in the courses table records the information of courses taught by each teacher, 
-- the associated id is teacher_id. Now we know that teachers are divided into teams by nationality, 
-- and each team is formed with at least 3 teachers, if a team of teachers leave collectively, 
-- the information in the teachers table is recorded as continuous and the same nationality, 
-- now we need to count how many students need to wait for a new teacher if each team of teachers leave collectively, 
-- and the result is sorted by team nationality. He needs to create a view v_courses_teachers to solve this problem, students come to help Zhang San boy ~

-- Table Definition 1: teachers (Teachers table)
-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	Lecturer's name
-- email	varchar	Instructor's email
-- age	int	lecturer's age
-- country	varchar	tutor's nationality

-- Table Definition 2: courses (Course List)
-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course_name
-- student_count	int	total_students
-- created_at	date	Course creation time
-- teacher_id	int unsigned	instructor id

-- Result View Definition: v_courses_teachers(view)
-- column_name	type	comment
-- country	varchar	nationality of the departing faculty team
-- student_count	int	Total number of students waiting for new teachers

-- Example
-- Enter data:

-- courses table.

-- id	name	student_count	created_at	teacher_id
-- 1	'Advanced Algorithms'	880	'2020-6-1 09:03:12'	4
-- 2	'System Design'	1350	'2020-7-18 10:03:12'	8
-- 3	'Django'	780	'2020-2-29 12:03:12'	2
-- 4	'Web'	340	'2020-4-22 13:03:12'	4
-- 5	'Big Data'	700	'2020-9-11 16:03:12'	7
-- 6	'Artificial Intelligence'	1660	'2018-5-13 18:03:12'	3
-- 7	'Java P6+'	780	'2019-1-19 13:03:12'	3
-- 8	'Data Analysis'	500	'2019-7-12 13:03:12'	6
-- 10	'Object Oriented Design'	300	'2020-8-8 13:03:12'	4
-- 12	'Dynamic Programming'	2000	'2018-8-18 20:03:12'	1
-- teachers table.

-- id	name	email	age	country
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'JP'
-- 3	'Western Venom'	'western.venom@163.com'	28	'JP'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'
-- 5	'Linghu Chong'	None	18	'CN'
-- Returned results:

-- country	student_count
-- 'JP'	4740

create view 
	v_courses_teachers 
as
	select
		t.country, 
		sum(student_count) as student_count
    	from 
	    teachers t 
	left join courses c 
	on t.id = c.teacher_id
   	where (
	     -- 连续且国籍相同
        (
			t.country = (select country from teachers where id = t.id+1) and 
			t.country = (select country from teachers where id = t.id+2)
	   	) or (
		   	t.country = (select country from teachers where id = t.id-1) and 
			t.country = (select country from teachers where id = t.id+1)
		) or (
			t.country = (select country from teachers where id = t.id-1) and 
			t.country = (select country from teachers where id = t.id-2)
		)
    )
    group by t.country