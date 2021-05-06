FROM golang:1.15.4 as builder
COPY . /src/
RUN cd /src && GOPROXY=https://goproxy.io go mod download
RUN cd /src/cmd/supervisord && CGO_ENABLED=0 go build -o supervisord
RUN cd /src/cmd/supervisor-ctl && CGO_ENABLED=0 go build -o supervisor-ctl

# final stage
FROM ineva/alpine:3.9
WORKDIR /app
COPY --from=builder /src/cmd/supervisord/supervisord /app/supervisord
COPY --from=builder /src/cmd/supervisor-ctl/supervisor-ctl /app/supervisor-ctl
ENTRYPOINT /app/supervisord -c config/supervisor.yaml