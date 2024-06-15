# base go image
FROM golang:1.22-alpine as builder
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 go build -o frontendApp ./cmd/web
RUN chmod +x /app/frontendApp

# build a tiny docker image
FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/frontendApp /app
CMD [ "/app/frontendApp" ]