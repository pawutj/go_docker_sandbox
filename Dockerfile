#docker build -t $name . 
#docker run $name

FROM golang:1.19-alpine

WORKDIR /app

COPY . .

CMD ["go","test","-v","./..."]