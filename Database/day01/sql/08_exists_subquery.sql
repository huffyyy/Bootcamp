-- subquery
SELECT * FROM hr.employees WHERE EXISTS (
	SELECT d.department_id
	FROM hr.departments as d
	JOIN hr.employees as e ON e.department_id = d.department_id
	GROUP BY d.department_id, department_name
)