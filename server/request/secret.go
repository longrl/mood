package request

type SecretRequest struct {
	Key    string `json:"key"`
	Id     string `json:"id"`
	Answer string `json:"answer"`
}
