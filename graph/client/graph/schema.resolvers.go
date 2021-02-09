package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"
	"fmt"

	db "github.com/JIeeiroSst/store/component"
	"github.com/JIeeiroSst/store/graph/client/graph/generated"
	"github.com/JIeeiroSst/store/graph/client/graph/model"
	"github.com/JIeeiroSst/store/models"
	"github.com/JIeeiroSst/store/utils/jwt"
	"github.com/JIeeiroSst/store/utils/tranfer"
)

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

func (r *mutationResolver) UpdateProfile(ctx context.Context, input *model.InputProfile) (*model.ResultCheck, error) {
	data := models.Profile{
		UserID:    input.UserID,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Address:   input.Address,
		Phone:     input.Phone,
	}

	err := db.GetConn().Model(models.Profile{}).Where("id = ? ", input.ID).Updates(data)
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

func (r *mutationResolver) RefreshToken(ctx context.Context, token *string) (*string, error) {
	id := jwt.GetIdUser(*token)
	s := fmt.Sprintf("%f", id)
	sEnc := base64.StdEncoding.EncodeToString([]byte(s))
	return &sEnc, nil
}

func (r *queryResolver) News(ctx context.Context) ([]*model.News, error) {
	var news []models.News
	var data *model.News
	var result = make([]*model.News, 0)
	var tag model.Tags
	db.GetConn().Find(&news)
	for _, new := range news {
		data = &model.News{
			ID:          *new.ID,
			Title:       new.Title,
			Description: new.Description,
			Image:       new.Image,
			CreatedAt:   tranfer.DeferTimeToString(new.CreatedAt),
			ViewCount:   new.ViewCount,
			Content:     new.Content,
			TagID:       new.TagID,
			Active:      new.Active,
		}
		db.GetConn().Where("id = ?", new.TagID).Find(&tag)
		data.Tags = append(data.Tags, &tag)
		result = append(result, data)
	}
	return result, nil
}

func (r *queryResolver) Tags(ctx context.Context) ([]*model.Tags, error) {
	var tags []models.Tags
	var data *model.Tags
	var result = make([]*model.Tags, 0)
	db.GetConn().Find(&tags)
	for _, tag := range tags {
		data = &model.Tags{
			ID:   tag.ID,
			Name: tag.Name,
		}
		result = append(result, data)
	}
	return result, nil
}

func (r *queryResolver) NewTags(ctx context.Context) ([]*model.NewTag, error) {
	var newTag []*models.NewTag
	var data *model.NewTag
	var result = make([]*model.NewTag, 0)
	var tag model.Tags
	var new model.News
	db.GetConn().Find(&newTag)
	for _, item := range newTag {
		data = &model.NewTag{
			ID:    item.ID,
			TagID: item.TagId,
			NewID: item.NewId,
			News:  nil,
			Tags:  nil,
		}

		db.GetConn().Where("id = ?", *item.TagId).Find(&tag)
		data.Tags = append(data.Tags, &tag)

		db.GetConn().Where("id = ?", *item.NewId).Find(&new)
		data.News = append(data.News, &new)
		result = append(result, data)
	}

	return result, nil
}

func (r *queryResolver) New(ctx context.Context, id *int) (*model.News, error) {
	var new models.News
	var result model.News
	var tag model.Tags

	db.GetConn().Where("id = ?", id).Find(&new)

	result = model.News{
		ID:          *new.ID,
		Title:       new.Title,
		Description: new.Description,
		Image:       new.Image,
		CreatedAt:   tranfer.DeferTimeToString(new.CreatedAt),
		ViewCount:   new.ViewCount,
		Content:     new.Content,
		TagID:       new.TagID,
		Active:      new.Active,
		Tags:        nil,
	}
	db.GetConn().Where("id = ?", new.TagID).Find(&tag)
	result.Tags = append(result.Tags, &tag)

	return &result, nil
}

func (r *queryResolver) Tag(ctx context.Context, id *int) (*model.Tags, error) {
	var tag models.Tags
	var result model.Tags
	db.GetConn().Where("id = ?", id).Find(&tag)
	result = model.Tags{
		ID:   tag.ID,
		Name: tag.Name,
	}
	return &result, nil
}

func (r *queryResolver) NewTag(ctx context.Context, id *int) (*model.NewTag, error) {
	var newTag models.NewTag
	var result model.NewTag
	var new model.News
	var tag model.Tags

	db.GetConn().Where("id = ?", id).Find(&newTag)
	result = model.NewTag{
		ID:    newTag.ID,
		TagID: newTag.TagId,
		NewID: newTag.NewId,
	}

	db.GetConn().Where("id = ?", newTag.NewId).Find(&new)
	result.News = append(result.News, &new)

	db.GetConn().Where("id = ?", newTag.TagId).Find(&tag)
	result.Tags = append(result.Tags, &tag)
	return &result, nil
}

func (r *queryResolver) Profile(ctx context.Context, id *int) (*model.Profile, error) {
	var profile models.Profile
	var result model.Profile

	db.GetConn().Where("user_id = ?", id).Find(&profile)
	result = model.Profile{
		ID:        profile.ID,
		UserID:    profile.UserID,
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		Address:   profile.Address,
		Phone:     profile.Phone,
		CreatedAt: tranfer.DeferTimeToString(profile.CreatedAt),
	}
	return &result, nil
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

func (r *mutationResolver) SendMail(ctx context.Context, email *string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) AddFriend(ctx context.Context, id *int) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}
