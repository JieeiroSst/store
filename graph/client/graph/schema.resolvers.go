package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/JIeeiroSst/store/graph/client/graph/generated"
	"github.com/JIeeiroSst/store/graph/client/graph/model"
)

func (r *mutationResolver) SendMail(ctx context.Context, email *string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddFriend(ctx context.Context, id *int) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateFeebBack(ctx context.Context, input *model.InputFeedBacks) (*model.ResultCheck, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) News(ctx context.Context) (*model.News, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tags(ctx context.Context) (*model.Tags, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) NewTags(ctx context.Context) (*model.NewTag, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) New(ctx context.Context, id *int) (*model.News, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tag(ctx context.Context, id *int) (*model.Tags, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) NewTag(ctx context.Context, id *int) (*model.NewTag, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Profile(ctx context.Context, id *int) (*model.Profile, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }