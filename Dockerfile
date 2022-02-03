FROM golang:latest

WORKDIR $GOPATH/src/fiber
COPY . $GOPATH/src/fiber
RUN GOPROXY="https://goproxy.io" GO111MODULE=on go build .

EXPOSE 8008
ENTRYPOINT ["./fiber"]