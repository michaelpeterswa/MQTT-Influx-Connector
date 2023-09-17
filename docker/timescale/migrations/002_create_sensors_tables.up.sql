CREATE TABLE IF NOT EXISTS sensors.bme280 
(
    time TIMESTAMPTZ NOT NULL,
    type   VARCHAR(20) NOT NULL,
    location VARCHAR(20) NOT NULL,
    room VARCHAR(20) NOT NULL,
    name VARCHAR(20) NOT NULL,
    field VARCHAR(20) NOT NULL,
    temperature DOUBLE PRECISION,
    humidity DOUBLE PRECISION,
    pressure DOUBLE PRECISION,
    rssi DOUBLE PRECISION,
    voltage DOUBLE PRECISION
);

SELECT create_hypertable('sensors.bme280', 'time');