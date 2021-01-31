package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	db "github.com/JIeeiroSst/store/component"
	"github.com/JIeeiroSst/store/graph/admin/graph/generated"
	"github.com/JIeeiroSst/store/graph/admin/graph/model"
	"github.com/JIeeiroSst/store/models"
)

func (r *mutationResolver) CreateNews(ctx context.Context, input *model.InputNews) (*model.ResultCheck, error) {
	data := models.News{
		Title:       input.Title,
		Image:       input.Image,
		Detail:      input.Detail,
		Content:     input.Content,
		TagID:       input.TagID,
		Description: input.Description,
	}
	err := db.GetConn().Create(&data)
	if err != nil {
		status = false
		message = "create failure "
		result.Status = &status
		result.Message = &message
	}

	status = true
	message = "created successfully"
	result.Status = &status
	result.Message = &message
	return &result, nil
}

func (r *mutationResolver) CreateTags(ctx context.Context, input *model.InputTags) (*model.ResultCheck, error) {
	data := models.Tags{
		Name: input.Name,
	}
	err := db.GetConn().Create(&data)
	if err != nil {
		status = false
		message = "create failure "
		result.Status = &status
		result.Message = &message
	}
	status = true
	message = "created successfully"
	result.Status = &status
	result.Message = &message
	return &result, nil
}

func (r *mutationResolver) CreateNewTag(ctx context.Context, input *model.InputNewTag) (*model.ResultCheck, error) {
	data := models.NewTag{
		TagId: input.TagID,
		NewId: input.NewID,
	}
	err := db.GetConn().Create(&data)
	if err != nil {
		status = false
		message = "create failure "
		result.Status = &status
		result.Message = &message
	}
	status = true
	message = "created successfully"
	result.Status = &status
	result.Message = &message

	return &result, nil
}

func (r *mutationResolver) CreateFeebBack(ctx context.Context, input *model.InputFeedBacks) (*model.ResultCheck, error) {
	data := models.FeedBacks{
		Name:    input.Name,
		Phone:   input.Phone,
		Email:   input.Email,
		Address: input.Address,
		Content: input.Content,
	}

	err := db.GetConn().Create(&data)
	if err != nil {
		status = false
		message = "create failure "
		result.Status = &status
		result.Message = &message
	}
	status = true
	message = "created successfully"
	result.Status = &status
	result.Message = &message

	return &result, nil
}

func (r *mutationResolver) CreateProfile(ctx context.Context, input *model.InputProfile) (*model.ResultCheck, error) {
	data := models.Profile{
		UserID:    input.UserID,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Address:   input.Address,
		Phone:     input.Phone,
	}
	err := db.GetConn().Create(&data)
	if err != nil {
		status = false
		message = "create failure "
		result.Status = &status
		result.Message = &message
	}

	status = true
	message = "created successfully"
	result.Status = &status
	result.Message = &message

	return &result, nil
}

func (r *mutationResolver) CreateMenues(ctx context.Context, input *model.InputMenues) (*model.ResultCheck, error) {
	data := model.Menues{
		Text:   input.Text,
		Link:   input.Link,
		Target: input.Target,
	}
	err := db.GetConn().Create(&data)
	if err != nil {
		status = false
		message = "create failure "
		result.Status = &status
		result.Message = &message
	}
	status = true
	message = "created successfully"
	result.Status = &status
	result.Message = &message

	return &result, nil
}

func (r *mutationResolver) UpdateNews(ctx context.Context, id *int, input *model.InputNews) (*model.ResultCheck, error) {
	data := models.News{
		Title:       input.Title,
		Image:       input.Image,
		Detail:      input.Detail,
		Content:     input.Content,
		TagID:       input.TagID,
		Description: input.Description,
	}

	err := db.GetConn().Model(models.News{}).Where("id = ?", id).Updates(data)
	if err != nil {
		status = false
		message = "update failure "
		result.Status = &status
		result.Message = &message
	}
	status = true
	message = "update successfully"
	result.Status = &status
	result.Message = &message

	return &result, nil
}

func (r *mutationResolver) UpdateTags(ctx context.Context, id *int, input *model.InputTags) (*model.ResultCheck, error) {
	data := models.Tags{
		Name: input.Name,
	}

	err := db.GetConn().Model(models.Tags{}).Where("id = ?", id).Updates(data)
	if err != nil {
		status = false
		message = "update failure "
		result.Status = &status
		result.Message = &message
	}

	status = true
	message = "update successfully"
	result.Status = &status
	result.Message = &message

	return &result, nil
}

func (r *mutationResolver) UpdateNewTag(ctx context.Context, id *int, input *model.InputNewTag) (*model.ResultCheck, error) {
	data := models.NewTag{
		TagId: input.TagID,
		NewId: input.NewID,
	}
	err := db.GetConn().Model(models.NewTag{}).Where("id = ?").Updates(data)
	if err != nil {
		status = false
		message = "update failure "
		result.Status = &status
		result.Message = &message
	}
	status = true
	message = "update successfully"
	result.Status = &status
	result.Message = &message

	return &result, nil
}

