# -=-=-=-=-=-=- Compile Image -=-=-=-=-=-=-

FROM golang:1 AS stage-compile

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./... && CGO_ENABLED=0 GOOS=linux go build ./cmd/mqtt-influx-connector

# -=-=-=-=- Final Distroless Image -=-=-=-=-

FROM gcr.io/distroless/static-debian11:latest-amd64 as stage-final

COPY --from=stage-compile /go/src/app/mqtt-influx-connector /
CMD ["/mqtt-influx-connector"]