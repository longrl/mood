package service

import "github.com/longrl/mood/model/secret"

type SecretService struct {
}

func (s *SecretService) Check(key string) (bool, *secret.Secret) {
	ok := secret.IsKey(key)
	if !ok {
		return false, nil
	}
	return true, secret.FindByKey(key)
}
