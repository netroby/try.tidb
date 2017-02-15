FROM golang:rc

ENV GOPATH /go
ADD . /go/src/try.tidb
RUN apt update; \
    apt install mysql-client; \
    cd /go/src/try.tidb;\
    go get -u -v ...;\
    go build .
    
CMD /go/src/try.tidb/try.tidb
