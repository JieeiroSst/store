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
	"github.com/JIeeiroSst/store/utils/tranfer"
)

func (r *mutationResolver) CreateNews(ctx context.Context, input *model.InputNews) (*model.ResultCheck, error) {
	active := false
	data := models.News{
		Title:       input.Title,
		Image:       input.Image,
		Detail:      input.Detail,
		Content:     input.Content,
		TagID:       input.TagID,
		Description: input.Description,
		Active:      tranfer.DeferBoolPonter(active),
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

func (r *mutationResolver) UpdateNews(ctx context.Context, input *model.InputNews) (*model.ResultCheck, error) {
	data := models.News{
		Title:       input.Title,
		Image:       input.Image,
		Detail:      input.Detail,
		Content:     input.Content,
		TagID:       input.TagID,
		Description: input.Description,
	}

	err := db.GetConn().Model(models.News{}).Where("id = ?", input.ID).Updates(data)
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

func (r *mutationResolver) UpdateTags(ctx context.Context, input *model.InputTags) (*model.ResultCheck, error) {
	data := models.Tags{
		Name: input.Name,
	}

	err := db.GetConn().Model(models.Tags{}).Where("id = ?", input.ID).Updates(data)
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

func (r *mutationResolver) UpdateNewTag(ctx context.Context, input *model.InputNewTag) (*model.ResultCheck, error) {
	data := models.NewTag{
		TagId: input.TagID,
		NewId: input.NewID,
	}
	err := db.GetConn().Model(models.NewTag{}).Where("id = ?",input.ID).Updates(data)
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

func (r *mutationResolver) UpdateFeebBack(ctx context.Context, input *model.InputFeedBacks) (*model.ResultCheck, error) {
	data := models.FeedBacks{
		Name:    input.Name,
		Phone:   input.Phone,
		Email:   input.Email,
		Address: input.Address,
		Content: input.Content,
	}

	err := db.GetConn().Model(models.FeedBacks{}).Where("id = ?", input.ID).Updates(data)
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

func (r *mutationResolver) UpdateMenues(ctx context.Context, input *model.InputMenues) (*model.ResultCheck, error) {
	data := model.Menues{
		Text:   input.Text,
		Link:   input.Link,
		Target: input.Target,
	}

	err := db.GetConn().Model(models.Menues{}).Where("id = ?", input.ID).Updates(data)
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
	err := db.GetConn().Delete(models.Menues{}, "id = ?", id)
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

func (r *mutationResolver) CheckNews(ctx context.Context, id *int) (*model.ResultCheck, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PublicNew(ctx context.Context, id *int) (*model.ResultCheck, error) {
	active := true
	data := models.News{Active: tranfer.DeferBoolPonter(active)}

	err := db.GetConn().Model(models.News{}).Where("id = ? ", id).Updates(data)
	if err != nil {
		status = false
		message = "public failure "
		result.Status = &status
		result.Message = &message
	}

	status = true
	message = "public successfully"
	result.Status = &status
	result.Message = &message

	return &result, nil
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

func (r *queryResolver) News(ctx context.Context) ([]*model.ResultNews, error) {
	var news []models.News
	var data *model.ResultNews
	var result = make([]*model.ResultNews, 0)
	db.GetConn().Find(&news)
	for _, new := range news {
		data = &model.ResultNews{
			ID:          new.ID,
			Title:       new.Title,
			Description: new.Description,
			Image:       new.Image,
			Detail:      new.Detail,
			CreatedAt:   tranfer.DeferTimeToString(new.CreatedAt),
			ViewCount:   new.ViewCount,
			Content:     new.Content,
			TagID:       new.TagID,
			Active:      new.Active,
		}
		result = append(result, data)
	}

	return result, nil
}

func (r *queryResolver) Tags(ctx context.Context) ([]*model.ResultTags, error) {
	var tags []models.Tags
	var data *model.ResultTags
	var result = make([]*model.ResultTags, 0)
	db.GetConn().Find(&tags)
	for _, tag := range tags {
		data = &model.ResultTags{
			ID:   tag.ID,
			Name: tag.Name,
		}
		result = append(result, data)
	}
	return result, nil
}

func (r *queryResolver) NewTags(ctx context.Context) ([]*model.ResultNewTag, error) {
	var newTags []models.NewTag
	var data *model.ResultNewTag
	var result = make([]*model.ResultNewTag, 0)
	db.GetConn().Find(&newTags)
	for _, item := range newTags {
		data = &model.ResultNewTag{
			ID:    item.ID,
			TagID: item.TagId,
			NewID: item.NewId,
		}
		result = append(result, data)
	}
	return result, nil
}

func (r *queryResolver) Feedbacks(ctx context.Context) ([]*model.ResultFeedBacks, error) {
	var feedBacks []models.FeedBacks
	var data *model.ResultFeedBacks
	var result = make([]*model.ResultFeedBacks, 0)
	db.GetConn().Find(&feedBacks)

	for _, item := range feedBacks {
		data = &model.ResultFeedBacks{
			ID:        item.ID,
			Name:      item.Name,
			Phone:     item.Phone,
			Email:     item.Email,
			Address:   item.Address,
			Content:   item.Content,
			CreatedAt: tranfer.DeferTimeToString(item.CreatedAt),
		}

		result = append(result, data)
	}
	return result, nil
}

func (r *queryResolver) Profiles(ctx context.Context) ([]*model.ResultProfile, error) {
	var profiles []models.Profile
	var data *model.ResultProfile
	var result = make([]*model.ResultProfile, 0)
	db.GetConn().Find(&profiles)

	for _, item := range profiles {
		data = &model.ResultProfile{
			ID:        item.ID,
			UserID:    item.UserID,
			FirstName: item.FirstName,
			LastName:  item.LastName,
			Address:   item.Address,
			Phone:     item.Phone,
			CreatedAt: tranfer.DeferTimeToString(item.CreatedAt),
		}
		result = append(result, data)
	}
	return result, nil
}

func (r *queryResolver) SystemConfigs(ctx context.Context) ([]*model.ResultSystemConfig, error) {
	var systemConfigs []models.SystemConfig
	var data *model.ResultSystemConfig
	var result = make([]*model.ResultSystemConfig, 0)
	for _, item := range systemConfigs {
		data = &model.ResultSystemConfig{
			ID:    item.ID,
			Name:  item.Name,
			Type:  item.Type,
			Value: item.Value,
		}
		result = append(result, data)
	}
	return result, nil
}

func (r *queryResolver) New(ctx context.Context, id *int) (*model.ResultNews, error) {
	var new models.News
	var result model.ResultNews

	db.GetConn().Where("id = ?", id).Find(&new)
	result = model.ResultNews{
		ID:          new.ID,
		Title:       new.Title,
		Description: new.Description,
		Image:       new.Image,
		Detail:      new.Detail,
		CreatedAt:   tranfer.DeferTimeToString(new.CreatedAt),
		ViewCount:   new.ViewCount,
		Content:     new.Content,
		TagID:       new.TagID,
		Active:      new.Active,
	}
	return &result, nil
}

func (r *queryResolver) Tag(ctx context.Context, id *int) (*model.ResultTags, error) {
	var tag models.Tags
	var result model.ResultTags

	db.GetConn().Where("id = ?", id).Find(&tag)
	result = model.ResultTags{
		ID:   tag.ID,
		Name: tag.Name,
	}
	return &result, nil
}

func (r *queryResolver) NewTag(ctx context.Context, id *int) (*model.ResultNewTag, error) {
	var newTag models.NewTag
	var result model.ResultNewTag

	db.GetConn().Where("id = ?", id).Find(&newTag)
	result = model.ResultNewTag{
		ID:    newTag.ID,
		TagID: newTag.TagId,
		NewID: newTag.NewId,
	}
	return &result, nil
}

func (r *queryResolver) Feedback(ctx context.Context, id *int) (*model.ResultFeedBacks, error) {
	var feedBack models.FeedBacks
	var result model.ResultFeedBacks

	db.GetConn().Where("id = ?", id).Find(&feedBack)

	result = model.ResultFeedBacks{
		ID:        feedBack.ID,
		Name:      feedBack.Name,
		Phone:     feedBack.Phone,
		Email:     feedBack.Email,
		Address:   feedBack.Address,
		Content:   feedBack.Content,
		CreatedAt: tranfer.DeferTimeToString(feedBack.CreatedAt),
	}
	return &result, nil
}

func (r *queryResolver) ProfileID(ctx context.Context, id *int) (*model.ResultProfile, error) {
	var profile models.Profile
	var result model.ResultProfile
	db.GetConn().Where("id = ?", id).Find(&profile)
	result = model.ResultProfile{
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

func (r *queryResolver) Systemconfig(ctx context.Context, id *int) (*model.ResultSystemConfig, error) {
	var systemConfig models.SystemConfig
	var result model.ResultSystemConfig
	db.GetConn().Where("id = ?", id).Find(&systemConfig)
	result = model.ResultSystemConfig{
		ID:    systemConfig.ID,
		Name:  systemConfig.Name,
		Type:  systemConfig.Type,
		Value: systemConfig.Value,
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
