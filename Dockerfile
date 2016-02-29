FROM golang
ADD . /go/src/github.com/golang/example/web
RUN go install github.com/golang/example/web
ENTRYPOINT /go/bin/web
EXPOSE 80
