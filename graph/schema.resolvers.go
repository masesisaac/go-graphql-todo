package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/masesisaac/go-graphql-todo/db"
	"github.com/masesisaac/go-graphql-todo/graph/generated"
	"github.com/masesisaac/go-graphql-todo/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo, err := db.AddTodo(input.Text)
	return convertTodoType(todo), err
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (bool, error) {
	done, err := db.DeleteTodo(id)
	return done, err
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, id string, done bool) (*model.Todo, error) {
	todo, err := db.UpdateTodoDone(id, done)
	return convertTodoType(todo), err
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	todo, err := db.GetTodo(id)
	return convertTodoType(todo), err
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todos, err := db.GetTodos()
	var _todos []*model.Todo
	for _, todo := range *todos {
		_todos = append(_todos, convertTodoType(&todo))
	}

	return _todos, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func convertTodoType(todo *db.Todo) *model.Todo {
	if todo == nil {
		return nil
	}
	return &model.Todo{
		ID:   fmt.Sprint(todo.ID),
		Text: todo.Text,
		Done: todo.Done,
	}
}
