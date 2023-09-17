all: migrate-down migrate-up

.PHONY: migrate-up
migrate-up:
	docker run -v ./docker/timescale/migrations:/migrations --network host migrate/migrate \
    -path=/migrations/ -database "postgres://postgres:root@<timescale address>/postgres?sslmode=disable" up

.PHONY: migrate-down
migrate-down:
	docker run -v ./docker/timescale/migrations:/migrations --network host migrate/migrate \
    -path=/migrations/ -database "postgres://postgres:root@<timescale address>/postgres?sslmode=disable" down -all