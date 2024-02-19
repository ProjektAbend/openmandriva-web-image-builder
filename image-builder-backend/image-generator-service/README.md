# ImageGeneratorService - Open Mandriva Web Image Builder

## About this service
This service fetches the imageConfigs created by the `api-gateway-service` from RabbitMQ and generates
the image using [os-image-builder](https://github.com/OpenMandrivaSoftware/os-image-builder/tree/master).

After generating, the image is passed to the `image-storage-service` in order to be stored and then downloaded by the user.


## How to start the service
> [!WARNING]  
> This service does not run locally on Windows

This service is using `os-image-builder` which cannot run on Windows.
The service is supposed to run inside a Docker Container, even during development.


First build an executable for linux which will be started inside the Docker Container:
```shell
GOOS=linux GOARCH=amd64 go build -o build/image-generator-service cmd/main.go
```

Running the Makefile will do the same thing:
```shell
make
```

After that, start the docker container:
```shell
docker-compose up -d
```
This will create a Docker Container based on the image which is defined in `Dockerfile`.
Before starting, the executable `build/image-generator-service` will be copied inside the container by mounting
the /build directory.

If you make changes to your code just repeat the steps above.

### Easy Start in GoLand
If you use GoLand you can use this run configuration instead of the steps above:
`start docker image-generator-service`. This will build the code and start the docker container automatically.
If you make changes to the code, just re-run this run configuration.