FROM golang:1.19-alpine as builder

WORKDIR /build
COPY go.mod .
RUN go mod download
RUN go mod tidy
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /logity cmd/main.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder logity /bin/logity

ENTRYPOINT ["/bin/logity"]