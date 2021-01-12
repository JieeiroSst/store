package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	db "github.com/JIeeiroSst/store/component"
	"github.com/JIeeiroSst/store/graph/generated"
	"github.com/JIeeiroSst/store/graph/model"
	"github.com/JIeeiroSst/store/models/news"
	"github.com/JIeeiroSst/store/models/tags"
	"github.com/JIeeiroSst/store/utils/logger"
	"github.com/JIeeiroSst/store/utils/tranfer"
	"github.com/manucorporat/try"
)

func (r *mutationResolver) CreateNews(ctx context.Context, input *model.InputNews) (*bool, error) {
	check := false
	try.This(func() {
		data := &news.News{
			Title:           tranfer.DeferString(input.Title),
			MetaTitle:       tranfer.DeferString(input.MetaTitle),
			Description:     tranfer.DeferString(input.Description),
			Image:           tranfer.DeferString(input.Image),
			CategoryId:      tranfer.DeferInt(input.CategoryID),
			Detail:          tranfer.DeferString(input.Detail),
			CreatedBy:       tranfer.DeferString(input.CreatedBy),
			ModifiedData:    tranfer.DeferString(input.ModifiedData),
			ModifiedBy:      tranfer.DeferString(input.ModifiedBy),
			MetaKeyWord:     tranfer.DeferString(input.MetaKeyWord),
			MetaDescription: tranfer.DeferString(input.MetaDescription),
			TagId:           tranfer.DeferInt(input.TagID),
			Content: 		 tranfer.DeferString(input.Content),
		}

		db.GetConn().Create(&data)
		check = true

	}).Finally(func() {
		logger.Log.Info("this must be printed after the catch")
	}).Catch(func(e try.E) {
		logger.Log.Info(e)
	})
	return &check, nil
}

func (r *mutationResolver) CreateTags(ctx context.Context, input *model.InputTags) (*bool, error) {
	check := false
	try.This(func() {
		data := &tags.Tags{
			Name:  tranfer.DeferString(input.Name),
			NewId: tranfer.DeferInt(input.NewID),
		}

		db.GetConn().Create(&data)
		check = true

	}).Finally(func() {
		logger.Log.Info("this must be printed after the catch")
	}).Catch(func(e try.E) {
		logger.Log.Info(e)
	})
	return &check, nil
}

func (r *mutationResolver) CreateNewTag(ctx context.Context, input *model.InputNewTag) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateCasbinRule(ctx context.Context, input *model.InputCasbinRule) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateCategories(ctx context.Context, input *model.InputCategories) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateContact(ctx context.Context, input *model.InputContact) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateFeebBack(ctx context.Context, input *model.InputFeedBacks) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateProduct(ctx context.Context, input *model.InputProduct) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateProfile(ctx context.Context, input *model.InputProfile) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMenues(ctx context.Context, input *model.InputMenues) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateProductCategories(ctx context.Context, input *model.InputProductCategory) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateSliders(ctx context.Context, input *model.InputSliders) (*bool, error) {
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

func (r *mutationResolver) UpdateCasbinRule(ctx context.Context, id *int, input *model.InputCasbinRule) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateCategories(ctx context.Context, id *int, input *model.InputCategories) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateContact(ctx context.Context, id *int, input *model.InputContact) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateFeebBack(ctx context.Context, id *int, input *model.InputFeedBacks) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateProduct(ctx context.Context, id *int, input *model.InputProduct) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateProfile(ctx context.Context, id *int, input *model.InputProfile) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMenues(ctx context.Context, id *int, input *model.InputMenues) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateProductCategories(ctx context.Context, id *int, input *model.InputProductCategory) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateSliders(ctx context.Context, id *int, input *model.InputSliders) (*bool, error) {
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

func (r *mutationResolver) RefreshToken(ctx context.Context, token *string) (*model.ResultToken, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CheckNews(ctx context.Context, id *int) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CheckUser(ctx context.Context, id *int) (*bool, error) {
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

func (r *queryResolver) Casbin(ctx context.Context) (*model.CasbinConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Categories(ctx context.Context) (*model.CategoriesConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Contacts(ctx context.Context) (*model.ContactConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Feedbacks(ctx context.Context) (*model.FeedBackConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Products(ctx context.Context) (*model.ProductConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Profile(ctx context.Context) (*model.ProfileConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SystemConfig(ctx context.Context) (*model.SystemConfigConnection, error) {
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
