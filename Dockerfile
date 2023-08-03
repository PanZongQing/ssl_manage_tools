FROM golang:1.19.7

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/ssl_manager_tools
COPY . $GOPATH/src/github.com/ssl_manager_tools

RUN go build ./server/main.go

EXPOSE 8080
ENTRYPOINT [ "./main" ]
