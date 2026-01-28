-- CTE (Common Table Expression) for department aggregation
WITH ct_depts AS (
    SELECT d.department_id,
           department_name,
           COUNT(employee_id) AS total_employee
    FROM hr.departments d
    JOIN hr.employees e ON e.department_id = d.department_id
    GROUP BY d.department_id, department_name
)
SELECT *
FROM ct_depts
WHERE total_employee > 5;