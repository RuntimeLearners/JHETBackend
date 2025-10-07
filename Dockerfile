FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
RUN go mod download && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

FROM scratch
COPY --from=builder /app/server .
EXPOSE 8080
ENTRYPOINT ["./server"]
