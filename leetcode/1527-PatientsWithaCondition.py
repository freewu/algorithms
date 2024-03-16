# 1527. Patients With a Condition
# Table: Patients
# +--------------+---------+
# | Column Name  | Type    |
# +--------------+---------+
# | patient_id   | int     |
# | patient_name | varchar |
# | conditions   | varchar |
# +--------------+---------+
# patient_id is the primary key for this table.
# 'conditions' contains 0 or more code separated by spaces.
# This table contains information of the patients in the hospital.

# Write an SQL query to report the patient_id, patient_name all conditions of patients who have Type I Diabetes.
# Type I Diabetes always starts with DIAB1 prefix
# Return the result table in any order.
# The query result format is in the following example.

# Example 1:
# Input:
# Patients table:
# +------------+--------------+--------------+
# | patient_id | patient_name | conditions   |
# +------------+--------------+--------------+
# | 1          | Daniel       | YFEV COUGH   |
# | 2          | Alice        |              |
# | 3          | Bob          | DIAB100 MYOP |
# | 4          | George       | ACNE DIAB100 |
# | 5          | Alain        | DIAB201      |
# +------------+--------------+--------------+
# Output:
# +------------+--------------+--------------+
# | patient_id | patient_name | conditions   |
# +------------+--------------+--------------+
# | 3          | Bob          | DIAB100 MYOP |
# | 4          | George       | ACNE DIAB100 |
# +------------+--------------+--------------+
# Explanation: Bob and George both have a condition that starts with DIAB1.

import pandas as pd

# contains + startswith
def find_patients(patients: pd.DataFrame) -> pd.DataFrame:
    #return patients.query("conditions like '%DIAB1%'")
    #return patients[patients["conditions"].str.contains("DIAB1")]
    #return patients[patients["conditions"].str.match(r"DIAB1") | patients["conditions"].str.match(r"\s+DIAB1")]
    # Type I Diabetes always starts with DIAB1 prefix
    filter = (patients["conditions"].str.contains(" DIAB1") | patients["conditions"].str.startswith("DIAB1"))
    return patients[filter]

# contains + regex
def find_patients1(patients: pd.DataFrame) -> pd.DataFrame:
    return patients[patients['conditions'].str.contains(r'\bDIAB1', regex=True)]

# regex + match
def find_patients2(patients: pd.DataFrame) -> pd.DataFrame:
    return patients[patients['conditions'].str.match(r'.*\bDIAB1.*')]


if __name__ == "__main__":
    data = [[1, 'Daniel', 'YFEV COUGH'], [2, 'Alice', ''], [3, 'Bob', 'DIAB100 MYOP'], [4, 'George', 'ACNE DIAB100'], [5, 'Alain', 'DIAB201']]
    patients = pd.DataFrame(data, columns=['patient_id', 'patient_name', 'conditions']).astype({'patient_id':'int64', 'patient_name':'object', 'conditions':'object'})
    print(find_patients(patients))
    print(find_patients1(patients))
    print(find_patients2(patients))