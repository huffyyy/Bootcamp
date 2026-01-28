-- in operator
SELECT 
	r.region_id, region_name,
	country_id, country_name
FROM hr.regions as r
JOIN hr.countries as c ON r.region_id = c.region_id
WHERE country_id IN ('CN', 'CH', 'CA')

-- exist operator
SELECT * FROM hr.departments as d WHERE EXISTS (
	SELECT l.location_id
	FROM hr.regions as r
	JOIN hr.countries as c ON r.region_id = c.region_id
	JOIN hr.locations as l ON l.country_id = c.country_id
)