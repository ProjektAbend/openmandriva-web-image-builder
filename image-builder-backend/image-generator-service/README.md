# ImageGeneratorService - Open Mandriva Web Image Builder

## About this service
what does it do
- reads imageConfig from rabbitmq
- generates image
- gives update about progress of the generating image
- sends image to image storage service

does not run on windows!! it uses os-image-builder
service runs inside docker, even during development -> open mandriva image


## How to start the service

run
```shell
    GOOS=linux GOARCH=amd64 go build -o build/image-generator-service cmd/main.go
```
it generates an executable for linux which will be started inside the docker container

you could also just run the Makefile:
```shell
make
```

After that, start the docker container:
```shell
  docker-compose up
```

If you use intelliJ you can use this run configuration instead of the steps above:
`start docker image-generator-service`. This will build the code and start the docker container

If you make changes to the code, just re-run this run configuration or follow the steps above.