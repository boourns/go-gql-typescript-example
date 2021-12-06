# go-gql-typescript-example
Example of strongly typed go/graphql/typescript web application

# Overview
This is an example web application.

On the server it uses:
- Golang
- https://gqlgen.com to generate the graphql API from a graphql schema
- https://github.com/boourns/scaffold to generate a SQLite interface from Go structs

The client is built with:
- Typescript
- Webpack
- https://www.graphql-code-generator.com to generate typescript types from the graphql schema

# Install and Build
```bash
git clone git@github.com:boourns/go-gql-typescript-example.git
cd go-gql-typescript-example/js
yarn
yarn build
cd ../server
go get
go run server.go
```

# Generating code from GraphQL Schema
Edit `./server/graph/graphql.schemas`
run `./tools/generate.sh` to regenerate the go and typescript bindings

# Generating Model layer
A basic SQLite interface is generated using Scaffold.  If you modify the structs, you need to regenerate the model layer, and write a migration to update the existing schema.
To regenerate the SQL layer
```bash
cd server/graph/model
go get github.com/boourns/scaffold@549f411c4bac527df427315ba98913c6613c3bdf
go run github.com/boourns/scaffold model -in=user.go
go run github.com/boourns/scaffold model -in=todo.go
```

