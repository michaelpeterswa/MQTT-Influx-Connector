all: build

build:
	docker build -t michaelpeterswa/mqtt-influx-connector .

publish:
	docker push michaelpeterswa/mqtt-influx-connector
