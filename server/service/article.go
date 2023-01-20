package service

import (
	"github.com/longrl/mood/model/article"
	"github.com/longrl/mood/request"
	"github.com/longrl/mood/response"
	"strconv"
)

type ArticleService struct {
}

// List 博客列表
func (s *ArticleService) List() []response.BlogVO {
	vos := make([]response.BlogVO, 0)
	articles := article.FindAll()
	for _, article := range articles {
		vo := response.BlogVO{
			Id:       strconv.Itoa(int(article.ID)),
			Title:    article.Title,
			Top:      strconv.Itoa(int(article.Top)),
			Image:    article.Image,
			Content:  article.Content,
			MdPath:   article.MdPath,
			MusicUrl: article.MusicUrl,
			Category: strconv.Itoa(article.Category),
			Date:     article.CreatedAt.Format("2006-01-02"),
		}
		vos = append(vos, vo)
	}
	return vos
}

// Detail 博客详情
func (s *ArticleService) Detail(id string) response.BlogVO {
	article := article.FindById(id)
	vo := response.BlogVO{
		Id:       strconv.Itoa(int(article.ID)),
		Title:    article.Title,
		Top:      strconv.Itoa(int(article.Top)),
		Image:    article.Image,
		Content:  article.Content,
		MdPath:   article.MdPath,
		MusicUrl: article.MusicUrl,
		Category: strconv.Itoa(article.Category),
		Date:     article.CreatedAt.Format("2006-01-02"),
	}
	return vo
}

// Top 博客置顶
func (s *ArticleService) Top(id string) bool {
	return article.UpdateTopById(id)
}

// Add 新增博客
func (s *ArticleService) Add(req *request.ArticleRequest) bool {
	category, _ := strconv.Atoi(req.Category)
	top, _ := strconv.Atoi(req.Top)
	a := &article.Article{
		Title:    req.Title,
		Top:      int8(top),
		Image:    req.Image,
		Content:  req.Content,
		MdPath:   req.MdPath,
		MusicUrl: req.MusicUrl,
		Category: category,
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
	category, _ := strconv.Atoi(req.Category)
	id, _ := strconv.Atoi(req.Id)
	top, _ := strconv.Atoi(req.Top)
	a := &article.Article{
		Title:    req.Title,
		Top:      int8(top),
		Image:    req.Image,
		Content:  req.Content,
		MdPath:   req.MdPath,
		MusicUrl: req.MusicUrl,
		Category: category,
	}

	a.ID = uint(id)
	return article.Update(a)
}

func (s *ArticleService) Delete(id string) bool {
	return article.Delete(id)
}

func (s *ArticleService) GetTop() response.BlogVO {
	article := article.FindByTop()
	vo := response.BlogVO{
		Id:       strconv.Itoa(int(article.ID)),
		Title:    article.Title,
		Top:      strconv.Itoa(int(article.Top)),
		Image:    article.Image,
		Content:  article.Content,
		MdPath:   article.MdPath,
		MusicUrl: article.MusicUrl,
		Category: strconv.Itoa(article.Category),
		Date:     article.CreatedAt.Format("2006-01-02"),
	}
	return vo
}
