INSERT INTO
    sensors.renogychargecontroller (
        time,
        type,
        location,
        room,
        name,
        field,
        battery_capacity_soc,
        battery_voltage,
        charging_current,
        controller_temperature,
        battery_temperature,
        street_light_load_volatge,
        street_light_load_current,
        street_light_load_power,
        solar_panel_voltage,
        solar_panel_current,
        charging_power,
        battery_minimum_voltage_current_day,
        battery_maximum_voltage_current_day,
        maximum_charging_current_current_day,
        maximum_discharging_current_current_day,
        maximum_charging_power_current_day,
        maximum_discharging_power_current_day,
        charging_amp_hours_current_day,
        discharging_amp_hours_current_day,
        power_generation_current_day,
        power_consumption_current_day,
        total_operating_days,
        total_battery_over_discharges,
        total_battery_full_charges,
        total_charging_amp_hours,
        total_discharging_amp_hours,
        cumulative_power_generation,
        cumulative_power_consumption,
        street_light_stauts,
        street_light_brightness,
        charging_state,
        controller_faults,
        rssi
    )
VALUES
    ( 
        $1, 
        $2, 
        $3, 
        $4, 
        $5, 
        $6, 
        $7, 
        $8, 
        $9, 
        $10, 
        $11, 
        $12, 
        $13, 
        $14, 
        $15, 
        $16, 
        $17, 
        $18, 
        $19, 
        $20, 
        $21, 
        $22, 
        $23, 
        $24, 
        $25, 
        $26, 
        $27, 
        $28, 
        $29, 
        $30, 
        $31, 
        $32, 
        $33, 
        $34, 
        $35, 
        $36, 
        $37, 
        $38, 
        $39
    )