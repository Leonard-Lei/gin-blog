package auth_service

import "gin-blog/models"

type Auth struct {
	ID       int
	Username string
	Password string

	CreateBy   int
	CreateTime string
	UpdateBy   int
	State      int

	PageNum  int
	PageSize int
}

func (a *Auth) Check() (bool, error) {
	return models.CheckAuth(a.Username, a.Password)
}

func (a *Auth) Add() error {
	auth := map[string]interface{}{
		"username":  a.Username,
		"password":  a.Password,
		"create_by": a.CreateBy,
		"state":     a.State,
	}

	if err := models.AddAuth(auth); err != nil {
		return err
	}

	return nil
}

func (a *Auth) Edit() error {
	return models.EditAuth(a.ID, map[string]interface{}{
		"id":        a.ID,
		"username":  a.Username,
		"password":  a.Password,
		"state":     a.State,
		"update_by": a.UpdateBy,
	})
}

func (a *Auth) Get() (*models.Auth, error) {

	auth, err := models.GetAuth(a.ID)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func (a *Auth) GetAll() ([]*models.Auth, error) {

	auths, err := models.GetAuths(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}
	return auths, nil
}

func (a *Auth) Delete() error {
	return models.DeleteAuth(a.ID)
}

func (a *Auth) ExistByID() (bool, error) {
	return models.ExistAuthByID(a.ID)
}

func (a *Auth) Count() (int, error) {
	return models.GetAuthTotal(a.getMaps())
}

func (a *Auth) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["delete_flag"] = 0
	if a.State != -1 {
		maps["state"] = a.State
	}
	return maps
}
