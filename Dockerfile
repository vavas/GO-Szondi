FROM golang:1.8


# Load dependencies
RUN go get -u github.com/kardianos/govendor

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
