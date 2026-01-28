CREATE VIEW vw_region_country_locations as 
	SELECT r.region_id, region_name, c.country_id, country_name, l.location_id, street_address, city
	FROM hr.regions as r
	JOIN hr.countries as c ON r.region_id = c.region_id
	JOIN hr.locations as l ON c.country_id = l.country_id

SELECT * FROM vw_region_country_location