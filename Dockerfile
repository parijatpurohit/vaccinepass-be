FROM golang:1.14

WORKDIR /go/src

# Install protoc
RUN apt-get update && apt-get install unzip
RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.9.1/protoc-3.9.1-linux-x86_64.zip && \
    unzip -o protoc-3.9.1-linux-x86_64.zip -d /usr/local bin/protoc && \
    unzip -o protoc-3.9.1-linux-x86_64.zip -d /usr/local include/* && \
    rm -rf protoc-3.9.1-linux-x86_64.zip

# install deps
RUN go get -u github.com/golang/protobuf/protoc-gen-go && \
    go get -u gopkg.in/yaml.v2 && \
    go get -u gorm.io/gorm && \
    go get -u google.golang.org/grpc && \
    go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
    go get -u golang.org/x/net/http2 && \
    go get -u golang.org/x/sys/unix

RUN mkdir -p github.com/parijatpurohit
# Setup sidecar generator
RUN git clone https://github.com/parijatpurohit/sidecar-sql.git github.com/parijatpurohit/sidecar-sql
RUN chmod +x github.com/parijatpurohit/sidecar-sql/generator.sh
WORKDIR /go/src/github.com/parijatpurohit/sidecar-sql

# Generate code
RUN git clone https://github.com/parijatpurohit/vaccinepass-be.git /go/src/github.com/parijatpurohit/vaccinepass
RUN ./generator.sh /go/src/github.com/parijatpurohit/vaccinepass

WORKDIR /go/src/github.com/parijatpurohit/vaccinepass

CMD ["go", "run", "cmd/main.go"]