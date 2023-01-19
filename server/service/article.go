package service

import (
	"github.com/longrl/mood/model/article"
	"github.com/longrl/mood/request"
	"github.com/longrl/mood/response"
)

type ArticleService struct {
}

// List 博客列表
func (s *ArticleService) List() []article.Article {
	return article.FindAll()
}

// Detail 博客详情
func (s *ArticleService) Detail(id string) *article.Article {
	return article.FindById(id)
}

// Top 博客置顶
func (s *ArticleService) Top(id string) bool {
	return article.UpdateTopById(id)
}

// Add 新增博客
func (s *ArticleService) Add(req *request.ArticleRequest) bool {
	a := &article.Article{
		Title:    req.Title,
		Top:      req.Top,
		Image:    req.Image,
		Content:  req.Content,
		MdPath:   req.MdPath,
		MusicUrl: req.MusicUrl,
		Category: req.Category,
	}
	return article.Add(a)
}

// Archive 博客归档
func (s *ArticleService) Archive() []response.ArchiveVo {
	vos := make([]response.ArchiveVo, 0)

	articles := article.FindAll()
	for _, article := range articles {
		vo := response.ArchiveVo{
			Title: article.Title,
			Date:  article.CreatedAt,
		}
		vos = append(vos, vo)
	}
	return vos
}

// Update 更新博客
func (s *ArticleService) Update(req *request.ArticleRequest) bool {
	a := &article.Article{
		Title:    req.Title,
		Top:      req.Top,
		Image:    req.Image,
		Content:  req.Content,
		MdPath:   req.MdPath,
		MusicUrl: req.MusicUrl,
		Category: req.Category,
	}

	a.ID = uint(req.Id)
	return article.Update(a)
}

func (s *ArticleService) Delete(id string) bool {
	return article.Delete(id)
}
