FROM golang:1.21.5

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -v -a -installsuffix cgo -o finance ./

EXPOSE 8080

CMD [ "./finance" ]