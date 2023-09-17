INSERT INTO
    sensors.tsl2561 (
        time,
        type,
        location,
        room,
        name,
        field,
        lux,
        rssi
    )
VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8)