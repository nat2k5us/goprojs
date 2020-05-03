FROM golang:alpine
ADD . /go/src/github.com/nat2k5us/gohello
RUN go install github.com/nat2k5us/gohello
CMD ["/go/bin/gohello"]
EXPOSE 3000