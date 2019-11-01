package messageHandler

import (
	"context"

	"github.com/liuyuexclusive/future.srv.basic/model"
	message "github.com/liuyuexclusive/future.srv.basic/proto/message"
	"github.com/liuyuexclusive/utils/dbutil"

	"github.com/ahmetb/go-linq"
	"github.com/jinzhu/gorm"
)

type Handler struct {
	message.MessageHandler
}

func (e *Handler) Send(ctx context.Context, req *message.SendRequest, rsp *message.Response) error {
	return dbutil.Open(func(db *gorm.DB) error {
		messageToList := make([]model.MessageTo, 0)
		if len(req.ToList) > 0 {
			for _, v := range req.ToList {
				messageToList = append(messageToList, model.MessageTo{To: v, Status: uint(message.ChangeStatusRequest_Unread)})
			}
		} else { // send to all users
			var users []model.User
			db.Select("name").Find(&users)
			for _, v := range users {
				messageToList = append(messageToList, model.MessageTo{To: v.Name, Status: uint(message.ChangeStatusRequest_Unread)})
			}
		}
		message := &model.Message{From: req.From, Title: req.Title, Content: req.Content, MessageToList: messageToList}
		db.Create(&message)
		return nil
	})
}
func (e *Handler) ChangeStatus(ctx context.Context, req *message.ChangeStatusRequest, rsp *message.Response) error {
	return dbutil.Open(func(db *gorm.DB) error {
		if req.Id != 0 {
			db.Model(model.MessageTo{}).Where("id=?", req.Id).Update(model.MessageTo{Status: uint(req.Status)})
		} else {
			db.Model(model.MessageTo{}).Where("`to`=?", req.To).Update(model.MessageTo{Status: uint(req.Status)})
		}
		return nil
	})
}

func (e *Handler) Init(ctx context.Context, req *message.InitRequest, rsp *message.InitResponse) error {
	return dbutil.Open(func(db *gorm.DB) error {
		var listAll []model.MessageTo
		db.Preload("Message").Where("`to`=?", req.To).Find(&listAll)
		rsp.To = req.To

		getList := func(status uint) []*message.InitResponse_Message {
			var result []*message.InitResponse_Message
			linq.From(listAll).Where(func(c interface{}) bool { return c.(model.MessageTo).Status == status }).Select(func(c interface{}) interface{} {
				x := c.(model.MessageTo)
				return &message.InitResponse_Message{Id: int64(x.ID), From: x.Message.From, Title: x.Message.Title}
			}).ToSlice(&result)
			return result
		}

		rsp.Unread = getList(uint(message.ChangeStatusRequest_Unread))
		rsp.Readed = getList(uint(message.ChangeStatusRequest_Readed))
		rsp.Trash = getList(uint(message.ChangeStatusRequest_Trash))

		return nil
	})
}

func (e *Handler) Get(ctx context.Context, req *message.GetRequest, rsp *message.GetResponse) error {
	return dbutil.Open(func(db *gorm.DB) error {
		var messageTo model.MessageTo
		db.Preload("Message").First(&messageTo)
		rsp.Id = int64(messageTo.ID)
		rsp.From = messageTo.Message.From
		rsp.Title = messageTo.Message.Title
		rsp.Content = messageTo.Message.Content
		return nil
	})
}
