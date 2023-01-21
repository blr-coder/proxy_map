package usecases

import (
	"proxy_map/Internal/infrastructure/repository"
	"proxy_map/Internal/infrastructure/repository/map_store"
)

type ProxyUseCase struct {
	storage repository.IProxyRepository
}

func NewProxyUseCase(proxyMap *map_store.ProxyMap) *ProxyUseCase {
	return &ProxyUseCase{storage: proxyMap}
}
