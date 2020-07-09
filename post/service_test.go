package post

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type postDaoMockImpl struct{}

func (pd *postDaoMockImpl) GetPost(ctx context.Context, id uint64) (*Post, error) {
	if id == 0 {
		return nil, nil
	}
	return &Post{ID: id}, nil
}

var curID uint64 = 0

func (pd *postDaoMockImpl) CreatePost(ctx context.Context, p *Post) (uint64, error) {
	curID++
	return curID, nil
}

func TestServiceImpl_GetPost(t *testing.T) {
	svc := NewService(&postDaoMockImpl{})
	post, _ := svc.GetPost(context.Background(), 0)
	assert.Nil(t, post)
	post, _ = svc.GetPost(context.Background(), 1)
	assert.Equal(t, &Post{ID: 1}, post)
}

func TestServiceImpl_CreatePost(t *testing.T) {
	svc := NewService(&postDaoMockImpl{})
	id, err := svc.CreatePost(context.Background(), &Post{})
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), id)
	id, err = svc.CreatePost(context.Background(), &Post{})
	assert.Equal(t, uint64(2), id)
}
