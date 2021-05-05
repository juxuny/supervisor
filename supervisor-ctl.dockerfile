FROM golang:1.15.4 as builder
COPY . /src/
RUN cd /src && GOPROXY=https://goproxy.io go mod download
RUN cd /src/cmd/proxy && CGO_ENABLED=0 go build -o proxy
RUN cd /src/cmd/proxy-ctl && CGO_ENABLED=0 go build -o proxy-ctl
RUN cd /src/cmd/supervisor-ctl && CGO_ENABLED=0 go build -o supervisor-ctl

#FROM ineva/docker-envsubst:stable AS tools

# final stage
FROM docker:stable
RUN apk add --update --no-cache libintl gettext ca-certificates
WORKDIR /app
ENV PATH="/app:${PATH}"
COPY --from=builder /src/cmd/supervisor-ctl/supervisor-ctl /app/supervisor-ctl
COPY --from=builder /src/cmd/proxy-ctl/proxy-ctl /app/proxy-ctl
#COPY --from=tools /usr/bin/envsubst /usr/bin/envsubst