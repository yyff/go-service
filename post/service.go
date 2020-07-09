package post

import (
	"context"
)

type Service interface {
	GetPost(context.Context, uint64) (*Post, error)
	CreatePost(context.Context, *Post) (uint64, error)
}

type ServiceImpl struct {
	dao PostDao
}

func NewService(postDao PostDao) *ServiceImpl {
	return &ServiceImpl{dao: postDao}
}

func (si *ServiceImpl) GetPost(ctx context.Context, id uint64) (*Post, error) {
	return si.dao.GetPost(ctx, id)
}
func (si *ServiceImpl) CreatePost(ctx context.Context, post *Post) (uint64, error) {
	return si.dao.CreatePost(ctx, post)
}
