package secret

import "github.com/longrl/mood/config"

func IsKey(secret string) bool {
	var count int64
	config.DB.Model(Secret{}).Where("secret = ?", secret).Count(&count)
	return count > 0
}

func FindByKey(key string) (s Secret) {
	config.DB.Model(Secret{}).Where("secret = ?", key).First(&s)
	return
}
