-- tampilkan salary minimum tiap departement(min)
SELECT d.department_id, department_name, min(salary) as min_salary
FROM hr.departments as d
JOIN hr.employees as e ON e.department_id = d.department_id
GROUP BY d.department_id, department_name
ORDER BY min_salary
