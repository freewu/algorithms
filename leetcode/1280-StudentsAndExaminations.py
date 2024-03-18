# 1280. Students and Examinations
# Table: Students
# +---------------+---------+
# | Column Name   | Type    |
# +---------------+---------+
# | student_id    | int     |
# | student_name  | varchar |
# +---------------+---------+
# student_id is the primary key (column with unique values) for this table.
# Each row of this table contains the ID and the name of one student in the school.
 
# Table: Subjects
# +--------------+---------+
# | Column Name  | Type    |
# +--------------+---------+
# | subject_name | varchar |
# +--------------+---------+
# subject_name is the primary key (column with unique values) for this table.
# Each row of this table contains the name of one subject in the school.
 
# Table: Examinations
# +--------------+---------+
# | Column Name  | Type    |
# +--------------+---------+
# | student_id   | int     |
# | subject_name | varchar |
# +--------------+---------+
# There is no primary key (column with unique values) for this table. It may contain duplicates.
# Each student from the Students table takes every course from the Subjects table.
# Each row of this table indicates that a student with ID student_id attended the exam of subject_name.
 
# Write a solution to find the number of times each student attended each exam.
# Return the result table ordered by student_id and subject_name.
# The result format is in the following example.

# Example 1:
# Input: 
# Students table:
# +------------+--------------+
# | student_id | student_name |
# +------------+--------------+
# | 1          | Alice        |
# | 2          | Bob          |
# | 13         | John         |
# | 6          | Alex         |
# +------------+--------------+
# Subjects table:
# +--------------+
# | subject_name |
# +--------------+
# | Math         |
# | Physics      |
# | Programming  |
# +--------------+
# Examinations table:
# +------------+--------------+
# | student_id | subject_name |
# +------------+--------------+
# | 1          | Math         |
# | 1          | Physics      |
# | 1          | Programming  |
# | 2          | Programming  |
# | 1          | Physics      |
# | 1          | Math         |
# | 13         | Math         |
# | 13         | Programming  |
# | 13         | Physics      |
# | 2          | Math         |
# | 1          | Math         |
# +------------+--------------+
# Output: 
# +------------+--------------+--------------+----------------+
# | student_id | student_name | subject_name | attended_exams |
# +------------+--------------+--------------+----------------+
# | 1          | Alice        | Math         | 3              |
# | 1          | Alice        | Physics      | 2              |
# | 1          | Alice        | Programming  | 1              |
# | 2          | Bob          | Math         | 1              |
# | 2          | Bob          | Physics      | 0              |
# | 2          | Bob          | Programming  | 1              |
# | 6          | Alex         | Math         | 0              |
# | 6          | Alex         | Physics      | 0              |
# | 6          | Alex         | Programming  | 0              |
# | 13         | John         | Math         | 1              |
# | 13         | John         | Physics      | 1              |
# | 13         | John         | Programming  | 1              |
# +------------+--------------+--------------+----------------+
# Explanation: 
# The result table should contain all students and all subjects.
# Alice attended the Math exam 3 times, the Physics exam 2 times, and the Programming exam 1 time.
# Bob attended the Math exam 1 time, the Programming exam 1 time, and did not attend the Physics exam.
# Alex did not attend any exams.
# John attended the Math exam 1 time, the Physics exam 1 time, and the Programming exam 1 time.

import pandas as pd

# agg
def students_and_examinations(students: pd.DataFrame, subjects: pd.DataFrame, examinations: pd.DataFrame) -> pd.DataFrame:
    return examinations\
        .groupby(['student_id', 'subject_name'], as_index=False)\
        .agg(attended_exams=('student_id', 'count'))\
        .merge(students.assign(t=1).merge(subjects.assign(t=1)).drop('t', axis=1), how='right')\
        .fillna(pd.DataFrame({'attended_exams': [0] * students.size * subjects.size}))\
        .sort_values(['student_id', 'subject_name'])\
        [['student_id', 'student_name', 'subject_name', 'attended_exams']]


def students_and_examinations1(students: pd.DataFrame, subjects: pd.DataFrame, examinations: pd.DataFrame) -> pd.DataFrame:
    students['key'] = 1
    subjects['key'] = 1
    examinations['key'] = 1
    # 按 'student_id','subject_name' 统计出数量
    examinations = examinations.groupby(['student_id','subject_name'])['key'].count().reset_index()
    # join subjects & examinations 两张表
    df = students.merge(subjects,on='key').merge(examinations,on=['student_id','subject_name'],how = 'left') 
    #print(df)
    df['attended_exams'] = df['key_y'].fillna(0)
    return df[['student_id','student_name','subject_name','attended_exams']].sort_values(by = ['student_id','subject_name'],ascending=[True,True])

if __name__ == "__main__":
    data = [[1, 'Alice'], [2, 'Bob'], [13, 'John'], [6, 'Alex']]
    students = pd.DataFrame(data, columns=['student_id', 'student_name']).astype({'student_id':'Int64', 'student_name':'object'})
    data = [['Math'], ['Physics'], ['Programming']]
    subjects = pd.DataFrame(data, columns=['subject_name']).astype({'subject_name':'object'})
    data = [[1, 'Math'], [1, 'Physics'], [1, 'Programming'], [2, 'Programming'], [1, 'Physics'], [1, 'Math'], [13, 'Math'], [13, 'Programming'], [13, 'Physics'], [2, 'Math'], [1, 'Math']]
    examinations = pd.DataFrame(data, columns=['student_id', 'subject_name']).astype({'student_id':'Int64', 'subject_name':'object'})
    print(students_and_examinations(students, subjects, examinations))
    print(students_and_examinations1(students, subjects, examinations))