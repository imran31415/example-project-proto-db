package main

import (
	auth "github.com/imran31415/example-project-proto-db/auth"
	configGenerator "github.com/imran31415/proto-db-translator/config_generator"
	translator "github.com/imran31415/proto-db-translator/translator"
	db "github.com/imran31415/proto-db-translator/translator/db"

	"google.golang.org/protobuf/proto"
)

func main() {
	conn := db.DefaultMysqlConnection()
	conn.DbName = "example_project_proto_db"
	t := translator.NewTranslator(conn)
	// .GenerateModels does the following:
	//   1. Takes each of the protos and generate the SQL create table statement,
	//   2. Execute the statements to generate all the tables based on the protobuf annotations
	///  3. With the created SQL tables, generate the Go CRUD models.
	t.GenerateModels("../generated_models", []proto.Message{&auth.User{}, &auth.Role{}, &auth.UserRole{}})

	// .GenerateConfig will create a /config directory with a config.go that supplies a valid config object to supply to the API
	configGenerator.GenerateConfig("../config")

}
