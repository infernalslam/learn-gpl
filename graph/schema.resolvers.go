package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"github.com/infernalslam/gql-simple/graph/generated"
	"github.com/infernalslam/gql-simple/graph/model"
)

var todoDB []*model.Todo
var userDB = []model.User{{
	ID:   "1",
	Name: "tak",
}}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	rand.Seed(time.Now().UnixNano())

	var user model.User
	for _, u := range userDB {
		if u.ID == input.UserID {
			user = u
		}
	}

	obj := model.Todo{
		ID:   strconv.Itoa(rand.Intn(1000-0+1) + 0),
		Text: input.Text,
		Done: true,
		User: &user,
	}

	todoDB = append(todoDB, &obj)

	return &obj, nil

}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return todoDB, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
