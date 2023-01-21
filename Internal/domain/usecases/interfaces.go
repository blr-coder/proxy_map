package usecases

import "proxy_map/Internal/domain/models"

type IProxyUseCase interface {
	Proxy(proxyRequest *models.ProxyRequest) error
	ProxyMap() (map[*models.ProxyRequest]*models.ProxyResponse, error)
}
