package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	db "github.com/JIeeiroSst/store/component"
	"github.com/JIeeiroSst/store/graph/generated"
	"github.com/JIeeiroSst/store/graph/model"
	casbin "github.com/JIeeiroSst/store/models/casbin_rules"
	"github.com/JIeeiroSst/store/models/categories"
	"github.com/JIeeiroSst/store/models/contacts"
	feedBacks "github.com/JIeeiroSst/store/models/feed-backs"
	"github.com/JIeeiroSst/store/models/menues"
	newTags "github.com/JIeeiroSst/store/models/new-tags"
	"github.com/JIeeiroSst/store/models/news"
	productCategory "github.com/JIeeiroSst/store/models/product_category"
	"github.com/JIeeiroSst/store/models/products"
	"github.com/JIeeiroSst/store/models/profiles"
	"github.com/JIeeiroSst/store/models/sliders"
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
	check := false
	try.This(func() {
		data := &newTags.NewTag{
			TagId: tranfer.DeferInt(input.TagID),
			Name:  tranfer.DeferString(input.Name),
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

func (r *mutationResolver) CreateCasbinRule(ctx context.Context, input *model.InputCasbinRule) (*bool, error) {
	check := false
	try.This(func() {
		id ,_ :=strconv.Atoi(input.ID)

		data := &casbin.CasbinRules{
			Id:    id,
			PType: tranfer.DeferString(input.PType),
			V0:    tranfer.DeferString(input.V0),
			V1:    tranfer.DeferString(input.V1),
			V2:    tranfer.DeferString(input.V2),
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

func (r *mutationResolver) CreateCategories(ctx context.Context, input *model.InputCategories) (*bool, error) {
	check := false
	try.This(func() {
		data := &categories.Categories{
			Name:            tranfer.DeferString(input.Name),
			MetaTitle:       tranfer.DeferString(input.MetaTitle),
			CreatedBy:       tranfer.DeferString(input.ModifiedBy),
			MetaKeyword:     tranfer.DeferString(input.MetaKeyword),
			MetaDescription: tranfer.DeferString(input.MetaDescription),
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

func (r *mutationResolver) CreateContact(ctx context.Context, input *model.InputContact) (*bool, error) {
	check := false
	try.This(func() {
		data := &contacts.Contacts{
			Content: tranfer.DeferString(input.Content),
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

func (r *mutationResolver) CreateFeebBack(ctx context.Context, input *model.InputFeedBacks) (*bool, error) {
	check := false
	try.This(func() {
		data := &feedBacks.FeedBacks{
			Name:      tranfer.DeferString(input.Name),
			Phone:     tranfer.DeferString(input.Phone),
			Email:     tranfer.DeferString(input.Email),
			Address:   tranfer.DeferString(input.Address),
			Content:   tranfer.DeferString(input.Content),
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

func (r *mutationResolver) CreateProduct(ctx context.Context, input *model.InputProduct) (*bool, error) {
	check := false
	try.This(func() {
		data := &products.Products{
			Code:         tranfer.DeferString(input.Code),
			Name:         tranfer.DeferString(input.Name),
			Title:        tranfer.DeferString(input.Title),
			Description:  tranfer.DeferString(input.Description),
			Images:       tranfer.DeferString(input.Images),
			Price:        tranfer.DeferInt(input.Price),
			Vat:          tranfer.DeferInt(input.Vat),
			CategoryId:   tranfer.DeferInt(input.CatedgoryID),
			Detail:       tranfer.DeferString(input.Detail),
			CreatedBy:    tranfer.DeferString(input.CreatedBy),
			ModifiedBy:   tranfer.DeferString(input.ModifiedBy),
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

func (r *mutationResolver) CreateProfile(ctx context.Context, input *model.InputProfile) (*bool, error) {
	check := false
	try.This(func() {
		id , _ :=strconv.Atoi(input.UserID)
		data := &profiles.Profiles{
			UserId:       id,
			FirstName:    tranfer.DeferString(input.FirstName),
			LastName:     tranfer.DeferString(input.LastName),
			Address:      tranfer.DeferString(input.Address),
			Phone:        tranfer.DeferString(input.Phone),
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

func (r *mutationResolver) CreateMenues(ctx context.Context, input *model.InputMenues) (*bool, error) {
	check := false
	try.This(func() {
		data := &menues.Menues{
			Text:         tranfer.DeferString(input.Text),
			Link:         tranfer.DeferString(input.Link),
			Target:       tranfer.DeferString(input.Target),
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

func (r *mutationResolver) CreateProductCategories(ctx context.Context, input *model.InputProductCategory) (*bool, error) {
	check := false
	try.This(func() {
		data := &productCategory.ProductCategory{
			Name:         tranfer.DeferString(input.Name),
			MetaTitle:    tranfer.DeferString(input.MetaTitle),
			Title:        tranfer.DeferString(input.Title),
			CreatedBy:    tranfer.DeferString(input.CreatedBy),
			ModifiedBy:   tranfer.DeferString(input.ModifiedBy),
			Description:  tranfer.DeferString(input.Description),
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

func (r *mutationResolver) CreateSliders(ctx context.Context, input *model.InputSliders) (*bool, error) {
	check := false
	try.This(func() {
		data := &sliders.Sliders{
			Image:        tranfer.DeferString(input.Image),
			Link:         tranfer.DeferString(input.Link),
			Description:  tranfer.DeferString(input.Description),
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

func (r *mutationResolver) UpdateNews(ctx context.Context, id *int, input *model.InputNews) (*bool, error) {
	check := false
	try.This(func() {
		db.GetConn().Where("id = ?", id).First(&news.News{})
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

		db.GetConn().Save(&data)

		check = true

	}).Finally(func() {
		logger.Log.Info("this must be printed after the catch")
	}).Catch(func(e try.E) {
		logger.Log.Info(e)
	})
	return &check, nil
}

func (r *mutationResolver) UpdateTags(ctx context.Context, id *int, input *model.InputTags) (*bool, error) {
	check := false
	try.This(func() {
		db.GetConn().Where("id = ?", id).First(&tags.Tags{})
		data := &tags.Tags{
			Name:  tranfer.DeferString(input.Name),
			NewId: tranfer.DeferInt(input.NewID),
		}

		db.GetConn().Save(&data)

		check = true

	}).Finally(func() {
		logger.Log.Info("this must be printed after the catch")
	}).Catch(func(e try.E) {
		logger.Log.Info(e)
	})
	return &check, nil
}

func (r *mutationResolver) UpdateNewTag(ctx context.Context, id *int, input *model.InputNewTag) (*bool, error) {
	check := false
	try.This(func() {
		db.GetConn().Where("id = ?", id).First(&newTags.NewTag{})
		data := &newTags.NewTag{
			TagId: tranfer.DeferInt(input.TagID),
			Name:  tranfer.DeferString(input.Name),
		}

		db.GetConn().Save(&data)

		check = true

	}).Finally(func() {
		logger.Log.Info("this must be printed after the catch")
	}).Catch(func(e try.E) {
		logger.Log.Info(e)
	})
	return &check, nil
}

func (r *mutationResolver) UpdateCasbinRule(ctx context.Context, id *int, input *model.InputCasbinRule) (*bool, error) {
	check := false
	try.This(func() {
		db.GetConn().Where("id = ?", id).First(&casbin.CasbinRules{})
		data := &casbin.CasbinRules{
			PType: tranfer.DeferString(input.PType),
			V0:    tranfer.DeferString(input.V0),
			V1:    tranfer.DeferString(input.V1),
			V2:    tranfer.DeferString(input.V2),
		}

		db.GetConn().Save(&data)

		check = true

	}).Finally(func() {
		logger.Log.Info("this must be printed after the catch")
	}).Catch(func(e try.E) {
		logger.Log.Info(e)
	})
	return &check, nil
}

func (r *mutationResolver) UpdateCategories(ctx context.Context, id *int, input *model.InputCategories) (*bool, error) {
	check := false
	try.This(func() {
		db.GetConn().Where("id = ?", id).First(&categories.Categories{})
		data := &categories.Categories{
			Name:            tranfer.DeferString(input.Name),
			MetaTitle:       tranfer.DeferString(input.MetaTitle),
			CreatedBy:       tranfer.DeferString(input.ModifiedBy),
			MetaKeyword:     tranfer.DeferString(input.MetaKeyword),
			MetaDescription: tranfer.DeferString(input.MetaDescription),
		}

		db.GetConn().Save(&data)

		check = true

	}).Finally(func() {
		logger.Log.Info("this must be printed after the catch")
	}).Catch(func(e try.E) {
		logger.Log.Info(e)
	})
	return &check, nil
}

func (r *mutationResolver) UpdateContact(ctx context.Context, id *int, input *model.InputContact) (*bool, error) {
	check := false
	try.This(func() {
		db.GetConn().Where("id = ?", id).First(&contacts.Contacts{})
		data := &contacts.Contacts{
			Content: tranfer.DeferString(input.Content),
		}

		db.GetConn().Save(&data)

		check = true

	}).Finally(func() {
		logger.Log.Info("this must be printed after the catch")
	}).Catch(func(e try.E) {
		logger.Log.Info(e)
	})
	return &check, nil
}

func (r *mutationResolver) UpdateFeebBack(ctx context.Context, id *int, input *model.InputFeedBacks) (*bool, error) {
	check := false
	try.This(func() {
		db.GetConn().Where("id = ?", id).First(&feedBacks.FeedBacks{})
		data := &feedBacks.FeedBacks{
			Name:      tranfer.DeferString(input.Name),
			Phone:     tranfer.DeferString(input.Phone),
			Email:     tranfer.DeferString(input.Email),
			Address:   tranfer.DeferString(input.Address),
			Content:   tranfer.DeferString(input.Content),
		}

		db.GetConn().Save(&data)

		check = true

	}).Finally(func() {
		logger.Log.Info("this must be printed after the catch")
	}).Catch(func(e try.E) {
		logger.Log.Info(e)
	})
	return &check, nil
}

func (r *mutationResolver) UpdateProduct(ctx context.Context, id *int, input *model.InputProduct) (*bool, error) {
	check := false
	try.This(func() {
		db.GetConn().Where("id = ?", id).First(&products.Products{})
		data := &products.Products{
			Code:         tranfer.DeferString(input.Code),
			Name:         tranfer.DeferString(input.Name),
			Title:        tranfer.DeferString(input.Title),
			Description:  tranfer.DeferString(input.Description),
			Images:       tranfer.DeferString(input.Images),
			Price:        tranfer.DeferInt(input.Price),
			Vat:          tranfer.DeferInt(input.Vat),
			CategoryId:   tranfer.DeferInt(input.CatedgoryID),
			Detail:       tranfer.DeferString(input.Detail),
			CreatedBy:    tranfer.DeferString(input.CreatedBy),
			ModifiedBy:   tranfer.DeferString(input.ModifiedBy),
		}

		db.GetConn().Save(&data)

		check = true

	}).Finally(func() {
		logger.Log.Info("this must be printed after the catch")
	}).Catch(func(e try.E) {
		logger.Log.Info(e)
	})
	return &check, nil
}

func (r *mutationResolver) UpdateProfile(ctx context.Context, id *int, input *model.InputProfile) (*bool, error) {
	check := false
	try.This(func() {
		db.GetConn().Where("id = ?", id).First(&profiles.Profiles{})
		data := &profiles.Profiles{
			FirstName:    tranfer.DeferString(input.FirstName),
			LastName:     tranfer.DeferString(input.LastName),
			Address:      tranfer.DeferString(input.Address),
			Phone:        tranfer.DeferString(input.Phone),
		}

		db.GetConn().Save(&data)

		check = true

	}).Finally(func() {
		logger.Log.Info("this must be printed after the catch")
	}).Catch(func(e try.E) {
		logger.Log.Info(e)
	})
	return &check, nil
}

func (r *mutationResolver) UpdateMenues(ctx context.Context, id *int, input *model.InputMenues) (*bool, error) {
	check := false
	try.This(func() {
		db.GetConn().Where("id = ?", id).First(&menues.Menues{})
		data := &menues.Menues{
			Text:         tranfer.DeferString(input.Text),
			Link:         tranfer.DeferString(input.Link),
			Target:       tranfer.DeferString(input.Target),
		}

		db.GetConn().Save(&data)

		check = true

	}).Finally(func() {
		logger.Log.Info("this must be printed after the catch")
	}).Catch(func(e try.E) {
		logger.Log.Info(e)
	})
	return &check, nil
}

func (r *mutationResolver) UpdateProductCategories(ctx context.Context, id *int, input *model.InputProductCategory) (*bool, error) {
	check := false
	try.This(func() {
		db.GetConn().Where("id = ?", id).First(&productCategory.ProductCategory{})
		data := &productCategory.ProductCategory{
			Name:         tranfer.DeferString(input.Name),
			MetaTitle:    tranfer.DeferString(input.MetaTitle),
			Title:        tranfer.DeferString(input.Title),
			CreatedBy:    tranfer.DeferString(input.CreatedBy),
			ModifiedBy:   tranfer.DeferString(input.ModifiedBy),
			Description:  tranfer.DeferString(input.Description),
		}

		db.GetConn().Save(&data)

		check = true

	}).Finally(func() {
		logger.Log.Info("this must be printed after the catch")
	}).Catch(func(e try.E) {
		logger.Log.Info(e)
	})
	return &check, nil
}

func (r *mutationResolver) UpdateSliders(ctx context.Context, id *int, input *model.InputSliders) (*bool, error) {
	check := false
	try.This(func() {
		db.GetConn().Where("id = ?", id).First(&sliders.Sliders{})
		data := &sliders.Sliders{
			Image:        tranfer.DeferString(input.Image),
			Link:         tranfer.DeferString(input.Link),
			Description:  tranfer.DeferString(input.Description),
		}

		db.GetConn().Save(&data)

		check = true

	}).Finally(func() {
		logger.Log.Info("this must be printed after the catch")
	}).Catch(func(e try.E) {
		logger.Log.Info(e)
	})
	return &check, nil
}

func (r *mutationResolver) DeleteNews(ctx context.Context, id *int) (*bool, error) {
	check := false
	try.This(func() {
		db.GetConn().Where("id = ?",id).Delete(&news.News{})
		check = true

	}).Finally(func() {
		logger.Log.Info("this must be printed after the catch")
	}).Catch(func(e try.E) {
		logger.Log.Info(e)
	})
	return &check, nil
}

func (r *mutationResolver) DeleteTags(ctx context.Context, id *int) (*bool, error) {
	check := false
	try.This(func() {
		db.GetConn().Where("id = ?",id).Delete(&tags.Tags{})
		check = true

	}).Finally(func() {
		logger.Log.Info("this must be printed after the catch")
	}).Catch(func(e try.E) {
		logger.Log.Info(e)
	})
	return &check, nil
}

func (r *mutationResolver) DeleteNewTag(ctx context.Context, id *int) (*bool, error) {
	check := false
	try.This(func() {
		db.GetConn().Where("id = ?",id).Delete(&newTags.NewTag{})
		check = true

	}).Finally(func() {
		logger.Log.Info("this must be printed after the catch")
	}).Catch(func(e try.E) {
		logger.Log.Info(e)
	})
	return &check, nil
}

func (r *mutationResolver) DeleteCasbinRule(ctx context.Context, id *int) (*bool, error) {
	check := false
	try.This(func() {
		db.GetConn().Where("id = ?",id).Delete(&casbin.CasbinRules{})
		check = true

	}).Finally(func() {
		logger.Log.Info("this must be printed after the catch")
	}).Catch(func(e try.E) {
		logger.Log.Info(e)
	})
	return &check, nil
}

func (r *mutationResolver) DeleteCategories(ctx context.Context, id *int) (*bool, error) {
	check := false
	try.This(func() {
		db.GetConn().Where("id = ?",id).Delete(&categories.Categories{})
		check = true

	}).Finally(func() {
		logger.Log.Info("this must be printed after the catch")
	}).Catch(func(e try.E) {
		logger.Log.Info(e)
	})
	return &check, nil
}

func (r *mutationResolver) DeleteContact(ctx context.Context, id *int) (*bool, error) {
	check := false
	try.This(func() {
		db.GetConn().Where("id = ?",id).Delete(&contacts.Contacts{})
		check = true

	}).Finally(func() {
		logger.Log.Info("this must be printed after the catch")
	}).Catch(func(e try.E) {
		logger.Log.Info(e)
	})
	return &check, nil
}

func (r *mutationResolver) DeleteFeebBack(ctx context.Context, id *int) (*bool, error) {
	check := false
	try.This(func() {
		db.GetConn().Where("id = ?",id).Delete(&feedBacks.FeedBacks{})
		check = true

	}).Finally(func() {
		logger.Log.Info("this must be printed after the catch")
	}).Catch(func(e try.E) {
		logger.Log.Info(e)
	})
	return &check, nil
}

func (r *mutationResolver) DeleteProduct(ctx context.Context, id *int) (*bool, error) {
	check := false
	try.This(func() {
		db.GetConn().Where("id = ?",id).Delete(&products.Products{})
		check = true

	}).Finally(func() {
		logger.Log.Info("this must be printed after the catch")
	}).Catch(func(e try.E) {
		logger.Log.Info(e)
	})
	return &check, nil
}

func (r *mutationResolver) DeleteProfile(ctx context.Context, userID *int) (*bool, error) {
	check := false
	try.This(func() {
		db.GetConn().Where("id = ?",userID).Delete(&profiles.Profiles{})
		check = true

	}).Finally(func() {
		logger.Log.Info("this must be printed after the catch")
	}).Catch(func(e try.E) {
		logger.Log.Info(e)
	})
	return &check, nil
}

func (r *mutationResolver) DeleteMenues(ctx context.Context, id *int) (*bool, error) {
	check := false
	try.This(func() {
		db.GetConn().Where("id = ?",id).Delete(&menues.Menues{})
		check = true

	}).Finally(func() {
		logger.Log.Info("this must be printed after the catch")
	}).Catch(func(e try.E) {
		logger.Log.Info(e)
	})
	return &check, nil
}

func (r *mutationResolver) DeleteProductCategories(ctx context.Context, id *int) (*bool, error) {
	check := false
	try.This(func() {
		db.GetConn().Where("id = ?",id).Delete(&productCategory.ProductCategory{})
		check = true

	}).Finally(func() {
		logger.Log.Info("this must be printed after the catch")
	}).Catch(func(e try.E) {
		logger.Log.Info(e)
	})
	return &check, nil
}

func (r *mutationResolver) DeleteSliders(ctx context.Context, id *int) (*bool, error) {
	check := false
	try.This(func() {
		db.GetConn().Where("id = ?",id).Delete(&sliders.Sliders{})
		check = true

	}).Finally(func() {
		logger.Log.Info("this must be printed after the catch")
	}).Catch(func(e try.E) {
		logger.Log.Info(e)
	})
	return &check, nil
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
