FROM golang:1.20-alpine

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o go-weather-bot ./cmd/main.go

CMD [ "./go-weather-bot" ]
