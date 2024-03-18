# 2356. Number of Unique Subjects Taught by Each Teacher
# Table: Teacher
# +-------------+------+
# | Column Name | Type |
# +-------------+------+
# | teacher_id  | int  |
# | subject_id  | int  |
# | dept_id     | int  |
# +-------------+------+
# (subject_id, dept_id) is the primary key (combinations of columns with unique values) of this table.
# Each row in this table indicates that the teacher with teacher_id teaches the subject subject_id in the department dept_id.
 
# Write a solution to calculate the number of unique subjects each teacher teaches in the university.
# Return the result table in any order.
# The result format is shown in the following example.

# Example 1:
# Input: 
# Teacher table:
# +------------+------------+---------+
# | teacher_id | subject_id | dept_id |
# +------------+------------+---------+
# | 1          | 2          | 3       |
# | 1          | 2          | 4       |
# | 1          | 3          | 3       |
# | 2          | 1          | 1       |
# | 2          | 2          | 1       |
# | 2          | 3          | 1       |
# | 2          | 4          | 1       |
# +------------+------------+---------+
# Output:  
# +------------+-----+
# | teacher_id | cnt |
# +------------+-----+
# | 1          | 2   |
# | 2          | 4   |
# +------------+-----+
# Explanation: 
# Teacher 1:
#   - They teach subject 2 in departments 3 and 4.
#   - They teach subject 3 in department 3.
# Teacher 2:
#   - They teach subject 1 in department 1.
#   - They teach subject 2 in department 1.
#   - They teach subject 3 in department 1.
#   - They teach subject 4 in department 1.

import pandas as pd

# def count_unique_subjects(teacher: pd.DataFrame) -> pd.DataFrame:
#     # 取值 
#     teacher = teacher[["teacher_id","subject_id"]]
#     # 去重
#     teacher = teacher.drop_duplicates()
#     # 统计
#     teacher = teacher.groupby('teacher_id').count().reset_index('cnt')
#     return teacher

def count_unique_subjects(teacher: pd.DataFrame) -> pd.DataFrame:
    # 取值 
    teacher = teacher[['teacher_id', 'subject_id']]
    # 去重
    teacher = teacher.drop_duplicates()
    # 按 teacher_id 分组统计并设定名称
    return teacher.groupby('teacher_id').size().reset_index(name='cnt')

# nunique
def count_unique_subjects1(teacher: pd.DataFrame) -> pd.DataFrame:
    group = teacher.groupby("teacher_id")['subject_id'].nunique().reset_index()
    return group.rename(columns={'subject_id': 'cnt'})
    

if __name__ == "__main__":
    data = [[1, 2, 3], [1, 2, 4], [1, 3, 3], [2, 1, 1], [2, 2, 1], [2, 3, 1], [2, 4, 1]]
    teacher = pd.DataFrame(data, columns=['teacher_id', 'subject_id', 'dept_id']).astype({'teacher_id':'Int64', 'subject_id':'Int64', 'dept_id':'Int64'})
    print(count_unique_subjects(teacher))
    print(count_unique_subjects1(teacher))

# value_counts( )函数
# 在pandas中，value_counts常用于数据表的计数及排序，它可以用来查看数据表中，指定列里有多少个不同的数据值，并计算每个不同值有在该列中的个数，同时还能根据需要进行排序。
# 函数体
#     value_counts(values,sort=True, ascending=False, normalize=False,bins=None,dropna=True)

# 主要参数：
#     sort=True： 是否要进行排序；默认进行排序
#     ascending=False： 默认降序排列；
#     normalize=False： 是否要对计算结果进行标准化并显示标准化后的结果，默认是False。
#     bins=None： 可以自定义分组区间，默认是否；
#     dropna=True：是否删除缺失值nan，默认删除
