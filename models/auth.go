package models

import "github.com/jinzhu/gorm"

type Auth struct {
	Model

	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	CreateBy int    `json:"create_by"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`

	State int `json:"state"`
}

// CheckAuth checks if authentication information exists
func CheckAuth(username, password string) (bool, error) {
	var auth Auth
	err := db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.ID > 0 {
		return true, nil
	}

	return false, nil
}

// ExistAuthByID checks if an auth exists based on ID
func ExistAuthByID(id int) (bool, error) {
	var auth Auth
	err := db.Select("id").Where("id = ? AND delete_flag = ? ", id, 0).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetAuthTotal gets the total number of auth based on the constraints
func GetAuthTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Auth{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetAuths gets a list of auths based on paging constraints
func GetAuths(pageNum int, pageSize int, maps interface{}) ([]*Auth, error) {
	var auth []*Auth
	//err := db.Preload("Article").Where(maps).Offset(pageNum).Limit(pageSize).Find(&auth).Error
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return auth, nil
}

// GetAuth Get a single auth based on ID
func GetAuth(id int) (*Auth, error) {
	var auth Auth
	err := db.Where("id = ? AND delete_flag = ?", id, 0).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	err = db.Model(&auth).Related(&auth.ID).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &auth, nil
}

// EditAuth modify a single auth
func EditAuth(id int, data interface{}) error {
	if err := db.Model(&Auth{}).Where("id = ? AND delete_flag = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

//Add a Auth
func AddAuth(data map[string]interface{}) error {
	auth := Auth{
		Username: data["username"].(string),
		Password: data["password"].(string),
		CreateBy: data["create_by"].(int),
		State:    data["state"].(int),
	}

	if err := db.Create(&auth).Error; err != nil {
		return err
	}

	return nil
}

// DeleteAuth delete a single auth
func DeleteAuth(id int) error {
	if err := db.Where("id = ?", id).Delete(Auth{}).Error; err != nil {
		return err
	}

	return nil
}

// CleanAllAuth clear all Auth
func CleanAllAuth() error {
	if err := db.Unscoped().Where("delete_flag != ?", 0).Delete(&Auth{}).Error; err != nil {
		return err
	}

	return nil
}
