package errs

import (
	"fmt"
	"proxy_map/Internal/domain/models"
)

type KeyNotFoundErr struct {
	key *models.ProxyRequest
}

func NewKeyNotFoundErr(key *models.ProxyRequest) *KeyNotFoundErr {
	return &KeyNotFoundErr{key: key}
}

func (e *KeyNotFoundErr) Error() string {
	return fmt.Sprintf("key %v not found", e.key)
}
