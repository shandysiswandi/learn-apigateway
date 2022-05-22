FROM golang:1.17-alpine as builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM alpine:3.15
COPY --from=builder /build/server .
EXPOSE 8080
ENTRYPOINT [ "./server" ]