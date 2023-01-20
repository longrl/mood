package article

import (
	"github.com/longrl/mood/config"
)

func FindById(id string) *Article {
	var a = &Article{}
	config.DB.Model(&Article{}).Where("id = ?", id).First(a)
	return a
}

func UpdateTopById(id string) bool {
	// 先将原先的文章取消置顶
	config.DB.Model(&Article{}).Where("top = ?", 1).Update("top", 0)
	res := config.DB.Model(&Article{}).Where("id = ?", id).Update("top", 1)
	return res.RowsAffected > 0
}

func Add(article *Article) bool {
	res := config.DB.Create(article)
	return res.RowsAffected > 0
}

func FindAll() (articles []Article) {
	config.DB.Find(&articles)
	return
}

func Update(article *Article) bool {
	res := config.DB.Updates(article)
	return res.RowsAffected > 0
}

func Delete(id string) bool {
	// 软删除
	res := config.DB.Where("id = ?", id).Delete(&Article{})
	return res.RowsAffected > 0
}

func FindByTop() *Article {
	var a = &Article{}
	config.DB.Model(&Article{}).Where("top = ?", 1).First(a)
	return a
}
