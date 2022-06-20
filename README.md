# Producer RMQ - Golang

Go based counterpart to `ktor-rmq producer` server.

The easiest way to get `RabbitMQ` locally on machines is to pull the [docker image](https://hub.docker.com/_/rabbitmq).
There are two options:
- Base image: `docker pull rabbitmq:3`
- Base image with `Management plugin`: `docker pull rabbitmq:3-management`

The management plugin provides a web based dashboard to inspect queue information. Accessed at port `15672`.

Since we can expose docker ports and map them to `localhost` ports, in development terms our apps can locally interact with
the `RabbitMQ` framework, and later they can be deployed to their own containers.

## Run the application

*Instead of hardcoding values, this application relies on environment variables which are easily passed to docker containers, and are also 
easily retreived in code.*

1. Pull `RabbitMQ` image from dockerhub
2. Create local docker network `docker network create <network_name>`
3. Run `RabbitMQ` image `docker run -d --rm --net <network_name> --hostname <host_name> --name <container_name> <image_name>`
   1. `<host_name>` is **important** as our applications need to know this name to work.
4. If downloaded base `RabbitMQ` image, enable `management` plugin
   ```
    # Run docker image
   docker run -d --rm --net <network_name> --hostname <host_name> --name <container_name> <image_name>
   
    # Access local container terminal
   docker exec -it <container_name> bash
   
    # Enable management plugin
   rabbitmq-plugins enable rabbitmq_management
    ```
5. Build and run docker image for Producer
    ```
    # Open terminal in project root folder
   docker build --no-cache --tag <name>:<tag> .
   
   # Run image in container - Make sure to assign same network
   docker run -d --rm --net <network_name> -e RABBIT_HOST=<host_name> -e RABBIT_PORT=5672 -e RABBIT_USERNAME=guest -e RABBIT_PASSWORD=guest -p 80:80 --name <container_name> <name>:<tag>
   
   # Container exposes port 80, send POST request to localhost:80
    ```
