package post

import (
	"context"

	"github.com/yyff/go-service/dao"
)

const (
	getPostSQL    = "SELECT * FROM post WHERE id=?"
	insertPostSQL = "INSERT INTO post(userid, title, content) VALUES(?, ?, ?)"
)

type PostDao interface {
	GetPost(context.Context, uint64) (*Post, error)
	CreatePost(context.Context, *Post) (uint64, error)
}

type postDaoImpl struct {
	d *dao.Dao
}

func NewPostDao(d *dao.Dao) *postDaoImpl {
	return &postDaoImpl{d}
}

func (pd *postDaoImpl) GetPost(ctx context.Context, id uint64) (*Post, error) {
	rows, err := pd.d.DB.Queryx(getPostSQL, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}
	post := &Post{}
	err = rows.StructScan(post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (pd *postDaoImpl) CreatePost(ctx context.Context, p *Post) (uint64, error) {
	res, err := pd.d.DB.ExecContext(ctx, insertPostSQL, p.UserID, p.Title, p.Content)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}
