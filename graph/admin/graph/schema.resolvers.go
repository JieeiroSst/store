package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/JIeeiroSst/store/graph/admin/graph/generated"
	"github.com/JIeeiroSst/store/graph/admin/graph/model"
)

func (r *mutationResolver) CreateNews(ctx context.Context, input *model.InputNews) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTags(ctx context.Context, input *model.InputTags) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateNewTag(ctx context.Context, input *model.InputNewTag) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateFeebBack(ctx context.Context, input *model.InputFeedBacks) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateProfile(ctx context.Context, input *model.InputProfile) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMenues(ctx context.Context, input *model.InputMenues) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateNews(ctx context.Context, id *int, input *model.InputNews) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTags(ctx context.Context, id *int, input *model.InputTags) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateNewTag(ctx context.Context, id *int, input *model.InputNewTag) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateFeebBack(ctx context.Context, id *int, input *model.InputFeedBacks) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateProfile(ctx context.Context, id *int, input *model.InputProfile) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMenues(ctx context.Context, id *int, input *model.InputMenues) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteNews(ctx context.Context, id *int) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTags(ctx context.Context, id *int) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteNewTag(ctx context.Context, id *int) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCasbinRule(ctx context.Context, id *int) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCategories(ctx context.Context, id *int) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteContact(ctx context.Context, id *int) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteFeebBack(ctx context.Context, id *int) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteProduct(ctx context.Context, id *int) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteProfile(ctx context.Context, userID *int) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMenues(ctx context.Context, id *int) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteProductCategories(ctx context.Context, id *int) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteSliders(ctx context.Context, id *int) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CheckNews(ctx context.Context, id *int) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PublicNew(ctx context.Context, id *int) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SendMail(ctx context.Context, email *string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddFriend(ctx context.Context, id *int) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) News(ctx context.Context) (*model.NewsConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tags(ctx context.Context) (*model.TagsConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) NewTags(ctx context.Context) (*model.NewTagsConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Feedbacks(ctx context.Context) (*model.FeedBackConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Profile(ctx context.Context) (*model.ProfileConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SystemConfig(ctx context.Context) (*model.SystemConfigConnection, error) {
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

func (r *queryResolver) Feedback(ctx context.Context, id *int) (*model.FeedBacks, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ProfileID(ctx context.Context, id *int) (*model.Profile, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Systemconfig(ctx context.Context, id *int) (*model.SystemConfig, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) RefreshToken(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }