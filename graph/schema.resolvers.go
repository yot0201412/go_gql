package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/yot0201412/go_gql/graph/model"
	"github.com/yot0201412/go_gql/sqlc"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	param := sqlc.CreateTodoParams{
		UserID: ToInt32(input.UserID),
		Text:   input.Text,
	}
	if err := r.Repo.CreateTodo(ctx, param); err != nil {
		return nil, err
	}
	return &model.Todo{
		ID:   string(1),
		Text: input.Text,
		Done: false,
	}, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	if err := r.Repo.CreateUser(ctx, input.Name); err != nil {
		return nil, err
	}
	return &model.User{
		ID:   ToString(1),
		Name: input.Name,
	}, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todos, _ := r.Repo.GetTodosByUserID(ctx, sql.NullInt32{0, false})
	var result []*model.Todo
	for _, todo := range todos {
		result = append(result, &model.Todo{
			ID:     ToString(todo.ID),
			Text:   todo.Text,
			Done:   todo.Done,
			UserID: ToString(todo.UserID),
		})
	}
	return result, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, input string) (*model.User, error) {
	user, err := r.Repo.GetUser(ctx, ToInt32(input))
	if err != nil {
		return nil, err
	}
	gUser := &model.User{
		ID:   ToString(user.ID),
		Name: user.Name,
	}
	return gUser, nil
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	user, err := r.Repo.GetUser(ctx, ToInt32(obj.UserID))
	fmt.Println(user)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:   ToString(user.ID),
		Name: user.Name,
	}, nil
}

// Todos is the resolver for the todos field.
func (r *userResolver) Todos(ctx context.Context, obj *model.User) ([]*model.Todo, error) {
	todos, err := r.Repo.GetTodosByUserID(ctx,
		sql.NullInt32{
			ToInt32(obj.ID),
			ToInt32(obj.ID) != 0,
		})
	if err != nil {
		return nil, err
	}
	var result []*model.Todo
	for _, todo := range todos {
		result = append(result, &model.Todo{
			ID:     ToString(todo.ID),
			Text:   todo.Text,
			Done:   todo.Done,
			UserID: ToString(todo.UserID),
		})
	}
	return result, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
