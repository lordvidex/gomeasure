FROM golang:alpine AS builder
WORKDIR /go/src/github.com/lordvidex/gomeasure
COPY . .
RUN go mod download
RUN go build -o bin/gomeasure main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/github.com/lordvidex/gomeasure/bin/gomeasure .
RUN ls -la
ENTRYPOINT ["./gomeasure"]