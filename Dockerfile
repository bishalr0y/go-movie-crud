FROM golang:1.22.2-alpine3.19 as builder

WORKDIR /home/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . . 

RUN go build -v -o ./bin/


# Multistage build

FROM gcr.io/distroless/static-debian12:latest

COPY --from=builder /home/app/bin/go-movie-crud .

EXPOSE 8080

CMD [ "./go-movie-crud" ]