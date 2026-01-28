SELECT r.region_id, region_name, country_id, country_name
FROM hr.regions as r
JOIN hr.countries as c ON r.region_id = c.region_id