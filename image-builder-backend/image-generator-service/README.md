# ImageGeneratorService - Open Mandriva Web Image Builder

## About this service
what does it do
- reads imageConfig from rabbitmq
- generates image
- gives update about progress of the generating image
- sends image to image storage service

## How to start the service
service runs inside docker, even during development
https://blog.jetbrains.com/go/2021/04/30/how-to-use-docker-to-compile-go-from-goland/#enter-the-docker-container
