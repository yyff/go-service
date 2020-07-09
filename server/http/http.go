package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yyff/go-service/conf"
	"github.com/yyff/go-service/dao"
	"github.com/yyff/go-service/post"
	_ "github.com/yyff/go-service/post"
)

var (
	postSvc post.Service
	// userSvc user.Service
)

func InitServices(d *dao.Dao) {
	postSvc = post.NewService(post.NewPostDao(d))
}

func Run(config *conf.Config) {
	// init the gin server by config
	router := gin.Default()
	router.GET("/post/:id", Get)
	router.POST("/post", Create)

	hConfig := config.Http
	http := &http.Server{
		Addr:         fmt.Sprintf(":%d", hConfig.Port),
		Handler:      router,
		ReadTimeout:  time.Duration(hConfig.ReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(hConfig.WriteTimeout) * time.Millisecond,
	}

	if err := http.ListenAndServe(); err != nil {
		panic(err)
	}
}
