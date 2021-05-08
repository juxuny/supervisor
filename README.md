Supervisor
=================


### Install

```shell
go get github.com/juxuny/supervisor
go install github.com/juxuny/supervisor/cmd/supervisord
go install github.com/juxuny/supervisor/cmd/supervisorctl
```

### Generate Certification

```shell
rm *.pem

# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=CN/ST=Guangdong/L=Dongguan/O=Root/OU=Root/CN=*.juxuny.com/emailAddress=juxuny@163.com"

echo "CA's self-signed certificate"
openssl x509 -in ca-cert.pem -noout -text

# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=CN/ST=Guangdong/L=Dongguan/O=Server/OU=Server/CN=*.juxuny.com/emailAddress=juxuny@163.com"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 -req -in server-req.pem -days 365 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf

echo "Server's signed certificate"
openssl x509 -in server-cert.pem -noout -text
#openssl x509 -req -in client-req.pem -days 365 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem -extfile client-ext.cnf
```

### Configuration

1. config/supervisor.yaml
```yaml
supervisor:
  proxy_image: "juxuny/supervisor:proxy-v0.0.1"
  control_port: 50060
  docker_host: unix:///var/run/docker.sock
  store_dir: tmp
  cert_file: "cert/server-cert.pem"
  cert_key_file: "cert/server-key.pem"
```

2. tmp/deploy-web.yaml
```yaml
deploy:
  service_port: 8080
  proxy_port: 8090
  name: web
  image: juxuny/go-web
  tag: latest
  mounts:
  - host_path: ./tmp
    mount_path: /html
  env_data: QUNDRVNTX0tFWT0iMTIzIDQ1NiIKU0VDUkVUPSAxMjM0NTc3OA==
  envs:
  - key: PORT
    value: "8080"
  version: 9
  entrypoint: []
  health_check:
    type: 0
    path: "/"
    port: 8080
```

* health_check.type = 0: `HTTP GET` request
* health_check.type = 1: TCP Connection Test 


### Start Service

```shell
supervisord -c config/supervisor.yaml
```

### Deploy Web Server

```shell
# apply deployment config
supervisor-ctl apply --host 127.0.0.1:50060 --cert-file cert/ca-cert.pem --file tmp/deploy-web.yaml --timeout 300

# stop service
supervisor-ctl stop --host 127.0.0.1:50060 --cert-file cert/ca-cert.pem --name web
```


### Deploy Supervisord Via docker-compose

```yaml
version: '3.4'

x-default: &default
  logging:
    options:
      max-size: "5M"
      max-file: "5"

services:
  supervisord:
    environment:
      HOST_PWD: ${PWD}
    image: juxuny/supervisor:srv-v0.0.11
    restart: always
    volumes:
      - ./cert:/app/cert
      - ./tmp:/app/tmp
      - ./config:/app/config
      - /var/run/docker.sock:/var/run/docker.sock
    ports:
      - "50060:50060"
    <<: *default

```