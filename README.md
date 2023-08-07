# Golang-Rabbitmq-Consumer

Golang-Rabbitmq-Consumer is a Golang app for Rabbitmq infrastructure validation.

## Installation

Use Golang 1.19. If you doesn't have Golang on your machine, install [gvm](https://github.com/moovweb/gvm#installing) then install Golang.

```bash
gvm install go1.19
gvm use go1.19
```

### Install dependencies

```bash
go mod tidy
```

## Environment Variables
To run this app, you must set this environment variables:
- RABBITMQ_HOST     - IP or DNS for rabbitmq service;
- RABBITMQ_PORT     - Port number for rabbitmq service;
- RABBITMQ_VHOST    - Virtual host for rabbitmq service;
- RABBITMQ_USERNAME - Username for rabbitmq service;
- RABBITMQ_PASSWORD - Password for rabbitmq service;
- RABBITMQ_QUEUE    - Queue name to consume;

## Running locally

```bash
go run main.go
```

## Build Container Image
### Using Docker
```bash
docker build -f Containerfile -t einstein/devops/golang-rabbitmq-consumer:0.0.1 .
```
### Using Buildah
```bash
buildah bud -f Containerfile -t einstein/devops/golang-rabbitmq-consumer:0.0.1 .
```

## Running Container Image Locally
### Using Docker
```bash
docker container run --rm --name golang-consumer -it -e RABBITMQ_HOST=172.16.10.2 -e RABBITMQ_PORT=5672 -e RABBITMQ_VHOST=testvhost -e RABBITMQ_USERNAME=guest -e RABBITMQ_PASSWORD=guest -e RABBITMQ_QUEUE=devops.test einstein/devops/golang-rabbitmq-consumer:0.0.1
```
### Using Podman
```bash
podman container run --rm --name golang-consumer -it -e RABBITMQ_HOST=172.16.10.2 -e RABBITMQ_PORT=5672 -e RABBITMQ_VHOST=testvhost -e RABBITMQ_USERNAME=guest -e RABBITMQ_PASSWORD=guest -e RABBITMQ_QUEUE=devops.test einstein/devops/golang-rabbitmq-consumer:0.0.1
```