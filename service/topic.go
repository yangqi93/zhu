package service

import (
	"time"
	"zhu/models"

	"github.com/gin-gonic/gin"
)

type Topic struct {
	Id          int32
	Name        string
	Title       string
	Description string
	Cover       string
	Url         string
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
			Name:        v.Name,
			Title:       v.Title,
			Description: v.Description,
			Cover:       v.Image,
			Url:         v.Url,
			CreatedAt:   time.Unix(v.CreatedAt, 0).Format("2006-01-02 15:04:05"),
		})
	}
	return topics, nil
}

func AddTopic(ctx *gin.Context, name, title, url, info string) error {

	_, err := models.Insert[models.Topic](ctx, models.GetDB(), &models.Topic{
		Name:        name,
		Title:       title,
		Url:         url,
		Description: info,
	})
	if err != nil {
		return err
	}

	return nil
}
