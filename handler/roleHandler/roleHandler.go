package roleHandler

import (
	"context"
	"errors"

	"github.com/liuyuexclusive/future.srv.basic/model"
	role "github.com/liuyuexclusive/future.srv.basic/proto/role"
	"github.com/liuyuexclusive/utils/db"

	"github.com/jinzhu/gorm"
)

type Handler struct {
	role.RoleHandler
}

func (e *Handler) Get(ctx context.Context, req *role.GetRequest, rsp *role.GetResponse) error {
	panic("not implemented")
}

func (e *Handler) AddOrUpdate(ctx context.Context, req *role.RoleAddOrUpdateRequest, rsp *role.Response) error {
	return db.Open(func(db *gorm.DB) error {
		if req.Id == 0 {
			db.Create(&model.Role{Name: req.Name})
		} else {
			var entity model.Role
			db.Where("id=?", req.Id).First(&entity)
			if entity.ID == 0 {
				return errors.New("无效的ID")
			}
			entity.Name = req.Name
			db.Save(&entity)
		}
		return nil
	})
}
