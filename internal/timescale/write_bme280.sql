INSERT INTO
    sensors.bme280 (
        time,
        type,
        location,
        room,
        name,
        field,
        temperature,
        humidity,
        pressure,
        rssi,
        voltage
    )
VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8,  $9, $10, $11)