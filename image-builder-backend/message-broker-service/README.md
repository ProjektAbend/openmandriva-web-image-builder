# MessageBrokerService - Open Mandriva Web Image Builder

## About this service
- queues


## How to start the service
Start the service inside a Docker container using:
```shell
docker-compose up -d
```

After a few seconds you can access the RabbitMQ Management UI: http://localhost:15672.

Use the username and password from the `docker-compose.yml`.


### Running Tests
Test sending or consuming messages from the queue with the scripts located under `/test`.
In order to run the test scripts you first have to install the dependencies with
```shell
go mod tidy
```
within the `/test` directory.

Execute `test_send.go` to send a hello world message to the queue.
```shell
go run test_send.go
```
You should see the message in the Management UI afterwards.
The message counter should be increased by 1.

Execute `test_consume.go` to consume the message you send earlier.
```shell
go run test_consume.go
```
After the message was dequeued the message counter should be at 0 again.