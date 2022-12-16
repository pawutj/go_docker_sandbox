FROM golang:1.19-alpine

WORKDIR /app

# COPY . .

CMD CGO_ENABLED=0 go test -v ./...


#docker images
#docker image rm *


#docker build -t $name . 
#docker run $name