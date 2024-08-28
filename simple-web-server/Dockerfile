FROM golang:1.18 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -v -o /app/server

FROM gcr.io/distroless/base-debian11

COPY --from=builder /app/server /app/server

EXPOSE 3000

CMD ["/app/server"]
