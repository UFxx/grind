FROM golang:1.25.4 AS builder
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/app cmd/app/main.go

FROM alpine:3.21 AS runtime
WORKDIR /src
COPY --from=builder /out/app /src/app

EXPOSE 8080
CMD ["/src/app"]