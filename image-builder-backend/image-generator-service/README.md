# ImageGeneratorService - Open Mandriva Web Image Builder

## About this service
This service fetches the imageConfigs created by the `api-gateway-service` from RabbitMQ and generates
the image using [os-image-builder](https://github.com/OpenMandrivaSoftware/os-image-builder/tree/master).

After generating, the image is passed to the `image-storage-service` in order to be stored and then downloaded by the user.


## How to start the service
> [!WARNING]  
> You have to run this service with `--mock` in order to run it on windows

This service is using `os-image-builder` which cannot run on Windows.
Run the service in `MOCK MODE` to use it on Windows.


### Run in MOCK MODE
Run this service in MOCK MODE with this:
```shell
go run cmd/main.go --mock
```
With `--mock` it will create a fake .iso file instead of using the real `os-image-builder`.
If you use GoLand you can use this run configuration instead of the command above:
`start mock image-generator-service`.


### Run the real thing
You can run the real thing if you develop on a Linux system.
The service is supposed to run inside a Docker Container, even during development.
Inside the Docker Container, the `os-image-builder` is automatically installed.


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

If you use GoLand you can use this run configuration instead of the steps above:
`start docker image-generator-service`. This will build the code and start the docker container automatically.
If you make changes to the code, just re-run this run configuration.