FROM golang:1.10-alpine3.7 as source
ENV GOPATH=/go
RUN go get -u github.com/bannzai/constructor/

FROM golang:1.10-alpine3.7 as binary
ENV GOPATH=/go
ENV BIN_PATH=$GOPATH/bin/constructor/
COPY --from=source $BIN_PATH $BIN_PATH
WORKDIR BIN_PATH

CMD ["./constructor"]


