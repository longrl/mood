package article

import "github.com/longrl/mood/config"

func FindById(id int) (a *Article) {
	config.DB.Model(&Article{}).Where("id = ?", id).First(a)
	return
}

func UpdateTopById(id int) bool {
	res := config.DB.Model(&Article{}).Where("id = ?", id).Update("top", 1)
	return res.RowsAffected > 0
}

func Add(article Article) bool {
	res := config.DB.Create(&article)
	return res.RowsAffected > 0
}

func FindAll() (articles []Article) {
	config.DB.Find(&articles)
	return
}
