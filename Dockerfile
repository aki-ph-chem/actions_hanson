FROM golang:1.22.4 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o actions_hanson 

FROM scratch
COPY --from=builder /app/actions_hanson /app
ENTRYPOINT ["/app"]
