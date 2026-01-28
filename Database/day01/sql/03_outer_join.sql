-- left join
SELECT 
	c.country_id, country_name,
	location_id, street_address, city, state_province
FROM hr.countries as c
LEFT JOIN hr.locations as l ON c.country_id = l.country_id

-- right join
SELECT
	c.country_id, country_name,
	location_id, street_address, city, state_province
FROM hr.countries as c
RIGHT JOIN hr.locations as l ON c.country_id = l.country_id

-- outer join
SELECT 
	c.country_id, country_name,
	location_id, street_address, city, state_province
FROM hr.countries as c
FULL OUTER JOIN hr.locations as l ON c.country_id = l.country_id
WHERE LOWER(country_name) LIKE LOWER('c%')