
# Simple script to generate the protobuf definitions
# Used internally for pony
# protoc --go_out=. proto/api.proto
# protoc --go_out=pkg/ --go_opt=paths=source_relative \
#     --go-grpc_out=pkg/ --go-grpc_opt=paths=source_relative \
#     proto/api.proto

protoc --go_out=pkg/ --go-grpc_out=pkg/ proto/api.proto
