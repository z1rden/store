FROM golang AS builder
LABEL stage=builder

WORKDIR /loms

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./bin/loms ./cmd/loms/main.go

FROM ubuntu

COPY --from=builder /loms/bin/loms /loms
COPY --from=builder /loms/config/ /config/
COPY --from=builder /loms/pkg/swagger /pkg/swagger
COPY --from=builder /loms/pkg/swagger-ui /pkg/swagger-ui
COPY --from=builder /loms/migrations /migrations

CMD [ "/loms" ]
