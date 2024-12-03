npm install --save-dev ts-proto
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
mkdir ts_out



export PATH=$PATH:$HOME/go/bin
protoc -I . \
       --plugin=protoc-gen-ts_proto=$(pwd)/node_modules/.bin/protoc-gen-ts_proto \
       --ts_proto_out=./ts_out \
       --ts_proto_opt=esModuleInterop=true,outputEncodeMethods=false,outputJsonMethods=false,outputClientImpl=false \
       --go_out=./ \
       --go-grpc_out=./ \
       --grpc-gateway_out=./ \
       --proto_path=./protobuf-db/proto \
       --proto_path=./proto \
       ./proto/auth.proto