package main

import (
	"database/sql"
	"github.com/boourns/dbutil"
	"go-gql-typescript-example/graph"
	"go-gql-typescript-example/graph/generated"
	"go-gql-typescript-example/graph/model"
	"go-gql-typescript-example/lib/migrations"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
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
		http.ServeFile(w, r, "./static/" + r.URL.Path)
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
	if err != nil { panic(err) }

	migrations.DefineMigration(db, 2, CreateTodoMigration)
	if err != nil { panic(err) }

	if err != nil {
		panic(err)
	}
	return db
}

func CreateUserMigration(tx dbutil.DBLike) error {
	return model.CreateUserTable(tx)
}

func CreateTodoMigration(tx dbutil.DBLike) error {
	return model.CreateTodoTable(tx)
}