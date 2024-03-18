# 596. Classes More Than 5 Students
# Table: Courses
# +-------------+---------+
# | Column Name | Type    |
# +-------------+---------+
# | student     | varchar |
# | class       | varchar |
# +-------------+---------+
# (student, class) is the primary key column for this table.
# Each row of this table indicates the name of a student and the class in which they are enrolled.
#  
# Write an SQL query to report all the classes that have at least five students.
# Return the result table in any order.
# The query result format is in the following example.

# Example 1:
# Input:
# Courses table:
# +---------+----------+
# | student | class    |
# +---------+----------+
# | A       | Math     |
# | B       | English  |
# | C       | Math     |
# | D       | Biology  |
# | E       | Math     |
# | F       | Computer |
# | G       | Math     |
# | H       | Math     |
# | I       | Math     |
# +---------+----------+
# Output:
# +---------+
# | class   |
# +---------+
# | Math    |
# +---------+
# Explanation:
# - Math has 6 students, so we include it.
# - English has 1 student, so we do not include it.
# - Biology has 1 student, so we do not include it.
# - Computer has 1 student, so we do not include it.

import pandas as pd

def find_classes(courses: pd.DataFrame) -> pd.DataFrame:
    # 统计
    courses = courses.groupby("class")["student"].size().reset_index(name="count")
    # 过滤
    filter = courses["count"] >= 5
    # 取值
    return courses[filter][["class"]]

def find_classes1(courses: pd.DataFrame) -> pd.DataFrame:
    courses = courses.groupby('class')['student'].nunique().reset_index()
    courses.rename(columns={'student': 'student_cnt'}, inplace=True)
    return courses.loc[courses['student_cnt'] >= 5, ['class']]

if __name__ == "__main__":
    data = [['A', 'Math'], ['B', 'English'], ['C', 'Math'], ['D', 'Biology'], ['E', 'Math'], ['F', 'Computer'], ['G', 'Math'], ['H', 'Math'], ['I', 'Math']]
    courses = pd.DataFrame(data, columns=['student', 'class']).astype({'student':'object', 'class':'object'})
    print(find_classes(courses))
    print(find_classes1(courses))
