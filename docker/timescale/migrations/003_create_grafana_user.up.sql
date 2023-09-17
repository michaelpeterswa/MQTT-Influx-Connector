CREATE USER grafana WITH PASSWORD 'grafana';

GRANT USAGE ON SCHEMA sensors TO grafana;

GRANT SELECT ON sensors.bme280 TO grafana;

ALTER ROLE grafana SET search_path = 'sensors';