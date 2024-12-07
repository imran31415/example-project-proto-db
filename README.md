# About
# Automating the Database Code Layer Leveraging Protobuf Annotations and Code Generation  

---

## Overview  
We strive to optimize the developer experience by integrating the Application and Database layers seamlessly. Specifically, we propose a solution where the **Database layer** is described **upfront** within the **API specification** itself. By establishing a **single source of truth**, we can leverage code generation tools to automatically generate much of the code that developers currently implement manually for the database layer.  

---

## Problem Statement  
When implementing an API and Database, developers often deal with **multiple "source of truth" specs** – sometimes as many as **three layers**:  
1. **`.proto` files** defining the API schema,  
2. **SQL scripts** defining the database tables, and  
3. **Business logic** implementing the API-to-Database calls.  

This fragmented workflow introduces complexity and redundancy. For example, adding a new column to a table or API involves:  
1. Updating the `.proto` file,  
2. Running `protoc` to regenerate stubs,  
3. Updating database schema models,  
4. Running SQL migrations to modify the database,  
5. Updating the API implementation to handle the new field.  

This process violates the **"Don't Repeat Yourself" (DRY)** principle and creates **error-prone manual steps** across multiple layers.  

---

## Solution  

We propose an **API-first approach** where the **API spec** serves as the **single source of truth** for both the API and database layers. By augmenting Protobuf definitions with **custom database annotations**, we can capture database schema and constraints **directly** in the `.proto` file.  

This enriched Protobuf specification enables us to use **code generation** to automate the following:  

1. **Generate the SQL Database Schema**  
2. **Validate the SQL Spec Against a Target Database**  
3. **Generate Basic CRUD Logic** (Go models with database interactions)  
4. **Generate gRPC Server Implementations**  
5. **Generate SQL Migrations** Between Two Versions of Protobuf Files  

---

## New Workflow  

The development process is simplified to:  

1. **Update the Protobuf File**: Add/modify fields with DB annotations.  
2. **Run Code Generation**: Use the translator to generate:  
   - SQL schema,  
   - Go models, and  
   - gRPC API implementation.  
3. **Iterate on the Generated Implementation**: Add any necessary business logic.  

This eliminates redundant tasks, reduces the cognitive load on developers, and ensures consistency between the API and database layers.  

---

## Example  

We have created a **prototype** to demonstrate this approach:  

### Step 1: Define Protobuf With Annotations  

Here is an example of a `User` entity:  

\`\`\`proto
syntax = "proto3";

package auth;

import "protobuf-db/db-annotations.proto";

// Message for the User entity
message User {
  int32 id = 1 [
    (db_annotations.db_column) = "id",
    (db_annotations.db_primary_key) = true,
    (db_annotations.db_column_type) = DB_TYPE_INT,
    (db_annotations.db_constraints) = DB_CONSTRAINT_NOT_NULL
  ];

  string username = 2 [
    (db_annotations.db_column) = "username",
    (db_annotations.db_column_type) = DB_TYPE_VARCHAR,
    (db_annotations.db_constraints) = DB_CONSTRAINT_UNIQUE
  ];

  string email = 3 [
    (db_annotations.db_column) = "email",
    (db_annotations.db_column_type) = DB_TYPE_VARCHAR,
    (db_annotations.db_constraints) = DB_CONSTRAINT_NOT_NULL
  ];

  string password = 4 [
    (db_annotations.db_column) = "password",
    (db_annotations.db_column_type) = DB_TYPE_VARCHAR
  ];
}
\`\`\`

---

### Step 2: Generate SQL Schema and Go Models  

By running the translator, the SQL table statements and Go models are generated automatically:  

**Run the Translator Code**:  

\`\`\`go
t.GenerateModels("../generated_models", []proto.Message{
    &auth.User{}, 
    &auth.Role{}, 
    &auth.UserRole{},
})
\`\`\`

**Generated SQL Table Statement**:  

\`\`\`sql
CREATE TABLE users (
    id INT PRIMARY KEY NOT NULL,
    username VARCHAR(255) UNIQUE,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255)
);
\`\`\`

**Generated Go Model**:  

\`\`\`go
type User struct {
    ID       int32  `db:"id" json:"id"`
    Username string `db:"username" json:"username"`
    Email    string `db:"email" json:"email"`
    Password string `db:"password" json:"password"`
}
\`\`\`

---

### Step 3: Generate gRPC Server Implementation  

The translator also generates a skeleton gRPC server implementation to handle CRUD operations:  

**Generated gRPC Server Code**:  

\`\`\`go
func (s *Server) CreateUser(ctx context.Context, req *auth.User) (*auth.User, error) {
    // Example implementation: Insert into the database
    query := \`INSERT INTO users (username, email, password) VALUES (?, ?, ?)\`
    _, err := s.db.ExecContext(ctx, query, req.Username, req.Email, req.Password)
    if err != nil {
        return nil, err
    }
    return req, nil
}
\`\`\`

---

## The Power of a Single Source of Truth  

By consolidating the database schema and API definitions into a **single `.proto` file** with annotations, we achieve:  
- **Reduced Boilerplate**: Auto-generated SQL schema, Go models, and gRPC code.  
- **Consistency**: A single authoritative source eliminates inconsistencies between API and database layers.  
- **Faster Development**: Fewer manual steps mean faster iterations and fewer errors.  

---

## Future Possibilities  

This tool is a foundation for further automation and improvements:  

1. **Automatic Caching Layer**: Generate cache configurations and logic directly from the `.proto` file.  
2. **Advanced SQL Migrations**: Auto-generate SQL migration scripts by comparing two versions of a `.proto` file.  
3. **Integration Into CI/CD Pipelines**: Validate schema changes, ensuring that new fields or constraints are valid SQL operations.  
4. **Enhanced Application Logic**: Use Protobuf as a central schema for even more layers of the application stack.  

---

## Conclusion  

By leveraging **Protobuf annotations** and code generation, we streamline the development process by consolidating the API and database definitions into a single **source of truth**. This approach reduces redundancy, ensures consistency, and accelerates development while maintaining the flexibility to extend into more advanced automation in the future.  

Check out the repositories:  
- [protobuf-db](https://github.com/imran31415/protobuf-db)  
- [proto-db-translator](https://github.com/imran31415/proto-db-translator)  

Give it a try and see how much time you can save in your workflow! 🚀  





## Build

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
