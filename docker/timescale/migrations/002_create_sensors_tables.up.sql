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

SELECT create_hypertable('sensors.bme280', 'time', if_not_exists => TRUE);

CREATE TABLE IF NOT EXISTS sensors.tsl2561 
(
    time TIMESTAMPTZ NOT NULL,
    type   VARCHAR(20) NOT NULL,
    location VARCHAR(20) NOT NULL,
    room VARCHAR(20) NOT NULL,
    name VARCHAR(20) NOT NULL,
    field VARCHAR(20) NOT NULL,
    lux DOUBLE PRECISION,
    rssi DOUBLE PRECISION
);

SELECT create_hypertable('sensors.tsl2561', 'time', if_not_exists => TRUE);

CREATE TABLE IF NOT EXISTS sensors.pmsa003i
(
    time TIMESTAMPTZ NOT NULL,
    type   VARCHAR(20) NOT NULL,
    location VARCHAR(20) NOT NULL,
    room VARCHAR(20) NOT NULL,
    name VARCHAR(20) NOT NULL,
    field VARCHAR(20) NOT NULL,
    pm10s DOUBLE PRECISION,
    pm25s DOUBLE PRECISION,
    pm100s DOUBLE PRECISION,
    pm10e DOUBLE PRECISION,
    pm25e DOUBLE PRECISION,
    pm100e DOUBLE PRECISION,
    pm03um DOUBLE PRECISION,
    pm05um DOUBLE PRECISION,
    pm10um DOUBLE PRECISION,
    pm25um DOUBLE PRECISION,
    pm50um DOUBLE PRECISION,
    pm100um DOUBLE PRECISION,
    rssi DOUBLE PRECISION
);

SELECT create_hypertable('sensors.pmsa003i', 'time', if_not_exists => TRUE);

CREATE TABLE IF NOT EXISTS sensors.renogychargecontroller
(
    time TIMESTAMPTZ NOT NULL,
    type   VARCHAR(20) NOT NULL,
    location VARCHAR(20) NOT NULL,
    room VARCHAR(20) NOT NULL,
    name VARCHAR(20) NOT NULL,
    field VARCHAR(20) NOT NULL,
    battery_capacity_soc INT,
    battery_voltage DOUBLE PRECISION,
    charging_current DOUBLE PRECISION,
    controller_temperature INT,
    battery_temperature INT,
    street_light_load_volatge DOUBLE PRECISION,
    street_light_load_current DOUBLE PRECISION,
    street_light_load_power DOUBLE PRECISION,
    solar_panel_voltage DOUBLE PRECISION,
    solar_panel_current DOUBLE PRECISION,
    charging_power DOUBLE PRECISION,
    battery_minimum_voltage_current_day DOUBLE PRECISION,
    battery_maximum_voltage_current_day DOUBLE PRECISION,
    maximum_charging_current_current_day DOUBLE PRECISION,
    maximum_discharging_current_current_day DOUBLE PRECISION,
    maximum_charging_power_current_day DOUBLE PRECISION,
    maximum_discharging_power_current_day DOUBLE PRECISION,
    charging_amp_hours_current_day DOUBLE PRECISION,
    discharging_amp_hours_current_day DOUBLE PRECISION,
    power_generation_current_day DOUBLE PRECISION,
    power_consumption_current_day DOUBLE PRECISION,
    total_operating_days INT,
    total_battery_over_discharges INT,
    total_battery_full_charges INT,
    total_charging_amp_hours DOUBLE PRECISION,
    total_discharging_amp_hours DOUBLE PRECISION,
    cumulative_power_generation DOUBLE PRECISION,
    cumulative_power_consumption DOUBLE PRECISION,
    street_light_stauts BOOLEAN,
    street_light_brightness INT,
    charging_state VARCHAR(40),
    controller_faults VARCHAR(40)[],
    rssi DOUBLE PRECISION
);

SELECT create_hypertable('sensors.renogychargecontroller', 'time', if_not_exists => TRUE);