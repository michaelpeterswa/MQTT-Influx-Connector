-- CREATE USER grafana WITH PASSWORD 'grafana';

GRANT USAGE ON SCHEMA sensors TO grafana;

GRANT SELECT ON sensors.bme280 TO grafana;
GRANT SELECT ON sensors.tsl2561 TO grafana;
GRANT SELECT ON sensors.pmsa003i TO grafana;
GRANT SELECT ON sensors.renogychargecontroller TO grafana;

-- ALTER ROLE grafana SET search_path to sensors,public;