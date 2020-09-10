FROM golang:alpine as builder
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"
ADD main.go  .
RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM alpine
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"
COPY --from=builder /go/app /app
ENTRYPOINT ["/app"]