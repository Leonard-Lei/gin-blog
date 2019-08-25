package category_service

import (
	"gin-blog/models"
)

type Category struct {
	ID         int
	Name       string
	State      int
	CreateBy   int
	CreateTime string
	UpdateBy   int

	PageNum  int
	PageSize int
}

func (a *Category) Add() error {
	comment := map[string]interface{}{
		"name":      a.Name,
		"create_by": a.CreateBy,
		"state":     a.State,
	}

	if err := models.AddCategory(comment); err != nil {
		return err
	}

	return nil
}

func (a *Category) Edit() error {
	return models.EditCategory(a.ID, map[string]interface{}{
		"id":        a.ID,
		"state":     a.State,
		"update_by": a.UpdateBy,
	})
}

func (a *Category) Get() (*models.Category, error) {

	category, err := models.GetCategory(a.ID)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (a *Category) GetAll() ([]*models.Category, error) {

	categorys, err := models.GetCategorys(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}
	return categorys, nil
}

func (a *Category) Delete() error {
	return models.DeleteCategory(a.ID)
}

func (a *Category) ExistByID() (bool, error) {
	return models.ExistCategoryByID(a.ID)
}

func (a *Category) ExistByName() (bool, error) {
	return models.ExistCategoryByName(a.Name)
}
func (a *Category) Count() (int, error) {
	return models.GetCategoryTotal(a.getMaps())
}

func (a *Category) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["delete_flag"] = 0
	if a.State != -1 {
		maps["state"] = a.State
	}
	return maps
}
