# ApiGatewayService - Open Mandriva Web Image Builder

## About this service
This is the entry point for the Frontend to send requests to.
It handles the requests from the Frontend and makes sure it arrives at the correct destination.
It communicates with the other microservices in the Backend through RabbitMQ.

The Service is written in Go using the Gin Framework.
The API controller is generated using the openApi specification and [oapi codegen](https://github.com/deepmap/oapi-codegen).


## How to start the service
Before starting the service, the API controller has to be generated from the `openapi.yml` first.

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
If you use goland you can also use this run configuration: `start api-gateway-service` instead.

## Known Issues

### Packages don't resolve in the IDE
If you are working with a Jetbrains IDE like Goland, you may have the problem that the packages which are 
imported at the beginning of every `.go` file cannot be resolved.
A reason could be that your IDE is not setup for a multi modular Go project.
The solution would be to activate `Go modules integration` in your IDE.
