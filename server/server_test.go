package main

import (
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"

	"go-gql-typescript-example/graph"
	"go-gql-typescript-example/graph/generated"
	"go-gql-typescript-example/graph/model"

	"github.com/stretchr/testify/require"

	"log"
	"os"
	"testing"
)

var user *model.User

func TestMain(m *testing.M) {
	graph.Database = openAndMigrateDatabase("file::memory:?cache=shared")

	user = createUser("bob")

	if user == nil {
		log.Fatalf("failed to create default user")
	}

	os.Exit(m.Run())
}

func TestFetchTodos(t *testing.T) {
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})))

	t.Run("Fetch user list", func(t *testing.T) {
		// response struct must match the shape of the graphql JSON response
		var resp struct {
			Users []*model.User `json:"users"`
		}

		query := `
query {
  users {
    name
  }
}
`
		c.MustPost(query, &resp, func(bd *client.Request) {
			bd.Variables = make(map[string]interface{})
			bd.Variables["name"] = "bob"
		})

		require.Equal(t, 1, len(resp.Users))

	})

}

func createUser(name string) *model.User { // response struct must match the shape of the graphql JSON response
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})))

	var resp struct {
		CreateUser *model.User `json:"createUser"`
	}

	query := `
	mutation($name: String!) {
		createUser(name: $name) {
			id
			name
		}
	}`

	c.MustPost(query, &resp, func(bd *client.Request) {
		bd.Variables = make(map[string]interface{})
		bd.Variables["name"] = "bob"
	})

	return resp.CreateUser
}
