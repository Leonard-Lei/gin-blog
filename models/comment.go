package models

import "github.com/jinzhu/gorm"

type Comment struct {
	Model

	ArticleID int    `json:"article_id"`
	Content   string `json:"content"`
	CreateBy  int    `json:"create_by"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	State     int    `json:"state"`
}

// ExistCommentByID checks if an comment exists based on ID
func ExistCommentByID(id int) (bool, error) {
	var comment Comment
	err := db.Select("id").Where("id = ? AND delete_flag = ? ", id, 0).First(&comment).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if comment.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetCommentTotal gets the total number of comment based on the constraints
func GetCommentTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Comment{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetComments gets a list of comments based on paging constraints
func GetComments(pageNum int, pageSize int, maps interface{}) ([]*Comment, error) {
	var comment []*Comment
	//err := db.Preload("Article").Where(maps).Offset(pageNum).Limit(pageSize).Find(&comment).Error
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&comment).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return comment, nil
}

// GetComment Get a single comment based on ID
func GetComment(id int) (*Comment, error) {
	var comment Comment
	err := db.Where("id = ? AND delete_flag = ?", id, 0).First(&comment).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	err = db.Model(&comment).Related(&comment.ID).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &comment, nil
}

// EditComment modify a single comment
func EditComment(id int, data interface{}) error {
	if err := db.Model(&Comment{}).Where("id = ? AND delete_flag = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

//Add a Comment
func AddComment(data map[string]interface{}) error {
	comment := Comment{
		ArticleID: data["article_id"].(int),
		Content:   data["content"].(string),
		CreateBy:  data["create_by"].(int),
		Nickname:  data["nickname"].(string),
		Email:     data["email"].(string),
		State:     data["state"].(int),
	}

	if err := db.Create(&comment).Error; err != nil {
		return err
	}

	return nil
}

// DeleteComment delete a single comment
func DeleteComment(id int) error {
	if err := db.Where("id = ?", id).Delete(Comment{}).Error; err != nil {
		return err
	}

	return nil
}

// CleanAllComment clear all Comment
func CleanAllComment() error {
	if err := db.Unscoped().Where("delete_flag != ?", 0).Delete(&Comment{}).Error; err != nil {
		return err
	}

	return nil
}
