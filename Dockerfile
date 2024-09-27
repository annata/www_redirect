FROM golang:1.22-alpine as builder
WORKDIR /www_redirect
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs=false -ldflags="-w -s"

FROM buildpack-deps:curl
WORKDIR /www_redirect
COPY --from=builder /www_redirect/www_redirect .
CMD ["/www_redirect/www_redirect"]