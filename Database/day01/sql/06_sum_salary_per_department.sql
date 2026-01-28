-- tampilkan total salary tiap departement (sum)
SELECT d.department_id, department_name, sum(salary) as total_salary
FROM hr.departments as d
JOIN hr.employees as e ON e.department_id = d.department_id
GROUP BY d.department_id, department_name
ORDER BY total_salary 