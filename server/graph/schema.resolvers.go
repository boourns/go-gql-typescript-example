package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"database/sql"
	"fmt"
	"go-gql-typescript-example/graph/generated"
	"go-gql-typescript-example/graph/model"
	"strconv"
	"strings"

	"github.com/boourns/dblib"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	// do whatever authorization of context or associated objects

	// input validation
	if input.Text == "" {
		return nil, fmt.Errorf("todo text must not be blank")
	}

	uid, e := strconv.ParseInt(input.UserID, 10, 64)
	if e != nil {
		return nil, fmt.Errorf("userId must be an integer")
	}

	todo := &model.Todo{
		Done:   false,
		Text:   input.Text,
		UserID: uid,
	}

	// use a transaction just for demonstration purposes even though we're doing a single SQL statement
	err := dblib.Transact(Database, func(tx *sql.Tx) error {
		// if this inner function returns an error, the entire block is rolled back

		return todo.Insert(Database)
	})

	return todo, err
}

func (r *mutationResolver) CreateUser(ctx context.Context, name string) (*model.User, error) {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		return nil, fmt.Errorf("name must not be blank")
	}
	user := &model.User{
		Name: name,
	}
	err := dblib.Transact(Database, func(tx *sql.Tx) error {
		// if this inner function returns an error, the entire block is rolled back

		return user.Insert(Database)
	})

	return user, err
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return model.SelectTodo(Database, "")
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return model.SelectUser(Database, "")
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return obj.User(Database)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
