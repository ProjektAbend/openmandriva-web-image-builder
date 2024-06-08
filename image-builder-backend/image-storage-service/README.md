# ImageStorageService - Open Mandriva Web Image Builder

## About the service
This service is responsible for the management of generated .iso files.
Other services use this service to upload or download files from it.
Every file gets an expiration date. When this date is reached, the file gets deleted.


## How to start the service
Before starting the service, the API controller has to be generated from the `./openapi.yml` first.

### Generate API from openapi.yml file
The `Makefile` automates the process of generating the API controller.
In this directory, just run:
```shell
make
```
after that you should see a new `.go` file inside _/cmd/api_: `interface.gen.go`

### Start server
When the API controller is generated, the server can be started:
```shell
go run main.go
```

If you use goland you can also use this run configuration: `start image-storage-service` instead.