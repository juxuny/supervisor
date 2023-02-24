FROM golang:1.15.4 as builder
COPY . /src/
RUN cd /src && GOPROXY=https://goproxy.cn go mod download
RUN cd /src/cmd/proxy && CGO_ENABLED=0 go build -o proxy

# final stage
FROM juxuny/alpine:3.13.5
WORKDIR /app
COPY --from=builder /src/cmd/proxy/proxy /app/proxy
ENTRYPOINT /app/proxy -e