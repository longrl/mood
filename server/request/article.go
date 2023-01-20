package request

type ArticleRequest struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Top      string `json:"top"`
	Image    string `json:"image"`
	Content  string `json:"content"`
	MdPath   string `json:"md_path"`
	MusicUrl string `json:"music_url"`
	Category string `json:"category"`
}
