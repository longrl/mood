package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/longrl/mood/request"
	"github.com/longrl/mood/response"
	"github.com/longrl/mood/service"
)

// List 博客列表接口
func List(c *gin.Context) {
	s := service.ArticleService{}
	articles := s.List()
	response.Data(c, articles)
}

// Detail 博客详情接口
func Detail(c *gin.Context) {
	id := c.Param("id")
	s := service.ArticleService{}
	article := s.Detail(id)
	response.Data(c, article)
}

// UpdateDetail 更新博客内容接口
func UpdateDetail(c *gin.Context) {
	req := request.ArticleRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		response.Error(c, err, "传入参数格式有问题")
		return
	}
	s := service.ArticleService{}
	ok := s.Update(&req)
	if !ok {
		response.Error(c, errors.New("mood: 博客内容更新失败"))
		return
	}
	response.Success(c)
}

// Archive 博客归档接口
func Archive(c *gin.Context) {
	s := service.ArticleService{}
	vos := s.Archive()
	response.Data(c, vos)
}

// Top 博客置顶接口
func Top(c *gin.Context) {
	id := c.Param("id")
	s := service.ArticleService{}
	ok := s.Top(id)
	if !ok {
		response.Error(c, errors.New("mood: 置顶博客失败"))
		return
	}
	response.Success(c)
}

func GetTop(c *gin.Context) {
	s := service.ArticleService{}
	vo := s.GetTop()
	response.Data(c, vo)
}

// Add 新增博客接口
func Add(c *gin.Context) {
	req := request.ArticleRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		response.Error(c, err, "传入参数格式有问题")
		return
	}
	s := service.ArticleService{}
	ok := s.Add(&req)
	if !ok {
		response.Error(c, errors.New("mood: 博客内容新增失败"))
		return
	}
	response.Success(c)
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	s := service.ArticleService{}
	ok := s.Delete(id)
	if !ok {
		response.Error(c, errors.New("mood: 删除博客失败"))
		return
	}
	response.Success(c)
}
