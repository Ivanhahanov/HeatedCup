# HeatedCup
Simple GO API for managing cup using mqtt protocol

### Simple run
For correct work, you need mqtt broker, for example _Mosquitto_. 

You can run app using follow command:
```shell
go run Api/main.go
```

Example request:
```shell
curl --header "Content-Type: application/json" -XPOST  --data '{"command":"on"}' http://localhost:8000
```
### Containers
#### Build from Dockerfile
Build image:
```shell
docker build -t heated_cap Api/
```
Run container:
```shell
docker run --rm -p8000:8000 heated_cap
```
localhost:8000 
#### Docker Compose
```shell
docker-compose up -d
```