FROM golang:1.15.4 as builder
COPY . /src/
RUN cd /src && GOPROXY=https://goproxy.io go mod download
RUN cd /src/cmd/supervisord && CGO_ENABLED=0 go build -o supervisord

# final stage
FROM ineva/alpine:3.9
WORKDIR /app
COPY --from=builder /src/cmd/supervisord/supervisord /app/supervisord
ENTRYPOINT /app/supervisord -c config/supervisor.yaml