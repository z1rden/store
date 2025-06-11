FROM golang AS builder
LABEL stage=builder

WORKDIR /cart

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./bin/cart ./cmd/cart/main.go

FROM ubuntu

COPY --from=builder /cart/bin/cart /cart
COPY --from=builder /cart/config/config_cart.yaml /config/

CMD [ "/cart" ]