-- tampilkan jumlah pegawai tiap departement (count)
SELECT d.department_id, department_name, count(employee_id) as total_employee
FROM hr.departments as d
JOIN hr.employees as e ON e.department_id = d.department_id
GROUP BY d.department_id, department_name
ORDER BY total_employee desc