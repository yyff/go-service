package http

import (
	"net/http"
	"strconv"

	"github.com/yyff/go-service/post"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Get(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	post, err := postSvc.GetPost(c.Request.Context(), id)
	if err != nil {
		log.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if post == nil {
		log.Infof("can't get post by id[%v]", id)
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, post)
}

func Create(c *gin.Context) {
	post := &post.Post{}
	err := c.BindJSON(post)
	if err != nil {
		log.Warn(err)
		c.Status(http.StatusBadRequest)
		return
	}
	id, err := postSvc.CreatePost(c.Request.Context(), post)
	if err != nil {
		log.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, &map[string]uint64{"id": id})
}
