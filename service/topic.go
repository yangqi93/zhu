package service

import (
	"github.com/gin-gonic/gin"
	"time"
	"zhu/models"
)

type Topic struct {
	Id          int32
	Title       string
	Description string
	Cover       string
	CreatedAt   string
}

func GetTopicList(ctx *gin.Context, page, pageSize int) ([]*Topic, error) {
	offset := (page - 1) * pageSize
	list, err := models.GetList[models.Topic](ctx, models.GetDB(), map[string]interface{}{}, []string{"*"}, offset, pageSize, "id desc")
	if err != nil {
		return nil, err
	}
	var topics []*Topic
	for _, v := range list {
		topics = append(topics, &Topic{
			Id:          int32(v.ID),
			Title:       v.Name,
			Description: v.Description,
			Cover:       v.Image,
			CreatedAt:   time.Unix(v.CreatedAt, 0).Format("2006-01-02 15:04:05"),
		})
	}
	return topics, nil
}
