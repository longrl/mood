package request

type ArticleRequest struct {
	Id       int
	Title    string `json:"title"`
	Top      int8   `json:"top"`
	Image    string `json:"image"`
	Content  string `json:"content"`
	MdPath   string `json:"md_path"`
	MusicUrl string `json:"music_url"`
	Category int    `json:"category"`
}
