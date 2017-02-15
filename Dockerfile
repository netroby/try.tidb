FROM golang:rc

ENV GOPATH /go
ADD ./ /go/src/try.tidb
RUN apt update; \
    apt install  -y mysql-client; \
    cd /go/src/try.tidb;\
    go get -u -v ...;\
    go build .
WORKDIR /go/src/try.tidb    
CMD ./try.tidb
