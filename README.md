

Run protoc referencing the path of the cloned repo `--proto_path=./protobuf-db/proto \`


```bash 
export PATH=$PATH:$HOME/go/bin
protoc -I . \
       --plugin=protoc-gen-ts_proto=$(pwd)/node_modules/.bin/protoc-gen-ts_proto \
       --go_out=./ \
       --go-grpc_out=./ \
       --grpc-gateway_out=./ \
       --proto_path=./protobuf-db/proto \
       --proto_path=./proto \
       ./proto/auth.proto

```
