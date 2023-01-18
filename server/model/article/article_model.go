package article

import "github.com/longrl/mood/model"

type Article struct {
	model.BaseModel

	Top      int8   `json:"top"`
	Image    string `json:"image"`
	Content  string `json:"content"`
	MdPath   string `json:"md_path"`
	MusicUrl string `json:"music_url"`
	Category int    `json:"category"`
}
