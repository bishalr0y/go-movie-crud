FROM golang:1.22.2-alpine3.19

WORKDIR /home/app

EXPOSE 8080

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . . 

RUN go build -v -o ./bin/

CMD [ "./bin/go-movie-crud" ]