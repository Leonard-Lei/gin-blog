package models

import "github.com/jinzhu/gorm"

type Message struct {
	Model

	Content  string `json:"content"`
	CreateBy int    `json:"create_by"`
	State    int    `json:"state"`
}

// ExistMessageByID checks if an message exists based on ID
func ExistMessageByID(id int) (bool, error) {
	var message Message
	err := db.Select("id").Where("id = ? AND delete_flag = ? ", id, 0).First(&message).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if message.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetMessageTotal gets the total number of message based on the constraints
func GetMessageTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Message{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetMessages gets a list of messages based on paging constraints
func GetMessages(pageNum int, pageSize int, maps interface{}) ([]*Message, error) {
	var message []*Message
	//err := db.Preload("Article").Where(maps).Offset(pageNum).Limit(pageSize).Find(&message).Error
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&message).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return message, nil
}

// GetMessage Get a single message based on ID
func GetMessage(id int) (*Message, error) {
	var message Message
	err := db.Where("id = ? AND delete_flag = ?", id, 0).First(&message).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	err = db.Model(&message).Related(&message.ID).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &message, nil
}

// EditMessage modify a single message
func EditMessage(id int, data interface{}) error {
	if err := db.Model(&Message{}).Where("id = ? AND delete_flag = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

//Add a Message
func AddMessage(data map[string]interface{}) error {
	message := Message{
		Content:  data["content"].(string),
		CreateBy: data["create_by"].(int),
		State:    data["state"].(int),
	}

	if err := db.Create(&message).Error; err != nil {
		return err
	}

	return nil
}

// DeleteMessage delete a single message
func DeleteMessage(id int) error {
	if err := db.Where("id = ?", id).Delete(Message{}).Error; err != nil {
		return err
	}

	return nil
}

// CleanAllMessage clear all Message
func CleanAllMessage() error {
	if err := db.Unscoped().Where("delete_flag != ?", 0).Delete(&Message{}).Error; err != nil {
		return err
	}

	return nil
}
