FROM golang:alpine as builder
WORKDIR /app
COPY . ./
RUN go get && go build -v -o shrt_server
RUN ls

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/shrt_server /app/server
RUN ls
COPY --from=builder /app/config.yaml /app/config.yaml
ENTRYPOINT ["/app/server"]