package repository

import "proxy_map/Internal/domain/models"

type IProxyRepository interface {
	Save(k *models.ProxyRequest, v *models.ProxyResponse) error
	Get(k *models.ProxyRequest) (*models.ProxyResponse, error)
	All() (map[*models.ProxyRequest]*models.ProxyResponse, error)
}
