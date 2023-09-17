package timescale

import (
	"context"
	"fmt"
	"time"

	_ "embed"

	"github.com/jackc/pgx/v5"
	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/structs"
)

//go:embed write_bme280.sql
var writeBME280SQL string

type TimescaleConn struct {
	Conn *pgx.Conn
}

func InitTimescale(ctx context.Context, settings *structs.MQTTInfluxConnectorSettings) (*TimescaleConn, error) {
	conn, err := pgx.Connect(ctx, settings.TimescaleAddress)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	return &TimescaleConn{
		Conn: conn,
	}, nil
}

func (t *TimescaleConn) Close(ctx context.Context) {
	t.Conn.Close(ctx)
}

func (t *TimescaleConn) WriteBME280SensorData(ctx context.Context, reading structs.BME280, st structs.SubTopic) error {
	_, err := t.Conn.Exec(ctx, writeBME280SQL, time.Now(), st.Type, st.Location, st.Room, st.Name, st.Field, reading.Temperature, reading.Humidity, reading.Pressure, reading.RSSI, reading.Voltage)
	if err != nil {
		return fmt.Errorf("unable to write bme280 data to timescale: %w", err)
	}

	return nil
}
