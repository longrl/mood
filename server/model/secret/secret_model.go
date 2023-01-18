package secret

import "github.com/longrl/mood/model"

type Secret struct {
	model.BaseModel

	Secret string `json:"-"`
	Role   int8   `json:"-"`
}
