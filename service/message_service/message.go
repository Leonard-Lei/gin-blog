package message_service

import (
	"gin-blog/models"
)

type Message struct {
	ID         int
	Content    string
	State      int
	CreateBy   int
	CreateTime string
	UpdateBy   int

	PageNum  int
	PageSize int
}

func (a *Message) Add() error {
	message := map[string]interface{}{
		"content":   a.Content,
		"create_by": a.CreateBy,
		"state":     a.State,
	}

	if err := models.AddMessage(message); err != nil {
		return err
	}

	return nil
}

func (a *Message) Edit() error {
	return models.EditMessage(a.ID, map[string]interface{}{
		"id":        a.ID,
		"content":   a.Content,
		"state":     a.State,
		"update_by": a.UpdateBy,
	})
}

func (a *Message) Get() (*models.Message, error) {

	message, err := models.GetMessage(a.ID)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (a *Message) GetAll() ([]*models.Message, error) {

	messages, err := models.GetMessages(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (a *Message) Delete() error {
	return models.DeleteMessage(a.ID)
}

func (a *Message) ExistByID() (bool, error) {
	return models.ExistMessageByID(a.ID)
}

func (a *Message) Count() (int, error) {
	return models.GetMessageTotal(a.getMaps())
}

func (a *Message) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["delete_flag"] = 0
	if a.State != -1 {
		maps["state"] = a.State
	}
	return maps
}
