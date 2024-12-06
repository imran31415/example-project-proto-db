

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


This is an example project to show how to use protobuf-db and proto-db-translator. 

1. See the proto/* files for the messages and annotations
2. Example usage is in generate/main.go
```go
func main() {
	conn := db.DefaultMysqlConnection()
	conn.DbName = "example_project_proto_db"
	t := translator.NewTranslator(conn)
	// .GenerateModels does the following:
	//   1. Takes each of the protos and generate the SQL create table statement,
	//   2. Execute the statements to generate all the tables based on the protobuf annotations
	///  3. With the created SQL tables, generate the Go CRUD models.
	t.GenerateModels("../generated_models", []proto.Message{&auth.User{}, &auth.Role{}, &auth.UserRole{}})
}
```
