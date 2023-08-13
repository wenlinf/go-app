FROM golang:latest

WORKDIR /app

COPY . .

COPY my_app.go ./

RUN go build -o my-app

CMD ["./my-app"]
