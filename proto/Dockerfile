FROM golang:1.20.4-bullseye
RUN apt-get update -y \
    && apt-get upgrade \
    && apt-get install -y \
    wget \
    unzip

WORKDIR /tmp
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v23.1/protoc-23.1-linux-aarch_64.zip \
    && unzip protoc-23.1-linux-aarch_64.zip -d protoc
    && mv protoc/bin/protoc /usr/local/bin/protoc

# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
