package main

import (
	_ "github.com/mattn/go-sqlite3"

	"database/sql"
	"go-gql-typescript-example/graph"
	"go-gql-typescript-example/graph/generated"
	"go-gql-typescript-example/graph/model"
	"log"
	"net/http"
	"os"

	"github.com/boourns/dblib/migrations"

	_ "github.com/mattn/go-sqlite3"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/boourns/dblib"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	graph.Database = openAndMigrateDatabase("./todo.db")

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	http.HandleFunc("/js/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/"+r.URL.Path)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func openAndMigrateDatabase(filename string) *sql.DB {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		log.Fatal(err)
	}

	err = migrations.CreateMigrationsTable(db)
	if err != nil {
		log.Printf("%q\n", err)
	}

	err = migrations.DefineMigration(db, 1, CreateUserMigration)
	if err != nil {
		panic(err)
	}

	err = migrations.DefineMigration(db, 2, CreateTodoMigration)
	if err != nil {
		panic(err)
	}

	return db
}

func CreateUserMigration(tx dblib.DBLike) error {
	return model.CreateUserTable(tx)
}

func CreateTodoMigration(tx dblib.DBLike) error {
	return model.CreateTodoTable(tx)
}
