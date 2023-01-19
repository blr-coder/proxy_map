package core

import (
	"proxy_map/core/core_errors"
	"proxy_map/core/models"
	"sync"
)

type ProxyMap struct {
	sync.RWMutex
	storage map[*models.ProxyRequest]*models.ProxyResponse
}

func NewProxyMap() *ProxyMap {
	return &ProxyMap{storage: make(map[*models.ProxyRequest]*models.ProxyResponse)}
}

func (m *ProxyMap) Save(k *models.ProxyRequest, v *models.ProxyResponse) error {
	m.Lock()
	m.storage[k] = v
	m.Unlock()
	return nil
}

func (m *ProxyMap) Get(k *models.ProxyRequest) (*models.ProxyResponse, error) {
	m.RLock()
	val, ok := m.storage[k]
	m.RUnlock()
	if !ok {
		return nil, core_errors.NewKeyNotFoundErr(k)
	}
	return val, nil
}

func (m *ProxyMap) All() (map[*models.ProxyRequest]*models.ProxyResponse, error) {
	return m.storage, nil
}
