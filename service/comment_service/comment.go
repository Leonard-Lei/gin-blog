package comment_service

import (
	"gin-blog/models"
)

type Comment struct {
	ID         int
	ArticleID  int
	Content    string
	State      int
	CreateBy   int
	CreateTime string
	UpdateBy   int

	PageNum  int
	PageSize int
}

func (a *Comment) Add() error {
	comment := map[string]interface{}{
		"article_id": a.ArticleID,
		"content":    a.Content,
		"create_by":  a.CreateBy,
		"state":      a.State,
	}

	if err := models.AddComment(comment); err != nil {
		return err
	}

	return nil
}

func (a *Comment) Edit() error {
	return models.EditComment(a.ID, map[string]interface{}{
		"id":        a.ID,
		"content":   a.Content,
		"state":     a.State,
		"update_by": a.UpdateBy,
	})
}

func (a *Comment) Get() (*models.Comment, error) {

	comment, err := models.GetComment(a.ID)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (a *Comment) GetAll() ([]*models.Comment, error) {

	comments, err := models.GetComments(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (a *Comment) Delete() error {
	return models.DeleteComment(a.ID)
}

func (a *Comment) ExistByID() (bool, error) {
	return models.ExistCommentByID(a.ID)
}

func (a *Comment) Count() (int, error) {
	return models.GetCommentTotal(a.getMaps())
}

func (a *Comment) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["delete_flag"] = 0
	if a.State != -1 {
		maps["state"] = a.State
	}
	return maps
}
