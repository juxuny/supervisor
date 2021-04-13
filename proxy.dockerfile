FROM golang:1.15.4 as builder
COPY . /src/
RUN cd /src && GOPROXY=https://goproxy.io go mod download
RUN cd /src/cmd/proxy && CGO_ENABLED=0 go build -o proxy

# final stage
FROM ineva/alpine:3.9
WORKDIR /app
COPY --from=builder /src/cmd/proxy/proxy /app/proxy
ENTRYPOINT /app/proxy -e