func (r *mutationResolver) UpdateFeebBack(ctx context.Context, id *int, input *model.InputFeedBacks) (*model.ResultCheck, error) {
	data := models.FeedBacks{
		Name:    input.Name,
		Phone:   input.Phone,
		Email:   input.Email,
		Address: input.Address,
		Content: input.Content,
	}

	err := db.GetConn().Model(models.FeedBacks{}).Where("id = ?", id).Updates(data)
	if err != nil {
		status = false
		message = "update failure "
		result.Status = &status
		result.Message = &message
	}
	status = true
	message = "update successfully"
	result.Status = &status
	result.Message = &message

	return &result, nil
}

func (r *mutationResolver) UpdateProfile(ctx context.Context, id *int, input *model.InputProfile) (*model.ResultCheck, error) {
	data := models.Profile{
		UserID:    input.UserID,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Address:   input.Address,
		Phone:     input.Phone,
	}

	err := db.GetConn().Model(models.Profile{}).Where("id = ? ", id).Updates(data)
	if err != nil {
		status = false
		message = "update failure "
		result.Status = &status
		result.Message = &message
	}

	status = true
	message = "update successfully"
	result.Status = &status
	result.Message = &message

	return &result, nil
}

func (r *mutationResolver) UpdateMenues(ctx context.Context, id *int, input *model.InputMenues) (*model.ResultCheck, error) {
	data := model.Menues{
		Text:   input.Text,
		Link:   input.Link,
		Target: input.Target,
	}

	err := db.GetConn().Model(models.Menues{}).Where("id = ?", id).Updates(data)
	if err != nil {
		status = false
		message = "update failure "
		result.Status = &status
		result.Message = &message
	}
	status = true
	message = "update successfully"
	result.Status = &status
	result.Message = &message

	return &result, nil
}

func (r *mutationResolver) DeleteNews(ctx context.Context, id *int) (*model.ResultCheck, error) {
	err := db.GetConn().Delete(models.News{}, "id = ?", id)
	if err != nil {
		status = false
		message = "delete failure "
		result.Status = &status
		result.Message = &message
	}
	status = true
	message = "delete successfully"
	result.Status = &status
	result.Message = &message

	return &result, nil
}

func (r *mutationResolver) DeleteTags(ctx context.Context, id *int) (*model.ResultCheck, error) {
	err := db.GetConn().Delete(models.Tags{}, "id = ?", id)
	if err != nil {
		status = false
		message = "delete failure "
		result.Status = &status
		result.Message = &message
	}
	status = true
	message = "delete successfully"
	result.Status = &status
	result.Message = &message

	return &result, nil
}

func (r *mutationResolver) DeleteNewTag(ctx context.Context, id *int) (*model.ResultCheck, error) {
	err := db.GetConn().Delete(models.NewTag{}, "id = ?", id)
	if err != nil {
		status = false
		message = "delete failure "
		result.Status = &status
		result.Message = &message
	}
	status = true
	message = "delete successfully"
	result.Status = &status
	result.Message = &message

	return &result, nil
}

func (r *mutationResolver) DeleteFeebBack(ctx context.Context, id *int) (*model.ResultCheck, error) {
	err := db.GetConn().Delete(models.FeedBacks{}, "id = ?", id)
	if err != nil {
		status = false
		message = "delete failure "
		result.Status = &status
		result.Message = &message
	}
	status = true
	message = "delete successfully"
	result.Status = &status
	result.Message = &message

	return &result, nil
}

func (r *mutationResolver) DeleteProfile(ctx context.Context, id *int) (*model.ResultCheck, error) {
	err := db.GetConn().Delete(models.Profile{}, "id = ?", id)
	if err != nil {
		status = false
		message = "delete failure "
		result.Status = &status
		result.Message = &message
	}
	status = true
	message = "delete successfully"
	result.Status = &status
	result.Message = &message

	return &result, nil
}

func (r *mutationResolver) DeleteMenues(ctx context.Context, id *int) (*model.ResultCheck, error) {
	err := db.GetConn().Delete(models.Menues{},"id = ?",id)
	if err !=nil {
		status = false
		message = "delete failure "
		result.Status = &status
		result.Message = &message
	}
	status = true
	message = "delete successfully"
	result.Status = &status
	result.Message = &message

	return &result, nil
}

func (r *mutationResolver) CheckNews(ctx context.Context, id *int) (*model.ResultCheck, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PublicNew(ctx context.Context, id *int) (*model.ResultCheck, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SendMail(ctx context.Context, email *string) (*model.ResultCheck, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddFriend(ctx context.Context, id *int) (*model.ResultCheck, error) {
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var (
	status  bool
	message string
	result  model.ResultCheck
)
