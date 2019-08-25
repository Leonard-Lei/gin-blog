package models

import "github.com/jinzhu/gorm"

type Category struct {
	Model

	Name     string `json:"name"`
	CreateBy int    `json:"create_by"`
	State    int    `json:"state"`
}

// ExistCategoryByID checks if an category exists based on ID
func ExistCategoryByID(id int) (bool, error) {
	var category Category
	err := db.Select("id").Where("id = ? AND delete_flag = ? ", id, 0).First(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if category.ID > 0 {
		return true, nil
	}

	return false, nil
}

// ExistCategoryByName checks if an category exists based on Name
func ExistCategoryByName(name string) (bool, error) {
	var category Category
	err := db.Select("name").Where("name = ? AND delete_flag = ? ", name, 0).First(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if category.Name != "" {
		return true, nil
	}

	return false, nil
}

// GetCategoryTotal gets the total number of category based on the constraints
func GetCategoryTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Category{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetCategory gets a list of categorys based on paging constraints
func GetCategorys(pageNum int, pageSize int, maps interface{}) ([]*Category, error) {
	var category []*Category
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return category, nil
}

// GetCategoryt Get a single category based on ID
func GetCategory(id int) (*Category, error) {
	var category Category
	err := db.Where("id = ? AND delete_flag = ?", id, 0).First(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	err = db.Model(&category).Related(&category.ID).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &category, nil
}

// EditCategory modify a single comment
func EditCategory(id int, data interface{}) error {
	if err := db.Model(&Category{}).Where("id = ? AND delete_flag = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

//Add a Category
func AddCategory(data map[string]interface{}) error {
	category := Category{
		Name:     data["name"].(string),
		CreateBy: data["create_by"].(int),
		State:    data["state"].(int),
	}

	if err := db.Create(&category).Error; err != nil {
		return err
	}

	return nil
}

// DeleteCategory delete a single category
func DeleteCategory(id int) error {
	if err := db.Where("id = ?", id).Delete(Category{}).Error; err != nil {
		return err
	}

	return nil
}

// CleanAllCategory clear all category
func CleanAllCategory() error {
	if err := db.Unscoped().Where("delete_flag != ?", 0).Delete(&Category{}).Error; err != nil {
		return err
	}

	return nil
}
